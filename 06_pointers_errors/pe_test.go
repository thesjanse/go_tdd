package pointers

import "testing"

func TestWallet(t *testing.T) {
	assert := func(t testing.TB, actual, expected Bitcoin) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual: '%s'; Expected: '%s'", actual, expected)
		}
	}

	t.Run("Deposit to empty wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)
		assert(t, actual, expected)
	})

	t.Run("Withdraw from non-empty wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(0)
		assert(t, actual, expected)
	})
}
