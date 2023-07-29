package tenants

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"xhantos/common"
)

type TenantHandler struct {
	service *TenantService
}

func NewTenantHandler(service *TenantService) *TenantHandler {
	return &TenantHandler{
		service: service,
	}
}

func (h *TenantHandler) CreateTenant(ctx *fiber.Ctx) error {
	var tenantData Tenant // Implement TenantData struct to hold request data
	if err := ctx.BodyParser(&tenantData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	err := common.ValidateStruct(tenantData)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdTenant, err := h.service.CreateTenant(ctx.Context(), &Tenant{
		Alias:  tenantData.Alias,
		APIKey: tenantData.APIKey,
		Status: tenantData.Status,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create tenant",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdTenant)
}

func (h *TenantHandler) UpdateTenant(ctx *fiber.Ctx) error {
	// Parse the request body to get the updated tenant data
	var updatedTenant Tenant
	if err := ctx.BodyParser(&updatedTenant); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := common.ValidateStruct(updatedTenant)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.service.UpdateTenant(ctx.Context(), &updatedTenant)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update tenant",
		})
	}

	// Return the updated tenant as the response
	return ctx.Status(fiber.StatusOK).JSON(updatedTenant)
}

func (h *TenantHandler) GetTenantByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tenant ID",
		})
	}

	tenant, err := h.service.GetTenantByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tenant not found",
		})
	}
	return ctx.JSON(tenant)
}

func (h *TenantHandler) DeleteTenant(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tenant ID",
		})
	}

	err = h.service.DeleteTenant(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Tenant not found",
		})
	}
	return ctx.JSON(fiber.Map{"message": "successfully deleted"})
}

func (h *TenantHandler) GetTenantsByQuery(ctx *fiber.Ctx) error {
	queryParams := ctx.Locals("queryParams").(common.QueryParams)

	tenants, err := h.service.GetTenantsByQuery(ctx.Context(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tenants",
		})
	}
	return ctx.JSON(tenants)
}
