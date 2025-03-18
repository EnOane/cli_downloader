package lib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

const dest = "./temp_test"

type exp struct {
	id       string
	provider string
	link     string
	size     int64
	hash     string
}

var expected = [...]exp{
	{
		id:       "X-xPsJfIWK0",
		provider: "youtube",
		link:     "https://youtube.com/shorts/X-xPsJfIWK0",
		size:     int64(1566683),
		hash:     "69cceb62c6cb4f4a3059cae20e04deeb79328be93052d332664b3fa96dd16e74",
	},
	{
		id:       "-46638176_456239535",
		provider: "vk",
		link:     "https://vk.com/clip-46638176_456239535",
		size:     int64(7789825),
		hash:     "bfa838ff6b857948a943d1d0af6bdc8ba045de3a1479d1d46b43a4dd2253c414",
	},
	{
		id:       "ce0d3b5fddbb6829282d7a406f9df882",
		provider: "rutube",
		link:     "https://rutube.ru/shorts/ce0d3b5fddbb6829282d7a406f9df882",
		size:     int64(6829427),
		hash:     "60eda595333b1e744e3f042c4387ef074265bcbecac2553d95174f211ec256c3",
	},
}

// TODO: реорганизовать структуру ответа при потоке и файле

// TODO: мок загрузки файла

func TestGetVideoMetadata(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()
		actual, err := GetVideoMetadata(expected.link)

		assert.Nil(t, err)
		assert.Equal(t, expected.id, actual.Id)
	}

	for _, ex := range expected {
		go testF(ex, &wg)
	}

	wg.Wait()
}

func TestGetHashVideo(t *testing.T) {
	defer os.RemoveAll(dest)

	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()

		filePath, err := DownloadAndSave(expected.link, dest)
		actual, err := GetHashVideo(filePath)

		assert.Nil(t, err)
		assert.Equal(t, expected.hash, actual)
	}

	for _, ex := range expected {
		go testF(ex, &wg)
	}

	wg.Wait()
}

func TestGetVideoFileSize(t *testing.T) {
	defer os.RemoveAll(dest)

	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()

		filePath, err := DownloadAndSave(expected.link, dest)
		actual, err := GetVideoFileSize(filePath)

		assert.Nil(t, err)
		assert.Equal(t, expected.size, actual)
	}

	for _, ex := range expected {
		go testF(ex, &wg)
	}

	wg.Wait()
}

func TestDownloadStream(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()

		ch, filename := DownloadStream(expected.link)

		actual := make([]byte, 0, expected.size)
		for bytes := range ch {
			actual = append(actual, bytes...)
		}

		assert.Equal(t, expected.size, int64(len(actual)))
		assert.NotEmpty(t, filename)
	}

	for _, ex := range expected {
		go testF(ex, &wg)
	}

	wg.Wait()
}

// short тесты
// если тесты дольше трех минут - беда
// t.Skip
// префиксы к тесту Test, Benchmark

// по пайплайну доставки - линтеры, solar, unit тесты

func TestDownloadAndSaveParallel(t *testing.T) {
	tmpDir := t.TempDir()

	for _, ex := range expected {
		t.Run("DownloadAndSaveParallel "+ex.provider, func(t *testing.T) {
			t.Parallel()
			actual, err := DownloadAndSave(ex.link, tmpDir)
			hash, _ := GetHashVideo(actual)

			assert.Nil(t, err)
			assert.Equal(t, ex.hash, hash)
		})
	}
}
