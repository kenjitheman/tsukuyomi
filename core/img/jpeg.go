package img

import (
	"fmt"
	"image/jpeg"
	"os"
)

// Quality 0-100

func main(fileName string, quality int) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening JPEG file:", err)
		return
	}
	defer inputFile.Close()

	inputImage, err := jpeg.Decode(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error decoding JPEG file:", err)
		return
	}

	outputFile, err := os.Create("output.jpg")
	if err != nil {
		fmt.Println("[ERROR] error creating JPEG file:", err)
		return
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, inputImage, &jpeg.Options{Quality: quality})
	if err != nil {
		fmt.Println("[ERROR] error encoding JPEG file:", err)
		return
	}

	fmt.Println("[SUCCESS] JPEG compression has been completed")
}
