package convert

import (
	"encoding/hex"
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
	0x68: "Eng",         //
	0xf5: "ExpReg",      //
	0x3a: ".",           //
	0x41: "A",           //
	0x42: "B",           //
	0x43: "C",           //
	0x44: "D",           //
	0x45: "E",           //
	0x46: "F",           //
	0x47: "G",           //
	0x48: "H",           //
	0x49: "I",           //
	0x4a: "J",           //
	0x4b: "K",           //
	0x4c: "L",           //
	0x4d: "M",           //
	0x4e: "N",           //
	0x4f: "O",           //
	0x50: "P",           //
	0x51: "Q",           //
	0x52: "R",           //
	0x53: "S",           //
	0x54: "T",           //
	0x55: "U",           //
	0x56: "V",           //
	0x57: "W",           //
	0x58: "X",           //
	0x59: "Y",           //
	0x5a: "Z",           //
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
	0x2e: "l",              //
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

var reverse_tokens = map[string][]byte{
	"det(":           {179},
	"ClrDraw":        {133},
	"Shade(":         {164},
	"A":              {65},
	"!=":             {111},
	">":              {108},
	",":              {43},
	"C":              {67},
	"sqrt(":          {188},
	"pi":             {172},
	"ln(":            {190},
	"fMin(":          {39},
	"Degree":         {101},
	"M":              {77},
	"Z":              {90},
	"If ":            {206},
	"Else":           {208},
	":":              {62},
	"W":              {87},
	"B":              {66},
	"U":              {85},
	"Repeat ":        {210},
	"Menu(":          {230},
	"Get(":           {232},
	"Tangent(":       {167},
	"tan(":           {198},
	"/":              {131},
	"(":              {16},
	" or ":           {60},
	"^2":             {13},
	"CubicReg":       {46},
	"e^(":            {191},
	"Q":              {81},
	"<":              {110},
	"<=":             {109},
	"Input ":         {220},
	"Circle(":        {165},
	"Pt-On(":         {158},
	"cosh(":          {202},
	"\"":             {42},
	"Lbl ":           {214},
	"->":             {4},
	"Ans":            {114},
	"K":              {75},
	"N":              {78},
	"O":              {79},
	"Dispgraph":      {223},
	"Pxl-On(":        {161},
	"Pxl-Change(":    {163},
	".":              {58},
	"Disp ":          {222},
	"theta":          {91},
	"ExpReg":         {245},
	"E":              {69},
	"\n":             {63},
	"augment(":       {20},
	"DependAuto":     {124},
	"I":              {73},
	"L":              {76},
	"Then":           {207},
	"ClrTable":       {251},
	"arccos(":        {197},
	"arccosh(":       {203},
	"P":              {80},
	"log(":           {192},
	"arcsin(":        {195},
	"Horizontal ":    {166},
	"solve(":         {34},
	"D":              {68},
	"F":              {70},
	"+":              {112},
	"Stop":           {217},
	"Pt-Change(":     {160},
	"=":              {106},
	">=":             {107},
	"arctan(":        {199},
	"nDeriv(":        {37},
	"BoxPlot":        {5},
	"Pause ":         {216},
	"fMax(":          {40},
	"fnInt(":         {36},
	"S":              {83},
	"V":              {86},
	")":              {17},
	"!":              {45},
	"Goto ":          {215},
	"prgm":           {95},
	">DMS":           {1},
	"While ":         {209},
	"Pt-Off(":        {159},
	"dim(":           {181},
	"-":              {113},
	"Output(":        {224},
	"Pxl-Off(":       {162},
	"RecallPic ":     {153},
	"^":              {240},
	"}":              {9},
	"]":              {7},
	">Frac":          {3},
	"DispTable":      {229},
	"getKey":         {173},
	"DrawInv ":       {168},
	"StorePic ":      {152},
	"RecallGDB ":     {155},
	"sin(":           {194},
	"abs(":           {178},
	"G":              {71},
	"cos(":           {196},
	"[":              {6},
	">Dec":           {2},
	"^3":             {15},
	"T":              {84},
	"End":            {212},
	" xor ":          {61},
	"^-1":            {12},
	"Eng":            {104},
	"H":              {72},
	"J":              {74},
	"R":              {82},
	"Y":              {89},
	"Return":         {213},
	"Text(":          {147},
	"not(":           {184},
	"IS>(":           {218},
	"DrawF ":         {169},
	" and ":          {64},
	"StoreGDB ":      {154},
	"X":              {88},
	"Wait ":          {239},
	"Prompt":         {221},
	"ClrHome":        {225},
	"pxl-Test(":      {19},
	"DependAsk":      {125},
	" ":              {41},
	"For(":           {211},
	"Send(":          {231},
	"?":              {175},
	"ClrList ":       {250},
	"*":              {130},
	"DS<(":           {219},
	"Line(":          {156},
	"Vertical ":      {157},
	"{":              {8},
	"i":              {44},
	"GraphStyle(":    {0xbb, 69},
	"String->Equ(":   {0xbb, 86},
	"ANOVA(":         {0xbb, 89},
	"Archive ":       {0xbb, 104},
	"x^2-Test(":      {0xbb, 64},
	"ClrAllLists":    {0xbb, 82},
	"DiagnosticOn":   {0xbb, 102},
	"Equ>String(":    {0xbb, 85},
	"a+bi":           {0xbb, 79},
	"conj(":          {0xbb, 37},
	"dbd(":           {0xbb, 7},
	"e":              {0xbb, 49},
	"DelVar ":        {0xbb, 84},
	"angle(":         {0xbb, 40},
	"binomcdf(":      {0xbb, 22},
	"binompdf(":      {0xbb, 21},
	"x^2cdf(":        {0xbb, 19},
	"x^pdf(":         {0xbb, 29},
	"cumSum(":        {0xbb, 41},
	">Eff(":          {0xbb, 6},
	"expr(":          {0xbb, 42},
	"bal(":           {0xbb, 2},
	"Clear Entries":  {0xbb, 87},
	"DiagnosticOff":  {0xbb, 103},
	"ExprOff":        {0xbb, 81},
	"ExprOn":         {0xbb, 80},
	"GREEN":          {0xef, 69},
	"WHITE":          {0xef, 75},
	"DARKGRAY":       {0xef, 79},
	"AUTO":           {0xef, 59},
	"CENTER":         {0xef, 147},
	"checkTmr(":      {0xef, 2},
	"GraphColor(":    {0xef, 101},
	"ORANGE":         {0xef, 70},
	"l":              {0xef, 46},
	"Sigma(":         {0xef, 51},
	"logBASE(":       {0xef, 52},
	"Dot-Thin":       {0xef, 117},
	"BLUE":           {0xef, 65},
	"BLACK":          {0xef, 67},
	"MAGENTA":        {0xef, 68},
	"GRAY":           {0xef, 78},
	"CLASSIC":        {0xef, 56},
	"RED":            {0xef, 66},
	"NAVY":           {0xef, 72},
	"x^2GOF-Test(":   {0xef, 20},
	"DEC":            {0xef, 60},
	"LTGRAY":         {0xef, 76},
	"BackgroundOff ": {0xef, 100},
	"DetectAsymOff":  {0xef, 107},
	"toString(":      {0xef, 151},
	"LTBLUE":         {0xef, 73},
	"MEDGRAY":        {0xef, 77},
	"TextColor(":     {0xef, 103},
	"dayOfWk(":       {0xef, 6},
	"DetectAsymOn":   {0xef, 106},
	"OpenLib(":       {0xef, 17},
	"eval(":          {0xef, 152},
	"BROWN":          {0xef, 71},
	"YELLOW":         {0xef, 74},
	"BorderColor":    {0xef, 108},
	"ClockOff":       {0xef, 15},
	"ClockOn":        {0xef, 16},
	"ExecLib":        {0xef, 18},
	"BackgroundOn ":  {0xef, 91},
	"piecewise(":     {0xef, 166},
	"Ymin":           {0x63, 12},
	"Ymax":           {0x63, 13},
	"Xres":           {0x63, 54},
	"deltaY":         {0x63, 39},
	"TraceStep":      {0x63, 56},
	"Xmax":           {0x63, 11},
	"Yscl":           {0x63, 3},
	"deltaX":         {0x63, 38},
	"XFact":          {0x63, 40},
	"Yfact":          {0x63, 41},
	"Xmin":           {0x63, 10},
	"Xscl":           {0x63, 2},
	"L4":             {0x5d, 3},
	"L5":             {0x5d, 4},
	"L6":             {0x5d, 5},
	"L1":             {0x5d, 0},
	"L2":             {0x5d, 1},
	"L3":             {0x5d, 2},
	"AxesOff":        {0x7e, 9},
	"AxesOn":         {0x7e, 8},
	"CoordOff":       {0x7e, 5},
	"CoordOn":        {0x7e, 4},
	"Dot-Thick":      {0x7e, 7},
}

func Eightxp_to_txt(from_path string, to_path string) {
	from_path = strings.TrimSpace(from_path)   // Remove whitespace
	if !strings.HasSuffix(from_path, ".8xp") { // If file path doesn't have ".8xp" suffix,
		from_path = from_path + ".8xp" // Append it
	}

	to_path = strings.TrimSpace(to_path)          // Remove whitespace
	to_path = strings.TrimSuffix(to_path, ".txt") // Remove ".txt" suffix

	var program_metadata [4]string
	var program_data []byte
	byte_data, err := os.ReadFile(from_path) // Read file data
	if err != nil {
		log.Fatal("Failed to read file data: ", err)
	}
	if len(byte_data) > 76 { // If data is more than 76 bytes long,
		program_metadata[0] = string(byte_data[60:67])                  // Store bytes 60 - 67 (program name)
		program_metadata[1] = string(byte_data[11:52])                  // Store bytes 11 - 52 (transmission comment)
		program_metadata[2] = hex.EncodeToString([]byte{byte_data[59]}) // Store byte 59 (type id)
		program_metadata[3] = hex.EncodeToString([]byte{byte_data[69]}) // Store bytes 69 (flag)
		program_data = byte_data[74 : len(byte_data)-2]                 // Store bytes 74 - end-2 (program), remove the first 74 bytes (program metadata) and last 2 bytes (checksum)
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
			//builder.WriteString(" (" + strconv.FormatInt(int64(val), 16) + ") ") // Uncomment to see hex equivalent
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
		_, err = file.WriteString(program_metadata[i])
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

	longest_command_length := 0
	for key := range reverse_tokens {
		if len(key) > longest_command_length {
			longest_command_length = len(key)
		}
	}

	var program_byte_data []byte
	i := 0
	n := longest_command_length
	for i < len(byte_data) {
		if i+n > len(byte_data) {
			n = len(byte_data) - i
		}
		if n == -1 {
			log.Fatal("Unknown command at position ", i)
		}
		command_bytes := byte_data[i : i+n]
		arr, ok := reverse_tokens[string(command_bytes)]
		if ok {
			for j := 0; j < len(arr); j++ {
				fmt.Println(strconv.FormatInt(int64(arr[j]), 16))
				program_byte_data = append(program_byte_data, arr[j])
			}
			i += n
			n = longest_command_length
		} else {
			n -= 1
		}
	}
	fmt.Println(program_byte_data)
}
