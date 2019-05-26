package main

import (
	"encoding/json"
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"net/http"
	"time"
	"math/big"
	"strconv"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
	"image"
	"image/jpeg"
	"log"
	"os"
)


type Configuration struct {
	PrivateKey string
	ContractAddress string
}

type coinsData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price_usd"`
}


var endpoint = "https://rinkeby.infura.io/fYe8qCnWi6TXZAXOVof9"
var gasL uint64 = 300000
var gasP int64 = 9000000000

func setKey(c *cli.Context){
	if c.NArg() < 1{
		log.Fatal("Please provide a key")
	}
	provided := c.Args().First()
	_, err := crypto.HexToECDSA(provided)
	if err != nil {
		log.Fatal("Invalid private key. Please provide a valid key")
	}
	contract, err := getContractAddress()
	if err != nil {
		log.Println("Error getting contract address", err)
	}
	newConfig := Configuration{}
	newConfig.PrivateKey = c.Args().First()
	newConfig.ContractAddress = contract
	file, err := os.OpenFile("conf.json", os.O_RDWR, 0755)
	if err !=nil {
		log.Fatal("Error opening conf file", err)

	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(newConfig)
	if err !=nil{
		log.Fatal("Error encoding configuration file", err)
	}
}

func getKey() (string, error){
	file, err := os.OpenFile("conf.json", os.O_RDWR | os.O_CREATE,  0755)
	if err != nil {
		log.Fatal("No conf file found", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("error decoding config", err)
	}
	if configuration.PrivateKey == "" {
		log.Fatal("No key in configuration file")
		return "",  nil
	}
	return configuration.PrivateKey, nil
}


func getContractAddress() (string, error){
	file, err := os.OpenFile("conf.json", os.O_RDWR | os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("No conf file found", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("Error getting contract address in configuration", err)
		return "", err
	}
	return configuration.ContractAddress, nil
}

func showConf(c *cli.Context){
	contract, err := getContractAddress()
	if err != nil{
		log.Fatal("Unable to get the contract address")
	}
	key, err := getKey()
	if err != nil{
		log.Fatal("Unable to get the key in the conf file")
	}
	fmt.Println("current key:", key)
	fmt.Println("current address:", contract)

}

//Deploys a new impro contract and writes it's address on the config file
func deploy(c *cli.Context){
	keyString, err :=  getKey()
	if err != nil {
		log.Fatalln("Unable to get key for deploying", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Fatalln("Error impossible to decode key to valid ECDSA key", err)
	}
	auth := bind.NewKeyedTransactor(key)
	blockchain, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalln("Unable to connect to network", err)
	}
	address, _, _, err := DeployImpro(auth, blockchain)
	if err != nil {
		log.Fatalln("Unable to deploy")
	}
	fmt.Println("contract deployed at :", address.Hex())
	newConfig := Configuration{}
	newConfig.PrivateKey = keyString
	newConfig.ContractAddress = address.Hex()
	file, err := os.OpenFile("conf.json", os.O_RDWR | os.O_CREATE, os.ModeAppend)
	if err !=nil {
		log.Fatalln("Error opening conf file", err)

	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(newConfig)
	if err !=nil{
		log.Fatal("Error encoding configuration file", err)
	}
}

func returnImproContract()(*Impro, error){
	contractAddress, err := getContractAddress()
	if err != nil {
		log.Fatalln("Unable to get contract address", err)
	}
	address := common.HexToAddress(contractAddress)
	blockchain, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalln("Unable to connect to network ", err)
	}
	contract, err := NewImpro(address, blockchain)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func registerImage(c *cli.Context){
	if c.NArg() < 2 {
		log.Fatalln("Please provide a valid filename and price for the image")
	}
	priceString := c.Args().Get(1)
	price, err := strconv.ParseFloat(priceString, 64)
	ethPrice := getEthPrice()
	var ratio float64
	ratio  = price/ethPrice
	//convert price from dollars to weis
	sendingPrice  := int(ratio*1000000000000000000)
	if sendingPrice == 0 {
		fmt.Println("You entered a price of 0. Switching to non buyable image")
		price = -1
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Fatalln("Error getting the impro contract object", err)
	}
	keyString, err := getKey()
	if err != nil {
		log.Fatalln("Error getting key", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Fatalln("Error impossible to decode key to valid ECDSA key", err)
	}
	img, err := getImage(c.Args().First())
	if err != nil {
		log.Fatal("Invalid image name. Please provide a valid name.")
	}
	hash, err := makePerceptionHash(img)
	if err != nil {
		log.Fatal("Error in generating hash")
	}
	auth := bind.NewKeyedTransactor(key)
	options := bind.TransactOpts{
		auth.From,
		auth.Nonce,
		auth.Signer,
		big.NewInt(0),
		big.NewInt(gasP),
		gasL,
		auth.Context,
	}

	tx, err := contract.Register(&options, hash, big.NewInt(int64(sendingPrice)))
	if err != nil {
		log.Fatal("Error with the register transaction : ", err)
	}
	fmt.Println("transaction successful: ", tx.Hash().Hex())
	fmt.Println("perceptual hash uploaded:", hash)
}

func getTimestamp(c *cli.Context){
	if c.NArg() < 1 {
		log.Fatalln("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	timestamp, err := contract.GetTimestamp(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting timestamp")
	}
	fmt.Println("timestamp:", time.Unix(timestamp.Int64(), 0))
}

func getPrice(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	price, err := contract.GetPrice(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting price")
	}
	if price.Int64() == -1{
		fmt.Println("This is image can't be bought")

	} else {
		var marketPrice float64
		marketPrice = getEthPrice()*float64(price.Int64())/1e18
		fmt.Println(marketPrice, "$")
	}
}

func returnPrice(hash string)(*big.Int, error){
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
		return nil, err
	}
	price, err := contract.GetPrice(&bind.CallOpts{}, hash)
	if err != nil {
		log.Println("Error getting timestamp")
		return nil, err
	}
	return price, nil
}

func getOwner(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	owner, err := contract.GetOwner(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting timestamp")
	}
	fmt.Println(owner.Hex())
}

func exists(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	exist, err := contract.Exists(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting existence of contract")
	}
	if exist {
		fmt.Println("The hash ", c.Args().First(), " exists on the current impro contract.")
	} else {
		fmt.Println("hash was not yet uploaded")
	}
}

func buy(c *cli.Context) {
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	hash := c.Args().First()
	price, err := returnPrice(hash)
	if err != nil {
		log.Fatal("Error getting price of image")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	keyString, err := getKey()
	if err != nil {
		log.Println("Error getting key", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Error impossible to decode key to valid ECDSA key", err)
	}
	auth := bind.NewKeyedTransactor(key)
	options := bind.TransactOpts{
		auth.From,
		auth.Nonce,
		auth.Signer,
		price,
		big.NewInt(gasP),
		gasL,
		auth.Context,
	}
	tx, err := contract.Buy(&options, hash)
	if err != nil {
		log.Fatal("Error issuing buy transaction")
	}
	fmt.Println("Succesful buy transaction! Transaction receipt: ", tx.Hash().Hex())
}

func transfer(c *cli.Context) {
	if c.NArg() < 2{
		log.Println("Please provide a valid hash and address of new owner")
	}
	hash := c.Args().First()
	addressString := c.Args().Get(1)
	address := common.HexToAddress(addressString)
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract", err)
	}
	keyString, err := getKey()
	if err != nil {
		log.Println("Error registering image", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Error impossible to decode key to valid ECDSA key", err)
	}

	auth := bind.NewKeyedTransactor(key)
	options := bind.TransactOpts{
		auth.From,
		auth.Nonce,
		auth.Signer,
		big.NewInt(0),
		big.NewInt(gasP),
		gasL,
		auth.Context,
	}
	tx, err := contract.Transfer(&options, hash, address)
	if err != nil {
		log.Fatal("Error issuing buy transaction")
	}
	fmt.Println("Succesful transfer transaction! Transaction receipt: ", tx.Hash().Hex())
}

func changePrice(c *cli.Context) {
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash and new price")
	}
	hash := c.Args().First()
	priceString := c.Args().Get(1)
	price, err := strconv.ParseFloat(priceString, 64)
	ethPrice := getEthPrice()
	//convert price from dollars to weis
	var ratio float64
	ratio = price/ethPrice
	sendingPrice := int(ratio*1000000000000000000)
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error getting impro contract object", err)
	}
	keyString, err := getKey()
	if err != nil {
		log.Println("Error registering image", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Error impossible to decode key to valid ECDSA key", err)
	}

	auth := bind.NewKeyedTransactor(key)
	options := bind.TransactOpts{
		auth.From,
		auth.Nonce,
		auth.Signer,
		big.NewInt(0),
		big.NewInt(gasP),
		gasL,
		auth.Context,
	}
	tx, err := contract.ChangePrice(&options, hash, big.NewInt(int64(sendingPrice)))
	if err != nil {
		log.Fatal("Error issuing buy transaction")
	}
	fmt.Println("Succesfully changed price. Transaction receipt: ", tx.Hash().Hex())
}


func whoami(c *cli.Context){
	keyString, err := getKey()
	if err != nil {
		log.Println("Unable to get the key", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Unable to convert to hex ", err)
	}
	fmt.Println("Currently logged in with address :", crypto.PubkeyToAddress(key.PublicKey).Hex())
}
func getImage(fileName string) (image.Image, error){
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func makePerceptionHash(image image.Image)(string, error){
	hash, err := goimagehash.PerceptionHash(image)
	if err != nil {
		return "", err
	}
	return hash.ToString(), nil
}


func displayPHash(c *cli.Context){
	if c.NArg() < 1 {
		log.Fatal("Please provide a valid image name")
	}
	image, err := getImage(c.Args().First())
	if err != nil {
		log.Fatal("Error getting the image")
	}
	hash, err := goimagehash.PerceptionHash(image)
	if err != nil {
		log.Println("Error creating the phash of the image", err)
	}
	fmt.Println(hash.ToString())
}

func getEthPrice() float64{
	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/?convert=USD")
	if err != nil {
		log.Fatal("Error getting the price of ether")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var coins []coinsData
	err = json.Unmarshal(body, &coins)
	if err != nil {
		log.Fatal(err)
	}
	priceString := coins[0].Price
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil{
		log.Println("Error converting string to int", err)
	}
	return price
}



func main() {
	app := cli.NewApp()
	app.Name = "impro CLI"
	app.Usage = "interact with the impro smart contract"
	app.Version = "1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "buy",
			Usage: "Buys the right to an image",
			ArgsUsage: "perceptual hash of the image",
			Action: buy,
		},
		{
			Name: "change",
			Usage: "Changes the price of an image",
			ArgsUsage: "perceptual hash of the image and new price",
			Action: changePrice,
		},
		{
			Name: "deploy",
			Usage: "Deploys a new impro contract",
			Action: deploy,
		},
		{
			Name: "exists",
			Usage: "Verifies if the hash was already uploaded",
			ArgsUsage: "perceptual hash of the image",
			Action: exists,
		},
		{
			Name: "key",
			Usage: "Sets the key for the current session",
			ArgsUsage: "ethereum private key",
			Action: setKey,
		},
		{
			Name: "owner",
			Usage: "Gets owner of an image",
			ArgsUsage: "perceptual hash of the image",
			Action: getOwner,
		},
		{
			Name: "phash",
			Usage: "Displays the pHash of an image",
			ArgsUsage: "name of file",
			Action: displayPHash,
		},
		{
			Name: "price",
			Usage: "Gets price of an image",
			ArgsUsage: "perceptual hash of the image",
			Action: getPrice,
		},
		{
			Name: "register",
			Usage: "Registers an image. Given a filename, will create the perceptual hash and upload it to blockchain with a price.",
			Action: registerImage,
		},
		{
			Name: "show",
			Usage: "Displays configuration file",
			Action: showConf,
		},
		{
			Name: "timestamp",
			Usage: "Gets timestamp of an image",
			ArgsUsage: "perceptual hash of the image",
			Action: getTimestamp,
		},
		{
			Name: "transfer",
			Usage: "Transfers the ownership of an image",
			ArgsUsage: "perceptual hash of the image and new owner address",
			Action: transfer,
		},
		{
			Name: "whoami",
			Usage: "Displays address of current user",
			Action: whoami,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
