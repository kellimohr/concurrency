package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	done := make(chan bool)

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir("./files")

	for _, file := range files {
		fileArray, err := ioutil.ReadFile(file.Name())

		if err != nil {
			log.Fatal(err)
		}

		// Send to count function
		s := string(fileArray[:])

		{
			go WordCount(s, done)
		}
	}

}

// WordCount counts the number of words in a string
func WordCount(s string, done chan bool) {
	fmt.Printf("Word Count=%d \n", len(s))
}
