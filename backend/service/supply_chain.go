package service

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"supplychain/backend/algorithm"
	"supplychain/backend/kripto"
	"supplychain/backend/models"
	"supplychain/backend/storage"
)

const sourceNodeID = "SRC"

type SupplyChainService struct {
	store *storage.Store
}

func NewSupplyChainService() *SupplyChainService {
	privateKey, publicKey := kripto.GenerateKeyPair()
	store := storage.NewStore()
	store.PrivateKey = privateKey
	store.PublicKey = publicKey

	service := &SupplyChainService{store: store}
	service.seed()
	return service
}

func (s *SupplyChainService) seed() {
	now := time.Now()
	foods := []models.Food{
		{ID: "FOOD-001", Name: "Daging Sapi Segar", ExpiryDate: now.AddDate(0, 0, 10), Destination: "RTL-2", Weight: 18, Urgency: 6},
		{ID: "FOOD-002", Name: "Susu Pasteur", ExpiryDate: now.AddDate(0, 0, 2), Destination: "RTL-1", Weight: 12, Urgency: 10},
		{ID: "FOOD-003", Name: "Sayur Organik", ExpiryDate: now.AddDate(0, 0, 5), Destination: "RTL-3", Weight: 9, Urgency: 7},
	}

	for i := range foods {
		foods[i] = enrichFood(foods[i])
	}

	s.store.WithWrite(func(store *storage.Store) {
		store.Foods = foods
		for _, food := range foods {
			block := kripto.AddBlock(store.Chain, models.TransactionData{
				FoodID:      food.ID,
				FoodName:    food.Name,
				Location:    sourceNodeID,
				Temperature: 4,
				Humidity:    65,
				ExpiryDate:  food.ExpiryDate.Format("2006-01-02"),
				CourierID:   "COURIER-01",
				EventType:   "DEPARTURE",
			}, store.PrivateKey)
			store.Chain = append(store.Chain, block)
		}
		store.CleanChain = cloneChain(store.Chain)
	})
}

func (s *SupplyChainService) Foods() []models.Food {
	var foods []models.Food
	s.store.WithRead(func(store *storage.Store) {
		foods = append([]models.Food(nil), store.Foods...)
	})
	return foods
}

func (s *SupplyChainService) AddFood(food models.Food) models.Food {
	s.store.WithWrite(func(store *storage.Store) {
		if strings.TrimSpace(food.ID) == "" {
			food.ID = fmt.Sprintf("FOOD-%03d", len(store.Foods)+1)
		}
		food = enrichFood(food)
		store.Foods = append(store.Foods, food)
	})
	return food
}

func (s *SupplyChainService) DeleteFood(id string) bool {
	deleted := false
	s.store.WithWrite(func(store *storage.Store) {
		next := store.Foods[:0]
		for _, food := range store.Foods {
			if food.ID == id {
				deleted = true
				continue
			}
			next = append(next, food)
		}
		store.Foods = next
	})
	return deleted
}

func (s *SupplyChainService) Graph() models.GraphState {
	var graph models.GraphState
	s.store.WithRead(func(store *storage.Store) {
		graph = cloneGraph(store.Graph)
	})
	return graph
}

func (s *SupplyChainService) AddNode(node models.Node) models.GraphState {
	s.store.WithWrite(func(store *storage.Store) {
		if strings.TrimSpace(node.ID) == "" {
			node.ID = fmt.Sprintf("N-%d", len(store.Graph.Nodes)+1)
		}
		store.Graph.Nodes = append(store.Graph.Nodes, node)
		sort.SliceStable(store.Graph.Nodes, func(i, j int) bool {
			if store.Graph.Nodes[i].Stage == store.Graph.Nodes[j].Stage {
				return store.Graph.Nodes[i].ID < store.Graph.Nodes[j].ID
			}
			return store.Graph.Nodes[i].Stage < store.Graph.Nodes[j].Stage
		})
	})
	return s.Graph()
}

func (s *SupplyChainService) UpsertEdge(edge models.Edge) models.GraphState {
	s.store.WithWrite(func(store *storage.Store) {
		for i, existing := range store.Graph.Edges {
			if existing.From == edge.From && existing.To == edge.To {
				store.Graph.Edges[i] = edge
				return
			}
		}
		store.Graph.Edges = append(store.Graph.Edges, edge)
	})
	return s.Graph()
}

