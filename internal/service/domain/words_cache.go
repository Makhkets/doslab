package domain

type WordsCache struct {
	// Для получения количества "Встреч" используется id поста
	// по нему получается массив сколько было встреч
	PostComments map[int][]PostComments
}

type WordsCacheRepository interface {
	UpdateWordsCache(cache *WordsCache) error
	GetCommentsById(cache *WordsCache, id int) []PostComments
}

type WordsCacheUseCase interface {
	UpdateWordsCache(cache *WordsCache) error
	GetCommentsById(cache *WordsCache, id int) []PostComments
}
