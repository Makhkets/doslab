package post

import (
	"doslab/internal/service/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ServicePost struct {
	postUseCase       domain.PostUseCase
	wordsCacheUseCase domain.WordsCacheUseCase
	cache             *domain.WordsCache
}

func NewServicePost(postUseCase domain.PostUseCase, wordsCacheUseCase domain.WordsCacheUseCase, cache *domain.WordsCache) *ServicePost {
	return &ServicePost{
		postUseCase,
		wordsCacheUseCase,
		cache,
	}
}

func (t ServicePost) Comments(c *gin.Context) {
	postId := c.Param("id")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Error params",
			"result": "Bad",
		})
		return
	}

	iPostId, err := strconv.Atoi(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "Error server",
			"result": "Bad",
		})
		return
	}

	comments := t.wordsCacheUseCase.GetCommentsById(t.cache, iPostId)

	c.JSON(http.StatusOK, gin.H{
		"msg":    "Server work",
		"result": "Ok",
		"data":   comments,
	})
}
