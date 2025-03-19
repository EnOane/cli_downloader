package vkvideo

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/google/uuid"
)

type VkVideoService struct {
	lib interfaces.DownloaderLib
}

func NewVkVideoService(lib interfaces.DownloaderLib) *VkVideoService {
	return &VkVideoService{lib}
}

func (v *VkVideoService) DownloadAndSave(videoUrl, destPath string) (string, error) {
	id := uuid.New().String()

	return v.lib.DownloadAndSave(videoUrl, id, destPath)
}

func (v *VkVideoService) DownloadStream(videoUrl string) (<-chan []byte, string) {
	id := uuid.New().String()

	return v.lib.DownloadStream(videoUrl, id)
}
