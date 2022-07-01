package shopeego


//=======================================================
// Object Raw Type - GlobalProductGetCategory
//=======================================================
type GlobalProductGetCategory struct {
// category_list is 
CategoryList []CategoryList `json:"category_list"`
}
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetCategory `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetAttributes
//=======================================================
type GlobalProductGetAttributes struct {
// attribute_list is Attribute info list.
AttributeList []AttributeList `json:"attribute_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetAttributes `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetBrandList
//=======================================================
type GlobalProductGetBrandList struct {
// brand_list is 
BrandList []BrandList `json:"brand_list"`
// has_next_page is  This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
HasNextPage bool `json:"has_next_page,omitempty"`
// next_offset is If has_next_page is true, this value need set to next request.offset
NextOffset int `json:"next_offset,omitempty"`
// is_mandatory is Whether is mandatory.
IsMandatory bool `json:"is_mandatory,omitempty"`
// input_type is Input type: DROP_DOWN  TEXT_FILED COMBO_BOX.
InputType string `json:"input_type,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetBrandList `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalItemNameLengthLimit
//=======================================================
type GlobalItemNameLengthLimit struct {
// min_limit is Global item name length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item name length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalItemImageCountLimit
//=======================================================
type GlobalItemImageCountLimit struct {
// min_limit is Global item image count min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item image count max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalItemDescriptionLengthLimit
//=======================================================
type GlobalItemDescriptionLengthLimit struct {
// min_limit is Global item description length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item description length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimit
//=======================================================
type GlobalProductGetGlobalItemLimit struct {
// price_limit is 
PriceLimit PriceLimit `json:"price_limit"`
// stock_limit is 
StockLimit StockLimit `json:"stock_limit"`
// global_item_name_length_limit is 
GlobalItemNameLengthLimit GlobalItemNameLengthLimit `json:"global_item_name_length_limit"`
// global_item_image_count_limit is 
GlobalItemImageCountLimit GlobalItemImageCountLimit `json:"global_item_image_count_limit"`
// global_item_description_length_limit is 
GlobalItemDescriptionLengthLimit GlobalItemDescriptionLengthLimit `json:"global_item_description_length_limit"`
// tier_variation_name_length_limit is 
TierVariationNameLengthLimit TierVariationNameLengthLimit `json:"tier_variation_name_length_limit"`
// tier_variation_option_length_limit is 
TierVariationOptionLengthLimit TierVariationOptionLengthLimit `json:"tier_variation_option_length_limit"`
// text_length_multiplier is Length ratio of Chinese characters to English characters in parameter verification. len(text)=len(Chinese characters)*text_length_multiplier+len(English characters )
TextLengthMultiplier float64 `json:"text_length_multiplier,omitempty"`
// extended_description_limit is 
ExtendedDescriptionLimit ExtendedDescriptionLimit `json:"extended_description_limit"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetGlobalItemLimit `json:"response"`
}


//=======================================================
// Object Raw Type - DaysToShipRangeList
//=======================================================
type DaysToShipRangeList struct {
// min_limit is Days to ship min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Days to ship max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetDtsLimit
//=======================================================
type GlobalProductGetDtsLimit struct {
// days_to_ship_range_list is If the length of the range list is greater than one, the final result is the union of multiple intervals
DaysToShipRangeList []DaysToShipRangeList `json:"days_to_ship_range_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetDtsLimit `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalItemList
//=======================================================
type GlobalItemList struct {
// global_item_id is Shopee's unique identifier for an global item.
GlobalItemID int `json:"global_item_id,omitempty"`
// update_time is Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
UpdateTime int `json:"update_time,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemList
//=======================================================
type GlobalProductGetGlobalItemList struct {
// global_item_list is 
GlobalItemList []GlobalItemList `json:"global_item_list"`
// total_count is Total global item count.
TotalCount int `json:"total_count,omitempty"`
// has_next_page is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
HasNextPage bool `json:"has_next_page,omitempty"`
// offset is If has_next_page is true, this value need set to next request.offset.
Offset string `json:"offset,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetGlobalItemList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfo
//=======================================================
type GlobalProductGetGlobalItemInfo struct {
// global_item_list is 
GlobalItemList []GlobalItemList `json:"global_item_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetGlobalItemInfo `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalProductAddGlobalItem
//=======================================================
type GlobalProductAddGlobalItem struct {
// global_item_id is Id of added global item.
GlobalItemID int `json:"global_item_id,omitempty"`
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
    AttributeList []AttributeList `json:"attribute_list,omitempty"`
    // description_info is New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is Values: See Data Definition- description_type (normal , extended). If you want to use extended_description, this field must be inputed
    DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// GlobalProductAddGlobalItemResponse
//=======================================================
type GlobalProductAddGlobalItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductAddGlobalItem `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductUpdateGlobalItem
//=======================================================
type GlobalProductUpdateGlobalItem struct {
// global_item_id is Id of updated global item.
GlobalItemID int `json:"global_item_id,omitempty"`
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
    AttributeList []AttributeList `json:"attribute_list,omitempty"`
    // description_info is New description field. New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is Values: See Data Definition- description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed
    DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// GlobalProductUpdateGlobalItemResponse
//=======================================================
type GlobalProductUpdateGlobalItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductUpdateGlobalItem `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - FailureDeleteItem
//=======================================================
type FailureDeleteItem struct {
// shop_id is The id of shop corresponding to the related item failed to delete.
ShopID int `json:"shop_id,omitempty"`
// item_id is The id of related item failed to delete.
ItemID int `json:"item_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductDeleteGlobalItem
//=======================================================
type GlobalProductDeleteGlobalItem struct {
// failure_delete_item is If delete failed, this field shows the details.
FailureDeleteItem []FailureDeleteItem `json:"failure_delete_item"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductDeleteGlobalItem `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalModel
//=======================================================
type GlobalModel struct {
// original_price is Original price of global model.
OriginalPrice float64 `json:"original_price,omitempty"`
// normal_stock is Normal stock of global model.
NormalStock int `json:"normal_stock,omitempty"`
// global_model_sku is Sku of global model. model_sku length information needs to be no more than 100 characters.
GlobalModelSku string `json:"global_model_sku,omitempty"`
// tier_index is Tier index of global model. Index starts from 0.
TierIndex []int `json:"tier_index,omitempty"`
}
//=======================================================
// GlobalProductInitTierVariationRequest
//=======================================================
type GlobalProductInitTierVariationRequest struct {
    // tier_variation is  Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []TierVariation `json:"tier_variation"`
    // global_model is Model info list, model number at most 50
    GlobalModel []GlobalModel `json:"global_model"`
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductInitTierVariationResponse
//=======================================================
type GlobalProductInitTierVariationResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// GlobalProductUpdateTierVariationRequest
//=======================================================
type GlobalProductUpdateTierVariationRequest struct {
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []TierVariation `json:"tier_variation"`
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
}
//=======================================================
// GlobalProductUpdateTierVariationResponse
//=======================================================
type GlobalProductUpdateTierVariationResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// GlobalProductAddGlobalModelRequest
//=======================================================
type GlobalProductAddGlobalModelRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // global_model is Global model setting list. Limit is  [1,50].
    GlobalModel []GlobalModel `json:"global_model"`
}
//=======================================================
// GlobalProductAddGlobalModelResponse
//=======================================================
type GlobalProductAddGlobalModelResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// GlobalProductUpdateGlobalModelRequest
//=======================================================
type GlobalProductUpdateGlobalModelRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // global_model is Sku setting for global model. Limit is [1,50].
    GlobalModel []GlobalModel `json:"global_model"`
}
//=======================================================
// GlobalProductUpdateGlobalModelResponse
//=======================================================
type GlobalProductUpdateGlobalModelResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - Failures
//=======================================================
type Failures struct {
// shop_id is Failed shop id.
ShopID int `json:"shop_id,omitempty"`
// item_id is Failed item id.
ItemID int `json:"item_id,omitempty"`
// model_id is Failed model id.
ModelID int `json:"model_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductDeleteGlobalModel
//=======================================================
type GlobalProductDeleteGlobalModel struct {
// global_model_id is Global model id.
GlobalModelID int `json:"global_model_id,omitempty"`
// failures is 
Failures []Failures `json:"failures"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductDeleteGlobalModel `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelList
//=======================================================
type GlobalProductGetGlobalModelList struct {
// tier_variation is Tier variation information of global item.
TierVariation []TierVariation `json:"tier_variation"`
// global_model is Global models.
GlobalModel []GlobalModel `json:"global_model"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetGlobalModelList `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalProductSupportSizeChart
//=======================================================
type GlobalProductSupportSizeChart struct {
// support_size_chart is If category support size chart.
SupportSizeChart bool `json:"support_size_chart,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductSupportSizeChart `json:"response"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - GlobalProductCreatePublishTask
//=======================================================
type GlobalProductCreatePublishTask struct {
// publish_task_id is The id of publish task.
PublishTaskID int `json:"publish_task_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductCreatePublishTask `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - PublishableShop
//=======================================================
type PublishableShop struct {
// shop_id is Id of publishable shop.
ShopID int `json:"shop_id,omitempty"`
// shop_region is Region of published shop.
ShopRegion string `json:"shop_region,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetPublishableShop
//=======================================================
type GlobalProductGetPublishableShop struct {
// publishable_shop is Detail of publishable shops.
PublishableShop []PublishableShop `json:"publishable_shop"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetPublishableShop `json:"response"`
}


//=======================================================
// Object Raw Type - Success
//=======================================================
type Success struct {
// region is The region of published item.
Region string `json:"region,omitempty"`
// shop_id is The shop id of published item.
ShopID string `json:"shop_id,omitempty"`
// item_id is The id of published item.
ItemID string `json:"item_id,omitempty"`
}


//=======================================================
// Object Raw Type - Failed
//=======================================================
type Failed struct {
// failed_reason is Failed reason.
FailedReason string `json:"failed_reason,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetPublishTaskResult
//=======================================================
type GlobalProductGetPublishTaskResult struct {
// publish_status is Status of publish task.
PublishStatus string `json:"publish_status,omitempty"`
// success is If publish task is successful, this field shows the published results.
Success Success `json:"success"`
// failed is If publish task is failed, this field shows the failed reason.
Failed Failed `json:"failed"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetPublishTaskResult `json:"response"`
}


//=======================================================
// Object Raw Type - PublishedItem
//=======================================================
type PublishedItem struct {
// shop_id is Shop id corresponding to the published item.
ShopID int `json:"shop_id,omitempty"`
// shop_region is Region of shop.
ShopRegion string `json:"shop_region,omitempty"`
// item_id is Id of published item.
ItemID int `json:"item_id,omitempty"`
// item_status is <p>Status of published item.Applicable values: 0.DELETED(Item is deleted by seller himself),1.NORMAL, 2.BANNED,3.REVIEWING,4.INVALID(Shopee Admin deleted),5.INVALID_HIDE(Shopee Admin delete confirmed),6.BLACKLISTED(Offensive_hide),8.NORMAL_UNLIST</p>
ItemStatus int `json:"item_status,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetPublishedList
//=======================================================
type GlobalProductGetPublishedList struct {
// published_item is Detail of published items.
PublishedItem []PublishedItem `json:"published_item"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetPublishedList `json:"response"`
}
//=======================================================
// GlobalProductUpdatePriceRequest
//=======================================================
type GlobalProductUpdatePriceRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // price_list is Price setting for global model. Limit is [1,50].
    PriceList []PriceList `json:"price_list"`
}
//=======================================================
// GlobalProductUpdatePriceResponse
//=======================================================
type GlobalProductUpdatePriceResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// GlobalProductUpdateStockRequest
//=======================================================
type GlobalProductUpdateStockRequest struct {
    // global_item_id is ID of global item.
    GlobalItemID int `json:"global_item_id"`
    // stock_list is Stock setting for global model. Limit is [1,50].
    StockList []StockList `json:"stock_list"`
}
//=======================================================
// GlobalProductUpdateStockResponse
//=======================================================
type GlobalProductUpdateStockResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - ShopSyncList
//=======================================================
type ShopSyncList struct {
// shop_id is Id of shop.
ShopID int `json:"shop_id,omitempty"`
// shop_region is TW TH MY BR IN SG VN
ShopRegion string `json:"shop_region,omitempty"`
// name_and_description is sync name and description
NameAndDescription bool `json:"name_and_description,omitempty"`
// media_information is sync media information
MediaInformation bool `json:"media_information,omitempty"`
// tier_variation_name_and_option is sync tier variation
TierVariationNameAndOption bool `json:"tier_variation_name_and_option,omitempty"`
// price is sync price
Price bool `json:"price,omitempty"`
// days_to_ship is sync days to ship info
DaysToShip bool `json:"days_to_ship,omitempty"`
}
//=======================================================
// GlobalProductSetSyncFieldRequest
//=======================================================
type GlobalProductSetSyncFieldRequest struct {
    // shop_sync_list is Length limit is [1,50].
    ShopSyncList []ShopSyncList `json:"shop_sync_list"`
}
//=======================================================
// GlobalProductSetSyncFieldResponse
//=======================================================
type GlobalProductSetSyncFieldResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - ItemIdMap
//=======================================================
type ItemIdMap struct {
// item_id is Id of item.
ItemID int `json:"item_id,omitempty"`
// global_item_id is Id of global item.
GlobalItemID int `json:"global_item_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemID
//=======================================================
type GlobalProductGetGlobalItemID struct {
// item_id_map is 
ItemIdMap []ItemIdMap `json:"item_id_map"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetGlobalItemID `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalProductCategoryRecommend
//=======================================================
type GlobalProductCategoryRecommend struct {
// category_id is Shopee's unique identifier for a category.
CategoryID []int `json:"category_id,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductCategoryRecommend `json:"response"`
}


//=======================================================
// Object Raw Type - GlobalProductGetRecommendAttribute
//=======================================================
type GlobalProductGetRecommendAttribute struct {
// attribute_list is Attribute info list.
AttributeList []AttributeList `json:"attribute_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response GlobalProductGetRecommendAttribute `json:"response,omitempty"`
}