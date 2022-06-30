package shopeego
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The name of the error raised in the request. The error will be empty if the request succeeds.
    Error string `json:"error,omitempty"`
    // message is The detailed description of the error. The message will be empty if the request succeeds.
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The name of the error raised in the request. The error will be empty if the request succeeds.
    Error string `json:"error,omitempty"`
    // message is The detailed description of the error. The message will be empty if the request succeeds.
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The name of the error raised in the request. The error will be empty if the request succeeds.
    Error string `json:"error,omitempty"`
    // message is The detailed description of the error. The message will be empty if the request succeeds.
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The name of the error raised in the request. The error will be empty if the request succeeds. 
    Error string `json:"error,omitempty"`
    // message is The detailed description of the error. The message will be empty if the request succeeds.
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is 
    Error string `json:"error,omitempty"`
    // message is 
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is 
    Error string `json:"error,omitempty"`
    // message is 
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is 
    Error string `json:"error,omitempty"`
    // message is 
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}