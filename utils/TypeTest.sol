pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

contract TypeTest {
    constructor() public {}

    function Bool(bool input) public {}
    function Bools(bool[] memory input) public {}
    function Bool2(bool[2] memory input) public {}
    function Bool4(bool[4] memory input) public {}
    function Bool2s(bool[2][] memory input) public {}
    function Bool4s(bool[4][] memory input) public {}
    function Bool44s(bool[4][4][] memory input) public {}
    function Bool222(bool[2][2][2] memory input) public {}
    function Bool2222(bool[2][2][2][2] memory input) public {}
    function Bools2(bool[][2] memory input) public {}
    function Boolss(bool[][] memory input) public {}
    function Bools2s(bool[][2][] memory input) public {}

    function Int8(int8 input) public {}
    function Int8s(int8[] memory input) public {}
    function Int8_2(int8[2] memory input) public {}
    function Int8_4(int8[4] memory input) public {}

    function Int16(int16 input) public {}
    function Int16s(int16[] memory input) public {}
    function Int16_2(int16[2] memory input) public {}
    function Int16_4(int16[4] memory input) public {}

    function Int32(int32 input) public {}
    function Int32s(int32[] memory input) public {}
    function Int32_2(int32[2] memory input) public {}
    function Int32_4(int32[4] memory input) public {}

    function Int64(int64 input) public {}
    function Int64s(int64[] memory input) public {}
    function Int64_2(int64[2] memory input) public {}
    function Int64_4(int64[4] memory input) public {}
    function Int64_22(int64[2][2] memory input) public {}

    function Uint8(uint8 input) public {}
    function Uint8s(uint8[] memory input) public {}
    function Uint8_2(uint8[2] memory input) public {}
    function Uint8_4(uint8[4] memory input) public {}

    function Uint16(uint16 input) public {}
    function Uint16s(uint16[] memory input) public {}
    function Uint16_2(uint16[2] memory input) public {}
    function Uint16_4(uint16[4] memory input) public {}

    function Uint32(uint32 input) public {}
    function Uint32s(uint32[] memory input) public {}
    function Uint32_2(uint32[2] memory input) public {}
    function Uint32_4(uint32[4] memory input) public {}

    function Uint64(uint64 input) public {}
    function Uint64s(uint64[] memory input) public {}
    function Uint64_2(uint64[2] memory input) public {}
    function Uint64_4(uint64[4] memory input) public {}

    function Int128(int128 input) public {}
    function Int128s(int128[] memory input) public {}
    function Int128_2(int128[2] memory input) public {}
    function Int128_4(int128[4] memory input) public {}

    function Uint128(uint128 input) public {}
    function Uint128s(uint128[] memory input) public {}
    function Uint128_2(uint128[2] memory input) public {}
    function Uint128_4(uint128[4] memory input) public {}

    function Int256(int256 input) public {}
    function Int256s(int256[] memory input) public {}
    function Int256_2(int256[2] memory input) public {}
    function Int256_4(int256[4] memory input) public {}

    function Uint256(uint256 input) public {}
    function Uint256s(uint256[] memory input) public {}
    function Uint256_2(uint256[2] memory input) public {}
    function Uint256_4(uint256[4] memory input) public {}
    function Uint256_22(uint256[2][2] memory input) public {}

    /*function Fixed(fixed8x8 input) public {}
    function Fixed64x18(fixed64x18 input) public {}
    function Fixed128x18(fixed128x18 input) public {}
    function Fixed256x80(fixed256x80 input) public {}

    function Ufixed(ufixed8x8 input) public {}
    function Ufixed64x18(ufixed64x18 input) public {}
    function Ufixed128x18(ufixed128x18 input) public {}
    function Ufixed256x80(ufixed256x80 input) public {}*/

    function Address(address input) public {}
    function Addresses(address[] memory input) public {}
    function Address2(address[2] memory input) public {}
    function Address4(address[4] memory input) public {}
    function Address22(address[2][2] memory input) public {}

    function String(string memory input) public {}
    function Strings(string[] memory input) public {}
    function String2(string[2] memory input) public {}
    function String4(string[4] memory input) public {}
    function String22(string[2][2] memory input) public {}

    function Bytes(bytes memory input) public {}
    function Bytes1(bytes1 input) public {}
    function Bytes2(bytes2 input) public {}
    function Bytes8(bytes8 input) public {}
    function Bytes32(bytes32 input) public {}

    function Byte(byte input) public {}
    function Byten(byte[] memory input) public {}
    function Byte2(byte[2] memory input) public {}
    function Byte8(byte[8] memory input) public {}
    function Byte64(byte[64] memory input) public {}
    function Byte128(byte[128] memory input) public {}
    function Byte128128(byte[128][128] memory input) public {}

    function Bytes1n(bytes1[] memory input) public {}
    function Bytes2n(bytes2[] memory input) public {}
    function Bytes8n(bytes8[] memory input) public {}
    function Bytes32n(bytes32[] memory input) public {}
    function Bytes32_22(bytes32[][] memory input) public {}

    function Bytes1_2(bytes1[2] memory input) public {}
    function Bytes2_2(bytes2[2] memory input) public {}
    function Bytes8_2(bytes8[2] memory input) public {}
    function Bytes32_2(bytes32[2] memory input) public {}
    function Bytes32_128(bytes32[128] memory input) public {}
    function Bytes32_128128(bytes32[128][128] memory input) public {}

    function Bytes1_4(bytes1[4] memory input) public {}
    function Bytes2_4(bytes2[4] memory input) public {}
    function Bytes8_4(bytes8[4] memory input) public {}
    function Bytes32_4(bytes32[4] memory input) public {}
    function Bytes32_4n(bytes32[4][] memory input) public {}
}
