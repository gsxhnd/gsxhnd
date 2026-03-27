package filetreehandler

type AddNodeRequest struct {
	ParentPath string `json:"parent_path"`
	Name       string `json:"name"`
	IsDir      bool   `json:"is_dir"`
	FileID     uint64 `json:"file_id,omitempty"`
}

type RemoveNodeRequest struct {
	Path string `json:"path"`
}

type RenameNodeRequest struct {
	OldPath string `json:"old_path"`
	NewName string `json:"new_name"`
}

type MoveNodeRequest struct {
	SourcePath    string `json:"source_path"`
	TargetDirPath string `json:"target_dir_path"`
}
