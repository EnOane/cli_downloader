package rutube

import "github.com/EnOane/cli_downloader/internal/lib"

func Download(videoUrl, destPath string) (string, error) {
	return lib.DownloadAndSave(videoUrl, destPath)
}
