package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var (
	baseFileName = ".env"
)

// WriteEnvWithMultipleData is generate .env file
func WriteEnvWithMultipleData(envData map[string][]string, isTomlBindFormat bool) (string, error) {

	fp, err := os.Create(baseFileName)
	if err != nil {
		return "", errors.Wrapf(err, "os.Create(%s)", baseFileName)
	}
	defer fp.Close()
	writer := bufio.NewWriter(fp)

	for key, value := range envData {
		row := fmt.Sprintf("export %s=%v", key, commaSeparatedToString(value, isTomlBindFormat))
		if _, err := writer.WriteString(row + "\n"); err != nil {
			return "", errors.Wrapf(err, " writer.WriteString(%s)", row)
		}
	}

	if err := writer.Flush(); err != nil {
		return "", errors.Wrapf(err, "writer.Flush(%s)", baseFileName)
	}

	return baseFileName, nil
}

func commaSeparatedToString(values []string, isTomlBindFormat bool) string {
	var str = ""
	for _, v := range values {
		if isTomlBindFormat {
			str += fmt.Sprintf("\"\\\"%v\"\\\",", v)
		} else {
			str += fmt.Sprintf("%v,", v)
		}
	}
	return strings.TrimRight(str, ",")
}
