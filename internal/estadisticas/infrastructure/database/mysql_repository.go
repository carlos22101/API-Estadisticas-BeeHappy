package estadisticasdb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"estadisticas-api/internal/estadisticas/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) GetEstadisticasDia(ctx context.Context, sensorID *int, macRaspberry *string) ([]domain.EstadisticasDia, error) {
	query := `SELECT id, mac_raspberry, id_sensor, fecha, valor_minimo, valor_maximo, valor_promedio, cantidad_lecturas, fecha_calculo 
			  FROM estadisticas_dia`
	args := []interface{}{}
	conditions := []string{}

	if sensorID != nil {
		conditions = append(conditions, "id_sensor = ?")
		args = append(args, *sensorID)
	}

	if macRaspberry != nil {
		conditions = append(conditions, "mac_raspberry = ?")
		args = append(args, *macRaspberry)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY fecha DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error consultando estadísticas diarias: %w", err)
	}
	defer rows.Close()

	var estadisticas []domain.EstadisticasDia
	for rows.Next() {
		var est domain.EstadisticasDia
		err := rows.Scan(&est.ID, &est.MacRaspberry, &est.IDSensor, &est.Fecha, &est.ValorMinimo, 
						&est.ValorMaximo, &est.ValorPromedio, &est.CantidadLecturas, &est.FechaCalculo)
		if err != nil {
			return nil, fmt.Errorf("error escaneando estadísticas diarias: %w", err)
		}
		estadisticas = append(estadisticas, est)
	}

	return estadisticas, nil
}

func (r *MySQLRepository) GetEstadisticasSemana(ctx context.Context, sensorID *int, macRaspberry *string) ([]domain.EstadisticasSemana, error) {
	query := `SELECT id, mac_raspberry, id_sensor, ano_semana, fecha_inicio, fecha_fin, valor_minimo, 
			  valor_maximo, valor_promedio, cantidad_lecturas, fecha_calculo 
			  FROM estadisticas_semana`
	args := []interface{}{}
	conditions := []string{}

	if sensorID != nil {
		conditions = append(conditions, "id_sensor = ?")
		args = append(args, *sensorID)
	}

	if macRaspberry != nil {
		conditions = append(conditions, "mac_raspberry = ?")
		args = append(args, *macRaspberry)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY fecha_inicio DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error consultando estadísticas semanales: %w", err)
	}
	defer rows.Close()

	var estadisticas []domain.EstadisticasSemana
	for rows.Next() {
		var est domain.EstadisticasSemana
		err := rows.Scan(&est.ID, &est.MacRaspberry, &est.IDSensor, &est.AnoSemana, &est.FechaInicio, &est.FechaFin,
						&est.ValorMinimo, &est.ValorMaximo, &est.ValorPromedio, &est.CantidadLecturas, &est.FechaCalculo)
		if err != nil {
			return nil, fmt.Errorf("error escaneando estadísticas semanales: %w", err)
		}
		estadisticas = append(estadisticas, est)
	}

	return estadisticas, nil
}

func (r *MySQLRepository) GetEstadisticasMes(ctx context.Context, sensorID *int, macRaspberry *string) ([]domain.EstadisticasMes, error) {
	query := `SELECT id, mac_raspberry, id_sensor, ano_mes, fecha_inicio, fecha_fin, valor_minimo, 
			  valor_maximo, valor_promedio, cantidad_lecturas, fecha_calculo 
			  FROM estadisticas_mes`
	args := []interface{}{}
	conditions := []string{}

	if sensorID != nil {
		conditions = append(conditions, "id_sensor = ?")
		args = append(args, *sensorID)
	}

	if macRaspberry != nil {
		conditions = append(conditions, "mac_raspberry = ?")
		args = append(args, *macRaspberry)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY fecha_inicio DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error consultando estadísticas mensuales: %w", err)
	}
	defer rows.Close()

	var estadisticas []domain.EstadisticasMes
	for rows.Next() {
		var est domain.EstadisticasMes
		err := rows.Scan(&est.ID, &est.MacRaspberry, &est.IDSensor, &est.AnoMes, &est.FechaInicio, &est.FechaFin,
						&est.ValorMinimo, &est.ValorMaximo, &est.ValorPromedio, &est.CantidadLecturas, &est.FechaCalculo)
		if err != nil {
			return nil, fmt.Errorf("error escaneando estadísticas mensuales: %w", err)
		}
		estadisticas = append(estadisticas, est)
	}

	return estadisticas, nil
}

func (r *MySQLRepository) GetEstadisticasAnio(ctx context.Context, sensorID *int, macRaspberry *string) ([]domain.EstadisticasAnio, error) {
	query := `SELECT id, mac_raspberry, id_sensor, ano, fecha_inicio, fecha_fin, valor_minimo, 
			  valor_maximo, valor_promedio, cantidad_lecturas, fecha_calculo 
			  FROM estadisticas_ano`
	args := []interface{}{}
	conditions := []string{}

	if sensorID != nil {
		conditions = append(conditions, "id_sensor = ?")
		args = append(args, *sensorID)
	}

	if macRaspberry != nil {
		conditions = append(conditions, "mac_raspberry = ?")
		args = append(args, *macRaspberry)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY ano DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error consultando estadísticas anuales: %w", err)
	}
	defer rows.Close()

	var estadisticas []domain.EstadisticasAnio
	for rows.Next() {
		var est domain.EstadisticasAnio
		err := rows.Scan(&est.ID, &est.MacRaspberry, &est.IDSensor, &est.Ano, &est.FechaInicio, &est.FechaFin,
						&est.ValorMinimo, &est.ValorMaximo, &est.ValorPromedio, &est.CantidadLecturas, &est.FechaCalculo)
		if err != nil {
			return nil, fmt.Errorf("error escaneando estadísticas anuales: %w", err)
		}
		estadisticas = append(estadisticas, est)
	}

	return estadisticas, nil
}