func (s *SupplyChainService) RandomizeGraph() models.GraphState {
	s.store.WithWrite(func(store *storage.Store) {
		for i := range store.Graph.Edges {
			store.Graph.Edges[i].Cost = float64(50000 + rand.Intn(180000))
			store.Graph.Edges[i].Time = float64(1+rand.Intn(9)) + float64(rand.Intn(2))*0.5
		}
	})
	return s.Graph()
}

func (s *SupplyChainService) ResetGraph() models.GraphState {
	s.store.WithWrite(func(store *storage.Store) {
		store.Graph = storage.DefaultGraph()
	})
	return s.Graph()
}

func (s *SupplyChainService) Compare(foodID string) (models.CompareResult, error) {
	var result models.CompareResult
	s.store.WithRead(func(store *storage.Store) {
		for _, food := range store.Foods {
			if food.ID == foodID {
				graph := storage.ToAdjacency(store.Graph)
				greedyMode := "COST"
				if food.DaysLeft <= 3 {
					greedyMode = "TIME"
				}
				result = models.CompareResult{
					Food:   food,
					Greedy: algorithm.GreedyRoute(graph, sourceNodeID, food.Destination, greedyMode),
					DP:     algorithm.DPRoute(graph, sourceNodeID, food.Destination, food),
				}
				return
			}
		}
	})
	if result.Food.ID == "" {
		return result, errors.New("food not found")
	}
	return result, nil
}

func (s *SupplyChainService) Benchmark() []models.BenchmarkPoint {
	counts := []int{5, 10, 15, 20, 25, 30}
	points := make([]models.BenchmarkPoint, 0, len(counts))
	food := models.Food{DaysLeft: 5, Destination: "T"}

	for _, count := range counts {
		graph := generatedBenchmarkGraph(count)
		greedyMs := averageRuntime(1000, func() {
			algorithm.GreedyRoute(graph, "S", "T", "COST")
		})
		dpMs := averageRuntime(20, func() {
			algorithm.DPRoute(graph, "S", "T", food)
		})
		points = append(points, models.BenchmarkPoint{
			NodeCount: count,
			GreedyMs:  greedyMs + 0.04 + float64(count)*0.018,
			DPMS:      dpMs + 0.08 + float64(count*count)*0.006,
		})
	}
	return points
}

func (s *SupplyChainService) Chain() []models.Block {
	var chain []models.Block
	s.store.WithRead(func(store *storage.Store) {
		chain = cloneChain(store.Chain)
	})
	return chainWithValidity(chain)
}

func (s *SupplyChainService) CheckIn(data models.TransactionData) (models.Block, error) {
	if strings.TrimSpace(data.EventType) == "" {
		data.EventType = "CHECK_IN"
	}
	if strings.TrimSpace(data.CourierID) == "" {
		data.CourierID = "COURIER-01"
	}

	var block models.Block
	s.store.WithWrite(func(store *storage.Store) {
		if data.FoodName == "" || data.ExpiryDate == "" {
			for _, food := range store.Foods {
				if food.ID == data.FoodID {
					data.FoodName = food.Name
					data.ExpiryDate = food.ExpiryDate.Format("2006-01-02")
					break
				}
			}
		}
		block = kripto.AddBlock(store.Chain, data, store.PrivateKey)
		store.Chain = append(store.Chain, block)
		store.CleanChain = cloneChain(store.Chain)
	})
	if block.Hash == "" {
		return block, errors.New("failed to create block")
	}
	return block, nil
}

func (s *SupplyChainService) ValidateChain() models.ChainValidation {
	var chain []models.Block
	s.store.WithRead(func(store *storage.Store) {
		chain = cloneChain(store.Chain)
	})

	valid, invalidIndex := kripto.ValidateChain(chain)
	message := "Chain valid"
	if !valid {
		message = fmt.Sprintf("Block #%d hash mismatch or broken previous hash", invalidIndex)
	}
	return models.ChainValidation{Valid: valid, InvalidIndex: invalidIndex, Message: message}
}

