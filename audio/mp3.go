package tsukuyomi

import (
	"fmt"
	"os"

	mp3 "github.com/hajimehoshi/go-mp3"
)

func CompressMp3(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening input file:", err)
	}
	defer inputFile.Close()

	pcm, _, err := mp3.Decode(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding input file:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating output file:", err)
	}
	defer outputFile.Close()

	encoder := mp3.NewEncoder(outputFile, mp3.BitRate192kbps) // change the bitrate as needed

	_, err = encoder.Write(pcm)
	if err != nil {
		fmt.Println("[ERROR] error encoding to MP3:", err)
	}

	encoder.Close()

	fmt.Println("[SUCCESS] MP3 re-encoding has been completed. Output file:", fileName)
}
