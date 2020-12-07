package container

import (
	"github.com/fgrosse/goldi"

	SharedConfig "github.com/time-off-system/shared/config"
	SharedDatabase "github.com/time-off-system/shared/database"
	SharedMiddleware "github.com/time-off-system/shared/middleware"
)

// GetDefaultContainer ...
func GetDefaultContainer() *goldi.Container {
	registry := goldi.NewTypeRegistry()
	config := make(map[string]interface{})
	container := goldi.NewContainer(registry, config)

	// - register dependency
	container.RegisterType("shared.middleware", SharedMiddleware.NewMiddleware)
	container.RegisterType("shared.config", SharedConfig.GetDefaultImmutableConfig)
	container.RegisterType("shared.database", SharedDatabase.NewMysql, "@shared.config")

	return container
}
