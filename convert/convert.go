package convert

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/mehonje/TI-Convert/tokens"
)

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
				s, ok := tokens.Tokens_bb[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			case 0xef:
				s, ok := tokens.Tokens_ef[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString("Wait ") // Add "Wait" command (0xef) if no
				}
			case 0x63:
				s, ok := tokens.Tokens_63[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			case 0x5d:
				s, ok := tokens.Tokens_5d[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString("/") // Add division operator (0x5d) if no
				}
			case 0x7e:
				s, ok := tokens.Tokens_7e[next_val] // Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			case 0xaa:
				s, ok := tokens.Tokens_aa[next_val] //Check if mapping exists
				if ok {
					builder.WriteString(s) // Replace if yes,
					step = 2
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			default:
				_, ok := tokens.Tokens[val] // Check if mapping exists
				if ok {
					builder.WriteString(tokens.Tokens[val]) // Replace if yes,
				} else {
					builder.WriteString(string(val)) // Turn into string if no
				}
			}
		} else { // 1-byte token
			s, ok := tokens.Tokens[val] // Check if mapping exists
			if ok {
				builder.WriteString(s) // Replace if yes,
			} else {
				builder.WriteString(string(val)) // Turn into string if no
			}
		}
		{
			//var hex_string string = fmt.Sprintf("<%02x>", val)
			//builder.WriteString(hex_string) // Uncomment to see hex equivalent
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
	from_path = strings.TrimSpace(from_path)          // Remove whitespace
	from_path = strings.TrimSuffix(from_path, ".txt") // Remove ".txt" suffix

	to_path = strings.TrimSpace(to_path)     // Remove whitespace
	if !strings.HasSuffix(to_path, ".8xp") { // If file path doesn't have ".8xp" suffix,
		to_path = to_path + ".8xp" // Append it
	}

	var metadata []string
	metadata_file, err := os.Open(from_path + ".meta")
	if err != nil {
		log.Fatal("Failed to open file ", from_path, ".meta: ", err)
	}
	defer metadata_file.Close()
	scanner := bufio.NewScanner(metadata_file)
	for scanner.Scan() {
		metadata = append(metadata, scanner.Text())
	}
	if len(metadata) < 4 {
		log.Fatal("Metadata file must contain at least 4 lines")
	}
	if len(metadata[0]) > 8 {
		log.Fatal("Program name (line 1) in \"", from_path, ".meta\" cannot be more than 8 characters")
	}
	if len(metadata[1]) > 42 {
		log.Fatal("Program comment (line 2) in \"", from_path, ".meta\" cannot be more than 42 characters")
	}

	var program_byte_data []byte

	program_byte_data = append(program_byte_data, 0x2a, 0x2a, 0x54, 0x49, 0x38, 0x33, 0x46, 0x2a) // Append signature
	program_byte_data = append(program_byte_data, 0x1a, 0x0a)                                     // Append signature_part_2
	program_byte_data = append(program_byte_data, 0x0a)                                           // Append mystery byte
	{                                                                                             // Append comment
		comment_padded := make([]byte, 42)
		copy(comment_padded, []byte(metadata[1]))
		program_byte_data = append(program_byte_data, comment_padded...)
	}
	program_byte_data = append(program_byte_data, 0x00, 0x00) // Append placeholder meta_and_body_length. Set later on
	program_byte_data = append(program_byte_data, 0x0d)       // Append flag
	program_byte_data = append(program_byte_data, 0x00)       // Append unknown
	program_byte_data = append(program_byte_data, 0x00, 0x00) // Append placeholder body_and_checksum_length. Set later
	{                                                         // Append file_type
		b, err := hex.DecodeString(metadata[2])
		if err != nil {
			log.Fatal("Failed to convert string\"", metadata[2], "\"to byte")
		}
		program_byte_data = append(program_byte_data, b[0])
	}
	{ // Append program_name
		name_padded := make([]byte, 8)
		copy(name_padded, []byte(metadata[0]))
		program_byte_data = append(program_byte_data, name_padded...)
	}
	program_byte_data = append(program_byte_data, 0x00) // Append version
	{                                                   // Append is_archived
		b, err := hex.DecodeString(metadata[3])
		if err != nil {
			log.Fatal("Failed to convert string\"", metadata[3], "\"to byte")
		}
		program_byte_data = append(program_byte_data, b[0])
	}
	program_byte_data = append(program_byte_data, 0x00, 0x00) // Append placeholder body_and_checksum_length_2. Set later
	program_byte_data = append(program_byte_data, 0x00, 0x00) // Append placeholder body_length. Set later

	var body_length uint16 = 2
	{ // Append program data
		byte_data, err := os.ReadFile(from_path + ".txt") // Read program data
		if err != nil {
			log.Fatal("Failed to read program data: ", err)
		}

		longest_command_length := 0
		for key := range tokens.Reverse_tokens {
			if len(key) > longest_command_length {
				longest_command_length = len(key)
			}
		}

		i := 0
		n := longest_command_length
		for i < len(byte_data) {
			if i+n > len(byte_data) {
				n = len(byte_data) - i
			}
			if n <= 0 {
				log.Fatal("Unknown command at position ", i)
			}
			command_bytes := byte_data[i : i+n]
			arr, ok := tokens.Reverse_tokens[string(command_bytes)]
			if ok {
				for j := 0; j < len(arr); j++ {
					body_length++
					program_byte_data = append(program_byte_data, arr[j])
				}
				i += n
				n = longest_command_length
			} else {
				n -= 1
			}
		}
	}

	{ // Set meta_and_body_length
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(len(program_byte_data)-57))
		program_byte_data[53] = buf[0]
		program_byte_data[54] = buf[1]
	}
	{ // Set body_and_checksum_length
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, body_length)
		program_byte_data[57] = buf[0]
		program_byte_data[58] = buf[1]
	}
	{ // Set body_and_checksum_length_2
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, body_length)
		program_byte_data[70] = buf[0]
		program_byte_data[71] = buf[1]
	}
	{ // Set body_length
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, body_length-2)
		program_byte_data[72] = buf[0]
		program_byte_data[73] = buf[1]
	}
	{ // Append checksum
		var checksum uint16 = 0
		for i := 55; i < len(program_byte_data); i++ {
			checksum += uint16(program_byte_data[i])
		}
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, checksum)
		program_byte_data = append(program_byte_data, buf[0], buf[1])
	}

	err = os.WriteFile(to_path, program_byte_data, 0644)
	if err != nil {
		log.Fatal("Failed to create", to_path, ", err")
	}
}

