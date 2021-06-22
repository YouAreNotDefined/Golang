package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// display file list
	err := filepath.Walk("root", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("path: %#v\n", path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	// server listening
	port := "8080"
	http.Handle("/", http.FileServer(http.Dir("root/")))
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}
