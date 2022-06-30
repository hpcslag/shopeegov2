package shopeego


//=======================================================
// Object Raw Type - CategoryList
//=======================================================
type CategoryList struct {
// category_id is ID for category.
CategoryID int `json:"category_id,omitempty"`
// parent_category_id is ID for parent category.
ParentCategoryID int `json:"parent_category_id,omitempty"`
// original_category_name is Default name for category.
OriginalCategoryName string `json:"original_category_name,omitempty"`
// display_category_name is Display name dependent on display name.
DisplayCategoryName string `json:"display_category_name,omitempty"`
// has_children is Whether this category has active children category.
HasChildren bool `json:"has_children,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetCategory
//=======================================================
type ProductGetCategory struct {
// category_list is 
CategoryList CategoryList `json:"category_list"`
}
//=======================================================
// ProductGetCategoryRequest
//=======================================================
type ProductGetCategoryRequest struct {
    // language is <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ;  BR: en / pt-br ;  MX: en / es-mx ; PL: pl ; CO: en/es-CO ; CL: en/es-CL ; FR: en/fr ; ES: en/es-ES ; AR:en / es-ar Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data</p>
    Language string `json:"language,omitempty"`
}
//=======================================================
// ProductGetCategoryResponse
//=======================================================
type ProductGetCategoryResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetCategory `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ParentAttributeList
//=======================================================
type ParentAttributeList struct {
// parent_attribute_id is ID of parent attribute.
ParentAttributeID int `json:"parent_attribute_id,omitempty"`
// parent_value_id is ID of parent attribute value.
ParentValueID int `json:"parent_value_id,omitempty"`
}


//=======================================================
// Object Raw Type - ParentBrandList
//=======================================================
type ParentBrandList struct {
// parent_brand_id is ID of parent brand.
ParentBrandID int `json:"parent_brand_id,omitempty"`
}


//=======================================================
// Object Raw Type - AttributeValueList
//=======================================================
type AttributeValueList struct {
// value_id is ID of attribute value.
ValueID int `json:"value_id,omitempty"`
// original_value_name is Original name of value.
OriginalValueName string `json:"original_value_name,omitempty"`
// display_value_name is Display name of value.
DisplayValueName string `json:"display_value_name,omitempty"`
// value_unit is Unit of value(quantitative attribute only).
ValueUnit string `json:"value_unit,omitempty"`
// parent_attribute_list is 
ParentAttributeList ParentAttributeList `json:"parent_attribute_list"`
// parent_brand_list is 
ParentBrandList ParentBrandList `json:"parent_brand_list"`
}


//=======================================================
// Object Raw Type - AttributeList
//=======================================================
type AttributeList struct {
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
// attribute_unit is All applicable attribute units
AttributeUnit []string `json:"attribute_unit,omitempty"`
// attribute_value_list is Value list of this attribute.
AttributeValueList AttributeValueList `json:"attribute_value_list"`
}


//=======================================================
// Object Raw Type - ProductGetAttributes
//=======================================================
type ProductGetAttributes struct {
// attribute_list is Attribute info list.
AttributeList AttributeList `json:"attribute_list"`
}
//=======================================================
// ProductGetAttributesRequest
//=======================================================
type ProductGetAttributesRequest struct {
    // language is <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ;  BR: en / pt-br ;  MX: en / es-mx ; PL: pl ; CO: en/es-CO ; CL: en/es-CL ; FR: en/fr ; ES: en/es-ES ; AR: en / es-ar . Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data</p>
    Language string `json:"language,omitempty"`
    // category_id is ID of category.
    CategoryID int `json:"category_id"`
}
//=======================================================
// ProductGetAttributesResponse
//=======================================================
type ProductGetAttributesResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetAttributes `json:"response"`
}


