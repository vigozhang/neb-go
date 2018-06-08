package utils

import (
	"math/big"
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/satori/go.uuid"
)

func GetIntWithDefault(value int, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func GetStringWithDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func GetBytesWithDefault(value []byte, defaultValue []byte) []byte {
	if value == nil {
		return defaultValue
	}
	return value
}

func UuidStringFromBytes(input []byte) string {
	uuid, err := uuid.FromBytes(input)
	if err != nil {
		log.Println("uuid error:", err)
	}
	return uuid.String()
}

func BytesFromHexString(hexstring string) []byte {
	intvalue, _ := big.NewInt(0).SetString(hexstring, 16)
	return intvalue.Bytes()
}

func BigIntFromHexString(hexstring string) *big.Int {
	intvalue, _ := big.NewInt(0).SetString(hexstring, 16)
	return intvalue
}

func PadToBigEndian(value []byte, digit int) []byte {
	bytes := digit / 8

	i := len(value)

	if i > bytes {
		return value[(i - bytes):]
	}

	for i < bytes {
		value = append([]byte{0x00}, value...)
		i++
	}
	return value
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func UInt64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func IntToBytes(i int) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func EncodeToJsonBytes(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}

func To128BitBytes(u *big.Int) ([16]byte, error) {
	var res [16]byte
	bs := u.Bytes()
	l := len(bs)
	if l == 0 {
		return res, nil
	}
	idx := 16 - len(bs)
	if idx < 16 {
		copy(res[idx:], bs)
	}
	return res, nil
}

func To128BitSlice(u *big.Int) []byte {
	bytes, err := To128BitBytes(u)
	if err != nil {
		log.Println("To128BitSlice error:", err)
	}
	return bytes[:]
}
