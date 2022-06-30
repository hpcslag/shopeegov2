package shopeego
//=======================================================
// BundleDealAddBundleDealRequest
//=======================================================
type BundleDealAddBundleDealRequest struct {
    // rule_type is The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
    RuleType int `json:"rule_type"`
    // discount_value is The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
    DiscountValue float64 `json:"discount_value"`
    // fix_price is The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
    FixPrice float64 `json:"fix_price"`
    // discount_percentage is The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
    DiscountPercentage int `json:"discount_percentage"`
    // min_amount is The quantity of items that need buyer to combine purchased
    MinAmount int `json:"min_amount"`
    // start_time is The time when bundle deal activity start.The start time must be later than current time.
    StartTime int `json:"start_time"`
    // end_time is The time when bundle deal activity end. The end time must be 1 hour later than start time.
    EndTime int `json:"end_time"`
    // name is Title of the bundle deal
    Name string `json:"name"`
    // purchase_limit is Maximum number of bundle deals that can be bought by a buyer.
    PurchaseLimit int `json:"purchase_limit"`
}
//=======================================================
// BundleDealAddBundleDealResponse
//=======================================================
type BundleDealAddBundleDealResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// BundleDealAddBundleDealItemRequest
//=======================================================
type BundleDealAddBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items added in this bundle deal promotion.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// BundleDealAddBundleDealItemResponse
//=======================================================
type BundleDealAddBundleDealItemResponse struct {
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
// BundleDealGetBundleDealListRequest
//=======================================================
type BundleDealGetBundleDealListRequest struct {
    // page_size is Data paging, representing the data size of each page, the maximum is 1000, the default is 20
    PageSize int `json:"page_size,omitempty"`
    // time_status is The Status of bundle deal，all=1；upcoming=2；ongoing=3，expired=4 , the default is 1
    TimeStatus int `json:"time_status,omitempty"`
    // page_no is Data paging, represents the page number, starting from 1, the default is 1
    PageNo int `json:"page_no,omitempty"`
}
//=======================================================
// BundleDealGetBundleDealListResponse
//=======================================================
type BundleDealGetBundleDealListResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is  The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// BundleDealGetBundleDealRequest
//=======================================================
type BundleDealGetBundleDealRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
}
//=======================================================
// BundleDealGetBundleDealResponse
//=======================================================
type BundleDealGetBundleDealResponse struct {
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
// BundleDealGetBundleDealItemRequest
//=======================================================
type BundleDealGetBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
}
//=======================================================
// BundleDealGetBundleDealItemResponse
//=======================================================
type BundleDealGetBundleDealItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// BundleDealUpdateBundleDealRequest
//=======================================================
type BundleDealUpdateBundleDealRequest struct {
    // rule_type is The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
    RuleType int `json:"rule_type,omitempty"`
    // discount_value is  The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
    DiscountValue float64 `json:"discount_value,omitempty"`
    // fix_price is The amount of the buyer needs to spend to purchase a bundle deal.Need to input it when the bundle deal rule type is 1
    FixPrice float64 `json:"fix_price,omitempty"`
    // discount_percentage is The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
    DiscountPercentage int `json:"discount_percentage,omitempty"`
    // min_amount is The quantity of items that need buyer to combine purchased
    MinAmount int `json:"min_amount,omitempty"`
    // start_time is The time when bundle deal activity start.The start time must be later than current time.
    StartTime int `json:"start_time,omitempty"`
    // end_time is The time when bundle deal activity end. The end time must be later than current time.
    EndTime int `json:"end_time,omitempty"`
    // name is Title of the bundle deal
    Name string `json:"name,omitempty"`
    // purchase_limit is Maximum number of bundle deals that can be bought by a buyer.
    PurchaseLimit int `json:"purchase_limit,omitempty"`
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
}
//=======================================================
// BundleDealUpdateBundleDealResponse
//=======================================================
type BundleDealUpdateBundleDealResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// BundleDealUpdateBundleDealItemRequest
//=======================================================
type BundleDealUpdateBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items added in this bundle deal promotion.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// BundleDealUpdateBundleDealItemResponse
//=======================================================
type BundleDealUpdateBundleDealItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is  Detail informations you are querying.
    Response Response `json:"response"`
}
//=======================================================
// BundleDealEndBundleDealRequest
//=======================================================
type BundleDealEndBundleDealRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
}
//=======================================================
// BundleDealEndBundleDealResponse
//=======================================================
type BundleDealEndBundleDealResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is  The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}
//=======================================================
// BundleDealDeleteBundleDealRequest
//=======================================================
type BundleDealDeleteBundleDealRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
}
//=======================================================
// BundleDealDeleteBundleDealResponse
//=======================================================
type BundleDealDeleteBundleDealResponse struct {
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
// BundleDealDeleteBundleDealItemRequest
//=======================================================
type BundleDealDeleteBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items deleted in this bundle deal promotion.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// BundleDealDeleteBundleDealItemResponse
//=======================================================
type BundleDealDeleteBundleDealItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}