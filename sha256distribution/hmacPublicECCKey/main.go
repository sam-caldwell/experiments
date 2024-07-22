package main

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
)

const numHashes = 1000000000

func main() {
	// Initialize counters for each byte value (0-255)
	byteCount := make([]int, 256)

	// Loop to generate Ed25519 public keys and calculate hashes
	for i := 0; i < numHashes; i++ {
		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println("Error generating Ed25519 key:", err)
			return
		}

		// Calculate HMAC of the public key with SHA256
		h := hmac.New(sha256.New, []byte(privateKey))
		h.Write(publicKey)
		hmacResult := h.Sum(nil)

		// Analysis (optional)
		for _, b := range hmacResult {
			byteCount[b]++
		}
		fmt.Printf("progress: %4.2f%%\n", float64(100*i)/float64(numHashes))
	}

	// Calculate expected count per byte (assuming uniform distribution)
	expectedCount := float64(numHashes) * 32 / 256 // 32 bytes per hash

	minDiff := math.MaxFloat64
	maxDiff := -math.MaxFloat64
	totalAbsoluteDifference := 0.0
	totalCount := 0
	hitCount := 0
	hitIndexes := []int{}

	// Track zero-hits
	zeroHits := 0
	zeroHitIndexes := []int{}

	// Analyze and print byte distribution
	fmt.Println("Byte Distribution Analysis:")
	fmt.Println("Byte | Count  | Expected       | Difference")
	for b, count := range byteCount {
		difference := math.Abs(float64(count) - expectedCount)
		totalCount += count
		AbsoluteDifference := difference
		totalAbsoluteDifference += AbsoluteDifference
		if AbsoluteDifference < minDiff {
			minDiff = difference
		}
		if AbsoluteDifference > maxDiff {
			maxDiff = difference
		}
		if count == 0 {
			zeroHits++
			zeroHitIndexes = append(zeroHitIndexes, b)
		}
		if AbsoluteDifference == 0 {
			hitCount++
			hitIndexes = append(hitIndexes, b)
		}
		fmt.Printf("%02x   | %6d | %15.4f | %15.2f\n", b, count, expectedCount, difference)
	}

	averageAbsoluteDifference := totalAbsoluteDifference / 256
	fmt.Printf("Average Absolute Difference: %.6f\n", averageAbsoluteDifference)
	fmt.Printf("Minimum Difference: %.6f\n", minDiff)
	fmt.Printf("Maximum Difference: %.6f\n", maxDiff)
	fmt.Printf("Total count: %d\n", totalCount)
	fmt.Printf("Zero hits: %d\n", zeroHits)
	fmt.Printf("Zero hit indexes: %v\n", zeroHitIndexes)
	fmt.Printf("     hits: %d\n", hitCount)
	fmt.Printf("     hit indexes: %v\n", hitIndexes)
	fmt.Println("public key bytes")
}
