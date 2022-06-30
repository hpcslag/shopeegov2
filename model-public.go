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
    // 通用的 Response 回傳參數
    V2UnityResponse

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
    // 通用的 Response 回傳參數
    V2UnityResponse

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
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - PublicGetRefreshTokenByUpgradeCode
//=======================================================
type PublicGetRefreshTokenByUpgradeCode struct {
// success_shop_id_list is <p>The list of shop id which get the refresh_token successfully.<br /></p>
SuccessShopIdList []int `json:"success_shop_id_list,omitempty"`
// refresh_token is <p>Use refresh_token to obtain new access_token. Valid for each shop_id and merchant_id respectively one-time use, expires in 30 days.<br /></p>
RefreshToken string `json:"refresh_token,omitempty"`
// failure_list is <p>The failure information list<br /></p>
FailureList FailureList `json:"failure_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is <p>Detail informations you are querying.<br /></p>
    Response PublicGetRefreshTokenByUpgradeCode `json:"response"`
}