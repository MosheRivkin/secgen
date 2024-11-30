package help

func IsHelp(args []string) (bool, string) {
	if len(args) > 0 && args[len(args)-1] == "-h" {
		return true, helpText
	}
	return false, ""
}
