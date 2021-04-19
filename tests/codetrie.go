package tests

type Metadata struct {
	Version    uint8
	CodeHash   []byte `ssz-size:"32"`
	CodeLength uint16
}

type Chunk struct {
	FIO  uint8
	Code []byte `ssz-size:"32"` // Last chunk is right-padded with zeros
}

type Chunk40 struct {
	FIO  uint8
	Code []byte `ssz-size:"40"`
}

type CodeTrieSmall struct {
	Metadata *Metadata
	Chunks   []*Chunk `ssz-max:"4"`
}

type CodeTrieBig struct {
	Metadata *Metadata
	Chunks   []*Chunk `ssz-max:"1024"`
}

type CodeTrie40 struct {
	Metadata *Metadata
	Chunks   []*Chunk40 `ssz-max:"1024"`
}
