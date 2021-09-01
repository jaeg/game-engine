package system

import (
	"github.com/jaeg/game-engine/entity"
)

// System base system interface
type SystemInterface interface {
	Update(world interface{}, entity *entity.Entity) error
}