//=======================================================
// Object Raw Type - BrandList
//=======================================================
type BrandList struct {
// original_brand_name is Original name of brand
OriginalBrandName string `json:"original_brand_name,omitempty"`
// brand_id is 
BrandID int `json:"brand_id,omitempty"`
// display_brand_name is Display name of brand
DisplayBrandName string `json:"display_brand_name,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetBrandList
//=======================================================
type ProductGetBrandList struct {
// brand_list is 
BrandList BrandList `json:"brand_list"`
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
// ProductGetBrandListRequest
//=======================================================
type ProductGetBrandListRequest struct {
    // offset is Specifies the starting entry of data to return in the current call. Default is 0. If data is more than one page,this field needs to be replaced with "next_offset" to request,and the offset can be some entry to start next call.
    Offset int `json:"offset,omitempty"`
    // page_size is the size of one page.Max=100
    PageSize int `json:"page_size"`
    // category_id is ID of category.
    CategoryID int `json:"category_id"`
    // status is Brand status , 1: normal brand, 2: pending brand
    Status int `json:"status"`
    // language is <p>If language is not uploaded, the default language=en, the following are the languages supported by different markets SG: en ; MY: en / ms-my / zh-hans ; TH: en / th ; VN: en / vi ; PH: en ; TW: en / zh-hant ; ID: en / id ;  BR: en / pt-br ;  MX: en / es-mx ; PL: pl ; CO: en/es-CO ; CL: en/es-CL ; FR: en/fr ; ES: en/es-ES ; AR: en/ es-ar . Note: For markets that have already launched global tree, Crossboard shop only support returning en and zh-hans language data</p>
    Language string `json:"language,omitempty"`
}
//=======================================================
// ProductGetBrandListResponse
//=======================================================
type ProductGetBrandListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetBrandList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - DaysToShipLimit
//=======================================================
type DaysToShipLimit struct {
// min_limit is Days to ship min limit
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Days to ship max limit
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetDtsLimit
//=======================================================
type ProductGetDtsLimit struct {
// days_to_ship_limit is 
DaysToShipLimit DaysToShipLimit `json:"days_to_ship_limit"`
// non_pre_order_days_to_ship is Non pre order days to ship
NonPreOrderDaysToShip int `json:"non_pre_order_days_to_ship,omitempty"`
}
//=======================================================
// ProductGetDtsLimitRequest
//=======================================================
type ProductGetDtsLimitRequest struct {
    // category_id is Shopee's unique identifier for an category.
    CategoryID int `json:"category_id"`
}
//=======================================================
// ProductGetDtsLimitResponse
//=======================================================
type ProductGetDtsLimitResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetDtsLimit `json:"response"`
}


//=======================================================
// Object Raw Type - PriceLimit
//=======================================================
type PriceLimit struct {
// min_limit is Item price max limit.
MinLimit float64 `json:"min_limit,omitempty"`
// max_limit is Item price min limit.
MaxLimit float64 `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - WholesalePriceThresholdPercentage
//=======================================================
type WholesalePriceThresholdPercentage struct {
// min_limit is Item wholesale price percentage of original price min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item wholesale price percentage of original price min limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - StockLimit
//=======================================================
type StockLimit struct {
// min_limit is Item stock min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item stock max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ItemNameLengthLimit
//=======================================================
type ItemNameLengthLimit struct {
// min_limit is Item name length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item name length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ItemImageCountLimit
//=======================================================
type ItemImageCountLimit struct {
// min_limit is Item image count min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item image count max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ItemDescriptionLengthLimit
//=======================================================
type ItemDescriptionLengthLimit struct {
// min_limit is Item description length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item description length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - TierVariationNameLengthLimit
//=======================================================
type TierVariationNameLengthLimit struct {
// min_limit is Item tier variation name length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item tier variation name length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - TierVariationOptionLengthLimit
//=======================================================
type TierVariationOptionLengthLimit struct {
// min_limit is Item tier variation option length min limit.
MinLimit int `json:"min_limit,omitempty"`
// max_limit is Item tier variation option length max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ItemCountLimit
//=======================================================
type ItemCountLimit struct {
// max_limit is Item count max limit.
MaxLimit int `json:"max_limit,omitempty"`
}


//=======================================================
// Object Raw Type - ExtendedDescriptionLimit
//=======================================================
type ExtendedDescriptionLimit struct {
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
// description_image_aspect_ratio_min is length min limit for item extended description image aspect  ( aspect_ratio= image width / image hight )
DescriptionImageAspectRatioMin float64 `json:"description_image_aspect_ratio_min,omitempty"`
// description_image_aspect_ratio_max is length max limit for item extended description image aspect ( aspect_ratio= image width / image hight )
DescriptionImageAspectRatioMax float64 `json:"description_image_aspect_ratio_max,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetItemLimit
//=======================================================
type ProductGetItemLimit struct {
// price_limit is 
PriceLimit PriceLimit `json:"price_limit"`
// wholesale_price_threshold_percentage is 
WholesalePriceThresholdPercentage WholesalePriceThresholdPercentage `json:"wholesale_price_threshold_percentage"`
// stock_limit is 
StockLimit StockLimit `json:"stock_limit"`
// item_name_length_limit is 
ItemNameLengthLimit ItemNameLengthLimit `json:"item_name_length_limit"`
// item_image_count_limit is 
ItemImageCountLimit ItemImageCountLimit `json:"item_image_count_limit"`
// item_description_length_limit is 
ItemDescriptionLengthLimit ItemDescriptionLengthLimit `json:"item_description_length_limit"`
// tier_variation_name_length_limit is 
TierVariationNameLengthLimit TierVariationNameLengthLimit `json:"tier_variation_name_length_limit"`
// tier_variation_option_length_limit is 
TierVariationOptionLengthLimit TierVariationOptionLengthLimit `json:"tier_variation_option_length_limit"`
// item_count_limit is 
ItemCountLimit ItemCountLimit `json:"item_count_limit"`
// extended_description_limit is 
ExtendedDescriptionLimit ExtendedDescriptionLimit `json:"extended_description_limit"`
}
//=======================================================
// ProductGetItemLimitRequest
//=======================================================
type ProductGetItemLimitRequest struct {
}
//=======================================================
// ProductGetItemLimitResponse
//=======================================================
type ProductGetItemLimitResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetItemLimit `json:"response"`
}


//=======================================================
// Object Raw Type - Item
//=======================================================
type Item struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// item_status is Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, UNLIST and BANNED.
ItemStatus string `json:"item_status,omitempty"`
// update_time is The update time of item.
UpdateTime int `json:"update_time,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetItemList
//=======================================================
type ProductGetItemList struct {
// item is list of item info with item_id/ item_status/ update_time
Item Item `json:"item"`
// total_count is total count of all items
TotalCount int `json:"total_count,omitempty"`
// has_next_page is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
HasNextPage bool `json:"has_next_page,omitempty"`
// next_offset is if has_next_page is true, this value need set to next request.offset
NextOffset int `json:"next_offset,omitempty"`
}
//=======================================================
// ProductGetItemListRequest
//=======================================================
type ProductGetItemListRequest struct {
    // offset is Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
    Offset int `json:"offset"`
    // page_size is the size of one page.Max=100
    PageSize int `json:"page_size"`
    // update_time_from is  The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range. 
    UpdateTimeFrom int `json:"update_time_from,omitempty"`
    // update_time_to is The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range
    UpdateTimeTo int `json:"update_time_to,omitempty"`
    // item_status is <p>NORMAL/BANNED/DELETED/UNLIST</p><p>If you want to search multiple status, please upload the url like this: item_status=NORMAL&amp;item_status=BANNED</p>
    ItemStatus []string `json:"item_status"`
}
//=======================================================
// ProductGetItemListResponse
//=======================================================
type ProductGetItemListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetItemList `json:"response"`
}


//=======================================================
// Object Raw Type - PriceInfo
//=======================================================
type PriceInfo struct {
// currency is The three-digit code representing the currency unit used for the item in Shopee Listings.
Currency string `json:"currency,omitempty"`
// original_price is The original price of the item in the listing currency.
OriginalPrice float64 `json:"original_price,omitempty"`
// current_price is The current price of the item in the listing currency. If product under a onging promotion, current_price will be the promotion price
CurrentPrice float64 `json:"current_price,omitempty"`
// inflated_price_of_original_price is The After-tax original price of the item in the listing currency.
InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price,omitempty"`
// inflated_price_of_current_price is The After-tax current price of the item in the listing currency.
InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price,omitempty"`
// sip_item_price is The price of the item in sip.If item is for CNSC primary shop, this field will not be returned
SipItemPrice float64 `json:"sip_item_price,omitempty"`
// sip_item_price_source is  source of sip' price. ( auto or manual).If item is for CNSC SIP primary shop, this field will not be returned
SipItemPriceSource string `json:"sip_item_price_source,omitempty"`
}


//=======================================================
// Object Raw Type - StockInfo
//=======================================================
type StockInfo struct {
// stock_type is The stock type. Applicable values: See Data Definition- StockType.
StockType int `json:"stock_type,omitempty"`
// stock_location_id is location_id of the stock.
StockLocationID string `json:"stock_location_id,omitempty"`
// current_stock is  Current available inventory, if item under promotion, it will be promotion stock, if not, it will be normal_stock
CurrentStock int `json:"current_stock,omitempty"`
// normal_stock is  The stock set by the seller.
NormalStock int `json:"normal_stock,omitempty"`
// reserved_stock is Promotion stock. Sellers can set Promotion stock for some promotion, which can only be used during the event. If the item with multiple promotion, this value is the total number of locked stocks for multiple promotions
ReservedStock int `json:"reserved_stock,omitempty"`
}


//=======================================================
// Object Raw Type - Image
//=======================================================
type Image struct {
// image_url_list is List of image url.
ImageUrlList []string `json:"image_url_list,omitempty"`
// image_id_list is List of image id.
ImageIdList []string `json:"image_id_list,omitempty"`
}


//=======================================================
// Object Raw Type - Dimension
//=======================================================
type Dimension struct {
// package_length is The length of package for this single item, the unit is CM.
PackageLength int `json:"package_length,omitempty"`
// package_width is The width of package for this single item, the unit is CM.
PackageWidth int `json:"package_width,omitempty"`
// package_height is The height of package for this single item, the unit is CM.
PackageHeight int `json:"package_height,omitempty"`
}


//=======================================================
// Object Raw Type - LogisticInfo
//=======================================================
type LogisticInfo struct {
// logistic_id is The identity of logistic channel.
LogisticID int `json:"logistic_id,omitempty"`
// logistic_name is The name of logistic.
LogisticName string `json:"logistic_name,omitempty"`
// enabled is Related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item.
Enabled bool `json:"enabled,omitempty"`
// shipping_fee is Only needed when logistics fee_type = CUSTOM_PRICE.
ShippingFee float64 `json:"shipping_fee,omitempty"`
// size_id is If specify logistic fee_type is SIZE_SELECTION size_id is required.
SizeID int `json:"size_id,omitempty"`
// is_free is when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
IsFree bool `json:"is_free,omitempty"`
// estimated_shipping_fee is Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty"`
}


//=======================================================
// Object Raw Type - PreOrder
//=======================================================
type PreOrder struct {
// is_pre_order is  Pre-order will be set true.
IsPreOrder bool `json:"is_pre_order,omitempty"`
// days_to_ship is The days to ship. Only work for pre-order, it means this value should be bigger than 7.
DaysToShip int `json:"days_to_ship,omitempty"`
}


//=======================================================
// Object Raw Type - Wholesales
//=======================================================
type Wholesales struct {
// min_count is The min count of this tier wholesale.
MinCount int `json:"min_count,omitempty"`
// max_count is The max count of this tier wholesale.
MaxCount int `json:"max_count,omitempty"`
// unit_price is The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
UnitPrice float64 `json:"unit_price,omitempty"`
// inflated_price_of_unit_price is The After-tax Price of the wholesale show to buyer.
InflatedPriceOfUnitPrice float64 `json:"inflated_price_of_unit_price,omitempty"`
}


//=======================================================
// Object Raw Type - VideoInfo
//=======================================================
type VideoInfo struct {
// video_url is Url of video.
VideoUrl string `json:"video_url,omitempty"`
// thumbnail_url is Thumbnail of video.
ThumbnailUrl string `json:"thumbnail_url,omitempty"`
// duration is Duration of video.
Duration int `json:"duration,omitempty"`
}


//=======================================================
// Object Raw Type - Brand
//=======================================================
type Brand struct {
// brand_id is Id of brand.
BrandID int `json:"brand_id,omitempty"`
// original_brand_name is Original name of brand.
OriginalBrandName string `json:"original_brand_name,omitempty"`
}


//=======================================================
// Object Raw Type - ComplaintPolicy
//=======================================================
type ComplaintPolicy struct {
// warranty_time is <p>Time for a warranty claim.Value should be in one of ONE_YEAR TWO_YEARS OVER_TWO_YEARS.<br /></p>
WarrantyTime string `json:"warranty_time,omitempty"`
// exclude_entrepreneur_warranty is <p>If True means "I exclude warranty complaints for entrepreneur"<br /></p>
ExcludeEntrepreneurWarranty bool `json:"exclude_entrepreneur_warranty,omitempty"`
// complaint_address_id is <p>The identity of complaint address.<br /></p>
ComplaintAddressID int `json:"complaint_address_id,omitempty"`
// additional_information is <p>Additional information for complaint policy<br /></p>
AdditionalInformation string `json:"additional_information,omitempty"`
}


//=======================================================
// Object Raw Type - TaxInfo
//=======================================================
type TaxInfo struct {
// ncm is <p>Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves.(<b><font color="#c24f4a">only for BR region</font></b>)<br /></p>
Ncm string `json:"ncm,omitempty"`
// diff_state_cfop is <p>Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice. (<b><font color="#c24f4a">only for BR region</font></b>)<br /></p>
DiffStateCfop string `json:"diff_state_cfop,omitempty"`
// csosn is <p>Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations. (<b><font color="#c24f4a">only for BR region</font></b>)<br /></p>
Csosn string `json:"csosn,omitempty"`
// origin is <p>Product source, domestic or foreig (<b><font color="#c24f4a">only for BR region)</font></b><br /></p>
Origin string `json:"origin,omitempty"`
// cest is <p>(<font color="#c24f4a" style><b>only for BR region</b></font>)<br /></p>
Cest string `json:"cest,omitempty"`
// measure_unit is <p>(<b><font color="#c24f4a">only for BR region</font></b>)<br /></p>
MeasureUnit string `json:"measure_unit,omitempty"`
// invoice_option is <p>Value shuold be one of <b>NO_INVOICES VAT_MARGIN_SCHEME_INVOICES</b> / <b>VAT_INVOICES /</b> <b>NON_VAT_INVOICES</b> and if value is NON_VAT_INVOICE vat_rate should be null (<b><font color="#c24f4a">only for PL region</font></b>)<br /></p>
InvoiceOption string `json:"invoice_option,omitempty"`
// vat_rate is <p>Value should be one of <b>0%</b> / <b>5%</b> / <b>8%</b> / <b>23%</b>&nbsp;/ <b>NO_VAT_RATE</b> (<b><font color="#c24f4a">only for PL region</font></b>)<br /></p>
VatRate string `json:"vat_rate,omitempty"`
// hs_code is <p>HS Code (<b><font color="#c24f4a">Only for IN region</font></b>)<br /></p>
HsCode string `json:"hs_code,omitempty"`
// tax_code is <p>Tax Code (<b><font color="#c24f4a">Only for IN region</font></b>)<br /></p>
TaxCode string `json:"tax_code,omitempty"`
}


//=======================================================
// Object Raw Type - SummaryInfo
//=======================================================
type SummaryInfo struct {
// total_reserved_stock is <p>total reserved stock<br /></p>
TotalReservedStock int `json:"total_reserved_stock,omitempty"`
// total_available_stock is <p>total available stock<br /></p>
TotalAvailableStock int `json:"total_available_stock,omitempty"`
}


//=======================================================
// Object Raw Type - SellerStock
//=======================================================
type SellerStock struct {
// location_id is 
LocationID string `json:"location_id,omitempty"`
// stock is <p>stock in the current warehouse<br /></p>
Stock int `json:"stock,omitempty"`
}


//=======================================================
// Object Raw Type - ShopeeStock
//=======================================================
type ShopeeStock struct {
// location_id is 
LocationID string `json:"location_id,omitempty"`
// stock is <p>stock in the current warehouse<br /></p>
Stock int `json:"stock,omitempty"`
}


//=======================================================
// Object Raw Type - StockInfoV2
//=======================================================
type StockInfoV2 struct {
// summary_info is <p>stock summary info<br /></p>
SummaryInfo SummaryInfo `json:"summary_info"`
// seller_stock is <p>seller stock<br /></p>
SellerStock SellerStock `json:"seller_stock"`
// shopee_stock is <p>shopee stock<br /></p>
ShopeeStock ShopeeStock `json:"shopee_stock"`
}


//=======================================================
// Object Raw Type - ItemList
//=======================================================
type ItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// category_id is Shopee's unique identifier for a category.
CategoryID int `json:"category_id,omitempty"`
// item_name is Name of the item in local language.
ItemName string `json:"item_name,omitempty"`
// description is if description_type is normal , Description information will be returned through this field，else description will be empty
Description string `json:"description,omitempty"`
// item_sku is An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
ItemSku string `json:"item_sku,omitempty"`
// create_time is Timestamp that indicates the date and time that the item was created.
CreateTime int `json:"create_time,omitempty"`
// update_time is Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
UpdateTime int `json:"update_time,omitempty"`
// attribute_list is 
AttributeList AttributeList `json:"attribute_list"`
// price_info is If the item has models, price_info will not be returned. Please get the price of each model through the get_model_list api
PriceInfo PriceInfo `json:"price_info"`
// stock_info is <p>if the item has models, this field will not be returned, please get it through get_model_list api.</p><p><b><font color="#c24f4a">Please use the stock_info_v2 field instead, we will deprecate this field in the future.</font></b><br /></p>
StockInfo StockInfo `json:"stock_info"`
// image is 
Image Image `json:"image"`
// weight is The net weight of this item, the unit is KG.
Weight string `json:"weight,omitempty"`
// dimension is The dimension of this item.
Dimension Dimension `json:"dimension"`
// logistic_info is The logistics list.
LogisticInfo LogisticInfo `json:"logistic_info"`
// pre_order is 
PreOrder PreOrder `json:"pre_order"`
// wholesales is The wholesales tier list.
Wholesales Wholesales `json:"wholesales"`
// condition is Is it second-hand.
Condition string `json:"condition,omitempty"`
// size_chart is Url of size chart image.
SizeChart string `json:"size_chart,omitempty"`
// item_status is Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, BANNED and UNLIST.
ItemStatus string `json:"item_status,omitempty"`
// has_model is Does it contain model.
HasModel bool `json:"has_model,omitempty"`
// promotion_id is 
PromotionID int `json:"promotion_id,omitempty"`
// video_info is Info of video list.
VideoInfo VideoInfo `json:"video_info"`
// brand is 
Brand Brand `json:"brand"`
// item_dangerous is This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
ItemDangerous int `json:"item_dangerous,omitempty"`
// complaint_policy is <p>Time for a warranty claim.Value should be in one of ONE_YEAR TWO_YEARS OVER_TWO_YEARS.<br /></p>
ComplaintPolicy ComplaintPolicy `json:"complaint_policy"`
// tax_info is <p>Tax information<br /></p>
TaxInfo TaxInfo `json:"tax_info"`
// stock_info_v2 is <p>new stock object.<br /></p><p>Please check this FAQ for more detail:<a style="font-size:14px;">https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230</a></p><a href="https://open.shopee.com/faq?top=162&amp;sub=166&amp;page=1&amp;faq=230" target="_blank"></a>
StockInfoV2 StockInfoV2 `json:"stock_info_v2"`
}


//=======================================================
// Object Raw Type - ImageInfo
//=======================================================
type ImageInfo struct {
// image_id is Image id
ImageID string `json:"image_id,omitempty"`
// image_url is Image url.
ImageUrl string `json:"image_url,omitempty"`
}


//=======================================================
// Object Raw Type - FieldList
//=======================================================
type FieldList struct {
// field_type is Type of extended description field ：values: See Data Definition- description_field_type (text , image).
FieldType string `json:"field_type,omitempty"`
// text is If field_type is text, text information will be returned through this field.
Text string `json:"text,omitempty"`
// image_info is If field_type is image, image url will be returned through this field.
ImageInfo ImageInfo `json:"image_info"`
}


//=======================================================
// Object Raw Type - ExtendedDescription
//=======================================================
type ExtendedDescription struct {
// field_list is Field of extended description
FieldList FieldList `json:"field_list"`
}


//=======================================================
// Object Raw Type - DescriptionInfo
//=======================================================
type DescriptionInfo struct {
// extended_description is  If description_type is extended , Description information will be returned through this field.
ExtendedDescription ExtendedDescription `json:"extended_description"`
}


//=======================================================
// Object Raw Type - ProductGetItemBaseInfo
//=======================================================
type ProductGetItemBaseInfo struct {
// item_list is 
ItemList ItemList `json:"item_list"`
// description_info is New description  field. Only whitelist sellers can use it.
DescriptionInfo DescriptionInfo `json:"description_info"`
// description_type is Type of description : values: See Data Definition- description_type (normal , extended).
DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// ProductGetItemBaseInfoRequest
//=======================================================
type ProductGetItemBaseInfoRequest struct {
    // item_id_list is item_id list; limit [0,50]
    ItemIdList []int `json:"item_id_list"`
    // need_tax_info is if true will response tax_info
    NeedTaxInfo bool `json:"need_tax_info,omitempty"`
    // need_complaint_policy is if true will response complaint_policy
    NeedComplaintPolicy bool `json:"need_complaint_policy,omitempty"`
}
//=======================================================
// ProductGetItemBaseInfoResponse
//=======================================================
type ProductGetItemBaseInfoResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetItemBaseInfo `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ProductGetItemExtraInfo
//=======================================================
type ProductGetItemExtraInfo struct {
// item_list is extra info of item list.
ItemList ItemList `json:"item_list"`
}
//=======================================================
// ProductGetItemExtraInfoRequest
//=======================================================
type ProductGetItemExtraInfoRequest struct {
    // item_id_list is  item_id list, limit [0,50]
    ItemIdList []int `json:"item_id_list"`
}
//=======================================================
// ProductGetItemExtraInfoResponse
//=======================================================
type ProductGetItemExtraInfoResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetItemExtraInfo `json:"response"`
}


//=======================================================
// Object Raw Type - Wholesale
//=======================================================
type Wholesale struct {
// min_count is Minimum count of this tier
MinCount int `json:"min_count,omitempty"`
// max_count is Maximum count of this tier
MaxCount int `json:"max_count,omitempty"`
// unit_price is <p>Unit price of this tier.</p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller</font></b>:&nbsp;Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.<br /></p>
UnitPrice float64 `json:"unit_price,omitempty"`
}


//=======================================================
// Object Raw Type - Images
//=======================================================
type Images struct {
// image_id_list is ID of image
ImageIdList []string `json:"image_id_list,omitempty"`
// image_url_list is Display URL of image
ImageUrlList []string `json:"image_url_list,omitempty"`
}


//=======================================================
// Object Raw Type - Attributes
//=======================================================
type Attributes struct {
// attribute_id is Attribute ID
AttributeID int `json:"attribute_id,omitempty"`
// attribute_value_list is 
AttributeValueList AttributeValueList `json:"attribute_value_list"`
}


//=======================================================
// Object Raw Type - ProductAddItem
//=======================================================
type ProductAddItem struct {
// description is Description of item
Description string `json:"description,omitempty"`
// weight is Item weight
Weight int `json:"weight,omitempty"`
// pre_order is Pre order setting
PreOrder PreOrder `json:"pre_order"`
// item_name is Item name
ItemName string `json:"item_name,omitempty"`
// images is Item images
Images Images `json:"images"`
// item_status is Item status
ItemStatus string `json:"item_status,omitempty"`
// price_info is Item price info
PriceInfo PriceInfo `json:"price_info"`
// logistic_info is Logistic setting
LogisticInfo LogisticInfo `json:"logistic_info"`
// stock_info is Stock info
StockInfo StockInfo `json:"stock_info"`
// item_id is Item ID
ItemID int `json:"item_id,omitempty"`
// attributes is Item attributes
Attributes Attributes `json:"attributes"`
// category_id is Category ID
CategoryID int `json:"category_id,omitempty"`
// dimension is Item dimension
Dimension Dimension `json:"dimension"`
// condition is Item condition, could be NEW or USED
Condition string `json:"condition,omitempty"`
// video_info is Item video
VideoInfo VideoInfo `json:"video_info"`
// wholesale is Wholesale setting
Wholesale Wholesale `json:"wholesale"`
// brand is 
Brand Brand `json:"brand"`
// item_dangerous is This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
ItemDangerous int `json:"item_dangerous,omitempty"`
// description_info is New description field. Only whitelist sellers can use it. If item with extended_description this field will return, otherwise do not return.
DescriptionInfo DescriptionInfo `json:"description_info"`
// description_type is <p>description_type (normal , extended).</p>
DescriptionType string `json:"description_type,omitempty"`
// complaint_policy is Complaint Policy for item. Only returned for local PL sellers.
ComplaintPolicy ComplaintPolicy `json:"complaint_policy"`
// seller_stock is <p>seller stock<br /></p>
SellerStock SellerStock `json:"seller_stock"`
}
//=======================================================
// ProductAddItemRequest
//=======================================================
type ProductAddItemRequest struct {
    // original_price is <p>Item price</p><p><b><font color="#c24f4a">For CO local VAT responsible seller：</font></b>Please remember the price you set in here must be VAT inclusive. If you have any doubts on how to calculate VAT for your product please refer to the Seller Education Hub（https://seller.shopee.com.co/edu/article/13565）<br /></p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller:&nbsp;</font></b><span style="font-size:14px;"></span><span style="font-size:14px;">Sellers can set the price with two decimal place,&nbsp;</span><span style="font-size:14px;">other regions can only set the price as an integer.</span></p>
    OriginalPrice float64 `json:"original_price"`
    // description is if description_type is normal , Description information should be set by this field.
    Description string `json:"description"`
    // weight is Weight of item
    Weight float64 `json:"weight,omitempty"`
    // item_name is Item name
    ItemName string `json:"item_name"`
    // item_status is Item status, could be UNLIST or NORMAL
    ItemStatus string `json:"item_status,omitempty"`
    // dimension is Item dimension
    Dimension Dimension `json:"dimension,omitempty"`
    // normal_stock is Item stock
    NormalStock int `json:"normal_stock"`
    // logistic_info is Logistic channel setting
    LogisticInfo LogisticInfo `json:"logistic_info"`
    // attribute_list is <p>This field is optional(expect Indonesia) depending on the specific attribute under different categories. Should call get_attribute api to get attribute first. Must contain all all mandatory attribute.</p>
    AttributeList AttributeList `json:"attribute_list,omitempty"`
    // category_id is ID of category
    CategoryID int `json:"category_id"`
    // image is Item images
    Image Image `json:"image"`
    // pre_order is Pre order setting
    PreOrder PreOrder `json:"pre_order,omitempty"`
    // item_sku is SKU tag of item
    ItemSku string `json:"item_sku,omitempty"`
    // condition is Condition of item, could be USED or NEW
    Condition string `json:"condition,omitempty"`
    // wholesale is <p>Wholesale setting.</p>
    Wholesale Wholesale `json:"wholesale,omitempty"`
    // video_upload_id is Video upload ID returned from video uploading API. Only accept one video_upload_id.
    VideoUploadID []string `json:"video_upload_id,omitempty"`
    // brand is 
    Brand Brand `json:"brand,omitempty"`
    // item_dangerous is This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
    ItemDangerous int `json:"item_dangerous,omitempty"`
    // tax_info is Tax information
    TaxInfo TaxInfo `json:"tax_info,omitempty"`
    // complaint_policy is Complaint Policy for item. Only required for local PL sellers, ignored otherwise.
    ComplaintPolicy ComplaintPolicy `json:"complaint_policy,omitempty"`
    // description_info is New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is <p>description_type (normal , extended). If you want to use extended_description, this field must be inputed</p>
    DescriptionType string `json:"description_type,omitempty"`
    // seller_stock is <p>seller stock（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
    SellerStock SellerStock `json:"seller_stock,omitempty"`
}
//=======================================================
// ProductAddItemResponse
//=======================================================
type ProductAddItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductAddItem `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ProductUpdateItem
//=======================================================
type ProductUpdateItem struct {
// description is Item description.
Description string `json:"description,omitempty"`
// weight is Item weight.
Weight float64 `json:"weight,omitempty"`
// pre_order is 
PreOrder PreOrder `json:"pre_order"`
// item_name is Item name.
ItemName string `json:"item_name,omitempty"`
// item_status is Item status.
ItemStatus string `json:"item_status,omitempty"`
// images is Item images.
Images Images `json:"images"`
// logistic_info is 
LogisticInfo LogisticInfo `json:"logistic_info"`
// item_id is ID of item.
ItemID int `json:"item_id,omitempty"`
// category_id is ID of item category.
CategoryID int `json:"category_id,omitempty"`
// dimension is 
Dimension Dimension `json:"dimension"`
// condition is Item condition, could be USED or NEW.
Condition string `json:"condition,omitempty"`
// brand is 
Brand Brand `json:"brand"`
// item_dangerous is This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
ItemDangerous int `json:"item_dangerous,omitempty"`
// complaint_policy is Complaint policy
ComplaintPolicy ComplaintPolicy `json:"complaint_policy"`
// description_info is New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
DescriptionInfo DescriptionInfo `json:"description_info"`
// description_type is <p>description_type (normal , extended).</p>
DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// ProductUpdateItemRequest
//=======================================================
type ProductUpdateItemRequest struct {
    // description is Description of item.
    Description string `json:"description,omitempty"`
    // weight is Weight of item.
    Weight float64 `json:"weight,omitempty"`
    // pre_order is Pre Order setting.
    PreOrder PreOrder `json:"pre_order,omitempty"`
    // item_name is Item name.
    ItemName string `json:"item_name,omitempty"`
    // attribute_list is Item attributes.
    AttributeList AttributeList `json:"attribute_list,omitempty"`
    // image is Images of item.
    Image Image `json:"image,omitempty"`
    // item_sku is SKU tag for item.
    ItemSku string `json:"item_sku,omitempty"`
    // item_status is Item status, could be UNLIST or NORMAL.
    ItemStatus string `json:"item_status,omitempty"`
    // logistic_info is Logistic channel setting.
    LogisticInfo LogisticInfo `json:"logistic_info,omitempty"`
    // wholesale is Wholesale setting.
    Wholesale Wholesale `json:"wholesale,omitempty"`
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // category_id is ID of category.
    CategoryID int `json:"category_id,omitempty"`
    // dimension is Dimension of item.
    Dimension Dimension `json:"dimension,omitempty"`
    // condition is Condition of item, could be NEW or USED.
    Condition string `json:"condition,omitempty"`
    // video_upload_id is Video upload ID returned from video uploading API.
    VideoUploadID string `json:"video_upload_id,omitempty"`
    // brand is 
    Brand Brand `json:"brand,omitempty"`
    // item_dangerous is This field is only applicable for local sellers in Indonesia and Malaysia. Use this field to identify whether a product is a dangerous product. 0 for non-dangerous product and 1 for dangerous product. For more information, please visit the market's respective Seller Education Hub.
    ItemDangerous int `json:"item_dangerous,omitempty"`
    // tax_info is Tax information
    TaxInfo TaxInfo `json:"tax_info,omitempty"`
    // complaint_policy is Complaint Policy for item. Only required for local PL sellers, ignored otherwise.
    ComplaintPolicy ComplaintPolicy `json:"complaint_policy,omitempty"`
    // description_info is New description field. Only whitelist sellers can use it. If you use the field, please upload the description_type=extended otherwise api will return error. If you don't use this field, you don't need to upload the description_type or upload description_type=normal
    DescriptionInfo DescriptionInfo `json:"description_info,omitempty"`
    // description_type is <p>description_type (normal , extended). If you want to use extended_description or change description type ,this field must be inputed</p>
    DescriptionType string `json:"description_type,omitempty"`
}
//=======================================================
// ProductUpdateItemResponse
//=======================================================
type ProductUpdateItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductUpdateItem `json:"response,omitempty"`
}
//=======================================================
// ProductDeleteItemRequest
//=======================================================
type ProductDeleteItemRequest struct {
    // item_id is The identity of product item.
    ItemID int `json:"item_id"`
}
//=======================================================
// ProductDeleteItemResponse
//=======================================================
type ProductDeleteItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - OptionList
//=======================================================
type OptionList struct {
// option is Option name
Option string `json:"option,omitempty"`
// image is Option image
Image Image `json:"image,omitempty"`
}


//=======================================================
// Object Raw Type - TierVariation
//=======================================================
type TierVariation struct {
// name is Tier variation name
Name string `json:"name,omitempty"`
// option_list is Tier variation option info list
OptionList OptionList `json:"option_list"`
}


//=======================================================
// Object Raw Type - Model
//=======================================================
type Model struct {
// tier_index is Tier index of this model
TierIndex []int `json:"tier_index,omitempty"`
// normal_stock is Normal stock of this model
NormalStock int `json:"normal_stock,omitempty"`
// original_price is <p>Original price of this model.</p><p><b><font color="#c24f4a">For CO local VAT responsible seller：</font></b>Please remember the price you set in here must be VAT inclusive. If you have any doubts on how to calculate VAT for your product please refer to the Seller Education Hub（https://seller.shopee.com.co/edu/article/13565）<br /></p>
OriginalPrice float64 `json:"original_price,omitempty"`
// model_sku is Seller SKU of this model, model_sku length information needs to be no more than 100 characters.
ModelSku string `json:"model_sku,omitempty"`
// seller_stock is <p>new stock info（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
SellerStock SellerStock `json:"seller_stock,omitempty"`
}


//=======================================================
// Object Raw Type - ProductInitTierVariation
//=======================================================
type ProductInitTierVariation struct {
// item_id is ID of item
ItemID int `json:"item_id,omitempty"`
// tier_variation is Variations of item
TierVariation TierVariation `json:"tier_variation"`
// model is 
Model Model `json:"model"`
}
//=======================================================
// ProductInitTierVariationRequest
//=======================================================
type ProductInitTierVariationRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation TierVariation `json:"tier_variation"`
    // model is Model info list, model number at most 50
    Model Model `json:"model"`
}
//=======================================================
// ProductInitTierVariationResponse
//=======================================================
type ProductInitTierVariationResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductInitTierVariation `json:"response"`
}
//=======================================================
// ProductUpdateTierVariationRequest
//=======================================================
type ProductUpdateTierVariationRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation TierVariation `json:"tier_variation"`
}
//=======================================================
// ProductUpdateTierVariationResponse
//=======================================================
type ProductUpdateTierVariationResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - ProductGetModelList
//=======================================================
type ProductGetModelList struct {
// tier_variation is Variation config of item.
TierVariation TierVariation `json:"tier_variation"`
// model is Model list.
Model Model `json:"model"`
}
//=======================================================
// ProductGetModelListRequest
//=======================================================
type ProductGetModelListRequest struct {
    // item_id is The ID of the item
    ItemID int `json:"item_id"`
}
//=======================================================
// ProductGetModelListResponse
//=======================================================
type ProductGetModelListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetModelList `json:"response"`
}


//=======================================================
// Object Raw Type - ModelList
//=======================================================
type ModelList struct {
// tier_index is Tier index of model
TierIndex []int `json:"tier_index,omitempty"`
// normal_stock is Normal stock for model
NormalStock int `json:"normal_stock,omitempty"`
// original_price is Normal stock for price
OriginalPrice float64 `json:"original_price,omitempty"`
// model_sku is Seller sku, model_sku length information needs to be no more than 100 characters.
ModelSku string `json:"model_sku,omitempty"`
// seller_stock is <p>new stock info for model（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
SellerStock SellerStock `json:"seller_stock,omitempty"`
}


//=======================================================
// Object Raw Type - ProductAddModel
//=======================================================
type ProductAddModel struct {
// model is 
Model Model `json:"model"`
}
//=======================================================
// ProductAddModelRequest
//=======================================================
type ProductAddModelRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // model_list is Model list
    ModelList ModelList `json:"model_list"`
}
//=======================================================
// ProductAddModelResponse
//=======================================================
type ProductAddModelResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductAddModel `json:"response"`
}
//=======================================================
// ProductUpdateModelRequest
//=======================================================
type ProductUpdateModelRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // model is Length should be between 1 to 50
    Model Model `json:"model"`
}
//=======================================================
// ProductUpdateModelResponse
//=======================================================
type ProductUpdateModelResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// ProductDeleteModelRequest
//=======================================================
type ProductDeleteModelRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // model_id is ID of model.
    ModelID int `json:"model_id"`
}
//=======================================================
// ProductDeleteModelResponse
//=======================================================
type ProductDeleteModelResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - ProductSupportSizeChart
//=======================================================
type ProductSupportSizeChart struct {
// support_size_chart is Can sizechart be set for this category.
SupportSizeChart bool `json:"support_size_chart,omitempty"`
}
//=======================================================
// ProductSupportSizeChartRequest
//=======================================================
type ProductSupportSizeChartRequest struct {
    // category_id is Category ID
    CategoryID int `json:"category_id"`
}
//=======================================================
// ProductSupportSizeChartResponse
//=======================================================
type ProductSupportSizeChartResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductSupportSizeChart `json:"response"`
}
//=======================================================
// ProductUpdateSizeChartRequest
//=======================================================
type ProductUpdateSizeChartRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // size_chart is ID of size chart image
    SizeChart string `json:"size_chart"`
}
//=======================================================
// ProductUpdateSizeChartResponse
//=======================================================
type ProductUpdateSizeChartResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - FailureList
//=======================================================
type FailureList struct {
// item_id is Failed item id
ItemID int `json:"item_id,omitempty"`
// failed_reason is Failed reason
FailedReason string `json:"failed_reason,omitempty"`
}


