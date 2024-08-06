package urchain

type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

type Utxo struct {
	Address     string `json:"address"`
	TxId        string `json:"txId"`
	OutputIndex int    `json:"outputIndex"`
	Height      int    `json:"height"`
	ScriptHash  string `json:"scriptHash"`
	Script      string `json:"script"`
	Satoshis    int    `json:"satoshis"`
	Type        string `json:"type"`
	Time        int    `json:"time"`
}

type TxHistory struct {
	TxId     string `json:"txId"`
	Satoshis int    `json:"satoshis"`
	Income   bool   `json:"income"`
	Height   int    `json:"height"`
	Time     int    `json:"time"`
}

type NoteTxHistory struct {
	TxId        string `json:"txId"`
	Amount      string `json:"amount"`
	Income      bool   `json:"income"`
	Tick        string `json:"tick"`
	Op          string `json:"op"`
	Dec         int    `json:"dec"`
	Height      int    `json:"height"`
	Time        int    `json:"time"`
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
}

type NoteTxHistoryResp struct {
	Total int              `json:"total"`
	Data  []*NoteTxHistory `json:"data"`
}

type BroadcastResult struct {
	Success bool   `json:"success"`
	TxId    string `json:"txId,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type BlockHeader struct {
	Hash   string `json:"hash"`
	Height int    `json:"height"`
}

type TokenBalance struct {
	Tick        string `json:"tick"`
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

type Token struct {
	*TokenBalance
	ScriptHash string `json:"scriptHash"`
	Dec        int    `json:"dec"`
	P          string `json:"p"`
}

type N20TokenInfo struct {
	Inpoint       string `json:"inpoint"`
	Blockchain    string `json:"blockchain"`
	TxId          string `json:"txId"`
	InputIndex    int    `json:"inputIndex"`
	Height        int    `json:"height"`
	BlockHash     string `json:"blockHash"`
	BlockTime     int    `json:"blockTime"`
	IndexInBlock  int    `json:"indexInBlock"`
	IndexInChain  int    `json:"indexInChain"`
	AccountId     string `json:"accountId"`
	Address       string `json:"address"`
	P             string `json:"p"`
	Tick          string `json:"tick"`
	Max           string `json:"max"`
	Lim           string `json:"lim"`
	Dec           int    `json:"dec"`
	Uri           string `json:"uri"`
	Idx           string `json:"idx"`
	Code          string `json:"code"`
	Sch           string `json:"sch"`
	Time          int    `json:"time"`
	DeployOptions string `json:"deployOptions"`
}

type N20TokenInfoResp struct {
	Total int             `json:"total"`
	Data  []*N20TokenInfo `json:"data"`
}

type NoteContract struct {
	Hash         string `json:"hash"`
	Inpoint      string `json:"inpoint"`
	Blockchain   string `json:"blockchain"`
	TxId         string `json:"txId"`
	InputIndex   int    `json:"inputIndex"`
	Height       int    `json:"height"`
	BlockHash    string `json:"blockHash"`
	BlockTime    int    `json:"blockTime"`
	IndexInBlock int    `json:"indexInBlock"`
	IndexInChain int    `json:"indexInChain"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	AccountId    string `json:"accountId"`
	Time         int    `json:"time"`
	Status       string `json:"status"`
	AbiJson      any    `json:"abiJson"`
	Creater      string `json:"creater"`
}

type Health struct {
	Status              string `json:"status"`
	LastHealthChecktime int    `json:"lastHealthChecktime"`
	Height              int    `json:"height"`
	Merkleroot          string `json:"merkleroot"`
}

type TokenUtxo struct {
	TxId        string `json:"txId"`
	Outpoint    string `json:"outpoint"`
	OutputIndex int    `json:"outputIndex"`
	Script      string `json:"script"`
	Satoshis    int    `json:"satoshis"`
	Holder      string `json:"holder"`
	Amount      string `json:"amount"`
	Height      int    `json:"height"`
	Tick        string `json:"tick"`
	Dec         int    `json:"dec"`
	P           string `json:"p"`
	Time        int    `json:"time"`
}

type Block struct {
	Hash              string `json:"hash"`
	Height            int    `json:"height"`
	Version           int    `json:"version"`
	VersionHex        string `json:"versionHex"`
	Time              int    `json:"time"`
	Mediantime        int    `json:"mediantime"`
	Nonce             string `json:"nonce"`
	Bits              string `json:"bits"`
	Difficulty        string `json:"difficulty"`
	Chainwork         string `json:"chainwork"`
	NumTx             int    `json:"num_tx"`
	PreviousBlockHash string `json:"previousblockhash"`
	NextBlockHash     string `json:"nextblockhash"`
	BlockSize         int    `json:"blockSize"`
	Coinbase          string `json:"coinbase"`
	BlockSubsidy      string `json:"blockSubsidy"`
	Merkleroot        string `json:"merkleroot"`
	TotalFee          string `json:"totalFee"`
	Miner             string `json:"miner"`
	Reorg             any    `json:"reorg"`
	Blockchain        string `json:"blockchain"`
	NoteMerkelRoot    string `json:"noteMerkelRoot"`
}

