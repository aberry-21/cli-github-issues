package editor

import (
	"fmt"
	"os"
	"os/exec"
)

type Editor struct {
	editorPath string
}

// NewEditor creates a new instance of the Editor with name of editor program.
func NewEditor(name string) (*Editor, error) {
	const op = "editor.NewEditor"

	// find editor program in PATH
	path, err := exec.LookPath(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Editor{path}, nil
}

// OpenFile opens file in editor
func (e *Editor) OpenFile(path string) error {
	const op = "editor.OpenFile"

	// create file
	cmd := exec.Command(e.editorPath, path)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// SaveInputDataToFile saves input data in temporary file and returns its data.
func SaveInputDataToFile(editorName string) ([]byte, error) {
	const op = "editor.SaveInputDataToFile"

	// create editor
	editor, err := NewEditor(editorName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// create temporary file
	f, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// clean up
	defer os.Remove(f.Name())
	defer f.Close()

	// run editor with temporary file
	if err := editor.OpenFile(f.Name()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// read temporary file after editor program closed
	data, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}
