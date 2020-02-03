// Copyright (c) 2018 Clearmatics Technologies Ltd

package utils_test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/clearmatics/ion-cli/utils"
	"github.com/stretchr/testify/assert"
)

// var EXPECTEDINTERFACE = "[6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177 1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347 0000000000000000000000000000000000000000 53580584816f617295ea26c0e17641e0120cab2f0a8ffb53a866fd53aa8e8c2d 56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421 56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421 00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 02 01 47c94c 00 58ee45da d783010600846765746887676f312e372e33856c696e757800000000000000009f1efa1efa72af138c915966c639544a0255e6288e188c22ce9168c10dbe46da3d88b4aa065930119fb886210bf01a084fde5d3bc48d8aa38bca92e4fcc5215100 0000000000000000000000000000000000000000000000000000000000000000 0000000000000000]"

var EXPECTEDRLPBLOCK = "f90256a06341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940000000000000000000000000000000000000000a053580584816f617295ea26c0e17641e0120cab2f0a8ffb53a866fd53aa8e8c2da056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002018347c94c008458ee45dab861d783010600846765746887676f312e372e33856c696e757800000000000000009f1efa1efa72af138c915966c639544a0255e6288e188c22ce9168c10dbe46da3d88b4aa065930119fb886210bf01a084fde5d3bc48d8aa38bca92e4fcc5215100a00000000000000000000000000000000000000000000000000000000000000000880000000000000000"

var TESTBLOCK = `{"parentHash": "0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177", "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347", "miner": "0x0000000000000000000000000000000000000000","stateRoot": "0x53580584816f617295ea26c0e17641e0120cab2f0a8ffb53a866fd53aa8e8c2d","transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","receiptsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", "difficulty": "0x2", "number": "0x1","gasLimit": "0x47c94c",	"gasUsed": "0x0",	"timestamp": "0x58ee45da",	"extraData": "0xd783010600846765746887676f312e372e33856c696e757800000000000000009f1efa1efa72af138c915966c639544a0255e6288e188c22ce9168c10dbe46da3d88b4aa065930119fb886210bf01a084fde5d3bc48d8aa38bca92e4fcc5215100",	"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000","nonce": "0x0000000000000000"}`

// EncodePrefix calculate prefix of the entire signed block
func Test_EncodePrefix(t *testing.T) {
	var blockHeader utils.Header
	err := json.Unmarshal([]byte(TESTBLOCK), &blockHeader)
	assert.Nil(t, err)

	prefix := utils.EncodePrefix(blockHeader)

	assert.Equal(t, "\x02\x14", string(prefix))
}

// EncodeExtraData calculate prefix of the extraData field without the signatures
func Test_EncodeExtraData(t *testing.T) {
	var blockHeader utils.Header
	err := json.Unmarshal([]byte(TESTBLOCK), &blockHeader)
	assert.Nil(t, err)

	prefix := utils.EncodeExtraData(blockHeader)

	assert.Equal(t, "\xa0", string(prefix))
}

// EncodeBlock rlp encodes the martialled JSON struct
func Test_EncodeBlock(t *testing.T) {
	var blockHeader utils.Header
	err := json.Unmarshal([]byte(TESTBLOCK), &blockHeader)
	assert.Nil(t, err)

	blockInterface := utils.GenerateInterface(blockHeader)

	rlpBlock := utils.EncodeBlock(blockInterface)

	strRlpBlock := hex.EncodeToString(rlpBlock)
	assert.Equal(t, EXPECTEDRLPBLOCK, strRlpBlock)
}
