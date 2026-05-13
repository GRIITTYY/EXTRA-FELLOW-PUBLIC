package main

func ParseArgs(args []string) (string, string, error) {

	if len(args) != 2 {
		return "", "", errorMessage
	}

	return  args[0], args[1], nil
}
