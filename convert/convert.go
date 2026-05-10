package convert

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var tokens = map[byte]string{ // Normal tokens
	0x3f: "\n",          // Newline
	0x82: "*",           // Multiplication
	0x83: "/",           // Division
	0x70: "+",           // Addition
	0x71: "-",           // Subtraction
	0x10: "(",           // Left Parenthesis
	0x11: ")",           // Right Parenthesis
	0x29: " ",           // Space
	0x2a: "\"",          // Quotation mark
	0x2d: "!",           // Factorial
	0xce: "If ",         // If
	0xcf: "Then",        // Then
	0xd0: "Else",        // Else
	0xd3: "For(",        // For(
	0xd1: "While ",      // While
	0xd2: "Repeat ",     // Repeat
	0xd4: "End",         // End
	0xd8: "Pause ",      // Pause
	0xd6: "Lbl ",        // Label
	0xd7: "Goto ",       // Goto
	0xef: "Wait ",       // Wait
	0xda: "IS>(",        // Increment and skip if greater than
	0xdb: "DS<(",        // Decrement and skip if less than
	0xe6: "Menu(",       // Menu
	0x5f: "prgm",        // Program
	0xd5: "Return",      // Return
	0xd9: "Stop",        // Stop
	0xdc: "Input ",      //
	0xdd: "Prompt",      //
	0xde: "Disp ",       //
	0xdf: "Dispgraph",   //
	0xe5: "DispTable",   //
	0xe0: "Output(",     //
	0xad: "getKey",      //
	0xe1: "ClrHome",     // Clear home
	0xfb: "ClrTable",    // Clear table
	0xe8: "Get(",        //
	0xe7: "Send(",       //
	0x85: "ClrDraw",     //
	0x9c: "Line(",       //
	0xa6: "Horizontal ", //
	0x9d: "Vertical ",   //
	0xa7: "Tangent(",    //
	0xa9: "DrawF ",      //
	0xa4: "Shade(",      //
	0xa8: "DrawInv ",    //
	0xa5: "Circle(",     //
	0x93: "Text(",       //
	0x9e: "Pt-On(",      //
	0x9f: "Pt-Off(",     //
	0xa0: "Pt-Change(",  //
	0xa1: "Pxl-On(",     //
	0xa2: "Pxl-Off(",    //
	0xa3: "Pxl-Change(", //
	0x13: "pxl-Test(",   //
	0x98: "StorePic ",   //
	0x99: "RecallPic ",  //
	0x9a: "StoreGDB ",   //
	0x9b: "RecallGDB ",  //
	0x6a: "=",           // Equal
	0x6f: "!=",          // Not equal
	0x6c: ">",           //
	0x6e: "<",           //
	0x6b: ">=",          //
	0x6d: "<=",          //
	0x40: " and ",       //
	0x3c: " or ",        //
	0x3d: " xor ",       //
	0xb8: "not(",        //
	0xc2: "sin(",        //
	0xc4: "cos(",        //
	0xc6: "tan(",        //
	0xf0: "^",           //
	0x0d: "^2",          //
	0x0c: "^-1",         //
	0xbc: "sqrt(",       //
	0xac: "pi",          //
	0x08: "{",           //
	0x09: "}",           //
	0x06: "[",           //
	0x07: "]",           //
	0x5b: "theta",       //
	0x2c: "i",           //
	0xaf: "?",           //
	0x04: "->",          //
	0xbe: "ln(",         //
	0xc0: "log(",        //
	0xc3: "arcsin(",     //
	0xc5: "arccos(",     //
	0xc7: "arctan(",     //
	0x2b: ",",           //
	0x3e: ":",           //
	0x03: ">Frac",       //
	0x02: ">Dec",        //
	0x0f: "^3",          //
	0x27: "fMin(",       //
	0x28: "fMax(",       //
	0x25: "nDeriv(",     //
	0x24: "fnInt(",      //
	0x22: "solve(",      //
	0xb5: "dim(",        //
	0xb2: "abs(",        //
	0x72: "Ans",         //
	0x14: "augment(",    //
	0x05: "BoxPlot",     //
	0xfa: "ClrList ",    //
	0xca: "cosh(",       //
	0xcb: "arccosh(",    //
	0x2e: "CubicReg",    //
	0x65: "Degree",      //
	0x7d: "DependAsk",   //
	0x7c: "DependAuto",  //
	0xb3: "det(",        //
	0x01: ">DMS",        //
	0xbf: "e^(",         //
	0x3b: "E",           //
	0x68: "Eng",         //
	0xf5: "ExpReg",      //
	0x3a: ".",           //
}

