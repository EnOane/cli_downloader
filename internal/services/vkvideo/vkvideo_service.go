package vkvideo

import "github.com/EnOane/cli_downloader/internal/lib"

func Download(videoUrl, destPath string) (string, error) {
	return lib.DownloadAndSave(videoUrl, destPath)
}

func DownloadStream(videoUrl string) (<-chan []byte, string) {
	return lib.DownloadStream(videoUrl)
}
