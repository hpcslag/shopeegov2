package shopeego


//=======================================================
// Object Raw Type - VoucherAddVoucher
//=======================================================
type VoucherAddVoucher struct {
// voucher_id is The unique identifier for the created voucher.
VoucherID int `json:"voucher_id,omitempty"`
}
//=======================================================
// VoucherAddVoucherRequest
//=======================================================
type VoucherAddVoucherRequest struct {
    // voucher_name is The name of the voucher.
    VoucherName string `json:"voucher_name"`
    // voucher_code is The code of the voucher.
    VoucherCode string `json:"voucher_code"`
    // start_time is <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
    StartTime int `json:"start_time"`
    // end_time is The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
    EndTime int `json:"end_time"`
    // voucher_type is The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
    VoucherType int `json:"voucher_type"`
    // reward_type is The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
    RewardType int `json:"reward_type"`
    // usage_quantity is The number of times for this particular voucher could be used.
    UsageQuantity int `json:"usage_quantity"`
    // min_basket_price is The minimum spend required for using this voucher. 
    MinBasketPrice float64 `json:"min_basket_price"`
    // discount_amount is The discount amount set for this particular voucher. Only fill in when you are creating a fix amount voucher.
    DiscountAmount float64 `json:"discount_amount,omitempty"`
    // percentage is The discount percentage set for this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.
    Percentage int `json:"percentage,omitempty"`
    // max_price is The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are creating a discount percentage voucher or coins cashback voucher.
    MaxPrice float64 `json:"max_price,omitempty"`
    // display_channel_list is The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
    DisplayChannelList []int `json:"display_channel_list,omitempty"`
    // item_id_list is The list of items which is applicable for the voucher. Only fill in when you are creating a product type of voucher.
    ItemIdList []int `json:"item_id_list,omitempty"`
    // display_start_time is The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
    DisplayStartTime int `json:"display_start_time,omitempty"`
}
//=======================================================
// VoucherAddVoucherResponse
//=======================================================
type VoucherAddVoucherResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detailed informations you are querying.
    Response VoucherAddVoucher `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - VoucherDeleteVoucher
//=======================================================
type VoucherDeleteVoucher struct {
// voucher_id is The unique identifier for the voucher it is being deleted.
VoucherID int `json:"voucher_id,omitempty"`
}
//=======================================================
// VoucherDeleteVoucherRequest
//=======================================================
type VoucherDeleteVoucherRequest struct {
    // voucher_id is The unique identifier for the voucher you want to delete.
    VoucherID int `json:"voucher_id"`
}
//=======================================================
// VoucherDeleteVoucherResponse
//=======================================================
type VoucherDeleteVoucherResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response VoucherDeleteVoucher `json:"response"`
}


//=======================================================
// Object Raw Type - VoucherEndVoucher
//=======================================================
type VoucherEndVoucher struct {
// voucher_id is The unique identifier for the voucher it is being ended.
VoucherID int `json:"voucher_id,omitempty"`
}
//=======================================================
// VoucherEndVoucherRequest
//=======================================================
type VoucherEndVoucherRequest struct {
    // voucher_id is The unique identifier for the voucher you want to end now.
    VoucherID int `json:"voucher_id"`
}
//=======================================================
// VoucherEndVoucherResponse
//=======================================================
type VoucherEndVoucherResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detailed informations you are querying.
    Response VoucherEndVoucher `json:"response"`
}


//=======================================================
// Object Raw Type - VoucherUpdateVoucher
//=======================================================
type VoucherUpdateVoucher struct {
// voucher_id is The unique identifier of the voucher which is being updated.
VoucherID int `json:"voucher_id,omitempty"`
}
//=======================================================
// VoucherUpdateVoucherRequest
//=======================================================
type VoucherUpdateVoucherRequest struct {
    // voucher_id is The unique identifier of the voucher which is going to be updated.
    VoucherID int `json:"voucher_id"`
    // voucher_name is The name of the voucher
    VoucherName string `json:"voucher_name,omitempty"`
    // start_time is <p>The timing from when the voucher is valid; so buyer is allowed to claim and to use. Field can only be updated if voucher has not started.</p>
    StartTime int `json:"start_time,omitempty"`
    // end_time is The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
    EndTime int `json:"end_time,omitempty"`
    // usage_quantity is The number of times for this particular voucher could be used.
    UsageQuantity int `json:"usage_quantity,omitempty"`
    // min_basket_price is The minimum spend required for using this voucher. 
    MinBasketPrice float64 `json:"min_basket_price,omitempty"`
    // discount_amount is The discount amount set for this particular voucher. Only fill in when you are updating a fix amount voucher.
    DiscountAmount float64 `json:"discount_amount,omitempty"`
    // percentage is The discount percentage set for this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
    Percentage int `json:"percentage,omitempty"`
    // max_price is The max amount of discount/value a user can enjoy by using this particular voucher. Only fill in when you are updating a discount percentage voucher or coins cashback voucher.
    MaxPrice float64 `json:"max_price,omitempty"`
    // display_channel_list is The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
    DisplayChannelList []int `json:"display_channel_list,omitempty"`
    // item_id_list is The list of items which is applicable for the voucher. Only fill in when you are updating a product type of voucher.
    ItemIdList []int `json:"item_id_list,omitempty"`
    // display_start_time is The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
    DisplayStartTime int `json:"display_start_time,omitempty"`
}
//=======================================================
// VoucherUpdateVoucherResponse
//=======================================================
type VoucherUpdateVoucherResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detailed informations you are querying.
    Response VoucherUpdateVoucher `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - VoucherGetVoucher
