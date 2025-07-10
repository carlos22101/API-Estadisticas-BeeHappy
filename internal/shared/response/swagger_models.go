package response

import (
	"estadisticas-api/internal/estadisticas/domain"
)

// SuccessResponse modelo base para respuestas exitosas
type SuccessResponse struct {
	Status    string      `json:"status" example:"success"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string      `json:"version,omitempty" example:"v1.0"`
}

// ErrorResponse modelo para respuestas de error
type ErrorResponse struct {
	Status    string `json:"status" example:"error"`
	Message   string `json:"message" example:"Error obteniendo estadísticas"`
	Timestamp string `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string `json:"version,omitempty" example:"v1.0"`
}

// EstadisticasDiaResponse respuesta específica para estadísticas diarias
type EstadisticasDiaResponse struct {
	Status    string                    `json:"status" example:"success"`
	Data      []domain.EstadisticasDia  `json:"data"`
	Timestamp string                    `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                    `json:"version,omitempty" example:"v1.0"`
}

// EstadisticasSemanaResponse respuesta específica para estadísticas semanales
type EstadisticasSemanaResponse struct {
	Status    string                      `json:"status" example:"success"`
	Data      []domain.EstadisticasSemana `json:"data"`
	Timestamp string                      `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                      `json:"version,omitempty" example:"v1.0"`
}

// EstadisticasMesResponse respuesta específica para estadísticas mensuales
type EstadisticasMesResponse struct {
	Status    string                   `json:"status" example:"success"`
	Data      []domain.EstadisticasMes `json:"data"`
	Timestamp string                   `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                   `json:"version,omitempty" example:"v1.0"`
}

// EstadisticasAnioResponse respuesta específica para estadísticas anuales
type EstadisticasAnioResponse struct {
	Status    string                    `json:"status" example:"success"`
	Data      []domain.EstadisticasAnio `json:"data"`
	Timestamp string                    `json:"timestamp,omitempty" example:"2025-07-07T12:30:00Z"`
	Version   string                    `json:"version,omitempty" example:"v1.0"`
}