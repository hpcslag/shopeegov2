package shopeego


//=======================================================
// Object Raw Type - AddOnDealAddAddOnDeal
//=======================================================
type AddOnDealAddAddOnDeal struct {
// add_on_deal_id is Shopee's unique identifier for an add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealAddAddOnDeal `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - MainItemList
//=======================================================
type MainItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// status is The status of add on deal item：enable = 1；disable =2
Status int `json:"status,omitempty"`
}


//=======================================================
// Object Raw Type - AddOnDealAddAddOnDealMainItem
//=======================================================
type AddOnDealAddAddOnDealMainItem struct {
// main_item_list is The main items added in this add on deal promotion.
MainItemList []MainItemList `json:"main_item_list"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
//=======================================================
// AddOnDealAddAddOnDealMainItemRequest
//=======================================================
type AddOnDealAddAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // main_item_list is The main items added in this add on deal promotion.
    MainItemList []MainItemList `json:"main_item_list"`
}
//=======================================================
// AddOnDealAddAddOnDealMainItemResponse
//=======================================================
type AddOnDealAddAddOnDealMainItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealAddAddOnDealMainItem `json:"response"`
}


//=======================================================
// Object Raw Type - SubItemList
//=======================================================
type SubItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// model_id is Shopee's unique identifier for a model.
ModelID int `json:"model_id,omitempty"`
// status is The status of add on deal item：enable = 1；disable =2
Status int `json:"status,omitempty"`
// sub_item_input_price is Add-on discount price before tax
SubItemInputPrice float64 `json:"sub_item_input_price,omitempty"`
// sub_item_limit is The purchase limit of sub item.
SubItemLimit int `json:"sub_item_limit,omitempty"`
}


//=======================================================
// Object Raw Type - AddOnDealAddAddOnDealSubItem
//=======================================================
type AddOnDealAddAddOnDealSubItem struct {
// sub_item_list is The sub items added in this add on deal promotion.
SubItemList []SubItemList `json:"sub_item_list"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
//=======================================================
// AddOnDealAddAddOnDealSubItemRequest
//=======================================================
type AddOnDealAddAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []SubItemList `json:"sub_item_list"`
}
//=======================================================
// AddOnDealAddAddOnDealSubItemResponse
//=======================================================
type AddOnDealAddAddOnDealSubItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealAddAddOnDealSubItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealDeleteAddOnDeal
//=======================================================
type AddOnDealDeleteAddOnDeal struct {
// add_on_deal_id is Shopee's unique identifier for an add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealDeleteAddOnDeal `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealDeleteAddOnDealMainItem
//=======================================================
type AddOnDealDeleteAddOnDealMainItem struct {
// main_item_list is The main items added in this add on deal promotion.
MainItemList []int `json:"main_item_list,omitempty"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealDeleteAddOnDealMainItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealDeleteAddOnDealSubItem
//=======================================================
type AddOnDealDeleteAddOnDealSubItem struct {
// sub_item_list is The sub items added in this add on deal promotion.
SubItemList []SubItemList `json:"sub_item_list"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
}
//=======================================================
// AddOnDealDeleteAddOnDealSubItemRequest
//=======================================================
type AddOnDealDeleteAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []SubItemList `json:"sub_item_list"`
}
//=======================================================
// AddOnDealDeleteAddOnDealSubItemResponse
//=======================================================
type AddOnDealDeleteAddOnDealSubItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealDeleteAddOnDealSubItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealList
//=======================================================
type AddOnDealList struct {
// start_time is The time when add on deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when add on deal activity end
EndTime int `json:"end_time,omitempty"`
// promotion_type is The type of add on deal：add on discount =0；gift with mini spend=1
PromotionType int `json:"promotion_type,omitempty"`
// purchase_min_spend is The minimum purchase amount that needs to be met to buy the gift with min.Spend
PurchaseMinSpend float64 `json:"purchase_min_spend,omitempty"`
// add_on_deal_id is Shopee's unique identifier for an add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// per_gift_num is Number of gifts that buyers can get
PerGiftNum int `json:"per_gift_num,omitempty"`
// promotion_purchase_limit is Max. number of add-on products that a customer can purchase per order.
PromotionPurchaseLimit int `json:"promotion_purchase_limit,omitempty"`
// add_on_deal_name is Title of the add on deal
AddOnDealName string `json:"add_on_deal_name,omitempty"`
// source is The create source of bundle deal：Seller=1，shopee admin=0
Source int `json:"source,omitempty"`
// sub_item_prioriry is The display sequence of sub item in buyer side
SubItemPrioriry []int `json:"sub_item_prioriry,omitempty"`
}


