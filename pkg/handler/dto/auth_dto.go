package handler

type LoginRequest struct {
	Login    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
