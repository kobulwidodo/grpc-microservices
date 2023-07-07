package entity

type CreateUserParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
