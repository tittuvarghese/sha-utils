package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"bufio"
)

func hashString(s string, algorithm string, bitSize int) string {
	var hash []byte
	switch algorithm {
	case "SHA1":
		h := sha1.New()
		h.Write([]byte(s))
		hash = h.Sum(nil)
	case "SHA256":
		h := sha256.New()
		h.Write([]byte(s))
		hash = h.Sum(nil)
	case "SHA512":
		h := sha512.New()
		h.Write([]byte(s))
		hash = h.Sum(nil)
	default:
		fmt.Println("Unsupported algorithm:", algorithm)
		return ""
	}

	// Convert hash to hexadecimal string
	return fmt.Sprintf("%x", hash)
}

func main() {
	// Define command line flags
	inputFile := flag.String("input", "input.txt", "Path to the input file")
	outputFile := flag.String("output", "output.txt", "Path to the output file")
	algorithm := flag.String("algorithm", "SHA256", "Hashing algorithm (SHA1, SHA256, SHA512)")
	bitSize := flag.Int("bits", 256, "Bit size of the hash (128, 256, 512)")
	flag.Parse()

	// Validate bit size
	if *bitSize != 128 && *bitSize != 256 && *bitSize != 512 {
		fmt.Println("Unsupported bit size:", *bitSize)
		return
	}

	// Open input file
	inFile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inFile.Close()

	// Open output file
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Process each line in the input file
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		hashed := hashString(line, *algorithm, *bitSize)
		if hashed != "" {
			_, err := outFile.WriteString(hashed + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}
