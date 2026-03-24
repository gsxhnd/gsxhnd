package filetreehandler

import "github.com/gofiber/fiber/v2"

func (h *handler) GetTree(c *fiber.Ctx) error {
	tree, err := h.svc.GetTree(c.UserContext())
	if err != nil {
		h.log.Warn("failed to get tree")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(tree)
}
