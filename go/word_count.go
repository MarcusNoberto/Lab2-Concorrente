package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// Count the number of words in `fileContent`.
func wc(fileContent string) int {
	words := strings.Fields(fileContent)
	return len(words)
}

// Count the number of words in the file at `filePath`.
func wc_file(filePath string, wg *sync.WaitGroup, wordCount chan int) {
	defer wg.Done()
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	wordCount <- wc(string(fileContent))
}

// Count the number of words in all files directly within `directoryPath`.
// Files in subdirectories are not considered.
func wc_dir(directoryPath string, wg *sync.WaitGroup, wordCount chan int) {
	defer wg.Done()
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := directoryPath + "/" + file.Name()
			wg.Add(1)
			go wc_file(filePath, wg, wordCount)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " root_directory_path")
		return
	}
	rootPath := os.Args[1]

	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wordCount := make(chan int)

	for _, file := range files {
		if file.IsDir() {
			directoryPath := rootPath + "/" + file.Name()
			wg.Add(1)
			go wc_dir(directoryPath, &wg, wordCount)
		}
	}

	go func() {
		wg.Wait()
		close(wordCount)
	}()

	numberOfWords := 0

	for count := range wordCount {
		numberOfWords += count
	}

	fmt.Println(numberOfWords)
}

