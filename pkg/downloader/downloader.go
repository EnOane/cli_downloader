package downloader

import (
	"fmt"
	"github.com/EnOane/cli_downloader/internal/lib"
	"github.com/EnOane/cli_downloader/internal/services/rutube"
	"github.com/EnOane/cli_downloader/internal/services/vkvideo"
	"github.com/EnOane/cli_downloader/internal/services/youtube"
	"github.com/rs/zerolog/log"
	"net/url"
	"strings"
	"time"
)

// TODO: custom errors type
// TODO: metadata file
// TODO: вынести в const провайдеров

type downloader struct {
	yt *youtube.YoutubeService
	vk *vkvideo.VkVideoService
	rt *rutube.RutubeService
}

var _dl *downloader

func getDownloader() *downloader {
	if _dl != nil {
		return _dl
	}

	downloadLib := lib.NewLib()

	_dl = &downloader{
		yt: youtube.NewYoutubeService(downloadLib),
		vk: vkvideo.NewVkVideoService(downloadLib),
		rt: rutube.NewRutubeService(downloadLib),
	}

	return _dl
}

// DownloadVideo загрузка видео с rutube, vk, youtube с сохранением файла
func DownloadVideo(videoUrl *url.URL, destPath string) (string, error) {
	dl := getDownloader()
	provider := prepareProviderData(videoUrl)
	return downloadAndSave(dl, videoUrl, destPath, provider)
}

// DownloadStreamVideo загрузка видео с rutube, vk, youtube потоком
func DownloadStreamVideo(videoUrl *url.URL) (<-chan []byte, string, error) {
	dl := getDownloader()
	provider := prepareProviderData(videoUrl)
	return downloadStream(dl, videoUrl, provider)
}

// prepareProviderData возвращает наименование провайдера
func prepareProviderData(videoUrl *url.URL) string {
	host := videoUrl.Host
	host = strings.ReplaceAll(host, "www.", "")

	provider := strings.Split(host, ".")[0]
	return provider
}

// downloadAndSave логика скачивания и сохранения
func downloadAndSave(dl *downloader, videoUrl *url.URL, destPath string, provider string) (string, error) {
	// время выполнения
	exStart := time.Now()

	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", provider))

	// строковое значение url
	videoUrlStr := videoUrl.String()

	// имя сохраненного файла
	var filenamePath string
	var err error

	// эмуляция разной логики провайдеров
	switch provider {
	case "rutube":
		filenamePath, err = dl.rt.DownloadAndSave(videoUrlStr, destPath)
	case "vk", "vkvideo":
		filenamePath, err = dl.vk.DownloadAndSave(videoUrlStr, destPath)
	case "youtube":
		filenamePath, err = dl.yt.DownloadAndSave(videoUrlStr, destPath)
	default:
		return "", fmt.Errorf("download video from provider %v not supported %w", provider, err)
	}

	// обработка ошибок
	if err != nil {
		return "", err
	}

	log.Info().Msg(fmt.Sprintf("video was downloaded in %v to path '%v'", time.Since(exStart), filenamePath))

	return filenamePath, err
}

// downloadStream логика скачивания потоком
func downloadStream(dl *downloader, videoUrl *url.URL, provider string) (<-chan []byte, string, error) {
	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", provider))

	// строковое значение url
	videoUrlStr := videoUrl.String()

	var filename string
	var in <-chan []byte

	// эмуляция разной логики провайдеров
	switch provider {
	case "rutube":
		in, filename = dl.rt.DownloadStream(videoUrlStr)
	case "vk", "vkvideo":
		in, filename = dl.vk.DownloadStream(videoUrlStr)
	case "youtube":
		in, filename = dl.yt.DownloadStream(videoUrlStr)
	default:
		return nil, "", fmt.Errorf("download video from provider %v not supported", provider)
	}

	return in, filename, nil
}
