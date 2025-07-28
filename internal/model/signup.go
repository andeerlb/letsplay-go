package model

type SignUpAuthServer struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignUp struct {
	SignUpAuthServer
	PreferredSport GameInfo   `json:"preferredSport" binding:"required"`
	OtherSports    []GameInfo `json:"otherSports"`
}

type GameInfo struct {
	Type     string `json:"type" binding:"required"`
	Position string `json:"position" binding:"required"`
}
