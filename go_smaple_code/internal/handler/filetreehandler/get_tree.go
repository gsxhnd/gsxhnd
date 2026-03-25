package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetTree(c *fiber.Ctx) error {
	tree, err := h.svc.GetTree(c.UserContext())
	if err != nil {
		h.log.Warn("failed to get tree")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(tree, nil))
}
