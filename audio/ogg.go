package tsukuyomi

import (
	"fmt"
	"os"

	"github.com/mccoyst/ogg"
)

func CompressOgg(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening input file:", err)
	}
	defer inputFile.Close()

	oggDecoder, err := ogg.NewDecoder(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error creating Ogg decoder:", err)
	}
	defer oggDecoder.Close()

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating output file:", err)
	}
	defer outputFile.Close()

	oggEncoder, err := ogg.NewEncoder(outputFile, ogg.Options{
		Serial:  oggDecoder.Serial, // reuse the same serial number
		Codec:   ogg.VorbisCodec,
		Quality: 0.6, // 0.1 - 1.0
		// specify other desired encoding options here
	})
	if err != nil {
		fmt.Println("[ERROR] error creating Ogg encoder:", err)
	}
	defer oggEncoder.Close()

	for {
		page, err := oggDecoder.ReadNextPage()
		if err != nil {
			break
		}
		oggEncoder.WritePage(page)
	}

	fmt.Println("[SUCCESS] Ogg re-encoding has been completed. Output file:", fileName)
}
