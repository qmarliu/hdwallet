package hdwallet_test

import (
	"fmt"
	"testing"

	"github.com/qmarliu/hdwallet"
)

const (
	account1 = "TLECxwhM8a52Yabsuadgqe5zJUFtecmRys"
	pvkey1   = "4839230f0cb046f33b3b2a9ba253d52d74dc134f1f0cd1c07e63d1f1f3e35387"

	account2 = "TXVZGJgmbaNnjL4xJj89mHAWPuh5RcGzyq"
	pvkey2   = "1e75584210196f6f637cfc92b0f248c9a25cec88572d7392abe0a6add40222ac"
)

func TestWallet(t *testing.T) {
	mnemonic := "situate embrace eagle shift start position gorilla history cry calm electric frost"
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		t.Fatal(err)
	}
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.TRON), hdwallet.AddressIndex(0))
	address, _ := wallet.GetAddress()
	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
	addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()
	fmt.Println(
		"address: ", account1,
		",\n result: ", address, addressP2WPKH, addressP2WPKHInP2SH,
	)
	fmt.Println("pvkey: ", pvkey1,
		",\n result:",
		wallet.GetKey().PrivateHex())
	if pvkey1 != wallet.GetKey().PrivateHex() {
		t.Fatal("pvkey fail")
	}

}
