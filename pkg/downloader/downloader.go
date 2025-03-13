package downloader

import (
	"errors"
	"fmt"
	"github.com/EnOane/cli_downloader/internal/services/rutube"
	"github.com/EnOane/cli_downloader/internal/services/vkvideo"
	"github.com/EnOane/cli_downloader/internal/services/youtube"
	"github.com/rs/zerolog/log"
	"net/url"
	"strings"
	"time"
)

// TODO: custom filename
// TODO: custom errors type

// DownloadVideo загрузка видео с rutube, vk, youtube
func DownloadVideo(videoUrl *url.URL, destPath string) (string, error) {
	provider := prepareProviderData(videoUrl)
	return downloadAndSave(videoUrl, destPath, provider)
}

func DownloadStreamVideo(videoUrl *url.URL) (<-chan []byte, string, error) {
	provider := prepareProviderData(videoUrl)
	return downloadStream(videoUrl, provider)
}

// prepareProviderData возвращает наименование провайдера
func prepareProviderData(videoUrl *url.URL) string {
	host := videoUrl.Host
	host = strings.ReplaceAll(host, "www.", "")

	provider := strings.Split(host, ".")[0]
	return provider
}

// downloadAndSave логика скачивания и сохранения
func downloadAndSave(videoUrl *url.URL, destPath string, provider string) (string, error) {
	// время выполнения
	exStart := time.Now()
	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", provider))

	// имя сохраненного файла
	var filenamePath string
	var err error

	// строковое значение url
	videoUrlStr := videoUrl.String()

	// эмуляция разной логики провайдеров
	switch provider {
	// TODO: вынести в const
	// TODO: реализовать DI
	case "rutube":
		filenamePath, err = rutube.Download(videoUrlStr, destPath)
	case "vk", "vkvideo":
		filenamePath, err = vkvideo.Download(videoUrlStr, destPath)
	case "youtube":
		filenamePath, err = youtube.Download(videoUrlStr, destPath)
	default:
		return "", errors.New(fmt.Sprintf("download video from provider %v not supported", provider))
	}

	// обработка ошибок
	if err != nil {
		return "", err
	}

	log.Info().Msg(fmt.Sprintf("video was downloaded in %v to path '%v'", time.Since(exStart), filenamePath))

	return filenamePath, err
}

func downloadStream(videoUrl *url.URL, provider string) (<-chan []byte, string, error) {
	// время выполнения
	exStart := time.Now()
	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", provider))

	// строковое значение url
	videoUrlStr := videoUrl.String()

	var filename string
	var in <-chan []byte

	// эмуляция разной логики провайдеров
	switch provider {
	// TODO: вынести в const
	// TODO: реализовать DI
	case "rutube":
		in, filename = rutube.DownloadStream(videoUrlStr)
	case "vk", "vkvideo":
		in, filename = vkvideo.DownloadStream(videoUrlStr)
	case "youtube":
		in, filename = youtube.DownloadStream(videoUrlStr)
	default:
		return nil, "", errors.New(fmt.Sprintf("download video from provider %v not supported", provider))
	}

	log.Info().Msg(fmt.Sprintf("video was downloaded in %v to path '%v'", time.Since(exStart), filename))

	return in, filename, nil
}
