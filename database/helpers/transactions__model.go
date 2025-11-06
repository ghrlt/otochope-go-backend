package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	UID              uuid.UUID
	UserUID          uuid.UUID
	CartUID          uuid.UUID
	AmountInEURCents int64
	Amount           string
	Currency         string
	Status           string
	Reference        string
	ChargeReference  string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (t *Transaction) GetUser() (*User, error) {
	return GetUserByUID(t.UserUID)
}

func (t *Transaction) GetCart() (*Cart, error) {
	return GetCartByUID(t.CartUID)
}

func parseTransactionRows(rows *sql.Rows) ([]Transaction, error) {
	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.UID, &transaction.UserUID, &transaction.CartUID, &transaction.AmountInEURCents, &transaction.Amount, &transaction.Currency, &transaction.Status, &transaction.Reference, &transaction.ChargeReference, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func parseTransactionRow(rows *sql.Rows) (*Transaction, error) {
	var transaction Transaction
	if !rows.Next() {
		return nil, nil // No transaction found
	}
	if err := rows.Scan(&transaction.UID, &transaction.UserUID, &transaction.CartUID, &transaction.AmountInEURCents, &transaction.Amount, &transaction.Currency, &transaction.Status, &transaction.Reference, &transaction.ChargeReference, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
		return nil, err
	}
	return &transaction, nil
}