//=======================================================
// Object Raw Type - SuccessList
//=======================================================
type SuccessList struct {
// item_id is Success item id 
ItemID int `json:"item_id,omitempty"`
// unlist is Whether the item is unlisted
Unlist bool `json:"unlist,omitempty"`
}


//=======================================================
// Object Raw Type - ProductUnlistItem
//=======================================================
type ProductUnlistItem struct {
// failure_list is 
FailureList FailureList `json:"failure_list"`
// success_list is 
SuccessList SuccessList `json:"success_list"`
}
//=======================================================
// ProductUnlistItemRequest
//=======================================================
type ProductUnlistItemRequest struct {
    // item_list is Length should be between 1 to 50.
    ItemList ItemList `json:"item_list"`
}
//=======================================================
// ProductUnlistItemResponse
//=======================================================
type ProductUnlistItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductUnlistItem `json:"response"`
}


//=======================================================
// Object Raw Type - PriceList
//=======================================================
type PriceList struct {
// model_id is 0 for no model item.
ModelID int `json:"model_id,omitempty"`
// original_price is <p>Original price for this model.</p><p><b><font color="#c24f4a">For CO local VAT responsible seller：</font></b>Please remember the price you set in here must be VAT inclusive. If you have any doubts on how to calculate VAT for your product please refer to the Seller Education Hub（https://seller.shopee.com.co/edu/article/13565）<br /></p><p><b><font color="#c24f4a">For&nbsp;SG/MY/BR/MX/PL/ES/AR seller:&nbsp;</font></b>Sellers can set the price with two decimal place,&nbsp;other regions can only set the price as an integer.<br /></p>
OriginalPrice float64 `json:"original_price,omitempty"`
}


