package convert

import "../tokens"
import "core:encoding/endian"
import "core:encoding/hex"
import "core:fmt"
import "core:path/filepath"
import "core:os"
import "core:strings"

Conversion_Error :: enum {
  File_Size_Too_Small,
  File_Read_Error,
  File_Write_Error,
  Name_Too_Long,
  Comment_Too_Long,
  Hex_Decoding_Error,
  Unknown_Command,
}

eightxp_to_txt :: proc(from_path, to_path: string, debug: bool) -> Conversion_Error {
  eightxp_path: string = from_path
  from_path_ext := filepath.ext(from_path)
  if from_path_ext == "" {
    eightxp_path = fmt.tprintf("%s.8xp", from_path)
  } else if from_path_ext != ".8xp" {
    eightxp_path = from_path[:len(from_path) - len(from_path_ext)]
    eightxp_path = fmt.tprintf("%s.8xp", eightxp_path)
  }
  
  txt_path := to_path
  meta_path := to_path
  to_path_ext := filepath.ext(to_path)
  if to_path_ext == "" {
    txt_path = fmt.tprintf("%s.txt", to_path)
    meta_path = fmt.tprintf("%s.meta", to_path)
  } else if to_path_ext != "" {
    tmp := to_path[:len(to_path) - len(to_path_ext)]
    txt_path = tmp
    meta_path = tmp
    txt_path = fmt.tprintf("%s.txt", txt_path)
    meta_path = fmt.tprintf("%s.meta", meta_path)
  }

  program_meta: [4]string
  program_data: []byte

  byte_data, err := os.read_entire_file(eightxp_path, context.allocator)
  defer delete(byte_data)
  if err != nil {
    return Conversion_Error.File_Read_Error
  }

  if len(byte_data) < 76 {
    return Conversion_Error.File_Size_Too_Small
  }

  program_meta[0] = string(byte_data[60:67]) // store bytes 60 - 67 to metadata (program name)
  program_meta[1] = string(byte_data[11:52]) // store bytes 11 - 52 to metadat (transmission comment)
  program_meta[2] = fmt.tprintf("%02x", byte_data[59]) // store byte 59 to metadata (type id)
  program_meta[3] = fmt.tprintf("%02x", byte_data[69]) // store byte 69 to metadata (flag)
  
  if debug {
    fmt.printfln("Program name: %s", program_meta[0])
    fmt.printfln("Transmission comment: %s", program_meta[1])
    fmt.printfln("Type id: %s", program_meta[2])
    fmt.printfln("Flag: %s", program_meta[3])
  }

  program_data = byte_data[74:len(byte_data) - 2] // store bytes 74 - end-2 to program data (program data)

  meta_builder: strings.Builder
  strings.builder_init(&meta_builder)
  defer strings.builder_destroy(&meta_builder)

  for meta in program_meta {
    for meta_byte in meta {
      if meta_byte == 0x00 {
        break
      }
      strings.write_byte(&meta_builder, u8(meta_byte))
    }
    strings.write_byte(&meta_builder, '\n')
  }

  meta_string := strings.to_string(meta_builder)
  final_bytes := transmute([]byte)meta_string

  err = os.write_entire_file(meta_path, final_bytes)
  if err != nil {
    return Conversion_Error.File_Write_Error
  }

  builder := strings.builder_make()
  defer strings.builder_destroy(&builder)

  idx := 0
  for idx < len(program_data) {
    curr_byte := program_data[idx]
    next_byte: byte

    step := 1

    if idx + 1 < len(program_data) {
      next_byte = program_data[idx + 1]

      switch curr_byte {
        case 0xbb:
          s, ok := tokens.tokens_bb[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case 0xef:
          s, ok := tokens.tokens_ef[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case 0x63:
          s, ok := tokens.tokens_63[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case 0x53:
          s, ok := tokens.tokens_5d[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case 0x7e:
          s, ok := tokens.tokens_7e[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case 0xaa:
          s, ok := tokens.tokens_aa[next_byte]
          if ok {
            write_known_command(&builder, s, debug)
            step = 1
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
        case:
          s, ok := tokens.normal_tokens[curr_byte]
          if ok {
            write_known_command(&builder, s, debug)
          } else {
            write_unknown_command(&builder, curr_byte, debug)
          }
      }
    } else {
      s, ok := tokens.normal_tokens[curr_byte]
      if ok {
        write_known_command(&builder, s, debug)
      } else {
        write_unknown_command(&builder, curr_byte, debug)
      }
    }

    idx += step
  }

  program_string := strings.to_string(builder)
  err = os.write_entire_file(txt_path, program_string)
  if err != nil {
    return Conversion_Error.File_Write_Error
  }


  return nil
}

write_unknown_command :: proc(builder: ^strings.Builder, b: byte, debug: bool) {
  strings.write_string(builder, fmt.tprintf("%c", b))
  if debug {
    strings.write_string(builder, fmt.tprintf("<%02x>", b))
    fmt.printfln(`Write unknown command: %c<%s>`, b, fmt.tprintf("<%02x>", b))
  }
}

write_known_command :: proc(builder: ^strings.Builder, s: string, debug: bool) {
  strings.write_string(builder, fmt.tprintf("%s", s))
  if debug {
    new_s, alloc := strings.replace_all(s, "\n", "\\n", context.temp_allocator)
    fmt.printfln(`Write command: %s`, new_s)
  }
}

txt_to_eightxp :: proc(from_path, to_path: string) -> Conversion_Error {
  txt_path := from_path
  meta_path := from_path
  from_path_ext := filepath.ext(from_path)
  if from_path_ext == "" {
    txt_path = fmt.tprintf("%s.txt", from_path)
    meta_path = fmt.tprintf("%s.meta", from_path)
  } else if from_path_ext != "" {
    tmp := from_path[:len(from_path) - len(from_path_ext)]
    txt_path = tmp
    meta_path = tmp
    txt_path = fmt.tprintf("%s.txt", txt_path)
    meta_path = fmt.tprintf("%s.meta", meta_path)
  }

  eightxp_path: string = to_path
  to_path_ext := filepath.ext(to_path)
  if to_path_ext == "" {
    eightxp_path = fmt.tprintf("%s.8xp", to_path)
  } else if to_path_ext != ".8xp" {
    eightxp_path = to_path[:len(to_path) - len(to_path_ext)]
    eightxp_path = fmt.tprintf("%s.8xp", eightxp_path)
  }

  // read meta
  meta: [4]string
  {
    meta_file, err := os.read_entire_file(meta_path, context.allocator)
    if err != nil {
      return Conversion_Error.File_Read_Error
    }
    defer delete(meta_file)
    meta_string := string(meta_file)

    // read first 4 lines
    i := 0
    for line in strings.split_lines_iterator(&meta_string) {
      meta[i] = line
      i += 1
      if i > 3 {
        break
      }
    }
  }

  if len(meta[0]) > 8 {
    return Conversion_Error.Name_Too_Long
  }
  if len(meta[1]) > 42 {
    return Conversion_Error.Comment_Too_Long
  }

  program_name: [8]byte // null-padded name
  for i := 0; i < len(meta[0]); i += 1 {
    program_name[i] = meta[0][i]
  }
  program_comment: [42]byte // null-padded comment
  for i := 0; i < len(meta[1]); i += 1 {
    program_comment[i] = meta[1][i]
  }
  
  file_bytes: [dynamic]byte
  defer delete(file_bytes)

  append(&file_bytes, 0x2a, 0x2a, 0x54, 0x49, 0x38, 0x33, 0x46, 0x2a) // append signature
  append(&file_bytes, 0x1a, 0x0a) // append signature_part_2
  append(&file_bytes, 0x0a) // mystery_byte
  append(&file_bytes, ..program_comment[:]) // append comment
  append(&file_bytes, 0x00, 0x00) // append placeholder meta_and_body_length. Set later
  append(&file_bytes, 0x0d)       // Append flag
	append(&file_bytes, 0x00)       // Append unknown
	append(&file_bytes, 0x00, 0x00) // Append placeholder body_and_checksum_length. Set later
  { // append file_type
    decoded_file_type, ok := hex.decode(transmute([]u8)meta[2], context.allocator)
    defer delete(decoded_file_type)
    if !ok {
      return Conversion_Error.Hex_Decoding_Error
    }
    append(&file_bytes, ..decoded_file_type[:]) 
  }
  append(&file_bytes, ..program_name[:]) // append program_name
  append(&file_bytes, 0x00) // append version
  { // append is_archived
    decoded_archive_state, ok := hex.decode(transmute([]u8)meta[3], context.allocator)
    defer delete(decoded_archive_state)
    if !ok {
      return Conversion_Error.Hex_Decoding_Error
    }
    append(&file_bytes, ..decoded_archive_state[:]) 
  }
  append(&file_bytes, 0x00, 0x00) // append placeholder body_and_checksum_length_2. Set later
  append(&file_bytes, 0x00, 0x00) // append placeholder body_length. Set later

  body_length: u16 = 2
  { // append program data
    program_bytes, err := os.read_entire_file(txt_path, context.allocator)
    if err != nil {
      return Conversion_Error.File_Read_Error
    }
    defer delete(program_bytes)

    longest_command_length := 0
    for key in tokens.reverse_tokens {
      if len(key) > longest_command_length {
        longest_command_length = len(key)
      }
    }

    i := 0
    n := longest_command_length
    for i < len(program_bytes) {
      if i + n > len(program_bytes) {
        n = len(program_bytes) - i
      }
      if n <= 0 {
        return  Conversion_Error.Unknown_Command
      }

      command_bytes := program_bytes[i:i+n]
      arr, ok := tokens.reverse_tokens[string(command_bytes)]
      if ok {
        for j := 0; j < len(arr); j += 1 {
          body_length += 1
          append(&file_bytes, arr[j])
        }
        i += n
        n = longest_command_length
      } else {
        n -= 1
      }
    }
  }

  buf: [2]u8
  { // set meta_and_body_length
    endian.put_u16(buf[:], .Little, u16(len(file_bytes) - 57))
    file_bytes[53] = buf[0]
		file_bytes[54] = buf[1]
  }
  { // set body_and_checksum_length
    endian.put_u16(buf[:], .Little, body_length)
    file_bytes[57] = buf[0]
		file_bytes[58] = buf[1]
  }
  { // set_body_and_checksum_length_2
    endian.put_u16(buf[:], .Little, body_length)
    file_bytes[70] = buf[0]
		file_bytes[71] = buf[1]
  }
  { // set body_length
    endian.put_u16(buf[:], .Little, body_length - 2)
    file_bytes[72] = buf[0]
		file_bytes[73] = buf[1]
  }
  { // append checksum
    checksum: u16 = 0
    for i := 55; i < len(file_bytes); i += 1 {
      checksum += u16(file_bytes[i])
    }
    endian.put_u16(buf[:], .Little, checksum)
    append(&file_bytes, buf[0], buf[1])
  }

  err := os.write_entire_file_from_bytes(eightxp_path, file_bytes[:])
  if err != nil {
    return Conversion_Error.File_Write_Error
  }

  return nil
}

