package main

import (

	"encoding/json"
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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


var endpoint = "https://rinkeby.infura.io/fYe8qCnWi6TXZAXOVof9"
var gasL uint64 = 3000000
var gasP int64 = 90000000000

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
		log.Println("Error setKey", err)
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
		fmt.Println("error:", err)
	}
	if configuration.PrivateKey == "" {
		//Should create a new error here
		return "",  nil
		log.Fatal("No key in configuration file")
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
		fmt.Println("error:", err)
		return "", err
	}
	return configuration.ContractAddress, nil
}



func showConf(c *cli.Context){
	contract, err := getContractAddress()
	key, err := getKey()
	if err != nil{
		log.Fatal("Something went wrong")
	}
	fmt.Println("current key:", key)
	fmt.Println("current address:", contract)

}

//Deploys a new impro contract and writes it's address on the config file
func deploy(c *cli.Context){
	keyString, err :=  getKey()
	if err != nil {
		log.Println("Error deploy, getKey", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Error impossible to decode key to valid ECDSA key", err)
	}
	auth := bind.NewKeyedTransactor(key)
	blockchain, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	address, _, _, err := DeployImpro(auth, blockchain)
	if err != nil {
		log.Println("Unable to deploy")
	}
	fmt.Println("The contract was deployed at :", address.Hex())
	newConfig := Configuration{}
	newConfig.PrivateKey = keyString
	newConfig.ContractAddress = address.Hex()
	file, err := os.OpenFile("conf.json", os.O_RDWR | os.O_CREATE, os.ModeAppend)
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

func returnImproContract()(*Impro, error){
	contractAddress, err := getContractAddress()
	if err != nil {
		log.Println("Error returnContractObject, getContractAddress", err)
	}
	address := common.HexToAddress(contractAddress)
	blockchain, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	contract, err := NewImpro(address, blockchain)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func registerImage(c *cli.Context){
	if c.NArg() < 2 {
		log.Println("Please provide a valid filename and price for the image")
	}
	priceString := c.Args().Get(0)
	price, err := strconv.Atoi(priceString)
	if int(price) == 0 {
		log.Println("You entered a price of 0. Anyone can get the rights to your image. Switching to non buyable image")
		price = -1
	}

	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error registering image", err)
	}
	keyString, err := getKey()
	if err != nil {
		log.Println("Error registering image", err)
	}
	key, err := crypto.HexToECDSA(keyString)
	if err != nil {
		log.Println("Error impossible to decode key to valid ECDSA key", err)
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

	tx, err := contract.Register(&options, hash, big.NewInt(int64(price)))
	if err != nil {
		log.Fatal("Error with the register transaction : ", err)
	}
	fmt.Println("Transaction was successful. Tx hash : ", tx.Hash())
}


func getTimestamp(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error registering image", err)
	}
	timestamp, err := contract.GetTimestamp(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting timestamp")
	}
	fmt.Println(time.Unix(timestamp.Int64(), 0))
	fmt.Println(timestamp)
}


func getPrice(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error registering image", err)
	}
	price, err := contract.GetPrice(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting timestamp")
	}
	fmt.Println(price)
}


func getOwner(c *cli.Context){
	if c.NArg() < 1 {
		log.Println("Please provide a valid hash")
	}
	contract, err := returnImproContract()
	if err != nil {
		log.Println("Error registering image", err)
	}
	owner, err := contract.GetOwner(&bind.CallOpts{}, c.Args().First())
	if err != nil {
		log.Println("Error getting timestamp")
	}
	fmt.Println(owner.Hex())
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



func main() {

	app := cli.NewApp()
	app.Name = "impro CLI"
	//app.Usage = "interact with the impro smart contract"
	app.Version = "1"

	app.Commands = []cli.Command{
		{
			Name: "key",
			Usage: "set the key for the current session",
			ArgsUsage: "ethereum private key",
			Action: setKey,
		},
		{
			Name: "show",
			Usage: "Displays configuration file",
			Action: showConf,
		},
		{
			Name: "deploy",
			Usage: "deploys a new impro contract",
			Action: deploy,
		},
		{
			Name: "register",
			Usage: "registers an image given a filename and a price",
			Action: registerImage,
		},
		{
			Name: "price",
			Usage: "gets price of an image",
			Action: getPrice,
		},
		{
			Name: "owner",
			Usage: "gets owner of an image",
			Action: getOwner,
		},
		{
			Name: "timestamp",
			Usage: "gets timestamp of an image",
			Action: getTimestamp,
		},

	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
