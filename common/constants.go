package common

import (
	"time"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	BlockRange  = int64(500)
	NodeTimeout = 10 * time.Second
	BlockTime   = 2 * time.Second
)

var (
	ZeroAddress = etherCommon.HexToAddress("0x0000000000000000000000000000000000000000")
)
