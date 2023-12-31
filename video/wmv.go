package tsukuyomi

import (
	"fmt"
	"os"
	"os/exec"
)

func CompressWmv(fileName string) {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		fileName,
		"-c:v",
		"libx264",
		"-crf",
		"23",
		"-c:a",
		"aac",
		"-b:a",
		"128k",
		fileName,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("[ERROR] error running FFmpeg:", err)
	}

	fmt.Println("[SUCCESS] WMV compression has been completed. Compressed file:", fileName)
}
