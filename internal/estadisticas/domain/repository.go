package domain

import "context"

type EstadisticasRepository interface {
	GetEstadisticasDia(ctx context.Context, sensorID *int, macRaspberry *string) ([]EstadisticasDia, error)
	GetEstadisticasSemana(ctx context.Context, sensorID *int, macRaspberry *string) ([]EstadisticasSemana, error)
	GetEstadisticasMes(ctx context.Context, sensorID *int, macRaspberry *string) ([]EstadisticasMes, error)
	GetEstadisticasAnio(ctx context.Context, sensorID *int, macRaspberry *string) ([]EstadisticasAnio, error)
}