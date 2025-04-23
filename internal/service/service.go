package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/featsci/go1fl-sprint6-final/pkg/morse"
)

func ServiceMorse(x string) (string, error) {
	data, err := os.ReadFile(x)
	if err != nil {
		// if errors.Is(err, os.ErrNotExist) {
		// 	return "", nil
		// }
		return "", err
	}
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

	fmt.Println(t)
	// if strings.ContainsRune(data) {
	// 	fmt.Println("Yes - .")
	// }

	return t, nil
}
