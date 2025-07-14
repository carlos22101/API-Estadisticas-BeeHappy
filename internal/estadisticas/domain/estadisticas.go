package domain

import "time"

type EstadisticasDia struct {
	ID               int       `json:"id" db:"id"`
	MacRaspberry     string    `json:"mac_raspberry" db:"mac_raspberry"`     
	IDSensor         int       `json:"id_sensor" db:"id_sensor"`
	Fecha            time.Time `json:"fecha" db:"fecha"`
	ValorMinimo      float64   `json:"valor_minimo" db:"valor_minimo"`
	ValorMaximo      float64   `json:"valor_maximo" db:"valor_maximo"`
	ValorPromedio    float64   `json:"valor_promedio" db:"valor_promedio"`
	CantidadLecturas int       `json:"cantidad_lecturas" db:"cantidad_lecturas"`
	FechaCalculo     time.Time `json:"fecha_calculo" db:"fecha_calculo"`
}

type EstadisticasSemana struct {
	ID               int       `json:"id" db:"id"`
	MacRaspberry     string    `json:"mac_raspberry" db:"mac_raspberry"`      
	IDSensor         int       `json:"id_sensor" db:"id_sensor"`
	AnoSemana        string    `json:"ano_semana" db:"ano_semana"`
	FechaInicio      time.Time `json:"fecha_inicio" db:"fecha_inicio"`
	FechaFin         time.Time `json:"fecha_fin" db:"fecha_fin"`
	ValorMinimo      float64   `json:"valor_minimo" db:"valor_minimo"`
	ValorMaximo      float64   `json:"valor_maximo" db:"valor_maximo"`
	ValorPromedio    float64   `json:"valor_promedio" db:"valor_promedio"`
	CantidadLecturas int       `json:"cantidad_lecturas" db:"cantidad_lecturas"`
	FechaCalculo     time.Time `json:"fecha_calculo" db:"fecha_calculo"`
}

type EstadisticasMes struct {
	ID               int       `json:"id" db:"id"`
	MacRaspberry     string    `json:"mac_raspberry" db:"mac_raspberry"`      
	IDSensor         int       `json:"id_sensor" db:"id_sensor"`
	AnoMes           string    `json:"ano_mes" db:"ano_mes"`
	FechaInicio      time.Time `json:"fecha_inicio" db:"fecha_inicio"`
	FechaFin         time.Time `json:"fecha_fin" db:"fecha_fin"`
	ValorMinimo      float64   `json:"valor_minimo" db:"valor_minimo"`
	ValorMaximo      float64   `json:"valor_maximo" db:"valor_maximo"`
	ValorPromedio    float64   `json:"valor_promedio" db:"valor_promedio"`
	CantidadLecturas int       `json:"cantidad_lecturas" db:"cantidad_lecturas"`
	FechaCalculo     time.Time `json:"fecha_calculo" db:"fecha_calculo"`
}

type EstadisticasAnio struct {
	ID               int       `json:"id" db:"id"`
	MacRaspberry     string    `json:"mac_raspberry" db:"mac_raspberry"`      
	IDSensor         int       `json:"id_sensor" db:"id_sensor"`
	Ano              int       `json:"ano" db:"ano"`
	FechaInicio      time.Time `json:"fecha_inicio" db:"fecha_inicio"`
	FechaFin         time.Time `json:"fecha_fin" db:"fecha_fin"`
	ValorMinimo      float64   `json:"valor_minimo" db:"valor_minimo"`
	ValorMaximo      float64   `json:"valor_maximo" db:"valor_maximo"`
	ValorPromedio    float64   `json:"valor_promedio" db:"valor_promedio"`
	CantidadLecturas int       `json:"cantidad_lecturas" db:"cantidad_lecturas"`
	FechaCalculo     time.Time `json:"fecha_calculo" db:"fecha_calculo"`
}