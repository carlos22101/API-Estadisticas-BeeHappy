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

// GetEstadisticasDia obtiene estadísticas diarias
// @Summary Obtener estadísticas diarias
// @Description Retorna las estadísticas calculadas por día para las colmenas y sensores especificados
// @Tags Estadísticas Diarias
// @Accept json
// @Produce json
// @Param sensor_id query int false "ID del sensor para filtrar las estadísticas" minimum(1) example(1)
// @Param colmena_id query int false "ID de la colmena para filtrar las estadísticas" minimum(1) example(1)
// @Success 200 {object} response.SuccessResponse{data=[]domain.EstadisticasDia} "Estadísticas diarias obtenidas exitosamente"
// @Failure 400 {object} response.ErrorResponse "Parámetros de consulta inválidos"
// @Failure 500 {object} response.ErrorResponse "Error interno del servidor"
// @Router /estadisticas/dia [get]
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

// GetEstadisticasSemana obtiene estadísticas semanales
// @Summary Obtener estadísticas semanales
// @Description Retorna las estadísticas calculadas por semana para las colmenas y sensores especificados
// @Tags Estadísticas Semanales
// @Accept json
// @Produce json
// @Param sensor_id query int false "ID del sensor para filtrar las estadísticas" minimum(1) example(1)
// @Param colmena_id query int false "ID de la colmena para filtrar las estadísticas" minimum(1) example(1)
// @Success 200 {object} response.SuccessResponse{data=[]domain.EstadisticasSemana} "Estadísticas semanales obtenidas exitosamente"
// @Failure 400 {object} response.ErrorResponse "Parámetros de consulta inválidos"
// @Failure 500 {object} response.ErrorResponse "Error interno del servidor"
// @Router /estadisticas/semana [get]
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

// GetEstadisticasMes obtiene estadísticas mensuales
// @Summary Obtener estadísticas mensuales
// @Description Retorna las estadísticas calculadas por mes para las colmenas y sensores especificados
// @Tags Estadísticas Mensuales
// @Accept json
// @Produce json
// @Param sensor_id query int false "ID del sensor para filtrar las estadísticas" minimum(1) example(1)
// @Param colmena_id query int false "ID de la colmena para filtrar las estadísticas" minimum(1) example(1)
// @Success 200 {object} response.SuccessResponse{data=[]domain.EstadisticasMes} "Estadísticas mensuales obtenidas exitosamente"
// @Failure 400 {object} response.ErrorResponse "Parámetros de consulta inválidos"
// @Failure 500 {object} response.ErrorResponse "Error interno del servidor"
// @Router /estadisticas/mes [get]
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

// GetEstadisticasAnio obtiene estadísticas anuales
// @Summary Obtener estadísticas anuales
// @Description Retorna las estadísticas calculadas por año para las colmenas y sensores especificados
// @Tags Estadísticas Anuales
// @Accept json
// @Produce json
// @Param sensor_id query int false "ID del sensor para filtrar las estadísticas" minimum(1) example(1)
// @Param colmena_id query int false "ID de la colmena para filtrar las estadísticas" minimum(1) example(1)
// @Success 200 {object} response.SuccessResponse{data=[]domain.EstadisticasAnio} "Estadísticas anuales obtenidas exitosamente"
// @Failure 400 {object} response.ErrorResponse "Parámetros de consulta inválidos"
// @Failure 500 {object} response.ErrorResponse "Error interno del servidor"
// @Router /estadisticas/anio [get]
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