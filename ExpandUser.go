package ExpandUser

import (
	"os"
	"strings"
)

func ExpandUser(x string) (string, error) {
	// region TODO
	// TODO: can be cached?
	sep := string(os.PathSeparator)
	tildesep := "~" + sep
	Ltildesep := len(tildesep)
	// endregion

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if x == "~" {
		return home, nil
	} else if strings.HasPrefix(x, tildesep) {
		return home + sep + x[Ltildesep:], nil
	}
	return x, nil
}
