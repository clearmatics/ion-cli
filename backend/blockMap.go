package backend

import (
	"encoding/json"
	"fmt"
	"github.com/clearmatics/ion-cli/backend/ethereum"
)

// implement JSON Unmarshaller to unmarshal into the correct type implementing the Block interface
func (b *BlockMap) UnmarshalJSON(data []byte) error {
	blocks := make(map[string]json.RawMessage)

	err := json.Unmarshal(data, &blocks)
	if err != nil {
		return err
	}

	result := make(BlockMap)

	for k, v := range blocks{
		switch k {
		case "eth":
			header := &ethereum.EthBlockHeader{}
			err := json.Unmarshal(v, &header)
			if err != nil {
				fmt.Println(err)
				return err
			}
			result["eth"] = header
		}
	}

	// assign to interface
	*b = result

	return nil
}
