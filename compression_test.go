package main

import (
    "bytes"
    "compress/gzip"
    "io/ioutil"
    "testing"
)

func TestCompressDecompress(t *testing.T) {
    originalData := []byte("This is some test data to compress and decompress")
    compressedData, err := compress(originalData)
    if err != nil {
        t.Fatalf("Failed to compress data: %v", err)
    }

    decompressedData, err := decompress(compressedData)
    if err != nil {
        t.Fatalf("Failed to decompress data: %v", err)
    }

    if !bytes.Equal(originalData, decompressedData) {
        t.Fatalf("Decompressed data does not match original data\nOriginal: %s\nDecompressed: %s", originalData, decompressedData)
    }
}

func TestCompressionLevel(t *testing.T) {
    originalData := []byte("This is some test data to compress with different levels")
    
    for level := gzip.HuffmanOnly; level <= gzip.BestCompression; level++ {
        config.ConfigData.Compression.Level = level

        compressedData, err := compress(originalData)
        if err != nil {
            t.Fatalf("Failed to compress data at level %d: %v", level, err)
        }

        decompressedData, err := decompress(compressedData)
        if err != nil {
            t.Fatalf("Failed to decompress data at level %d: %v", level, err)
        }

        if !bytes.Equal(originalData, decompressedData) {
            t.Fatalf("Decompressed data does not match original data at level %d\nOriginal: %s\nDecompressed: %s", level, originalData, decompressedData)
        }
    }
}
