package controller_v1

import (
	"book/models"
	"book/repositories"
	"book/security"
	"book/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRepo repositories.UserRepo
}

func (controller *UserController) SignIn(c *gin.Context) {

	var account models.Account

	if err := c.ShouldBindBodyWithJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	if !utils.ValidateEmail(account.Email) || utils.IsEmpty(account.Password) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Invalid sign in information"))
		return
	}

	var user models.User
	if err := controller.UserRepo.GetByEmail(account.Email, &user); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse(err.Error()))
		return
	}

	accessToken, err := security.NewAccessToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
		return
	}

	refreshToken, err := security.NewRefreshToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
		return
	}

	userInfo := models.UserInfo{
		Name:         user.Name,
		Email:        user.Email,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, models.DataAndMessageResponse(userInfo, "Login successful"))

}

func (controller *UserController) SignUp(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	if !utils.ValidateEmail(user.Email) || utils.IsEmpty(user.Password) || utils.IsEmpty(user.Name) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Invalid sign up information"))
		return
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(err.Error()))
		return
	}

	if err := user.GenarateID(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
		return
	}

	if err := controller.UserRepo.Save(&user); err != nil {
		c.JSON(http.StatusConflict, models.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.MessageResponse("Account created successfully"))

}

func (controller *UserController) RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token invalid"))
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	jwtMapClaims, err := security.ParseRefreshToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token invalid"))
		return
	}

	exp, err := jwtMapClaims.GetExpirationTime()

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token invalid"))
		return
	}

	if exp == nil || exp.UTC().Unix() < time.Now().UTC().Unix() {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Refesh token expired"))
		return
	}

	userId, err := jwtMapClaims.GetSubject()
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token invalid"))
		return
	}

	var user models.User
	if err := controller.UserRepo.GetById(userId, &user); err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse("Token invalid"))
		return
	}

	accessToken, err := security.NewAccessToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(err.Error()))
		return
	}

	userInfo := models.UserInfo{
		Name:         "",
		Email:        "",
		Token:        accessToken,
		RefreshToken: "",
	}

	c.JSON(http.StatusOK, models.DataAndMessageResponse(userInfo, "Refresh token successful"))

}
