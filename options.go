package urchain

type Options struct {

	// testnet4 or mainnet
	Testnet bool

	// urchain api key
	ApiKey string

	// http debug
	Debug bool
}

func (o *Options) GetHost() string {
	if o.Testnet {
		return "https://btc-testnet4.urchain.com/api/"
	}
	return "https://btc.urchain.com/api/"
}
