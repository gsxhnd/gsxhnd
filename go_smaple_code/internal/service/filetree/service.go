package filetree

import (
	"context"
	"sync"

	"go_sample_code/internal/errno"
	pkgtree "go_sample_code/pkg/filetree"
)

type Service interface {
	AddNode(ctx context.Context, parentPath, name string, isDir bool, fileID uint64) (*pkgtree.FileNode, error)
	RemoveNode(ctx context.Context, path string) (*pkgtree.FileNode, error)
	RenameNode(ctx context.Context, oldPath, newName string) (*pkgtree.FileNode, error)
	GetAllFiles(ctx context.Context) ([]pkgtree.FileEntry, error)
	GetTree(ctx context.Context) (*pkgtree.FileNode, error)
}

type service struct {
	mu   sync.RWMutex
	tree string
}

func NewService() Service {
	return &service{
		tree: `{"name":"root","is_dir":true,"children":[]}`,
	}
}

func (s *service) AddNode(_ context.Context, parentPath, name string, isDir bool, fileID uint64) (*pkgtree.FileNode, error) {
	if parentPath == "" {
		parentPath = "/"
	}
	if name == "" {
		return nil, errno.RequestValidateError.WithData("name cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	tree, err := pkgtree.Deserialize(s.tree)
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
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

	serialized, err := tree.Serialize()
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}
	s.tree = serialized

	return tree, nil
}

func (s *service) RemoveNode(_ context.Context, path string) (*pkgtree.FileNode, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tree, err := pkgtree.Deserialize(s.tree)
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}

	if err := tree.RemoveNode(path); err != nil {
		return nil, errno.RequestValidateError.WithData(err.Error())
	}

	serialized, err := tree.Serialize()
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}
	s.tree = serialized

	return tree, nil
}

func (s *service) RenameNode(_ context.Context, oldPath, newName string) (*pkgtree.FileNode, error) {
	if newName == "" {
		return nil, errno.RequestValidateError.WithData("new_name cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	tree, err := pkgtree.Deserialize(s.tree)
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}

	if err := tree.RenameNode(oldPath, newName); err != nil {
		return nil, errno.RequestValidateError.WithData(err.Error())
	}

	serialized, err := tree.Serialize()
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}
	s.tree = serialized

	return tree, nil
}

func (s *service) GetAllFiles(ctx context.Context) ([]pkgtree.FileEntry, error) {
	tree, err := s.GetTree(ctx)
	if err != nil {
		return nil, err
	}
	return tree.GetAllFiles(), nil
}

func (s *service) GetTree(_ context.Context) (*pkgtree.FileNode, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tree, err := pkgtree.Deserialize(s.tree)
	if err != nil {
		return nil, errno.DataConversionError.WithData(err.Error())
	}
	return tree, nil
}
