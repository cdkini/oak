package helper

const (
	OakRootEnvVar = "OAK_ROOT"
	OakEditor     = "nvim"
	// Green         = "\033[32m"
	// Red           = "\033[31m"
	// Blue          = "\033[34m"
	// Reset         = "\033[0m"
	DeletedDir = ".deleted"
)

var OakDependencies = [...]string{"fd", "rg", "fzf", "bat", OakEditor}
