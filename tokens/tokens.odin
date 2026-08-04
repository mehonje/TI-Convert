package tokens

import "base:runtime"
import "core:fmt"

normal_tokens: map[byte]string
tokens_bb: map[byte]string
tokens_ef: map[byte]string
tokens_63: map[byte]string
tokens_5d: map[byte]string
tokens_7e: map[byte]string
tokens_aa: map[byte]string

reverse_tokens: map[string][]byte

@(init)
init_tokens :: proc "contextless" () {
  context = runtime.default_context()
  init_normal_tokens()
  init_bb_tokens()
  init_ef_tokens()
  init_63_tokens()
  init_5d_tokens()
  init_7e_tokens()
  init_aa_tokens()

  init_reverse_tokens()
}

init_normal_tokens :: proc() {
  normal_tokens = make(map[byte]string)

	normal_tokens[0x3f] = "\n"
	normal_tokens[0x82] = "*"
  normal_tokens[0x83] = "/"
	normal_tokens[0x70] = "+"
	normal_tokens[0x71] = "-"
	normal_tokens[0x10] = "("
	normal_tokens[0x11] = ")"
	normal_tokens[0x29] = " "
	normal_tokens[0x2a] = "\""
	normal_tokens[0x2d] = "!"
	normal_tokens[0xce] = "If "
	normal_tokens[0xcf] = "Then"
	normal_tokens[0xd0] = "Else"
	normal_tokens[0xd3] = "For("
	normal_tokens[0xd1] = "While "
	normal_tokens[0xd2] = "Repeat "
	normal_tokens[0xd4] = "End"
	normal_tokens[0xd8] = "Pause "
	normal_tokens[0xd6] = "Lbl "
	normal_tokens[0xd7] = "Goto "
	normal_tokens[0xef] = "Wait "
	normal_tokens[0xda] = "IS>("
	normal_tokens[0xdb] = "DS<("
	normal_tokens[0xe6] = "Menu("
	normal_tokens[0x5f] = "prgm"
	normal_tokens[0xd5] = "Return"
	normal_tokens[0xd9] = "Stop"
	normal_tokens[0xdc] = "Input "
	normal_tokens[0xdd] = "Prompt"
	normal_tokens[0xde] = "Disp "
	normal_tokens[0xdf] = "Dispgraph"
	normal_tokens[0xe5] = "DispTable"
	normal_tokens[0xe0] = "Output("
	normal_tokens[0xad] = "getKey"
	normal_tokens[0xe1] = "ClrHome"
	normal_tokens[0xfb] = "ClrTable"
	normal_tokens[0xe8] = "Get("
	normal_tokens[0xe7] = "Send("
	normal_tokens[0x85] = "ClrDraw"
	normal_tokens[0x9c] = "Line("
	normal_tokens[0xa6] = "Horizontal "
	normal_tokens[0x9d] = "Vertical "
	normal_tokens[0xa7] = "Tangent("
	normal_tokens[0xa9] = "DrawF "
	normal_tokens[0xa4] = "Shade("
	normal_tokens[0xa8] = "DrawInv "
	normal_tokens[0xa5] = "Circle("
	normal_tokens[0x93] = "Text("
	normal_tokens[0x9e] = "Pt-On("
	normal_tokens[0x9f] = "Pt-Off("
	normal_tokens[0xa0] = "Pt-Change("
	normal_tokens[0xa1] = "Pxl-On("
	normal_tokens[0xa2] = "Pxl-Off("
	normal_tokens[0xa3] = "Pxl-Change("
	normal_tokens[0x13] = "pxl-Test("
	normal_tokens[0x98] = "StorePic "
	normal_tokens[0x99] = "RecallPic "
	normal_tokens[0x9a] = "StoreGDB "
	normal_tokens[0x9b] = "RecallGDB "
	normal_tokens[0x6a] = "="
	normal_tokens[0x6f] = "!="
	normal_tokens[0x6c] = ">"
	normal_tokens[0xf0] = "^"
	normal_tokens[0x0d] = "^2"
	normal_tokens[0x0c] = "^-1"
	normal_tokens[0xbc] = "sqrt("
	normal_tokens[0xac] = "pi"
	normal_tokens[0x08] = "{"
	normal_tokens[0x09] = "}"
	normal_tokens[0x06] = "["
	normal_tokens[0x07] = "]"
	normal_tokens[0x5b] = "theta"
	normal_tokens[0x2c] = "i"
	normal_tokens[0xaf] = "?"
	normal_tokens[0x04] = "->"
	normal_tokens[0xbe] = "ln("
	normal_tokens[0xc0] = "log("
	normal_tokens[0xc3] = "arcsin("
	normal_tokens[0xc5] = "arccos("
	normal_tokens[0xc7] = "arctan("
	normal_tokens[0x2b] = ","
	normal_tokens[0x3e] = " ="
	normal_tokens[0x03] = ">Frac"
	normal_tokens[0x02] = ">Dec"
	normal_tokens[0x0f] = "^3"
	normal_tokens[0x27] = "fMin("
	normal_tokens[0x28] = "fMax("
	normal_tokens[0x25] = "nDeriv("
	normal_tokens[0x24] = "fnInt("
	normal_tokens[0x22] = "solve("
	normal_tokens[0xb5] = "dim("
	normal_tokens[0xb2] = "abs("
	normal_tokens[0x72] = "Ans"
	normal_tokens[0x14] = "augment("
	normal_tokens[0x05] = "BoxPlot"
	normal_tokens[0xfa] = "ClrList "
	normal_tokens[0xca] = "cosh("
	normal_tokens[0xcb] = "arccosh("
	normal_tokens[0x2e] = "CubicReg"
	normal_tokens[0x65] = "Degree"
	normal_tokens[0x7d] = "DependAsk"
	normal_tokens[0x7c] = "DependAuto"
	normal_tokens[0xb3] = "det("
	normal_tokens[0x01] = ">DMS"
	normal_tokens[0xbf] = "e^("
	normal_tokens[0x68] = "Eng"
	normal_tokens[0xf5] = "ExpReg"
	normal_tokens[0x3a] = "."
	normal_tokens[0x41] = "A"
	normal_tokens[0x42] = "B"
	normal_tokens[0x43] = "C"
	normal_tokens[0x44] = "D"
	normal_tokens[0x45] = "E"
	normal_tokens[0x46] = "F"
	normal_tokens[0x47] = "G"
	normal_tokens[0x48] = "H"
	normal_tokens[0x49] = "I"
	normal_tokens[0x4a] = "J"
	normal_tokens[0x4b] = "K"
	normal_tokens[0x4c] = "L"
	normal_tokens[0x4d] = "M"
	normal_tokens[0x4e] = "N"
	normal_tokens[0x4f] = "O"
	normal_tokens[0x50] = "P"
	normal_tokens[0x51] = "Q"
	normal_tokens[0x52] = "R"
	normal_tokens[0x53] = "S"
	normal_tokens[0x54] = "T"
	normal_tokens[0x55] = "U"
	normal_tokens[0x56] = "V"
	normal_tokens[0x57] = "W"
	normal_tokens[0x58] = "X"
	normal_tokens[0x59] = "Y"
	normal_tokens[0x5a] = "Z"
	normal_tokens[0x30] = "0"
	normal_tokens[0x31] = "1"
	normal_tokens[0x32] = "2"
	normal_tokens[0x33] = "3"
	normal_tokens[0x34] = "4"
	normal_tokens[0x35] = "5"
	normal_tokens[0x36] = "6"
	normal_tokens[0x37] = "7"
	normal_tokens[0x38] = "8"
	normal_tokens[0x39] = "9"
	normal_tokens[0x0a] = "getTime"
}

