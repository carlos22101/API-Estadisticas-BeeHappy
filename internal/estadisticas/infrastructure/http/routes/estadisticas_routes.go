package routes

import (
	"estadisticas-api/internal/estadisticas/infrastructure/http/handlers"
	"github.com/gin-gonic/gin"
)

func SetupEstadisticasRoutes(router *gin.Engine, handler *handlers.EstadisticasHandler) {
	api := router.Group("/api/v1")
	{
		estadisticas := api.Group("/estadisticas")
		{
			estadisticas.GET("/dia", handler.GetEstadisticasDia)
			estadisticas.GET("/semana", handler.GetEstadisticasSemana)
			estadisticas.GET("/mes", handler.GetEstadisticasMes)
			estadisticas.GET("/anio", handler.GetEstadisticasAnio)
		}
	}
}