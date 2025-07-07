package usecases

import (
	"context"
	"estadisticas-api/internal/estadisticas/domain"
)

type GetEstadisticasSemanaUseCase struct {
	repo domain.EstadisticasRepository
}

func NewGetEstadisticasSemanaUseCase(repo domain.EstadisticasRepository) *GetEstadisticasSemanaUseCase {
	return &GetEstadisticasSemanaUseCase{repo: repo}
}

func (uc *GetEstadisticasSemanaUseCase) Execute(ctx context.Context, sensorID *int, colmenaID *int) ([]domain.EstadisticasSemana, error) {
	return uc.repo.GetEstadisticasSemana(ctx, sensorID, colmenaID)
}