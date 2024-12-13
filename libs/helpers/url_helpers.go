package helpers

import (
	netUrl "net/url"
	"path"
	"strings"
)

func JoinPathToEndpoints(upstreamPath string, endpoints []string) ([]string, error) {
	urls := make([]string, 0, len(endpoints))

	for _, endpoint := range endpoints {
		parsedEndpoint, err := netUrl.Parse(endpoint)
		if err != nil {
			return urls, err
		}

		parsedEndpoint.Path = path.Join(parsedEndpoint.Path, upstreamPath)

		urls = append(urls, parsedEndpoint.String())
	}

	return urls, nil
}

func GetTargetUrl(requestUrl string, routePath string, targetUrl string) (string, error) {
	parsedRequestUrl, err := netUrl.Parse(requestUrl)
	if err != nil {
		return "", err
	}

	parsedTargetUrl, err := netUrl.Parse(targetUrl)
	if err != nil {
		return "", err
	}

	parsedTargetUrl.RawQuery = parsedRequestUrl.RawQuery

	// Wildcard
	if strings.HasSuffix(routePath, "*") {
		pathToRemove := strings.Replace(routePath, "*", "", 1)

		requestPath := strings.Replace(parsedRequestUrl.Path, pathToRemove, "", 1)

		parsedTargetUrl.Path = path.Join(parsedTargetUrl.Path, requestPath)

		return parsedTargetUrl.String(), nil
	}

	requestPath := strings.Replace(parsedRequestUrl.Path, routePath, "", 1)

	parsedTargetUrl.Path = path.Join(parsedTargetUrl.Path, requestPath)

	return parsedTargetUrl.String(), nil
}