var tokens_bb = map[byte]string{ // 2-byte tokens
	0x45: "GraphStyle(",   //
	0x54: "DelVar ",       //
	0x2a: "expr(",         //
	0x56: "String->Equ(",  //
	0x4f: "a+bi",          //
	0x28: "angle(",        //
	0x59: "ANOVA(",        //
	0x68: "Archive ",      //
	0x02: "bal(",          //
	0x16: "binomcdf(",     //
	0x15: "binompdf(",     //
	0x13: "x^2cdf(",       //
	0x1d: "x^pdf(",        //
	0x40: "x^2-Test(",     //
	0x57: "Clear Entries", //
	0x52: "ClrAllLists",   //
	0x25: "conj(",         //
	0x29: "cumSum(",       //
	0x07: "dbd(",          //
	0x67: "DiagnosticOff", //
	0x66: "DiagnosticOn",  //
	0x31: "e",             //
	0x06: ">Eff(",         //
	0x55: "Equ>String(",   //
	0x51: "ExprOff",       //
	0x50: "ExprOn",        //
}

var tokens_ef = map[byte]string{
	0x65: "GraphColor(",    //
	0x11: "OpenLib(",       //
	0x12: "ExecLib",        //
	0x98: "eval(",          //
	0x97: "toString(",      //
	0x41: "BLUE",           // Blue
	0x42: "RED",            // Red
	0x43: "BLACK",          // Black
	0x44: "MAGENTA",        // Magenta
	0x45: "GREEN",          // Green
	0x46: "ORANGE",         // Orange
	0x47: "BROWN",          // Brown
	0x48: "NAVY",           // Navy
	0x49: "LTBLUE",         // Light blue
	0x4a: "YELLOW",         // Yellow
	0x4b: "WHITE",          // White
	0x4c: "LTGRAY",         // Light grey
	0x4d: "MEDGRAY",        // Medium grey
	0x4e: "GRAY",           // Grey
	0x4f: "DARKGRAY",       // Dark grey
	0x67: "TextColor(",     //
	0x5b: "BackgroundOn ",  //
	0x64: "BackgroundOff ", //
	0x2e: "L",              //
	0x33: "Sigma(",         //
	0x34: "logBASE(",       //
	0xa6: "piecewise(",     //
	0x3B: "AUTO",           //
	0x6c: "BorderColor",    //
	0x93: "CENTER",         //
	0x02: "checkTmr(",      //
	0x14: "x^2GOF-Test(",   //
	0x38: "CLASSIC",        //
	0x0f: "ClockOff",       //
	0x10: "ClockOn",        //
	0x06: "dayOfWk(",       //
	0x3c: "DEC",            //
	0x6b: "DetectAsymOff",  //
	0x6a: "DetectAsymOn",   //
	0x75: "Dot-Thin",       //
}

var tokens_63 = map[byte]string{
	0x0a: "Xmin",      //
	0x0b: "Xmax",      //
	0x02: "Xscl",      //
	0x0c: "Ymin",      //
	0x0d: "Ymax",      //
	0x03: "Yscl",      //
	0x36: "Xres",      //
	0x26: "deltaX",    //
	0x27: "deltaY",    //
	0x28: "XFact",     //
	0x29: "Yfact",     //
	0x38: "TraceStep", //
}

var tokens_5d = map[byte]string{
	0x00: "L1", //
	0x01: "L2", //
	0x02: "L3", //
	0x03: "L4", //
	0x04: "L5", //
	0x05: "L6", //
}

