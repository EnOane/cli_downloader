package di

import (
	"github.com/EnOane/cli_downloader/internal/core/interfaces"
	"github.com/EnOane/cli_downloader/internal/lib"
	"github.com/EnOane/cli_downloader/internal/services/rutube"
	"github.com/EnOane/cli_downloader/internal/services/vkvideo"
	"github.com/EnOane/cli_downloader/internal/services/youtube"
	"github.com/EnOane/cli_downloader/pkg/downloader"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

var Container *dig.Container

func MakeDIContainer() {
	Container = dig.New()

	makeProviders()
}

func makeProviders() {
	Container.Provide(func() interfaces.DownloaderLib {
		return lib.NewLib()
	})
	Container.Provide(func(l interfaces.DownloaderLib) interfaces.DownloaderProvider {
		return youtube.NewYoutubeService(l)
	})
	Container.Provide(func(l interfaces.DownloaderLib) interfaces.DownloaderProvider {
		return vkvideo.NewVkVideoService(l)
	})
	Container.Provide(func(l interfaces.DownloaderLib) interfaces.DownloaderProvider {
		return rutube.NewRutubeService(l)
	})
	Container.Provide(func(yt, vk, rt interfaces.DownloaderProvider) interfaces.Downloader {
		return downloader.NewDownloader(yt, vk, rt)
	})
}

func Inject[T any]() T {
	if Container == nil {
		MakeDIContainer()
	}

	var dep T

	err := Container.Invoke(func(d T) { dep = d })
	if err != nil {
		log.Fatal().Err(err)
	}

	return dep
}
