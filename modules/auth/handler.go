package auth

import (
	"fmt"
	"github.com/authorizerdev/authorizer-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"xhantos/common"
)

type AuthHandler struct {
	service    *AuthService
	authorizer *authorizer.AuthorizerClient
	db         *sqlx.DB
}

func NewAuthHandler(service *AuthService, authorizer *authorizer.AuthorizerClient, db *sqlx.DB) *AuthHandler {
	return &AuthHandler{
		service:    service,
		authorizer: authorizer,
		db:         db,
	}
}

func (h AuthHandler) Login(ctx *fiber.Ctx) error {
	loginInput := new(authorizer.LoginInput)

	if err := ctx.BodyParser(loginInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err})
	}

	err := common.ValidateStruct(loginInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	login, err := h.authorizer.Login(loginInput)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": login.AccessToken, "user": login.User})
}

type RegisterInput struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required" `
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	Tenant          string `json:"tenant" validate:"required"`
}

func (h AuthHandler) Register(ctx *fiber.Ctx) error {
	registerInput := new(RegisterInput)
	if err := ctx.BodyParser(registerInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err})
	}

	err := common.ValidateStruct(registerInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := h.authorizer.SignUp(&authorizer.SignUpInput{
		Email:           registerInput.Email,
		Password:        registerInput.Password,
		ConfirmPassword: registerInput.ConfirmPassword,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.service.setTenant(ctx.Context(), res.User.ID, registerInput.Tenant)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	fmt.Println("Tenant update successful.")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"result": res})
}

func (h AuthHandler) ForgotPassword(ctx *fiber.Ctx) error {
	forgotPasswordInput := new(authorizer.ForgotPasswordInput)
	if err := ctx.BodyParser(forgotPasswordInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err})
	}
	password, err := h.authorizer.ForgotPassword(forgotPasswordInput)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": password.Message})
}
