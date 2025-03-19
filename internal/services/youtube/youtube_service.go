package youtube

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type YoutubeService struct {
	lib interfaces.DownloaderLib
}

func NewYoutubeService(lib interfaces.DownloaderLib) *YoutubeService {
	return &YoutubeService{lib}
}

func (y *YoutubeService) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return y.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (y *YoutubeService) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return y.lib.DownloadStream(videoUrl, id)
}
