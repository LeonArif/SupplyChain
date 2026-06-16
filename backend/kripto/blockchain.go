package kripto

import (
	"supplychain/backend/models"
	"time"
)

func AddBlock(chain []models.Block, data models.TransactionData, privateKey string) models.Block {
	index := len(chain)
	prevHash := ""
	if len(chain) > 0 {
		prevHash = chain[len(chain)-1].Hash
	}

	block := models.Block{
		Index:     index,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
		PrevHash:  prevHash,
		Valid:     true,
	}
	block.Hash = HashBlock(block)
	block.Signature = SignData(block.Hash, privateKey)
	return block
}

func ValidateChain(chain []models.Block) (bool, int) {
	for i, block := range chain {
		expectedHash := HashBlock(block)
		if block.Hash != expectedHash {
			return false, i
		}
		if i == 0 {
			if block.PrevHash != "" {
				return false, i
			}
			continue
		}
		if block.PrevHash != chain[i-1].Hash {
			return false, i
		}
	}
	return true, -1
}

func TamperBlock(chain []models.Block, index int, newData models.TransactionData) []models.Block {
	if index < 0 || index >= len(chain) {
		return chain
	}

	chain[index].Data = newData
	valid, invalidIndex := ValidateChain(chain)
	for i := range chain {
		chain[i].Valid = valid || i < invalidIndex
	}
	return chain
}
