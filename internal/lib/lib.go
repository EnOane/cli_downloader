package lib

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"os/exec"
)

const format = "mp4"

type VideoMetadata struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// GetVideoMetadata возвращает метаданные видео
func GetVideoMetadata(videoUrl string) (*VideoMetadata, error) {
	var out bytes.Buffer

	cmd := exec.Command("yt-dlp", "--dump-json", videoUrl)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Error().Err(err).Msgf("error get metadata for: %v", videoUrl)
		return nil, fmt.Errorf("error get metadata for: %w; videoUrl: %v", err, videoUrl)
	}

	var videoInfo VideoMetadata

	err = json.Unmarshal(out.Bytes(), &videoInfo)
	if err != nil {
		log.Error().Err(err).Msgf("error parse metadata to json - %v", videoUrl)
		return nil, fmt.Errorf("error parse metadata to json: %w; videoUrl: %v", err, videoUrl)
	}

	log.Info().Msgf("received metadata for: %v; %v", videoInfo, videoUrl)

	return &videoInfo, err
}

// GetHashVideo возвращает hash видео
func GetHashVideo(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("error open file - %v", filePath)
		return "", fmt.Errorf("error open file: %w; path: %v", err, filePath)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		log.Error().Err(err).Msgf("error copy hash from file - %v", filePath)
		return "", fmt.Errorf("error copy hash from file: %w", err)
	}
	hash := hex.EncodeToString(hasher.Sum(nil))

	log.Info().Msgf("calc hash: %v for: %v", hash, filePath)

	return hash, nil
}

// GetVideoFileSize возвращает размер файла
func GetVideoFileSize(filePath string) (int64, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		log.Error().Err(err).Msgf("error get file size - %v", filePath)
		return 0, fmt.Errorf("error get file size: %w", err)
	}

	size := stat.Size()

	log.Info().Msgf("get file size %v for: %v", size, filePath)

	return size, err
}

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

	// TODO: обработка ошибки в горутине
	go func() {
		defer close(out)

		// Запускаем yt-dlp для потоковой загрузки видео
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
