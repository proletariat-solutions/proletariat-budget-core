package coreentity

import (
	"errors"
	"time"
)

type Account struct {
	ID                 *string          `json:"id"`
	Name               string           `json:"name"`
	Type               AccountType      `json:"type"`
	Currency           *Currency        `json:"currency"`
	InitialBalance     float32          `json:"initial_balance"`
	CurrentBalance     float32          `json:"current_balance"`
	OverdraftLimit     float32          `json:"overdraft_limit"`
	Description        *string          `json:"description"`
	Institution        *string          `json:"institution"`
	AccountNumber      *string          `json:"account_number"`
	AccountInformation *string          `json:"account_information"`
	OwnerID            *string          `json:"owner_id"`
	Owner              *HouseholdMember `json:"owner,omitempty"`
	Active             bool             `json:"active"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

// Account domain errors
var (
	ErrAccountNotFound                    = errors.New("account not found")
	ErrAccountInactive                    = errors.New("account is inactive")
	ErrInsufficientBalance                = errors.New("insufficient account balance")
	ErrAccountHasTransactions             = errors.New("cannot delete account with existing transactions")
	ErrAccountAlreadyActive               = errors.New("account is already active")
	ErrAccountAlreadyInactive             = errors.New("account is already inactive")
	ErrAccountHasActiveRecurrencePatterns = errors.New("account has active recurrence patterns")
	ErrAccountHasActiveSavingsGoals       = errors.New("account has active savings goals")
	ErrAccountCurrencyRequired            = errors.New("account currency is required")
	ErrAccountOwnerRequired               = errors.New("account owner is required")
	ErrAccountNameRequired                = errors.New("account name is required")
)

type AccountType string

const (
	AccountTypeBank       AccountType = "bank"
	AccountTypeCash       AccountType = "cash"
	AccountTypeCrypto     AccountType = "crypto"
	AccountTypeInvestment AccountType = "investment"
	AccountTypeOther      AccountType = "other"
)

func (a *Account) Validate() error {
	if a.Name == "" {
		return ErrAccountNameRequired
	}
	if a.Currency == nil || a.Currency.ID == "" {
		return ErrAccountCurrencyRequired
	}
	if a.Owner == nil || a.Owner.ID == "" {
		return ErrAccountOwnerRequired
	}

	return nil
}

// DebitBalance debits the account balance (for expenditures and outgoing transfers)
func (a *Account) DebitBalance(amount float32) error {
	if !a.HasSufficientBalance(amount) {
		return ErrInsufficientBalance
	}
	a.CurrentBalance -= amount
	a.UpdatedAt = time.Now()

	return nil
}

// CreditBalance credits the account balance (for income and incoming transfers)
func (a *Account) CreditBalance(amount float32) {
	a.CurrentBalance += amount
	a.UpdatedAt = time.Now()
}

// SetActive sets the account active status
func (a *Account) SetActive() error {
	if a.Active {
		return ErrAccountAlreadyActive
	}
	a.Active = true
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Account) SetInactive() error {
	if !a.Active {
		return ErrAccountAlreadyInactive
	}
	a.Active = false
	a.UpdatedAt = time.Now()

	return nil
}

// HasSufficientBalance checks if the account has sufficient balance for a transaction
func (a *Account) HasSufficientBalance(amount float32) bool {
	return a.CurrentBalance+a.OverdraftLimit > amount
}