//=======================================================
// Object Raw Type - ProductUpdatePrice
//=======================================================
type ProductUpdatePrice struct {
// failure_list is Fail model list.
FailureList FailureList `json:"failure_list"`
// success_list is Success model list.
SuccessList SuccessList `json:"success_list"`
}
//=======================================================
// ProductUpdatePriceRequest
//=======================================================
type ProductUpdatePriceRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // price_list is Length should be between 1 to 50.
    PriceList PriceList `json:"price_list"`
}
//=======================================================
// ProductUpdatePriceResponse
//=======================================================
type ProductUpdatePriceResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductUpdatePrice `json:"response"`
}


//=======================================================
// Object Raw Type - StockList
//=======================================================
type StockList struct {
// model_id is 0 for no model item.
ModelID int `json:"model_id,omitempty"`
// normal_stock is <p>Normal stock.</p><p><b><font color="#c24f4a">Please use the seller_stock field instead, we will deprecate this field in the future.</font></b><br /></p>
NormalStock int `json:"normal_stock,omitempty"`
// seller_stock is <p>new stock info（Please notice that stock(including Seller Stock and Shopee Stock) should be larger than or equal to real-time reserved stock）<br /></p>
SellerStock SellerStock `json:"seller_stock,omitempty"`
}


//=======================================================
// Object Raw Type - ProductUpdateStock
//=======================================================
type ProductUpdateStock struct {
// failure_list is Fail model list.
FailureList FailureList `json:"failure_list"`
// success_list is Success model list.
SuccessList SuccessList `json:"success_list"`
}
//=======================================================
// ProductUpdateStockRequest
//=======================================================
type ProductUpdateStockRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // stock_list is Length should be between 1 to 50.
    StockList StockList `json:"stock_list"`
}
//=======================================================
// ProductUpdateStockResponse
//=======================================================
type ProductUpdateStockResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductUpdateStock `json:"response"`
}


