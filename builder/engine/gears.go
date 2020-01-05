package engine

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func ShouldContinue() bool {

	var shouldContinue string
	fmt.Printf("\n\nChange detected do you want to deploy (Y/N): ")
	fmt.Scanln(&shouldContinue)

	return shouldContinue == "Y" || shouldContinue == "y"

}

func DoEvery(d time.Duration, f func(time.Time, string), action string) {
	for x := range time.Tick(d) {
		f(x, action)
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func CompareHash(old string, new string) int {
	return strings.Compare(old, new)
}

func ListCommands() string {
	return `
deployKube
	builds new images and deploys to local minikube

deployCompose
	builds new images and deploys with docker compose

default	
	empty string, just build new images

`
}

func Action() string {
	var action string
	fmt.Println(ListCommands())
	fmt.Printf("\n\nWould you like to use Kube or Compose ( deployKube || deployCompose ): ")
	fmt.Scanln(&action)
	switch action {
	case "deployKube":
		return action
	case "deployCompose":
		return action
	default:
		return ""
	}
}

func RunAction(action string) {
	log.Println(fmt.Sprintf("Now taking action, running: %v", action))

	cmd := exec.Command("/bin/sh", "-c", action)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		HandleErr(err)
	}
	log.Println(outb.String())
	log.Println(errb.String())
}
