package shopeego
//=======================================================
// GlobalProductGetCategoryRequest
//=======================================================
type GlobalProductGetCategoryRequest struct {
    // language is Display language. Language should be one of "zh-hans", "zh-hant", "ms-my", "en-my", "en", "id", "vi", "th", "pt-br".
    Language string `json:"language,omitempty"`
}
//=======================================================
// GlobalProductGetCategoryResponse
//=======================================================
type GlobalProductGetCategoryResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductGetAttributesRequest
//=======================================================
type GlobalProductGetAttributesRequest struct {
    // language is Language for display name zh-hans   zh-hant  ms-my   en-my   en   id   vi   th    pt-br.
    Language string `json:"language,omitempty"`
    // category_id is ID of category.
    CategoryID int `json:"category_id"`
    // need_region_mandatory is Do you want to return the mandatory of attribute_ Region attribute.
    NeedRegionMandatory bool `json:"need_region_mandatory,omitempty"`
}
//=======================================================
// GlobalProductGetAttributesResponse
//=======================================================
type GlobalProductGetAttributesResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductGetBrandListRequest
//=======================================================
type GlobalProductGetBrandListRequest struct {
    // offset is Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
    Offset int `json:"offset"`
    // page_size is the size of one page.
    PageSize int `json:"page_size"`
    // category_id is ID of category.
    CategoryID int `json:"category_id"`
    // status is Brand status , 1: normal brand, 2: pending brand.
    Status int `json:"status"`
}
//=======================================================
// GlobalProductGetBrandListResponse
//=======================================================
type GlobalProductGetBrandListResponse struct {
    // error is  Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetGlobalItemLimitRequest
//=======================================================
type GlobalProductGetGlobalItemLimitRequest struct {
}
//=======================================================
// GlobalProductGetGlobalItemLimitResponse
//=======================================================
type GlobalProductGetGlobalItemLimitResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning massage.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetDtsLimitRequest
//=======================================================
type GlobalProductGetDtsLimitRequest struct {
    // category_id is Id of category.
    CategoryID int `json:"category_id"`
}
//=======================================================
// GlobalProductGetDtsLimitResponse
//=======================================================
type GlobalProductGetDtsLimitResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetGlobalItemListRequest
//=======================================================
type GlobalProductGetGlobalItemListRequest struct {
    // offset is Specifies the starting entry of data to return in the current call. Default is null. if data is more than one page, the offset can be some entry to start next call.
    Offset string `json:"offset,omitempty"`
    // page_size is The size of one page. Limit is [1,50].
    PageSize int `json:"page_size"`
    // update_time_from is The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range.
    UpdateTimeFrom int `json:"update_time_from,omitempty"`
    // update_time_to is The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range
    UpdateTimeTo int `json:"update_time_to,omitempty"`
}
//=======================================================
// GlobalProductGetGlobalItemListResponse
//=======================================================
type GlobalProductGetGlobalItemListResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning massage.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductGetGlobalItemInfoRequest
//=======================================================
type GlobalProductGetGlobalItemInfoRequest struct {
    // global_item_id_list is Global item id list. Length limit is [1,20].
    GlobalItemIdList []int `json:"global_item_id_list"`
}
//=======================================================
// GlobalProductGetGlobalItemInfoResponse
//=======================================================
type GlobalProductGetGlobalItemInfoResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning massage.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductAddGlobalItemRequest
//=======================================================
type GlobalProductAddGlobalItemRequest struct {
    // category_id is Category id of global item.
    CategoryID int `json:"category_id"`
    // global_item_name is Name of global item.
    GlobalItemName string `json:"global_item_name"`
    // description is Description of global item.
    Description string `json:"description"`
    // global_item_sku is Sku of global item.
    GlobalItemSku string `json:"global_item_sku,omitempty"`
    // image is Image information of global item.
    Image Image `json:"image,omitempty"`
    // original_price is Original price of global item.
    OriginalPrice float64 `json:"original_price"`
    // normal_stock is Normal stock of global item.
    NormalStock int `json:"normal_stock"`
    // weight is Weight of global item.
    Weight float64 `json:"weight"`
    // dimension is Dimension information of global item.
    Dimension Dimension `json:"dimension,omitempty"`
    // pre_order is Preorder information of global item.
    PreOrder PreOrder `json:"pre_order"`
    // condition is Condition of global item, "NEW" or "USED" is available.
    Condition string `json:"condition,omitempty"`
    // video_upload_id is Video upload id of global item. Only accept one video_upload_id at most.
    VideoUploadID []string `json:"video_upload_id,omitempty"`
    // brand is 
    Brand Brand `json:"brand,omitempty"`
    // attribute_list is Item attributes.
    AttributeList []interface{} `json:"attribute_list,omitempty"`
    // description_info is New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is Values: See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed
    DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// GlobalProductAddGlobalItemResponse
//=======================================================
type GlobalProductAddGlobalItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductUpdateGlobalItemRequest
//=======================================================
type GlobalProductUpdateGlobalItemRequest struct {
    // global_item_id is Id of global item.
    GlobalItemID int `json:"global_item_id"`
    // category_id is Category id of global item.
    CategoryID int `json:"category_id,omitempty"`
    // global_item_name is Name of global item.
    GlobalItemName string `json:"global_item_name,omitempty"`
    // description is Description of global item.
    Description string `json:"description,omitempty"`
    // global_item_sku is Sku of global item.
    GlobalItemSku string `json:"global_item_sku,omitempty"`
    // weight is Weight of global item.
    Weight float64 `json:"weight,omitempty"`
    // dimension is Dimension information of global item.
    Dimension Dimension `json:"dimension,omitempty"`
    // pre_order is Preorder information of global item.
    PreOrder PreOrder `json:"pre_order,omitempty"`
    // condition is Condition of global item, "NEW" or "USED" is available.
    Condition string `json:"condition,omitempty"`
    // image is Image information of global item.
    Image Image `json:"image,omitempty"`
    // video_upload_id is Video upload id of global item.
    VideoUploadID []string `json:"video_upload_id,omitempty"`
    // brand is 
    Brand Brand `json:"brand,omitempty"`
    // attribute_list is Item attributes.
    AttributeList []interface{} `json:"attribute_list,omitempty"`
    // description_info is New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is Values: See Data Definition- description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed
    DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// GlobalProductUpdateGlobalItemResponse
//=======================================================
type GlobalProductUpdateGlobalItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductDeleteGlobalItemRequest
//=======================================================
type GlobalProductDeleteGlobalItemRequest struct {
    // global_item_id is The id of global item to delete.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductDeleteGlobalItemResponse
//=======================================================
type GlobalProductDeleteGlobalItemResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductInitTierVariationRequest
//=======================================================
type GlobalProductInitTierVariationRequest struct {
    // tier_variation is  Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []interface{} `json:"tier_variation"`
    // global_model is Model info list, model number at most 50
    GlobalModel []interface{} `json:"global_model"`
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductInitTierVariationResponse
//=======================================================
type GlobalProductInitTierVariationResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductUpdateTierVariationRequest
//=======================================================
type GlobalProductUpdateTierVariationRequest struct {
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []interface{} `json:"tier_variation"`
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductUpdateTierVariationResponse
//=======================================================
type GlobalProductUpdateTierVariationResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductAddGlobalModelRequest
//=======================================================
type GlobalProductAddGlobalModelRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // global_model is Global model setting list. Limit is  [1,50].
    GlobalModel []interface{} `json:"global_model"`
}
//=======================================================
// GlobalProductAddGlobalModelResponse
//=======================================================
type GlobalProductAddGlobalModelResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductUpdateGlobalModelRequest
//=======================================================
type GlobalProductUpdateGlobalModelRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // global_model is Sku setting for global model. Limit is [1,50].
    GlobalModel []interface{} `json:"global_model"`
}
//=======================================================
// GlobalProductUpdateGlobalModelResponse
//=======================================================
type GlobalProductUpdateGlobalModelResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductDeleteGlobalModelRequest
//=======================================================
type GlobalProductDeleteGlobalModelRequest struct {
    // global_item_id is Shopee's unique identifier for an global item.
    GlobalItemID int `json:"global_item_id"`
    // global_model_id is Shopee's unique identifier for an global model.
    GlobalModelID int `json:"global_model_id"`
}
//=======================================================
// GlobalProductDeleteGlobalModelResponse
//=======================================================
type GlobalProductDeleteGlobalModelResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning massage.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetGlobalModelListRequest
//=======================================================
type GlobalProductGetGlobalModelListRequest struct {
    // global_item_id is The id of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductGetGlobalModelListResponse
//=======================================================
type GlobalProductGetGlobalModelListResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductSupportSizeChartRequest
//=======================================================
type GlobalProductSupportSizeChartRequest struct {
    // category_id is Id of category.
    CategoryID int `json:"category_id"`
}
//=======================================================
// GlobalProductSupportSizeChartResponse
//=======================================================
type GlobalProductSupportSizeChartResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning massage.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductUpdateSizeChartRequest
//=======================================================
type GlobalProductUpdateSizeChartRequest struct {
    // global_item_id is Id of global item.
    GlobalItemID int `json:"global_item_id"`
    // size_chart is Image id of size chart.
    SizeChart string `json:"size_chart"`
}
//=======================================================
// GlobalProductUpdateSizeChartResponse
//=======================================================
type GlobalProductUpdateSizeChartResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}


//=======================================================
// Object Raw Type - Item
//=======================================================
type Item struct {
// item_name is Name of item.
ItemName string `json:"item_name,omitempty"`
// description is Description of item.
Description string `json:"description,omitempty"`
// item_status is Status of item.
ItemStatus string `json:"item_status,omitempty"`
// original_price is <p>Original price of item.</p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller:</font></b>&nbsp;Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.<br /></p>
OriginalPrice float64 `json:"original_price,omitempty"`
// image is Image information of item.
Image Image `json:"image,omitempty"`
// tier_variation is Tier variation information of item.
TierVariation []interface{} `json:"tier_variation,omitempty"`
// model is Model information of item.
Model []interface{} `json:"model,omitempty"`
// size_chart is Size chart of item.
SizeChart string `json:"size_chart,omitempty"`
// logistic is Logistic information of item.
Logistic []interface{} `json:"logistic,omitempty"`
// pre_order is Preorder information of item.
PreOrder PreOrder `json:"pre_order,omitempty"`
// description_info is New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
}
//=======================================================
// GlobalProductCreatePublishTaskRequest
//=======================================================
type GlobalProductCreatePublishTaskRequest struct {
    // global_item_id is Id of global item.
    GlobalItemID int `json:"global_item_id"`
    // shop_id is Id of shop to publish to.
    ShopID int `json:"shop_id"`
    // shop_region is Region of shop.
    ShopRegion string `json:"shop_region"`
    // item is Item information.
    Item Item `json:"item,omitempty"`
}
//=======================================================
// GlobalProductCreatePublishTaskResponse
//=======================================================
type GlobalProductCreatePublishTaskResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// GlobalProductGetPublishableShopRequest
//=======================================================
type GlobalProductGetPublishableShopRequest struct {
    // global_item_id is Id of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductGetPublishableShopResponse
//=======================================================
type GlobalProductGetPublishableShopResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetPublishTaskResultRequest
//=======================================================
type GlobalProductGetPublishTaskResultRequest struct {
    // publish_task_id is Id of publish task.
    PublishTaskID int `json:"publish_task_id"`
}
//=======================================================
// GlobalProductGetPublishTaskResultResponse
//=======================================================
type GlobalProductGetPublishTaskResultResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetPublishedListRequest
//=======================================================
type GlobalProductGetPublishedListRequest struct {
    // global_item_id is Id of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductGetPublishedListResponse
//=======================================================
type GlobalProductGetPublishedListResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductUpdatePriceRequest
//=======================================================
type GlobalProductUpdatePriceRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // price_list is Price setting for global model. Limit is [1,50].
    PriceList []interface{} `json:"price_list"`
}
//=======================================================
// GlobalProductUpdatePriceResponse
//=======================================================
type GlobalProductUpdatePriceResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductUpdateStockRequest
//=======================================================
type GlobalProductUpdateStockRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // stock_list is Stock setting for global model. Limit is [1,50].
    StockList []interface{} `json:"stock_list"`
}
//=======================================================
// GlobalProductUpdateStockResponse
//=======================================================
type GlobalProductUpdateStockResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductSetSyncFieldRequest
//=======================================================
type GlobalProductSetSyncFieldRequest struct {
    // shop_sync_list is Length limit is [1,50].
    ShopSyncList []interface{} `json:"shop_sync_list"`
}
//=======================================================
// GlobalProductSetSyncFieldResponse
//=======================================================
type GlobalProductSetSyncFieldResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning  message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// GlobalProductGetGlobalItemIDRequest
//=======================================================
type GlobalProductGetGlobalItemIDRequest struct {
    // shop_id is Id of shop.
    ShopID int `json:"shop_id"`
    // item_id_list is Item id list. Length limit is [1,20].
    ItemIdList []int `json:"item_id_list"`
}
//=======================================================
// GlobalProductGetGlobalItemIDResponse
//=======================================================
type GlobalProductGetGlobalItemIDResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductCategoryRecommendRequest
//=======================================================
type GlobalProductCategoryRecommendRequest struct {
    // global_item_name is name of item
    GlobalItemName string `json:"global_item_name"`
}
//=======================================================
// GlobalProductCategoryRecommendResponse
//=======================================================
type GlobalProductCategoryRecommendResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// GlobalProductGetRecommendAttributeRequest
//=======================================================
type GlobalProductGetRecommendAttributeRequest struct {
    // global_item_name is Name of item.
    GlobalItemName string `json:"global_item_name"`
    // category_id is ID of category.
    CategoryID int `json:"category_id"`
    // cover_image_id is ID of image.
    CoverImageID string `json:"cover_image_id,omitempty"`
}
//=======================================================
// GlobalProductGetRecommendAttributeResponse
//=======================================================
type GlobalProductGetRecommendAttributeResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}