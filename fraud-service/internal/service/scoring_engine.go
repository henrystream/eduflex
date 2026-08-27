package service

import "math"

type RiskScore struct {
	Score       float64 `json:"score"`
	Level       string  `json:"level"`
	Explanation string  `json:"explanation"`
}

type ScoringEngine struct{}

func NewScoringEngine() *ScoringEngine {
	return &ScoringEngine{}
}

func (e *ScoringEngine) ScoreStudent(country string, numAgreements int) RiskScore {
	score := 0.0

	if country != "AE" {
		score += 30
	}
	if numAgreements > 3 {
		score += 40
	}

	level := "LOW"
	if score >= 70 {
		level = "HIGH"
	} else if score >= 40 {
		level = "MEDIUM"
	}

	return RiskScore{
		Score:       math.Min(score, 100),
		Level:       level,
		Explanation: "Student risk based on country and number of agreements",
	}
}

func (e *ScoringEngine) ScoreSchool(tier string) RiskScore {
	score := 0.0
	switch tier {
	case "C":
		score = 60
	case "B":
		score = 30
	}

	level := "LOW"
	if score >= 70 {
		level = "HIGH"
	} else if score >= 40 {
		level = "MEDIUM"
	}

	return RiskScore{
		Score:       score,
		Level:       level,
		Explanation: "School risk based on tier",
	}
}

func (e *ScoringEngine) ScoreAgreement(principal float64, termMonths int) RiskScore {
	score := 0.0

	if principal > 100000 {
		score += 50
	}
	if termMonths > 36 {
		score += 30
	}

	level := "LOW"
	if score >= 70 {
		level = "HIGH"
	} else if score >= 40 {
		level = "MEDIUM"
	}

	return RiskScore{
		Score:       math.Min(score, 100),
		Level:       level,
		Explanation: "Agreement risk based on principal and term",
	}
}
