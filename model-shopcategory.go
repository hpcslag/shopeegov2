package shopeego


//=======================================================
// Object Raw Type - ShopCategoryAddShopCategory
//=======================================================
type ShopCategoryAddShopCategory struct {
// shop_category_id is ShopCategory's unique identifier.
ShopCategoryID int `json:"shop_category_id,omitempty"`
}
//=======================================================
// ShopCategoryAddShopCategoryRequest
//=======================================================
type ShopCategoryAddShopCategoryRequest struct {
    // name is ShopCategory's name.
    Name string `json:"name"`
    // sort_weight is ShopCategory's sort weight. The maximum number should be 2147483546.
    SortWeight int `json:"sort_weight,omitempty"`
}
//=======================================================
// ShopCategoryAddShopCategoryResponse
//=======================================================
type ShopCategoryAddShopCategoryResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryAddShopCategory `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ShopCategoryGetShopCategoryListShopCategorys
//=======================================================
type ShopCategoryGetShopCategoryListShopCategorys struct {
// shop_category_id is ShopCategory's unique identifier.
ShopCategoryID int `json:"shop_category_id,omitempty"`
// status is ShopCategory's status. Applicable values--1: 'NORMAL', 2: 'INACTIVE', 0: 'DELETED'
Status int `json:"status,omitempty"`
// name is ShopCategory's name.
Name string `json:"name,omitempty"`
// sort_weight is ShopCategory's sort weight.
SortWeight int `json:"sort_weight,omitempty"`
}


//=======================================================
// Object Raw Type - ShopCategoryGetShopCategoryList
//=======================================================
type ShopCategoryGetShopCategoryList struct {
// shop_categorys is ShopCategory's unique identifier.
ShopCategorys []ShopCategoryGetShopCategoryListShopCategorys `json:"shop_categorys"`
// total_count is This is to indicate the whole number of  in-shop categories under the shop.
TotalCount int `json:"total_count,omitempty"`
// more is This is to indicate whether the list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest.
More bool `json:"more,omitempty"`
}
//=======================================================
// ShopCategoryGetShopCategoryListRequest
//=======================================================
type ShopCategoryGetShopCategoryListRequest struct {
    // page_size is Specifies the starting entry of data to return in the current call. The parameter range of page_size should be [1, 2147483647]
    PageSize int `json:"page_size"`
    // page_no is Specifies the total returned data per entry. The parameter range of page_no should be [1, 100]
    PageNo int `json:"page_no"`
}
//=======================================================
// ShopCategoryGetShopCategoryListResponse
//=======================================================
type ShopCategoryGetShopCategoryListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryGetShopCategoryList `json:"response"`
}


//=======================================================
// Object Raw Type - ShopCategoryDeleteShopCategory
//=======================================================
type ShopCategoryDeleteShopCategory struct {
// shop_category_id is ShopCategory's unique identifier.
ShopCategoryID int `json:"shop_category_id,omitempty"`
// msg is The return message of the operation result 
Msg string `json:"msg,omitempty"`
}
//=======================================================
// ShopCategoryDeleteShopCategoryRequest
//=======================================================
type ShopCategoryDeleteShopCategoryRequest struct {
    // shop_category_id is ShopCategory's unique identifier.
    ShopCategoryID int `json:"shop_category_id"`
}
//=======================================================
// ShopCategoryDeleteShopCategoryResponse
//=======================================================
type ShopCategoryDeleteShopCategoryResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryDeleteShopCategory `json:"response"`
}


//=======================================================
// Object Raw Type - ShopCategoryUpdateShopCategory
//=======================================================
type ShopCategoryUpdateShopCategory struct {
// shop_category_id is This is to indicate whether the shop categories list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categories
ShopCategoryID int `json:"shop_category_id,omitempty"`
// name is ShopCategory's name.
Name string `json:"name,omitempty"`
// sort_weight is ShopCategory's sort weight.
SortWeight int `json:"sort_weight,omitempty"`
// status is ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
Status string `json:"status,omitempty"`
}
//=======================================================
// ShopCategoryUpdateShopCategoryRequest
//=======================================================
type ShopCategoryUpdateShopCategoryRequest struct {
    // shop_category_id is ShopCategory's unique identifier.
    ShopCategoryID int `json:"shop_category_id"`
    // name is ShopCategory's name.
    Name string `json:"name,omitempty"`
    // sort_weight is ShopCategory's sort weight.
    SortWeight int `json:"sort_weight,omitempty"`
    // status is ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
    Status string `json:"status,omitempty"`
}
//=======================================================
// ShopCategoryUpdateShopCategoryResponse
//=======================================================
type ShopCategoryUpdateShopCategoryResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryUpdateShopCategory `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ShopCategoryAddItemListInvalidItemIdList
//=======================================================
type ShopCategoryAddItemListInvalidItemIdList struct {
// item_id is The invalid item id.
ItemID int `json:"item_id,omitempty"`
// fail_error is The reason of the fail.
FailError string `json:"fail_error,omitempty"`
// fail_message is The detailed reason of the failure and the hints of error fixing
FailMessage string `json:"fail_message,omitempty"`
}


