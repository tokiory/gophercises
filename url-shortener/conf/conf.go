package conf

import (
	"fmt"
	"os"
	"strings"
)

func GetConfExtension(file *os.File) (ConfigExtension, error) {
	filename := file.Name()
	extensions := make([]string, 0, len(ConfigExtensionHashmap))

	for k, v := range ConfigExtensionHashmap {
		extensions = append(extensions, k)
		if strings.HasSuffix(filename, k) {
			return v, nil
		}
	}

	return -1, fmt.Errorf("file extension must be one of following %v", extensions)
}
