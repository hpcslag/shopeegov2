package shopeego
//=======================================================
// MerchantGetMerchantInfoRequest
//=======================================================
type MerchantGetMerchantInfoRequest struct {
    V2RequestAuthenticationParams
    

}
//=======================================================
// MerchantGetMerchantInfoResponse
//=======================================================
type MerchantGetMerchantInfoResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// MerchantGetShopListByMerchantRequest
//=======================================================
type MerchantGetShopListByMerchantRequest struct {
    V2RequestAuthenticationParams
    

    // page_no is Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.	
    PageNo int `json:"page_no"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.No more than 500.	
    PageSize int `json:"page_size"`
}
//=======================================================
// MerchantGetShopListByMerchantResponse
//=======================================================
type MerchantGetShopListByMerchantResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}