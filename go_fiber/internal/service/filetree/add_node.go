package filetree

import (
	"context"

	"go_sample_code/internal/errno"
	pkgtree "go_sample_code/pkg/filetree"
)

func (s *service) AddNode(ctx context.Context, parentPath, name string, isDir bool, fileID uint64) (*pkgtree.FileNode, error) {
	if parentPath == "" {
		parentPath = "/"
	}
	if name == "" {
		return nil, errno.RequestValidateError.WithData("name cannot be empty")
	}

	userID, tree, err := s.loadTree(ctx)
	if err != nil {
		return nil, err
	}

	if isDir {
		err = tree.AddDir(parentPath, name)
	} else {
		if fileID == 0 {
			return nil, errno.RequestValidateError.WithData("file_id must be greater than 0")
		}
		err = tree.AddFile(parentPath, name, fileID)
	}
	if err != nil {
		return nil, errno.RequestValidateError.WithData(err.Error())
	}

	if err := s.saveTree(ctx, userID, tree); err != nil {
		return nil, err
	}

	return tree, nil
}
