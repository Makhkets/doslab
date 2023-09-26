package app

import (
	"doslab/internal/config"
	"doslab/internal/db/postgresql"
	"doslab/internal/delivery/http/post"
	"doslab/internal/service/domain"
	"doslab/internal/service/repository"
	"doslab/internal/service/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Server() {
	cache := &domain.WordsCache{}
	cache.PostComments = make(map[int][]domain.PostComments, 1024)

	postRep := repository.NewPostRepository()
	postUseCase := usecase.NewPostUseCase(postRep)

	wordsCacheRep := repository.NewWordsCacheRepository(postRep)
	wordsCacheUseCase := usecase.NewWordsCacheUseCase(wordsCacheRep)

	postService := post.NewServicePost(postUseCase, wordsCacheUseCase, cache)

	// В начале обновляем далее каждые 5 минут запуска программы
	go func() {
		wordsCacheUseCase.UpdateWordsCache(cache)

		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			wordsCacheUseCase.UpdateWordsCache(cache)
		}
	}()

	// Initialize dbs
	config.InitializeConfig()
	postgresql.InitializeDb()

	// DataBase migrations
	err := postgresql.Migrations()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			postGroup := v1Group.Group("/post")
			{
				postGroup.GET("/:id/comments/statistics", postService.Comments)
			}
		}
	}

	router.Run(":8000")
}
