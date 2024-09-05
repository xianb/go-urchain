package urchain

import (
	"testing"
)

var cli = NewClient(Options{
	Testnet: false,
	ApiKey:  "1234567890",
	Debug:   true,
})

var (
	tokenAddressScriptHash   = "f6e1cbf01d9eca1cfa0b5e3a9b953b37e41538c94795b500ef4c4155902e7185"
	bitcoinAddressScriptHash = "16a108c4e6c8424e7fc463dab72d58b6ceeeada07ea6bfc9721247f820191182"
	tick                     = "NOTE"
	tokenAddress             = "bc1p7qgearx2g34jgd3uynz3fsjhppl5sj75yrhy4739lhmardlrmghskuy390"
	txid                     = "ca7ae00b0e81b7f89adf3c0e3283142cd6bc2bba9c3d521255910955a562ff00"
	contractHash             = "50b13619d4d936d7c5c7fb7dfbe752e33b85b33774e9e2b3779f16791fb1c749"
	blockHash                = "0000000000000000000052667a65a50429bfb3c4380a3e9afd1ccb3f6dd6e809"
)

func TestCheckHealth(t *testing.T) {
	h, err := cli.CheckUrchainHealth()
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(h))
}

func TestGetBalance(t *testing.T) {
	balance, err := cli.GetBalance(bitcoinAddressScriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetBalances(t *testing.T) {
	scriptHashs := []string{
		bitcoinAddressScriptHash,
	}
	balances, err := cli.GetBalances(scriptHashs)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balances))
}

func TestGetUnspentUtxos(t *testing.T) {
	scriptHashs := []string{
		bitcoinAddressScriptHash,
	}
	utxos, err := cli.GetUnspentUtxos(scriptHashs)
	if err != nil {
		t.Error(err)
	}

	t.Log(FormatStruct(utxos))
}

func TestGetTxHistory(t *testing.T) {
	history, err := cli.GetTxHistory(bitcoinAddressScriptHash)
	if err != nil {
		t.Error(err)
	}

	t.Log(FormatStruct(history))
}

func TestBroadcastTx(t *testing.T) {
	rawHex := ""
	res, err := cli.BroadcastTx(rawHex)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(res))
}

func TestGetBlockHeader(t *testing.T) {
	header, err := cli.GetBlockHeader()
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(header))
}

func TestRefreshUtxos(t *testing.T) {
	err := cli.RefreshUtxos(tokenAddressScriptHash)
	if err != nil {
		t.Error(err)
	}
}

func TestGetNoteTxHistory(t *testing.T) {
	history, err := cli.GetNoteTxHistory(tokenAddressScriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(history))
}

func TestGetNoteTokensByAddress(t *testing.T) {
	tokens, err := cli.GetTokensByAddress(tokenAddress)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tokens))
}

func TestGetNoteTokensByScriptHash(t *testing.T) {
	tokens, err := cli.GetTokensByScriptHash(tokenAddressScriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tokens))
}

func TestGetTokenBalanceByAddress(t *testing.T) {
	balance, err := cli.GetTokenBalanceByAddress(tokenAddress, tick)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetTokenBalanceByScriptHash(t *testing.T) {
	balance, err := cli.GetTokenBalanceByScriptHas(tokenAddressScriptHash, tick)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetAllN20Tokens(t *testing.T) {
	resp, err := cli.GetAllN20Tokens()
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(resp))
}

func TestGetTokenInfo(t *testing.T) {
	info, err := cli.GetTokenInfo("NOTE")
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(info))
}

func TestGetNoteContract(t *testing.T) {
	contract, err := cli.GetNoteContract(contractHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(contract))
}

func TestGetTokenUtxos(t *testing.T) {
	scriptHashs := []string{
		tokenAddressScriptHash,
	}
	utxos, err := cli.GetTokenUtxos(scriptHashs, tick)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(utxos))
}

func TestGetNoteBlock(t *testing.T) {
	block, err := cli.GetNoteBlock(blockHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(block))
}

func TestGetNoteTxInfo(t *testing.T) {
	tx, err := cli.GetNoteTxInfo(txid)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tx))
}