var tokens_7e = map[byte]string{
	0x09: "AxesOff",   //
	0x08: "AxesOn",    //
	0x05: "CoordOff",  //
	0x04: "CoordOn",   //
	0x07: "Dot-Thick", //
}

func Eightxp_to_txt(from_path string, to_path string) {
	from_path = strings.TrimSpace(from_path)   // Remove whitespace
	if !strings.HasSuffix(from_path, ".8xp") { // If file path doesn't have ".8xp" suffix,
		from_path = from_path + ".8xp" // Append it
	}

	to_path = strings.TrimSpace(to_path)          // Remove whitespace
	to_path = strings.TrimSuffix(to_path, ".txt") // Remove ".txt" suffix

	var program_metadata [4][]byte
	var program_data []byte
	byte_data, err := os.ReadFile(from_path) // Read file data
	if err != nil {
		log.Fatal("Failed to read file data: ", err)
	}
	if len(byte_data) > 76 { // If data is more than 76 bytes long,
		program_metadata[0] = byte_data[60:67]          // Store bytes 60-67 (program name)
		program_metadata[1] = byte_data[11:52]          // Store bytes 11-52 (transmission comment)
		program_metadata[2] = []byte{byte_data[59]}     // Store byte 59 (type id)
		program_metadata[3] = []byte{byte_data[69]}     // Store bytes 69 (flag)
		program_data = byte_data[74 : len(byte_data)-2] // Remove the first 74 bytes (program metadata) and last 2 bytes (checksum)
	}

	var builder strings.Builder
	for i := 0; i < len(program_data); {
		val := program_data[i]
		var next_val byte

		step := 1
		if i+1 < len(program_data) {
			next_val = program_data[i+1]

			switch val {
			case 0xbb:
				s, ok := tokens_bb[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			case 0xef:
				s, ok := tokens_ef[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString("Wait ") // Add "Wait" command (0xef) if no
				}
			case 0x6e:
				s, ok := tokens_63[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			case 0x5d:
				s, ok := tokens_5d[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString("/") // Add division operator (0x5d) if no
				}
			case 0x7e:
				s, ok := tokens_7e[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			default:
				_, ok := tokens[val] // Check if mapping exists
				if ok {
					builder.WriteString(tokens[val]) // Replace if yes,
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			}
		} else { // 1-byte token
			_, ok := tokens[val] // Check if mapping exists
			if ok {
				builder.WriteString(tokens[val]) // Replace if yes,
			} else {
				builder.WriteString(string(val)) // Turn into string if no
			}
		}
		{
			builder.WriteString(" (" + strconv.FormatInt(int64(val), 16) + ") ") // Uncomment to see hex equivalent
		}
		i += step
	}

	dir := filepath.Dir(to_path)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal("Failed to create directory ", filepath.Dir(to_path), ": ", err)
	}
	err = os.WriteFile(to_path+".txt", []byte(builder.String()), 0644)
	if err != nil {
		log.Fatal("Failed to create", to_path+".txt: ", err)
	}
	file, err := os.OpenFile(to_path+".meta", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to create ", to_path+".meta: ", err)
	}
	defer file.Close()
	for i := 0; i < 4; i++ {
		_, err = file.Write(program_metadata[i])
		if err != nil {
			log.Fatal("Failed to store program metadata: ", err)
		}
		_, err = file.Write([]byte{0x0a})
		if err != nil {
			log.Fatal("Failed to store program metadata: ", err)
		}
	}
}

func Txt_to_eightxp(from_path string, to_path string) {
	from_path = strings.TrimSpace(from_path)   // Remove whitespace
	if !strings.HasSuffix(from_path, ".txt") { // If file path doesn't have ".txt" suffix,
		from_path = from_path + ".txt" // Append it
	}

	to_path = strings.TrimSpace(to_path)     // Remove whitespace
	if !strings.HasSuffix(to_path, ".8xp") { // If file path doesn't have ".8xp" suffix,
		to_path = to_path + "8xp" // Append it
	}

	byte_data, err := os.ReadFile(from_path) // Read file data
	if err != nil {
		log.Fatal("Failed to read program data: ", err)
	}

	fmt.Println(byte_data)
}
