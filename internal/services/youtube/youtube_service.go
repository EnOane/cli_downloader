package youtube

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type Service struct {
	lib interfaces.DownloaderLib
}

func NewYoutubeService(lib interfaces.DownloaderLib) interfaces.DownloaderProvider {
	return &Service{lib}
}

func (y *Service) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return y.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (y *Service) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return y.lib.DownloadStream(videoUrl, id)
}
