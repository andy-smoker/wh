package server

type User struct {
	ID       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"` //  binding:"required" для валидации полей в теле запроса
	Username string `json:"username" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
