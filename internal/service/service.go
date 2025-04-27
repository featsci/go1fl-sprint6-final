package service

import (
	"strings"

	"github.com/featsci/go1fl-sprint6-final/pkg/morse"
)

func ServiceMorse(data string) (string, error) {

	isMorse := true
	t := ""
	// fmt.Println(string(data))
	for _, v := range data {
		// fmt.Println(string(v))
		if strings.Contains(string(v), ".") || strings.Contains(string(v), "-") || strings.Contains(string(v), " ") {
			continue
		} else {
			isMorse = false
			break
		}
	}
	// fmt.Println(isMorse)
	if isMorse {
		t = morse.ToText(string(data))
	} else {
		t = morse.ToMorse(string(data))
	}

	// fmt.Println(t)

	return t, nil
}
