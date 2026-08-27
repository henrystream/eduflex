package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/henrystream/eduflex/fraud-service/internal/client"
	"github.com/jackc/pgx/v5/pgtype"
)

type FraudService struct {
	students  *client.StudentClient
	schools   *client.SchoolClient
	financing *client.FinancingClient
	loans     *client.LoanClient
	ledger    *client.LedgerClient
	events    *client.EventsClient
	scoring   *ScoringEngine
}

func NewFraudService(
	s *client.StudentClient,
	sc *client.SchoolClient,
	f *client.FinancingClient,
	l *client.LoanClient,
	lg *client.LedgerClient,
	ev *client.EventsClient,
	se *ScoringEngine,
) *FraudService {
	return &FraudService{
		students:  s,
		schools:   sc,
		financing: f,
		loans:     l,
		ledger:    lg,
		events:    ev,
		scoring:   se,
	}
}

type FraudCheckResult struct {
	StudentRisk   RiskScore `json:"student_risk"`
	SchoolRisk    RiskScore `json:"school_risk"`
	AgreementRisk RiskScore `json:"agreement_risk"`
}

func (s *FraudService) CheckAgreementRisk(studentID, agreementID pgtype.UUID) (FraudCheckResult, error) {

	agreements, _ := s.financing.GetAgreement(studentID)
	/*if err != nil {
		return FraudCheckResult{}, err
	}*/
	var newAgree client.Agreement
	for _, agree := range agreements {
		if agree.ID == agreementID {
			newAgree = agree
			break
		}

	}
	student, _ := s.students.GetStudent(newAgree.StudentID.String())
	/*if err != nil {
		return FraudCheckResult{}, err
	}*/

	school, _ := s.schools.GetSchool(newAgree.SchoolID)

	/*if err != nil {
		return FraudCheckResult{}, err
	}*/

	// naive: count agreements by student via ledger entries
	entries, _ := s.ledger.ListByEvent("FINANCING_AGREEMENT", newAgree.StudentID.String())
	numAgreements := len(entries)

	numPrincipal := fmt.Sprintf("%v", newAgree.Principal)

	studentRisk := s.scoring.ScoreStudent(student.Country, numAgreements)
	schoolRisk := s.scoring.ScoreSchool(school.Tier)
	agreementRisk := s.scoring.ScoreAgreement(parseAmount(numPrincipal), newAgree.TermMonths)

	result := FraudCheckResult{
		StudentRisk:   studentRisk,
		SchoolRisk:    schoolRisk,
		AgreementRisk: agreementRisk,
	}

	s.events.Publish(
		"FRAUD_RISK_EVALUATED",
		agreementID,
		pgtype.Timestamptz{Time: time.Now(), Valid: true},
		result,
	)

	return result, nil
}

func parseAmount(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
