package utils

import "runtime"

type constants struct {
	Newline string
}

var Constants = constants{}

func init() {
	switch runtime.GOOS {
	case "windows":
		Constants.Newline = "\r\n"
		break
	default:
		Constants.Newline = "\n"
	}
}
