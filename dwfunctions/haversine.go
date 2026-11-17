package dwfunctions

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateHaversineFunction(dbpool *pgxpool.Pool) {
	query := `
	CREATE OR REPLACE FUNCTION dw.haversine_lat_long_metros_ponto_servico(
		p_lat NUMERIC,
		p_lon NUMERIC
	)
	RETURNS TABLE (
		ponto_servico_key BIGINT,
		distancia_metros NUMERIC
	)
	LANGUAGE sql
	STABLE
	AS $$
		SELECT 
			p.id_ponto_servico,
			(
				6371000 * 2 * ATAN2(
					SQRT(
						SIN(RADIANS(p.latitude - p_lat) / 2)^2 +
						COS(RADIANS(p_lat)) * COS(RADIANS(p.latitude)) *
						SIN(RADIANS(p.longitude - p_lon) / 2)^2
					),
					SQRT(1 - (
						SIN(RADIANS(p.latitude - p_lat) / 2)^2 +
						COS(RADIANS(p_lat)) * COS(RADIANS(p.latitude)) *
						SIN(RADIANS(p.longitude - p_lon) / 2)^2
					))
				)
			) AS distancia_metros
		FROM dw.dim_ponto_servico p
		WHERE p.latitude IS NOT NULL
		AND p.latitude  BETWEEN p_lat - 0.001 AND p_lat + 0.001
		AND p.longitude BETWEEN p_lon - 0.001 AND p_lon + 0.001
		ORDER BY distancia_metros ASC
		LIMIT 1;
	$$;
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
