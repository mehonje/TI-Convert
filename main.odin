package main

import "convert"
import "core:bufio"
import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

main :: proc() {
  fmt.print(`
0: Convert 8xp -> txt
1: Convert txt -> 8xp
Idx? `)

  scanner: bufio.Scanner
  stdin_stream := os.to_stream(os.stdin)
  bufio.scanner_init(&scanner, stdin_stream, context.allocator)
  idx: i8 = 0
  if bufio.scanner_scan(&scanner) {
    text := bufio.scanner_text(&scanner)

    value, ok := strconv.parse_int(text)
    if !ok {
      fmt.printfln("Invalid number \"%s\"", text)
      bufio.scanner_destroy(&scanner)
      return
    }

    idx = i8(value)
  }
  bufio.scanner_destroy(&scanner)

  switch idx {
    case 0:
      convert_8xp_to_txt()
    case 1:
      convert_txt_to_8xp()
  }
}

convert_8xp_to_txt :: proc() {
  scanner: bufio.Scanner
  stdin_stream := os.to_stream(os.stdin)
  bufio.scanner_init(&scanner, stdin_stream, context.allocator)

  fmt.print("\n8xp to convert: ")

  from_path: string
  if bufio.scanner_scan(&scanner) {
    from_path = strings.clone(bufio.scanner_text(&scanner), context.allocator)
  }

  fmt.print("Project name to output to: ")

  to_path: string
  if bufio.scanner_scan(&scanner) {
    to_path = strings.clone(bufio.scanner_text(&scanner), context.allocator)
  }

  bufio.scanner_destroy(&scanner)

  err := convert.eightxp_to_txt(from_path, to_path)
  if err != nil {
    fmt.println(err)
  }
}

convert_txt_to_8xp:: proc() {
  scanner: bufio.Scanner
  stdin_stream := os.to_stream(os.stdin)
  bufio.scanner_init(&scanner, stdin_stream, context.allocator)

  fmt.print("\nTxt to convert: ")

  from_path: string
  if bufio.scanner_scan(&scanner) {
    from_path = strings.clone(bufio.scanner_text(&scanner), context.allocator)
  }

  fmt.print("8xp to output to: ")

  to_path: string
  if bufio.scanner_scan(&scanner) {
    to_path = strings.clone(bufio.scanner_text(&scanner), context.allocator)
  }

  bufio.scanner_destroy(&scanner)

  err := convert.txt_to_eightxp(from_path, to_path)
  if err != nil {
    fmt.println(err)
  }
}

