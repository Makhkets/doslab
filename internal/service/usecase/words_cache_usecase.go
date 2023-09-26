package usecase

import "doslab/internal/service/domain"

type wordsCacheUseCase struct {
	wordsCacheRepository domain.WordsCacheRepository
}

func (w wordsCacheUseCase) GetCommentsById(cache *domain.WordsCache, id int) []domain.PostComments {
	return w.wordsCacheRepository.GetCommentsById(cache, id)
}

func (w wordsCacheUseCase) UpdateWordsCache(cache *domain.WordsCache) error {
	return w.wordsCacheRepository.UpdateWordsCache(cache)
}

func NewWordsCacheUseCase(wordsCacheRepository domain.WordsCacheRepository) domain.WordsCacheUseCase {
	return wordsCacheUseCase{wordsCacheRepository}
}
