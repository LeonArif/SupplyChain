package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"supplychain/backend/models"
	"supplychain/backend/service"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	svc := service.NewSupplyChainService()

	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type")
		context.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS")
		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	})

	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "ok"})
	})

	api := router.Group("/api")
	{
		api.GET("/foods", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.Foods())
		})

		api.POST("/foods", func(context *gin.Context) {
			var req foodRequest
			if err := context.ShouldBindJSON(&req); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			expiry, err := parseDate(req.ExpiryDate)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiryDate"})
				return
			}
			food := svc.AddFood(models.Food{
				ID:          req.ID,
				Name:        req.Name,
				ExpiryDate:  expiry,
				Destination: req.Destination,
				Weight:      req.Weight,
				Urgency:     req.Urgency,
			})
			context.JSON(http.StatusCreated, food)
		})

		api.DELETE("/foods/:id", func(context *gin.Context) {
			if !svc.DeleteFood(context.Param("id")) {
				context.JSON(http.StatusNotFound, gin.H{"error": "food not found"})
				return
			}
			context.Status(http.StatusNoContent)
		})

		api.GET("/graph", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.Graph())
		})

		api.POST("/graph/node", func(context *gin.Context) {
			var node models.Node
			if err := context.ShouldBindJSON(&node); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, svc.AddNode(node))
		})

		api.POST("/graph/edge", func(context *gin.Context) {
			var edge models.Edge
			if err := context.ShouldBindJSON(&edge); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, svc.UpsertEdge(edge))
		})

		api.POST("/graph/randomize", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.RandomizeGraph())
		})

		api.POST("/graph/reset", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.ResetGraph())
		})

		api.POST("/algo/compare", func(context *gin.Context) {
			var req compareRequest
			if err := context.ShouldBindJSON(&req); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result, err := svc.Compare(req.FoodID)
			if err != nil {
				context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, result)
		})

		api.POST("/algo/benchmark", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.Benchmark())
		})

		api.GET("/chain", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.Chain())
		})

		api.POST("/chain/checkin", func(context *gin.Context) {
			var data models.TransactionData
			if err := context.ShouldBindJSON(&data); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			block, err := svc.CheckIn(data)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusCreated, block)
		})

		api.GET("/chain/validate", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.ValidateChain())
		})

		api.POST("/chain/tamper", func(context *gin.Context) {
			var req tamperRequest
			if err := context.ShouldBindJSON(&req); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			chain, err := svc.Tamper(req.Index, req.Data)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, chain)
		})

		api.POST("/chain/restore", func(context *gin.Context) {
			context.JSON(http.StatusOK, svc.RestoreChain())
		})
	}

	return router
}

type foodRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ExpiryDate  string  `json:"expiryDate"`
	Destination string  `json:"destination"`
	Weight      float64 `json:"weight"`
	Urgency     int     `json:"urgency"`
}

type compareRequest struct {
	FoodID string `json:"foodId"`
}

type tamperRequest struct {
	Index int                    `json:"index"`
	Data  models.TransactionData `json:"data"`
}

func parseDate(value string) (time.Time, error) {
	if parsed, err := time.Parse("2006-01-02", value); err == nil {
		return parsed, nil
	}
	if millis, err := strconv.ParseInt(value, 10, 64); err == nil {
		return time.UnixMilli(millis), nil
	}
	return time.Parse(time.RFC3339, value)
}
