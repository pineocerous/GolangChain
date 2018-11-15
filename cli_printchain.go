package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain(nodeID string) {
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("===================================== Block %d =====================================\n", block.Height)
		fmt.Printf("Hash:		%x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Valid:		%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Previous block:	%x\n\n", block.PrevBlockHash)
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("---------------------------------------------------------------------------------------\n\n")
		fmt.Printf("===================================================================================\n")
		if len(block.PrevBlockHash) == 0 {
			fmt.Printf("Genesis Block Data: %v\n", genesisCoinbaseData)
			fmt.Printf("===================================================================================\n")
			break
		}
	}
}
