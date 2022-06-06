package controllers

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"klik/dto"
	"klik/service"
	"klik/utils"
	"klik/utils/session"
	"net/http"
)

type AuthController struct {
	Controller
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (c *AuthController) Index(ctx echo.Context) error {
	return echoview.Render(ctx, http.StatusOK, "auth/login", echo.Map{
		"title":        "Login Page",
		"flashMessage": session.GetFlashMessage(ctx),
	})
}

func (c *AuthController) LoginLos(ctx echo.Context) error {
	var loginDto dto.LoginDto
	if err := ctx.Bind(&loginDto); err != nil {
		session.SetFlashMessage(ctx, "error binding data", "error", nil)
		return ctx.Redirect(302, "/klik/auth/login")
	}
	if err := ctx.Validate(&loginDto); err != nil {
		session.SetFlashMessage(ctx, "validation Error", "error", nil)
		return ctx.Redirect(302, "/klik/auth/login")
	}

	//Search Email
	data, err := c.authService.FindUserByEmail(loginDto.Email)
	if err != nil {
		session.SetFlashMessage(ctx, err.Error(), "error", nil)
		return ctx.Redirect(302, "/klik/auth/login")
	}
	if !utils.CheckPasswordHash(loginDto.Password, data.Password) {
		session.SetFlashMessage(ctx, "wrong email or password", "error", nil)
		return ctx.Redirect(302, "/klik/auth/login")
	}
	//Search Branch
	userInfo := session.UserInfo{
		UserID:     data.UserID,
		Name:       data.Name,
		Email:      data.Email,
		TypeUser:   data.TypeUser,
		UserRoleID: data.UserRoleID,
	}
	if err := session.Manager.Set(ctx, session.SessionId, &userInfo); err != nil {
		session.SetFlashMessage(ctx, err.Error(), "error", nil)
		return ctx.Redirect(302, "/klik/auth/login")
	}
	session.SetFlashMessage(ctx, "login success", "success", nil)

	return ctx.Redirect(302, "/klik/admin/home")
}

func (c *AuthController) Logout(ctx echo.Context) error {
	err := session.Manager.Delete(ctx, session.SessionId)
	if err != nil {
		session.SetFlashMessage(ctx, err.Error(), "error", nil)
		return ctx.Redirect(302, "/klik/admin/home")
	}
	session.SetFlashMessage(ctx, "logout success", "success", nil)
	return ctx.Redirect(http.StatusFound, "/klik/auth/login")
}
