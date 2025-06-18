package domain

import (
	"iperf-app/internal/domain/entities"
)

type ScriptRepository interface {
	GetAll() ([]entities.Script, error)
	GetByName(name string) (*entities.Script, error)
	Save(script *entities.Script) error
	Delete(name string) error
}

type ConfigRepository interface {
	Load() (*entities.Config, error)
	Save(config *entities.Config) error
}

type ExecutionRepository interface {
	Save(execution *entities.Execution) error
	GetByID(id string) (*entities.Execution, error)
	GetHistory(limit int) ([]entities.Execution, error)
}
