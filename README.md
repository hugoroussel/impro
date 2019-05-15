#Impro
## Image securing using golang and ethereum smart-contracts

EE-552 class project

## What is impro?

Impro is a command line interface tool written in go. 
It lets you timestamp and upload a [perceptual hash](https://en.wikipedia.org/wiki/Perceptual_hashing) of an image
to the impro smart-contract.

```
➜  impro git:(master) ✗ ./impro
NAME:
   image protection CLI 

USAGE:
   impro [global options] command [command options] [arguments...]

VERSION:
   1

COMMANDS:
     key        Sets the key for the current session
     show       Displays configuration file
     deploy     Deploys a new impro contract
     register   Registers an image. Given a filename, will create the perceptual hash and upload it to blockchain with a price.
     price      Gets price of an image
     owner      Gets owner of an image
     timestamp  Gets timestamp of an image
     exists     Verifies if the hash was already uploaded
     buy        Buys the right to an image
     change     Changes the price of an image
     transfer   Transfers the ownership of an image
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## What is the impro smart contract? 

The impro smart-contract is written in solidity and runs on the Ethereum blockchain.

Currently the contract lets you register a perceptual hash and timestamps it. You can then transfer the rights to another account, change the price and sell it 
to others in a peer to peer  way. 


     
