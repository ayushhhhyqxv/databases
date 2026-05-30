package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ayushhhhyqxv/databases/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandOwner(),
		Balance:  utils.RandMoney(),
		Currency: utils.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, args.Owner, acc.Owner)
	require.Equal(t, args.Balance, acc.Balance)
	require.Equal(t, args.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account_1 := createRandomAccount(t)
	account_2, err := testQueries.GetAccount(context.Background(), account_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account_2)

	require.Equal(t, account_1.ID, account_2.ID)
	require.Equal(t, account_1.Owner, account_2.Owner)
	require.Equal(t, account_1.Balance, account_2.Balance)
	require.Equal(t, account_1.Currency, account_2.Currency)
	require.WithinDuration(t, account_1.CreatedAt.Time, account_2.CreatedAt.Time, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account_1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account_1.ID,
		Balance: utils.RandMoney(),
	}

	account_2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NotEmpty(t, account_2)
	require.NoError(t,err)

	require.Equal(t, account_1.ID, account_2.ID)
	require.Equal(t, account_1.Owner, account_2.Owner)
	require.Equal(t, arg.Balance, account_2.Balance)
	require.Equal(t, account_1.Currency, account_2.Currency)
	require.WithinDuration(t, account_1.CreatedAt.Time, account_2.CreatedAt.Time, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account_1:=createRandomAccount(t)
	err:= testQueries.DeleteAccount(context.Background(),account_1.ID)
	require.NoError(t,err)

	account_2,err1 := testQueries.GetAccount(context.Background(),account_1.ID)
	require.Error(t,err1)
	require.EqualError(t,err1,sql.ErrNoRows.Error())
	require.Empty(t,account_2)
}

func TestListAccounts(t *testing.T){
	for i:=0;i<10;i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit : 5,
		Offset: 5,
	}

	accounts,err:=testQueries.ListAccounts(context.Background(),arg)

	require.NoError(t,err)
	require.Len(t,accounts,5)

	for _, account := range accounts {
		require.NotEmpty(t,account)
	}

}
