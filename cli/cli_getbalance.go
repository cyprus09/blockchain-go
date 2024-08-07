package cli

import (
	"fmt"
	"log"

	"github.com/cyprus09/blockchain/blockchainstruct"
	"github.com/cyprus09/blockchain/utils"
	"github.com/cyprus09/blockchain/wallets"
)

func (cli *CLI) getBalance(address string, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := blockchainstruct.NewBlockchain(nodeID)
	UTXOSet := blockchainstruct.UTXOSet{Blockchain: bc}
	defer bc.DB.Close()

	balance := 0

	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
