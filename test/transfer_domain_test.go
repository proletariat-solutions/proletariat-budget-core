package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
)

func TestTransfer_SetTransferType_SetsTransferTypeSavingOperationWhenSavingOperationIsTrue(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(true)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeSavingOperation,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeTransferWhenAccountsHaveSameTypeAndSameCurrency(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeTransfer,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeCurrencyConversionWhenAccountsHaveDifferentCurrencies(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "EUR",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeCurrencyConversion,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeWithdrawalWhenDestinationAccountIsCashTypeAndCurrenciesAreTheSame(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeWithdrawal,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeDepositWhenSourceAccountIsCashTypeAndCurrenciesAreTheSame(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeDeposit,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeTransferWhenBothAccountsAreCashTypeAndCurrenciesAreTheSame(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeTransfer,
		transfer.TransferType,
	)
}
func TestTransfer_SetTransferType_SetsTransferTypeCurrencyConversionWhenSourceIsCashDestinationIsBankAndCurrenciesAreDifferent(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "EUR",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeCurrencyConversion,
		transfer.TransferType,
	)
}
func TestTransfer_SetTransferType_SetsTransferTypeCurrencyConversionWhenSourceIsBankDestinationIsCashAndCurrenciesAreDifferent(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "EUR",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeCurrencyConversion,
		transfer.TransferType,
	)
}
func TestTransfer_SetTransferType_SetsTransferTypeCurrencyConversionWhenBothAccountsAreCashTypeButCurrenciesAreDifferent(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "EUR",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeCurrencyConversion,
		transfer.TransferType,
	)
}

func TestTransfer_SetTransferType_SetsTransferTypeCurrencyConversionWhenAccountsHaveDifferentTypesAndDifferentCurrencies(t *testing.T) {
	// Arrange
	sourceAccount := &coreentity.Account{
		Type: coreentity.AccountTypeBank,
		Currency: &coreentity.Currency{
			ID: "USD",
		},
	}
	destinationAccount := &coreentity.Account{
		Type: coreentity.AccountTypeCash,
		Currency: &coreentity.Currency{
			ID: "EUR",
		},
	}
	transfer := &coreentity.Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
	}

	// Act
	transfer.SetTransferType(false)

	// Assert
	assert.Equal(
		t,
		coreentity.TransferTypeCurrencyConversion,
		transfer.TransferType,
	)
}
