package engines

import (
	"github.com/ethanmdavidson/ecs/core"
)

// defaultEngine is simple a composition of an defaultEntityManager and a defaultSystemManager.
type defaultEngine struct {
	entityManager core.EntityManager
	systemManager core.SystemManager
}

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *defaultEngine) Run() {
	shouldStop := false
	for !shouldStop {
		for _, system := range e.systemManager.Systems() {
			state := system.Process(e.entityManager)
			if state == core.StateEngineStop {
				shouldStop = true
				break
			}
		}
	}
}

// Tick calls the Process() method for each System exactly once
func (e *defaultEngine) Tick() {
	for _, system := range e.systemManager.Systems() {
		system.Process(e.entityManager)
	}
}

// Setup calls the Setup() method for each System
// and initializes ShouldEngineStop and ShouldEnginePause with false.
func (e *defaultEngine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown calls the Teardown() method for each System.
func (e *defaultEngine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}

// NewDefaultEngine creates a new Engine and returns its address.
func NewDefaultEngine(entityManager core.EntityManager, systemManager core.SystemManager) core.Engine {
	return &defaultEngine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}
