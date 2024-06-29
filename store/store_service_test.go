package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	_store := InitializeStore()
	testStoreService = _store
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// 元URLと短縮URLのマッピングを保存する
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// 元のURLを取得する
	retrievedUrl := RetrieveInitialUrl(shortURL)

	// 検証
	assert.Equal(t, initialLink, retrievedUrl)
}
