package tsukuyomi

import (
	"fmt"
	"os"

	"github.com/mewkiz/flac"
)

func CompressWav(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening input file:", err)
	}
	defer inputFile.Close()

	inputFileInfo, err := inputFile.Stat()
	if err != nil {
		fmt.Println("[ERROR] error getting input file info:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating output file:", err)
	}
	defer outputFile.Close()

	encoder, err := flac.NewEncoder(outputFile, flac.DefaultOptions())
	if err != nil {
		fmt.Println("[ERROR] error creating FLAC encoder:", err)
	}
	defer encoder.Close()

	err = encoder.Encode(inputFile, uint32(inputFileInfo.Size()))
	if err != nil {
		fmt.Println("[ERROR] error encoding to FLAC:", err)
	}

	fmt.Println(
		"[SUCCESS] WAV compression using FLAC has been completed. Compressed file:",
		fileName,
	)
}
