package rutube

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type RutubeService struct {
	lib interfaces.DownloaderLib
}

func NewRutubeService(lib interfaces.DownloaderLib) *RutubeService {
	return &RutubeService{lib}
}

func (r *RutubeService) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return r.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (r *RutubeService) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return r.lib.DownloadStream(videoUrl, id)
}
