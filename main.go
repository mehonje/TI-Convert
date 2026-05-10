package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ti_convert/convert"
)

func main() {
	fmt.Print("\n")
	fmt.Print("0: Convert 8xp -> txt\n")
	fmt.Print("1: Convert txt -> 8xp\n")
	fmt.Print("Idx? ")

	reader := bufio.NewReader(os.Stdin)
	option, err := reader.ReadString('\n') // Read input until newline
	print_error(err)
	option = strings.TrimSpace(option)

	switch option {
	case "0":
		convert_8xp_to_txt()
	case "1":
		convert_txt_to_8xp()
	}
}

func convert_8xp_to_txt() {
	fmt.Print("\nProgram file to convert: ") // Ask for input
	reader := bufio.NewReader(os.Stdin)
	from_path_input, err := reader.ReadString('\n') // Read input until newline
	print_error(err)
	fmt.Print("Path to output data to: ") // Ask for input
	reader = bufio.NewReader(os.Stdin)
	to_path_input, err := reader.ReadString('\n') // Read input until newline
	print_error(err)
	convert.Eightxp_to_txt(from_path_input, to_path_input)
}

func convert_txt_to_8xp() {
	fmt.Print("\nPath to .txt program file to convert: ") // Ask for input
	reader := bufio.NewReader(os.Stdin)
	from_path_input, err := reader.ReadString('\n') // Read input until newline
	print_error(err)
	fmt.Print("Path to output data to: ") // Ask for input
	reader = bufio.NewReader(os.Stdin)
	to_path_input, err := reader.ReadString('\n') // Read input until newline
	print_error(err)
	convert.Txt_to_eightxp(from_path_input, to_path_input)
}

func print_error(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
	}
}
