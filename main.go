package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter API endpoint (or type 'exit' to quit): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", err)
			continue
		}

		input = input[:len(input)-1] // remove new line character

		if input == "exit" {
			break
		}

		resp, err := http.Get(input)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Response:")
			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		} else {
			fmt.Println("Request failed with status code: ", resp.StatusCode)
		}
	}
}

