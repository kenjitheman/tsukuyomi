package tsukuyomi

import (
	"fmt"
	"os"
	"os/exec"
)

func CompressWebM(fileName string) {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		fileName,
		"-c:v",
		"libvpx-vp9",
		"-b:v",
		"1M",
		"-c:a",
		"libopus",
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

	fmt.Println("[SUCCESS] WebM compression has been completed. Compressed file:", fileName)
}
