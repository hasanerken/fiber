package tenants

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type TenantHandler struct {
	service *TenantService
}

func NewTenantHandler(service *TenantService) *TenantHandler {
	return &TenantHandler{
		service: service,
	}
}

//func (h *TenantHandler) GetAllTenants(c *fiber.Ctx) error {
//	tenants, err := h.service.GetTenants(c.Context())
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to fetch tenants",
//		})
//	}
//	return c.JSON(tenants)
//}
//
//func (h *TenantHandler) GetTenantByID(c *fiber.Ctx) error {
//	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid tenant ID",
//		})
//	}
//
//	tenant, err := h.service.GetTenantByID(c.Context(), id)
//	if err != nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "Tenant not found",
//		})
//	}
//	return c.JSON(tenant)
//}

func (h *TenantHandler) CreateTenant(c *fiber.Ctx) error {
	fmt.Print("run", c)
	var tenantData Tenant // Implement TenantData struct to hold request data
	if err := c.BodyParser(&tenantData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	createdTenant, err := h.service.CreateTenant(c.Context(), &Tenant{
		Alias:  tenantData.Alias,
		APIKey: tenantData.APIKey,
		Status: tenantData.Status,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create tenant",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdTenant)
}
