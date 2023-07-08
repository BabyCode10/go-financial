package controllers

import (
	"errors"
	"go-financial/app/models"
	"go-financial/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SCategoryService struct {
	category models.ICategory
}

func Category(category models.ICategory) *SCategoryService {
	return &SCategoryService{
		category: category,
	}
}

func (service SCategoryService) Index(context *gin.Context) {
	categories, err := service.category.Get()

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"data": &categories,
		},
	)
}

func (service SCategoryService) Create(context *gin.Context) {
	var request requests.SCategoryStoreRequest

	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	category, err := service.category.Store(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"data": &category,
		},
	)
}

func (service SCategoryService) Show(context *gin.Context) {
	categoryId := context.Param("category_id")

	category, err := service.category.Find(categoryId)

	if err == errors.New("Category not found.") {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Category not found.",
			},
		)

		context.Abort()

		return
	}

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"data": &category,
		},
	)
}

func (service SCategoryService) Update(context *gin.Context) {
	categoryId := context.Param("category_id")

	var request requests.SCategoryUpdateRequest

	err := context.ShouldBindJSON(&request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	category, err := service.category.Update(categoryId, &request)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"data": &category,
		},
	)
}

func (service SCategoryService) Delete(context *gin.Context) {
	categoryId := context.Param("category_id")

	category, err := service.category.Delete(categoryId)

	if err != nil {
		context.Error(err)

		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		context.Abort()

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"data": &category,
		},
	)
}
