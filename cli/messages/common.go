package messages

const (
	PROMPT_DIRECTORY  = "Enter starting directory: "
	INVALID_DIRECTORY = "directory %s is invalid, please re-enter the directory"

	PROMPT_OPTION      = "Choose an option: "
	LIST_OPTIONS       = LIST_OPTION_PARENT + LIST_OPTION_SUB + LIST_OPTION_ELSE
	LIST_OPTION_PARENT = ".. - Back to parent directory\n"
	LIST_OPTION_SUB    = "INDEX - Enter sub directory\n"
	LIST_OPTION_ELSE   = "exit | bye | c - Exit\n"
)
