package main

import (
	"flag"
	"github.com/EnOane/cli_downloader/pkg/downloader"
	"github.com/rs/zerolog/log"
	"net/url"
)

func main() {
	var (
		videoUrlStr, destPath string
	)

	flag.StringVar(&videoUrlStr, "url", "https://youtube.com/shorts/X-xPsJfIWK0", "link to video")
	flag.StringVar(&destPath, "dest", "./", "download folder")
	flag.Parse()

	videoUrl, err := parseUrl(videoUrlStr)
	if err != nil {
		log.Fatal().Msg("url is not valid")
	}

	_, err = downloader.DownloadVideo(videoUrl, destPath)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func parseUrl(videoUrlRaw string) (*url.URL, error) {
	return url.Parse(videoUrlRaw)
}
