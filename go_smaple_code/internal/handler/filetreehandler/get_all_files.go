package filetreehandler

import "github.com/gofiber/fiber/v2"

func (h *handler) GetAllFiles(c *fiber.Ctx) error {
	files, err := h.svc.GetAllFiles(c.UserContext())
	if err != nil {
		h.log.Warn("failed to get files")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"files": files,
	})
}
