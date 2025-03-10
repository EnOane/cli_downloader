package main

import (
	"flag"
	"github.com/EnOane/cli_downloader/downloader"
	"github.com/rs/zerolog/log"
	"net/url"
)

func main() {
	var (
		videoUrlStr, destPath string
	)

	flag.StringVar(&videoUrlStr, "url", "", "link to video")
	flag.StringVar(&destPath, "dest", "./", "download folder")
	flag.Parse()

	videoUrl, err := parseUrl(videoUrlStr)
	if err != nil {
		log.Fatal().Msg("url is not valid")
	}

	downloader.Download(videoUrl, destPath)
}

func parseUrl(videoUrlRaw string) (*url.URL, error) {
	return url.Parse(videoUrlRaw)
}
