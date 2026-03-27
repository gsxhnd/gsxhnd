package filetree

import (
	"context"

	"go_sample_code/internal/errno"
	pkgtree "go_sample_code/pkg/filetree"
)

func (s *service) RenameNode(ctx context.Context, oldPath, newName string) (*pkgtree.FileNode, error) {
	if newName == "" {
		return nil, errno.RequestValidateError.WithData("new_name cannot be empty")
	}

	userID, tree, err := s.loadTree(ctx)
	if err != nil {
		return nil, err
	}

	if err := tree.RenameNode(oldPath, newName); err != nil {
		return nil, errno.RequestValidateError.WithData(err.Error())
	}

	if err := s.saveTree(ctx, userID, tree); err != nil {
		return nil, err
	}

	return tree, nil
}
