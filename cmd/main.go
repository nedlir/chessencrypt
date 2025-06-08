package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nedlir/chessencrypt/chess/pgn"
	"github.com/nedlir/chessencrypt/utils/fileshandler"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./cmd/main.go encode <inputFilePath>  <outputDirectoryPath>")
	fmt.Println("  go run ./cmd/main.go decode <inputDirectoryPath> <outputFilePath>")
	fmt.Println("  go run ./cmd/main.go --help | -h")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing command")
		printUsage()
		return
	}

	switch cmd := os.Args[1]; cmd {
	case "--help", "-h":
		printUsage()

	case "encode":
		if len(os.Args) != 4 {
			fmt.Println("Error: encode requires inputFilePath and outputDirectoryPath")
			printUsage()
			return
		}
		inputFile, outputDir := os.Args[2], os.Args[3]

		if err := fileshandler.CreateDir(outputDir); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		data, err := fileshandler.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Error reading %q: %v\n", inputFile, err)
			return
		}
		fmt.Printf("Input file size: %d bytes\n", len(data))

		encoder := pgn.NewPGNEncoder()
		part := 1
		for offset := 0; offset < len(data); offset += fileshandler.ChunkSize {
			end := offset + fileshandler.ChunkSize
			if end > len(data) {
				end = len(data)
			}
			chunk := data[offset:end]

			fileName := fmt.Sprintf(fileshandler.PartNamePattern, part)
			outPath := filepath.Join(outputDir, fileName)

			pgnText := encoder.BytesToPgn(chunk, part)
			if err := fileshandler.WriteFile(outPath, []byte(pgnText)); err != nil {
				fmt.Printf("Error writing %q: %v\n", outPath, err)
				return
			}
			fmt.Printf("Created %s (%d bytes)\n", outPath, len(chunk))
			part++
		}
		fmt.Printf("\nFinished encoding into %d files in %s\n", part-1, outputDir)

	case "decode":
		if len(os.Args) != 4 {
			fmt.Println("Error: decode requires inputDirectoryPath and outputFilePath")
			printUsage()
			return
		}
		inputDir, outputFile := os.Args[2], os.Args[3]

		entries, err := fileshandler.ListDir(inputDir)
		if err != nil {
			fmt.Printf("Error reading directory %q: %v\n", inputDir, err)
			return
		}

		decoder := pgn.NewPGNDecoder()
		var allData []byte

		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != fileshandler.PartExtension {
				continue
			}
			path := filepath.Join(inputDir, entry.Name())
			pgnBytes, err := fileshandler.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading %q: %v\n", path, err)
				return
			}
			decoded := decoder.PGNToBytes(string(pgnBytes))
			allData = append(allData, decoded...)
			fmt.Printf("Decoded %s: %d bytes\n", path, len(decoded))
		}

		if err := fileshandler.WriteFile(outputFile, allData); err != nil {
			fmt.Printf("Error writing output file %q: %v\n", outputFile, err)
			return
		}
		fmt.Printf("\nDecoding complete! Output written to %s (%d bytes)\n", outputFile, len(allData))

	default:
		fmt.Printf("Error: Unknown command %q\n", cmd)
		printUsage()
	}
}
