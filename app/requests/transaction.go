package requests

type STransactionStoreRequest struct {
	CategoryId string `json:"category_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	Note       string `json:"note" binding:"required"`
	Amount     int32  `json:"amount" binding:"required"`
}

type STransactionUpdateRequest struct {
	CategoryId string `json:"category_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Currency   string `json:"currency" binding:"required"`
	Note       string `json:"note" binding:"required"`
	Amount     int32  `json:"amount" binding:"required"`
}
