package account

import (
	"testing"
	"math/big"
	"github.com/satori/go.uuid"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount()
	if account != nil {
		t.Logf("NewAccount success,account address is %s", account.GetAddressString())
	} else {
		t.Error("NewAccount failed")
	}
}

func TestIsValidAddress(t *testing.T) {
	addressNormal := "n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Lf"
	valid := IsValidAddress(addressNormal)
	if valid {
		t.Log("n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Lf is valid")
	} else {
		t.Error("n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Lf valid failed")
	}

	addressContract := "n1pfgMmLKSue8N8vkzcdQeSLwrdHWGG73d4"
	valid = IsValidAddress(addressContract)
	if valid {
		t.Log("n1pfgMmLKSue8N8vkzcdQeSLwrdHWGG73d4 is valid")
	} else {
		t.Error("n1pfgMmLKSue8N8vkzcdQeSLwrdHWGG73d4 valid failed")
	}

	addressInvalid1 := "xbhkiSljnDJklDNsdf"
	valid = IsValidAddress(addressInvalid1)
	if !valid {
		t.Log("xbhkiSljnDJklDNsdf is not valid")
	} else {
		t.Error("xbhkiSljnDJklDNsdf valid failed")
	}

	addressInvalid2 := "n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Ls"
	valid = IsValidAddress(addressInvalid2)
	if !valid {
		t.Log("n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Ls is not valid")
	} else {
		t.Error("n1aR9eGmPn6KikNU2bBRKhwxgrtzR2Le9Ls valid failed")
	}
}

func TestAccount_ToKeyString(t *testing.T) {
	acc := NewAccount()

	accPriv, _ := big.NewInt(0).SetString("ac3773e06ae74c0fa566b0e421d4e391333f31aef90b383f0c0e83e4873609d6", 16)
	acc.SetPrivateKey(accPriv.Bytes())

	salt, _ := big.NewInt(0).SetString("acded3c4c3f7655cbcd38eea82fdc939251354f21c14cd7faec05c73aa9f70e3", 16)
	iv, _ := big.NewInt(0).SetString("e90bd4e307153dc61632ce428606620c", 16)

	id, _ := uuid.FromString("0c560f8c-4c0e-4584-9a26-fb57d1dfd0e5")

	keyopts := KeyOptions{
		Salt: salt.Bytes(),
		Iv:   iv.Bytes(),
		Uuid: id.Bytes(),
	}

	accToKeyString, err := acc.ToKeyString("password0192837465DlK", &keyopts)

	if err != nil {
		t.Error("TestAccount_ToKeyString failed")
	} else {
		t.Logf("TestAccount_ToKeyString result: %s", accToKeyString)
	}
}

func TestAccount_FromKey(t *testing.T) {
	keyjson := `{"version":4,"id":"0c560f8c-4c0e-4584-9a26-fb57d1dfd0e5","address":"n1LfrjZzXDCcHhNV2r6F6dUS5Zxi7P8xC45","crypto":{"ciphertext":"f49c508fea24f211b708eced749e57756af046207e896f4fb5332d664d33d531","cipherparams":{"iv":"e90bd4e307153dc61632ce428606620c"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"acded3c4c3f7655cbcd38eea82fdc939251354f21c14cd7faec05c73aa9f70e3","n":4096,"r":8,"p":1,"c":0,"prf":""},"mac":"b87684d55dda7a4f76dcf3f30c98ac42f1958e6e4a9b80d63276aeb6c41a93d7","machash":"sha3256"}}`
	acc := NewAccount()
	fromAcc, err := acc.FromKey(keyjson, "password0192837465DlK", true)

	if err != nil {
		t.Error("TestAccount_FromKey failed")
	} else {
		t.Logf("TestAccount_FromKey address:%s , private key: %s, pub key: %s",
			fromAcc.GetAddressString(),
			fromAcc.GetPrivateKeyString(),
			fromAcc.GetPublicKeyString())
	}

}
