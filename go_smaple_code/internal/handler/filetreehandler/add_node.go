package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) AddNode(c *fiber.Ctx) error {
	var req AddNodeRequest
	if err := c.BodyParser(&req); err != nil {
		err = errno.RequestParserError.WithData(err.Error())
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	tree, err := h.svc.AddNode(requestContext(c), req.ParentPath, req.Name, req.IsDir, req.FileID)
	if err != nil {
		h.log.Warn("failed to add node")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(tree, nil))
}
