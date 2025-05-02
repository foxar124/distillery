package cosign_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	"github.com/glamorousis/distillery/pkg/cosign"
)

func TestParsePublicKey(t *testing.T) {
	// Generate a test ECDSA key pair
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key: %v", err)
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		t.Fatalf("Failed to marshal public key: %v", err)
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	// Test parsing the public key
	parsedPubKey, err := cosign.ParsePublicKey(pubKeyPEM)
	if err != nil {
		t.Fatalf("Failed to parse public key: %v", err)
	}
	if !parsedPubKey.Equal(&privKey.PublicKey) {
		t.Fatalf("Parsed public key does not match original")
	}
}

func TestVerifySignature(t *testing.T) {
	// Generate a test ECDSA key pair
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key: %v", err)
	}

	// Create test data and sign it
	data := []byte("test data")
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)

	sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash)
	if err != nil {
		t.Fatalf("Failed to sign data: %v", err)
	}

	// Encode the signature in base64
	signatureBase64 := base64.StdEncoding.EncodeToString(sig)
	fmt.Println("Signature:", signatureBase64)

	// Test verifying the signature
	valid, err := cosign.VerifySignature(&privKey.PublicKey, hash, []byte(signatureBase64))
	if err != nil {
		t.Fatalf("Failed to verify signature: %v", err)
	}
	if !valid {
		t.Fatalf("Signature verification failed")
	}
}

func TestVerifyChecksumSignature(t *testing.T) {
	// Read the contents of checksums.txt.pem
	publicKeyContentEncoded, err := os.ReadFile("testdata/checksums.txt.pem")
	if err != nil {
		t.Fatalf("Failed to read public key file: %v", err)
	}

	// Decode the base64-encoded public key
	publicKeyContent, err := base64.StdEncoding.DecodeString(string(publicKeyContentEncoded))
	if err != nil {
		t.Fatalf("Failed to decode base64 public key: %v", err)
	}

	// Read the contents of checksums.txt.sig
	signatureContent, err := os.ReadFile("testdata/checksums.txt.sig")
	if err != nil {
		t.Fatalf("Failed to read signature file: %v", err)
	}

	// Read the contents of checksums.txt
	dataContent, err := os.ReadFile("testdata/checksums.txt")
	if err != nil {
		t.Fatalf("Failed to read data file: %v", err)
	}

	// Decode the PEM-encoded public key
	pubKey, err := cosign.ParsePublicKey(publicKeyContent)
	if err != nil {
		t.Fatalf("Failed to parse public key: %v", err)
	}

	dataHash := cosign.HashData(dataContent)

	// Verify the signature
	valid, err := cosign.VerifySignature(pubKey, dataHash, signatureContent)
	if err != nil {
		t.Fatalf("Failed to verify signature: %v", err)
	}
	if !valid {
		t.Fatalf("Signature verification failed")
	}
}

func TestVerifyChecksumSignaturePublicKey(t *testing.T) {
	// Read the contents of checksums.txt.pem
	publicKeyContent, err := os.ReadFile("testdata/release.pub")
	if err != nil {
		t.Fatalf("Failed to read public key file: %v", err)
	}

	// Read the contents of checksums.txt.sig
	signatureContent, err := os.ReadFile("testdata/release.sig")
	if err != nil {
		t.Fatalf("Failed to read signature file: %v", err)
	}

	// Decode the PEM-encoded public key
	pubKey, err := cosign.ParsePublicKey(publicKeyContent)
	if err != nil {
		t.Fatalf("Failed to parse public key: %v", err)
	}

	dataHashEncoded, err := os.ReadFile("testdata/release.sha256")
	if err != nil {
		t.Fatalf("Failed to read data file: %v", err)
	}

	dataHash, err := base64.StdEncoding.DecodeString(string(dataHashEncoded))
	if err != nil {
		t.Fatalf("Failed to decode base64 data hash: %v", err)
	}

	// Verify the signature
	valid, err := cosign.VerifySignature(pubKey, dataHash, signatureContent)
	if err != nil {
		t.Fatalf("Failed to verify signature: %v", err)
	}
	if !valid {
		t.Fatalf("Signature verification failed")
	}
}