//=======================================================
// Object Raw Type - ProductBoostItem
//=======================================================
type ProductBoostItem struct {
// failure_list is 
FailureList FailureList `json:"failure_list"`
// success_list is 
SuccessList SuccessList `json:"success_list"`
}
//=======================================================
// ProductBoostItemRequest
//=======================================================
type ProductBoostItemRequest struct {
    // item_id_list is Shopee's unique identifier for an item, limit:[1,5]
    ItemIdList []int `json:"item_id_list"`
}
//=======================================================
// ProductBoostItemResponse
//=======================================================
type ProductBoostItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductBoostItem `json:"response"`
}


//=======================================================
// Object Raw Type - ProductGetBoostedList
//=======================================================
type ProductGetBoostedList struct {
// item_list is 
ItemList ItemList `json:"item_list"`
}
//=======================================================
// ProductGetBoostedListRequest
//=======================================================
type ProductGetBoostedListRequest struct {
}
//=======================================================
// ProductGetBoostedListResponse
//=======================================================
type ProductGetBoostedListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetBoostedList `json:"response"`
}


//=======================================================
// Object Raw Type - ProductGetItemPromotion
//=======================================================
type ProductGetItemPromotion struct {
// success_list is Success item promotion info.
SuccessList SuccessList `json:"success_list"`
// failure_list is Fail item promotion info.
FailureList FailureList `json:"failure_list"`
}
//=======================================================
// ProductGetItemPromotionRequest
//=======================================================
type ProductGetItemPromotionRequest struct {
    // item_id_list is Item ID list, can send 1 to 50 items.
    ItemIdList []int `json:"item_id_list"`
}
//=======================================================
// ProductGetItemPromotionResponse
//=======================================================
type ProductGetItemPromotionResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetItemPromotion `json:"response"`
}


