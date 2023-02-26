package address

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetEthAddress(pkHex string) string {
	privateKey, _ := crypto.HexToECDSA(pkHex)

	if privateKey != nil {
		publicKey := privateKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		return strings.ToLower(address[2:])
	}
	return ""
}

func HexToPrivateKey(hexKey string) (*btcec.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}

	privateKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	return privateKey, nil
}

func HexToPubKeyHash(hexKey string) []byte {
	privateKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil
	}

	privateKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	if err != nil {
		return nil
	}
	pubKey := privateKey.PubKey()
	pubKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())
	return pubKeyHash
}

func GetP2PKHAddress(pubKeyHash []byte) string {
	address, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error in getP2PKHAddress: ", err)
	}
	return strings.ToLower(address.EncodeAddress())
}

func GetBech32Address(pubKeyHash []byte) string {
	bech32Address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error in getP2PKHAddress: ", err)
	}
	return strings.ToLower(bech32Address.EncodeAddress())
}

func P2shAddress(hexKey string) string {
	// Convert hexadecimal string to byte array
	privateKey, err := hex.DecodeString(hexKey)
	if err != nil {
		return ""
	}

	// Convert the byte array to a private key
	_, pubKey := btcec.PrivKeyFromBytes(privateKey)

	// Get the hash of the public key
	publicKeyHash := btcutil.Hash160(pubKey.SerializeCompressed())

	// Concatenate the network prefix and public key hash to form the payload
	payload := []byte{0x05}
	payload = append(payload, publicKeyHash...)

	// Perform SHA-256 hash on the payload
	h := sha256.New()
	h.Write(payload)
	checksum := h.Sum(nil)

	// Perform another SHA-256 hash on the first hash
	h = sha256.New()
	h.Write(checksum)
	checksum = h.Sum(nil)[:4]

	// Concatenate the payload and checksum to form the full payload
	fullPayload := append(payload, checksum...)

	// Convert the full payload to base58 encoded address
	address := base58.Encode(fullPayload)

	return strings.ToLower(address)
}
