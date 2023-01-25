package pointers

import "testing"

func TestWallet(t *testing.T) {
	assert := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()
		if actual != expected {
			t.Errorf("Actual: '%s'; Expected: '%s'", actual, expected)
		}
	}

	assertError := func(t testing.TB, actual, expected error) {
		t.Helper()
		if actual == nil {
			t.Fatal("Didn't got an error!")
		}
		if actual != expected {
			t.Errorf("Actual: '%q' but expected: '%q", actual, expected)
		}
	}

	t.Run("Deposit to empty wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assert(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw from non-empty wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		assert(t, wallet, Bitcoin(0))
	})

	t.Run("Withdraw more than current balance", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20))
		assertError(t, err, ErrorInsufficientFunds)
		assert(t, wallet, startingBalance)
	})
}