//=======================================================
// Object Raw Type - SipItemPrice
//=======================================================
type SipItemPrice struct {
// model_id is 0 for no model item.
ModelID int `json:"model_id,omitempty"`
// sip_item_price is SIP item price.
SipItemPrice float64 `json:"sip_item_price,omitempty"`
}
//=======================================================
// ProductUpdateSipItemPriceRequest
//=======================================================
type ProductUpdateSipItemPriceRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // sip_item_price is 
    SipItemPrice SipItemPrice `json:"sip_item_price,omitempty"`
}
//=======================================================
// ProductUpdateSipItemPriceResponse
//=======================================================
type ProductUpdateSipItemPriceResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - ProductSearchItem
//=======================================================
type ProductSearchItem struct {
// item_id_list is List of  item ID.
ItemIdList []int `json:"item_id_list,omitempty"`
// total_count is Total num of items match condation.
TotalCount int `json:"total_count,omitempty"`
// next_offset is If has_next_page is true, this value need set to next request.offset
NextOffset string `json:"next_offset,omitempty"`
}
//=======================================================
// ProductSearchItemRequest
//=======================================================
type ProductSearchItemRequest struct {
    // offset is Specifies the starting entry of data to return in the current call. Default is empty. if data is more than one page, the offset can be some entry to start next call.
    Offset string `json:"offset,omitempty"`
    // page_size is the size of one page.
    PageSize int `json:"page_size"`
    // item_name is name of item.
    ItemName string `json:"item_name,omitempty"`
    // attribute_status is 1:get item lack of requires attribute.   2:get item lack of optional attribute.
    AttributeStatus int `json:"attribute_status,omitempty"`
}
//=======================================================
// ProductSearchItemResponse
//=======================================================
type ProductSearchItemResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductSearchItem `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - CommentReply
//=======================================================
type CommentReply struct {
// reply is The content of reply.
Reply string `json:"reply,omitempty"`
// hidden is The comment reply is hidden or not.
Hidden bool `json:"hidden,omitempty"`
}


//=======================================================
// Object Raw Type - ItemCommentList
//=======================================================
type ItemCommentList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSn string `json:"order_sn,omitempty"`
// comment_id is The identity of comment.
CommentID string `json:"comment_id,omitempty"`
// comment is The content of the comment.
Comment string `json:"comment,omitempty"`
// buyer_username is The username of the buyer who posted the comment.
BuyerUsername string `json:"buyer_username,omitempty"`
// item_id is The commented item's id
ItemID int `json:"item_id,omitempty"`
// model_id is <p>Shopee's unique identifier for a model of an item.It will only return 0 now.</p>
ModelID int `json:"model_id,omitempty"`
// rating_star is Buyer's rating for the item.
RatingStar int `json:"rating_star,omitempty"`
// editable is The editable status of the comment. The value may be one of  EXPIRED/EDITABLE/HAVE_EDIT_ONCE.
Editable string `json:"editable,omitempty"`
// hidden is The comment is hidden or not.
Hidden bool `json:"hidden,omitempty"`
// create_time is The create time of the comment.
CreateTime int `json:"create_time,omitempty"`
// comment_reply is The reply of the comment.
CommentReply CommentReply `json:"comment_reply"`
}


//=======================================================
// Object Raw Type - ProductGetComment
//=======================================================
type ProductGetComment struct {
// more is <p>This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments. But only respond 500 comments at most through OpenAPI, if there are more than 500, this field "more" also respond "true".</p>
More bool `json:"more,omitempty"`
// item_comment_list is The comment data list of the items.
ItemCommentList ItemCommentList `json:"item_comment_list"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// ProductGetCommentRequest
//=======================================================
type ProductGetCommentRequest struct {
    // item_id is The identity of product item.
    ItemID int `json:"item_id,omitempty"`
    // comment_id is The identity of comment.
    CommentID int `json:"comment_id,omitempty"`
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. The limit of page_size if between 1 and 100.
    PageSize int `json:"page_size"`
}
//=======================================================
// ProductGetCommentResponse
//=======================================================
type ProductGetCommentResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response ProductGetComment `json:"response"`
}


