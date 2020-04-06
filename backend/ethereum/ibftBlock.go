package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
)

type IBFTBlockHeader struct {
	Header *types.Header `json:"header"`
	ProposerSignature string  `json:"proposer_signature"`
	ValidatorSignature string `json:"validator_signature"`
}

type IstanbulExtra struct {
	Validators    []common.Address
	Seal          []byte
	CommittedSeal [][]byte
}

var (
	// IstanbulDigest represents a hash of "Istanbul practical byzantine fault tolerance"
	// to identify whether the block is from Istanbul consensus engine
	IstanbulDigest = common.HexToHash("0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365")

	IstanbulExtraVanity = 32 // Fixed number of extra-data bytes reserved for validator vanity
	IstanbulExtraSeal   = 65 // Fixed number of extra-data bytes reserved for validator seal

	// ErrInvalidIstanbulHeaderExtra is returned if the length of extra-data is less than 32 bytes
	ErrInvalidIstanbulHeaderExtra = errors.New("invalid istanbul header extra-data")
)


func (b *IBFTBlockHeader) Marshal() (header []byte, err error) {
	return json.Marshal(b)
}

// calculate and assign the rlp form of the header
func (b *IBFTBlockHeader) RlpEncode() (err error) {
	bValidator, err := EncodeCommitBlock(b.Header)
	if err != nil {
		return
	}

	b.ValidatorSignature = hex.EncodeToString(bValidator)

	bProposer, err := EncodeProposalBlock(b.Header)
	if err != nil {
		return
	}

	b.ProposerSignature = hex.EncodeToString(bProposer)
	return nil
}

func (b *IBFTBlockHeader) GetByNumber(rpcURL string, number string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, err := GetClient(rpcURL)
	//returnIfError(err)

	b.Header, _, err = eth.GetBlockByNumber(number)

	if err != nil {
		return err
	}

	return nil
}

func (b *IBFTBlockHeader) GetByHash(rpcURL string, hash string) (err error) {
	fmt.Println("Connecting to the RPC client..")

	eth, err := GetClient(rpcURL)
	//returnIfError(err)

	b.Header, _, err = eth.GetBlockByHash(hash)
	if err != nil {
		return err
	}

	return nil
}

func (b *IBFTBlockHeader) Print () {
	fmt.Println("Block hash:", b.Header.Hash().Hex())

}

// EncodeRLP serializes ist into the Ethereum RLP format.
func (ist *IstanbulExtra) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		ist.Validators,
		ist.Seal,
		ist.CommittedSeal,
	})
}

// DecodeRLP implements rlp.Decoder, and load the istanbul fields from a RLP stream.
func (ist *IstanbulExtra) DecodeRLP(s *rlp.Stream) error {
	var istanbulExtra struct {
		Validators    []common.Address
		Seal          []byte
		CommittedSeal [][]byte
	}
	if err := s.Decode(&istanbulExtra); err != nil {
		return err
	}
	ist.Validators, ist.Seal, ist.CommittedSeal = istanbulExtra.Validators, istanbulExtra.Seal, istanbulExtra.CommittedSeal
	return nil
}

// encodeProposalBlock returns the block signed by the block proposer of an IBFT chain
func EncodeProposalBlock(block *types.Header) ([]byte, error) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[IstanbulExtraVanity:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		return []byte{}, err
	}

	// remove proposal seal and commit seals
	istanbulExtra.Seal = make([]byte, 0)
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return []byte{}, err
	}
	block.Extra = append(block.Extra[:IstanbulExtraVanity], encodedIstanbulExtra[:]...)

	return rlp.EncodeToBytes(&block)
}

// encodeCommitBlock returns the block signed by the validators of an IBFT chain
func EncodeCommitBlock(block *types.Header) ([]byte, error) {
	// extract istanbul extraData from the block header
	istanbul := block.Extra[IstanbulExtraVanity:]

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(istanbul, &istanbulExtra)
	if err != nil {
		return []byte{}, err
	}

	// remove commit seals
	istanbulExtra.CommittedSeal = make([][]byte, 0)

	// Encode istanbulExtra
	encodedIstanbulExtra, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return []byte{}, err
	}
	block.Extra = append(block.Extra[:32], encodedIstanbulExtra[:]...)
	// fmt.Printf("%x\n", block.Extra)

	return rlp.EncodeToBytes(&block)
}


// ExtractIstanbulExtra extracts all values of the IstanbulExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractIstanbulExtra(h *types.Header) (*IstanbulExtra, error) {
	if len(h.Extra) < IstanbulExtraVanity {
		return nil, ErrInvalidIstanbulHeaderExtra
	}

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(h.Extra[IstanbulExtraVanity:], &istanbulExtra)
	if err != nil {
		return nil, err
	}
	return istanbulExtra, nil
}

// IstanbulFilteredHeader returns a filtered header which some information (like seal, committed seals)
// are clean to fulfill the Istanbul hash rules. It returns nil if the extra-data cannot be
// decoded/encoded by rlp.
func IstanbulFilteredHeader(h *types.Header, keepSeal bool) *types.Header {
	newHeader := types.CopyHeader(h)
	istanbulExtra, err := ExtractIstanbulExtra(newHeader)
	if err != nil {
		return nil
	}

	if !keepSeal {
		istanbulExtra.Seal = []byte{}
	}
	istanbulExtra.CommittedSeal = [][]byte{}

	payload, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return nil
	}

	newHeader.Extra = append(newHeader.Extra[:IstanbulExtraVanity], payload...)

	return newHeader
}

