package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("[wallet-test error] got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))

		want := Bitcoin(10.0)
		assertBalance(t, wallet, want)

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		err := wallet.Withdraw(Bitcoin(5.0))

		if err != nil {
			t.Errorf("[wallet-test error] should not return error")
		}

		want := Bitcoin(5.0)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw-error", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		err := wallet.Withdraw(Bitcoin(15.0))
		if err == nil {
			t.Errorf("[wallet-test error] expected error")
		}
	})
}
