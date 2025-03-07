package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenNote(f string, tags []string) {

	cmd := exec.Command(OakEditor, initFile(f, tags))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		Error("Something went wrong when opening file '%s'\n", f)
	}
}

func initFile(f string, tags []string) string {
	return f
}

func FzfOpen(query, reloadCmd, previewCmd string) {}

func FzfSelect(query, reloadCmd, previewCmd string) {}

func Rg(query string) {}

func Fd(query string) {}

func InitProj() string {
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
