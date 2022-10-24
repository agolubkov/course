package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	_, err := TouchFile("./go.mod")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	fmt.Println("done")
}

func TouchFile(f string) ([]byte, error) {
	file, err := os.OpenFile(f, 0, 0o666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return nil, nil
}
