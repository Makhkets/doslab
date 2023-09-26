package cache

import "doslab/internal/service/domain"

type ServiceWordsCache struct {
	wordsCacheService domain.WordsCacheUseCase
}

func NewServiceWordsCache(wordsCacheService domain.WordsCacheUseCase) ServiceWordsCache {
	return ServiceWordsCache{wordsCacheService}
}
