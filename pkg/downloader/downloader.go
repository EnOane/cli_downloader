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
	filename, err := execute(videoUrl, destPath, provider)
	return filename, err
}

// prepareProviderData возвращает наименование провайдера
func prepareProviderData(videoUrl *url.URL) string {
	host := videoUrl.Host
	host = strings.ReplaceAll(host, "www.", "")

	provider := strings.Split(host, ".")[0]
	return provider
}

// execute логика скачивания и сохранения
func execute(videoUrl *url.URL, destPath string, provider string) (string, error) {
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
