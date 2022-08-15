package tools

import (
	"bytes"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/bech32"
	"golang.org/x/crypto/sha3"
	"regexp"
	"strings"
)

const addressChecksumLen = 4

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	if address == "" || len(address) < 5 {
		return false
	}
	if strings.HasPrefix(address, "0x") {
		address = address[2:]
	}

	if IsBtcTestNetAddress(address) || IsNotSupportBtcAddress(address) {
		return false
	}

	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	if bytes.Compare(actualChecksum, targetChecksum) == 0 {
		return true
	} else {
		_, _, err := bech32.Decode(address)
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

// 判断是否为btc testnet地址see : https://en.bitcoin.it/wiki/List_of_address_prefixes
var BTC_TEST_NET_ADDRESS_PREFIX = []string{"m", "n", "2", "9", "c", "tpub", "tprv", "tb1"}

func IsBtcTestNetAddress(address string) bool {
	for _, address_prefix := range BTC_TEST_NET_ADDRESS_PREFIX {
		if strings.HasPrefix(address, address_prefix) {
			return true
		}
	}

	return false
}

// 判断是否为btc 不支持的地址前缀，see : https://en.bitcoin.it/wiki/List_of_address_prefixes
var BTC_NOT_SUPPORT_ADDRESS_PREFIX = []string{"5", "K", "L", "M", "xpub", "xprv"}

func IsNotSupportBtcAddress(address string) bool {
	for _, address_prefix := range BTC_NOT_SUPPORT_ADDRESS_PREFIX {
		if strings.HasPrefix(address, address_prefix) {
			return true
		}
	}

	return false
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// 电子邮箱正则
var EMAIL_REGEXP = regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)

func ValidateEmail(email string) bool {
	return EMAIL_REGEXP.MatchString(email)
}

// see https://github.com/ChainSafe/web3.js/blob/2279a67e07/packages/web3-utils/src/utils.js checkAddressChecksum checkAddressChecksum
var ETH_REGEXP = regexp.MustCompile(`^(0x|0X)?[0-9a-fA-F]{40}$`)
var LOWER_ETH_REGEXP = regexp.MustCompile(`^(0x|0X)?[0-9a-f]{40}$`)
var UPPER_ETH_REGEXP = regexp.MustCompile(`^(0x|0X)?[0-9A-F]{40}$`)

func ValidateEthAddress(ethAddress string) bool {
	if ethAddress == "" || len(ethAddress) < 40 {
		return false
	}
	if ETH_REGEXP.MatchString(ethAddress) {
		noPrefixAddress := ethAddress[2:]
		ethAddressLower := strings.ToLower(noPrefixAddress)
		sha := sha3.NewLegacyKeccak256()
		sha.Write([]byte(ethAddressLower))
		hash := sha.Sum(nil)

		result := []byte(ethAddressLower)
		for i := 0; i < len(result); i++ {
			hashByte := hash[i/2]
			if i%2 == 0 {
				hashByte = hashByte >> 4
			} else {
				hashByte &= 0xf
			}
			if result[i] > '9' && hashByte > 7 {
				result[i] -= 32
			}
		}

		if string(result) == noPrefixAddress ||
			LOWER_ETH_REGEXP.MatchString(noPrefixAddress) && strings.ToLower(string(result)) == noPrefixAddress ||
			UPPER_ETH_REGEXP.MatchString(noPrefixAddress) && strings.ToUpper(string(result)) == noPrefixAddress {
			return true
		}
	}

	return false
}
