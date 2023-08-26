package tsukuyomi

import (
	"fmt"
	"image/gif"
	"os"
)

func CompressGif(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening GIF file:", err)
	}
	defer inputFile.Close()

	inputGif, err := gif.DecodeAll(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding GIF file:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating GIF file:", err)
	}
	defer outputFile.Close()

	outputGif := gif.GIF{
		Image:     inputGif.Image,
		Delay:     inputGif.Delay,
		LoopCount: inputGif.LoopCount,
	}

	err = gif.EncodeAll(outputFile, &outputGif)
	if err != nil {
		fmt.Println("[ERROR] error encoding GIF file:", err)
	}

	fmt.Println("[SUCCESS] GIF compression has been completed")
}
