package storage

import (
	"sync"

	"supplychain/backend/models"
)

type Store struct {
	mu         sync.RWMutex
	Foods      []models.Food
	Graph      models.GraphState
	Chain      []models.Block
	CleanChain []models.Block
	PrivateKey string
	PublicKey  string
}

func NewStore() *Store {
	return &Store{
		Foods:      []models.Food{},
		Graph:      DefaultGraph(),
		Chain:      []models.Block{},
		CleanChain: []models.Block{},
	}
}

func DefaultGraph() models.GraphState {
	nodes := []models.Node{
		{ID: "SRC", Name: "Pabrik Utama", Stage: 0},
		{ID: "WH-JKT", Name: "Gudang Jakarta", Stage: 1},
		{ID: "WH-SBY", Name: "Gudang Surabaya", Stage: 1},
		{ID: "HUB-BDG", Name: "Hub Bandung", Stage: 2},
		{ID: "HUB-SMG", Name: "Hub Semarang", Stage: 2},
		{ID: "RTL-1", Name: "Retailer 1", Stage: 3},
		{ID: "RTL-2", Name: "Retailer 2", Stage: 3},
		{ID: "RTL-3", Name: "Retailer 3", Stage: 3},
	}

	edges := []models.Edge{
		{From: "SRC", To: "WH-JKT", Cost: 80000, Time: 5},
		{From: "SRC", To: "WH-SBY", Cost: 120000, Time: 7},
		{From: "WH-JKT", To: "HUB-BDG", Cost: 60000, Time: 3},
		{From: "WH-JKT", To: "HUB-SMG", Cost: 150000, Time: 4},
		{From: "WH-SBY", To: "HUB-BDG", Cost: 180000, Time: 6},
		{From: "WH-SBY", To: "HUB-SMG", Cost: 80000, Time: 3},
		{From: "HUB-BDG", To: "RTL-1", Cost: 70000, Time: 2},
		{From: "HUB-BDG", To: "RTL-2", Cost: 400000, Time: 2.5},
		{From: "HUB-SMG", To: "RTL-2", Cost: 70000, Time: 4},
		{From: "HUB-SMG", To: "RTL-3", Cost: 88000, Time: 3.5},
	}

	return models.GraphState{Nodes: nodes, Edges: edges}
}

func ToAdjacency(graph models.GraphState) models.Graph {
	adj := models.Graph{}
	for _, node := range graph.Nodes {
		adj[node.ID] = []models.Edge{}
	}
	for _, edge := range graph.Edges {
		if edge.From == "" {
			continue
		}
		adj[edge.From] = append(adj[edge.From], edge)
	}
	return adj
}

func (s *Store) WithRead(fn func(*Store)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fn(s)
}

func (s *Store) WithWrite(fn func(*Store)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fn(s)
}
