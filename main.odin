package main

import "convert"
import "core:fmt"
import "core:path/filepath"
import "core:os"

main :: proc() {
  args := os.args
  if len(args) < 2 {
    fmt.println("Please provide at least an input file path")
    return
  }

  input_path := args[1]

  output_path := input_path
  if len(args) >= 3 {
    output_path = args[2]
  }

  input_ext := filepath.ext(input_path)
  
  switch input_ext {
    case ".8xp":
      err := convert.eightxp_to_txt(input_path, output_path)

      if err != nil {
        fmt.println(err)
      }
    case ".txt":
      err := convert.txt_to_eightxp(input_path, output_path)

      if err != nil {
        fmt.println(err)
      }
    case:
      fmt.printfln("Unsupported file extension \"%s\"", input_ext)
  }
}