init_bb_tokens :: proc() {
  tokens_bb = make(map[byte]string)

	tokens_bb[0x45] = "GraphStyle("
	tokens_bb[0x54] = "DelVar "
	tokens_bb[0x2a] = "expr("
	tokens_bb[0x56] = "String->Equ("
	tokens_bb[0x4f] = "a+bi"
	tokens_bb[0x28] = "angle("
	tokens_bb[0x59] = "ANOVA("
	tokens_bb[0x68] = "Archive "
	tokens_bb[0x02] = "bal("
  tokens_bb[0x16] = "binomcdf("
	tokens_bb[0x15] = "binompdf("
	tokens_bb[0x13] = "x^2cdf("
	tokens_bb[0x1d] = "x^pdf("
	tokens_bb[0x40] = "x^2-Test("
	tokens_bb[0x57] = "Clear Entries"
	tokens_bb[0x52] = "ClrAllLists"
	tokens_bb[0x25] = "conj("
	tokens_bb[0x29] = "cumSum("
	tokens_bb[0x07] = "dbd("
	tokens_bb[0x67] = "DiagnosticOff"
	tokens_bb[0x66] = "DiagnosticOn"
	tokens_bb[0x31] = "e"
	tokens_bb[0x06] = ">Eff("
	tokens_bb[0x55] = "Equ>String("
	tokens_bb[0x51] = "ExprOff"
	tokens_bb[0x50] = "ExprOn"
}

