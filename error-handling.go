package main

import (
    "errors"
    "fmt"
    "os"
)

func openFile(filename string) error {
    _, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    return nil
}

func main() {
    err := openFile("nonexistent.txt")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File does not exist.")
    } else {
        fmt.Println("Error:", err)
    }
}
