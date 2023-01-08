package entity

// username, user_email, user_password, user_activation_status,status
type User struct {
	UserID           int    `json:"user_id"`
	Username         string `json:"username"`
	UserEmail        string `json:"user_email"`
	Password         string `json:"user_password"`
	ActivationStatus string `json:"user_activation_status"`
	Status           int    `json:"status"`
}
type Status struct {
	Status string `json:"status" db:"status" binding:"required"`
}
