package shopeego


//=======================================================
// Object Raw Type - DiscountAddDiscount
//=======================================================
type DiscountAddDiscount struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
}
//=======================================================
// DiscountAddDiscountRequest
//=======================================================
type DiscountAddDiscountRequest struct {
    // discount_name is Title of the discount.
    DiscountName string `json:"discount_name"`
    // start_time is The time when discount activity start.The start time must be 1 hour later than current time.
    StartTime int `json:"start_time"`
    // end_time is <p>The time when discount activity end.The end time must be 1 hour later than start time,and the discount period must be less than 180 days.</p>
    EndTime int `json:"end_time"`
}
//=======================================================
// DiscountAddDiscountResponse
//=======================================================
type DiscountAddDiscountResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response DiscountAddDiscount `json:"response"`
}


//=======================================================
// Object Raw Type - ErrorList
//=======================================================
type ErrorList struct {
// item_id is  Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// model_id is Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
ModelID int `json:"model_id,omitempty"`
// fail_message is Indicate error details if one element hit error.
FailMessage string `json:"fail_message,omitempty"`
// fail_error is Indicate error type if one element hit error.
FailError string `json:"fail_error,omitempty"`
}


//=======================================================
// Object Raw Type - DiscountAddDiscountItem
//=======================================================
type DiscountAddDiscountItem struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// count is The number of items that add successfully.
Count int `json:"count,omitempty"`
// error_list is Indicate error details.
ErrorList []ErrorList `json:"error_list"`
}
//=======================================================
// DiscountAddDiscountItemRequest
//=======================================================
type DiscountAddDiscountItemRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // item_list is The items added in this discount promotion.
    ItemList []ItemList `json:"item_list"`
}
//=======================================================
// DiscountAddDiscountItemResponse
//=======================================================
type DiscountAddDiscountItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountAddDiscountItem `json:"response"`
}


//=======================================================
// Object Raw Type - DiscountDeleteDiscount
//=======================================================
type DiscountDeleteDiscount struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// modify_time is The time when discount has been deleted.
ModifyTime int `json:"modify_time,omitempty"`
}
//=======================================================
// DiscountDeleteDiscountRequest
//=======================================================
type DiscountDeleteDiscountRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
}
//=======================================================
// DiscountDeleteDiscountResponse
//=======================================================
type DiscountDeleteDiscountResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response DiscountDeleteDiscount `json:"response"`
}


//=======================================================
// Object Raw Type - DiscountDeleteDiscountItem
//=======================================================
type DiscountDeleteDiscountItem struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// error_list is Detail informations about error.
ErrorList []ErrorList `json:"error_list"`
}
//=======================================================
// DiscountDeleteDiscountItemRequest
//=======================================================
type DiscountDeleteDiscountItemRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // item_id is Shopee's unique identifier for an item.
    ItemID int `json:"item_id"`
    // model_id is Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
    ModelID int `json:"model_id,omitempty"`
}
//=======================================================
// DiscountDeleteDiscountItemResponse
//=======================================================
type DiscountDeleteDiscountItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountDeleteDiscountItem `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - DiscountGetDiscount
//=======================================================
type DiscountGetDiscount struct {
// status is The status of discount promotion
Status string `json:"status,omitempty"`
// discount_name is Title of the discount.
DiscountName string `json:"discount_name,omitempty"`
// item_list is The items selected in this discount.
ItemList []ItemList `json:"item_list"`
// start_time is The time when discount activity start.
StartTime int `json:"start_time,omitempty"`
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// end_time is The time when discount activity end.
EndTime int `json:"end_time,omitempty"`
// more is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
More bool `json:"more,omitempty"`
}
//=======================================================
// DiscountGetDiscountRequest
//=======================================================
type DiscountGetDiscountRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // page_no is Specifies the page number of data to return in the current call. Starting from 1. if data is more than one page, the page_no can be some entry to start next call.
    PageNo int `json:"page_no"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call), and the "page_no" to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
    PageSize int `json:"page_size"`
}
//=======================================================
// DiscountGetDiscountResponse
//=======================================================
type DiscountGetDiscountResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountGetDiscount `json:"response"`
}


