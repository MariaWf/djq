package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"github.com/pkg/errors"
)

// 公钥和私钥可以从文件中读取
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
//7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
//Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
//AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
//ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
//XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
///jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
//IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
//4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
//DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
//9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
//DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
//AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
//-----END RSA PRIVATE KEY-----
//`)
//
//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
//ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
//wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
//AUeJ6PeW+DAkmJWF6QIDAQAB
//-----END PUBLIC KEY-----
//`)

//1024bit
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQCHOv95x+E250MsyboOYeNHTphASEbMvjtzPVvJw5uGnxUXrq3s
//GW875SFJLSXMH7rJQlHyDnZt1e/2ZZRI4l89ld86utwXQ5YFVKRuxYwbtpLNGXMy
//z++Jhu2rybY9+dQq017JXgOtwwn8y+M3Ml+PmaIYpw4jf+D1mFeDKXjvgwIDAQAB
//AoGAEx+wyVQO8Wi7AzZz7VzCe28N8OlHueSdG6rttBbJg7wzi2nUhrMCHAJHNsX/
//tmI2VAfg9s48yUOb78hls/jvZt1XgRzS47wLydvcp4cYEZ6ekzStVa2bV6P8wjX1
//Zqc1UXFHnY9gapNuH8o7iR0oX4TafT9SSRFL+uEodq/B8wECQQDLuPgfSW1qp+KJ
//6w5amHSRbzqwJ39cf90NQoLfoGf1AYhJQtfAkHzR98fAcFsPiAmfuDBxdi20WhHK
//94CTyXHjAkEAqe6d+NMP96iyb4CaUf4hKOPhbkCymKRz0EepAUxmaNCgz3dZZA1g
//yX9HtRDgGDGpoT9BizwGDFqwnXEHjPB94QJATiBQX0c8g9OAaB3RsmKXCZMbcaSk
//Digm8MfaAsK0O1xsFJRiw0Fl7OvWGfG8qjckYbE0Or70hh6ohirmj0aIuwJBAIjX
//XaDtHhOaZq6BykOyuLM75uIo+WkQLS8RNdiU1HcdYhIPkw2N9F3uwCjgAQWaoHX4
//AWkFGf3C3iy6w5DMbSECQQCuXOANOOAdrJLdnjVLaQQOccrdcj2gw4kTvb2nIm8M
//geuMXV65kekwJrxEkmVjZNTzYqqMhnuiTb04JTx7NcCP
//-----END RSA PRIVATE KEY-----
//`)
//
//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHOv95x+E250MsyboOYeNHTphA
//SEbMvjtzPVvJw5uGnxUXrq3sGW875SFJLSXMH7rJQlHyDnZt1e/2ZZRI4l89ld86
//utwXQ5YFVKRuxYwbtpLNGXMyz++Jhu2rybY9+dQq017JXgOtwwn8y+M3Ml+PmaIY
//pw4jf+D1mFeDKXjvgwIDAQAB
//-----END PUBLIC KEY-----
//`)

//512bit
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIIBOgIBAAJBAJfkiIRBIr+3KTx8bLBIkYWa7PTJdCTdt/MXBq3wv+FrXwkpoCTf
//U6JuSPmNPHshhlxBzTzTqNrlvLj7mjzeMCMCAwEAAQJAJtK77gn0BsqbGKG46iny
//QBMEQ+EF9bJSJSkahPUHJmd1cwz/4VcTbs25vGbVgfaUkmvGQkhBDiynOhJFhh1N
//OQIhAMngL78RvybgC2qZO/Ahs1CZhsJeTLRhl2Xmegf2hjz/AiEAwJ2+feZ4a6JS
//HFCg+8RU2Q5UJi2y5xw5RgP72q2ceN0CICsiRrFpplE1H9tYAHGPkdPP6fZP1c3Y
//6FNyinJ3HzmbAiEAvcPro/BHHYvLJ7hMEh9IlJwoTYDibA0DZDaSj8xRh0UCIFLH
//dQZSFkXazJo6WN+UolqkRzmUCIq1+KQFoeWdKp0b
//-----END RSA PRIVATE KEY-----
//`)
//
//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJfkiIRBIr+3KTx8bLBIkYWa7PTJdCTd
//t/MXBq3wv+FrXwkpoCTfU6JuSPmNPHshhlxBzTzTqNrlvLj7mjzeMCMCAwEAAQ==
//-----END PUBLIC KEY-----
//`)

