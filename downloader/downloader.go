package downloader

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// TODO: custom filename
// TODO: custom errors type

// Download загрузка видео с rutube, vk, youtube
func Download(videoUrl *url.URL, filename, destPath string) (string, error) {
	host := videoUrl.Host
	host = strings.ReplaceAll(host, "www.", "")

	provider := strings.Split(host, ".")[0]

	// имя сохраненного файла
	var filenamePath string
	var execErr error

	// обобщенная функция для скачивания
	execute := func(download func(u string, f, p string) (string, error)) {
		val, err := download(videoUrl.String(), filename, destPath)
		if err != nil {
			execErr = err
		}

		filenamePath = val
	}

	// время выполнения
	exStart := time.Now()
	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", host))

	// эмуляция разной логики провайдеров
	switch provider {
	case "rutube":
		execute(downloadFromRutube)
	case "vk", "vkvideo":
		execute(downloadFromVKVideo)
	case "youtube":
		execute(downloadFromYoutube)
	default:
		return "", errors.New(fmt.Sprintf("download video from provider %v not supported", provider))
	}

	if execErr == nil {
		// время выполнения
		since := time.Since(exStart)

		if destPath == "./" {
			log.Info().Msg(fmt.Sprintf("video was downloaded in %v", since))
		} else {
			log.Info().Msg(fmt.Sprintf("video was downloaded in %v to path '%v'", since, destPath))
		}

		fmt.Println()
	} else {
		return "", execErr
	}

	return filenamePath, execErr
}

func downloadFromRutube(videoUrl string, filename, path string) (string, error) {
	return downloadVideo(videoUrl, filename, path)
}

func downloadFromVKVideo(videoUrl string, filename, path string) (string, error) {
	return downloadVideo(videoUrl, filename, path)
}

func downloadFromYoutube(videoUrl string, filename, path string) (string, error) {
	return downloadVideo(videoUrl, filename, path)
}

func downloadVideo(videoUrl string, filename, path string) (string, error) {
	//var title string
	//
	//if filename == "" {
	//	title = "%(title)s.%(ext)s"
	//} else {
	//	title = filename + ".%(ext)s"
	//}

	// TODO: проверка кодов ошибок
	cmd := exec.Command("yt-dlp", "--print", "filename", "-o", "%(title)s.%(ext)s", "--path", path, videoUrl)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("download video err: %v", string(output)))
	}

	// TODO: проверка вывода - имя файла
	p := string(output)
	name := filepath.Base(p)
	fileNameWithoutExt := strings.TrimSuffix(name, filepath.Ext(name))

	return fileNameWithoutExt, nil
}
