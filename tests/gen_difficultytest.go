// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"math/big"

	"github.com/elastos/Elastos.ELA.SideChain.ETH/common"
	"github.com/elastos/Elastos.ELA.SideChain.ETH/common/math"
)

var _ = (*difficultyTestMarshaling)(nil)

func (d DifficultyTest) MarshalJSON() ([]byte, error) {
	type DifficultyTest struct {
		ParentTimestamp    *math.HexOrDecimal256 `json:"parentTimestamp"`
		ParentDifficulty   *math.HexOrDecimal256 `json:"parentDifficulty"`
		UncleHash          common.Hash           `json:"parentUncles"`
		CurrentTimestamp   *math.HexOrDecimal256 `json:"currentTimestamp"`
		CurrentBlockNumber math.HexOrDecimal64   `json:"currentBlockNumber"`
		CurrentDifficulty  *math.HexOrDecimal256 `json:"currentDifficulty"`
	}
	var enc DifficultyTest
	enc.ParentTimestamp = (*math.HexOrDecimal256)(d.ParentTimestamp)
	enc.ParentDifficulty = (*math.HexOrDecimal256)(d.ParentDifficulty)
	enc.UncleHash = d.UncleHash
	enc.CurrentTimestamp = (*math.HexOrDecimal256)(d.CurrentTimestamp)
	enc.CurrentBlockNumber = math.HexOrDecimal64(d.CurrentBlockNumber)
	enc.CurrentDifficulty = (*math.HexOrDecimal256)(d.CurrentDifficulty)
	return json.Marshal(&enc)
}

func (d *DifficultyTest) UnmarshalJSON(input []byte) error {
	type DifficultyTest struct {
		ParentTimestamp    *math.HexOrDecimal256 `json:"parentTimestamp"`
		ParentDifficulty   *math.HexOrDecimal256 `json:"parentDifficulty"`
		UncleHash          *common.Hash          `json:"parentUncles"`
		CurrentTimestamp   *math.HexOrDecimal256 `json:"currentTimestamp"`
		CurrentBlockNumber *math.HexOrDecimal64  `json:"currentBlockNumber"`
		CurrentDifficulty  *math.HexOrDecimal256 `json:"currentDifficulty"`
	}
	var dec DifficultyTest
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentTimestamp != nil {
		d.ParentTimestamp = (*big.Int)(dec.ParentTimestamp)
	}
	if dec.ParentDifficulty != nil {
		d.ParentDifficulty = (*big.Int)(dec.ParentDifficulty)
	}
	if dec.UncleHash != nil {
		d.UncleHash = *dec.UncleHash
	}
	if dec.CurrentTimestamp != nil {
		d.CurrentTimestamp = (*big.Int)(dec.CurrentTimestamp)
	}
	if dec.CurrentBlockNumber != nil {
		d.CurrentBlockNumber = uint64(*dec.CurrentBlockNumber)
	}
	if dec.CurrentDifficulty != nil {
		d.CurrentDifficulty = (*big.Int)(dec.CurrentDifficulty)
	}
	return nil
}
