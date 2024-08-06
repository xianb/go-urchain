package urchain

import "testing"

func TestGetNoteTxHistory(t *testing.T) {
	history, err := cli.GetNoteTxHistory(scriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(history))
}

func TestGetNoteTokensByAddress(t *testing.T) {
	tokens, err := cli.GetNoteTokensByAddress(address)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tokens))
}

func TestGetNoteTokensByScriptHash(t *testing.T) {
	tokens, err := cli.GetNoteTokensByScriptHash(scriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tokens))
}

func TestGetTokenBalanceByAddress(t *testing.T) {
	balance, err := cli.GetTokenBalanceByAddress(address, tickNote)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetTokenBalanceByScriptHas(t *testing.T) {
	balance, err := cli.GetTokenBalanceByScriptHas(scriptHash, tickNote)
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
	info, err := cli.GetTokenInfo(tickNote)
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

func TestCheckHealth(t *testing.T) {
	h, err := cli.CheckUrchainHealth()
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(h))
}

func TestGetTokenUtxos(t *testing.T) {
	scriptHashs := []string{
		scriptHash,
	}
	utxos, err := cli.GetTokenUtxos(scriptHashs, tickNote)
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

func TestParseNoteTx(t *testing.T) {
	tx, err := cli.ParseNoteTx(txid)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(tx))
}