//=======================================================
// Object Raw Type - ShopCategoryAddItemList
//=======================================================
type ShopCategoryAddItemList struct {
// invalid_item_id_list is List of invalid item ids.
InvalidItemIdList []ShopCategoryAddItemListInvalidItemIdList `json:"invalid_item_id_list"`
// shop_category_id is ShopCategory's unique identifier.
ShopCategoryID int `json:"shop_category_id,omitempty"`
// current_count is Count of items under this shop category after deletion.
CurrentCount int `json:"current_count,omitempty"`
}
//=======================================================
// ShopCategoryAddItemListRequest
//=======================================================
type ShopCategoryAddItemListRequest struct {
    // shop_category_id is ShopCategory's unique identifier.
    ShopCategoryID int `json:"shop_category_id"`
    // item_list is Shopee's unique identifiers list for an item. Max. 100 items to be deleted per request.
    ItemList []int `json:"item_list"`
}
//=======================================================
// ShopCategoryAddItemListResponse
//=======================================================
type ShopCategoryAddItemListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryAddItemList `json:"response"`
}


//=======================================================
// Object Raw Type - ShopCategoryGetItemList
//=======================================================
type ShopCategoryGetItemList struct {
// item_list is A list of Shopee's unique identifiers for items.
ItemList []int `json:"item_list,omitempty"`
// total_count is This is to indicate the whole number of items under the shop category.
TotalCount int `json:"total_count,omitempty"`
// more is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
More bool `json:"more,omitempty"`
}
//=======================================================
// ShopCategoryGetItemListRequest
//=======================================================
type ShopCategoryGetItemListRequest struct {
    // shop_category_id is ShopCategory's unique identifier.
    ShopCategoryID int `json:"shop_category_id"`
    // page_size is Specifies the starting entry of data to return in the current call. Default is 1000. The input range of page_size is [0, 1000]
    PageSize int `json:"page_size,omitempty"`
    // page_no is If many items are available to retrieve, you may need to call this api multiple times to retrieve all the data. And the default will be 0. page_size*page_no should be [0, 2147483446].
    PageNo int `json:"page_no,omitempty"`
}
//=======================================================
// ShopCategoryGetItemListResponse
//=======================================================
type ShopCategoryGetItemListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryGetItemList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ShopCategoryDeleteItemList
//=======================================================
type ShopCategoryDeleteItemList struct {
// shop_category_id is ShopCategory's unique identifier.
ShopCategoryID int `json:"shop_category_id,omitempty"`
// invalid_item_id is The list of item ids which are invalid; In other words, the item ids not being under the category.
InvalidItemID []int `json:"invalid_item_id,omitempty"`
// current_count is count of items under this shop category after deleting
CurrentCount int `json:"current_count,omitempty"`
}
//=======================================================
// ShopCategoryDeleteItemListRequest
//=======================================================
type ShopCategoryDeleteItemListRequest struct {
    // shop_category_id is The list of items need to be deleted. To note that the items which can be deleted successfully should be under this category.
    ShopCategoryID int `json:"shop_category_id"`
    // item_list is ShopCategory's unique identifier.
    ItemList []int `json:"item_list"`
}
//=======================================================
// ShopCategoryDeleteItemListResponse
//=======================================================
type ShopCategoryDeleteItemListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopCategoryDeleteItemList `json:"response"`
}