//=======================================================
// Object Raw Type - DiscountList
//=======================================================
type DiscountList struct {
// status is The status of discount.
Status string `json:"status,omitempty"`
// discount_name is Title of the discount.
DiscountName string `json:"discount_name,omitempty"`
// start_time is The time when discount activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when discount activity end.
EndTime int `json:"end_time,omitempty"`
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// source is Source of the discount. 7: live stream, 1: admin, 0: others
Source int `json:"source,omitempty"`
}


//=======================================================
// Object Raw Type - DiscountGetDiscountList
//=======================================================
type DiscountGetDiscountList struct {
// discount_list is The discounts created in this shop.
DiscountList []DiscountList `json:"discount_list"`
// more is <p>This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.<br /></p>
More bool `json:"more,omitempty"`
}
//=======================================================
// DiscountGetDiscountListRequest
//=======================================================
type DiscountGetDiscountListRequest struct {
    // discount_status is The status filter for retriveing discount list. Available value: upcoming/ongoing/expired/all.
    DiscountStatus string `json:"discount_status"`
    // page_no is Specifies the starting entry of data to return in the current call. if data is more than one page, the offset can be some entry to start next call.
    PageNo int `json:"page_no"`
    // page_size is If many items are available to retrieve, you may need to call GetDiscountsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
    PageSize int `json:"page_size"`
    // update_time_from is The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
    UpdateTimeFrom int `json:"update_time_from,omitempty"`
    // update_time_to is The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The maximum date range that may be specified with the update_time_from and update_time_to fields is 30 days.
    UpdateTimeTo int `json:"update_time_to,omitempty"`
}
//=======================================================
// DiscountGetDiscountListResponse
//=======================================================
type DiscountGetDiscountListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountGetDiscountList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - DiscountUpdateDiscount
//=======================================================
type DiscountUpdateDiscount struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// modify_time is The time when discount is updated.
ModifyTime int `json:"modify_time,omitempty"`
}
//=======================================================
// DiscountUpdateDiscountRequest
//=======================================================
type DiscountUpdateDiscountRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // discount_name is Title of the discount.
    DiscountName string `json:"discount_name,omitempty"`
    // end_time is The time when discount activity end. The end time must be 1 hour later than start time.
    EndTime int `json:"end_time,omitempty"`
    // start_time is The time when discount activity start. The new start time must later than original start time.
    StartTime int `json:"start_time,omitempty"`
}
//=======================================================
// DiscountUpdateDiscountResponse
//=======================================================
type DiscountUpdateDiscountResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountUpdateDiscount `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - DiscountUpdateDiscountItem
//=======================================================
type DiscountUpdateDiscountItem struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// count is The number of items that modify successfully.
Count int `json:"count,omitempty"`
// error_list is Error list of this discount.
ErrorList []ErrorList `json:"error_list"`
}
//=======================================================
// DiscountUpdateDiscountItemRequest
//=======================================================
type DiscountUpdateDiscountItemRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // item_list is The items selected to this discount. You can update at most 50 items per call.
    ItemList []ItemList `json:"item_list"`
}
//=======================================================
// DiscountUpdateDiscountItemResponse
//=======================================================
type DiscountUpdateDiscountItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountUpdateDiscountItem `json:"response"`
}


//=======================================================
// Object Raw Type - DiscountEndDiscount
//=======================================================
type DiscountEndDiscount struct {
// discount_id is Shopee's unique identifier for a discount activity.
DiscountID int `json:"discount_id,omitempty"`
// modify_time is The time to track the modified time.
ModifyTime int `json:"modify_time,omitempty"`
}
//=======================================================
// DiscountEndDiscountRequest
//=======================================================
type DiscountEndDiscountRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
}
//=======================================================
// DiscountEndDiscountResponse
//=======================================================
type DiscountEndDiscountResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response DiscountEndDiscount `json:"response"`
}