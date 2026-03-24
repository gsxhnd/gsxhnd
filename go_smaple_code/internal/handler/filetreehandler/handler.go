package filetreehandler

import (
	filetreeservice "go_sample_code/internal/service/filetree"
	"go_sample_code/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	AddNode(c *fiber.Ctx) error
	RemoveNode(c *fiber.Ctx) error
	RenameNode(c *fiber.Ctx) error
	GetAllFiles(c *fiber.Ctx) error
	GetTree(c *fiber.Ctx) error
}

type handler struct {
	log logger.Logger
	svc filetreeservice.Service
}

func NewHandler(log logger.Logger, svc filetreeservice.Service) Handler {
	return &handler{log: log, svc: svc}
}
