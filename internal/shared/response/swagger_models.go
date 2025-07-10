package response

import (
	"estadisticas-api/internal/estadisticas/domain"
)


type SuccessResponse struct {
	Status    string      `json:"status" example:"success"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string      `json:"version,omitempty" example:"v1.0"`
}


type ErrorResponse struct {
	Status    string `json:"status" example:"error"`
	Message   string `json:"message" example:"Error obteniendo estadísticas"`
	Timestamp string `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string `json:"version,omitempty" example:"v1.0"`
}


type EstadisticasDiaResponse struct {
	Status    string                    `json:"status" example:"success"`
	Data      []domain.EstadisticasDia  `json:"data"`
	Timestamp string                    `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                    `json:"version,omitempty" example:"v1.0"`
}

type EstadisticasSemanaResponse struct {
	Status    string                      `json:"status" example:"success"`
	Data      []domain.EstadisticasSemana `json:"data"`
	Timestamp string                      `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                      `json:"version,omitempty" example:"v1.0"`
}


type EstadisticasMesResponse struct {
	Status    string                   `json:"status" example:"success"`
	Data      []domain.EstadisticasMes `json:"data"`
	Timestamp string                   `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                   `json:"version,omitempty" example:"v1.0"`
}


type EstadisticasAnioResponse struct {
	Status    string                    `json:"status" example:"success"`
	Data      []domain.EstadisticasAnio `json:"data"`
	Timestamp string                    `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                    `json:"version,omitempty" example:"v1.0"`
}