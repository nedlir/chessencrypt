package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nedlir/chessencrypt/chess/pgn"
	"github.com/nedlir/chessencrypt/utils/fileshandler"
)

func Run() error {
	if len(os.Args) < 2 {
		printUsage()
		return errors.New("missing command")
	}

	cmd := os.Args[1]
	args := os.Args[2:]
	switch cmd {
	case "encode":
		return runEncode(args)
	case "decode":
		return runDecode(args)
	case "--help", "-h":
		printUsage()
		return nil
	default:
		printUsage()
		return fmt.Errorf("unknown command %q", cmd)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  app encode <inputFile> <outputDir>")
	fmt.Println("  app decode <inputDir> <outputFile>")
	fmt.Println("  app --help | -h")
}

func runEncode(args []string) error {
	if len(args) != 2 {
		printUsage()
		return errors.New("encode requires <inputFile> and <outputDir>")
	}
	inFile, outDir := args[0], args[1]

	if err := fileshandler.CreateDir(outDir); err != nil {
		return err
	}

	data, err := fileshandler.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	enc := pgn.NewPGNEncoder()
	part := 1
	for off := 0; off < len(data); off += fileshandler.ChunkSize {
		end := off + fileshandler.ChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[off:end]
		name := fmt.Sprintf(fileshandler.PartNamePattern, part)
		path := filepath.Join(outDir, name)

		text := enc.BytesToPgn(chunk, part)
		if err := fileshandler.WriteFile(path, []byte(text)); err != nil {
			return fmt.Errorf("write error: %w", err)
		}
		fmt.Printf("Wrote %s (%d bytes)\n", path, len(chunk))
		part++
	}
	fmt.Printf("Encoded into %d parts\n", part-1)
	return nil
}

func runDecode(args []string) error {
	if len(args) != 2 {
		printUsage()
		return errors.New("decode requires <inputDir> and <outputFile>")
	}
	inDir, outFile := args[0], args[1]

	entries, err := fileshandler.ListDir(inDir)
	if err != nil {
		return fmt.Errorf("dir error: %w", err)
	}

	dec := pgn.NewPGNDecoder()
	var all []byte

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != fileshandler.PartExtension {
			continue
		}
		data, err := fileshandler.ReadFile(filepath.Join(inDir, entry.Name()))
		if err != nil {
			return fmt.Errorf("read error: %w", err)
		}
		chunk := dec.PGNToBytes(string(data))
		all = append(all, chunk...)
		fmt.Printf("Decoded %s (%d bytes)\n", entry.Name(), len(chunk))
	}

	if err := fileshandler.WriteFile(outFile, all); err != nil {
		return fmt.Errorf("write error: %w", err)
	}
	fmt.Printf("Decoded output written to %s (%d bytes)\n", outFile, len(all))
	return nil
}
