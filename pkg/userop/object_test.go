package userop_test

import (
	"math/big"
	"testing"

	"github.com/stackup-wallet/stackup-bundler/internal/testutils"
)

// TestUserOperationGetGasPrice verifies that (*UserOperation).GetGasPrice returns the correct effective gas
// price given a base fee.
func TestUserOperationGetGasPrice(t *testing.T) {
	bf := big.NewInt(3)
	op := testutils.MockValidInitUserOp()

	// basefee + MPF > MF
	want := big.NewInt(4)
	op.MaxFeePerGas = big.NewInt(4)
	op.MaxPriorityFeePerGas = big.NewInt(3)
	if op.GetGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetGasPrice(bf).Int64(), want.Int64())
	}

	// basefee + MPF == MF
	want = big.NewInt(5)
	op.MaxFeePerGas = big.NewInt(5)
	op.MaxPriorityFeePerGas = big.NewInt(2)
	if op.GetGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetGasPrice(bf).Int64(), want.Int64())
	}

	// basefee + MPF < MF
	want = big.NewInt(4)
	op.MaxFeePerGas = big.NewInt(6)
	op.MaxPriorityFeePerGas = big.NewInt(1)
	if op.GetGasPrice(bf).Cmp(want) != 0 {
		t.Fatalf("got %d, want %d", op.GetGasPrice(bf).Int64(), want.Int64())
	}
}
