package x16r

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestX16R(t *testing.T) {
	targetHex := "0000000691ce7d3da5423a2f3d66b883333e0796ae1c4f0fd35f742c2206a633"
	targetBytes, _ := hex.DecodeString(targetHex)
	headerHex := "00000030ff31798d59640f93d351b4e0a313ceb13d9eac81f9a44ba9c855a27e000000008fc3525a59d2a3ebe4c9e8e486ec9697e6d5b1a2435ddd920f775672c1c4cd11278b9c5c49b4071d064eddcf"
	headerBytes, _ := hex.DecodeString(headerHex)
	hash := X16R(headerBytes)
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}
	if !bytes.Equal(hash, targetBytes) {
		t.Fatalf("wrong hash result, got %s, should be %s", hex.EncodeToString(hash), targetHex)
	}
}
