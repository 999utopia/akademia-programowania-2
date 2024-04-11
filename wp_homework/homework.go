// This program prompts the user for a number, 
// rounds it to the nearest multiple of ten, 
// generates a random key of the length of the rounded number, 
// displays a countdown along with key generation, and then writes the rounded number, 
// generated key, and countdown to a text file named "result.txt".

package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number < 1 {
		err := errors.New("Number must be greater than 0")
		fmt.Println("Error:", err)
		return
	}

	roundedNumber := int(math.Ceil(float64(number) / 10.0) * 10.0)

	countdown := ""
	for i := number; i > 0; i-- {
		countdown += fmt.Sprintf("%d... ", i)
		fmt.Println(i, "...")
		time.Sleep(time.Second)
	}
	fmt.Println("0...")

	key := make([]byte, roundedNumber)
	_, err = rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	keyStr := string(key)
	text := fmt.Sprintf("Rounded number: %d\nGenerated key: %s\nCountdown: %s", roundedNumber, keyStr, countdown)
	err = ioutil.WriteFile("result.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Program finished. Results saved in result.txt")
}