//=======================================================
type VoucherGetVoucher struct {
// voucher_id is The unique identifier of the voucher whose details are returned.
VoucherID int `json:"voucher_id,omitempty"`
// voucher_code is Voucher Code
VoucherCode string `json:"voucher_code,omitempty"`
// voucher_name is Voucher Name
VoucherName string `json:"voucher_name,omitempty"`
// voucher_type is The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
VoucherType int `json:"voucher_type,omitempty"`
// reward_type is The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
RewardType int `json:"reward_type,omitempty"`
// usage_quantity is The number of times for this particular voucher could be used.
UsageQuantity int `json:"usage_quantity,omitempty"`
// current_usage is Up till now, how many times has this particular voucher already been used.
CurrentUsage int `json:"current_usage,omitempty"`
// start_time is The timing from when the voucher is valid; so buyer is allowed to claim and to use. 
StartTime int `json:"start_time,omitempty"`
// end_time is The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
EndTime int `json:"end_time,omitempty"`
// is_admin is If the voucher is created by Shopee or not.
IsAdmin bool `json:"is_admin,omitempty"`
// voucher_purpose is The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership
VoucherPurpose int `json:"voucher_purpose,omitempty"`
// display_channel_list is The FE channel where the voucher will be displayed. The available values are: 1: display_all, 2: order page, 3: feed, 4: live streaming,   [] (empty - which is hidden).
DisplayChannelList []int `json:"display_channel_list,omitempty"`
// min_basket_price is The minimum spend required for using this voucher. 
MinBasketPrice float64 `json:"min_basket_price,omitempty"`
// percentage is The discount percentage set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
Percentage int `json:"percentage,omitempty"`
// max_price is The max amount of discount/value a user can enjoy by using this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
MaxPrice float64 `json:"max_price,omitempty"`
// discount_amount is The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
DiscountAmount float64 `json:"discount_amount,omitempty"`
// cmt_voucher_status is The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
CmtVoucherStatus int `json:"cmt_voucher_status,omitempty"`
// item_id_list is The list of items which is applicable for the voucher. Only return a value when it is a product type of voucher.
ItemIdList []int `json:"item_id_list,omitempty"`
// display_start_time is The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
DisplayStartTime int `json:"display_start_time,omitempty"`
}
//=======================================================
// VoucherGetVoucherRequest
//=======================================================
type VoucherGetVoucherRequest struct {
    // voucher_id is The unique identifier of a voucher used to query the voucher details.
    VoucherID int `json:"voucher_id"`
}
//=======================================================
// VoucherGetVoucherResponse
//=======================================================
type VoucherGetVoucherResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detailed informations you are querying.
    Response VoucherGetVoucher `json:"response"`
}


