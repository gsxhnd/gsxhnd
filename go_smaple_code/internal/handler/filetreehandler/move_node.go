package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) MoveNode(c *fiber.Ctx) error {
	var req MoveNodeRequest
	if err := c.BodyParser(&req); err != nil {
		err = errno.RequestParserError.WithData(err.Error())
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	tree, err := h.svc.MoveNode(requestContext(c), req.SourcePath, req.TargetDirPath)
	if err != nil {
		h.log.Warn("failed to move node")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(tree, nil))
}
