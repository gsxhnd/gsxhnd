package filetreehandler

import "github.com/gofiber/fiber/v2"

func (h *handler) AddNode(c *fiber.Ctx) error {
	var req AddNodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	tree, err := h.svc.AddNode(c.UserContext(), req.ParentPath, req.Name, req.IsDir, req.FileID)
	if err != nil {
		h.log.Warn("failed to add node")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "node added successfully",
		"tree":    tree,
	})
}
