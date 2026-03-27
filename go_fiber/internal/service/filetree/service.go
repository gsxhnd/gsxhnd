package filetree

import (
	"context"

	"go_sample_code/internal/errno"
	pkgtree "go_sample_code/pkg/filetree"
)

type Service interface {
	AddNode(ctx context.Context, parentPath, name string, isDir bool, fileID uint64) (*pkgtree.FileNode, error)
	RemoveNode(ctx context.Context, path string) (*pkgtree.FileNode, error)
	RenameNode(ctx context.Context, oldPath, newName string) (*pkgtree.FileNode, error)
	MoveNode(ctx context.Context, sourcePath, targetDirPath string) (*pkgtree.FileNode, error)
	GetAllFiles(ctx context.Context) ([]pkgtree.FileEntry, error)
	GetTree(ctx context.Context) (*pkgtree.FileNode, error)
}

type service struct {
	maxDepth int
}

func NewService(maxDepth int) (Service, error) {
	if maxDepth <= 0 {
		return nil, errno.RequestValidateError.WithData("max depth must be greater than 0")
	}

	return &service{
		maxDepth: maxDepth,
	}, nil
}
