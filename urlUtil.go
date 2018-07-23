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
	cleanQueryString := regexp.MustCompile("\\?")
	queryString = cleanQueryString.ReplaceAllString(queryString, "")

	queryParamsStrings := strings.Split(queryString, "&")
	queryParams := map[string]string{}

	for _, queryString := range queryParamsStrings {
		kvp := strings.Split(queryString, "=")
		k, v := kvp[0], kvp[1]
		queryParams[k] = v
	}
	return queryParams
}