package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	route "api_micro/router"
	gears "api_micro/middelware"
)

func main() {
	port := os.Getenv("API_PORT")
	log.Printf("returned env port :[%s]", port)


	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)


			mSys := fmt.Sprintf("total bytes of memory obtained from the OS: %.4v|s",float64(m.Sys) / 1024 / 1024)
			mHeap := fmt.Sprintf("bytes of allocated heap objects: %.4v|s",float64(m.HeapAlloc) / 1024 / 1024)

			err := gears.WriteString("out.log",[]string{mSys, mHeap})
			if err != nil {
				log.Fatalf(err.Error())
			}

			time.Sleep(5 * time.Second)
		}
	}()

	if port == "" {
		defaultPort := "3030"
		log.Println("no env var set for port, defaulting to " + defaultPort)
		r := route.Router()
		log.Fatal(http.ListenAndServe(":"+defaultPort, r))

	} else {
		log.Println("starting server on port " + port)
		r := route.Router()
		log.Fatal(http.ListenAndServe(":"+port, r))
	}
}