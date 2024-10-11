package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const certPath = "./cert"
const developEnv = "https://acme-staging-v02.api.letsencrypt.org/directory"
const productEnv = "https://acme-v02.api.letsencrypt.org/directory"

func main() {
	privateKey, err := generatePrivateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	_, err = generateCSR(privateKey)
	if err != nil {
		fmt.Println("Error generating CSR:", err)
		return
	}

	fmt.Println("CSR generated and saved as certificate_request.csr")
	err = submitCSR("certificate_request.csr")
	if err != nil {
		fmt.Println("Error submitting CSR:", err)
		return
	}

	fmt.Println("CSR successfully submitted to CA.")
}

func generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Save the private key to a file
	privateKeyFile, err := os.Create(path.Join(certPath, "private_key.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to create private key file: %v", err)
	}
	defer privateKeyFile.Close()

	// Encode the private key to PEM format
	pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return privateKey, nil
}

func generateCSR(privateKey *rsa.PrivateKey) ([]byte, error) {
	csrTemplate := &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "www.example.com", // 替换为你的域名
			Organization: []string{"Example Organization"},
		},
		DNSNames: []string{"www.example.com", "example.com"},
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, csrTemplate, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate CSR: %v", err)
	}

	// Save the CSR to a file
	csrFile, err := os.Create(path.Join(certPath, "certificate_request.csr"))
	if err != nil {
		return nil, fmt.Errorf("failed to create CSR file: %v", err)
	}
	defer csrFile.Close()

	pem.Encode(csrFile, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})

	return csrBytes, nil
}

// 提交 CSR 给证书颁发机构 (CA)
// 将生成的 CSR 文件提交给证书颁发机构（例如，Let's Encrypt、DigiCert、Comodo 等）。
// CA 通过这个 CSR 验证你的身份，并生成签名后的 SSL 证书。
func submitCSR(csrFile string) error {
	csrBytes, err := os.ReadFile(path.Join(certPath, csrFile))
	if err != nil {
		return fmt.Errorf("failed to read CSR file: %v", err)
	}

	resp, err := http.Post(developEnv, "application/x-pem-file", bytes.NewBuffer(csrBytes))
	if err != nil {
		return fmt.Errorf("failed to submit CSR: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	fmt.Println("CSR submitted, response:", string(body))
	return nil
}

func confirmAuth() {
	//CA 验证你是否是该域名的真正所有者。常见的验证方式有：
	//DNS 验证：你需要在 DNS 记录中添加指定的 TXT 记录，CA 会查询并验证此记录。
	//文件验证：CA 要求你在服务器的某个位置放置一个特定文件，CA 会访问该文件以确认你的域名控制权。
	//邮箱验证：CA 向你提供的邮箱地址（通常是域名所有者的邮箱）发送验证邮件。
	//CA 通过验证后，会为你颁发 SSL 证书。
}

func getCert() {
	//当 CA 验证通过后，CA 会生成并签发 SSL 证书文件，并提供给你。
	//证书文件通常包含：
	//证书文件 (.crt 或 .pem 格式)：这是你可以在服务器上安装的证书。
	//中间证书：有时 CA 也会提供中间证书，确保证书链的完整性。
}

func installCert() {
	//将你获得的证书安装在你的服务器上，与私钥一起使用。
	//安装过程取决于你的服务器类型。比如：
	//对于 Nginx 或 Apache，可以将证书路径和私钥路径配置在服务器的配置文件中。
	//对于 Kubernetes、云服务器，可能需要将其配置为密钥或秘密资源。
}

func forceReload() {
	//安装完 SSL 证书后，需要在 Web 服务器上配置 HTTPS。
	//确保服务器正确引用了私钥文件和 SSL 证书文件。
}

func renewal() {
	//续订
}
