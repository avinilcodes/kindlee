package db

import "context"

type State struct {
	Id                  string `db:"id"`
	Rank                string `db:"ranking"`
	State               string `db:"state"`
	Score               string `db:"total_score"`
	PersonalResidential string `db:"personal_residencial_safety"`
	Financial           string `db:"financial_safety"`
	Road                string `db:"road_safety"`
	Workplace           string `db:"workplace_safety"`
	Healthcare          string `db:"healthcare_safety"`
	MonthlyRentSpare    string `db:"monthly_rent_spare"`
}

const findAllStatsQuery = `SELECT * FROM stats`

func (s *store) ListStats(ctx context.Context) (stats []State, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &stats, findAllStatsQuery)
	})
	return
}
