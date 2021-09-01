package system

import (
	"log"

	"github.com/jaeg/game-engine/entity"
)

// HelloWorldSystem .
type HelloWorldSystem struct {
}

// Update .
func (c HelloWorldSystem) Update(world interface{}, entity *entity.Entity) {
	log.Println("hello world")
}
