package controllers

import (
	"errors"
	"go-financial/app/models"
	"go-financial/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type STransactionService struct {
	transaction models.ITransaction
}

func Transaction(transaction models.ITransaction) *STransactionService {
	return &STransactionService{
		transaction: transaction,
	}
}

func (service STransactionService) Index(context *gin.Context) {
	userId := context.GetString("user_id")

	transactions, err := service.transaction.GetByUser(userId)

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
			"data": &transactions,
		},
	)
}

func (service STransactionService) Create(context *gin.Context) {
	userId := context.GetString("user_id")

	var request requests.STransactionStoreRequest

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

	transaction, err := service.transaction.Store(userId, &request)

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
			"data": &transaction,
		},
	)
}

func (service STransactionService) Show(context *gin.Context) {
	userId := context.GetString("user_id")
	transactionId := context.Param("transaction_id")

	transaction, err := service.transaction.FindByUser(userId, transactionId)

	if err == errors.New("Transaction not found.") {
		context.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Transaction not found.",
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
			"data": &transaction,
		},
	)
}

func (service STransactionService) Update(context *gin.Context) {
	userId := context.GetString("user_id")
	transactionId := context.Param("transaction_id")

	var request requests.STransactionUpdateRequest

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

	transaction, err := service.transaction.Update(userId, transactionId, &request)

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
			"data": &transaction,
		},
	)
}

func (service STransactionService) Delete(context *gin.Context) {
	userId := context.GetString("user_id")
	transactionId := context.Param("transaction_id")

	transaction, err := service.transaction.Delete(userId, transactionId)

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
			"data": &transaction,
		},
	)
}
