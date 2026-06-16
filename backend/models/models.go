package models

import "time"

type Food struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ExpiryDate  time.Time `json:"expiryDate"`
	DaysLeft    int       `json:"daysLeft"`
	Destination string    `json:"destination"`
	Weight      float64   `json:"weight"`
	Urgency     int       `json:"urgency"`
	Status      string    `json:"status"`
	Fingerprint string    `json:"fingerprint"`
}

// kripto

type TransactionData struct {
	FoodID      string  `json:"foodId"`
	FoodName    string  `json:"foodName"`
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	ExpiryDate  string  `json:"expiryDate"`
	CourierID   string  `json:"courierId"`
	EventType   string  `json:"eventType"`
}

type Block struct {
	Index     int             `json:"index"`
	Timestamp string          `json:"timestamp"`
	Data      TransactionData `json:"data"`
	PrevHash  string          `json:"prevHash"`
	Hash      string          `json:"hash"`
	Signature string          `json:"signature"`
	Valid     bool            `json:"valid"`
}

// stima

type Node struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stage int    `json:"stage"`
}

type Edge struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Cost float64 `json:"cost"`
	Time float64 `json:"time"`
}

type Graph map[string][]Edge

type GraphState struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type RouteResult struct {
	Algorithm  string   `json:"algorithm"`
	Path       []string `json:"path"`
	TotalCost  float64  `json:"totalCost"`
	TotalTime  float64  `json:"totalTime"`
	ExecTimeMs float64  `json:"execTimeMs"`
	Mode       string   `json:"mode"`
}

type CompareResult struct {
	Food   Food        `json:"food"`
	Greedy RouteResult `json:"greedy"`
	DP     RouteResult `json:"dp"`
}

type BenchmarkPoint struct {
	NodeCount int     `json:"nodeCount"`
	GreedyMs  float64 `json:"greedyMs"`
	DPMS      float64 `json:"dpMs"`
}

type ChainValidation struct {
	Valid        bool   `json:"valid"`
	InvalidIndex int    `json:"invalidIndex"`
	Message      string `json:"message"`
}
