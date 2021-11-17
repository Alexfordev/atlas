// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package genesis

import (
	"encoding/json"
	"github.com/mapprotocol/atlas/helper/decimal/bigintstr"
	"github.com/mapprotocol/atlas/helper/decimal/fixed"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var _ = (*ReserveParametersMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r ReserveParameters) MarshalJSON() ([]byte, error) {
	type ReserveParameters struct {
		TobinTaxStalenessThreshold uint64               `json:"tobinTaxStalenessThreshold"`
		DailySpendingRatio         *fixed.Fixed         `json:"dailySpendingRatio"`
		AssetAllocations           AssetAllocationList  `json:"assetAllocations"`
		TobinTax                   *fixed.Fixed         `json:"tobinTax"`
		TobinTaxReserveRatio       *fixed.Fixed         `json:"tobinTaxReserveRatio"`
		Spenders                   []common.Address     `json:"spenders"`
		OtherAddresses             []common.Address     `json:"otherAddresses"`
		InitialBalance             *bigintstr.BigIntStr `json:"initialBalance"`
		FrozenAssetsStartBalance   *bigintstr.BigIntStr `json:"frozenAssetsStartBalance"`
		FrozenAssetsDays           uint64               `json:"frozenAssetsDays"`
	}
	var enc ReserveParameters
	enc.TobinTaxStalenessThreshold = r.TobinTaxStalenessThreshold
	enc.DailySpendingRatio = r.DailySpendingRatio
	enc.AssetAllocations = r.AssetAllocations
	enc.TobinTax = r.TobinTax
	enc.TobinTaxReserveRatio = r.TobinTaxReserveRatio
	enc.Spenders = r.Spenders
	enc.OtherAddresses = r.OtherAddresses
	enc.InitialBalance = (*bigintstr.BigIntStr)(r.InitialBalance)
	enc.FrozenAssetsStartBalance = (*bigintstr.BigIntStr)(r.FrozenAssetsStartBalance)
	enc.FrozenAssetsDays = r.FrozenAssetsDays
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *ReserveParameters) UnmarshalJSON(input []byte) error {
	type ReserveParameters struct {
		TobinTaxStalenessThreshold *uint64              `json:"tobinTaxStalenessThreshold"`
		DailySpendingRatio         *fixed.Fixed         `json:"dailySpendingRatio"`
		AssetAllocations           *AssetAllocationList `json:"assetAllocations"`
		TobinTax                   *fixed.Fixed         `json:"tobinTax"`
		TobinTaxReserveRatio       *fixed.Fixed         `json:"tobinTaxReserveRatio"`
		Spenders                   []common.Address     `json:"spenders"`
		OtherAddresses             []common.Address     `json:"otherAddresses"`
		InitialBalance             *bigintstr.BigIntStr `json:"initialBalance"`
		FrozenAssetsStartBalance   *bigintstr.BigIntStr `json:"frozenAssetsStartBalance"`
		FrozenAssetsDays           *uint64              `json:"frozenAssetsDays"`
	}
	var dec ReserveParameters
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.TobinTaxStalenessThreshold != nil {
		r.TobinTaxStalenessThreshold = *dec.TobinTaxStalenessThreshold
	}
	if dec.DailySpendingRatio != nil {
		r.DailySpendingRatio = dec.DailySpendingRatio
	}
	if dec.AssetAllocations != nil {
		r.AssetAllocations = *dec.AssetAllocations
	}
	if dec.TobinTax != nil {
		r.TobinTax = dec.TobinTax
	}
	if dec.TobinTaxReserveRatio != nil {
		r.TobinTaxReserveRatio = dec.TobinTaxReserveRatio
	}
	if dec.Spenders != nil {
		r.Spenders = dec.Spenders
	}
	if dec.OtherAddresses != nil {
		r.OtherAddresses = dec.OtherAddresses
	}
	if dec.InitialBalance != nil {
		r.InitialBalance = (*big.Int)(dec.InitialBalance)
	}
	if dec.FrozenAssetsStartBalance != nil {
		r.FrozenAssetsStartBalance = (*big.Int)(dec.FrozenAssetsStartBalance)
	}
	if dec.FrozenAssetsDays != nil {
		r.FrozenAssetsDays = *dec.FrozenAssetsDays
	}
	return nil
}
