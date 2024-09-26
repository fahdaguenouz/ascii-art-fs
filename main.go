package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		if len(args) == 1 {
			fmt.Println("please enter the string to print and the banner 'go run . [STRING] [BANNER]'")
		} else {
			fmt.Println("Please enter this format  'go run . [STRING] [BANNER]'")
		}
		return
	}

	input := args[0]
	banner := args[1]
	fileName:=""
	standard := "./art/standard.txt"
	shadow := "./art/shadow.txt"
	thinkertoy := "./art/thinkertoy.txt"

	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		if banner=="standard" {
			fileName = standard
		}else if banner == "shadow"{
			fileName = shadow
		}else if banner=="thinkertoy"{
			fileName = thinkertoy
		}
		// Open the ASCII art file
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		var asciiArt []string

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			asciiArt = append(asciiArt, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
			return
		}

		
		var result string
		lines := strings.Split(input, "\\n")

		for _, line := range lines {
			if line == "" {
				result += "\n"
				continue
			}
			// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
			for i := 1; i <= 8; i++ {
				for _, r := range line {
					// Ensure the character is within the valid ASCII range
					if r < 32 || r > 126 {
						fmt.Println("Please enter a valide character between ascii code 32 and 126")
						return
					}
					index := 9*(int(r)-32) + i 
					result += asciiArt[index]
				}
				result += "\n" // Add newline after finishing the current row of the line
			}
		}
			fmt.Print(result)
	}else{
		fmt.Println("Invalid banner. Please choose from: standard, shadow, thinkertoy")
        return
	}
}
