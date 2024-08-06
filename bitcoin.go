package urchain

import (
	"encoding/json"
	"errors"
)

// GetBalance get bitcoin balance of a script hash
func (c *Client) GetBalance(scriptHash string) (*Balance, error) {
	param := map[string]string{
		"scriptHash": scriptHash,
	}
	resp, err := c.client.R().SetBody(param).Post("balance")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(Balance)
	err = json.Unmarshal(resp.Body(), data)
	return data, err
}

// GetBalances get bitcoin balances of multiple script hashes
func (c *Client) GetBalances(scriptHashs []string) (*Balance, error) {
	param := map[string]any{
		"scriptHashs": scriptHashs,
	}
	resp, err := c.client.R().SetBody(param).Post("wallet-balance")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(Balance)
	err = json.Unmarshal(resp.Body(), data)
	return data, err
}

// GetUnspentUtxos get unspent bitcoin utxos of multiple script hashes
func (c *Client) GetUnspentUtxos(scriptHashs []string) ([]*Utxo, error) {
	param := map[string]any{
		"scriptHashs": scriptHashs,
	}
	resp, err := c.client.R().SetBody(param).Post("utxos")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := make([]*Utxo, 0)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// GetTxHistory get bitcoin transaction history of a script hash
func (c *Client) GetTxHistory(scriptHash string) ([]*TxHistory, error) {
	param := map[string]any{
		"scriptHash": scriptHash,
	}
	resp, err := c.client.R().SetBody(param).Post("history")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := make([]*TxHistory, 0)
	err = json.Unmarshal(resp.Body(), &data)
	return data, err
}

// BroadcastTx broadcast a raw bitcoin transaction
func (c *Client) BroadcastTx(rawHex string) (*BroadcastResult, error) {
	param := map[string]any{
		"rawHex": rawHex,
	}
	resp, err := c.client.R().SetBody(param).Post("broadcast")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(BroadcastResult)
	err = json.Unmarshal(resp.Body(), data)
	return data, err
}

// GetBlockHeader get bitcoin block header
func (c *Client) GetBlockHeader() (*BlockHeader, error) {
	resp, err := c.client.R().Post("best-header")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(string(resp.Body()))
	}
	data := new(BlockHeader)
	err = json.Unmarshal(resp.Body(), data)
	return data, err
}

// RefreshUtxos refresh utxos in node cache pool
func (c *Client) RefreshUtxos(scriptHash string) error {
	param := map[string]any{
		"scriptHash": scriptHash,
	}
	resp, err := c.client.R().SetBody(param).Post("refresh")
	if err != nil {
		return err
	}
	if resp.IsError() {
		return errors.New(string(resp.Body()))
	}
	return nil
}
