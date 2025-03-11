package downloader

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

// Download загрузка видео с rutube, vk, youtube
func Download(videoUrl *url.URL, destPath string) {
	host := videoUrl.Host
	host = strings.ReplaceAll(host, "www.", "")

	provider := strings.Split(host, ".")[0]

	execute := func(download func(u string, p string)) {
		download(videoUrl.String(), destPath)
	}

	exStart := time.Now()
	log.Info().Msg(fmt.Sprintf("download video from '%v' has been started", host))

	// эмуляция разной логики провайдеров
	switch provider {
	case "rutube":
		execute(downloadFromRutube)
	case "vk":
		execute(downloadFromVKVideo)
	case "vkvideo":
		execute(downloadFromVKVideo)
	case "youtube":
		execute(downloadFromYoutube)
	default:
		log.Fatal().Msg(fmt.Sprintf("download video from %v not supported", provider))
	}

	since := time.Since(exStart)

	if destPath == "./" {
		log.Info().Msg(fmt.Sprintf("video was downloaded in %v", since))
	} else {
		log.Info().Msg(fmt.Sprintf("video was downloaded in %v to path '%v'", since, destPath))
	}

	fmt.Println()
}

func downloadFromRutube(videoUrl string, path string) {
	downloadVideo(videoUrl, path)
}

func downloadFromVKVideo(videoUrl string, path string) {
	downloadVideo(videoUrl, path)
}

func downloadFromYoutube(videoUrl string, path string) {
	downloadVideo(videoUrl, path)
}

func downloadVideo(videoUrl string, path string) {
	cmd := exec.Command("yt-dlp", "-f", "best", "-o", "%(title)s.%(ext)s", "--path", path, videoUrl)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("downloadVideo err: %v", err))
	}
}
