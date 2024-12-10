package dependencies

import (
	"fmt"
	"sync"

	"github.com/saulova/seam/libs/interfaces"
)

var lock = &sync.Mutex{}

type DependencyContainer struct {
	dependencies map[string]interface{}
	logger       interfaces.LoggerInterface
}

var dependencyContainerInstance *DependencyContainer

func GetDependencyContainer() *DependencyContainer {
	if dependencyContainerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if dependencyContainerInstance == nil {
			dependencyContainerInstance = &DependencyContainer{
				dependencies: make(map[string]interface{}),
				logger:       nil,
			}
		}
	}

	return dependencyContainerInstance
}

func SetDependencyContainer(dependencyContainer *DependencyContainer) {
	lock.Lock()
	defer lock.Unlock()

	dependencyContainerInstance = dependencyContainer
}

func (c *DependencyContainer) SetLogger(logger interfaces.LoggerInterface) {
	c.logger = logger
}

func (c *DependencyContainer) AddDependency(dependencyId string, value interface{}) {
	if c.logger != nil {
		c.logger.Debug("adding dependency", "id", dependencyId)
	}

	c.dependencies[dependencyId] = value

	if c.logger != nil {
		c.logger.Debug("dependency added")
	}
}

func (c *DependencyContainer) HasDependency(dependencyId string) bool {
	return c.dependencies[dependencyId] != nil
}

func (c *DependencyContainer) GetDependency(dependencyId string) interface{} {
	if c.logger != nil {
		c.logger.Debug("get dependency", "id", dependencyId)
	}

	if c.dependencies[dependencyId] == nil {
		if c.logger != nil {
			c.logger.Fatal("missing dependency", "dependencyId", dependencyId)
		}

		panic(fmt.Sprintf("missing dependency: %s", dependencyId))
	}

	if c.logger != nil {
		c.logger.Debug("dependency found")
	}

	return c.dependencies[dependencyId]
}

func (c *DependencyContainer) Reset() {
	if c.logger != nil {
		c.logger.Debug("resetting dependency container")
	}
	c.dependencies = make(map[string]interface{})
	if c.logger != nil {
		c.logger.Debug("dependency container redefined")
	}
}
