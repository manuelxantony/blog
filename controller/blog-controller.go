package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manuelxantony/blog/blogerrors"
	"github.com/manuelxantony/blog/entity"
	"github.com/manuelxantony/blog/service"
)

type BlogController interface {
	ShowAll(ctx *gin.Context)
	ShowById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service service.BlogService
}

func New(service service.BlogService) BlogController {
	return &controller{
		service: service,
	}
}

func (b *controller) ShowAll(ctx *gin.Context) {
	posts, err := b.service.ShowAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (b *controller) ShowById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, blogerrors.ErrIDNotFound.Error())
	}

	post, err := b.service.ShowById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (b *controller) Create(ctx *gin.Context) {
	var post entity.BlogPost
	ctx.BindJSON(&post)
	id, err := b.service.Create(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	post.ID = id
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog post created",
		"post":    post,
	})
}

func (b *controller) Update(ctx *gin.Context) {
	var post entity.BlogPost
	ctx.BindJSON(&post)
	err := b.service.Update(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Blog post updated")
}

func (b *controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, blogerrors.ErrIDNotFound.Error())
	}
	err = b.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "Post deleted")
}
