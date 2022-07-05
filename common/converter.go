package common

import (
	"math/big"
	"strings"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

func HexToAddress(address string) etherCommon.Address {
	return etherCommon.HexToAddress(address)
}

func WeiToFloat(amountWei *big.Int, decimal int) float64 {
	amountWeiFloat := new(big.Float).SetInt(amountWei)
	decimalBigFloat := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil))
	amount := new(big.Float).Quo(amountWeiFloat, decimalBigFloat)
	rawResult, _ := amount.Float64()
	return rawResult
}

func FloatToWei(amount float64, decimal int) *big.Int {
	amountFloat := big.NewFloat(amount)
	decimalBigInt := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	weiFloat := big.NewFloat(0).Mul(amountFloat, big.NewFloat(0).SetInt(decimalBigInt))
	amountInt, _ := weiFloat.Int(nil)
	return amountInt
}

func StringToWei(amount string, decimal int) *big.Int {
	amountBig, _ := big.NewFloat(0).SetString(amount)
	decimalBigInt := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	weiFloat := big.NewFloat(0).Mul(amountBig, big.NewFloat(0).SetInt(decimalBigInt))
	amountInt, _ := weiFloat.Int(nil)
	return amountInt
}

func EtherToDecimal(ether *big.Int, decimal int64) *big.Int {
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimal), nil)
	return new(big.Int).Mul(ether, pow)
}

func WeiToDecimal(wei *big.Int, decimal int64) *big.Int {
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimal), nil)
	return new(big.Int).Quo(wei, pow)
}

func DecimalToWei(value *big.Int, decimal int64) *big.Int {
	unit := big.NewInt(1e18)
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	result := new(big.Int).Mul(value, unit)
	return new(big.Int).Quo(result, pow)
}

func NormalizeAddress(addr etherCommon.Address) string {
	return strings.ToLower(addr.Hex())
}
