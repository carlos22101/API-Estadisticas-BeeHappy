package usecases

import (
	"context"
	"estadisticas-api/internal/estadisticas/domain"
)

type GetEstadisticasDiaUseCase struct {
	repo domain.EstadisticasRepository
}

func NewGetEstadisticasDiaUseCase(repo domain.EstadisticasRepository) *GetEstadisticasDiaUseCase {
	return &GetEstadisticasDiaUseCase{repo: repo}
}

func (uc *GetEstadisticasDiaUseCase) Execute(ctx context.Context, sensorID *int, colmenaID *int) ([]domain.EstadisticasDia, error) {
	return uc.repo.GetEstadisticasDia(ctx, sensorID, colmenaID)
}