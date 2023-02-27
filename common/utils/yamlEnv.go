package utils

import (
	"os"
	"strings"
)

func ParseYamlEnv(yaml string) string {
	var result string
	for {
		f := strings.Index(yaml, "${")
		l := strings.Index(yaml, "}")
		if f == -1 || l == -1 {
			result += yaml
			break
		}
		result += yaml[:f]
		s := yaml[f+2 : l]
		yaml = yaml[l+1:]
		ss := strings.Split(s, ":")
		env := os.Getenv(ss[0])
		if env == "" {
			ss[0] = ss[1]
		} else {
			ss[0] = env
		}
		result += ss[0]
	}
	return result
}
