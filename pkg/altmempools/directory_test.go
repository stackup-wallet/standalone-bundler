package altmempools_test

import (
	"math/big"
	"testing"

	"github.com/stackup-wallet/stackup-bundler/internal/testutils"
	"github.com/stackup-wallet/stackup-bundler/pkg/altmempools"
)

func TestDirectoryHasSingleInvalidStorageAccessException(t *testing.T) {
	id := "1"
	alts := []*altmempools.AltMempoolConfig{
		{Id: id, Data: testutils.AltMempoolMock()},
	}
	dir, err := altmempools.New(testutils.ChainID, alts)
	if err != nil {
		t.Fatal("error initializing directory")
	}

	mempools := dir.HasInvalidStorageAccessException(
		"account",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
	)
	if len(mempools) != 1 || mempools[0] != id {
		t.Fatalf("got %v, want [1]", mempools)
	}
}

func TestDirectoryHasManyInvalidStorageAccessExceptions(t *testing.T) {
	id1 := "1"
	id2 := "2"
	alts := []*altmempools.AltMempoolConfig{
		{Id: id1, Data: testutils.AltMempoolMock()},
		{Id: id2, Data: testutils.AltMempoolMock()},
	}
	dir, err := altmempools.New(testutils.ChainID, alts)
	if err != nil {
		t.Fatal("error initializing directory")
	}

	mempools := dir.HasInvalidStorageAccessException(
		"account",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
	)
	if len(mempools) != 2 || mempools[0] != id1 && mempools[1] != id2 {
		t.Fatalf("got %v, want [1 2]", mempools)
	}
}

func TestDirectoryHasNoInvalidStorageAccessExceptions(t *testing.T) {
	id := "1"
	alts := []*altmempools.AltMempoolConfig{
		{Id: id, Data: testutils.AltMempoolMock()},
	}
	dir, err := altmempools.New(testutils.ChainID, alts)
	if err != nil {
		t.Fatal("error initializing directory")
	}

	mempools := dir.HasInvalidStorageAccessException(
		"paymaster",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
	)
	if len(mempools) != 0 {
		t.Fatalf("got %v, want []", mempools)
	}
}

func TestDirectoryIncompatibleChain(t *testing.T) {
	id := "1"
	chain := big.NewInt(2)
	alts := []*altmempools.AltMempoolConfig{
		{Id: id, Data: testutils.AltMempoolMock()},
	}
	dir, err := altmempools.New(chain, alts)
	if err != nil {
		t.Fatal("error initializing directory")
	}

	mempools := dir.HasInvalidStorageAccessException(
		"account",
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000",
	)
	if len(mempools) != 0 {
		t.Fatalf("got %v, want []", mempools)
	}
}
