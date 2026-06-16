package algorithm

import (
	"sort"
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

	path, totalCost, totalTime, ok := greedySearch(graph, source, destination, normalizedMode, []string{source}, map[string]bool{source: true})
	if !ok {
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

func greedySearch(graph models.Graph, current, destination, mode string, path []string, visited map[string]bool) ([]string, float64, float64, bool) {
	if current == destination {
		return append([]string(nil), path...), 0, 0, true
	}

	choices := make([]models.Edge, 0, len(graph[current]))
	for _, edge := range graph[current] {
		if visited[edge.To] {
			continue
		}
		choices = append(choices, edge)
	}

	sort.SliceStable(choices, func(i, j int) bool {
		iScore := choices[i].Cost
		jScore := choices[j].Cost
		if mode == "TIME" {
			iScore = choices[i].Time
			jScore = choices[j].Time
		}
		if iScore == jScore {
			return choices[i].To < choices[j].To
		}
		return iScore < jScore
	})

	for _, edge := range choices {
		visited[edge.To] = true
		nextPath := append(path, edge.To)
		resultPath, resultCost, resultTime, ok := greedySearch(graph, edge.To, destination, mode, nextPath, visited)
		if ok {
			return resultPath, resultCost + edge.Cost, resultTime + edge.Time, true
		}
		delete(visited, edge.To)
	}

	return nil, 0, 0, false
}
