package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenNote(root string, title string, tags []string) {
	f := initFile(root, title, tags)

	cmd := exec.Command(OakEditor, f)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		Error("Something went wrong when opening file '%s'\n", f)
	}
}

func initFile(root string, title string, tags []string) string {
	path := fmt.Sprintf("%s/%s.md", root, title)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			Error("Something went wrong when creating file '%s'\n", path)
		}
		defer f.Close()

		frontmatter := NewFrontmatter(title, tags)
		if _, err := f.WriteString(frontmatter.Render()); err != nil {
			Error("Something went wrong when writing to file '%s'\n", path)
		}
	}

	return path
}

func FzfOpen(query, reloadCmd, previewCmd string) {}

func FzfSelect(query, reloadCmd, previewCmd string) {}

func Rg(query string) {
	cmd := exec.Command("rg", "--color=always", "--line-number", "--no-heading", "--smart-case", query)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		Error("Something went wrong when executing command '%s'\n", query)
	}
}

func Fd(query string) {
	cmd := exec.Command("fd", "--color=always", "--type=f", "--hidden", "--exclude=.git", query)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		Error("Something went wrong when executing command '%s'\n", query)
	}

}

func GetOakRoot() string {
	for _, dep := range OakDependencies {
		if _, err := exec.LookPath(dep); err != nil {
			Error("Could not find dependency '%s'\n", dep)
		}
	}

	root, exists := os.LookupEnv(OakRootEnvVar)
	if !exists {
		Error("Root not set!\n")
	}

	info, err := os.Stat(root)
	if err != nil {
		Error("Root '%s' does not exist!\n", root)
	}

	if !info.IsDir() {
		Error("Root '%s' is not a valid directory!\n", root)
	}

	return root
}

func Error(format string, a ...any) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
