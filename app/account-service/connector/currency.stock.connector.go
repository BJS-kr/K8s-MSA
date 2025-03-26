package connector

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
)

var (
	ErrRequestLoanDenied error = errors.New("request loan denied")
	ErrReturnLoanDenied  error = errors.New("return loan denied")
)

type LoanServiceConnector interface {
	RequestLoan(ctx context.Context, currency string, amount int) error
	ReturnLoan(ctx context.Context, currency string, amount int) error
}

type LoanService struct {
	url *url.URL
}

func NewLoanServiceConnector() *LoanService {
	loanServiceAddr := os.Getenv("LOAN_SERVICE_ADDR")

	if loanServiceAddr == "" {
		log.Fatal("LOAN_SERVICE_ADDR is not set")
	}

	loanService := new(LoanService)
	loanServiceUrl, err := url.Parse(loanServiceAddr)

	if err != nil {
		log.Fatal("LOAN_SERVICE_ADDR is invalid address")
	}

	loanService.url = loanServiceUrl

	return loanService
}

func (l *LoanService) RequestLoan(ctx context.Context, currency string, amount int) error {
	return nil
}

func (*LoanService) ReturnLoan(ctx context.Context, currency string, amount int) error {
	return nil
}
