package models

type Plan struct {
        PlanId string `json:"plan_id"`
		PlanName string  `json:"plan_name"`
		Lob string `json:"lob"`
}

type Level struct {
        LevelId string `json:"level_id"`
		LevelName string  `json:"level_name"`
		LevelType string `json:"level_type"`
		Lob string `json:"lob"`
		LevelValues []LevelValue `json:"level_values"`
}

type LevelValue struct {
        Value string `json:"value"`
		ValueBasis string  `json:"value_basis"`
}

