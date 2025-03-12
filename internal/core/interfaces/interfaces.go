package interfaces

type Downloader interface {
	Download(videoUrl, destPath string) (string, error)
}
