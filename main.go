package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "net/http"
)

func generateAESKey() ([]byte, error) {
    key := make([]byte, 32) // AES-256 key length
    _, err := rand.Read(key)
    if err != nil {
        return nil, err
    }
    return key, nil
}

func encryptData(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

    return ciphertext, nil
}

func decryptData(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    return ciphertext, nil
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
    key, err := generateAESKey()
    if err != nil {
        http.Error(w, "Failed to generate encryption key", http.StatusInternalServerError)
        return
    }

    // Read input data from request body
    inputData, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

    encryptedData, err := encryptData(inputData, key)
    if err != nil {
        http.Error(w, "Encryption failed", http.StatusInternalServerError)
        return
    }

    // Respond with encrypted data as hex-encoded string
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(hex.EncodeToString(encryptedData)))
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
    key, err := generateAESKey() // Use the same key management strategy
    if err != nil {
        http.Error(w, "Failed to generate decryption key", http.StatusInternalServerError)
        return
    }

    // Read encrypted data (hex-encoded) from request body
    encryptedHex, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
        return
    }

    encryptedData, err := hex.DecodeString(string(encryptedHex))
    if err != nil {
        http.Error(w, "Failed to decode hex string", http.StatusBadRequest)
        return
    }

    decryptedData, err := decryptData(encryptedData, key)
    if err != nil {
        http.Error(w, "Decryption failed", http.StatusInternalServerError)
        return
    }

    // Respond with decrypted data
    w.Header().Set("Content-Type", "text/plain")
    w.Write(decryptedData)
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/encrypt", encryptHandler)
    http.HandleFunc("/decrypt", decryptHandler)
    http.ListenAndServe(":8080", nil)
}
