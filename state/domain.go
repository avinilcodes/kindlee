package state

type StateRequest struct {
	PersonalResidential int     `json:"personal_residencial_safety,string"`
	Financial           int     `json:"financial_safety,string"`
	Road                int     `json:"road_safety,string"`
	Workplace           int     `json:"workplace_safety,string"`
	Healthcare          int     `json:"healthcare_safety,string"`
	MonthlyRentSpare    float64 `json:"monthly_rent_spare,string"`
}

type StateResponse struct {
	State string `json:"best_suits"`
}
