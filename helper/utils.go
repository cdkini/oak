package helper

import (
	"errors"
	"os"
	"os/exec"
)

const (
	oakRootEnv = "OAK_ROOT"
	editor     = "nvim"
)

func GetOakRoot() string {
	root := os.Getenv(oakRootEnv)
	if root == "" {
		os.Exit(1)
	}
	if stat, err := os.Stat(root); err != nil || !stat.IsDir() {
		os.Exit(1)
	}

	return root
}

func OpenNote(path string, title string, tags []string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = initNote(path, title, tags); err != nil {
			return err
		}
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func initNote(path string, title string, tags []string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	frontmatter := NewFrontmatter(title, tags)
	_, err = f.WriteString(frontmatter.Render())
	return err
}
