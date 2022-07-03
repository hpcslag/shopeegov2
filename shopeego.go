package shopeego

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// There are errors in the input parameters
	ErrParams = errors.New("shopeego: there are errors in the input parameters")
	// The request is not authenticated. Ex: signature is wrong
	ErrAuth = errors.New("shopeego: the request is not authenticated. Ex: signature is wrong")
	// An error has occurred
	ErrServer = errors.New("shopeego: an error has occurred")
	// Too many request.Exceed 1000 request per min will be banned for an hour
	ErrTooManyRequest = errors.New("shopeego: too many request.Exceed 1000 request per min will be banned for an hour")
	// Not support action
	ErrNotSupport = errors.New("shopeego: not support action")
	// An inner error has occurred
	ErrInnerError = errors.New("shopeego: an inner error has occurred")
	// No Permission
	ErrPermission = errors.New("shopeego: no permission")
	// Not exist
	ErrNotExists = errors.New("shopeego: not exists")
)

const (
	// ClientVersionV1 for Open API V1, see: https://open.shopee.com/documents?module=63&type=2&id=53.
	ClientVersionV1 ClientVersion = iota
	// ClientVersionV2 for Open API V2, see: https://open.shopee.com/documents?module=63&type=2&id=56.
	ClientVersionV2
)

// ClientVersion is the API version of the client should be using.
type ClientVersion int

const (
	urlTestSandbox string = "https://partner.uat.shopeemobile.com/"
	urlStandard    string = "https://partner.shopeemobile.com/"
)

type ClientOptions struct {
	Secret string
	// 非必要
	//PartnerID int64
	// 非必要
	//ShopID int64
	//
	IsSandbox bool
	//
	Version ClientVersion
}

// CONSTS

func NewClient(opts *ClientOptions) Client {
	return &ShopeeClient{
		Secret:    opts.Secret,
		IsSandbox: opts.IsSandbox,
		Version:   opts.Version,
	}
}

