package shopeego
//=======================================================
// AddOnDealAddAddOnDealRequest
//=======================================================
type AddOnDealAddAddOnDealRequest struct {
    // add_on_deal_name is Title of the add on deal
    AddOnDealName string `json:"add_on_deal_name"`
    // start_time is The time when add on deal activity start.
    StartTime int `json:"start_time"`
    // end_time is The time when add on deal activity end
    EndTime int `json:"end_time"`
    // promotion_type is The type of add on deal：add on discount =0；gift with mini spend=1
    PromotionType int `json:"promotion_type"`
    // purchase_min_spend is The minimum purchase amount that needs to be met to buy the gift with min.Spend
    PurchaseMinSpend float64 `json:"purchase_min_spend,omitempty"`
    // per_gift_num is Number of gifts that buyers can get
    PerGiftNum int `json:"per_gift_num,omitempty"`
    // promotion_purchase_limit is promotion_purchase_limit
    PromotionPurchaseLimit int `json:"promotion_purchase_limit,omitempty"`
}
//=======================================================
// AddOnDealAddAddOnDealResponse
//=======================================================
type AddOnDealAddAddOnDealResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// AddOnDealAddAddOnDealMainItemRequest
//=======================================================
type AddOnDealAddAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // main_item_list is The main items added in this add on deal promotion.
    MainItemList []interface{} `json:"main_item_list"`
}
//=======================================================
// AddOnDealAddAddOnDealMainItemResponse
//=======================================================
type AddOnDealAddAddOnDealMainItemResponse struct {
    // error is  Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// AddOnDealAddAddOnDealSubItemRequest
//=======================================================
type AddOnDealAddAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []interface{} `json:"sub_item_list"`
}
//=======================================================
// AddOnDealAddAddOnDealSubItemResponse
//=======================================================
type AddOnDealAddAddOnDealSubItemResponse struct {
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
// AddOnDealDeleteAddOnDealRequest
//=======================================================
type AddOnDealDeleteAddOnDealRequest struct {
    // add_on_deal_id is Shopee's unique identifier for an add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
}
//=======================================================
// AddOnDealDeleteAddOnDealResponse
//=======================================================
type AddOnDealDeleteAddOnDealResponse struct {
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
// AddOnDealDeleteAddOnDealMainItemRequest
//=======================================================
type AddOnDealDeleteAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // main_item_list is  The main items added in this add on deal promotion.
    MainItemList []int `json:"main_item_list"`
}
//=======================================================
// AddOnDealDeleteAddOnDealMainItemResponse
//=======================================================
type AddOnDealDeleteAddOnDealMainItemResponse struct {
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
// AddOnDealDeleteAddOnDealSubItemRequest
//=======================================================
type AddOnDealDeleteAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []interface{} `json:"sub_item_list"`
}
//=======================================================
// AddOnDealDeleteAddOnDealSubItemResponse
//=======================================================
type AddOnDealDeleteAddOnDealSubItemResponse struct {
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
// AddOnDealGetAddOnDealListRequest
//=======================================================
type AddOnDealGetAddOnDealListRequest struct {
    // promotion_status is The Status of add on deal，default status is all
    PromotionStatus string `json:"promotion_status"`
    // page_no is The default page number is 1
    PageNo int `json:"page_no,omitempty"`
    // page_size is The default page size is 100
    PageSize int `json:"page_size,omitempty"`
}
//=======================================================
// AddOnDealGetAddOnDealListResponse
//=======================================================
type AddOnDealGetAddOnDealListResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// AddOnDealGetAddOnDealRequest
//=======================================================
type AddOnDealGetAddOnDealRequest struct {
    // add_on_deal_id is Shopee's unique identifier for an add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
}
//=======================================================
// AddOnDealGetAddOnDealResponse
//=======================================================
type AddOnDealGetAddOnDealResponse struct {
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
// AddOnDealGetAddOnDealMainItemRequest
//=======================================================
type AddOnDealGetAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
}
//=======================================================
// AddOnDealGetAddOnDealMainItemResponse
//=======================================================
type AddOnDealGetAddOnDealMainItemResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is  The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// AddOnDealGetAddOnDealSubItemRequest
//=======================================================
type AddOnDealGetAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
}
//=======================================================
// AddOnDealGetAddOnDealSubItemResponse
//=======================================================
type AddOnDealGetAddOnDealSubItemResponse struct {
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
// AddOnDealUpdateAddOnDealRequest
//=======================================================
type AddOnDealUpdateAddOnDealRequest struct {
    // add_on_deal_id is Shopee's unique identifier for an add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // start_time is The time when bundle deal activity start.The start time must be 1 hour than current time.
    StartTime int `json:"start_time,omitempty"`
    // end_time is The time when bundle deal activity end. The end time must be later than start time.
    EndTime int `json:"end_time,omitempty"`
    // purchase_min_spend is The minimum purchase amount that needs to be met to buy the gift with min.Spend
    PurchaseMinSpend float64 `json:"purchase_min_spend,omitempty"`
    // per_gift_num is Number of gifts that buyers can get
    PerGiftNum int `json:"per_gift_num,omitempty"`
    // promotion_purchase_limit is Max. number of add-on products that a customer can purchase per order.
    PromotionPurchaseLimit int `json:"promotion_purchase_limit,omitempty"`
    // sub_item_priority is The order of sub item
    SubItemPriority []int `json:"sub_item_priority,omitempty"`
    // add_on_deal_name is Title of the add on deal
    AddOnDealName string `json:"add_on_deal_name,omitempty"`
}
//=======================================================
// AddOnDealUpdateAddOnDealResponse
//=======================================================
type AddOnDealUpdateAddOnDealResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// AddOnDealUpdateAddOnDealMainItemRequest
//=======================================================
type AddOnDealUpdateAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // main_item_list is The main items added in this add on deal promotion.
    MainItemList []interface{} `json:"main_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealMainItemResponse
//=======================================================
type AddOnDealUpdateAddOnDealMainItemResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
//=======================================================
// AddOnDealUpdateAddOnDealSubItemRequest
//=======================================================
type AddOnDealUpdateAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []interface{} `json:"sub_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealSubItemResponse
//=======================================================
type AddOnDealUpdateAddOnDealSubItemResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
//=======================================================
// AddOnDealEndAddOnDealRequest
//=======================================================
type AddOnDealEndAddOnDealRequest struct {
    // add_on_deal_id is The identifier of the API request for error tracking
    AddOnDealID int `json:"add_on_deal_id"`
}
//=======================================================
// AddOnDealEndAddOnDealResponse
//=======================================================
type AddOnDealEndAddOnDealResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is The description of error code
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}