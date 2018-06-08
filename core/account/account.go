package account

import (
	"encoding/hex"
	"encoding/json"
	"bytes"
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"

	"github.com/vigozhang/neb-go/utils/secp256k1"
	"github.com/vigozhang/neb-go/utils/hash"
	"github.com/vigozhang/neb-go/utils/base58"
	"github.com/vigozhang/neb-go/utils"
)

type Account struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    []byte
}

type KeyOptions struct {
	Salt   []byte
	Iv     []byte
	Kdf    string
	Dklen  int
	C      int
	N      int
	R      int
	P      int
	Cipher string
	Uuid   []byte
}

type Key struct {
	Version int    `json:"version"`
	Id      string `json:"id"`
	Address string `json:"address"`
	Crypto  Crypto `json:"crypto"`
}

type Crypto struct {
	Ciphertext   string       `json:"ciphertext"`
	Cipherparams CipherParams `json:"cipherparams"`
	Cipher       string       `json:"cipher"`
	Kdf          string       `json:"kdf"`
	Kdfparams    KdfParams    `json:"kdfparams"`
	Mac          string       `json:"mac"`
	Machash      string       `json:"machash"`
}

type CipherParams struct {
	Iv string `json:"iv"`
}

type KdfParams struct {
	Dklen int    `json:"dklen"`
	Salt  string `json:"salt"`
	N     int    `json:"n"`
	R     int    `json:"r"`
	P     int    `json:"p"`
	C     int    `json:"c"`
	Prf   string `json:"prf"`
}

const (
	UncompressedPublicKeyPrefix = 0x04
	AddressPrefix               = 0x19
	NormalType                  = 0x57
	ContractType                = 0x58

	AddressLength       = 26
	AddressStringLength = 35

	KeyVersion3       = 3
	KeyCurrentVersion = 4
)

func NewAccount() *Account {
	privateKey := secp256k1.NewSeckey()
	publicKey := privateToPublicKey(privateKey)
	address := addressFromPublicKey(publicKey)

	return &Account{privateKey, publicKey, address}

}

func IsValidAddress(address string) bool {
	if len(address) != AddressStringLength {
		return false
	}

	addressByte := base58.Decode(address)
	if len(addressByte) != AddressLength {
		return false
	}

	if bytes.Compare(addressByte[0:1], []byte{AddressPrefix}) != 0 {
		return false
	}

	typeNum := addressByte[1:2]
	if bytes.Compare(typeNum, []byte{NormalType}) != 0 &&
		bytes.Compare(typeNum, []byte{ContractType}) != 0 {
		return false
	}
	content := addressByte[0:22]
	checksum := addressByte[len(addressByte)-4:]

	return bytes.Compare(hash.Sha3256(content)[0:4], checksum) == 0
}

func FromAddress(address string) (*Account, error) {
	acc := Account{}
	if IsValidAddress(address) {
		acc.Address = base58.Decode(address)
		return &acc, nil
	}
	return nil, errors.New("invalid address")
}

func (acc *Account) SetPrivateKey(privateKey []byte) {
	acc.PrivateKey = privateKey
	acc.PublicKey = privateToPublicKey(privateKey)
	acc.Address = addressFromPublicKey(acc.PublicKey)
}

func (acc *Account) GetPrivateKey() []byte {
	return acc.PrivateKey
}

func (acc *Account) GetPrivateKeyString() string {
	return hex.EncodeToString(acc.PrivateKey)
}

func (acc *Account) GetPublicKey() []byte {
	return acc.PublicKey
}

func (acc *Account) GetPublicKeyString() string {
	return hex.EncodeToString(acc.PublicKey)
}

func (acc *Account) GetAddress() []byte {
	return acc.Address
}

func (acc *Account) GetAddressString() string {
	return base58.Encode(acc.Address)
}

