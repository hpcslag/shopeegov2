package shopeego
//=======================================================
// PublicGetShopsByPartnerRequest
//=======================================================
type PublicGetShopsByPartnerRequest struct {
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
    PageSize int `json:"page_size,omitempty"`
    // page_no is Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
    PageNo int `json:"page_no,omitempty"`
}
//=======================================================
// PublicGetShopsByPartnerResponse
//=======================================================
type PublicGetShopsByPartnerResponse struct {
    // authed_shop_list is A list of shops authorized to the partner.
    AuthedShopList []interface{} `json:"authed_shop_list,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // more is This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
    More bool `json:"more,omitempty"`
}
//=======================================================
// PublicGetMerchantsByPartnerRequest
//=======================================================
type PublicGetMerchantsByPartnerRequest struct {
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
    PageSize int `json:"page_size,omitempty"`
    // page_no is Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
    PageNo int `json:"page_no,omitempty"`
}
//=======================================================
// PublicGetMerchantsByPartnerResponse
//=======================================================
type PublicGetMerchantsByPartnerResponse struct {
    // authed_merchant_list is A list of merchants authorized to the partner.
    AuthedMerchantList []interface{} `json:"authed_merchant_list,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // more is This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of datas.
    More bool `json:"more,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// PublicGetTokenByResendCodeRequest
//=======================================================
type PublicGetTokenByResendCodeRequest struct {
    // resend_code is the code in redirect url after you resend code in shop authorization management page. valid for one-time use, expires in 10minutes.
    ResendCode string `json:"resend_code"`
}
//=======================================================
// PublicGetTokenByResendCodeResponse
//=======================================================
type PublicGetTokenByResendCodeResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. 
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error.
    Message string `json:"message,omitempty"`
    // shop_id_list is Return when resend code in shop module
    ShopIdList []int `json:"shop_id_list,omitempty"`
    // merchant_id_list is Return when resend code in merchant module
    MerchantIdList []int `json:"merchant_id_list,omitempty"`
    // refresh_token is Use refresh_token to obtain new access_token. Valid for each shop_id and merchant_id respectively one-time use, expires in 30 days.
    RefreshToken string `json:"refresh_token,omitempty"`
    // access_token is The token for API access, using to identify your permission to the api. Valid for multiple use and expires in 4 hours.
    AccessToken string `json:"access_token,omitempty"`
    // expire_in is Access_token expiration time, unit is second.
    ExpireIn int `json:"expire_in,omitempty"`
}
//=======================================================
// PublicGetRefreshTokenByUpgradeCodeRequest
//=======================================================
type PublicGetRefreshTokenByUpgradeCodeRequest struct {
    // upgrade_code is <p>All the app who have the access to call open api V1.0 can have an upgrade_code after they apply in app details page.Each authorized shop can use this upgrade_code 3 times to get a V2 refresh token and start to call open api V2.0.<br /></p>
    UpgradeCode string `json:"upgrade_code"`
    // shop_id_list is <p>The list of shop id which you want to get the V2 refresh token. The limit is between 1 and 100.<br /></p>
    ShopIdList []int `json:"shop_id_list"`
}
//=======================================================
// PublicGetRefreshTokenByUpgradeCodeResponse
//=======================================================
type PublicGetRefreshTokenByUpgradeCodeResponse struct {
    // request_id is <p>The identifier for an API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // response is <p>Detail informations you are querying.<br /></p>
    Response Response `json:"response"`
}