package domain

import "context"

type EstadisticasRepository interface {
	GetEstadisticasDia(ctx context.Context, sensorID *int, colmenaID *int) ([]EstadisticasDia, error)
	GetEstadisticasSemana(ctx context.Context, sensorID *int, colmenaID *int) ([]EstadisticasSemana, error)
	GetEstadisticasMes(ctx context.Context, sensorID *int, colmenaID *int) ([]EstadisticasMes, error)
	GetEstadisticasAnio(ctx context.Context, sensorID *int, colmenaID *int) ([]EstadisticasAnio, error)
}