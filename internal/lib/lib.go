package lib

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os/exec"
)

const format = "mp4"

// DownloadAndSave
func DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()
	template := id + ".%(ext)s"

	cmd := exec.Command("yt-dlp", "-f", format, "-o", template, "--path", destPath, videoUrl)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// TODO: проверка кодов ошибок
		return "", errors.New(fmt.Sprintf("download video err: %v", string(output)))
	}

	return destPath + "/" + id + "." + format, nil
}
