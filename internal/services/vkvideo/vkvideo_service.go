package vkvideo

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type Service struct {
	lib interfaces.DownloaderLib
}

func NewVkVideoService(lib interfaces.DownloaderLib) interfaces.DownloaderProvider {
	return &Service{lib}
}

func (v *Service) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return v.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (v *Service) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return v.lib.DownloadStream(videoUrl, id)
}
