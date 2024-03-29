package authhttp

import (
	"Test_derictory/auth"
	"Test_derictory/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

type signInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {

	if c.Request.Method != "POST" {
		newErrorResponse(c, http.StatusMethodNotAllowed, "ForbiddenMethod")
		return
	}

	var input models.User2

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Info("тут 1")
		return
	}

	_, err := h.useCase.SignUp(c.Request.Context(), input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(303, "/auth/cong")

}

func (h *Handler) SignIn(c *gin.Context) {

	if c.Request.Method != "POST" {
		newErrorResponse(c, http.StatusMethodNotAllowed, "ForbiddenMethod")
	}

	var inp signInput

	if err := c.Bind(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.useCase.SignIn(c.Request.Context(), inp.Username, inp.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

<<<<<<< HEAD
	token, id, err := h.useCase.CreateTokens(c.Request.Context(), user.Username, user.Id) //Получаем токен

	saveErr := h.useCase.CreateAuth(c.Request.Context(), id, token) //Создаем запись в Redis по данному id
=======
	token, data, err := h.useCase.CreateTokens(c.Request.Context(), user)

	saveErr := h.useCase.CreateAuth(c.Request.Context(), data.Id, token)
>>>>>>> 63974d519d261fbf92fc379db93957af7697fe9a
	if saveErr != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.SetCookie("AccessToken", token.AccessToken, 60*60*24, "/", "localhost", false, true)
	c.SetCookie("RefreshToken", token.RefreshToken, 60*60*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})

	c.Redirect(303, "/api/home")

}
