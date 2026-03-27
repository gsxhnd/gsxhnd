package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RemoveNode(c *fiber.Ctx) error {
	var req RemoveNodeRequest
	if err := c.BodyParser(&req); err != nil {
		err = errno.RequestParserError.WithData(err.Error())
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	tree, err := h.svc.RemoveNode(requestContext(c), req.Path)
	if err != nil {
		h.log.Warn("failed to remove node")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(tree, nil))
}
