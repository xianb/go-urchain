package urchain

import "testing"

var cli = NewClient(Options{
	TestNetwork: false,
	ApiKey:      "1234567890",
	Debug:       false,
})

var (
	scriptHash   = "ef7a8442926595a3423e8ce19e7536cf826e3acc451861c8dc87411edfbea850"
	tickNote     = "NOTE"
	address      = "bc1pt6wzc5w0wrpa54c9tha7ef9hqysxng28ce2dv6fhj2vqyewg7azq3mktu9"
	txid         = "8766ab77f3680921be62798be5bd4ddae2452c2bc845077b4be9ce98708da54f"
	contractHash = "50b13619d4d936d7c5c7fb7dfbe752e33b85b33774e9e2b3779f16791fb1c749"
	blockHash    = "0000000000000000000019973b2778f08ad6d21e083302ff0833d17066921ebb"
)

func TestGetBalance(t *testing.T) {
	balance, err := cli.GetBalance(scriptHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetBalances(t *testing.T) {
	scriptHashs := []string{
		scriptHash,
	}
	balance, err := cli.GetBalances(scriptHashs)
	if err != nil {
		t.Error(err)
	}
	t.Log(FormatStruct(balance))
}

func TestGetUnspentUtxos(t *testing.T) {
	scriptHashs := []string{
		scriptHash,
	}
	utxos, err := cli.GetUnspentUtxos(scriptHashs)
	if err != nil {
		t.Error(err)
	}

	t.Log(FormatStruct(utxos))
}

func TestGetTxHistory(t *testing.T) {
	history, err := cli.GetTxHistory(scriptHash)
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
	err := cli.RefreshUtxos(scriptHash)
	if err != nil {
		t.Error(err)
	}
}
