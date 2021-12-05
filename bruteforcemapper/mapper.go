package bruteforcemapper

import "strings"

const domainName = "https://domain.com/"

type Mapper struct {
	Controllers []string
}

func NewMapper(c []string) *Mapper {
	return &Mapper{Controllers: c}
}

func (m *Mapper) IsValidController(url string) bool {
	var currentController string
	trimURL := url[len(domainName):]
	for _, controller := range m.Controllers {
		if strings.HasPrefix(trimURL, controller) {
			if currentController == "" || len(currentController) < len(controller) {
				currentController = controller
			}
		}
	}

	return currentController != ""
}
