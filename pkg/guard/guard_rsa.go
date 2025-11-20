package guard

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var RSAKeyPath = "./screet-key"

// GenRSA generates an RSA key pair if the key files don't exist.
func GenRSA() {
	// Create folder if not exists
	if err := os.MkdirAll(RSAKeyPath, 0700); err != nil {
		log.Println("Failed to create directory:", err)
		return
	}

	privateKeyPath := "private.pem"
	publicKeyPath := "public.pem"

	root, err := os.OpenRoot(RSAKeyPath)
	if err != nil {
		log.Println("Failed open root :", err)
	}

	_, errPublic := root.Stat(publicKeyPath)
	_, errPrivate := root.Stat(privateKeyPath)

	if os.IsNotExist(errPublic) || os.IsNotExist(errPrivate) {
		fmt.Println("Generating RSA key pair...")
		privateKey, publicKey, err := generateRSAKeyPair(2048)
		if err != nil {
			log.Println("Failed to generate RSA keys:", err)
			return
		}

		if err := savePEMKey(root, privateKeyPath, privateKey); err != nil {
			log.Println("Failed to save private key:", err)
		}
		if err := savePublicPEMKey(root, publicKeyPath, publicKey); err != nil {
			log.Println("Failed to save public key:", err)
		}
	}
}

// generateRSAKeyPair creates a new RSA private and public key pair.
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return key, &key.PublicKey, nil
}

// savePEMKey writes a private RSA key to a PEM file safely.
func savePEMKey(root *os.Root, filename string, key *rsa.PrivateKey) error {
	file, err := root.Create(filepath.Base(filename)) // 👈 Secure, scoped
	if err != nil {
		return err
	}
	defer file.Close()

	block := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	return pem.Encode(file, block)
}

// savePublicPEMKey writes a public RSA key to a PEM file safely.
func savePublicPEMKey(root *os.Root, filename string, pub *rsa.PublicKey) error {
	file, err := root.Create(filepath.Base(filename)) // 👈 Secure, scoped
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}
	block := &pem.Block{Type: "PUBLIC KEY", Bytes: bytes}
	return pem.Encode(file, block)
}