//=======================================================
// Object Raw Type - CommentList
//=======================================================
type CommentList struct {
// comment_id is The identity of comment.
CommentID int `json:"comment_id,omitempty"`
// comment is The content of the comment.
Comment string `json:"comment,omitempty"`
}


//=======================================================
// Object Raw Type - ResultList
//=======================================================
type ResultList struct {
// comment_id is The identity of comment.
CommentID int `json:"comment_id,omitempty"`
// fail_error is Indicate error details if one element hit error.
FailError string `json:"fail_error,omitempty"`
// fail_message is Indicate error type if one element hit error.
FailMessage string `json:"fail_message,omitempty"`
}


//=======================================================
// Object Raw Type - ProductReplyComment
//=======================================================
type ProductReplyComment struct {
// result_list is The result list of the request comment list.
ResultList ResultList `json:"result_list"`
}
//=======================================================
// ProductReplyCommentRequest
//=======================================================
type ProductReplyCommentRequest struct {
    // comment_list is The list of comment. The limit is between 1 and 100.
    CommentList CommentList `json:"comment_list"`
}
//=======================================================
// ProductReplyCommentResponse
//=======================================================
type ProductReplyCommentResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response ProductReplyComment `json:"response"`
}


//=======================================================
// Object Raw Type - ProductCategoryRecommend
//=======================================================
type ProductCategoryRecommend struct {
// category_id is Shopee's unique identifier for a category.
CategoryID []int `json:"category_id,omitempty"`
}
//=======================================================
// ProductCategoryRecommendRequest
//=======================================================
type ProductCategoryRecommendRequest struct {
    // item_name is name of item
    ItemName string `json:"item_name"`
}
//=======================================================
// ProductCategoryRecommendResponse
//=======================================================
type ProductCategoryRecommendResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductCategoryRecommend `json:"response"`
}


