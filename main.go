package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	var keySize int
	var privateKeyPath string
	var publicKeyPath string

	// Define command-line flags
	flag.IntVar(&keySize, "size", 2048, "The size of the generated RSA key in bits")
	flag.StringVar(&privateKeyPath, "private-key", "id_rsa.pem", "The path to the output file for the private key")
	flag.StringVar(&publicKeyPath, "public-key", "id_rsa.pub", "The path to the output file for the public key")

	// Parse command-line flags
	flag.Parse()

	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Printf("Error generating private key: %v\n", err)
		os.Exit(1)
	}

	// Encode private key as PEM block
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Output private key to file or stdout
	if privateKeyPath != "" {
		privateKeyFile, err := os.Create(privateKeyPath)
		if err != nil {
			fmt.Printf("Error creating private key file: %v\n", err)
			os.Exit(1)
		}
		defer privateKeyFile.Close()
		err = pem.Encode(privateKeyFile, privateKeyPEM)
		if err != nil {
			fmt.Printf("Error encoding private key: %v\n", err)
			os.Exit(1)
		}
	} else {
		err = pem.Encode(os.Stdout, privateKeyPEM)
		if err != nil {
			fmt.Printf("Error encoding private key: %v\n", err)
			os.Exit(1)
		}
	}

	// Generate public key from private key
	publicKey, err := sshPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Printf("Error generating public key: %v\n", err)
		os.Exit(1)
	}

	// Output public key to file or stdout
	if publicKeyPath != "" {
		publicKeyFile, err := os.Create(publicKeyPath)
		if err != nil {
			fmt.Printf("Error creating public key file: %v\n", err)
			os.Exit(1)
		}
		defer publicKeyFile.Close()
		err = pem.Encode(publicKeyFile, publicKey)
		if err != nil {
			fmt.Printf("Error encoding public key: %v\n", err)
			os.Exit(1)
		}
	} else {
		err = pem.Encode(os.Stdout, publicKey)
		if err != nil {
			fmt.Printf("Error encoding public key: %v\n", err)
			os.Exit(1)
		}
	}
}

// sshPublicKey encodes an RSA public key as an SSH public key
func sshPublicKey(publicKey *rsa.PublicKey) (*pem.Block, error) {
	pub, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pubBytes := ssh.MarshalAuthorizedKey(pub)
	return &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	}, nil
}
