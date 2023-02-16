package db

import (
	"context"
	"testing"

	"github.com/Lukmanern/go-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, accountID int64) Entry {
	args := CreateEntryParams{
		AccountID: accountID,
		Amount: util.RandomInt(-1000, 1000),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)
	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomEntry(t, account.ID)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.Error(t, err)
}

func TestGetEntry(t *testing.T) {
	account := CreateRandomAccount(t)
	entry := CreateRandomEntry(t, account.ID)
	require.NotEmpty(t, entry)
}

func TestListEntries(t *testing.T) {
	account := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, account.ID)
	}

	args := ListEntriesParams{
		AccountID: account.ID,
		Limit: 4,
		Offset: 4,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}