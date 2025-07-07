package usecases

import (
	"context"
	"estadisticas-api/internal/estadisticas/domain"
)

type GetEstadisticasMesUseCase struct {
	repo domain.EstadisticasRepository
}

func NewGetEstadisticasMesUseCase(repo domain.EstadisticasRepository) *GetEstadisticasMesUseCase {
	return &GetEstadisticasMesUseCase{repo: repo}
}

func (uc *GetEstadisticasMesUseCase) Execute(ctx context.Context, sensorID *int, colmenaID *int) ([]domain.EstadisticasMes, error) {
	return uc.repo.GetEstadisticasMes(ctx, sensorID, colmenaID)
}