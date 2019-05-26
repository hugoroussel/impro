# Impro
## Image securing using golang and ethereum smart-contracts

EE-552 Media Security class project

Impro for image protection not for improvisation

## What is impro?

Impro is a command line interface tool written in go. 
It lets you timestamp and upload a [perceptual hash](https://en.wikipedia.org/wiki/Perceptual_hashing) of an image, the goal being
to be able to provide a proof of ownership for a given image. One can then transfer or buy the image rights.

Disclaimer : this is a proof of concept application and should be used as such. 


```
➜  impro git:(master) ✗ ./impro
NAME:
   impro CLI - interact with the impro smart contract
   
VERSION:
   1

COMMANDS:
     buy        Buys the right to an image
     change     Changes the price of an image
     deploy     Deploys a new impro contract
     exists     Verifies if the hash was already uploaded
     key        Sets the key for the current session
     owner      Gets owner of an image
     phash      Displays the pHash of an image
     price      Gets price of an image
     register   Registers an image. Given a filename, will create the perceptual hash and upload it to blockchain with a price.
     show       Displays configuration file
     timestamp  Gets timestamp of an image
     transfer   Transfers the ownership of an image
     whoami     Displays address of current user
     help, h    Shows a list of commands or help for one command
```


## How to build?
```
git clone https://github.com/hugoroussel/impro
cd impro
go build
./impro
```
## What is the impro smart contract? 

The impro smart-contract is written in solidity and runs on the Ethereum blockchain.

Currently the contract lets you register a perceptual hash and timestamps it. You can then transfer the rights to another account, change the price and sell it 
to others in a peer to peer  way.

## Built with

- Solidity
- Golang
- [Go Image Hash](github.com/corona10/goimagehash) 




     
