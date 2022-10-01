package auth

import (
	"strings"
	"time"
	"web-chat/pkg/response"
	"web-chat/pkg/validation"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	registerHandler(*gin.Context)
	loginHandler(*gin.Context)
	changePasswordHandler(*gin.Context)
	forgotPasswordHandler(*gin.Context)
}

type authController struct {
	service AuthService
}

func NewAuthController(service AuthService) AuthController {
	return &authController{
		service: service,
	}
}

func trim(value string) string {
	return strings.Trim(value, " ")
}

func hashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func comparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func generateJWT(auth Auth) (string, error) {

	var jwtSecret = []byte("config.AppConfig.JWTSecret")
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authId"] = auth.AuthId
	claims["username"] = auth.Username
	claims["createdAt"] = auth.CreatedAt
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	token, err := jwtToken.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (ac *authController) registerHandler(ctx *gin.Context) {

	var authDTO AuthDTO
	if valid := validation.Bind(ctx, &authDTO); !valid {
		return
	}

	isPresent, _ := ac.service.checkAuth(ctx.Request.Context(), authDTO.Username)
	if isPresent.AuthId != "" {
		response.BadRequest(ctx, "account already exists")
		return
	}

	hashPassword, err := hashPassword(authDTO.Password)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}
	auth := Auth{
		AuthId:    ksuid.New().String(),
		Username:  trim(authDTO.Username),
		Password:  hashPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err = ac.service.register(ctx.Request.Context(), auth)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	token, err := generateJWT(auth)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}
	response.Created(ctx, token)

}
func (ac *authController) loginHandler(ctx *gin.Context) {

	var authDTO AuthDTO
	if valid := validation.Bind(ctx, &authDTO); !valid {
		return
	}

	auth := Auth{
		Username: trim(authDTO.Username),
		Password: trim(authDTO.Password),
	}

	err := ac.service.login(ctx.Request.Context(), auth)
	if err != nil {
		response.NotFound(ctx, err.Error())
		return
	}

	token, err := generateJWT(auth)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}
	response.OK(ctx, token)
}
func (ac *authController) changePasswordHandler(ctx *gin.Context) {

	var changePasswordDTO ChangePasswordDTO
	if valid := validation.Bind(ctx, &changePasswordDTO); !valid {
		return
	}

	auth, _ := ac.service.checkAuth(ctx.Request.Context(), changePasswordDTO.Username)
	if auth.AuthId == "" {
		response.BadRequest(ctx, "account not exists")
		return
	}

	err := comparePassword(auth.Password, changePasswordDTO.OldPassword)
	if err != nil {
		response.BadRequest(ctx, "invalid credentials")
		return
	}

	newPassword, err := hashPassword(changePasswordDTO.NewPassword)
	if err != nil {
		response.InternalServerError(ctx, "internal server error")
		return
	}

	auth.Password = newPassword

	err = ac.service.changePassword(ctx.Request.Context(), auth)
	if err != nil {
		response.InternalServerError(ctx, "internal server error")
		return
	}

	token, err := generateJWT(auth)
	if err != nil {
		response.InternalServerError(ctx, "internal server error")
		return
	}
	response.OK(ctx, token)

}
func (ac *authController) forgotPasswordHandler(ctx *gin.Context) {

	var forgotPasswordDTO ForgotPasswordDTO
	if valid := validation.Bind(ctx, &forgotPasswordDTO); !valid {
		return
	}

	isPresent, _ := ac.service.checkAuth(ctx.Request.Context(), forgotPasswordDTO.Username)
	if isPresent.AuthId == "" {
		response.BadRequest(ctx, "account not exists")
		return
	}

	response.OK(ctx, "reset password link send to your email. please check your email")

}
