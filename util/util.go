package util

import (
	"japapp/models"
	"math/rand"
	"time"
)

func Shuffle(vals []models.Question) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := len(vals)
	for i , _ := range vals {
		randIndex := r.Intn(n)
		vals[i], vals[randIndex] = vals[randIndex], vals[i]
	}
}
