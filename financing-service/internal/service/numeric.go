package service

import (
	"errors"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func numericToBigFloat(value pgtype.Numeric) (*big.Float, error) {
	if !value.Valid || value.Int == nil {
		return nil, errors.New("numeric value is required")
	}
	if value.NaN || value.InfinityModifier != pgtype.Finite {
		return nil, errors.New("numeric value must be finite")
	}

	result := new(big.Float).SetInt(value.Int)
	if value.Exp > 0 {
		power := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(value.Exp)), nil)
		result.Mul(result, new(big.Float).SetInt(power))
	} else if value.Exp < 0 {
		power := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(-value.Exp)), nil)
		result.Quo(result, new(big.Float).SetInt(power))
	}

	return result, nil
}