//=======================================================
// Object Raw Type - VoucherList
//=======================================================
type VoucherList struct {
// voucher_id is The unique identifier for a voucher.
VoucherID int `json:"voucher_id,omitempty"`
// voucher_code is Voucher Code
VoucherCode string `json:"voucher_code,omitempty"`
// voucher_name is Voucher Name
VoucherName string `json:"voucher_name,omitempty"`
// voucher_type is The type of the voucher. The available values are: 1: shop voucher, 2: product voucher.
VoucherType int `json:"voucher_type,omitempty"`
// reward_type is The reward type of the voucher. The available values are: 1: fix_amount voucher, 2: discount_percentage voucher, 3: coin_cashback voucher.
RewardType int `json:"reward_type,omitempty"`
// usage_quantity is The number of times for this particular voucher could be used.
UsageQuantity int `json:"usage_quantity,omitempty"`
// current_usage is Up till now, how many times has this particular voucher already been used.
CurrentUsage int `json:"current_usage,omitempty"`
// start_time is The timing from when the voucher is valid; so buyer is allowed to claim and to use. 
StartTime int `json:"start_time,omitempty"`
// end_time is The timing until when the voucher is still valid. Any time after this end_time, buyer is not allowed to claim or to use. 
EndTime int `json:"end_time,omitempty"`
// is_admin is If the voucher is created by Shopee or not.
IsAdmin bool `json:"is_admin,omitempty"`
// voucher_purpose is The use case for the voucher. The available values are: 0: normal; 1: welcome, 2: referral; 3: shop_follow; 4:shop_game, 5: free_gift, 6: membership
VoucherPurpose int `json:"voucher_purpose,omitempty"`
// discount_amount is The discount amount set for this particular voucher. Only when it is a fix amount voucher, api will return a value.
DiscountAmount float64 `json:"discount_amount,omitempty"`
// percentage is The discount percentage set for this particular voucher. Only when it is a discount percentage voucher or coins cashback voucher, api will return a value.
Percentage int `json:"percentage,omitempty"`
// cmt_voucher_status is The voucher status in CMT. The available values are: 1:review, 2: approved, 3:reject. Only when this voucher is attending CMT campaign and not being rejected, api will return a value.
CmtVoucherStatus int `json:"cmt_voucher_status,omitempty"`
// display_start_time is The timing of when voucher is displayed on shop pages; so buyer is allowed to claim.
DisplayStartTime int `json:"display_start_time,omitempty"`
}


//=======================================================
// Object Raw Type - VoucherGetVoucherList
//=======================================================
type VoucherGetVoucherList struct {
// more is This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.
More bool `json:"more,omitempty"`
// voucher_list is The list of voucher.
VoucherList VoucherList `json:"voucher_list"`
}
//=======================================================
// VoucherGetVoucherListRequest
//=======================================================
type VoucherGetVoucherListRequest struct {
    // page_no is Specifies the page number of data to return in the current call. Default to be 1 and allowed input is from 1 - 5000.
    PageNo int `json:"page_no,omitempty"`
    // page_size is Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.
    PageSize int `json:"page_size,omitempty"`
    // status is The status filter for retrieving voucher list. Available value: upcoming/ongoing/expired/all.
    Status string `json:"status"`
}
//=======================================================
// VoucherGetVoucherListResponse
//=======================================================
type VoucherGetVoucherListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detailed informations you are querying.
    Response VoucherGetVoucherList `json:"response"`
}