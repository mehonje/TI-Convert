package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"ti_convert/convert"
)

func main() {
	fmt.Print(`
0: Convert 8xp -> txt
1: Convert txt -> 8xp
Idx? `)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input: ", err)
		}
		return
	}
	option := strings.TrimSpace(scanner.Text())

	switch option {
	case "0":
		convert_8xp_to_txt()
	case "1":
		convert_txt_to_8xp()
	}
}

func convert_8xp_to_txt() {
	fmt.Print("\nProgram file to convert: ") // Ask for input
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input: ", err)
		}
		return
	}
	from_path_input := strings.TrimSpace(scanner.Text())
	fmt.Print("Path to output data to: ") // Ask for input
	scanner = bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input: ", err)
		}
		return
	}
	to_path_input := strings.TrimSpace(scanner.Text())
	convert.Eightxp_to_txt(from_path_input, to_path_input)
}

func convert_txt_to_8xp() {
	fmt.Print("\nPath to .txt program file to convert: ") // Ask for input
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input: ", err)
		}
		return
	}
	from_path_input := strings.TrimSpace(scanner.Text())
	fmt.Print("Path to output data to: ") // Ask for input
	scanner = bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal("Error reading input: ", err)
		}
		return
	}
	to_path_input := strings.TrimSpace(scanner.Text())
	convert.Txt_to_eightxp(from_path_input, to_path_input)
}

func print_error(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
	}
}
