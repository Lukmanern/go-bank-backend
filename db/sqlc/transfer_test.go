package db

import (
	"context"
	"testing"
	"time"

	"github.com/Lukmanern/go-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1ID, account2ID int64) Transfer {
	args := CreateTransferParams{
		FromAccountID: account1ID,
		ToAccountID: account2ID,
		Amount: util.RandomMoney(),
	}
	require.NotEqual(t, args.FromAccountID, args.ToAccountID)
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, transfer.Amount, args.Amount)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, account1ID)
	require.Equal(t, transfer.ToAccountID, account2ID)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	CreateRandomTransfer(t, account1.ID, account2.ID)
}

func TestGetTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	transfer1 := CreateRandomTransfer(t, account1.ID, account2.ID)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1.ID, account2.ID)
	}
	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 4,
		Offset: 4,
	}
	require.NotEqual(t, args.FromAccountID, args.ToAccountID)

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
