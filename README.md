# Huffman Compressor

This project is an implementation of a file compressor and decompressor based on the Huffman algorithm. It allows you to compress text files (`.txt`) into a compressed format (`.hoff`) and then decompress them back.

## Project Structure

- `cmd/main.go`: Main entry point to run the program.
- `internal/filemanager`: Contains functions for handling file encoding and decoding.
- `internal/tree`: Contains the implementation of the Huffman tree.
- `internal/utils`: Contains additional utilities.

## Usage

1) Compress
   `go run ./cmd/main.go -c /path/to/file.txt`
2) Decompress
   `go run ./cmd/main.go -d /path/to/file.hoff`
