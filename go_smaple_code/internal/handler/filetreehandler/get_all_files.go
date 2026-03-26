package filetreehandler

import (
	"go_sample_code/internal/errno"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetAllFiles(c *fiber.Ctx) error {
	files, err := h.svc.GetAllFiles(requestContext(c))
	if err != nil {
		h.log.Warn("failed to get files")
		decoded := errno.Decode(nil, err)
		return c.Status(decoded.GetHTTPStatus()).JSON(decoded)
	}

	return c.JSON(errno.Decode(files, nil))
}
