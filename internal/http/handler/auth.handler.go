package handler

import (
	"net/http"

	"github.com/Sandhya-Pratama/weather-app/internal/http/validator"
	"github.com/Sandhya-Pratama/weather-app/internal/service"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginService service.LoginUseCase
	tokenService service.TokenUseCase
}

func NewAuthHandler(
	loginService service.LoginUseCase,
	tokenService service.TokenUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginService: loginService,
		tokenService: tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	data := map[string]string{
		"access_token": accessToken,
	}

	sess, _ := session.Get("auth-sessions", ctx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
	}
	sess.Values["token"] = accessToken
	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, data)
}
