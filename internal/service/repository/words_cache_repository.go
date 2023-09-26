package repository

import (
	"doslab/internal/db/postgresql"
	"doslab/internal/service/domain"
	"fmt"
	"sort"
	"strings"
)

type wordsCacheRepository struct {
	postRepository domain.PostRepository
}

func (w wordsCacheRepository) GetCommentsById(cache *domain.WordsCache, id int) []domain.PostComments {
	return cache.PostComments[id]
}

func (w wordsCacheRepository) UpdateWordsCache(cache *domain.WordsCache) error {
	posts, err := w.postRepository.Posts()
	if err != nil {
		return err
	}

	for _, post := range posts {
		wordCache := make(map[string]int, 256)

		// Получаем комментарии к посту по его Id
		comments, err := w.postRepository.Comments(post.Id)
		if err != nil {
			return err
		}

		for _, v := range comments {
			// Убираем лишние символы
			v.Body = strings.Replace(v.Body, "\n", " ", -1)
			// Получаем слова в предложении
			words := strings.Split(v.Body, " ")

			// Далее каждое полученное слово кэшируем
			for _, word := range words {
				// Ищем слово в кэше
				_, ok := wordCache[word]
				if ok {
					// Нашли "Встречу"
					wordCache[word]++
				} else {
					// "Встреч" слова пока 0
					wordCache[word] = 0
				}
			}
		}

		for key, value := range wordCache {
			postComments := domain.PostComments{
				PostID: post.Id,
				Word:   key,
				Count:  value,
			}

			con, err := postgresql.DB.Exec("update statistics set count=$1 where post_id=$2 and word=$3;", postComments.Count, postComments.PostID, postComments.Word)
			if err != nil {
				fmt.Println(err)
			}

			count, err := con.RowsAffected()
			if err != nil {
				fmt.Println(err)
			}

			// Проверяем обновлены ли данные
			if count == 0 {
				_, err := postgresql.DB.Exec("insert into statistics (post_id,word,count) values($1,$2,$3);", postComments.PostID, postComments.Word, postComments.Count)
				if err != nil {
					fmt.Println(err)
				}
			}

			cache.PostComments[post.Id] = append(cache.PostComments[post.Id], postComments)
		}

		sort.Slice(cache.PostComments[post.Id], func(i, j int) bool {
			return cache.PostComments[post.Id][i].Count > cache.PostComments[post.Id][j].Count
		})
	}

	return nil
}

func NewWordsCacheRepository(postRep domain.PostRepository) domain.WordsCacheRepository {
	return wordsCacheRepository{postRep}
}
