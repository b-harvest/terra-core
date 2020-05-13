package keeper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/terra-project/core/x/wasm/internal/types"
)

func TestNewKeeper(t *testing.T) {
	_, _, keeper := CreateTestInput(t)
	require.NotNil(t, keeper)
}

func TestCodeInfo(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t)

	codeID := uint64(1)
	creatorAddr := addrFromUint64(codeID)
	expected := types.NewCodeInfo([]byte{1, 2, 3}, creatorAddr)
	keeper.SetCodeInfo(ctx, 1, expected)

	as, err := keeper.GetCodeInfo(ctx, codeID)
	require.NoError(t, err)
	require.Equal(t, expected, as)
}

func TestContractInfo(t *testing.T) {
	ctx, _, keeper := CreateTestInput(t)

	_, _, alice := keyPubAddr()
	_, _, bob := keyPubAddr()

	codeID := uint64(1)
	creatorAddr := addrFromUint64(codeID)
	contractAddr := keeper.generateContractAddress(ctx, codeID)

	initMsg := InitMsg{
		Verifier:    alice,
		Beneficiary: bob,
	}
	initMsgBz, err := json.Marshal(initMsg)
	require.NoError(t, err)

	expected := types.NewContractInfo(codeID, contractAddr, creatorAddr, initMsgBz)
	keeper.SetContractInfo(ctx, contractAddr, expected)

	as, err := keeper.GetContractInfo(ctx, contractAddr)
	require.NoError(t, err)
	require.Equal(t, expected, as)

	keeper.IterateContractInfo(ctx, func(contractInfo types.ContractInfo) bool {
		require.Equal(t, expected, contractInfo)
		return false
	})
}

func TestContractStore(t *testing.T) {
	models := []types.Model{
		{
			Key:   []byte("a"),
			Value: []byte("aa"),
		},
		{
			Key:   []byte("b"),
			Value: []byte("bb"),
		},
		{
			Key:   []byte("c"),
			Value: []byte("cc"),
		},
	}

	ctx, _, keeper := CreateTestInput(t)

	_, _, contractAddr := keyPubAddr()
	keeper.SetContractStore(ctx, contractAddr, models)

	i := 0
	for iter := keeper.GetContractStoreIterator(ctx, contractAddr); iter.Valid(); iter.Next() {
		require.Equal(t, models[i], types.Model{
			Key:   iter.Key(),
			Value: iter.Value(),
		})

		i++
	}
}
