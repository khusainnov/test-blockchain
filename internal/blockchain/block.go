package blockchain

type Block struct {
	// Hash - represents the hash of this block
	Hash []byte

	// Data - represents the data in this block (docs, images, etc)
	Data []byte

	// PrevHash - having previous block hash
	PrevHash []byte

	Nonce int
}

/*func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}*/

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}

	proof := NewProof(block)
	nonce, hash := proof.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//block.DeriveHash()

	return block
}

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}
