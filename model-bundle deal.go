package shopeego


//=======================================================
// Object Raw Type - BundleDealAddBundleDeal
//=======================================================
type BundleDealAddBundleDeal struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
}
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response BundleDealAddBundleDeal `json:"response"`
}


//=======================================================
// Object Raw Type - FailedList
//=======================================================
type FailedList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// fail_error is Indicate error type if one element hit error.
FailError string `json:"fail_error,omitempty"`
// fail_message is Indicate error details if one element hit error.
FailMessage string `json:"fail_message,omitempty"`
}


//=======================================================
// Object Raw Type - BundleDealAddBundleDealItem
//=======================================================
type BundleDealAddBundleDealItem struct {
// failed_list is Indicate error details.
FailedList []FailedList `json:"failed_list"`
// success_list is The list of succeed added items
SuccessList []SuccessList `json:"success_list"`
}
//=======================================================
// BundleDealAddBundleDealItemRequest
//=======================================================
type BundleDealAddBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items added in this bundle deal promotion.
    ItemList []ItemList `json:"item_list"`
}
//=======================================================
// BundleDealAddBundleDealItemResponse
//=======================================================
type BundleDealAddBundleDealItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response BundleDealAddBundleDealItem `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealRule
//=======================================================
type BundleDealRule struct {
// rule_type is The bundle deal rule type：FIX_PRICE = 1 ；DISCOUNT_PERCENTAGE = 2； DISCOUNT_VALUE = 3
RuleType int `json:"rule_type,omitempty"`
// discount_value is The deducted price when when buying a bundle deal. Need to input it when the bundle deal rule type is 3
DiscountValue float64 `json:"discount_value,omitempty"`
// fix_price is The amount of the buyer needs to spend to purchase a bundle deal. Need to input it when the bundle deal rule type is 1
FixPrice float64 `json:"fix_price,omitempty"`
// discount_percentage is The discount that the buyer can get when buying a bundle deal. Need to input it when the bundle deal rule type is 2
DiscountPercentage int `json:"discount_percentage,omitempty"`
// min_amount is The quantity of items that need buyer to combine purchased
MinAmount int `json:"min_amount,omitempty"`
}


//=======================================================
// Object Raw Type - BundleDealList
//=======================================================
type BundleDealList struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
// name is Title of the bundle deal
Name string `json:"name,omitempty"`
// start_time is The time when bundle deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when bundle deal activity end.
EndTime int `json:"end_time,omitempty"`
// bundle_deal_rule is 
BundleDealRule BundleDealRule `json:"bundle_deal_rule"`
// purchase_limit is Maximum number of bundle deals that can be bought by a buyer.
PurchaseLimit int `json:"purchase_limit,omitempty"`
}


//=======================================================
// Object Raw Type - BundleDealGetBundleDealList
//=======================================================
type BundleDealGetBundleDealList struct {
// bundle_deal_list is The list of bundle deal id
BundleDealList []BundleDealList `json:"bundle_deal_list"`
// more is this field shows whether there are more bundle deals in next page or not
More bool `json:"more,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response BundleDealGetBundleDealList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - BundleDealGetBundleDeal
//=======================================================
type BundleDealGetBundleDeal struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
// name is Title of the bundle deal
Name string `json:"name,omitempty"`
// start_time is The time when bundle deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when bundle deal activity end. 
EndTime int `json:"end_time,omitempty"`
// bundle_deal_rule is 
BundleDealRule BundleDealRule `json:"bundle_deal_rule"`
// purchase_limit is Maximum number of bundle deals that can be bought by a buyer.
PurchaseLimit int `json:"purchase_limit,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response BundleDealGetBundleDeal `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealGetBundleDealItem
//=======================================================
type BundleDealGetBundleDealItem struct {
// item_list is The list of bundle deal item
ItemList []ItemList `json:"item_list"`
// total_count is The number of  items in this bundle deal
TotalCount int `json:"total_count,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response BundleDealGetBundleDealItem `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealUpdateBundleDeal
//=======================================================
type BundleDealUpdateBundleDeal struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
// name is Title of the bundle deal
Name string `json:"name,omitempty"`
// start_time is The time when bundle deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when bundle deal activity end. 
EndTime int `json:"end_time,omitempty"`
// bundle_deal_rule is 
BundleDealRule BundleDealRule `json:"bundle_deal_rule"`
// purchase_limit is Maximum number of bundle deals that can be bought by a buyer.
PurchaseLimit int `json:"purchase_limit,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response BundleDealUpdateBundleDeal `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealUpdateBundleDealItem
//=======================================================
type BundleDealUpdateBundleDealItem struct {
// failed_list is Indicate error details.
FailedList []FailedList `json:"failed_list"`
// success_list is The list of succeed added items
SuccessList []SuccessList `json:"success_list"`
}
//=======================================================
// BundleDealUpdateBundleDealItemRequest
//=======================================================
type BundleDealUpdateBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items added in this bundle deal promotion.
    ItemList []ItemList `json:"item_list"`
}
//=======================================================
// BundleDealUpdateBundleDealItemResponse
//=======================================================
type BundleDealUpdateBundleDealItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is  Detail informations you are querying.
    Response BundleDealUpdateBundleDealItem `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealEndBundleDeal
//=======================================================
type BundleDealEndBundleDeal struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response BundleDealEndBundleDeal `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealDeleteBundleDeal
//=======================================================
type BundleDealDeleteBundleDeal struct {
// bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
BundleDealID int `json:"bundle_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response BundleDealDeleteBundleDeal `json:"response"`
}


//=======================================================
// Object Raw Type - BundleDealDeleteBundleDealItem
//=======================================================
type BundleDealDeleteBundleDealItem struct {
// failed_list is Indicate error details.
FailedList []FailedList `json:"failed_list"`
// success_list is The list of succeed added items
SuccessList []SuccessList `json:"success_list"`
}
//=======================================================
// BundleDealDeleteBundleDealItemRequest
//=======================================================
type BundleDealDeleteBundleDealItemRequest struct {
    // bundle_deal_id is Shopee's unique identifier for a bundle deal activity.
    BundleDealID int `json:"bundle_deal_id"`
    // item_list is The items deleted in this bundle deal promotion.
    ItemList []ItemList `json:"item_list"`
}
//=======================================================
// BundleDealDeleteBundleDealItemResponse
//=======================================================
type BundleDealDeleteBundleDealItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response BundleDealDeleteBundleDealItem `json:"response"`
}