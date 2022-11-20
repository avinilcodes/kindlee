package state

type StateRequest struct {
	PersonalResidential string `json:"personal_residencial_safety"`
	Financial           string `json:"financial_safety"`
	Road                string `json:"road_safety"`
	Workplace           string `json:"workplace_safety"`
	Healthcare          string `json:"healthcare_safety"`
	MonthlyRentSpare    string `json:"monthly_rent_spare"`
}
