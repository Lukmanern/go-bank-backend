package db

import (
	"context"
	"testing"
	"time"

	"github.com/Lukmanern/go-bank-backend/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:	util.RandomOwner(),
		Balance: 	util.RandomMoney(),
		Currency: 	util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account_1 := CreateRandomAccount(t)
	account_2, err := testQueries.GetAccount(context.Background(), account_1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account_2)
	
	require.Equal(t, account_1.ID, account_2.ID)
	require.Equal(t, account_1.Owner, account_2.Owner)
	require.Equal(t, account_1.Balance, account_2.Balance)
	require.Equal(t, account_1.Currency, account_2.Currency)
	require.Equal(t, account_1.CreatedAt, account_2.CreatedAt)
	require.WithinDuration(t, account_1.CreatedAt, account_2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account_1 := CreateRandomAccount(t)
	account_2, err := testQueries.GetAccount(context.Background(), account_1.ID)
	require.NoError(t, err)
	
	updateData := UpdateAccountParams{
		ID: account_2.ID,
		Balance: 999,
	}
	account_2, err = testQueries.UpdateAccount(context.Background(), updateData)
	
	require.NoError(t, err)
	require.Equal(t, account_2.ID, account_1.ID)
	require.Equal(t, account_2.Owner, account_1.Owner)
	require.Equal(t, account_2.Balance, updateData.Balance)
	require.Equal(t, account_2.Currency, account_1.Currency)
	require.Equal(t, account_2.CreatedAt, account_1.CreatedAt)
}

// func TestAccountsList(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		CreateRandomAccount(t)
// 	}
// }