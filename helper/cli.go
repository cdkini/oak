package helper

import (
	"errors"
	"os"
	"os/exec"
)

// TODO - Make these configurable
const (
	EDITOR = "nvim"
	GREP   = "rg"
	FIND   = "fd"
	FZF    = "fzf"
)

func CheckDependencies() error {
	for _, cmd := range []string{EDITOR, GREP, FIND} {
		if _, err := exec.LookPath(cmd); err != nil {
			return err
		}
	}
	return nil
}

func OpenNote(path string, title string, tags []string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = initNote(path, title, tags); err != nil {
			return err
		}
	}

	return openNoteInEditor(path)
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

func openNoteInEditor(path string) error {
	return runCommand(EDITOR, nil, path)
}

func Grep(root string, query string) error {
	return runCommand(GREP, &root, query)
}

func Find(root string, query string) error {
	return runCommand(FIND, &root, query)
}

func FZFOpen(root string, query string) error {
	// TODO - remove hardcoded nvim, rg, and bat
	reloadCmd := "reload:rg --column --color=always --smart-case {q} || :"
	previewCmd := "bat --style=full --color=always --highlight-line {2} {1}"
	args := []string{"--disabled", "--ansi", "--multi",
		"--bind", "start:" + reloadCmd,
		"--bind", "change:" + reloadCmd,
		"--bind", "enter:become:if [[ $FZF_SELECT_COUNT -eq 0 ]]; then nvim {1} +{2}; else nvim +cw -q {+f}; fi",
		"--delimiter", ":",
		"--preview", previewCmd,
		"--preview-window", "~4,+{2}+4/3,<80(up)",
		"--query", query,
	}
	return runCommand(FZF, &root, args...)
}

func runCommand(name string, dir *string, args ...string) error {
	cmd := exec.Command(name, args...)
	if dir != nil {
		cmd.Dir = *dir
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
