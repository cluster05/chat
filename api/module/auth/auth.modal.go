package auth

type Auth struct {
	AuthId    string `bson:"authId" json:"authId"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
	CreatedAt int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt" json:"updatedAt"`
}

type AuthDTO struct {
	Username string `json:"username" binding:"required,min=2,max=15"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type ChangePasswordDTO struct {
	Username    string `json:"username" binding:"required,min=2,max=15"`
	OldPassword string `json:"oldPassword" binding:"required,min=8,max=20"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=20"`
}

type ForgotPasswordDTO struct {
	Username string `json:"username" binding:"required,min=2,max=15"`
}

type JWTUser struct {
	AuthId   string `json:"authId"`
	Username string `json:"username"`
}
