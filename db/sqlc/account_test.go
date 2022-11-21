package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/harveysanders/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestGetAccount(t *testing.T) {
	account1 := createRandAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandAccount(t)
	newBalance := util.RandomAmount()

	updated, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      account.ID,
		Balance: newBalance,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, updated.ID, account.ID)
	require.Equal(t, updated.Owner, account.Owner)
	require.Equal(t, updated.Balance, newBalance)
	require.Equal(t, updated.Currency, account.Currency)
	require.NotZero(t, updated.CreatedAt)
	require.Equal(t, updated.CreatedAt, account.CreatedAt)
	//TODO: Updated At
}

func TestDeleteAccount(t *testing.T) {
	account := createRandAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	deleted, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deleted)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, a := range accounts {
		require.NotEmpty(t, a)
	}
}
