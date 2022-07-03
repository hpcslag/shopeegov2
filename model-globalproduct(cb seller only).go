package shopeego


//=======================================================
// Object Raw Type - GlobalProductGetCategoryCategory
//=======================================================
type GlobalProductGetCategoryCategory struct {
// category_id is 
CategoryID int `json:"category_id,omitempty"`
// parent_category_id is 
ParentCategoryID int `json:"parent_category_id,omitempty"`
// original_category_name is 
OriginalCategoryName string `json:"original_category_name,omitempty"`
// display_category_name is 
DisplayCategoryName string `json:"display_category_name,omitempty"`
// has_children is 
HasChildren bool `json:"has_children,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetCategory
//=======================================================
type GlobalProductGetCategory struct {
// category_list is 
CategoryList []GlobalProductGetCategoryCategory `json:"category_list"`
}
//=======================================================
// GlobalProductGetCategoryRequest
//=======================================================
type GlobalProductGetCategoryRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetAttributesAttributeAttributeValueParentAttribute
//=======================================================
type GlobalProductGetAttributesAttributeAttributeValueParentAttribute struct {
// parent_attribute_id is ID of parent attribute.
ParentAttributeID int `json:"parent_attribute_id,omitempty"`
// parent_value_id is ID of parent attribute value.
ParentValueID int `json:"parent_value_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetAttributesAttributeAttributeValueParentBrand
//=======================================================
type GlobalProductGetAttributesAttributeAttributeValueParentBrand struct {
// parent_brand_id is ID of parent brand.
ParentBrandID int `json:"parent_brand_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetAttributesAttributeAttributeValue
//=======================================================
type GlobalProductGetAttributesAttributeAttributeValue struct {
// value_id is ID of attribute value.
ValueID int `json:"value_id,omitempty"`
// original_value_name is Original name of value.
OriginalValueName string `json:"original_value_name,omitempty"`
// display_value_name is Display name of value.
DisplayValueName string `json:"display_value_name,omitempty"`
// value_unit is Unit of value.
ValueUnit string `json:"value_unit,omitempty"`
// parent_attribute_list is 
ParentAttributeList []GlobalProductGetAttributesAttributeAttributeValueParentAttribute `json:"parent_attribute_list"`
// parent_brand_list is 
ParentBrandList []GlobalProductGetAttributesAttributeAttributeValueParentBrand `json:"parent_brand_list"`
}


//=======================================================
// Object Raw Type - GlobalProductGetAttributesAttribute
//=======================================================
type GlobalProductGetAttributesAttribute struct {
// attribute_id is ID of attribute.
AttributeID int `json:"attribute_id,omitempty"`
// original_attribute_name is Attribute default name.
OriginalAttributeName string `json:"original_attribute_name,omitempty"`
// display_attribute_name is Attribute display name.
DisplayAttributeName string `json:"display_attribute_name,omitempty"`
// is_mandatory is Whether is mandatory.
IsMandatory bool `json:"is_mandatory,omitempty"`
// input_validation_type is Attribute input validation type : INT_TYPE STRING_TYPE ENUM_TYPE FLOAT_TYPE DATE_TYPE TIMESTAMP_TYPE
InputValidationType string `json:"input_validation_type,omitempty"`
// format_type is Attribute format type: NORMAL QUANTITATIVE .
FormatType string `json:"format_type,omitempty"`
// date_format_type is Attribute date format type:YEAR_MONTH_DATE YEAR_MONTH.
DateFormatType string `json:"date_format_type,omitempty"`
// input_type is Attribute input type: DROP_DOWN MULTIPLE_SELECT TEXT_FILED COMBO_BOX MULTIPLE_SELECT_COMBO_BOX.
InputType string `json:"input_type,omitempty"`
// attribute_unit is All applicable attribute units.
AttributeUnit []string `json:"attribute_unit,omitempty"`
// attribute_value_list is Value list of this attribute.
AttributeValueList []GlobalProductGetAttributesAttributeAttributeValue `json:"attribute_value_list"`
// mandatory_region is This attribute will be used to check the required regions.
MandatoryRegion []string `json:"mandatory_region,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetAttributes
//=======================================================
type GlobalProductGetAttributes struct {
// attribute_list is Attribute info list.
AttributeList []GlobalProductGetAttributesAttribute `json:"attribute_list"`
}
//=======================================================
// GlobalProductGetAttributesRequest
//=======================================================
type GlobalProductGetAttributesRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetBrandListBrand
//=======================================================
type GlobalProductGetBrandListBrand struct {
// brand_id is  Id of brand.
BrandID int `json:"brand_id,omitempty"`
// original_brand_name is Original name of brand
OriginalBrandName string `json:"original_brand_name,omitempty"`
// display_brand_name is  Display name of brand
DisplayBrandName string `json:"display_brand_name,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetBrandList
//=======================================================
type GlobalProductGetBrandList struct {
// brand_list is 
BrandList []GlobalProductGetBrandListBrand `json:"brand_list"`
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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetGlobalItemLimitPriceLimit
//=======================================================
type GlobalProductGetGlobalItemLimitPriceLimit struct {
// min_limit is Global item price min limit.
MinLimit float64 `json:"min_limit,omitempty,string"`
// max_limit is Global item price max limit.
MaxLimit float64 `json:"max_limit,omitempty,string"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitStockLimit
//=======================================================
type GlobalProductGetGlobalItemLimitStockLimit struct {
// min_limit is Global item stock min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item stock max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitGlobalItemNameLengthLimit
//=======================================================
type GlobalProductGetGlobalItemLimitGlobalItemNameLengthLimit struct {
// min_limit is Global item name length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item name length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitGlobalItemImageCountLimit
//=======================================================
type GlobalProductGetGlobalItemLimitGlobalItemImageCountLimit struct {
// min_limit is Global item image count min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item image count max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitGlobalItemDescriptionLengthLimit
//=======================================================
type GlobalProductGetGlobalItemLimitGlobalItemDescriptionLengthLimit struct {
// min_limit is Global item description length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item description length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitTierVariationNameLengthLimit
//=======================================================
type GlobalProductGetGlobalItemLimitTierVariationNameLengthLimit struct {
// min_limit is Global item tier variation name length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item tier variation name length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitTierVariationOptionLengthLimit
//=======================================================
type GlobalProductGetGlobalItemLimitTierVariationOptionLengthLimit struct {
// min_limit is Global item tier variation option length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Global item tier variation option length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimitExtendedDescriptionLimit
//=======================================================
type GlobalProductGetGlobalItemLimitExtendedDescriptionLimit struct {
// description_text_length_min is length min limit for item extended description text part, when one of the minimum limits for image and text is reached, the item can be added or updated successfully.
DescriptionTextLengthMin int `json:"description_text_length_min,omitempty"`
// description_text_length_max is length max limit for item extended description text part
DescriptionTextLengthMax int `json:"description_text_length_max,omitempty"`
// description_image_num_min is length min limit for item extended description image num, when one of the minimum limits for image and text is reached, the item can be added or updated successfully.
DescriptionImageNumMin int `json:"description_image_num_min,omitempty"`
// description_image_num_max is length max limit for item extended description image num
DescriptionImageNumMax int `json:"description_image_num_max,omitempty"`
// description_image_width_min is length min limit for item extended description image width
DescriptionImageWidthMin int `json:"description_image_width_min,omitempty"`
// description_image_height_min is length min limit for item extended description image hight
DescriptionImageHeightMin int `json:"description_image_height_min,omitempty"`
// description_image_aspect_ratio_min is length min limit for item extended description image aspect (image width / image hight )
DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min,omitempty,string"`
// description_image_aspect_ratio_max is length max limit for item extended description image aspect (image width / image hight )
DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max,omitempty,string"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemLimit
//=======================================================
type GlobalProductGetGlobalItemLimit struct {
// price_limit is 
PriceLimit GlobalProductGetGlobalItemLimitPriceLimit `json:"price_limit"`
// stock_limit is 
StockLimit GlobalProductGetGlobalItemLimitStockLimit `json:"stock_limit"`
// global_item_name_length_limit is 
GlobalItemNameLengthLimit GlobalProductGetGlobalItemLimitGlobalItemNameLengthLimit `json:"global_item_name_length_limit"`
// global_item_image_count_limit is 
GlobalItemImageCountLimit GlobalProductGetGlobalItemLimitGlobalItemImageCountLimit `json:"global_item_image_count_limit"`
// global_item_description_length_limit is 
GlobalItemDescriptionLengthLimit GlobalProductGetGlobalItemLimitGlobalItemDescriptionLengthLimit `json:"global_item_description_length_limit"`
// tier_variation_name_length_limit is 
TierVariationNameLengthLimit GlobalProductGetGlobalItemLimitTierVariationNameLengthLimit `json:"tier_variation_name_length_limit"`
// tier_variation_option_length_limit is 
TierVariationOptionLengthLimit GlobalProductGetGlobalItemLimitTierVariationOptionLengthLimit `json:"tier_variation_option_length_limit"`
// text_length_multiplier is Length ratio of Chinese characters to English characters in parameter verification. len(text)=len(Chinese characters)*text_length_multiplier+len(English characters )
TextLengthMultiplier float64 `json:"text_length_multiplier,omitempty,string"`
// extended_description_limit is 
ExtendedDescriptionLimit GlobalProductGetGlobalItemLimitExtendedDescriptionLimit `json:"extended_description_limit"`
}
//=======================================================
// GlobalProductGetGlobalItemLimitRequest
//=======================================================
type GlobalProductGetGlobalItemLimitRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetDtsLimitDaysToShipRange
//=======================================================
type GlobalProductGetDtsLimitDaysToShipRange struct {
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
DaysToShipRangeList []GlobalProductGetDtsLimitDaysToShipRange `json:"days_to_ship_range_list"`
}
//=======================================================
// GlobalProductGetDtsLimitRequest
//=======================================================
type GlobalProductGetDtsLimitRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetGlobalItemListGlobalItem
//=======================================================
type GlobalProductGetGlobalItemListGlobalItem struct {
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
GlobalItemList []GlobalProductGetGlobalItemListGlobalItem `json:"global_item_list"`
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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemStockInfo
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemStockInfo struct {
// stock_type is The stock type.
StockType int `json:"stock_type,omitempty"`
// stock_location_id is location_id of the stock.
StockLocationID string `json:"stock_location_id,omitempty"`
// normal_stock is The normal stock quantity of the variation in the listing currency.
NormalStock int `json:"normal_stock,omitempty"`
// reserved_stock is The reserved stock quantity of the variation in the listing currency.
ReservedStock int `json:"reserved_stock,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemPriceInfo
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemPriceInfo struct {
// currency is The three-digit code representing the currency unit used for the item in Shopee Listings.
Currency string `json:"currency,omitempty"`
// original_price is The original price of the item in the listing currency.
OriginalPrice float64 `json:"original_price,omitempty,string"`
// sip_item_price is SIP item price.
SipItemPrice float64 `json:"sip_item_price,omitempty,string"`
// sip_item_price_source is source of sip' price. ( auto or manual).
SipItemPriceSource string `json:"sip_item_price_source,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemImage
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemImage struct {
// image_id_list is List of image url.
ImageIdList []string `json:"image_id_list,omitempty"`
// image_url_list is List of image id.
ImageUrlList []string `json:"image_url_list,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemDimension
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemDimension struct {
// package_length is The length of package for this single item.
PackageLength int `json:"package_length,omitempty"`
// package_width is The width of package for this single item.
PackageWidth int `json:"package_width,omitempty"`
// package_height is The height of package for this single item.
PackageHeight int `json:"package_height,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemPreOrder
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemPreOrder struct {
// days_to_ship is Days to ship.
DaysToShip int `json:"days_to_ship,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemVideo
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemVideo struct {
// video_url is Url of video.
VideoUrl string `json:"video_url,omitempty"`
// thumbnail_url is Thumbnail of video.
ThumbnailUrl string `json:"thumbnail_url,omitempty"`
// duration is Duration of video.
Duration int `json:"duration,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemBrand
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemBrand struct {
// brand_id is Id of brand.
BrandID int `json:"brand_id,omitempty"`
// original_brand_name is Original name of brand.
OriginalBrandName string `json:"original_brand_name,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemAttributeAttributeValue
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemAttributeAttributeValue struct {
// value_id is Unique identifier for value of this item attribute.
ValueID int `json:"value_id,omitempty"`
// original_value_name is Value name of this item attribute.
OriginalValueName string `json:"original_value_name,omitempty"`
// value_unit is Value unit of this item attribute.
ValueUnit string `json:"value_unit,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemAttribute
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemAttribute struct {
// attribute_id is The Identify of each category.
AttributeID int `json:"attribute_id,omitempty"`
// original_attribute_name is The name of each attribute.
OriginalAttributeName string `json:"original_attribute_name,omitempty"`
// attribute_value_list is 
AttributeValueList []GlobalProductGetGlobalItemInfoGlobalItemAttributeAttributeValue `json:"attribute_value_list"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionFieldImageInfo
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionFieldImageInfo struct {
// image_id is Image id.
ImageID string `json:"image_id,omitempty"`
// image_url is Image url.
ImageUrl string `json:"image_url,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionField
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionField struct {
// field_type is Type of extended description field: values: See Data Definition- description_field_type (text , image).
FieldType string `json:"field_type,omitempty"`
// text is If field_type is text, text information will be returned through this field.
Text string `json:"text,omitempty"`
// image_info is If field_type is image, image url will be returned through this field.
ImageInfo GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionFieldImageInfo `json:"image_info"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescription
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescription struct {
// field_list is Field of extended description
FieldList []GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescriptionField `json:"field_list"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfo
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfo struct {
// extended_description is If description_type is extended , Description information will be returned through this field.
ExtendedDescription GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfoExtendedDescription `json:"extended_description"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfoGlobalItem
//=======================================================
type GlobalProductGetGlobalItemInfoGlobalItem struct {
// global_item_id is Shopee's unique identifier for an global item.
GlobalItemID int `json:"global_item_id,omitempty"`
// global_item_name is Name of the global item.
GlobalItemName string `json:"global_item_name,omitempty"`
// description is Description of the global item.
Description string `json:"description,omitempty"`
// global_item_sku is An global item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
GlobalItemSku string `json:"global_item_sku,omitempty"`
// global_item_status is The current status of the item. You can only query global product with normal status, otherwise api will return error.
GlobalItemStatus string `json:"global_item_status,omitempty"`
// create_time is Timestamp that indicates the date and time that the global item was created.
CreateTime int `json:"create_time,omitempty"`
// update_time is Timestamp that indicates the last time that there was a change in value of the global item.
UpdateTime int `json:"update_time,omitempty"`
// stock_info is If the item has models, this field will not be returned, please get it through get_model_list api.
StockInfo []GlobalProductGetGlobalItemInfoGlobalItemStockInfo `json:"stock_info"`
// price_info is If the item has models, price_info will not be returned. Please get the price of each model through the get_global_model_list api.
PriceInfo []GlobalProductGetGlobalItemInfoGlobalItemPriceInfo `json:"price_info"`
// image is 
Image GlobalProductGetGlobalItemInfoGlobalItemImage `json:"image"`
// weight is The weight of this item.
Weight float64 `json:"weight,omitempty,string"`
// dimension is 
Dimension GlobalProductGetGlobalItemInfoGlobalItemDimension `json:"dimension"`
// pre_order is 
PreOrder GlobalProductGetGlobalItemInfoGlobalItemPreOrder `json:"pre_order"`
// size_chart is Url of size chart image.
SizeChart string `json:"size_chart,omitempty"`
// condition is Is it second-hand.
Condition string `json:"condition,omitempty"`
// has_model is Does it contain model.
HasModel bool `json:"has_model,omitempty"`
// video is 
Video GlobalProductGetGlobalItemInfoGlobalItemVideo `json:"video"`
// category_id is Shopee's unique identifier for a category.
CategoryID int `json:"category_id,omitempty"`
// brand is 
Brand GlobalProductGetGlobalItemInfoGlobalItemBrand `json:"brand"`
// attribute_list is 
AttributeList []GlobalProductGetGlobalItemInfoGlobalItemAttribute `json:"attribute_list"`
// description_info is New description field.New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
DescriptionInfo GlobalProductGetGlobalItemInfoGlobalItemDescriptionInfo `json:"description_info"`
// description_type is Type of description : values: See Data Definition- description_type (normal , extended).
DescriptionType map[string]interface{} `json:"description_type,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalItemInfo
//=======================================================
type GlobalProductGetGlobalItemInfo struct {
// global_item_list is 
GlobalItemList []GlobalProductGetGlobalItemInfoGlobalItem `json:"global_item_list"`
}
//=======================================================
// GlobalProductGetGlobalItemInfoRequest
//=======================================================
type GlobalProductGetGlobalItemInfoRequest struct {
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    OriginalPrice float64 `json:"original_price,string"`
    // normal_stock is Normal stock of global item.
    NormalStock int `json:"normal_stock"`
    // weight is Weight of global item.
    Weight float64 `json:"weight,string"`
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
    V2RequestAuthenticationParams
    

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
    Weight float64 `json:"weight,omitempty,string"`
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
// Object Raw Type - GlobalProductDeleteGlobalItemFailureDeleteItem
//=======================================================
type GlobalProductDeleteGlobalItemFailureDeleteItem struct {
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
FailureDeleteItem []GlobalProductDeleteGlobalItemFailureDeleteItem `json:"failure_delete_item"`
}
//=======================================================
// GlobalProductDeleteGlobalItemRequest
//=======================================================
type GlobalProductDeleteGlobalItemRequest struct {
    V2RequestAuthenticationParams
    

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
OriginalPrice float64 `json:"original_price,omitempty,string"`
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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductDeleteGlobalModelFailure
//=======================================================
type GlobalProductDeleteGlobalModelFailure struct {
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
Failures []GlobalProductDeleteGlobalModelFailure `json:"failures"`
}
//=======================================================
// GlobalProductDeleteGlobalModelRequest
//=======================================================
type GlobalProductDeleteGlobalModelRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetGlobalModelListTierVariationOptionImage
//=======================================================
type GlobalProductGetGlobalModelListTierVariationOptionImage struct {
// image_url is Image url.
ImageUrl string `json:"image_url,omitempty"`
// image_id is Id of image.
ImageID string `json:"image_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelListTierVariationOption
//=======================================================
type GlobalProductGetGlobalModelListTierVariationOption struct {
// option is Tier option.
Option string `json:"option,omitempty"`
// image is Image information of tier.
Image GlobalProductGetGlobalModelListTierVariationOptionImage `json:"image"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelListTierVariation
//=======================================================
type GlobalProductGetGlobalModelListTierVariation struct {
// name is Tier name.
Name string `json:"name,omitempty"`
// option_list is Tier option list for corresponding tier name.
OptionList []GlobalProductGetGlobalModelListTierVariationOption `json:"option_list"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelListGlobalModelPriceInfo
//=======================================================
type GlobalProductGetGlobalModelListGlobalModelPriceInfo struct {
// original_price is Original price of global model.
OriginalPrice float64 `json:"original_price,omitempty,string"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelListGlobalModelStockInfo
//=======================================================
type GlobalProductGetGlobalModelListGlobalModelStockInfo struct {
// stock_type is Stock type. "1" means wms on hand, "2" means seller on hand.
StockType int `json:"stock_type,omitempty"`
// stock_location_id is Stock location id.
StockLocationID string `json:"stock_location_id,omitempty"`
// current_stock is Current stock.
CurrentStock int `json:"current_stock,omitempty"`
// normal_stock is Normal stock.
NormalStock int `json:"normal_stock,omitempty"`
// reserved_stock is Reserved stock.
ReservedStock int `json:"reserved_stock,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelListGlobalModel
//=======================================================
type GlobalProductGetGlobalModelListGlobalModel struct {
// global_model_id is Id of global model.
GlobalModelID int `json:"global_model_id,omitempty"`
// global_model_sku is Sku of global model.
GlobalModelSku string `json:"global_model_sku,omitempty"`
// price_info is Price info of global model.
PriceInfo GlobalProductGetGlobalModelListGlobalModelPriceInfo `json:"price_info"`
// stock_info is Stock info of global model.
StockInfo []GlobalProductGetGlobalModelListGlobalModelStockInfo `json:"stock_info"`
// tier_index is Tier index of global model.
TierIndex []int `json:"tier_index,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetGlobalModelList
//=======================================================
type GlobalProductGetGlobalModelList struct {
// tier_variation is Tier variation information of global item.
TierVariation []GlobalProductGetGlobalModelListTierVariation `json:"tier_variation"`
// global_model is Global models.
GlobalModel []GlobalProductGetGlobalModelListGlobalModel `json:"global_model"`
}
//=======================================================
// GlobalProductGetGlobalModelListRequest
//=======================================================
type GlobalProductGetGlobalModelListRequest struct {
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - ItemImage
//=======================================================
type ItemImage struct {
// image_id_list is Image id list of item.
ImageIdList []string `json:"image_id_list,omitempty"`
}


//=======================================================
// Object Raw Type - ItemTierVariationOptionImage
//=======================================================
type ItemTierVariationOptionImage struct {
// image_id is Image id.
ImageID string `json:"image_id,omitempty"`
}


//=======================================================
// Object Raw Type - ItemTierVariationOption
//=======================================================
type ItemTierVariationOption struct {
// option is Tier option.
Option string `json:"option,omitempty"`
// image is Image information for tier.
Image ItemTierVariationOptionImage `json:"image,omitempty"`
}


//=======================================================
// Object Raw Type - ItemTierVariation
//=======================================================
type ItemTierVariation struct {
// name is Tier name.
Name string `json:"name,omitempty"`
// option_list is Tier option list.
OptionList []ItemTierVariationOption `json:"option_list"`
}


//=======================================================
// Object Raw Type - ItemModel
//=======================================================
type ItemModel struct {
// tier_index is Tier index of model.
TierIndex []int `json:"tier_index,omitempty"`
// original_price is Original price of model.
OriginalPrice float64 `json:"original_price,omitempty,string"`
}


//=======================================================
// Object Raw Type - ItemLogistic
//=======================================================
type ItemLogistic struct {
// logistic_id is Logistic id.
LogisticID int `json:"logistic_id,omitempty"`
// enabled is If this logistic channel is enabled.
Enabled bool `json:"enabled,omitempty"`
// shipping_fee is Shipping fee.
ShippingFee float64 `json:"shipping_fee,omitempty,string"`
// size_id is Size id.
SizeID int `json:"size_id,omitempty"`
// is_free is If this logistic channel is free.
IsFree bool `json:"is_free,omitempty"`
}


//=======================================================
// Object Raw Type - ItemPreOrder
//=======================================================
type ItemPreOrder struct {
// is_pre_order is If this item is preorder.
IsPreOrder bool `json:"is_pre_order,omitempty"`
// days_to_ship is Days to ship, it's mandatory if is_pre_order is true.
DaysToShip int `json:"days_to_ship,omitempty"`
}


//=======================================================
// Object Raw Type - ItemDescriptionInfoExtendedDescriptionFieldImageInfo
//=======================================================
type ItemDescriptionInfoExtendedDescriptionFieldImageInfo struct {
// image_id is Image id.
ImageID string `json:"image_id,omitempty"`
}


//=======================================================
// Object Raw Type - ItemDescriptionInfoExtendedDescriptionField
//=======================================================
type ItemDescriptionInfoExtendedDescriptionField struct {
// field_type is Type of extended description field ：values: See Data Definition- description_field_type (text , image).
FieldType string `json:"field_type,omitempty"`
// text is If field_type is text， text information will be set by this field.
Text string `json:"text,omitempty"`
// image_info is If field_type is image，image url will be set by this field.
ImageInfo ItemDescriptionInfoExtendedDescriptionFieldImageInfo `json:"image_info,omitempty"`
}


//=======================================================
// Object Raw Type - ItemDescriptionInfoExtendedDescription
//=======================================================
type ItemDescriptionInfoExtendedDescription struct {
// field_list is Field of extended description.
FieldList []ItemDescriptionInfoExtendedDescriptionField `json:"field_list,omitempty"`
}


//=======================================================
// Object Raw Type - ItemDescriptionInfo
//=======================================================
type ItemDescriptionInfo struct {
// extended_description is If description_type is extended , Description information should be set by this field.
ExtendedDescription ItemDescriptionInfoExtendedDescription `json:"extended_description,omitempty"`
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
OriginalPrice float64 `json:"original_price,omitempty,string"`
// image is Image information of item.
Image ItemImage `json:"image,omitempty"`
// tier_variation is Tier variation information of item.
TierVariation []ItemTierVariation `json:"tier_variation,omitempty"`
// model is Model information of item.
Model []ItemModel `json:"model,omitempty"`
// size_chart is Size chart of item.
SizeChart string `json:"size_chart,omitempty"`
// logistic is Logistic information of item.
Logistic []ItemLogistic `json:"logistic,omitempty"`
// pre_order is Preorder information of item.
PreOrder ItemPreOrder `json:"pre_order,omitempty"`
// description_info is New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
DescriptionInfo ItemDescriptionInfo `json:"description_info,omitempty"`
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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetPublishableShopPublishableShop
//=======================================================
type GlobalProductGetPublishableShopPublishableShop struct {
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
PublishableShop []GlobalProductGetPublishableShopPublishableShop `json:"publishable_shop"`
}
//=======================================================
// GlobalProductGetPublishableShopRequest
//=======================================================
type GlobalProductGetPublishableShopRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetPublishTaskResultSucces
//=======================================================
type GlobalProductGetPublishTaskResultSucces struct {
// region is The region of published item.
Region string `json:"region,omitempty"`
// shop_id is The shop id of published item.
ShopID string `json:"shop_id,omitempty"`
// item_id is The id of published item.
ItemID string `json:"item_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetPublishTaskResultFailed
//=======================================================
type GlobalProductGetPublishTaskResultFailed struct {
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
Success GlobalProductGetPublishTaskResultSucces `json:"success"`
// failed is If publish task is failed, this field shows the failed reason.
Failed GlobalProductGetPublishTaskResultFailed `json:"failed"`
}
//=======================================================
// GlobalProductGetPublishTaskResultRequest
//=======================================================
type GlobalProductGetPublishTaskResultRequest struct {
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetPublishedListPublishedItem
//=======================================================
type GlobalProductGetPublishedListPublishedItem struct {
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
PublishedItem []GlobalProductGetPublishedListPublishedItem `json:"published_item"`
}
//=======================================================
// GlobalProductGetPublishedListRequest
//=======================================================
type GlobalProductGetPublishedListRequest struct {
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetGlobalItemIDItemIdMap
//=======================================================
type GlobalProductGetGlobalItemIDItemIdMap struct {
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
ItemIdMap []GlobalProductGetGlobalItemIDItemIdMap `json:"item_id_map"`
}
//=======================================================
// GlobalProductGetGlobalItemIDRequest
//=======================================================
type GlobalProductGetGlobalItemIDRequest struct {
    V2RequestAuthenticationParams
    

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
    V2RequestAuthenticationParams
    

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
// Object Raw Type - GlobalProductGetRecommendAttributeAttributeAttributeValue
//=======================================================
type GlobalProductGetRecommendAttributeAttributeAttributeValue struct {
// value_id is ID of attribute value.
ValueID int `json:"value_id,omitempty"`
}


//=======================================================
// Object Raw Type - GlobalProductGetRecommendAttributeAttribute
//=======================================================
type GlobalProductGetRecommendAttributeAttribute struct {
// attribute_id is ID of attribute.
AttributeID int `json:"attribute_id,omitempty"`
// attribute_value_list is Value list of this attribute.
AttributeValueList []GlobalProductGetRecommendAttributeAttributeAttributeValue `json:"attribute_value_list"`
}


//=======================================================
// Object Raw Type - GlobalProductGetRecommendAttribute
//=======================================================
type GlobalProductGetRecommendAttribute struct {
// attribute_list is Attribute info list.
AttributeList []GlobalProductGetRecommendAttributeAttribute `json:"attribute_list"`
}
//=======================================================
// GlobalProductGetRecommendAttributeRequest
//=======================================================
type GlobalProductGetRecommendAttributeRequest struct {
    V2RequestAuthenticationParams
    

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