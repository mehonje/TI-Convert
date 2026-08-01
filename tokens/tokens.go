package tokens

var Tokens = map[byte]string{ // Normal tokens
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
	0x30: "0",           //
	0x31: "1",           //
	0x32: "2",           //
	0x33: "3",           //
	0x34: "4",           //
	0x35: "5",           //
	0x36: "6",           //
	0x37: "7",           //
	0x38: "8",           //
	0x39: "9",           //
	0x0a: "getTime",     //
}

var Tokens_bb = map[byte]string{ // 2-byte tokens
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

var Tokens_ef = map[byte]string{
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
	0x09: "getDate",        //
}

var Tokens_63 = map[byte]string{
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

var Tokens_5d = map[byte]string{
	0x00: "L1", //
	0x01: "L2", //
	0x02: "L3", //
	0x03: "L4", //
	0x04: "L5", //
	0x05: "L6", //
}

var Tokens_7e = map[byte]string{
	0x09: "AxesOff",   //
	0x08: "AxesOn",    //
	0x05: "CoordOff",  //
	0x04: "CoordOn",   //
	0x07: "Dot-Thick", //
}

var Tokens_aa = map[byte]string{
	0x00: "Str1", //
	0x01: "Str2", //
	0x02: "Str3", //
	0x03: "Str4", //
	0x04: "Str5", //
	0x05: "Str6", //
	0x06: "Str7", //
	0x07: "Str8", //
	0x08: "Str9", //
	0x09: "Str0", //
}

var Reverse_tokens = map[string][]byte{}

func init() {
	for key, val := range Tokens {
		Reverse_tokens[val] = []byte{key}
	}
	for key, val := range Tokens_bb {
		Reverse_tokens[val] = []byte{0xbb, key}
	}
	for key, val := range Tokens_ef {
		Reverse_tokens[val] = []byte{0xef, key}
	}
	for key, val := range Tokens_63 {
		Reverse_tokens[val] = []byte{0x63, key}
	}
	for key, val := range Tokens_5d {
		Reverse_tokens[val] = []byte{0x5d, key}
	}
	for key, val := range Tokens_7e {
		Reverse_tokens[val] = []byte{0x7e, key}
	}
	for key, val := range Tokens_aa {
		Reverse_tokens[val] = []byte{0xaa, key}
	}
}

