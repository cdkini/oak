package helper

import (
	"errors"
	"os"
	"os/exec"
)

const (
	editor = "nvim"
)

func OpenNote(path string, title string, tags []string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = initNote(path, title, tags); err != nil {
			return err
		}
	}

	return openNote(path)
}

func initNote(path string, title string, tags []string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fileMetadata := NewFileMetadata(title, tags)
	_, err = f.WriteString(fileMetadata.Render())
	return err
}

func openNote(path string) error {
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
