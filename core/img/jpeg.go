package img

import (
	"fmt"
	"image/jpeg"
	"os"
)

func CompressJpeg(fileName string, quality int) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening JPEG file:", err)
	}
	defer inputFile.Close()

	inputImage, err := jpeg.Decode(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding JPEG file:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating JPEG file:", err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, inputImage, &jpeg.Options{Quality: quality}) // Quality 0-100
	if err != nil {
		fmt.Println("[ERROR] error encoding JPEG file:", err)
	}

	fmt.Println("[SUCCESS] JPEG compression has been completed")
}
