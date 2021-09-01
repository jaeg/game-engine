package system

import "github.com/jaeg/game-engine/entity"

type SystemManager struct {
	systems []SystemInterface
}

func (s *SystemManager) AddSystem(system SystemInterface) {
	if s.systems == nil {
		s.systems = make([]SystemInterface, 0)
	}

	s.systems = append(s.systems, system)
}

func (s *SystemManager) UpdateSystemsForEntity(world interface{}, entity *entity.Entity) error {
	for system := range s.systems {
		err := s.systems[system].Update(world, entity)
		if err != nil {
			return err
		}
	}
	return nil
}
