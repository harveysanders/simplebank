package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := createRandAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    100,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, int64(100), entry.Amount)
	require.NotZero(t, entry.CreatedAt)
	require.WithinDuration(t, time.Now(), entry.CreatedAt, time.Second)
	require.NotZero(t, account.ID)
}
