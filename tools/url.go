package tools

import "regexp"

var URL_REG = regexp.MustCompile("(http|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")

func IsValidUrl(url string) bool {
	if len(url) == 0 {
		return false
	}
	re := URL_REG.FindAllStringSubmatch(url, -1)
	if re != nil {
		return true
	}

	return false
}
