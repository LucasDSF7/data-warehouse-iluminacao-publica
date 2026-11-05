package utils

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed env.txt
var txt string

func LoadEnv() {
	for e := range strings.SplitSeq(txt, "\r\n") {
		if e == "" {
			continue
		}
		pair := strings.SplitN(e, "=", 2)
		os.Setenv(pair[0], pair[1])
	}
}
