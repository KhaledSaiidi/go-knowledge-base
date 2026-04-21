package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func processData(data []int) {
	start := time.Now()
	defer func() {
		fmt.Println(
			"Data processing completed in ",
			time.Since(start),
		)
	}()
	for _, d := range data {
		fmt.Printf("Processing data: %d\n", d)
		time.Sleep(time.Millisecond * 100)
	}
}

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went wrong!")
}

func main() {
	err := readFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := []int{1, 2, 3, 4, 5}
	processData(data)

	safeOperation()
}
