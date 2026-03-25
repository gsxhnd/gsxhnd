package filetree_test

import (
	"context"
	"errors"
	"testing"

	"go_sample_code/internal/errno"
	filetreeservice "go_sample_code/internal/service/filetree"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceAddNodeAndGetAllFiles(t *testing.T) {
	svc := filetreeservice.NewService()

	_, err := svc.AddNode(context.Background(), "/", "docs", true, 0)
	require.NoError(t, err)

	_, err = svc.AddNode(context.Background(), "/docs", "readme.md", false, 1001)
	require.NoError(t, err)

	files, err := svc.GetAllFiles(context.Background())
	require.NoError(t, err)
	require.Len(t, files, 1)
	assert.Equal(t, "/docs/readme.md", files[0].Path)
	assert.Equal(t, uint64(1001), files[0].FileID)
}

func TestServiceAddNodeValidation(t *testing.T) {
	svc := filetreeservice.NewService()

	_, err := svc.AddNode(context.Background(), "/", "", true, 0)
	require.Error(t, err)
	assert.Equal(t, errno.RequestValidateError.Message, err.Error())

	var eno errno.Errno
	require.True(t, errors.As(err, &eno))
	assert.Equal(t, errno.RequestValidateError.Code, eno.GetCode())
	assert.Equal(t, "name cannot be empty", eno.GetData())

	_, err = svc.AddNode(context.Background(), "/", "readme.md", false, 0)
	require.Error(t, err)
	assert.Equal(t, errno.RequestValidateError.Message, err.Error())
	require.True(t, errors.As(err, &eno))
	assert.Equal(t, errno.RequestValidateError.Code, eno.GetCode())
	assert.Equal(t, "file_id must be greater than 0", eno.GetData())
}

func TestServiceRenameAndRemoveNode(t *testing.T) {
	svc := filetreeservice.NewService()

	_, err := svc.AddNode(context.Background(), "/", "docs", true, 0)
	require.NoError(t, err)
	_, err = svc.AddNode(context.Background(), "/docs", "readme.md", false, 1001)
	require.NoError(t, err)

	_, err = svc.RenameNode(context.Background(), "/docs/readme.md", "README.md")
	require.NoError(t, err)

	tree, err := svc.GetTree(context.Background())
	require.NoError(t, err)
	require.NotNil(t, tree.FindNode("/docs/README.md"))
	require.Nil(t, tree.FindNode("/docs/readme.md"))

	_, err = svc.RemoveNode(context.Background(), "/docs/README.md")
	require.NoError(t, err)

	tree, err = svc.GetTree(context.Background())
	require.NoError(t, err)
	require.Nil(t, tree.FindNode("/docs/README.md"))
}