//=======================================================
// Object Raw Type - AddOnDealGetAddOnDealList
//=======================================================
type AddOnDealGetAddOnDealList struct {
// add_on_deal_list is The list of add on deal id
AddOnDealList []AddOnDealList `json:"add_on_deal_list"`
// more is This is to indicate whether the promotion list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of promotions.
More bool `json:"more,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealGetAddOnDealList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - AddOnDealGetAddOnDeal
//=======================================================
type AddOnDealGetAddOnDeal struct {
// start_time is The time when add on deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when add on deal activity end
EndTime int `json:"end_time,omitempty"`
// promotion_type is The type of add on deal：add on discount =0；gift with mini spend=1
PromotionType int `json:"promotion_type,omitempty"`
// purchase_min_spend is The minimum purchase amount that needs to be met to buy the gift with min.Spend
PurchaseMinSpend float64 `json:"purchase_min_spend,omitempty"`
// add_on_deal_id is Shopee's unique identifier for an add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// per_gift_num is Number of gifts that buyers can get
PerGiftNum int `json:"per_gift_num,omitempty"`
// sub_item_priority is The order of the sub item
SubItemPriority []int `json:"sub_item_priority,omitempty"`
// promotion_purchase_limit is Max. number of add-on products that a customer can purchase per order.
PromotionPurchaseLimit int `json:"promotion_purchase_limit,omitempty"`
// add_on_deal_name is Title of the add on deal
AddOnDealName string `json:"add_on_deal_name,omitempty"`
// source is 
Source int `json:"source,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealGetAddOnDeal `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealGetAddOnDealMainItem
//=======================================================
type AddOnDealGetAddOnDealMainItem struct {
// main_item_list is The main items added in this add on deal promotion.
MainItemList []MainItemList `json:"main_item_list"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealGetAddOnDealMainItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealGetAddOnDealSubItem
//=======================================================
type AddOnDealGetAddOnDealSubItem struct {
// sub_item_list is The sub items added in this add on deal promotion.
SubItemList []SubItemList `json:"sub_item_list"`
// add_on_deal_id is Shopee's unique identifier for add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealGetAddOnDealSubItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealUpdateAddOnDeal
//=======================================================
type AddOnDealUpdateAddOnDeal struct {
// start_time is The time when add on deal activity start.
StartTime int `json:"start_time,omitempty"`
// end_time is The time when add on deal activity end
EndTime int `json:"end_time,omitempty"`
// promotion_type is The type of add on deal：add on discount =0；gift with mini spend=1
PromotionType int `json:"promotion_type,omitempty"`
// purchase_min_spend is The minimum purchase amount that needs to be met to buy the gift with min.Spend
PurchaseMinSpend float64 `json:"purchase_min_spend,omitempty"`
// add_on_deal_id is Shopee's unique identifier for an add on deal activity.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// per_gift_num is Number of gifts that buyers can get
PerGiftNum int `json:"per_gift_num,omitempty"`
// promotion_purchase_limit is Max. number of add-on products that a customer can purchase per order.
PromotionPurchaseLimit int `json:"promotion_purchase_limit,omitempty"`
// add_on_deal_name is Title of the add on deal
AddOnDealName string `json:"add_on_deal_name,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealUpdateAddOnDeal `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - AddOnDealUpdateAddOnDealMainItem
//=======================================================
type AddOnDealUpdateAddOnDealMainItem struct {
// main_item_list is The main items added in this add on deal promotion.
MainItemList []MainItemList `json:"main_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealMainItemRequest
//=======================================================
type AddOnDealUpdateAddOnDealMainItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // main_item_list is The main items added in this add on deal promotion.
    MainItemList []MainItemList `json:"main_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealMainItemResponse
//=======================================================
type AddOnDealUpdateAddOnDealMainItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealUpdateAddOnDealMainItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealUpdateAddOnDealSubItem
//=======================================================
type AddOnDealUpdateAddOnDealSubItem struct {
// sub_item_list is The sub items added in this add on deal promotion.
SubItemList []SubItemList `json:"sub_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealSubItemRequest
//=======================================================
type AddOnDealUpdateAddOnDealSubItemRequest struct {
    // add_on_deal_id is Shopee's unique identifier for add on deal activity.
    AddOnDealID int `json:"add_on_deal_id"`
    // sub_item_list is The sub items added in this add on deal promotion.
    SubItemList []SubItemList `json:"sub_item_list"`
}
//=======================================================
// AddOnDealUpdateAddOnDealSubItemResponse
//=======================================================
type AddOnDealUpdateAddOnDealSubItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealUpdateAddOnDealSubItem `json:"response"`
}


//=======================================================
// Object Raw Type - AddOnDealEndAddOnDeal
//=======================================================
type AddOnDealEndAddOnDeal struct {
// add_on_deal_id is The identifier of the API request for error tracking
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response AddOnDealEndAddOnDeal `json:"response"`
}