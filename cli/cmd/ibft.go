package cmd

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	//"github.com/clearmatics/ion-cli/cli"
)

type IstanbulExtra struct {
	Validators    []common.Address
	Seal          []byte
	CommittedSeal [][]byte
}

func IBFTCommands() []*ishell.Cmd {
	return []*ishell.Cmd{
		//{
		//	Name: "getBlockByNumber_IBFT",
		//	Help: "use: \tgetBlockByNumber_IBFT [optional rpc url] [integer]\n\t\t\t\tdescription: Returns proposal, commital RLP-encoded block headers and commit seals by block number required for submission to IBFT validation from connected client or specified endpoint",
		//	Func: func(c *ishell.Context) {
		//		if len(c.Args) == 1 {
		//			if ethClient != nil {
		//				block, _, err := cli.GetBlockByNumber(ethClient, c.Args[0])
		//				if err != nil {
		//					c.Println(err)
		//					return
		//				}
		//				proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//				c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//				c.Printf("Commital Block:\n %+x\n", commitBlock)
		//				c.Printf("Commit Seals:\n %+x\n", seals)
		//			} else {
		//				c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc url] \n")
		//				return
		//			}
		//		} else if len(c.Args) == 2 {
		//			client, err := getClient(c.Args[0])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			block, _, err := cli.GetBlockByNumber(client, c.Args[1])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//			c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//			c.Printf("Commital Block:\n %+x\n", commitBlock)
		//			c.Printf("Commit Seals:\n %+x\n", seals)
		//		} else {
		//			c.Println("Usage: \tgetBlockByNumber_IBFT [optional rpc url] [integer]\n")
		//			return
		//		}
		//		c.Println("===============================================================")
		//	},
		//},
		//{
		//	Name: "getBlockByHash_IBFT",
		//	Help: "use: \tgetBlockByHash_IBFT [optional rpc url] [integer]\n\t\t\t\tdescription: Returns proposal, commital RLP-encoded block headers and commit seals by block hash required for submission to IBFT validation from connected client or specified endpoint",
		//	Func: func(c *ishell.Context) {
		//		if len(c.Args) == 1 {
		//			if ethClient != nil {
		//				block, _, err := cli.GetBlockByHash(ethClient, c.Args[0])
		//				if err != nil {
		//					c.Println(err)
		//					return
		//				}
		//				proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//				c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//				c.Printf("Commital Block:\n %+x\n", commitBlock)
		//				c.Printf("Commit Seals:\n %+x\n", seals)
		//			} else {
		//				c.Println("Please connect to a Client before invoking this function.\nUse \tconnectToClient [rpc url] \n")
		//				return
		//			}
		//		} else if len(c.Args) == 2 {
		//			client, err := getClient(c.Args[0])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			block, _, err := cli.GetBlockByHash(client, c.Args[1])
		//			if err != nil {
		//				c.Println(err)
		//				return
		//			}
		//			proposalBlock, commitBlock, seals := RlpEncodeIBFT(block)
		//			c.Printf("Proposal Block:\n %+x\n", proposalBlock)
		//			c.Printf("Commital Block:\n %+x\n", commitBlock)
		//			c.Printf("Commit Seals:\n %+x\n", seals)
		//		} else {
		//			c.Println("Usage: \tgetBlockByHash_IBFT [optional rpc url] [integer]\n")
		//			return
		//		}
		//		c.Println("===============================================================")
		//	},
		//},
	}
}

// RlpEncodeIBFT returns rlp encoded block header from an IBFT consensus chain
func RlpEncodeIBFT(blockHeader *types.Header) (proposalBlock []byte, commitBlock []byte, commitSeals []byte) {

	// Generate an interface to encode the blockheader without the signature in the extraData
	commitSeals = extractSeals(blockHeader)
	commitBlock = encodeCommitBlock(blockHeader)
	proposalBlock = encodeProposalBlock(blockHeader)

	return
}

func extractSeals(block *types.Header) (commitSeals []byte) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[32:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	commitSeals, err = rlp.EncodeToBytes(&istanbulExtra.CommittedSeal)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	return
}

// encodeProposalBlock returns the block signed by the block proposer of an IBFT chain
func encodeProposalBlock(block *types.Header) (encodedBlock []byte) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[32:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// remove proposal seal and commit seals
	istanbulExtra.Seal = make([]byte, 0)
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}
	block.Extra = append(block.Extra[:32], encodedIstanbulExtra[:]...)

	encodedBlock, err = rlp.EncodeToBytes(&block)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// return extradata to original so we can

	return
}

// encodeCommitBlock returns the block signed by the block proposer of an IBFT chain
func encodeCommitBlock(block *types.Header) (encodedBlock []byte) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[32:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	// remove proposal seal and commit seals
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}
	block.Extra = append(block.Extra[:32], encodedIstanbulExtra[:]...)
	// fmt.Printf("%x\n", block.Extra)

	encodedBlock, err = rlp.EncodeToBytes(&block)
	if err != nil {
		fmt.Println("can't RLP encode requested block:", err)
		return
	}

	return
}
