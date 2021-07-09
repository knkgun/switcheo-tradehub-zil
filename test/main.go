package main

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/contract"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
)

func main() {
	//privateKey := "e887faa2e702daa055e59ff9f94d2af9ded1b308fc30935bbc1b63dabbfb8b11"
	//deployer := &Deployer{
	//	PrivateKey:    privateKey,
	//	Host:          "https://polynetworkcc3dcb2-5-api.dev.z7a.xyz",
	//	ProxyPath:     "../contracts/ZilCrossChainManagerProxy.scilla",
	//	ImplPath:      "../contracts/ZilCrossChainManager.scilla",
	//	LockProxyPath: "../contracts/LockProxy.scilla",
	//}
	//wallet := account.NewWallet()
	//wallet.AddByPrivateKey(deployer.PrivateKey)
	//client := provider.NewProvider(deployer.Host)
	//proxy, impl, lockProxy, err := deployer.Deploy(wallet, client)
	//log.Printf("lock proxy address: %s\n", lockProxy)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//p := &polynetwork.Proxy{
	//	ProxyAddr:  proxy,
	//	ImplAddr:   impl,
	//	Wallet:     wallet,
	//	Client:     client,
	//	ChainId:    chainID,
	//	MsgVersion: msgVersion,
	//}
	//
	//_, err1 := p.UpgradeTo()
	//if err1 != nil {
	//	log.Fatalln(err1.Error())
	//}
	//
	//_, err2 := p.Unpause()
	//if err2 != nil {
	//	log.Fatalln(err2.Error())
	//}
	//
	//l := &polynetwork.LockProxy{
	//	Addr:       lockProxy,
	//	Wallet:     wallet,
	//	Client:     client,
	//	ChainId:    chainID,
	//	MsgVersion: msgVersion,
	//}
	//
	//tester := &Tester{p: p, l: l}
	//tester.InitGenesisBlock()
	////tester.ChangeBookKeeper()
	//tester.VerifierHeaderAndExecuteTx()
	//
	//// dummy ethereum contract address here
	//ethLockProxy := "0x05f4a42e251f2d52b8ed15e9fedaacfcef1fad27"
	//_, err3 := l.BindProxyHash("1", ethLockProxy)
	//if err3 != nil {
	//	log.Fatalln(err3.Error())
	//}
	//
	//_, err4 := l.BindAssetHash("0x0000000000000000000000000000000000000000", "1", ethLockProxy)
	//if err4 != nil {
	//	log.Fatalln(err4.Error())
	//}
	//
	//_, err5 := l.Lock("0x0000000000000000000000000000000000000000", "1", "0xd3573e0daa110b5498c54e93b66681fc0e0ff911", "100")
	//if err5 != nil {
	//	log.Fatalln(err5.Error())
	//}
	//
	//pubKey := keytools.GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), true)
	//address := keytools.GetAddressFromPublic(pubKey)
	//
	//_, err7 := l.SetManager("0x" + address)
	//if err7 != nil {
	//	log.Fatalln(err7.Error())
	//}
	//
	//// toAssetHash 0x05f4a42e251f2d52b8ed15e9fedaacfcef1fad27
	//// toAddressHash 0xd3573e0daa110b5498c54e93b66681fc0e0ff911
	//// amount 100
	//// txData 0x1405f4a42e251f2d52b8ed15e9fedaacfcef1fad2714d3573e0daa110b5498c54e93b66681fc0e0ff9110000000000000000000000000000000000000000000000000000000000000064
	//_, err6 := l.Unlock("0x1405f4a42e251f2d52b8ed15e9fedaacfcef1fad2714d3573e0daa110b5498c54e93b66681fc0e0ff9110000000000000000000000000000000000000000000000000000000000000064", "0x05f4a42e251f2d52b8ed15e9fedaacfcef1fad27", "1")
	//if err6 != nil {
	//	log.Fatalln(err6.Error())
	//}

	bug()
}

func bug() {
	bech32Addr := ""
	prikey := ""
	url := "https://dev-api.zilliqa.com"

	wallet := account.NewWallet()
	wallet.AddByPrivateKey(prikey)

	prov := provider.NewProvider(url)

	var transactions []*transaction.Transaction

	args := []core.ContractValue{
		{
			"tokenAddr",
			"ByStr20",
			"0x0000000000000000000000000000000000000000",
		},
		{
			"targetProxyHash",
			"ByStr",
			"0x0f71dda2b923e66d6c771f804ba64f4442e8d0b8",
		},
		{
			"toAddress",
			"ByStr",
			"0x5d775a8a0b4dff032cbc6a2514c139e3cee06998",
		},
		{
			"toAssetHash",
			"ByStr",
			"0x6574682e7a696c6c697161",
		},
		{
			"feeAddr",
			"ByStr",
			"0x989761fb0c0eb0c05605e849cae77d239f98ac7f",
		},
		{
			"amount",
			"Uint256",
			"10000",
		},
		{
			"feeAmount",
			"Uint256",
			"0",
		},
	}

	data := contract.Data{
		Tag: "lock",
		Params: args,
	}

	txn := &transaction.Transaction{
		Version:      strconv.FormatInt(int64(util.Pack(333, 1)), 10),
		SenderPubKey: util.EncodeHex(wallet.DefaultAccount.PublicKey),
		ToAddr:       bech32Addr,
		Amount:       "10000",
		GasPrice:     "2000000000",
		GasLimit:     "40000",
		Code:         "",
		Data:         data,
		Priority:     false,
	}
	transactions = append(transactions, txn)

	args = []core.ContractValue{
		{
			"tokenAddr",
			"ByStr20",
			"0xdce1262e3f9b987ec7e7008da8f1af837f7db2ed",
		},
		{
			"targetProxyHash",
			"ByStr",
			"0x0f71dda2b923e66d6c771f804ba64f4442e8d0b8",
		},
		{
			"toAddress",
			"ByStr",
			"0x5d775a8a0b4dff032cbc6a2514c139e3cee06998",
		},
		{
			"toAssetHash",
			"ByStr",
			"0x6574682e7a696c6c697161",
		},
		{
			"feeAddr",
			"ByStr",
			"0x989761fb0c0eb0c05605e849cae77d239f98ac7f",
		},
		{
			"amount",
			"Uint256",
			"20000000000000000",
		},
		{
			"feeAmount",
			"Uint256",
			"0",
		},
	}

	data = contract.Data{
		Tag: "lock",
		Params: args,
	}

	txn = &transaction.Transaction{
		Version:      strconv.FormatInt(int64(util.Pack(333, 1)), 10),
		SenderPubKey: util.EncodeHex(wallet.DefaultAccount.PublicKey),
		ToAddr:       bech32Addr,
		Amount:       "0",
		GasPrice:     "2000000000",
		GasLimit:     "40000",
		Code:         "",
		Data:         data,
		Priority:     false,
	}
	transactions = append(transactions, txn)

	wallet.SignBatch(transactions,*prov)
	rs := wallet.SendBatch(transactions,*prov)
	fmt.Println(rs)

}
