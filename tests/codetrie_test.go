package tests

import (
	"encoding/hex"
	"testing"
)

func TestVerifyMetadataProof(t *testing.T) {
	testCases := []struct {
		root  string
		proof []string
		leaf  string
		index int
		valid bool
	}{
		{
			root: "2a23ef2b7a7221eaac2ffb3842a506a981c009ca6c2fcbf20adbc595e56f1a93",
			proof: []string{
				"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
			},
			leaf:  "0100000000000000000000000000000000000000000000000000000000000000",
			index: 4,
			valid: true,
		},
		{
			root: "2a23ef2b7a7221eaac2ffb3842a506a981c009ca6c2fcbf20adbc595e56f1a93",
			proof: []string{
				"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
			},
			leaf:  "0100000000000000000000000000000000000000000000000000000000000000",
			index: 6,
			valid: false,
		},
		{
			root: "2a23ef2b7a7221eaac2ffb3842a506a981c009ca6c2fcbf20adbc595e56f1a93",
			proof: []string{
				"0100000000000000000000000000000000000000000000000000000000000000",
				"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
			},
			leaf:  "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			index: 5,
			valid: true,
		},
	}

	for _, c := range testCases {
		// Decode values from string to []byte
		root, err := hex.DecodeString(c.root)
		if err != nil {
			t.Errorf("Failed to decode root: %s\n", c.root)
		}
		proof := make([][]byte, len(c.proof))
		for i, p := range c.proof {
			b, err := hex.DecodeString(p)
			if err != nil {
				t.Errorf("Failed to decode proof element: %s\n", p)
			}
			proof[i] = b
		}
		leaf, err := hex.DecodeString(c.leaf)
		if err != nil {
			t.Errorf("Failed to decode leaf: %s\n", c.leaf)
		}

		// Verify proof
		ok, err := VerifyProof(root, proof, leaf, c.index)
		if err != nil {
			t.Errorf("Failed to verify proof: %v\n", err)
		}
		if ok != c.valid {
			t.Errorf("Incorrect proof verification: expected %v, got %v\n", c.valid, ok)
		}
	}
}

func TestVerifyCodeTrieProof(t *testing.T) {
	testCases := []struct {
		root  string
		proof []string
		leaf  string
		index int
		valid bool
	}{
		{
			root: "f1824b0084956084591ff4c91c11bcc94a40be82da280e5171932b967dd146e9",
			proof: []string{
				"35210d64853aee79d03f30cf0f29c1398706cbbcacaf05ab9524f00070aec91e",
				"f38a181470ef1eee90a29f0af0a9dba6b7e5d48af3c93c29b4f91fa11b777582",
			},
			leaf:  "0100000000000000000000000000000000000000000000000000000000000000",
			index: 7,
			valid: true,
		},
		{
			root: "f1824b0084956084591ff4c91c11bcc94a40be82da280e5171932b967dd146e9",
			proof: []string{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"0000000000000000000000000000000000000000000000000000000000000000",
				"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
				"0100000000000000000000000000000000000000000000000000000000000000",
				"f38a181470ef1eee90a29f0af0a9dba6b7e5d48af3c93c29b4f91fa11b777582",
			},
			leaf:  "6001000000000000000000000000000000000000000000000000000000000000",
			index: 49,
			valid: true,
		},
	}

	for _, c := range testCases {
		// Decode values from string to []byte
		root, err := hex.DecodeString(c.root)
		if err != nil {
			t.Errorf("Failed to decode root: %s\n", c.root)
		}
		proof := make([][]byte, len(c.proof))
		for i, p := range c.proof {
			b, err := hex.DecodeString(p)
			if err != nil {
				t.Errorf("Failed to decode proof element: %s\n", p)
			}
			proof[i] = b
		}
		leaf, err := hex.DecodeString(c.leaf)
		if err != nil {
			t.Errorf("Failed to decode leaf: %s\n", c.leaf)
		}

		// Verify proof
		ok, err := VerifyProof(root, proof, leaf, c.index)
		if err != nil {
			t.Errorf("Failed to verify proof: %v\n", err)
		}
		if ok != c.valid {
			t.Errorf("Incorrect proof verification: expected %v, got %v\n", c.valid, ok)
		}
	}
}

func TestVerifyCodeTrieMultiProof(t *testing.T) {
	testCases := []struct {
		root    string
		proof   []string
		leaves  []string
		indices []int
		valid   bool
	}{
		{
			root: "f1824b0084956084591ff4c91c11bcc94a40be82da280e5171932b967dd146e9",
			proof: []string{
				"0000000000000000000000000000000000000000000000000000000000000000",
				"0000000000000000000000000000000000000000000000000000000000000000",
				"f5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b",
				"0000000000000000000000000000000000000000000000000000000000000000",
				"0100000000000000000000000000000000000000000000000000000000000000",
				"f58f76419d9235451a8290a88ba380d852350a1843f8f26b8257a421633042b4",
			},
			leaves: []string{
				"0200000000000000000000000000000000000000000000000000000000000000",
				"6001000000000000000000000000000000000000000000000000000000000000",
			},
			indices: []int{10, 49},
			valid:   true,
		},
	}

	for _, c := range testCases {
		// Decode values from string to []byte
		root, err := hex.DecodeString(c.root)
		if err != nil {
			t.Errorf("Failed to decode root: %s\n", c.root)
		}
		proof := make([][]byte, len(c.proof))
		for i, p := range c.proof {
			b, err := hex.DecodeString(p)
			if err != nil {
				t.Errorf("Failed to decode proof element: %s\n", p)
			}
			proof[i] = b
		}
		leaves := make([][]byte, len(c.leaves))
		for i, l := range c.leaves {
			b, err := hex.DecodeString(l)
			if err != nil {
				t.Errorf("Failed to decode leaf: %s\n", l)
			}
			leaves[i] = b
		}

		// Verify proof
		ok, err := VerifyMultiproof(root, proof, leaves, c.indices)
		if err != nil {
			t.Errorf("Failed to verify proof: %v\n", err)
		}
		if ok != c.valid {
			t.Errorf("Incorrect proof verification: expected %v, got %v\n", c.valid, ok)
		}
	}
}