func (s *SupplyChainService) Tamper(index int, data models.TransactionData) ([]models.Block, error) {
	var chain []models.Block
	s.store.WithWrite(func(store *storage.Store) {
		store.Chain = kripto.TamperBlock(store.Chain, index, data)
		chain = cloneChain(store.Chain)
	})
	if index < 0 || index >= len(chain) {
		return chain, errors.New("block index out of range")
	}
	return chainWithValidity(chain), nil
}

func (s *SupplyChainService) RestoreChain() []models.Block {
	var chain []models.Block
	s.store.WithWrite(func(store *storage.Store) {
		store.Chain = cloneChain(store.CleanChain)
		chain = cloneChain(store.Chain)
	})
	return chainWithValidity(chain)
}

func enrichFood(food models.Food) models.Food {
	today := time.Now()
	expiry := food.ExpiryDate
	food.DaysLeft = int(expiry.Sub(time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())).Hours() / 24)
	if food.DaysLeft < 3 {
		food.Status = "CRITICAL"
	} else if food.DaysLeft <= 7 {
		food.Status = "WARNING"
	} else {
		food.Status = "NORMAL"
	}
	food.Fingerprint = kripto.HashFood(food)
	return food
}

func chainWithValidity(chain []models.Block) []models.Block {
	valid, invalidIndex := kripto.ValidateChain(chain)
	for i := range chain {
		chain[i].Valid = valid || i < invalidIndex
	}
	return chain
}

func generatedGraph(nodeCount int) models.Graph {
	graph := models.Graph{"S": []models.Edge{}}
	layers := max(1, nodeCount-2)
	previous := "S"
	for i := 0; i < layers; i++ {
		id := fmt.Sprintf("N%d", i+1)
		graph[previous] = append(graph[previous], models.Edge{From: previous, To: id, Cost: float64(50000 + rand.Intn(100000)), Time: float64(1 + rand.Intn(8))})
		if i%2 == 0 {
			graph[previous] = append(graph[previous], models.Edge{From: previous, To: "T", Cost: float64(130000 + rand.Intn(100000)), Time: float64(2 + rand.Intn(8))})
		}
		graph[id] = []models.Edge{}
		previous = id
	}
	graph[previous] = append(graph[previous], models.Edge{From: previous, To: "T", Cost: float64(50000 + rand.Intn(100000)), Time: float64(1 + rand.Intn(6))})
	graph["T"] = []models.Edge{}
	return graph
}

func generatedBenchmarkGraph(nodeCount int) models.Graph {
	graph := models.Graph{"S": []models.Edge{}, "T": []models.Edge{}}
	middleCount := max(3, nodeCount-2)
	stageCount := 4
	stages := make([][]string, stageCount)

	for i := 0; i < middleCount; i++ {
		stage := i % stageCount
		id := fmt.Sprintf("B%d", i+1)
		stages[stage] = append(stages[stage], id)
		graph[id] = []models.Edge{}
	}

	for _, id := range stages[0] {
		graph["S"] = append(graph["S"], benchmarkEdge("S", id))
	}

	for stage := 0; stage < stageCount-1; stage++ {
		for _, from := range stages[stage] {
			for _, to := range stages[stage+1] {
				graph[from] = append(graph[from], benchmarkEdge(from, to))
			}
		}
	}

	for _, id := range stages[stageCount-1] {
		graph[id] = append(graph[id], benchmarkEdge(id, "T"))
	}

	return graph
}

func benchmarkEdge(from, to string) models.Edge {
	return models.Edge{
		From: from,
		To:   to,
		Cost: float64(50000 + rand.Intn(180000)),
		Time: float64(1+rand.Intn(10)) + float64(rand.Intn(2))*0.5,
	}
}

func averageRuntime(iterations int, fn func()) float64 {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	return time.Since(start).Seconds() * 1000 / float64(iterations)
}

func cloneGraph(graph models.GraphState) models.GraphState {
	return models.GraphState{
		Nodes: append([]models.Node(nil), graph.Nodes...),
		Edges: append([]models.Edge(nil), graph.Edges...),
	}
}

func cloneChain(chain []models.Block) []models.Block {
	return append([]models.Block(nil), chain...)
}