//=======================================================
// Object Raw Type - ProductImage
//=======================================================
type ProductImage struct {
// image_id_list is Image Id of product image for this brand, max input num of file = 10 ,each file's length<=498. ID market is optional.	
ImageIdList []string `json:"image_id_list,omitempty"`
}


//=======================================================
// Object Raw Type - ProductRegisterBrand
//=======================================================
type ProductRegisterBrand struct {
// brand_id is The identity of brand.
BrandID int `json:"brand_id,omitempty"`
// original_brand_name is Brand name
OriginalBrandName string `json:"original_brand_name,omitempty"`
}
//=======================================================
// ProductRegisterBrandRequest
//=======================================================
type ProductRegisterBrandRequest struct {
    // original_brand_name is Brand name, length<=254.
    OriginalBrandName string `json:"original_brand_name"`
    // category_list is Category_id list for this brand, please input category in L1 or L2. Max input num of category_id is 50.
    CategoryList []int `json:"category_list"`
    // product_image is 
    ProductImage ProductImage `json:"product_image"`
    // app_logo_image_id is Image_id  of logo for  app client,please input hashcode of this picture.
    AppLogoImageID string `json:"app_logo_image_id,omitempty"`
    // brand_website is Official website of brand, length<=254.
    BrandWebsite string `json:"brand_website,omitempty"`
    // brand_description is The description of this brand, can input the information, length<=254.
    BrandDescription string `json:"brand_description,omitempty"`
    // additional_information is Additional notes or comment can seller can add, length<=254.
    AdditionalInformation string `json:"additional_information,omitempty"`
    // pc_logo_image_id is Image_id  of logo for  pc client,please input hashcode of this picture.
    PcLogoImageID string `json:"pc_logo_image_id,omitempty"`
    // brand_country is origin country of brand.
    BrandCountry string `json:"brand_country"`
}
//=======================================================
// ProductRegisterBrandResponse
//=======================================================
type ProductRegisterBrandResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductRegisterBrand `json:"response"`
}


//=======================================================
// Object Raw Type - ProductGetRecommendAttribute
//=======================================================
type ProductGetRecommendAttribute struct {
// attribute_list is Attribute info list.
AttributeList AttributeList `json:"attribute_list"`
}
//=======================================================
// ProductGetRecommendAttributeRequest
//=======================================================
type ProductGetRecommendAttributeRequest struct {
    // item_name is name of item
    ItemName string `json:"item_name"`
    // cover_image_id is Cover image id of item
    CoverImageID int `json:"cover_image_id,omitempty"`
    // category_id is ID of category
    CategoryID int `json:"category_id"`
}
//=======================================================
// ProductGetRecommendAttributeResponse
//=======================================================
type ProductGetRecommendAttributeResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ProductGetRecommendAttribute `json:"response"`
}