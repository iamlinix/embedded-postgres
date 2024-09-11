package embeddedpostgres

import (
	"fmt"
	"os/exec"
)

var (
	errNotSupported = fmt.Errorf("RunAsUser config parameter not supported on windows")
)

func setRunAs(*exec.Cmd, string) error {
	return errNotSupported
}

func chown(string, string) error {
	return errNotSupported
}
