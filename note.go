package urchain

import (
	"encoding/json"
	"errors"
)

// GetNoteTxHistory get note transaction history of a script hash
func (c *Client) GetNoteTxHistory(scriptHash string) ([]*NoteTxHistory, error) {
	param := map[string]any{
		"scriptHash": scriptHash,
	}
	resp, err := c.client.R().SetBody(param).Post("note-history")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(NoteTxHistoryResp)
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) getNoteTokens(param map[string]any) ([]*Token, error) {
	if len(param) == 0 {
		return nil, errors.New("scriptHash or address is required")
	}
	resp, err := c.client.R().SetBody(param).Post("token-list")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := make([]*Token, 0)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// GetNoteTokensByAddress get note tokens of an address
func (c *Client) GetNoteTokensByAddress(address string) ([]*Token, error) {
	param := map[string]any{
		"address": address,
	}
	return c.getNoteTokens(param)
}

// GetNoteTokensByScriptHash get note tokens of a script hash
func (c *Client) GetNoteTokensByScriptHash(scriptHash string) ([]*Token, error) {
	param := map[string]any{
		"scriptHash": scriptHash,
	}
	return c.getNoteTokens(param)
}

func (c *Client) getTokenBalance(param map[string]any) (*TokenBalance, error) {
	resp, err := c.client.R().SetBody(param).Post("token-balance")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(TokenBalance)
	err = json.Unmarshal(resp.Body(), data)
	return data, err
}

// GetTokenBalanceByAddress get token balance of an address
func (c *Client) GetTokenBalanceByAddress(address, tick string) (*TokenBalance, error) {
	param := map[string]any{
		"address": address,
		"tick":    tick,
	}
	return c.getTokenBalance(param)
}

// GetTokenBalanceByScriptHas get token balance of a script hash
func (c *Client) GetTokenBalanceByScriptHas(scriptHash, tick string) (*TokenBalance, error) {
	param := map[string]any{
		"scriptHash": scriptHash,
		"tick":       tick,
	}
	return c.getTokenBalance(param)
}

// GetAllN20Tokens get all N20 tokens
func (c *Client) GetAllN20Tokens() ([]*N20TokenInfo, error) {
	resp, err := c.client.R().Post("all-n20-tokens")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(N20TokenInfoResp)
	err = json.Unmarshal(resp.Body(), &data)
	return data.Data, err
}

// GetTokenInfo get token info by tick
func (c *Client) GetTokenInfo(tick string) (*N20TokenInfo, error) {
	param := map[string]any{
		"tick": tick,
	}
	resp, err := c.client.R().SetBody(param).Post("token-info")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(N20TokenInfo)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// GetNoteContract get note contract by hash
func (c *Client) GetNoteContract(contractHash string) (*NoteContract, error) {
	param := map[string]any{
		"hash": contractHash,
	}
	resp, err := c.client.R().SetBody(param).Post("contract")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(NoteContract)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// CheckUrchainHealth check urchain health
func (c *Client) CheckUrchainHealth() (*Health, error) {
	resp, err := c.client.R().Get("health")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(Health)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// GetTokenUtxos get token utxos
func (c *Client) GetTokenUtxos(scriptHashs []string, tick string) ([]*TokenUtxo, error) {
	param := map[string]any{
		"scriptHashs": scriptHashs,
		"tick":        tick,
	}
	resp, err := c.client.R().SetBody(param).Post("token-utxos")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := make([]*TokenUtxo, 0)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// GetNoteBlock get note block by hash
func (c *Client) GetNoteBlock(blockHash string) (*NoteBlock, error) {
	param := map[string]any{
		"hash": blockHash,
	}
	resp, err := c.client.R().SetBody(param).Post("note-block")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(NoteBlock)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// ParseNoteTx parse note transaction by txId
func (c *Client) ParseNoteTx(txId string) (*NoteTxInfo, error) {
	param := map[string]any{
		"txId": txId,
	}
	resp, err := c.client.R().SetBody(param).Post("note-tx")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(NoteTxInfo)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// DecodePSBT decode psbt
func (c *Client) DecodePSBT(psbtHex, blockchain string) (string, error) {
	param := map[string]any{
		"psbtHex":    psbtHex,
		"blockchain": blockchain,
	}
	resp, err := c.client.R().SetBody(param).Post("decode-psbt")
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", errors.New(string(resp.Body()))
	}
	return string(resp.Body()), nil
}
