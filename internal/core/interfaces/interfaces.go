package interfaces

import "net/url"

type VideoMetadata struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type DownloaderLib interface {
	DownloadAndSave(videoUrl, filename, destPath string) (string, error)
	DownloadStream(videoUrl, filename string) (<-chan []byte, string)
	GetVideoMetadata(videoUrl string) (*VideoMetadata, error)
	GetHashVideo(filePath string) (string, error)
	GetVideoFileSize(filePath string) (int64, error)
}

type DownloaderProvider interface {
	DownloadAndSave(videoUrl, destPath string) (string, error)
	DownloadStream(videoUrl string) (<-chan []byte, string)
}

type Downloader interface {
	DownloadVideo(videoUrl *url.URL, destPath string) (string, error)
	DownloadStreamVideo(videoUrl *url.URL) (<-chan []byte, string, error)
}
