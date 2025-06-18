package application

import "iperf-app/internal/domain"

type ScriptUseCase struct {
	scriptRepo domain.ScriptRepository
	configRepo domain.ConfigRepository
}
