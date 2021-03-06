package main
 
import (
    "bytes"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)
 
type License struct {
    KeyID         string `json:"key_id"`
    PrivateKey    string `json:"private_key"`
    Authorization string `json:"authorization"`
}
 
type DockerHeader struct {
    DockerKeyID     string `json:"X-DOCKER-KEY-ID"`
    DockerToken     string `json:"X-DOCKER-TOKEN"`
    DockerTimestamp string `json:"X-DOCKER-TIMESTAMP"`
}
 
func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: gen_token <license_file>\n")
        os.Exit(1)
    }
 
    var license License
    var err error
    if license, err = readLicense(os.Args[1]); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
 
    now := time.Now().Add(time.Hour).Format(time.RFC3339)
    token, err := generateToken(now, license.PrivateKey)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
 
    dockerHeader := &DockerHeader{
        DockerKeyID:     license.KeyID,
        DockerToken:     token,
        DockerTimestamp: now,
    }
 
    res, _ := json.Marshal(dockerHeader)
    fmt.Printf("%v\n", string(res))
    os.Exit(0)
}
 
func readLicense(filepath string) (License, error) {
    var license License
    file, err := os.Open(filepath)
    if err != nil {
        return license, err
    }
    defer file.Close()
 
    body, err := ioutil.ReadAll(file)
    if err != nil {
        return license, err
    }
 
    // There shouldn't be a BOM at the beginning of the license
    // file, but strip it out if there is
    body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
 
    err = json.Unmarshal(body, &license)
    if err != nil {
        return license, err
    }
 
    return license, nil
}
 
// GenerateToken generates a hash of the message with the privateKey via the
// sha256 algorithm.
func generateToken(message, privateKey string) (string, error) {
    key, err := base64.URLEncoding.DecodeString(privateKey)
    if err != nil {
        return "", err
    }
 
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
}
