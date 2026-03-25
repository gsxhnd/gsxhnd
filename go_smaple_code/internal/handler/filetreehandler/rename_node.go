package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RenameNode(c *fiber.Ctx) error {
	var req RenameNodeRequest
	if err := c.BodyParser(&req); err != nil {
		err = errno.RequestParserError.WithData(err.Error())
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	tree, err := h.svc.RenameNode(c.UserContext(), req.OldPath, req.NewName)
	if err != nil {
		h.log.Warn("failed to rename node")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(tree, nil))
}