func (acc *Account) ToKey(password string, opts *KeyOptions) (*Key, error) {
	salt := utils.GetBytesWithDefault(opts.Salt, utils.RandomCSPRNG(32))
	iv := utils.GetBytesWithDefault(opts.Iv, utils.RandomCSPRNG(16))
	kdf := utils.GetStringWithDefault(opts.Kdf, "scrypt")

	dklen := utils.GetIntWithDefault(opts.Dklen, 32)
	kdfparams := KdfParams{Dklen: dklen, Salt: hex.EncodeToString(salt)}

	var derivedKey []byte

	if kdf == "pbkdf2" {
		kdfparams.C = utils.GetIntWithDefault(opts.C, 262144)
		kdfparams.Prf = "hmac-sha256"

		derivedKey = pbkdf2.Key([]byte(password), salt, kdfparams.C, kdfparams.Dklen, sha256.New)

	} else if kdf == "scrypt" {
		kdfparams.N = utils.GetIntWithDefault(opts.N, 4096)
		kdfparams.R = utils.GetIntWithDefault(opts.R, 8)
		kdfparams.P = utils.GetIntWithDefault(opts.P, 1)
		derivedKey, _ = scrypt.Key([]byte(password), salt, kdfparams.N, kdfparams.R, kdfparams.P, kdfparams.Dklen)
	} else {
		return nil, errors.New("unsupported kdf")
	}

	cipher := utils.GetStringWithDefault(opts.Cipher, "aes-128-ctr")

	ciphertext, err := utils.OpensslEncrypt(acc.PrivateKey, cipher, derivedKey[0:16], iv)
	if err != nil {
		return nil, err
	}

	maccontent := append(derivedKey[16:32], ciphertext...)
	maccontent = append(maccontent, iv...)
	maccontent = append(maccontent, []byte(cipher)...)

	mac := hash.Sha3256(maccontent)

	cipherparams := CipherParams{hex.EncodeToString(iv)}
	crypto := Crypto{
		Ciphertext:   hex.EncodeToString(ciphertext),
		Cipherparams: cipherparams,
		Cipher:       cipher,
		Kdf:          kdf,
		Kdfparams:    kdfparams,
		Mac:          hex.EncodeToString(mac),
		Machash:      "sha3256",
	}

	key := Key{
		Version: KeyCurrentVersion,
		Id:      utils.UuidStringFromBytes(utils.GetBytesWithDefault(opts.Uuid, utils.RandomCSPRNG(16))),
		Address: acc.GetAddressString(),
		Crypto:  crypto,
	}

	return &key, nil
}

func (acc *Account) ToKeyString(password string, opts *KeyOptions) (string, error) {
	key, err := acc.ToKey(password, opts)
	if err != nil {
		return "", err
	}
	jsonbytes, err := json.Marshal(key)
	if err != nil {
		return "", err
	}
	return string(jsonbytes), nil
}

func (acc *Account) FromKey(input string, password string, nonStrict bool) (*Account, error) {
	key := Key{}
	err := json.Unmarshal([]byte(input), &key)
	if err != nil {
		return nil, err
	}

	if key.Version != KeyVersion3 && key.Version != KeyCurrentVersion {
		return nil, errors.New("not supported wallet version")
	}

	kdfparams := key.Crypto.Kdfparams

	var derivedKey []byte

	if key.Crypto.Kdf == "scrypt" {
		derivedKey, err = scrypt.Key([]byte(password), utils.BytesFromHexString(kdfparams.Salt), kdfparams.N, kdfparams.R, kdfparams.P, kdfparams.Dklen)
		if err != nil {
			return nil, err
		}
	} else if key.Crypto.Kdf == "pbkdf2" {
		if kdfparams.Prf != "hmac-sha256" {
			return nil, errors.New("unsupported parameters to PBKDF2")
		}
		derivedKey = pbkdf2.Key([]byte(password), utils.BytesFromHexString(kdfparams.Salt), kdfparams.C, kdfparams.Dklen, sha256.New)
	} else {
		return nil, errors.New("unsupported key derivation scheme")
	}

	ciphertext := utils.BytesFromHexString(key.Crypto.Ciphertext)

	maccontent := append(derivedKey[16:32], ciphertext...)
	if key.Version == KeyCurrentVersion {
		maccontent = append(maccontent, utils.BytesFromHexString(key.Crypto.Cipherparams.Iv)...)
		maccontent = append(maccontent, []byte(key.Crypto.Cipher)...)
	}

	mac := hash.Sha3256(maccontent)

	if hex.EncodeToString(mac) != key.Crypto.Mac {
		return nil, errors.New("key derivation failed - possibly wrong passphrase")
	}

	seed, _ := utils.OpensslDecrypt(ciphertext, key.Crypto.Cipher, derivedKey[0:16], utils.BytesFromHexString(key.Crypto.Cipherparams.Iv))

	for len(seed) < 32 {
		seed = append([]byte{0x00}, seed...)
	}

	acc.SetPrivateKey(seed)
	return acc, nil
}

func addressFromPublicKey(publicKey []byte) []byte {
	publicKey = append([]byte{UncompressedPublicKeyPrefix}, publicKey...)

	pubHash := hash.Sha3256(publicKey)
	pubHash = hash.Ripemd160(pubHash)

	content := append([]byte{AddressPrefix}, []byte{NormalType}...)
	content = append(content, pubHash...)

	checksum := hash.Sha3256(content)[0:4]
	address := append(content, checksum...)

	return address
}

func privateToPublicKey(privateKey []byte) []byte {
	publicKey, _ := secp256k1.GetPublicKey(privateKey)
	return publicKey[1:]
}
