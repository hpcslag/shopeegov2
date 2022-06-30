package shopeego


//=======================================================
// Object Raw Type - Response
//=======================================================
type Response struct {
// category_list is 
CategoryList []interface{} `json:"category_list,omitempty"`
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
    // error is  Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
// ProductGetItemLimitRequest
//=======================================================
type ProductGetItemLimitRequest struct {
}
//=======================================================
// ProductGetItemLimitResponse
//=======================================================
type ProductGetItemLimitResponse struct {
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is  Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit error. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}


//=======================================================
// Object Raw Type - Dimension
//=======================================================
type Dimension struct {
// package_height is <p>Package height, unit is cm</p>
PackageHeight int `json:"package_height,omitempty"`
// package_length is <p>Package height, unit is cm</p>
PackageLength int `json:"package_length,omitempty"`
// package_width is <p>Package height, unit is cm</p>
PackageWidth int `json:"package_width,omitempty"`
}


//=======================================================
// Object Raw Type - Image
//=======================================================
type Image struct {
// image_id_list is ID of image
ImageIdList []string `json:"image_id_list,omitempty"`
}


//=======================================================
// Object Raw Type - PreOrder
//=======================================================
type PreOrder struct {
// is_pre_order is Whether item is pre order
IsPreOrder bool `json:"is_pre_order,omitempty"`
// days_to_ship is The guaranteed days to ship orders. Please get the days_to_ship range from get_dts_limit api
DaysToShip int `json:"days_to_ship,omitempty"`
}


//=======================================================
// Object Raw Type - Brand
//=======================================================
type Brand struct {
// brand_id is Id of brand.
BrandID int `json:"brand_id,omitempty"`
}


//=======================================================
// Object Raw Type - TaxInfo
//=======================================================
type TaxInfo struct {
// ncm is Mercosur Common Nomenclature, it is a convention between Mercosur member countries to easily recognize goods, services and productive factors negotiated among themselves. (BR region)
Ncm string `json:"ncm,omitempty"`
// same_state_cfop is Tax Code of Operations and Installments for orders that seller and buyer are in the same state. It identifies a specific operation by category at the time of issuing the invoice.(BR region)
SameStateCfop string `json:"same_state_cfop,omitempty"`
// diff_state_cfop is Tax Code of Operations and Installments for orders that seller and buyer are in different states. It identifies a specific operation by category at the time of issuing the invoice.(BR region)
DiffStateCfop string `json:"diff_state_cfop,omitempty"`
// csosn is Code of Operation Status – Simples Nacional, code for company operations to identify the origin of the goods and the taxation regime of the operations.(BR region)
Csosn string `json:"csosn,omitempty"`
// origin is Product source, domestic or foreig (BR region)
Origin string `json:"origin,omitempty"`
// cest is (BR region)
Cest string `json:"cest,omitempty"`
// measure_unit is (BR region)
MeasureUnit string `json:"measure_unit,omitempty"`
// invoice_option is Value shuold be one of NO_INVOICES VAT_MARGIN_SCHEME_INVOICES VAT_INVOICES NON_VAT_INVOICES and if value is NON_VAT_INVOICE vat_rate should be null (PL region)
InvoiceOption string `json:"invoice_option,omitempty"`
// vat_rate is Value should be one of 0% 5% 8% 23% NO_VAT_RATE (PL region)
VatRate string `json:"vat_rate,omitempty"`
// hs_code is HS Code (Only for IN region)
HsCode string `json:"hs_code,omitempty"`
// tax_code is Tax Code (Only for IN region)
TaxCode string `json:"tax_code,omitempty"`
}


//=======================================================
// Object Raw Type - ComplaintPolicy
//=======================================================
type ComplaintPolicy struct {
// warranty_time is Value should be in one of ONE_YEAR TWO_YEARS OVER_TWO_YEARS.
WarrantyTime string `json:"warranty_time,omitempty"`
// exclude_entrepreneur_warranty is Whether to exclude warranty complaints for entrepreneurs.If True means "I exclude warranty complaints for entrepreneur"
ExcludeEntrepreneurWarranty bool `json:"exclude_entrepreneur_warranty,omitempty"`
// complaint_address_id is  Address for complaint. Fetch available addresses using v2.logistics.get_address_list, and use address_id  returned from it.
ComplaintAddressID int `json:"complaint_address_id,omitempty"`
// additional_information is Additional information for warranty claim. Should be less than 1000 characters.
AdditionalInformation string `json:"additional_information,omitempty"`
}


//=======================================================
// Object Raw Type - ExtendedDescription
//=======================================================
type ExtendedDescription struct {
// field_list is Field of extended description.
FieldList []interface{} `json:"field_list,omitempty"`
}


//=======================================================
// Object Raw Type - DescriptionInfo
//=======================================================
type DescriptionInfo struct {
// extended_description is If description_type is extended , Description information should be set by this field.
ExtendedDescription ExtendedDescription `json:"extended_description,omitempty"`
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
    LogisticInfo []interface{} `json:"logistic_info"`
    // attribute_list is <p>This field is optional(expect Indonesia) depending on the specific attribute under different categories. Should call get_attribute api to get attribute first. Must contain all all mandatory attribute.</p>
    AttributeList []interface{} `json:"attribute_list,omitempty"`
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
    Wholesale []interface{} `json:"wholesale,omitempty"`
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
    SellerStock []interface{} `json:"seller_stock,omitempty"`
}
//=======================================================
// ProductAddItemResponse
//=======================================================
type ProductAddItemResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
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
    AttributeList []interface{} `json:"attribute_list,omitempty"`
    // image is Images of item.
    Image Image `json:"image,omitempty"`
    // item_sku is SKU tag for item.
    ItemSku string `json:"item_sku,omitempty"`
    // item_status is Item status, could be UNLIST or NORMAL.
    ItemStatus string `json:"item_status,omitempty"`
    // logistic_info is Logistic channel setting.
    LogisticInfo []interface{} `json:"logistic_info,omitempty"`
    // wholesale is Wholesale setting.
    Wholesale []interface{} `json:"wholesale,omitempty"`
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
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// ProductInitTierVariationRequest
//=======================================================
type ProductInitTierVariationRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []interface{} `json:"tier_variation"`
    // model is Model info list, model number at most 50
    Model []interface{} `json:"model"`
}
//=======================================================
// ProductInitTierVariationResponse
//=======================================================
type ProductInitTierVariationResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ProductUpdateTierVariationRequest
//=======================================================
type ProductUpdateTierVariationRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // tier_variation is Tier variation info list.If you define a one-tier structure, the maximum number of options cannot exceed 50. If you define a two-tier structure, the number of options multiplied by the two tiers cannot exceed 50.
    TierVariation []interface{} `json:"tier_variation"`
}
//=======================================================
// ProductUpdateTierVariationResponse
//=======================================================
type ProductUpdateTierVariationResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
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
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
}
//=======================================================
// ProductAddModelRequest
//=======================================================
type ProductAddModelRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // model_list is Model list
    ModelList []interface{} `json:"model_list"`
}
//=======================================================
// ProductAddModelResponse
//=======================================================
type ProductAddModelResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is 
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ProductUpdateModelRequest
//=======================================================
type ProductUpdateModelRequest struct {
    // item_id is ID of item
    ItemID int `json:"item_id"`
    // model is Length should be between 1 to 50
    Model []interface{} `json:"model"`
}
//=======================================================
// ProductUpdateModelResponse
//=======================================================
type ProductUpdateModelResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
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
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
}
//=======================================================
// ProductUnlistItemRequest
//=======================================================
type ProductUnlistItemRequest struct {
    // item_list is Length should be between 1 to 50.
    ItemList []interface{} `json:"item_list"`
}
//=======================================================
// ProductUnlistItemResponse
//=======================================================
type ProductUnlistItemResponse struct {
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
// ProductUpdatePriceRequest
//=======================================================
type ProductUpdatePriceRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // price_list is Length should be between 1 to 50.
    PriceList []interface{} `json:"price_list"`
}
//=======================================================
// ProductUpdatePriceResponse
//=======================================================
type ProductUpdatePriceResponse struct {
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
// ProductUpdateStockRequest
//=======================================================
type ProductUpdateStockRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // stock_list is Length should be between 1 to 50.
    StockList []interface{} `json:"stock_list"`
}
//=======================================================
// ProductUpdateStockResponse
//=======================================================
type ProductUpdateStockResponse struct {
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
// ProductGetBoostedListRequest
//=======================================================
type ProductGetBoostedListRequest struct {
}
//=======================================================
// ProductGetBoostedListResponse
//=======================================================
type ProductGetBoostedListResponse struct {
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Waring message.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ProductUpdateSipItemPriceRequest
//=======================================================
type ProductUpdateSipItemPriceRequest struct {
    // item_id is ID of item.
    ItemID int `json:"item_id"`
    // sip_item_price is 
    SipItemPrice []interface{} `json:"sip_item_price,omitempty"`
}
//=======================================================
// ProductUpdateSipItemPriceResponse
//=======================================================
type ProductUpdateSipItemPriceResponse struct {
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
    // error is  Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is  Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is  Indicate waring details if hit waring. Empty if no waring happened.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}
//=======================================================
// ProductReplyCommentRequest
//=======================================================
type ProductReplyCommentRequest struct {
    // comment_list is The list of comment. The limit is between 1 and 100.
    CommentList []interface{} `json:"comment_list"`
}
//=======================================================
// ProductReplyCommentResponse
//=======================================================
type ProductReplyCommentResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
    // warning is Indicate warning message you should take care.
    Warning []string `json:"warning,omitempty"`
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
// Object Raw Type - ProductImage
//=======================================================
type ProductImage struct {
// image_id_list is Image Id of product image for this brand, max input num of file = 10 ,each file's length<=498. ID market is optional.	
ImageIdList []string `json:"image_id_list,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
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