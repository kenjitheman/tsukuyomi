package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "main.go/core/audio"
	// "main.go/core/img"
	// "main.go/core/video"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("welcome to tsukuyomi - modern file compressor")
	fmt.Println("choose what type of files you want to compress ! (audio, video or img)")

	input, _ := reader.ReadString('\n')
	file_type := strings.Split(strings.TrimSpace(input), " ")
	if file_type == nil {
		fmt.Println("[ERROR] error parsing input")
	}

	switch file_type[0] {
	case "audio":
		fmt.Print("enter command /format FILE_NAME\nexample: /mp3 man.mp3")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")
		if len(parts) < 3 {
			fmt.Println("[ERROR] invalid command format | usage: /format FILE_NAME")
			return
		}

		command := parts[0]
		fileName := parts[1]

		switch command {
		case "mp3":
			fmt.Println("mp3", fileName)
		case "aac":
			fmt.Println("aac", fileName)
		case "aiff":
			fmt.Println("aiff", fileName)
		case "alac":
			fmt.Println("alac", fileName)
		case "flac":
			fmt.Println("flac", fileName)
		case "ogg":
			fmt.Println("ogg", fileName)
		case "wav":
			fmt.Println("wav", fileName)
		case "wma":
			fmt.Println("wma", fileName)
		default:
			fmt.Println("unknown file format: ", fileName)
		}

	case "video":
		fmt.Print("enter command /format FILE_NAME\nexample: /mp4 man.mp4")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")
		if len(parts) < 3 {
			fmt.Println("[ERROR] invalid command format | usage: /format FILE_NAME")
			return
		}

		command := parts[0]
		fileName := parts[1]

		switch command {
		case "avi":
			fmt.Println("avi", fileName)
		case "flv":
			fmt.Println("flv", fileName)
		case "mkv":
			fmt.Println("mkv", fileName)
		case "mov":
			fmt.Println("mov", fileName)
		case "mp4":
			fmt.Println("mp4", fileName)
		case "webm":
			fmt.Println("webm", fileName)
		case "wmv":
			fmt.Println("wmv", fileName)
		case "h264":
			fmt.Println("h264", fileName)
		default:
			fmt.Println("unknown file format: ", fileName)
		}
	case "img":
		fmt.Print("enter command /format FILE_NAME\nexample: /png man.png")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")
		if len(parts) < 3 {
			fmt.Println("[ERROR] invalid command format | usage: /format FILE_NAME")
			return
		}

		command := parts[0]
		fileName := parts[1]

		switch command {
		case "png":
			fmt.Println("png", fileName)
		case "jpeg":
			fmt.Println("jpeg", fileName)
		case "webp":
			fmt.Println("webp", fileName)
		case "tiff":
			fmt.Println("tiff", fileName)
		case "jpeg2000":
			fmt.Println("jpeg2000", fileName)
		case "heif":
			fmt.Println("heif", fileName)
		case "bmp":
			fmt.Println("bmp", fileName)
		case "gif":
			fmt.Println("gif", fileName)
		default:
			fmt.Println("unknown file format: ", fileName)
		}

	default:
		fmt.Println("unknown instruction, usage: ")
	}
}
