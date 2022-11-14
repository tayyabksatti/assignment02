package assignment02

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	//"math/rand"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {
	// trans:=new(Transaction)
	//dataS := " "
	// for i :=0;i<len(Transaction);i++{
	// 	dataS.append(Transaction[i])
	// }
	var a int = rand.Intn(54222)
	fmt.Print(a)
	// blockhash := sha256.Sum256([]byte(dataS))
	// blockhash = hex.EncodeToString(blockhash[:])
	return a

}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	// trans:=new(Transaction)
	// dataS :=" "
	// for i :=0;i<len(Transaction);i++{
	// 	dataS.append(Transaction[i])
	// }
	// blockhash := sha256.Sum256([]byte(dataS))
	// blockhash = hex.EncodeToString(blockhash[:])
	// newoj := new(Block)
	// var a:= rand.Intn(54321)
	// //dataS.append(strconv(a))
	// newoj.Nonce=a
	// newoj.blockData=blockData
	// if chainHead == nil
	// {
	// 	newoj.PrevHash=NULL
	// }
	// else
	// {
	// 	newoj.PrevHash=
	// }
	// newoj.CurrentHash=blockhash

	var x *Block
	x = new(Block)
	var a int
	if chainHead != nil {
		a = GenerateNonce(chainHead.BlockData)
	} else {
		a = 0
	}

	x.Nonce = a
	x.PrevPointer = chainHead
	x.BlockData = append(x.BlockData, blockData...)
	x.BlockData = append(x.BlockData, blockData...)
	if chainHead != nil {
		x.PrevHash = CalculateHash(chainHead.BlockData, chainHead.Nonce)
	} else {
		x.PrevHash = "-"
	}
	x.CurrentHash = CalculateHash(x.BlockData, a)

	return x

}

func ListBlocks(chainHead *Block) {

	head := chainHead
	fmt.Println("Printing Blockchain, Starting from most recent transactions:\n")
	for head != nil {
		for _, v := range head.BlockData {
			fmt.Printf("Sender: %s Receiver: %s Amount: %d\n", v.Sender, v.Receiver, v.Amount)
		}
		if head.PrevPointer != nil {
			fmt.Println("|")
			fmt.Println("v")
		}

		head = head.PrevPointer
	}
	fmt.Println()
}

func (ChainObj *Block) DisplayTransactions(blockData []Transaction) {

	for i := range blockData {

		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 20), i+1, strings.Repeat("=", 20))

		fmt.Println("Recievr:", ChainObj.BlockData[i].Receiver)
		fmt.Println("Amount:", ChainObj.BlockData[i].Amount)
		fmt.Println("sender:", ChainObj.BlockData[i].Sender)
		fmt.Print("\n\n")

		fmt.Printf("%s Transactions %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))

		val, err := json.MarshalIndent(ChainObj.BlockData[i].TransactionID, "", "    ")
		if err != nil {
			fmt.Println("Error Occured.")
			panic(err)
		}

		fmt.Printf("%s\n\n", val)
	}

	fmt.Printf("%s END %s\n", strings.Repeat("=", 20), strings.Repeat("=", 20))
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	trans := new(Transaction)

	sum := sha256.Sum256([]byte(sender + receiver + strconv.Itoa((amount))))
	trans.TransactionID = hex.EncodeToString(sum[:])
	trans.Sender = sender
	trans.Receiver = receiver
	trans.Amount = amount
	return Transaction{TransactionID: "1", Sender: sender, Receiver: receiver, Amount: amount}
	//return trans
}
