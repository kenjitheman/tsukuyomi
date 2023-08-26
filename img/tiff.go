package tsukuyomi

import (
	"fmt"
	"image"
	"os"

	"github.com/fhs/gotiff/tiff/lzw"
)

func CompressTiff(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening input image:", err)
	}
	defer inputFile.Close()

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding input image:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating output TIFF:", err)
	}
	defer outputFile.Close()

	tiffWriter, err := lzw.NewWriter(outputFile, &lzw.Options{Predictor: true})
	if err != nil {
		fmt.Println("[ERROR] error creating TIFF writer:", err)
	}
	defer tiffWriter.Close()

	err = tiffWriter.Encode(inputImage, nil)
	if err != nil {
		fmt.Println("[ERROR] error encoding output TIFF:", err)
	}

	fmt.Println("[SUCCESS] TIFF compression has been completed")
}
