package filetree

import (
	"context"

	"go_sample_code/internal/errno"
	pkgtree "go_sample_code/pkg/filetree"
)

func (s *service) RemoveNode(ctx context.Context, path string) (*pkgtree.FileNode, error) {
	userID, tree, err := s.loadTree(ctx)
	if err != nil {
		return nil, err
	}

	if err := tree.RemoveNode(path); err != nil {
		return nil, errno.RequestValidateError.WithData(err.Error())
	}

	if err := s.saveTree(ctx, userID, tree); err != nil {
		return nil, err
	}

	return tree, nil
}
