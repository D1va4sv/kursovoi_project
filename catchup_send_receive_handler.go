package main

import (
    "bytes"
    "compress/gzip"
    "io"
    "log"
    "os"
    "path/filepath"
    "wal-g/internal/config"
)

func compress(data []byte) ([]byte, error) {
    var buf bytes.Buffer
    gz := gzip.NewWriter(&buf)
    gz.Level = config.ConfigData.Compression.Level
    if _, err := gz.Write(data); err != nil {
        return nil, err
    }
    if err := gz.Close(); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
    buf := bytes.NewBuffer(data)
    gz, err := gzip.NewReader(buf)
    if err != nil {
        return nil, err
    }
    defer gz.Close()
    var res bytes.Buffer
    if _, err := io.Copy(&res, gz); err != nil {
        return nil, err
    }
    return res.Bytes(), nil
}

func catchupSend(dataPath string) {
    data, err := ioutil.ReadFile(dataPath)
    if err != nil {
        log.Fatalf("Failed to read data: %v", err)
    }

    if config.ConfigData.Compression.Enabled {
        data, err = compress(data)
        if err != nil {
            log.Fatalf("Failed to compress data: %v", err)
        }
    }

    // шифрование данных
    encryptedData := encrypt(data) // предположим, что функция encrypt уже определена

    // отправка данных
    err = send(encryptedData) // предположим, что функция send уже определена
    if err != nil {
        log.Fatalf("Failed to send data: %v", err)
    }
}

func catchupReceive(destinationPath string) {
    // получение данных
    receivedData, err := receive() // предположим, что функция receive уже определена
    if err != nil {
        log.Fatalf("Failed to receive data: %v", err)
    }

    // расшифровка данных
    decryptedData := decrypt(receivedData) // предположим, что функция decrypt уже определена

    if config.ConfigData.Compression.Enabled {
        decryptedData, err = decompress(decryptedData)
        if err != nil {
            log.Fatalf("Failed to decompress data: %v", err)
        }
    }

    err = ioutil.WriteFile(destinationPath, decryptedData, 0644)
    if err != nil {
        log.Fatalf("Failed to write data: %v", err)
    }
}

func main() {
    // Использование функций catchupSend и catchupReceive в зависимости от аргументов командной строки или другого логики
    // Например:
    if len(os.Args) < 3 {
        log.Fatalf("Usage: %s <send|receive> <data path>", filepath.Base(os.Args[0]))
    }

    command := os.Args[1]
    dataPath := os.Args[2]

    if command == "send" {
        catchupSend(dataPath)
    } else if command == "receive" {
        catchupReceive(dataPath)
    } else {
        log.Fatalf("Unknown command: %s", command)
    }
}
