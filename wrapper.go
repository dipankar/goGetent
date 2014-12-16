package goGetent

import (
	"os/exec"
	"fmt"
)

// Make a call to getent and get an output
func get_ent(arg0 string,arg1 string) {
    app := "getent"
    cmd := exec.Command(app, arg0, arg1)
    stdout, err := cmd.Output()
    return string(stdout),err
}
