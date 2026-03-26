package filetreehandler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	filetreehandler "go_sample_code/internal/handler/filetreehandler"
	filetreeservice "go_sample_code/internal/service/filetree"
	"go_sample_code/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestApp(t *testing.T) *fiber.App {
	t.Helper()

	log, err := logger.NewLogger(logger.DefaultConfig())
	require.NoError(t, err)

	svc, err := filetreeservice.NewService(20)
	require.NoError(t, err)
	h := filetreehandler.NewHandler(log, svc)

	app := fiber.New()
	api := app.Group("/api/filetree")
	api.Post("/node", h.AddNode)
	api.Delete("/node", h.RemoveNode)
	api.Put("/rename", h.RenameNode)
	api.Put("/move", h.MoveNode)
	api.Get("/files", h.GetAllFiles)
	api.Get("/tree", h.GetTree)

	return app
}

func setRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-ID", "test-user")
}

func parseBody(t *testing.T, body io.Reader) map[string]any {
	t.Helper()

	var out map[string]any
	err := json.NewDecoder(body).Decode(&out)
	require.NoError(t, err)
	return out
}

func TestAddNodeAndQueryEndpoints(t *testing.T) {
	app := newTestApp(t)

	req := httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/","name":"docs","is_dir":true}`))
	setRequestHeaders(req)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/docs","name":"readme.md","is_dir":false,"file_id":1001}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("GET", "/api/filetree/files", nil)
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	body := parseBody(t, resp.Body)
	files, ok := body["data"].([]any)
	require.True(t, ok)
	require.Len(t, files, 1)

	entry, ok := files[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "/docs/readme.md", entry["path"])
	assert.Equal(t, float64(1001), entry["file_id"])

	req = httptest.NewRequest("GET", "/api/filetree/tree", nil)
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	treeBody := parseBody(t, resp.Body)
	tree, ok := treeBody["data"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "root", tree["name"])
	assert.Equal(t, true, tree["is_dir"])
}

func TestAddNodeValidationError(t *testing.T) {
	app := newTestApp(t)

	req := httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/","name":"readme.md","is_dir":false}`))
	setRequestHeaders(req)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body := parseBody(t, resp.Body)
	assert.Equal(t, float64(1003), body["code"])
	assert.Equal(t, "Request Validate Error", body["message"])
	assert.Equal(t, "file_id must be greater than 0", body["data"])
}

func TestAddNodeInvalidBody(t *testing.T) {
	app := newTestApp(t)

	req := httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":`))
	setRequestHeaders(req)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body := parseBody(t, resp.Body)
	assert.Equal(t, float64(1002), body["code"])
	assert.Equal(t, "Request Parser Error", body["message"])
	assert.NotNil(t, body["data"])
}

func TestRenameAndRemoveEndpoints(t *testing.T) {
	app := newTestApp(t)

	req := httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/","name":"docs","is_dir":true}`))
	setRequestHeaders(req)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/docs","name":"readme.md","is_dir":false,"file_id":1001}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("POST", "/api/filetree/node", strings.NewReader(`{"parent_path":"/","name":"archive","is_dir":true}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("PUT", "/api/filetree/rename", strings.NewReader(`{"old_path":"/docs/readme.md","new_name":"README.md"}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("PUT", "/api/filetree/move", strings.NewReader(`{"source_path":"/docs/README.md","target_dir_path":"/archive"}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("DELETE", "/api/filetree/node", strings.NewReader(`{"path":"/archive/README.md"}`))
	setRequestHeaders(req)
	resp, err = app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestMoveNodeValidationError(t *testing.T) {
	app := newTestApp(t)

	req := httptest.NewRequest("PUT", "/api/filetree/move", strings.NewReader(`{"source_path":"/missing.md","target_dir_path":"/docs"}`))
	setRequestHeaders(req)
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body := parseBody(t, resp.Body)
	assert.Equal(t, float64(1003), body["code"])
	assert.Equal(t, "Request Validate Error", body["message"])
	assert.NotNil(t, body["data"])
}
