package lib

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"os/exec"
)

const format = "mp4"

// DownloadAndSave скачивание видео в файл
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

// DownloadStream скачивание видео в поток
func DownloadStream(videoUrl string) (<-chan []byte, string) {
	out := make(chan []byte)

	// Запускаем yt-dlp для потоковой загрузки видео
	go func() {
		defer close(out)

		cmd := exec.Command("yt-dlp", "-f", format, "-o", "-", videoUrl)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal().Err(err)
		}

		// Запускаем процесс
		if err := cmd.Start(); err != nil {
			log.Fatal().Err(err)
		}

		buffer := make([]byte, 1024*1024)
		for {
			n, err := stdout.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal().Err(err).Msg("Ошибка при чтении данных")
				return
			}

			chunk := make([]byte, n)
			copy(chunk, buffer[:n])
			out <- chunk
		}

		// Дожидаемся завершения процесса
		if err := cmd.Wait(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	return out, uuid.New().String() + "." + format
}
