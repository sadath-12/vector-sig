package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ShoppingCart struct {
	Items      []string `json:"items"`
	Total      float64  `json:"total"`
	Discount   float64  `json:"discount"`
	FinalPrice float64  `json:"final_price"`
}

func main() {
	// Seed random number generator for message variations
	rand.Seed(time.Now().UnixNano())

	// Create a new logger instance.
	logger := logrus.New()

	// Set formatter to JSON for structured logging
	logger.Formatter = &logrus.JSONFormatter{}

	// Set the log level to control verbosity.
	logger.SetLevel(logrus.DebugLevel)

	// Bombard logs with a for loop
	for i := 0; i < 100000; i++ {

		// Info message with user object, shopping cart details, and performance metrics
		user := User{ID: rand.Intn(1000) + 1, Username: fmt.Sprintf("user_%d", i/3), Email: fmt.Sprintf("user_%d@example.com", i/3)}
		cart := ShoppingCart{
			Items:      []string{"Product A", "Product B", fmt.Sprintf("Mystery Item %d", i/3)},
			Total:      rand.Float64() * 100,
			Discount:   rand.Float64() * 20,
			FinalPrice: (rand.Float64() * 100) - (rand.Float64() * 20),
		}
		startTime := time.Now()
		memStats := runtime.MemStats{}
		runtime.ReadMemStats(&memStats)
		logger.WithFields(logrus.Fields{
			"source":          "Shopping Service",
			"operation":       "ProcessOrder",
			"user":            user,
			"cart":            cart,
			"processing_time": startTime,
			"memory_usage":    memStats.Alloc / 1024 / 1024, // Log memory usage in MB
			"message":         fmt.Sprintf("Order processed for user %s (cart total: %.2f, final price: %.2f)", user.Username, cart.Total, cart.FinalPrice),
		}).Info("Information message")

		cacheKey := "example_cache_key"

		logger.WithFields(logrus.Fields{
			"component": "Cache",
			"message":   fmt.Sprintf("Cache miss for key %s. Potential data inconsistency", cacheKey),
		}).Warn("Warning message")

		err := fmt.Errorf("example error occurred")

		logger.WithFields(logrus.Fields{
			"component": "Cache",
			"cache_key": cacheKey,
			"error":     err.Error(),
			"message":   fmt.Sprintf("Error unmarshalling cached data for key %s", cacheKey),
		}).Warn("Warning message")

		// Error message with detailed error object and stack
		stack := make([]byte, 4096)
		stack = stack[:runtime.Stack(stack, false)]
		logger.WithFields(logrus.Fields{
			"error":       err.Error(),
			"stack_trace": string(stack),
		}).Error("Error message")

		time.Sleep(1 * time.Second)
	}
}
