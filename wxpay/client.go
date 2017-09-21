package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"mimi/djq/config"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const bodyType = "application/xml; charset=utf-8"

// API客户端
type Client struct {
	stdClient *http.Client
	tlsClient *http.Client

	AppId     string
	MchId     string
	ApiKey    string
}

// 实例化API客户端
func NewDefaultClient() *Client {
	appId := config.Get("wxpay_appid") // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key") // 微信支付商户平台API密钥
	return &Client{
		stdClient: &http.Client{},
		AppId:     appId,
		MchId:     mchId,
		ApiKey:    apiKey,
	}
}

// 实例化API客户端
func NewClient(appId, mchId, apiKey string) *Client {
	return &Client{
		stdClient: &http.Client{},
		AppId:     appId,
		MchId:     mchId,
		ApiKey:    apiKey,
	}
}

// 设置请求超时时间
func (c *Client) SetTimeout(d time.Duration) {
	c.stdClient.Timeout = d
	if c.tlsClient != nil {
		c.tlsClient.Timeout = d
	}
}

// 附着商户证书
func (c *Client) WithCert(certFile, keyFile, rootcaFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(rootcaFile)
	if err != nil {
		return err
	}
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(data)
	if !ok {
		return errors.New("failed to parse root certificate")
	}
	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	trans := &http.Transport{
		TLSClientConfig: conf,
	}
	c.tlsClient = &http.Client{
		Transport: trans,
	}
	return nil
}

// 发送请求
func (c *Client) Post(url string, params Params, tls bool) (Params, error) {
	fmt.Println(url, params)
	var httpc *http.Client
	if tls {
		if c.tlsClient == nil {
			return nil, errors.New("tls client is not initialized")
		}
		httpc = c.tlsClient
	} else {
		httpc = c.stdClient
	}
	resp, err := httpc.Post(url, bodyType, c.Encode(params))
	if err != nil {
		return nil, err
	}
	return c.Decode(resp.Body), nil
}

// XML解码
func (c *Client) Decode(r io.Reader) Params {
	var (
		d      *xml.Decoder
		start  *xml.StartElement
		params Params
	)
	d = xml.NewDecoder(r)
	params = make(Params)
	for {
		tok, err := d.Token()
		if err != nil {
			break
		}
		switch t := tok.(type) {
		case xml.StartElement:
			start = &t
		case xml.CharData:
			if t = bytes.TrimSpace(t); len(t) > 0 {
				params.SetString(start.Name.Local, string(t))
			}
		}
	}
	return params
}

// XML编码
func (c *Client) Encode(params Params) io.Reader {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)
	return &buf
}

// 验证签名
func (c *Client) CheckSign(params Params) bool {
	return strings.TrimSpace(params.GetString("sign")) == strings.TrimSpace(c.Sign(params))
}

// 生成签名
func (c *Client) Sign(params Params) string {
	var keys = make([]string, 0, len(params))
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	buf.WriteString(`key=`)
	buf.WriteString(c.ApiKey)

	sum := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(sum[:])

	return strings.ToUpper(str)
}

// 解密退款结果信息
func (c *Client) Aes256EcbDecrypt(str string) (Params, error) {
	sum := md5.Sum([]byte(c.ApiKey))
	key := hex.EncodeToString(sum[:])
	key = strings.ToLower(key)
	bs, err := base64.StdEncoding.DecodeString(str)
	bb, err := myDecode(bs, key)
	if err != nil {
		return nil, err
	}
	return c.Decode(bytes.NewReader(bb)), nil
}

func myEncode(txt, key string) (dest []byte, err error) {
	dest = make([]byte, (len(txt) / len(key) + 1) * len(key))

	aesCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encrypter := cipher.NewECBEncrypter(aesCipher)

	encrypter.CryptBlocks(dest, []byte(txt))
	return
}

func myDecode(bs []byte, key string) (txt []byte, err error) {

	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return
	}
	blockModel := cipher.NewECBDecrypter(block)
	txt = make([]byte, len(bs))
	blockModel.CryptBlocks(txt, bs)
	return
}