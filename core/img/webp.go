package img

import (
	"fmt"
	"io"
	"os"

	"github.com/chai2010/webp"
)

func CompressWebp(fileName string, quality int) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[ERROR] error opening input WebP:", err)
	}
	defer inputFile.Close()

	inputData, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("[ERROR] error reading input WebP:", err)
	}

	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("[ERROR] error creating output WebP:", err)
	}
	defer outputFile.Close()

	options := &webp.Options{Lossless: false, Quality: quality} // Quality 0 - 100
	err = webp.Encode(outputFile, inputData, options)
	if err != nil {
		fmt.Println("[ERROR] error encoding output WebP:", err)
	}

	fmt.Println("[SUCCESS] WebP compression has been completed")
}
