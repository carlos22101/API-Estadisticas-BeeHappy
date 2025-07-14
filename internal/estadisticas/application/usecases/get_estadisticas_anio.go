package usecases

import (
	"context"
	"estadisticas-api/internal/estadisticas/domain"
)

type GetEstadisticasAnioUseCase struct {
	repo domain.EstadisticasRepository
}

func NewGetEstadisticasAnioUseCase(repo domain.EstadisticasRepository) *GetEstadisticasAnioUseCase {
	return &GetEstadisticasAnioUseCase{repo: repo}
}

func (uc *GetEstadisticasAnioUseCase) Execute(ctx context.Context, sensorID *int, macRaspberry *string) ([]domain.EstadisticasAnio, error) {
	return uc.repo.GetEstadisticasAnio(ctx, sensorID, macRaspberry)
}