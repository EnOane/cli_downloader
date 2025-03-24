package rutube

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type Service struct {
	lib interfaces.DownloaderLib
}

func NewRutubeService(lib interfaces.DownloaderLib) interfaces.DownloaderProvider {
	return &Service{lib}
}

func (r *Service) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return r.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (r *Service) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return r.lib.DownloadStream(videoUrl, id)
}
