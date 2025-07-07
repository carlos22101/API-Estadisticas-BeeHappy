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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

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

	log.Fatal(router.Run(":" + port))
}