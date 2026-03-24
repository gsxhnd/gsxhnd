package filetreehandler

import "github.com/gofiber/fiber/v2"

func (h *handler) RemoveNode(c *fiber.Ctx) error {
	var req RemoveNodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	tree, err := h.svc.RemoveNode(c.UserContext(), req.Path)
	if err != nil {
		h.log.Warn("failed to remove node")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "node removed successfully",
		"tree":    tree,
	})
}
