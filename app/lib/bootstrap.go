// Bootstrap file.
// Ready container.

package lib

import (
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/validation"
	"os"
	"regexp"
	"strings"
)

func Bootstrap(envPrefix string) *goldi.Container {
	registry := goldi.NewTypeRegistry()
	RegisterTypes(registry)

	config := generateConfig(envPrefix)

	container := goldi.NewContainer(registry, config)
	validator := validation.NewContainerValidator()
	validator.MustValidate(container)

	return container
}

func generateConfig(envPrefix string) map[string]interface{} {
	reg, _ := regexp.Compile("^" + envPrefix + "(.+)$")
	config := map[string]interface{}{}

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")

		if len(pair) == 2 && reg.MatchString(pair[0]) {
			// Removing prefix from the key.
			key := reg.ReplaceAllString(pair[0], "$1")
			config[key] = pair[1]
		}
	}

	return config
}
