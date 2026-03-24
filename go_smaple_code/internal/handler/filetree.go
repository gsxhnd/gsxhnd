package handler

import (
	"go_sample_code/pkg/filetree"

	"github.com/gofiber/fiber/v2"
)

type FileTreeHandler struct {
	tree *filetree.FileNode
}

func NewFileTreeHandler() *FileTreeHandler {
	return &FileTreeHandler{
		tree: &filetree.FileNode{
			Name:     "root",
			IsDir:    true,
			Children: []*filetree.FileNode{},
		},
	}
}

type AddFileRequest struct {
	DirPath  string `json:"dir_path"`
	FileName string `json:"file_name"`
	FileID   uint64 `json:"file_id"`
}

type AddDirRequest struct {
	ParentPath string `json:"parent_path"`
	DirName    string `json:"dir_name"`
}

type RemoveNodeRequest struct {
	Path string `json:"path"`
}

type RenameNodeRequest struct {
	OldPath string `json:"old_path"`
	NewName string `json:"new_name"`
}

func (h *FileTreeHandler) AddFile(c *fiber.Ctx) error {
	var req AddFileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.tree.AddFile(req.DirPath, req.FileName, req.FileID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "file added successfully",
	})
}

func (h *FileTreeHandler) AddDir(c *fiber.Ctx) error {
	var req AddDirRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.tree.AddDir(req.ParentPath, req.DirName); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "directory added successfully",
	})
}

func (h *FileTreeHandler) RemoveNode(c *fiber.Ctx) error {
	var req RemoveNodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.tree.RemoveNode(req.Path); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "node removed successfully",
	})
}

func (h *FileTreeHandler) RenameNode(c *fiber.Ctx) error {
	var req RenameNodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.tree.RenameNode(req.OldPath, req.NewName); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "node renamed successfully",
	})
}

func (h *FileTreeHandler) GetAllFiles(c *fiber.Ctx) error {
	files := h.tree.GetAllFiles()
	return c.JSON(fiber.Map{
		"files": files,
	})
}

func (h *FileTreeHandler) GetTree(c *fiber.Ctx) error {
	return c.JSON(h.tree)
}
