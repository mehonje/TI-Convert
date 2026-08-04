package main

import "convert"
import "core:fmt"
import "core:path/filepath"
import "core:os"
import "core:strings"

main :: proc() {
  args: [dynamic]string
  defer delete(args)
  for arg in os.args {
    append(&args, arg)
  }

  debug := false
  
  fmt.println(args)

  {
    i := 0
    for i < len(args) {
      if strings.has_prefix(args[i], "-") {
        switch args[i] {
          case "-debug":
            debug = true
        }
        ordered_remove(&args, i)
        i -= 1
      }
      i += 1
    }
  }

  fmt.println(args)

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
      err := convert.eightxp_to_txt(input_path, output_path, debug)

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