init_ef_tokens :: proc() {
  tokens_ef = make(map[byte]string)

	tokens_ef[0x65] = "GraphColor("
	tokens_ef[0x11] = "OpenLib("
	tokens_ef[0x12] = "ExecLib"
	tokens_ef[0x98] = "eval("
	tokens_ef[0x97] = "toString("
	tokens_ef[0x41] = "BLUE"
	tokens_ef[0x42] = "RED"
	tokens_ef[0x43] = "BLACK"
	tokens_ef[0x44] = "MAGENTA"
	tokens_ef[0x45] = "GREEN"
	tokens_ef[0x46] = "ORANGE"
	tokens_ef[0x47] = "BROWN"
	tokens_ef[0x48] = "NAVY"
	tokens_ef[0x49] = "LTBLUE"
	tokens_ef[0x4a] = "YELLOW"
	tokens_ef[0x4b] = "WHITE"
	tokens_ef[0x4c] = "LTGRAY"
	tokens_ef[0x4d] = "MEDGRAY"
	tokens_ef[0x4e] = "GRAY"
	tokens_ef[0x4f] = "DARKGRAY"
	tokens_ef[0x67] = "TextColor("
	tokens_ef[0x5b] = "BackgroundOn "
	tokens_ef[0x64] = "BackgroundOff "
	tokens_ef[0x2e] = "l"
	tokens_ef[0x33] = "Sigma("
	tokens_ef[0x34] = "logBASE("
	tokens_ef[0xa6] = "piecewise("
	tokens_ef[0x3B] = "AUTO"
	tokens_ef[0x6c] = "BorderColor"
	tokens_ef[0x93] = "CENTER"
	tokens_ef[0x02] = "checkTmr("
	tokens_ef[0x14] = "x^2GOF-Test("
	tokens_ef[0x38] = "CLASSIC"
	tokens_ef[0x0f] = "ClockOff"
	tokens_ef[0x10] = "ClockOn"
	tokens_ef[0x06] = "dayOfWk("
	tokens_ef[0x3c] = "DEC"
	tokens_ef[0x6b] = "DetectAsymOff"
	tokens_ef[0x6a] = "DetectAsymOn"
	tokens_ef[0x75] = "Dot-Thin"
	tokens_ef[0x09] = "getDate"
}

init_63_tokens :: proc() {
  tokens_63 = make(map[byte]string)

  tokens_63[0x0a] = "Xmin"
  tokens_63[0x0b] = "Xmax"
  tokens_63[0x02] = "Xscl"
  tokens_63[0x0c] = "Ymin"
  tokens_63[0x0d] = "Ymax"
  tokens_63[0x03] = "Yscl"
  tokens_63[0x36] = "Xres"
  tokens_63[0x26] = "deltaX"
  tokens_63[0x27] = "deltaY"
  tokens_63[0x28] = "XFact"
  tokens_63[0x39] = "Yfact"
  tokens_63[0x38] = "TraceStep"
}

init_5d_tokens :: proc() {
  tokens_5d = make(map[byte]string)

  tokens_5d[0x00] = "L1"
  tokens_5d[0x01] = "L2"
  tokens_5d[0x02] = "L3"
  tokens_5d[0x03] = "L4"
  tokens_5d[0x04] = "L5"
  tokens_5d[0x05] = "L6"
}

init_7e_tokens :: proc() {
  tokens_7e = make(map[byte]string)

  tokens_7e[0x09] = "AxesOff"
  tokens_7e[0x08] = "AxesOn"
  tokens_7e[0x05] = "CoordOff"
  tokens_7e[0x04] = "CoordOn"
  tokens_7e[0x07] = "Dot-Thick"
}

init_aa_tokens :: proc() {
  tokens_aa = make(map[byte]string)

  tokens_aa[0x00] = "Str1"
  tokens_aa[0x01] = "Str2"
  tokens_aa[0x02] = "Str3"
  tokens_aa[0x03] = "Str4"
  tokens_aa[0x04] = "Str5"
  tokens_aa[0x05] = "Str6"
  tokens_aa[0x06] = "Str7"
  tokens_aa[0x07] = "Str8"
  tokens_aa[0x08] = "Str9"
  tokens_aa[0x09] = "Str0"
}

init_reverse_tokens :: proc() {
  reverse_tokens = make(map[string][]byte)

  for key, val in normal_tokens {
    reverse_tokens[val] = make([]byte, 1)
    reverse_tokens[val][0] = key
  }
  for key, val in tokens_bb {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0xbb
    reverse_tokens[val][1] = key
  }
  for key, val in tokens_ef {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0xef
    reverse_tokens[val][1] = key
  }
  for key, val in tokens_63 {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0x63
    reverse_tokens[val][1] = key
  }
  for key, val in tokens_5d {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0x5d
    reverse_tokens[val][1] = key
  }
  for key, val in tokens_7e {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0x7e
    reverse_tokens[val][1] = key
  }
  for key, val in tokens_aa {
    reverse_tokens[val] = make([]byte, 2)
    reverse_tokens[val][0] = 0xaa
    reverse_tokens[val][1] = key
  }
}

