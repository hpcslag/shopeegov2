package shopeego
//=======================================================
// MerchantGetMerchantInfoRequest
//=======================================================
type MerchantGetMerchantInfoRequest struct {
}
//=======================================================
// MerchantGetMerchantInfoResponse
//=======================================================
type MerchantGetMerchantInfoResponse struct {
    // merchant_name is Name of the merchant.
    MerchantName string `json:"merchant_name,omitempty"`
    // is_cnsc is Use this filed to indicate whether the merchant is upgraded to CNSC.
    IsCnsc bool `json:"is_cnsc,omitempty"`
    // auth_time is The timestamp when the merchant was authorized to the partner.
    AuthTime int `json:"auth_time,omitempty"`
    // expire_time is Use this field to indicate the expiration date for merchant authorization.
    ExpireTime int `json:"expire_time,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // merchant_currency is <p>The three-digit code representing the currency unit used for the item in this merchant. If currency haven't been setting in CNSC/KRSC, it will be empty.China merchant support CNY and USD currently.Korea merchant support KRW and USD currently.</p>
    MerchantCurrency string `json:"merchant_currency,omitempty"`
    // merchant_region is <p>Region of the merchant.<br /></p>
    MerchantRegion string `json:"merchant_region,omitempty"`
    // is_upgraded_cbsc is <p>Use this filed to indicate whether this merchant is upgraded to cbsc.<br /></p>
    IsUpgradedCbsc bool `json:"is_upgraded_cbsc,omitempty"`
}
//=======================================================
// MerchantGetShopListByMerchantRequest
//=======================================================
type MerchantGetShopListByMerchantRequest struct {
    // page_no is Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.	
    PageNo int `json:"page_no"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.No more than 500.	
    PageSize int `json:"page_size"`
}
//=======================================================
// MerchantGetShopListByMerchantResponse
//=======================================================
type MerchantGetShopListByMerchantResponse struct {
    // is_cnsc is Use this filed to indicate whether the merchant is upgraded to CNSC.
    IsCnsc bool `json:"is_cnsc,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // shop_list is list of shop authorized to the partner and bound to the merchant. 	
    ShopList []interface{} `json:"shop_list,omitempty"`
    // more is This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
    More bool `json:"more,omitempty"`
}