// Client 定義了一個蝦皮的客戶端該有什麼功能。
type Client interface {
	SetAccessToken(t string) *ShopeeClient

	//=======================================================
	// Shop
	//=======================================================

	// GetShopInfo Use this call to get information of shop
	GetShopInfo(*GetShopInfoRequest) (*GetShopInfoResponse, error)
	// UpdateShopInfo Use this call to update information of shop
	UpdateShopInfo(*UpdateShopInfoRequest) (*UpdateShopInfoResponse, error)
	// Performance Shop performance includes the indexes from "My Performance" of Seller Center.
	Performance(*PerformanceRequest) (*PerformanceResponse, error)
	// SetShopInstallmentStatus Only for TW whitelisted shop.Use this API to set the installment status of shop.
	SetShopInstallmentStatus(*SetShopInstallmentStatusRequest) (*SetShopInstallmentStatusResponse, error)
	//
	AuthPartner(*AuthPartnerRequest) string

	//=======================================================
	// Auth (V2)
	//=======================================================

	// GetAccessToken Use this API and the code to obtain the access_token and refresh_token.
	GetAccessToken(*GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	// RefreshAccessToken Use this API to refresh the access_token after it expires.
	RefreshAccessToken(*RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)

	// import generated files
	V2I
}

// ShopeeClient represents a client to Shopee
type ShopeeClient struct {
	Secret      string
	accessToken string
	IsSandbox   bool
	Version     ClientVersion
}

// ResponseError defines a error response
type ResponseError struct {
	RequestID string `json:"request_id,omitempty"`
	Msg       string `json:"msg,omitempty"`
	ErrorType string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (e ResponseError) Error() string {
	// 如果 `msg` 是空的，就試試看 V2 的 `message`
	msg := e.Msg
	if msg == "" {
		msg = e.Message
	}
	return fmt.Sprintf("shopeego: %s - %s", e.ErrorType, msg)
}

//
func (s *ShopeeClient) getPath(method string) string {
	var host string
	if s.IsSandbox {
		host = urlTestSandbox
	} else {
		host = urlStandard
	}
	// 依照不同版本處理 API 前輟。
	switch s.Version {
	case ClientVersionV1:
		return fmt.Sprintf("%sapi/v1/%s", host, availablePaths[method])
	default:
		return fmt.Sprintf("%s%s", host, string(availablePaths[method])[1:]) // split first slash /api...
	}
}

// signV1 會簽署 V1 API。
func (s *ShopeeClient) signV1(url string, body []byte) string {
	h := hmac.New(sha256.New, []byte(s.Secret))
	io.WriteString(h, fmt.Sprintf("%s|%s", url, string(body)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// signV2 會簽署 V2 API。
func (s *ShopeeClient) signV2(url string, b []byte) string {
	h := hmac.New(sha256.New, []byte(s.Secret))
	p := s.getBodyPart(b)
	if s.IsSandbox {
		url = "/api/" + strings.TrimLeft(url, urlTestSandbox)
	} else {
		url = "/api/" + strings.TrimLeft(url, urlStandard)
	}
	// (順序是重要的) follow the orders https://open.shopee.com/documents/v2/%5B%E4%B8%AD%E6%96%87%E7%89%88%5D%20OpenAPI%202.0%20Overview?module=87&type=2
	io.WriteString(h, fmt.Sprintf("%d%s%d%s%d", p.partnerID, url, p.timestamp, s.accessToken, p.shopID))
	return fmt.Sprintf("%x", h.Sum(nil))
}

type bodyPart struct {
	partnerID int64
	timestamp int
	shopID    int64
}

//
func (s *ShopeeClient) getBodyPart(b []byte) bodyPart {
	var p bodyPart
	var t map[string]interface{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return p
	}
	if v, ok := t["partner_id"].(float64); ok {
		p.partnerID = int64(v)
	}
	if v, ok := t["timestamp"].(float64); ok {
		p.timestamp = int(v)
	}
	if v, ok := t["shopid"].(float64); ok {
		p.shopID = int64(v)
	}
	if v, ok := t["shop_id"].(float64); ok {
		p.shopID = int64(v)
	}
	return p
}

func (s *ShopeeClient) SetAccessToken(t string) *ShopeeClient {
	s.accessToken = t
	return s
}

//
func (s *ShopeeClient) makeV2Query(url string, b []byte) string {
	p := s.getBodyPart(b)
	q := fmt.Sprintf("?partner_id=%d&shop_id=%d&timestamp=%d&sign=%s", p.partnerID, p.shopID, p.timestamp, s.signV2(url, b))
	if s.accessToken != "" {
		q += fmt.Sprintf("&access_token=%s", s.accessToken)
	}
	return q
}

//
func (s *ShopeeClient) get(method string, in interface{}) ([]byte, error) {
	query, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	// 把 json 轉 params
	var queryMap map[string]interface{}
	err = json.Unmarshal(query, &queryMap)
	if err != nil {
		return nil, err
	}

	// pureHostPath for get request
	pureHostPath := s.getPath(method)

	url := pureHostPath + "?"

	if len(queryMap) > 0 {
		for key, value := range queryMap {
			switch v := value.(type) {
			case float64:
				url += fmt.Sprintf("%+v=%+v&", key, int(v))
			case []interface{}:
				if len(v) > 0 {
					combine := ""
					for index, item := range v {
						switch itemValue := item.(type) {
						case string:
							combine += fmt.Sprintf("%+v", itemValue)
						case float64:
							combine += fmt.Sprintf("%+v", int(itemValue))
						default:
							combine += fmt.Sprintf("%+v", itemValue)
						}
						if index < len(v)-1 {
							combine += ","
						}
					}
					combine += ""
					url += fmt.Sprintf("%+v=%+v&", key, combine)
				}
			default:
				url += fmt.Sprintf("%+v=%+v&", key, v)
			}
		}
	}

	// sign for url
	var sign string
	switch s.Version {
	// 如果是 V1 就在 Header 安插 Sign。
	case ClientVersionV1:
		sign = s.signV1(pureHostPath, query)
	// 如果是 V2 的 API，就在 Body 中自動安插 Sign。
	case ClientVersionV2:
		sign = s.signV2(pureHostPath, query)
	}
	// add sign to url
	url += fmt.Sprintf("sign=%s", sign)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	//Do request by native lib
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errResp ResponseError
	_ = json.Unmarshal(body, &errResp)
	if errResp.ErrorType != "" {
		return nil, errResp
	}

	return s.patchFloat(body), nil
}

//
func (s *ShopeeClient) post(method string, in interface{}) ([]byte, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	url := s.getPath(method)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")

	switch s.Version {
	// 如果是 V1 就在 Header 安插 Sign。
	case ClientVersionV1:
		req.Header.Add("Authorization", s.signV1(url, body))
	// 如果是 V2 的 API，就在 Body 中自動安插 Sign。
	case ClientVersionV2:
		req, err = http.NewRequest("POST", fmt.Sprintf("%s%s", url, s.makeV2Query(url, body)), bytes.NewReader(body))
		req.Header.Add("Authorization", s.signV2(url, body))
	}

	//Do request by native lib
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errResp ResponseError
	_ = json.Unmarshal(body, &errResp)
	if errResp.ErrorType != "" {
		return nil, errResp
	}

	return s.patchFloat(body), nil
}

// patchFloat 會修正無法將 JSON 的字串值轉換成 Float64 型態的錯誤，這主要是因為 Shopee 會在某些時候以 `""`（空字串）當作是 Float64 的零值，
// 但對於 Golang 來說，Float64 的零值必須是 `"0"`，所以這個函式會將 Raw JSON 裡的所有相關（請參閱 `replaces.go` 關鍵詞表） `""` 轉換為 `"0"` 以便正確解析。
//
// 這個函式修正了 `json: invalid use of ,string struct tag, trying to unmarshal unquoted value into float64` 錯誤。
// 相關 Issue： https://github.com/teacat/shopeego/issues/6
func (s *ShopeeClient) patchFloat(body []byte) []byte {
	replaceConcat := strings.Join(replaces, "|")
	for _, v := range replaces {
		body = []byte(strings.ReplaceAll(string(body), fmt.Sprintf(`"%s":""`, v), fmt.Sprintf(`"%s":"0"`, v)))
	}

	var r = regexp.MustCompile(fmt.Sprintf(`"(%s)":(.*?)(,|})`, replaceConcat))
	return []byte(r.ReplaceAllString(string(body), `"$1":"$2"$3`))
}

//=======================================================
// Shop
//=======================================================

// GetShopInfo Use this call to get information of shop
func (s *ShopeeClient) GetShopInfo(req *GetShopInfoRequest) (resp *GetShopInfoResponse, err error) {
	b, err := s.post("GetShopInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateShopInfo Use this call to update information of shop
func (s *ShopeeClient) UpdateShopInfo(req *UpdateShopInfoRequest) (resp *UpdateShopInfoResponse, err error) {
	b, err := s.post("UpdateShopInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// Performance Shop performance includes the indexes from "My Performance" of Seller Center.
func (s *ShopeeClient) Performance(req *PerformanceRequest) (resp *PerformanceResponse, err error) {
	b, err := s.post("Performance", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SetShopInstallmentStatus Only for TW whitelisted shop.Use this API to set the installment status of shop.
func (s *ShopeeClient) SetShopInstallmentStatus(req *SetShopInstallmentStatusRequest) (resp *SetShopInstallmentStatusResponse, err error) {
	b, err := s.post("SetShopInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AuthPartner for V2.
func (s *ShopeeClient) AuthPartner(req *AuthPartnerRequest) string {
	url := s.getPath("AuthPartner")
	switch s.Version {
	// V1
	case ClientVersionV1:
		h := sha256.Sum256([]byte(fmt.Sprintf("%s%s", s.Secret, req.Redirect)))
		token := fmt.Sprintf("%x", h[:])
		return fmt.Sprintf("%s?id=%d&token=%s&redirect=%s", url, req.PartnerID, token, req.Redirect)
	// V2
	default:
		h := hmac.New(sha256.New, []byte(s.Secret))
		io.WriteString(h, fmt.Sprintf("%d%s%d", req.PartnerID, "/api/v2/shop/auth_partner", req.Timestamp))
		token := fmt.Sprintf("%x", h.Sum(nil))
		return fmt.Sprintf("%s?partner_id=%d&redirect=%s&sign=%s&timestamp=%d", url, req.PartnerID, req.Redirect, token, req.Timestamp)
	}
}

//=======================================================
// Auth (V2)
//=======================================================

// GetAccessToken Use this API and the code to obtain the access_token and refresh_token.
func (s *ShopeeClient) GetAccessToken(req *GetAccessTokenRequest) (resp *GetAccessTokenResponse, err error) {
	b, err := s.post("GetAccessToken", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// RefreshAccessToken Use this API to refresh the access_token after it expires.
func (s *ShopeeClient) RefreshAccessToken(req *RefreshAccessTokenRequest) (resp *RefreshAccessTokenResponse, err error) {
	b, err := s.post("RefreshAccessToken", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}
