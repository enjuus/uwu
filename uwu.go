// UwU whats this
package uwu

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)
// Takes a string to be "translated" into uwu
func Translate(String string) (string, error) {
	var str bytes.Buffer
	String = strings.ToLower(String)
	if String == "" {
		return "", errors.New("uwu: string can't be empty")
	}
	for _, c := range String {
		var char string
		switch string(c) {
		case "l":
			char = "w"
		case "r":
			char = "w"
		default:
			char = string(c)
		}
		str.WriteString(char)
	}

	elements := map[int]string{
		0: "na",
		1: "ne",
		2: "ni",
		3: "no",
		4: "nu",
		5: "ma",
		6: "me",
		7: "mi",
		8: "mo",
		9: "mu",
	}

	returnString := str.String()
	index := 1
	for _, bit := range elements {
		returnString = strings.ReplaceAll(returnString, bit, bit[:index]+"y"+bit[index:])
	}

	returnString = fmt.Sprintf("%s uwu", returnString)

	return returnString, nil
}