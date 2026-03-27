package filetree

import (
	"context"

	pkgtree "go_sample_code/pkg/filetree"
)

func (s *service) GetTree(ctx context.Context) (*pkgtree.FileNode, error) {
	_, tree, err := s.loadTree(ctx)
	if err != nil {
		return nil, err
	}
	return tree, nil
}
