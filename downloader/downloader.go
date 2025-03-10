package downloader

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/url"
	"os/exec"
	"strings"
)

func Download(videoUrl *url.URL, destPath string) {
	host := videoUrl.Host
	t := strings.Split(host, ".")[0]

	execute := func(download func(u string, p string)) {
		download(videoUrl.String(), destPath)
	}

	log.Info().Msg(fmt.Sprintf("download video from %v has been started", t))

	switch t {
	case "rutube":
		execute(downloadFromRutube)
	case "vk":
		execute(downloadFromVKVideo)
	case "vkvideo":
		execute(downloadFromVKVideo)
	case "youtube":
		execute(downloadFromYoutube)
	default:
		log.Fatal().Msg(fmt.Sprintf("downloadVideo from %v not supported", t))
	}

	log.Info().Msg("video has been downloaded")
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
