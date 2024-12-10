package dependencies

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type DependencyContainer struct {
	dependencies map[string]interface{}
}

var dependencyContainerInstance *DependencyContainer

func GetDependencyContainer() *DependencyContainer {
	if dependencyContainerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if dependencyContainerInstance == nil {
			dependencyContainerInstance = &DependencyContainer{
				dependencies: make(map[string]interface{}),
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

func (c *DependencyContainer) AddDependency(dependencyId string, value interface{}) {
	c.dependencies[dependencyId] = value
}

func (c *DependencyContainer) HasDependency(dependencyId string) bool {
	return c.dependencies[dependencyId] != nil
}

func (c *DependencyContainer) GetDependency(dependencyId string) interface{} {
	if c.dependencies[dependencyId] == nil {
		panic(fmt.Sprintf("missing dependency: %s", dependencyId))
	}

	return c.dependencies[dependencyId]
}

func (c *DependencyContainer) Reset() {
	c.dependencies = make(map[string]interface{})
}
