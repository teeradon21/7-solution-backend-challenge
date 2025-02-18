package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func makeBeefSummary(data string) map[string]int {

	lowercase := strings.ToLower(data)
	words := strings.Fields(lowercase)

	summary := map[string]int{}

	for _, word := range words {
		name := strings.Trim(word, ",.")
		_, exist := summary[name]
		if !exist {
			summary[name] = 0
		}
		summary[name] += 1
	}

	return summary
}

func getBeefSummary(c *gin.Context) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	summary := makeBeefSummary(string(body))

	c.JSON(http.StatusOK, gin.H{
		"beef": summary,
	})
}

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.GET("/beef/summary", getBeefSummary)

	router.Run()
}
