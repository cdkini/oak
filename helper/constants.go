package helper

const (
	FZF           = "fzf"
	FindTool      = "fd"
	GrepTool      = "rg"
	PreviewTool   = "bat"
	OakEditor     = "nvim"
	OakRootEnvVar = "OAK_ROOT"
	DeletedDir    = ".deleted"
)

var OakDependencies = [...]string{FZF, FindTool, GrepTool, PreviewTool, OakEditor}
