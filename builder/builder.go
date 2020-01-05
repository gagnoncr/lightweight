package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"engine/engine"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var (
	registry  []string
	whiteList []string
	fileHash = make(map[string]string)
	wg sync.WaitGroup
)

func main() {
	doneChan := make(chan bool)
	go func(doneChan chan bool) {
		defer func() {
			doneChan <- true
		}()
		for {
			action := engine.Action()
			startup(fmt.Sprintf("./orche.sh build %s", action))
			time.Sleep(5 * time.Second)
			log.Println("done, starting new routine")
		}
	}(doneChan)
	<-doneChan
	wg.Wait()
}

func startup(action string) {
	log.Println("creating whitelist")
	createWhitelist()

	log.Println("building registry")
	buildRegistry()

	wg.Add(1)
	engine.DoEvery(5 * time.Second, verifyHashes, action)
	wg.Done()
}


func verifyHashes(t time.Time, action string) {
	log.Println("verifying hashes")
	for _, fn := range registry {
		oldHash := retrieveHash(fn)
		newHash := calculateHash(fn)
		if !(engine.CompareHash(oldHash, newHash) == 0) {
			log.Println("changed detected - updating hash, input required")
			update := engine.ShouldContinue()
			insertRecord(fn, newHash)
			if update == true {engine.RunAction(action)} else {continue}
		}
	}
}

func createWhitelist() {
	file, err := os.Open("../ignore")

	if err != nil {
		log.Println("no .ignore file found, race condition will ensue if jobs edit files -- will not create whitelist")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		log.Println(scanner.Text())
		whiteList = append(whiteList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateHash(absoluteFilePath string) string {
	f, err := os.Open(absoluteFilePath)
	engine.HandleErr(err)
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func insertRecord(absoluteFilePath string, hash string) {
	fileHash[absoluteFilePath] = hash
}

func retrieveHash(absoluteFilePath string) string {
	val, _ := fileHash[absoluteFilePath]
	return val
}

// we need to ignore .git
func recursiveDirectoryCrawl(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	engine.HandleErr(err)

	for _, f := range files {
		fileOrDir, err := os.Stat(dirName + "/" + f.Name())
		engine.HandleErr(err)
		switch mode := fileOrDir.Mode(); {
		case mode.IsDir():
			// keep looking for files
			if !(f.Name() == ".git") {
				recursiveDirectoryCrawl(dirName + "/" + f.Name())
			}

		case mode.IsRegular():
			// if the file is whitelisted, don't add it to the registry
			toAdd := true
			for _, whitelisted := range whiteList {
				if f.Name() == whitelisted {
					toAdd = false
				}
			}
			if toAdd {
				absolutePath := dirName + "/" + f.Name()
				registry = append(registry, absolutePath)
			}
		}
	}
}

func buildRegistry() {
	log.Println("starting directory scan")
	recursiveDirectoryCrawl("../")
	log.Println("computing hashes & creating map entries")
	for _, fn := range registry {
		hash := calculateHash(fn)
		insertRecord(fn, hash)
	}
}