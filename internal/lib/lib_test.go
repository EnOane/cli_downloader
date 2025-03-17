package lib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

const dest = "./temp_test"

type exp struct {
	id   string
	link string
	size int64
	hash string
}

var expected = [...]exp{
	{
		id:   "X-xPsJfIWK0",
		link: "https://youtube.com/shorts/X-xPsJfIWK0",
		size: int64(1566683),
		hash: "a9c739d0089b8cdca2fa0215e9b0c5ef727494608f531e4f86eb39ec2d4ec9ef",
	},
	{
		id:   "-46638176_456239535",
		link: "https://vk.com/clip-46638176_456239535",
		size: int64(7789825),
		hash: "bfa838ff6b857948a943d1d0af6bdc8ba045de3a1479d1d46b43a4dd2253c414",
	},
	{
		id:   "ce0d3b5fddbb6829282d7a406f9df882",
		link: "https://rutube.ru/shorts/ce0d3b5fddbb6829282d7a406f9df882",
		size: int64(6829427),
		hash: "60eda595333b1e744e3f042c4387ef074265bcbecac2553d95174f211ec256c3",
	},
}

// TODO: реорганизовать структуру ответа при потоке и файле

// TODO: мок загрузки файла

func TestDownloadAndSave(t *testing.T) {
	defer os.RemoveAll(dest)

	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()
		actual, err := DownloadAndSave(expected.link, dest)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}

	for _, expected := range expected {
		expected := expected
		go testF(expected, &wg)
	}

	wg.Wait()
}

func TestGetVideoMetadata(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(len(expected))

	testF := func(expected exp, wg *sync.WaitGroup) {
		defer wg.Done()
		actual, err := GetVideoMetadata(expected.link)

		assert.Nil(t, err)
		assert.Equal(t, expected.id, actual.Id)
	}

	for _, expected := range expected {
		expected := expected
		go testF(expected, &wg)
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

	for _, expected := range expected {
		expected := expected
		go testF(expected, &wg)
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

	for _, expected := range expected {
		expected := expected
		go testF(expected, &wg)
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

		file, _ := os.Create(filename)
		defer file.Close()

		file.Write(actual)

		assert.Equal(t, expected.size, int64(len(actual)))
		assert.NotEmpty(t, filename)
	}

	for _, expected := range expected {
		expected := expected
		go testF(expected, &wg)
	}

	wg.Wait()
}
