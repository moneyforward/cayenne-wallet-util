package env

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// ImportData ファイルから読み込む
func ImportData(fileName string) ([]string, error) {
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return nil, errors.Wrapf(err, "os.Open(%s)", fileName)
	}
	defer file.Close()

	var keys []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys, scanner.Text())
	}

	return keys, nil
}
