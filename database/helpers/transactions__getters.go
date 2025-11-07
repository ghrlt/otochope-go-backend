package helpers

import (
	"otochope/database"
	"time"

	"github.com/google/uuid"
)

type TransactionStats struct {
	TotalTransactions                  int
	TotalCompletedTransactions         int
	TotalAmountInEURCents              int64
	AverageTransactionAmountInEURCents int64
}

func GetTransactionsStatistics() (*TransactionStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) AS completed,
			SUM(amount_in_eur_cents) AS total_amount,
			AVG(amount_in_eur_cents) AS average_amount
		FROM transactions;
	`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats TransactionStats
	if rows.Next() {
		if err := rows.Scan(&stats.TotalTransactions, &stats.TotalCompletedTransactions, &stats.TotalAmountInEURCents, &stats.AverageTransactionAmountInEURCents); err != nil {
			return nil, err
		}
	}

	return &stats, nil
}

func GetAllTransactions() ([]Transaction, error) {
	query := "SELECT * FROM transactions"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsByUserUID(userUID uuid.UUID) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE user_uid = ?"
	rows, err := database.Query(query, userUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsByCartUID(cartUID uuid.UUID) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE cart_uid = ?"
	rows, err := database.Query(query, cartUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsByStatus(status string) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE status = ?"
	rows, err := database.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsByCurrency(currency string) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE currency = ?"
	rows, err := database.Query(query, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

// minAmount and maxAmount are in EUR cents
func GetAllTransactionsInAmountRange(minAmount int64, maxAmount int64) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE amount_in_eur_cents BETWEEN ? AND ?"
	rows, err := database.Query(query, minAmount, maxAmount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsCreatedAfter(date time.Time) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE created_at >= ?"
	rows, err := database.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsCreatedBefore(date time.Time) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE created_at <= ?"
	rows, err := database.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetAllTransactionsCreatedBetween(startDate time.Time, endDate time.Time) ([]Transaction, error) {
	query := "SELECT * FROM transactions WHERE created_at BETWEEN ? AND ?"
	rows, err := database.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRows(rows)
}

func GetTransactionByUID(transactionUID uuid.UUID) (*Transaction, error) {
	query := "SELECT * FROM transactions WHERE uid = ?"
	rows, err := database.Query(query, transactionUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRow(rows)
}

func GetTransactionByReference(reference string) (*Transaction, error) {
	query := "SELECT * FROM transactions WHERE reference = ?"
	rows, err := database.Query(query, reference)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRow(rows)
}

func GetTransactionByChargeReference(chargeReference string) (*Transaction, error) {
	query := "SELECT * FROM transactions WHERE charge_reference = ?"
	rows, err := database.Query(query, chargeReference)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseTransactionRow(rows)
}
