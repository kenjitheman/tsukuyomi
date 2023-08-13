package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"main.go/core/audio"
	"main.go/core/img"
	"main.go/core/video"
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
			audio.CompressMp3(fileName)
		case "aiff":
			audio.CompressAiff(fileName)
		case "flac":
			audio.CompressFlac(fileName)
		case "ogg":
			audio.CompressOgg(fileName)
		case "wav":
			audio.CompressWav(fileName)
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
			video.CompressAvi(fileName)
		case "flv":
			video.CompressFlv(fileName)
		case "mkv":
			video.CompressMkv(fileName)
		case "mov":
			video.CompressMov(fileName)
		case "mp4":
			video.CompressMp4(fileName)
		case "webm":
			video.CompressWebM(fileName)
		case "wmv":
			video.CompressWmv(fileName)
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
		quality, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch command {
		case "png":
			img.CompressPng(fileName)
		case "jpeg":
			fmt.Println(
				"jpeg cannot be compressed without quality loss, you need to specify quality (1-100)\nexample: man.jpeg, 50",
			)
			img.CompressJpeg(fileName, quality)
		case "webp":
			fmt.Println(
				"webp cannot be compressed without quality loss, you need to specify quality (1-100)\nexample: man.webp, 50",
			)
			img.CompressWebp(fileName, quality)
		case "tiff":
			img.CompressTiff(fileName)
		case "gif":
			img.CompressGif(fileName)
		default:
			fmt.Println("unknown file format: ", fileName)
		}

	default:
		fmt.Println("unknown instruction, usage: ")
	}
}
