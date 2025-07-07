package handlers

import (
	"strconv"

	"estadisticas-api/internal/estadisticas/application/usecases"
	"estadisticas-api/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type EstadisticasHandler struct {
	getDiaUseCase    *usecases.GetEstadisticasDiaUseCase
	getSemanaUseCase *usecases.GetEstadisticasSemanaUseCase
	getMesUseCase    *usecases.GetEstadisticasMesUseCase
	getAnioUseCase   *usecases.GetEstadisticasAnioUseCase
}

func NewEstadisticasHandler(
	getDiaUseCase *usecases.GetEstadisticasDiaUseCase,
	getSemanaUseCase *usecases.GetEstadisticasSemanaUseCase,
	getMesUseCase *usecases.GetEstadisticasMesUseCase,
	getAnioUseCase *usecases.GetEstadisticasAnioUseCase,
) *EstadisticasHandler {
	return &EstadisticasHandler{
		getDiaUseCase:    getDiaUseCase,
		getSemanaUseCase: getSemanaUseCase,
		getMesUseCase:    getMesUseCase,
		getAnioUseCase:   getAnioUseCase,
	}
}

func (h *EstadisticasHandler) GetEstadisticasDia(c *gin.Context) {
	sensorID := h.getSensorIDFromQuery(c)
	colmenaID := h.getColmenaIDFromQuery(c)
	
	estadisticas, err := h.getDiaUseCase.Execute(c.Request.Context(), sensorID, colmenaID)
	if err != nil {
		response.InternalError(c, "Error obteniendo estadísticas diarias")
		return
	}

	response.Success(c, estadisticas)
}

func (h *EstadisticasHandler) GetEstadisticasSemana(c *gin.Context) {
	sensorID := h.getSensorIDFromQuery(c)
	colmenaID := h.getColmenaIDFromQuery(c)
	
	estadisticas, err := h.getSemanaUseCase.Execute(c.Request.Context(), sensorID, colmenaID)
	if err != nil {
		response.InternalError(c, "Error obteniendo estadísticas semanales")
		return
	}

	response.Success(c, estadisticas)
}

func (h *EstadisticasHandler) GetEstadisticasMes(c *gin.Context) {
	sensorID := h.getSensorIDFromQuery(c)
	colmenaID := h.getColmenaIDFromQuery(c)
	
	estadisticas, err := h.getMesUseCase.Execute(c.Request.Context(), sensorID, colmenaID)
	if err != nil {
		response.InternalError(c, "Error obteniendo estadísticas mensuales")
		return
	}

	response.Success(c, estadisticas)
}

func (h *EstadisticasHandler) GetEstadisticasAnio(c *gin.Context) {
	sensorID := h.getSensorIDFromQuery(c)
	colmenaID := h.getColmenaIDFromQuery(c)
	
	estadisticas, err := h.getAnioUseCase.Execute(c.Request.Context(), sensorID, colmenaID)
	if err != nil {
		response.InternalError(c, "Error obteniendo estadísticas anuales")
		return
	}

	response.Success(c, estadisticas)
}

func (h *EstadisticasHandler) getSensorIDFromQuery(c *gin.Context) *int {
	sensorIDStr := c.Query("sensor_id")
	if sensorIDStr == "" {
		return nil
	}

	sensorID, err := strconv.Atoi(sensorIDStr)
	if err != nil {
		return nil
	}

	return &sensorID
}

func (h *EstadisticasHandler) getColmenaIDFromQuery(c *gin.Context) *int {
	colmenaIDStr := c.Query("colmena_id")
	if colmenaIDStr == "" {
		return nil
	}

	colmenaID, err := strconv.Atoi(colmenaIDStr)
	if err != nil {
		return nil
	}

	return &colmenaID
}