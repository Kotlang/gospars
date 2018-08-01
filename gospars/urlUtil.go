package gospars

import (
	"regexp"
	"strings"
)

func GetHashPath(path string) string {
	// remove #/ from path
	cleanPath := regexp.MustCompile("#")
	path = cleanPath.ReplaceAllString(path, "")

	return path
}

func MatchPathAndGetPathParams(configPath string, locationPath string) (bool, map[string]string) {
	configPathParams := strings.Split(configPath, "/")
	locationPathParams := strings.Split(locationPath, "/")

	if len(configPathParams) != len(locationPathParams) {
		return false, map[string]string {}
	}

	pathParamsMap := map[string]string {}
	for i, configPathParam := range configPathParams {
		if strings.HasPrefix(configPathParam, ":") {
			pathParamsMap[configPathParam] = locationPathParams[i]
			continue
		} else if configPathParam != locationPathParams[i] {
			return false, map[string]string {}
		}
	}
	return true, pathParamsMap
}

func GetQueryParams(queryString string) map[string]string {
	queryString = strings.Trim(queryString, " ")
	queryParams := map[string]string{}

	if len(queryString) == 0 {
		return queryParams
	}
	cleanQueryString := regexp.MustCompile("\\?")
	queryString = cleanQueryString.ReplaceAllString(queryString, "")

	queryParamsStrings := strings.Split(queryString, "&")

	for _, queryString := range queryParamsStrings {
		kvp := strings.Split(queryString, "=")
		var k, v string

		if len(kvp) == 0 {
			continue
		} else if len(kvp) == 1 {
			k, v = kvp[0], kvp[0]
		} else {
			k, v = kvp[0], kvp[1]
		}
		queryParams[k] = v
	}
	return queryParams
}