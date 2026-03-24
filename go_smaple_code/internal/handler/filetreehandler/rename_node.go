package filetreehandler

import "github.com/gofiber/fiber/v2"

func (h *handler) RenameNode(c *fiber.Ctx) error {
	var req RenameNodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	tree, err := h.svc.RenameNode(c.UserContext(), req.OldPath, req.NewName)
	if err != nil {
		h.log.Warn("failed to rename node")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "node renamed successfully",
		"tree":    tree,
	})
}
