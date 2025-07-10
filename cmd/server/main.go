package main

import (
	"log"
	"os"

	"estadisticas-api/configs"
	"estadisticas-api/internal/estadisticas/application/usecases"
	estadisticasdb "estadisticas-api/internal/estadisticas/infrastructure/database"
	"estadisticas-api/internal/estadisticas/infrastructure/http/handlers"
	"estadisticas-api/internal/estadisticas/infrastructure/http/middleware"
	"estadisticas-api/internal/estadisticas/infrastructure/http/routes"
	sharedDB "estadisticas-api/internal/shared/database"

	// ✨ SWAGGER IMPORTS
	_ "estadisticas-api/docs" // Importar documentación generada
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title BeeHappy Estadísticas API
// @version 1.0
// @description API para consultar estadísticas de colmenas inteligentes BeeHappy
// @description Esta API proporciona acceso a datos estadísticos agregados de sensores de colmenas
// @termsOfService https://beehappy.com/terms

// @contact.name Equipo BeeHappy
// @contact.url https://beehappy.com
// @contact.email api@beehappy.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @tag.name Estadísticas Diarias
// @tag.description Operaciones relacionadas con estadísticas calculadas por día

// @tag.name Estadísticas Semanales
// @tag.description Operaciones relacionadas con estadísticas calculadas por semana

// @tag.name Estadísticas Mensuales
// @tag.description Operaciones relacionadas con estadísticas calculadas por mes

// @tag.name Estadísticas Anuales
// @tag.description Operaciones relacionadas con estadísticas calculadas por año

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	// Configurar base de datos
	db, err := sharedDB.NewConnection(configs.GetDatabaseConfig())
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Verificar conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error verificando conexión a la base de datos:", err)
	}
	log.Println("✅ Conexión a base de datos establecida correctamente")

	// Inicializar repositorio
	estadisticasRepo := estadisticasdb.NewMySQLRepository(db)

	// Inicializar casos de uso
	getDiaUseCase := usecases.NewGetEstadisticasDiaUseCase(estadisticasRepo)
	getSemanaUseCase := usecases.NewGetEstadisticasSemanaUseCase(estadisticasRepo)
	getMesUseCase := usecases.NewGetEstadisticasMesUseCase(estadisticasRepo)
	getAnioUseCase := usecases.NewGetEstadisticasAnioUseCase(estadisticasRepo)

	// Inicializar handlers
	estadisticasHandler := handlers.NewEstadisticasHandler(
		getDiaUseCase,
		getSemanaUseCase,
		getMesUseCase,
		getAnioUseCase,
	)

	// Configurar servidor
	router := gin.Default()

	// Aplicar middleware CORS
	router.Use(middleware.CORSMiddleware())

	// ✨ CONFIGURAR SWAGGER
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar rutas
	routes.SetupEstadisticasRoutes(router, estadisticasHandler)

	// Información de configuración
	dbConfig := configs.GetDatabaseConfig()
	log.Printf("🐝 API BeeHappy - Estadísticas")
	log.Printf("📊 Base de datos: %s", dbConfig.Database)
	log.Printf("🔗 Host: %s:%s", dbConfig.Host, dbConfig.Port)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("📋 Documentación Swagger: http://localhost:%s/swagger/index.html", port)
	
	log.Fatal(router.Run(":" + port))
}