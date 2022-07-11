package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateNewTransection(t *testing.T) (Transaction, error) {
	arg := CreateTransactionParams{
		UserID: 1,
		CategoryID: 25,
		Notes: "salary",
		Status: TransactionStatueExpense,
		Amount: -85108,
	}

	transaction, err := testQueries.CreateTransaction(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.NotEmpty(t, transaction.Notes)
	if transaction.Status == TransactionStatueExpense {
		require.Negative(t, transaction.Amount)
	} else {
		require.Positive(t, transaction.Amount)
	}

	return transaction, err
}

func TestCreateTransaction(t *testing.T) {
	transaction, _ := CreateNewTransection(t)
	fmt.Printf(transaction.Notes, transaction.Amount, transaction.Status)
}

func TestGetTransaction(t *testing.T) {
	transaction, err := CreateNewTransection(t)

	require.NoError(t, err)
	transaction, err = testQueries.GetTransaction(ctx, transaction.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)
	require.NotZero(t, transaction.UpdatedAt)
	require.NotZero(t, transaction.UserID)
}

func TestUpdateTransaction(t *testing.T) {
	tranaction, err := CreateNewTransection(t)
	require.NoError(t, err)
	require.NotEmpty(t, tranaction)
	arg := UpdateTransactionParams{
		ID: tranaction.ID,
		Amount: -234,
		Notes: tranaction.Notes,
		CategoryID: tranaction.CategoryID,
		UpdatedAt: time.Now(),
	}

	err = testQueries.UpdateTransaction(ctx, arg)
	require.NoError(t, err)
}

func TestDeleteTransaction(t *testing.T) {
	transaction, err := CreateNewTransection(t)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	err = testQueries.DeleteTransaction(ctx, transaction.ID)
	require.NoError(t, err)
}

func TestListTransactionByStatus(t *testing.T) {
	for i := 0; i < 6; i++ {
		CreateNewTransection(t)
	}

	arg := ListTransactionByStatusParams{
		UserID: 1,
		Status: TransactionStatueExpense,
	}

	transactions, err := testQueries.ListTransactionByStatus(ctx, arg)
	require.NoError(t, err)

	for _, transaction := range(transactions) {
		require.NotEmpty(t, transaction)
	}

	fmt.Println(transactions)
}

func TestTotalAmount(t *testing.T) {
	transactions, err := testQueries.GetTotalAmount(ctx, 13)
	require.NoError(t, err)

	fmt.Println(transactions)
}

func TestListTransactionByCategory(t *testing.T) {
	arg := ListTransactionsByCategoryIDParams {
		CategoryID: 25,
		Limit: 5,
		Offset: 5,
	}
	transactions, err := testQueries.ListTransactionsByCategoryID(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, transactions)

	for _, transaction := range(transactions) {
		require.NotEmpty(t, transaction)
	}
}

func TestListTransactionByUserID(t *testing.T) {
	arg := ListTransactionsByUserIdParams {
		UserID: 1,
		Limit: 5,
		Offset: 5,
	}
	transactions, err := testQueries.ListTransactionsByUserId(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, transactions)

	for _, transaction := range(transactions) {
		require.NotEmpty(t, transaction)
	}
}
