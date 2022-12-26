package modules

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stackup-wallet/stackup-bundler/internal/testutils"
)

// TestAddDepositInfoToCtx verifies that stake info can be added to a context and later retrieved.
func TestAddDepositInfoToCtx(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	ctx := NewUserOpHandlerContext(op, common.HexToAddress("0x"), big.NewInt(1))

	entity := op.GetFactory()
	dep := testutils.StakedDepositInfo
	ctx.AddDepositInfo(entity, dep)

	if ctx.GetDepositInfo(entity) != dep {
		t.Fatal("Retrieved deposit info does not equal the original")
	}
}

// TestGetNilDepositInfoFromCtx calls (c *UserOpHandlerCtx).GetDepositInfo on an address that has not been
// set. Expects nil.
func TestGetNilDepositInfoFromCtx(t *testing.T) {
	op := testutils.MockValidInitUserOp()
	ctx := NewUserOpHandlerContext(op, common.HexToAddress("0x"), big.NewInt(1))

	if dep := ctx.GetDepositInfo(op.GetFactory()); dep != nil {
		t.Fatalf("got %+v, want nil", dep)
	}
}
