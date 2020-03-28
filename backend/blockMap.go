package backend

import (
	"encoding/json"
	"fmt"
)

// implement JSON Unmarshaller to unmarshal into the correct type implementing the Block interface
func (b Blocks) UnmarshalJSON(data []byte) error {

	var block = make(map[string]map[string]Block)

	json.Unmarshal(data, &block)
	fmt.Println(block)
	return nil
}
//	var block Block
//	fmt.Println("diocan")
//
//	err := json.Unmarshal(data, &block)
//	if err != nil {
//		return err
//	}
//	fmt.Println(block)
//	return nil
	//for _, block := range blocks {
	//	switch block.Type {
	//	case "eth":
	//		header := ethereum.EthBlockHeader{}
	//		json.Unmarshal(block.JsonHeader, &header)
	//		fmt.Println(header)
	//		//block.BlockInterface = &header
	//	}
	//}

	//result := make(BlockMap)

	//for k, v := range blocks{
	//	switch k {
	//	case "eth":
	//		s := reflect.ValueOf(v).Elem()
	//
	//		fmt.Println(s)
	//
	//		//header.Header = v["header"].(*types.Header)
	//		//fmt.Println(header)
	//		return nil
	//		//err := json.Unmarshal(v["header"], &header)
	//		//if err != nil {
	//		//	fmt.Println(err)
	//		//	return err
	//		//}
	//		//
	//		//fmt.Println(&header, header)
	//		//
	//		//result[k] = &header
	//	//case "rinkeby":
	//	//	header := &ethereum.CliqueBlockHeader{}
	//	//	err := json.Unmarshal(*v, &header)
	//	//	if err != nil {
	//	//		fmt.Println(err)
	//	//		return err
	//	//	}
	//	//	result["clique"] = header
	//	}
	//
	//}
	//
	//fmt.Println("result", result)
	//
	//// assign to interface
	//*b = result

	//return nil
//}

//func (b *BlockMap) MarshalJSON() ([]byte, error) {
//	fmt.Println("marshalling")
//	for k, v := range *b {
//		fmt.Println(k)
//		switch k {
//		case "eth":
//			var header ethereum.EthBlockHeader
//			return json.Marshal(header)
//		}
//	}
//}