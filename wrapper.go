package goGetent

import (
)

func main() {
	out, err := exec.Command("sh","-c","ls").Output()
}
