package img

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func CompressPng(fileName string) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening PNG file:", err)
		return
	}
	defer inputFile.Close()

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding PNG file:", err)
		return
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating PNG file:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, inputImage)
	if err != nil {
		fmt.Println("[ERROR] error encoding PNG file:", err)
		return
	}

	fmt.Println("[SUCCESS] PNG compression has been completed")
}
