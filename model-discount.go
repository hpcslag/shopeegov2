package shopeego
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
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // error is Error code
    Error string `json:"error,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// DiscountAddDiscountItemRequest
//=======================================================
type DiscountAddDiscountItemRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // item_list is The items added in this discount promotion.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// DiscountAddDiscountItemResponse
//=======================================================
type DiscountAddDiscountItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // response is 
    Response Response `json:"response"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// DiscountUpdateDiscountItemRequest
//=======================================================
type DiscountUpdateDiscountItemRequest struct {
    // discount_id is Shopee's unique identifier for a discount activity.
    DiscountID int `json:"discount_id"`
    // item_list is The items selected to this discount. You can update at most 50 items per call.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// DiscountUpdateDiscountItemResponse
//=======================================================
type DiscountUpdateDiscountItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}