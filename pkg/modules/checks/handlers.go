package checks

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stackup-wallet/stackup-bundler/pkg/entrypoint"
	"github.com/stackup-wallet/stackup-bundler/pkg/mempool"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules"
)

// StandaloneClient returns a UserOpHandler that relies on a given ethClient to run through all the standard
// client checks as specified in EIP-4337. This should be the first module in the stack.
func StandaloneClient(eth *ethclient.Client, mem *mempool.Interface) modules.UserOpHandlerFunc {
	return func(ctx *modules.UserOpHandlerCtx) error {
		entryPoint, err := entrypoint.NewEntrypoint(ctx.EntryPoint, eth)
		if err != nil {
			return err
		}

		if err := checkSender(eth, ctx.UserOp); err != nil {
			return err
		}
		if err := checkVerificationGasLimits(eth, ctx.UserOp); err != nil {
			return err
		}
		if err := checkPaymasterAndData(eth, ctx.UserOp, entryPoint); err != nil {
			return err
		}
		if err := checkCallGasLimit(eth, ctx.UserOp); err != nil {
			return err
		}
		if err := checkFeePerGas(eth, ctx.UserOp); err != nil {
			return err
		}
		if err := checkDuplicates(mem, ctx.UserOp, ctx.EntryPoint); err != nil {
			return err
		}

		return nil
	}
}
