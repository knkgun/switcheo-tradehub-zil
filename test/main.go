package main

import (
	"log"

	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/crosschain/polynetwork"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
)

func main() {
	// 0x7d48043742a1103042d327111746531ca26be9be
	// zil104yqgd6z5ygrqsknyug3w3jnrj3xh6d72g2ltv
	privateKey := "e887faa2e702daa055e59ff9f94d2af9ded1b308fc30935bbc1b63dabbfb8b11"
	deployer := &Deployer{
		PrivateKey:    privateKey,
		Host:          "https://dev-api.zilliqa.com",
		ProxyPath:     "../contracts/ZilCrossChainManagerProxy.scilla",
		ImplPath:      "../contracts/ZilCrossChainManager.scilla",
		LockProxyPath: "../contracts/LockProxySwitcheo.scilla",
	}
	wallet := account.NewWallet()
	wallet.AddByPrivateKey(deployer.PrivateKey)
	client := provider.NewProvider(deployer.Host)
	proxy, impl, lockProxy, err := deployer.Deploy(wallet, client)
	log.Printf("cross chain manager proxy address: %s\n", proxy)
	log.Printf("cross chain manager address: %s\n", impl)
	log.Printf("lock proxy address: %s\n", lockProxy)
	if err != nil {
		log.Fatalln(err.Error())
	}

	p := &polynetwork.Proxy{
		ProxyAddr:  proxy,
		ImplAddr:   impl,
		Wallet:     wallet,
		Client:     client,
		ChainId:    chainID,
		MsgVersion: msgVersion,
	}

	_, err1 := p.UpgradeTo()
	if err1 != nil {
		log.Fatalln(err1.Error())
	}

	_, err2 := p.Unpause()
	if err2 != nil {
		log.Fatalln(err2.Error())
	}

	l := &polynetwork.LockProxy{
		Addr:       lockProxy,
		Wallet:     wallet,
		Client:     client,
		ChainId:    chainID,
		MsgVersion: msgVersion,
	}

	tester := &Tester{p: p, l: l}
	tester.InitGenesisBlock()
	//tester.ChangeBookKeeper()
	tester.VerifierHeaderAndExecuteTx()

	_, err2 = l.UnPause()
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
}
