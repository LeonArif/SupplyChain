package kripto

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"supplychain/backend/models"
	"fmt"
)

func CalculateHash(data string) string {
	dataByte := []byte(data)
	hashData := sha256.Sum256(dataByte)
	return hex.EncodeToString(hashData[:])
}

func HashBlock(block models.Block) string {
	payload := fmt.Sprintf("%d|%s|%s|%s", block.Index, block.Timestamp, mustJSON(block.Data), block.PrevHash)
	return CalculateHash(payload)
}

// HashFood menghasilkan fingerprint produk
func HashFood(food models.Food) string {

	expiryStr := food.ExpiryDate.Format(time.RFC3339)

	return CalculateHash(fmt.Sprintf("%s|%s|%s",food.ID, food.Name, expiryStr))
}

func mustJSON(value any) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
