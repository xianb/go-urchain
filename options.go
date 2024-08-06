package urchain

type Options struct {

	// bitcoin testnet or mainnet
	TestNetwork bool

	// urchain api key
	ApiKey string

	//http debug
	Debug bool
}

func (o *Options) GetHost() string {
	if o.TestNetwork {
		return "https://btc-testnet4.urchain.com/api/"
	}
	return "https://btc.urchain.com/api/"
}