type NoteTx struct {
	TxId         string `json:"txId"`
	Blockchain   string `json:"blockchain"`
	Time         int    `json:"time"`
	NumInputs    int    `json:"num_inputs"`
	NumOutputs   int    `json:"num_outputs"`
	Height       int    `json:"height"`
	BlockHash    string `json:"blockHash"`
	BlockTime    int    `json:"blockTime"`
	IndexInBlock int    `json:"indexInBlock"`
	IndexInChain int    `json:"indexInChain"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	Payload      string `json:"payload"`
}

type NoteBlock struct {
	Block   *Block    `json:"block"`
	Txs     []*NoteTx `json:"txs"`
	TxCount int       `json:"txCount"`
}

type NotePayload struct {
	P    string `json:"p"`
	Op   string `json:"op"`
	Amt  any    `json:"amt"` // string or []string
	Tick string `json:"tick"`
}

type utxoInfo struct {
	TxId         string       `json:"txId"`
	Outpoint     string       `json:"outpoint"`
	OutputIndex  int          `json:"outputIndex"`
	Script       string       `json:"script"`
	Height       int          `json:"height"`
	Tick         string       `json:"tick,omitempty"`
	P            string       `json:"p,omitempty"`
	Dec          int          `json:"dec,omitempty"`
	Time         int          `json:"time"`
	BlockHash    string       `json:"blockHash"`
	BlockTime    int          `json:"blockTime"`
	IndexInBlock int          `json:"indexInBlock"`
	IndexInChain int          `json:"indexInChain"`
	Blockchain   string       `json:"blockchain,omitempty"`
	Sender       string       `json:"sender,omitempty"`
	Op           string       `json:"op,omitempty"`
	NotePayload  *NotePayload `json:"notePayload,omitempty"`
}

type TxUnlock struct {
	utxoInfo
	Inpoint        string   `json:"inpoint"`
	InputIndex     int      `json:"inputIndex"`
	PrevTxId       string   `json:"prevTxId"`
	SequenceNumber int      `json:"sequenceNumber"`
	Witnesses      []string `json:"witnesses"`
	Type           string   `json:"type"`
	FullScript     string   `json:"fullScript"`
	Amounts        []string `json:"amounts"`
}

type Txo struct {
	utxoInfo
	Satoshis   string `json:"satoshis"`
	Holder     string `json:"holder,omitempty"`
	Size       int    `json:"size"`
	ScriptHash string `json:"scriptHash"`
	Type       string `json:"type"`
	Pubkey     string `json:"pubkey"`
	Address    string `json:"address"`
	Amount     string `json:"amount,omitempty"`
}

type NoteTxInfo struct {
	TxId           string `json:"txId"`
	Blockchain     string `json:"blockchain"`
	Time           int    `json:"time"`
	NumInputs      int    `json:"num_inputs"`
	NumOutputs     int    `json:"num_outputs"`
	Height         int    `json:"height"`
	BlockHash      string `json:"blockHash"`
	BlockTime      int    `json:"blockTime"`
	IndexInBlock   int    `json:"indexInBlock"`
	IndexInChain   int    `json:"indexInChain"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	Payload        string `json:"payload"`
	Version        int    `json:"version"`
	NLockTime      int    `json:"nLockTime"`
	TxSize         int    `json:"txSize"`
	Fee            any    `json:"fee"`
	OutputSatoshis string `json:"outputSatoshis"`
	TxHex          string `json:"txHex"`
	MempoolTime    int    `json:"mempoolTime"`
	Rbf            bool   `json:"rbf"`
	Replaced       any    `json:"replaced"`

	Inputs []*struct {
		Unlock *TxUnlock `json:"Unlock"`
		TXO    *Txo      `json:"TXO"`
	} `json:"inputs"`

	Outputs []*struct {
		TXO    *Txo      `json:"TXO"`
		Unlock *TxUnlock `json:"Unlock"`

		utxoInfo
		Satoshis string `json:"satoshis,omitempty"`
		Holder   string `json:"holder,omitempty"`
		Amount   string `json:"amount,omitempty"`
	}
}
