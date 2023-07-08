package requests

type SCategoryStoreRequest struct {
	Name string `json:"name" binding:"required"`
}

type SCategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
