package algorithm

import (
	"strings"
	"supplychain/backend/models"
	"time"
)

func GreedyRoute(graph models.Graph, source, destination, mode string) models.RouteResult {
	start := time.Now()

	normalizedMode := strings.ToUpper(strings.TrimSpace(mode))
	if normalizedMode != "TIME" {
		normalizedMode = "COST"
	}

	path := []string{source}
	visited := map[string]bool{source: true}
	current := source
	totalCost := 0.0
	totalTime := 0.0

	for current != destination {
		edges := graph[current]
		bestIndex := -1
		bestScore := 0.0

		for i, edge := range edges {
			if visited[edge.To] {
				continue
			}
			score := edge.Cost
			if normalizedMode == "TIME" {
				score = edge.Time
			}
			if bestIndex == -1 || score < bestScore {
				bestIndex = i
				bestScore = score
			}
		}

		if bestIndex == -1 {
			break
		}

		edge := edges[bestIndex]
		totalCost += edge.Cost
		totalTime += edge.Time
		current = edge.To
		visited[current] = true
		path = append(path, current)
	}

	if current != destination {
		path = nil
		totalCost = 0
		totalTime = 0
	}

	return models.RouteResult{
		Algorithm:  "GREEDY",
		Path:       path,
		TotalCost:  totalCost,
		TotalTime:  totalTime,
		ExecTimeMs: time.Since(start).Seconds() * 1000,
		Mode:       normalizedMode,
	}
}