//2048bit
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQBwkBfXoaXcl+/pw0q8Wg4qM3nuREleXPQb9/pGhPCX9K0P+T82
wTLKj5aweEQ85vg4HTDS3eYVPaYarN5gELNjSpse/QsC2SHc47TQgUDz+EfzXFPk
kX6S7W+nw/+8y54AqfnuIjEMWuhHjF4BzwgP8ktb2xNqR3tEPB3gNrRPagLNCrqe
uKUI0ne2s19ML/6I+O9BB9oq3ejkdNwf/ae6B3XnPbn0LNjnt5ufTQrmr36aUE2W
OfmxK0ikAOanFJZCdrJlxh4u5cLp1QFqiNZd92J0wDR6Nwt2n2Wuu2kYotiAoS9H
Wk2Xl3tlf5VUvODGHftO9aXQspd7mig9VGxXAgMBAAECggEAYgfb+9u4C6n3M2r1
h5wWggJaynuGOjxSDJmmygu1TWG15xd3SkRv66Gp5v6Wz0OIIbaqcrr2SsFqDAlJ
kfh5gpvTBOXz1QMtTqaxLnWjD9bHMtbR6VodFfDbMQytiMr8TC+8jvwQI2Z4rPc+
V+zqZJjrsI/XFNQ0MnT4v/BkkFPnQBwjQ6I4jFgml+sQDnoHVaRxj0D2C9CxeNR+
UGoV8O4q5dj+pHdr4WpL8+lj+JguU8u9FD2ftXDNbMV7yppLDE8fLULzYkTdMff4
2at14mRSZPWyXyOvFUB3n0NXeQZ6GQ/u9KUpbYLUasGclSMDyBU8B5K8ZmA6E+U5
iFQCAQKBgQCwUdr7/ZVFp9v3p0npE2URES3afH8f1gCYSJPJVxvIIyugmdmvxTbi
L57uNGMW2SAbS1LD7Bko8Lg7ixniLUHYNGUXvV3xJ4yfYs9dgszoro3J6VfwyQiL
14zoiU+L0gxPiUkjkc5iGeZGcLnsKZtef7ZSajqseb+6Spn2+a2KFwKBgQCjbk6/
3k9vURmUYxTQ+nlp30u9Ux4nDOiqbr+wZGk1rLFCkBiKuV3EzbSdR13jd3J4zU26
5Q1eHO7UxgCCVN+Uq940m5ZoLLfaB0blnZ5T7C/JVptGDnvHQWwZpYrobvJOn8hk
Y45Cx6TRpUjP+v88yynoBjIjXtJsuAKEs2rXwQKBgQCEfZ3NyOlIJDB4Ue6KA5bo
Uj5gdTiZVZew1qimnJw1safw0GrxV4I8SIhmelsTKiikGd+hdWIaj3gA1i+m81cM
/aIRSqyyr4NNqVQk8krDdZb56Cu6tXWEj35Ephlt+XJiuciJv+CTet68g6xRt5oW
UmVQXJ6mR/44ICK9oTQ0FQKBgDBsTJT475xR8TxQJyjMLhE6ghx0diin4IAvFEjV
V375zgOk4kCoL6pSMMOxE5aEG7zDR9aqa02JURJmIFx9fUl1qv73Ypc3OLo3EcCS
9Qh0oKwNJzCYtgAR2E+5BvFbfhvhp+Rmww3yTVl4mZxWCC4hHCeNPmnQfJSN6OEc
r/mBAoGBAKeVovQaIeN4rT25zXPo29j4eI2WR1/lhNz1t2giYeETIEAXH81gtn6c
XSn+sHCodYTLtthOsVvJpQVtSon4t+z+9Xl5RRCqBq2aKyRZCxHRgFHyHQPLm068
40H2inoJysn6R8xiUapHiw6Jb8W4sHPoGmOLKpkcU09ZfMEhmCk9
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQBwkBfXoaXcl+/pw0q8Wg4q
M3nuREleXPQb9/pGhPCX9K0P+T82wTLKj5aweEQ85vg4HTDS3eYVPaYarN5gELNj
Spse/QsC2SHc47TQgUDz+EfzXFPkkX6S7W+nw/+8y54AqfnuIjEMWuhHjF4BzwgP
8ktb2xNqR3tEPB3gNrRPagLNCrqeuKUI0ne2s19ML/6I+O9BB9oq3ejkdNwf/ae6
B3XnPbn0LNjnt5ufTQrmr36aUE2WOfmxK0ikAOanFJZCdrJlxh4u5cLp1QFqiNZd
92J0wDR6Nwt2n2Wuu2kYotiAoS9HWk2Xl3tlf5VUvODGHftO9aXQspd7mig9VGxX
AgMBAAE=
-----END PUBLIC KEY-----
`)

const (
	Salt4Password = "djq.51zxiu.cn%!"
)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func EncryptPassword(str string) (string, error) {
	mid, err := RsaEncrypt([]byte(str))
	if err != nil {
		return str, errors.Wrap(err, "密码加密异常")
	}
	return base64.StdEncoding.EncodeToString(mid), nil
}

func DecryptPassword(str string) (string, error) {
	mid, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str, errors.Wrap(err, "密码加密异常")
	}
	mid2, err := RsaDecrypt(mid)
	if err != nil {
		return str, errors.Wrap(err, "密码加密异常")
	}
	return string(mid2), nil
}

func BuildPassword4DB(str string) string {
	s1 := str + Salt4Password
	s2 := []byte(s1)
	s3 := md5.Sum(s2)
	s4 := hex.EncodeToString(s3[:])
	return s4
	//return hex.EncodeToString(md5.Sum([]byte(str + Salt4Password))[:])
}

func GetPublicKey() []byte {
	return publicKey
}
