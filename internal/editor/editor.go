package editor

import (
	"fmt"
	"os"
	"os/exec"
)

type Editor struct {
	editorPath string
}

func NewEditor(name string) (*Editor, error) {
	const op = "editor.NewEditor"

	path, err := exec.LookPath(name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Editor{path}, nil
}

func (e *Editor) OpenFile(path string) error {
	const op = "editor.OpenFile"

	cmd := exec.Command(e.editorPath, path)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func SaveInputDataToFile(editorName string) ([]byte, error) {
	const op = "editor.SaveInputDataToFile"

	editor, err := NewEditor(editorName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	f, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// clean up
	defer os.Remove(f.Name())
	defer f.Close()

	if err := editor.OpenFile(f.Name()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return data, nil
}
