package system

import (
	"github.com/jaeg/game-engine/entity"
)

// SystemInterface - interface that represents a system, world is an interface and should be cast to whatever data
// structure the game is currently using or that the system cares about.
type SystemInterface interface {
	Update(world interface{}, entity *entity.Entity) error
}
