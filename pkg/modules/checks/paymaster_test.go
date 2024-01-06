package checks

import (
	"testing"

	"github.com/stackup-wallet/stackup-bundler/internal/testutils"
)

// TestNilPaymasterAndData calls checks.ValidatePaymasterAndData with no paymaster set. Expects nil.
func TestNilPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = []byte{}
	err := ValidatePaymasterAndData(op, testutils.MockGetCodeZero, testutils.NonStakedZeroDepositInfo)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}

// TestBadPaymasterAndData calls checks.ValidatePaymasterAndData with invalid paymaster set. Expects error.
func TestBadPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = []byte("1234")
	err := ValidatePaymasterAndData(op, testutils.MockGetCodeZero, testutils.NonStakedZeroDepositInfo)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestZeroByteCodePaymasterAndData calls checks.ValidatePaymasterAndData with paymaster contract not
// deployed. Expects error.
func TestZeroByteCodePaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.MockGetCodeZero, testutils.NonStakedZeroDepositInfo)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestNonStakedZeroDepositPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is not
// staked with zero deposit. Expects error.
func TestNonStakedZeroDepositPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.MockGetCode, testutils.NonStakedZeroDepositInfo)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestZeroDepositPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is staked but
// with zero deposit. Expects error.
func TestZeroDepositPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.MockGetCode, testutils.StakedZeroDepositInfo)

	if err == nil {
		t.Fatal("got nil, want err")
	}
}

// TestNotStakedPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is not staked and
// has sufficient deposit. Expects nil.
func TestNotStakedPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.MockGetCode, testutils.NonStakedDepositInfo)

	if err != nil {
		t.Fatalf("got %v, want nil", err)
	}
}

// TestPaymasterAndData calls checks.ValidatePaymasterAndData with paymaster that is staked and has sufficient
// deposit. Expects nil.
func TestPaymasterAndData(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	op.PaymasterAndData = op.Sender.Bytes()
	err := ValidatePaymasterAndData(op, testutils.MockGetCode, testutils.StakedDepositInfo)

	if err != nil {
		t.Fatalf("got err %v, want nil", err)
	}
}
