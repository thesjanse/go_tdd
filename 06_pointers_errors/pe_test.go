package pointers

import "testing"

func TestWallet(t *testing.T) {
	assert := func(t testing.TB, actual, expected Bitcoin) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual: '%d'; Expected: '%d'", actual, expected)
		}
	}

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	actual := wallet.Balance()
	expected := Bitcoin(10)
	assert(t, actual, expected)
}
