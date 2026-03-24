package filetree

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FileNode represents a file tree node (N-ary tree node)
type FileNode struct {
	Name     string      `json:"name"`
	IsDir    bool        `json:"is_dir"`
	FileID   uint64      `json:"file_id,omitempty"`
	Children []*FileNode `json:"children"`
}

type FileEntry struct {
	Path   string `json:"path"`
	FileID uint64 `json:"file_id"`
}

// Serialize converts the N-ary tree to a JSON string
func (f *FileNode) Serialize() (string, error) {
	data, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Deserialize converts a JSON string to an N-ary tree
func Deserialize(data string) (*FileNode, error) {
	var root *FileNode
	err := json.Unmarshal([]byte(data), &root)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// FindNode finds the node at the specified path
func (f *FileNode) FindNode(path string) *FileNode {
	if path == "" || path == "/" {
		return f
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := f

	for _, part := range parts {
		found := false
		for _, child := range current.Children {
			if child.Name == part {
				current = child
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}

	return current
}

// AddFile adds a file to the specified directory (requires file ID)
func (f *FileNode) AddFile(dirPath string, fileName string, fileID uint64) error {
	parent := f.FindNode(dirPath)
	if parent == nil || !parent.IsDir {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}

	for _, child := range parent.Children {
		if child.Name == fileName {
			return fmt.Errorf("file already exists: %s", fileName)
		}
	}

	parent.Children = append(parent.Children, &FileNode{
		Name:   fileName,
		IsDir:  false,
		FileID: fileID,
	})

	return nil
}

// AddDir adds a subdirectory to the specified directory
func (f *FileNode) AddDir(parentPath, dirName string) error {
	parent := f.FindNode(parentPath)
	if parent == nil || !parent.IsDir {
		return fmt.Errorf("directory does not exist: %s", parentPath)
	}

	for _, child := range parent.Children {
		if child.Name == dirName {
			return fmt.Errorf("directory already exists: %s", dirName)
		}
	}

	parent.Children = append(parent.Children, &FileNode{
		Name:     dirName,
		IsDir:    true,
		Children: []*FileNode{},
	})

	return nil
}

// RemoveNode removes the node at the specified path (file or directory)
func (f *FileNode) RemoveNode(path string) error {
	if path == "" || path == "/" {
		return fmt.Errorf("cannot remove root directory")
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	fileName := parts[len(parts)-1]
	parentPath := strings.Join(parts[:len(parts)-1], "/")

	parent := f.FindNode(parentPath)
	if parent == nil {
		return fmt.Errorf("parent directory does not exist: %s", parentPath)
	}

	for i, child := range parent.Children {
		if child.Name == fileName {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("node does not exist: %s", path)
}

// RenameNode renames a node
func (f *FileNode) RenameNode(oldPath, newName string) error {
	node := f.FindNode(oldPath)
	if node == nil {
		return fmt.Errorf("node does not exist: %s", oldPath)
	}

	node.Name = newName
	return nil
}

// GetFileID gets the ID of a file
func (f *FileNode) GetFileID(path string) (uint64, error) {
	node := f.FindNode(path)
	if node == nil {
		return 0, fmt.Errorf("node does not exist: %s", path)
	}
	if node.IsDir {
		return 0, fmt.Errorf("not a file: %s", path)
	}
	return node.FileID, nil
}

// GetAllFiles gets all file paths and FileIDs
func (f *FileNode) GetAllFiles() []FileEntry {
	var files []FileEntry
	var traverse func(node *FileNode, path string)

	traverse = func(node *FileNode, path string) {
		if !node.IsDir {
			files = append(files, FileEntry{Path: path, FileID: node.FileID})
			return
		}

		for _, child := range node.Children {
			childPath := path + "/" + child.Name
			traverse(child, childPath)
		}
	}

	for _, child := range f.Children {
		traverse(child, "/"+child.Name)
	}

	return files
}

// GetAllDirs gets all directory paths
func (f *FileNode) GetAllDirs() []string {
	var dirs []string
	var traverse func(node *FileNode, path string)

	traverse = func(node *FileNode, path string) {
		if node.IsDir {
			dirs = append(dirs, path)
		}

		for _, child := range node.Children {
			childPath := path + "/" + child.Name
			traverse(child, childPath)
		}
	}

	for _, child := range f.Children {
		traverse(child, "/"+child.Name)
	}

	return dirs
}
