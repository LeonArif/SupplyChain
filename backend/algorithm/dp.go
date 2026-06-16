package algorithm

import (
	"supplychain/backend/models"
	"time"
)

func DPRoute(graph models.Graph, source, destination string, food models.Food) models.RouteResult {
	start := time.Now()

	mode := "COST"
	if food.DaysLeft <= 3 {
		mode = "TIME"
	} else if food.DaysLeft <= 7 {
		mode = "HYBRID"
	}

	best := routeCandidate{Score: -1}
	dfsRoute(graph, source, destination, mode, []string{source}, map[string]bool{source: true}, 0, 0, &best)

	return models.RouteResult{
		Algorithm:  "DP",
		Path:       best.Path,
		TotalCost:  best.Cost,
		TotalTime:  best.Time,
		ExecTimeMs: time.Since(start).Seconds() * 1000,
		Mode:       mode,
	}
}

type routeCandidate struct {
	Path  []string
	Cost  float64
	Time  float64
	Score float64
}

func dfsRoute(graph models.Graph, current, destination, mode string, path []string, visited map[string]bool, totalCost, totalTime float64, best *routeCandidate) {
	if current == destination {
		score := routeScore(mode, totalCost, totalTime)
		if best.Score < 0 || score < best.Score {
			best.Score = score
			best.Cost = totalCost
			best.Time = totalTime
			best.Path = append([]string(nil), path...)
		}
		return
	}

	for _, edge := range graph[current] {
		if visited[edge.To] {
			continue
		}
		visited[edge.To] = true
		nextPath := append(path, edge.To)
		dfsRoute(graph, edge.To, destination, mode, nextPath, visited, totalCost+edge.Cost, totalTime+edge.Time, best)
		delete(visited, edge.To)
	}
}

func routeScore(mode string, cost, hours float64) float64 {
	switch mode {
	case "TIME":
		return hours
	case "HYBRID":
		return (cost / 100000.0) + (hours * 8)
	default:
		return cost
	}
}
