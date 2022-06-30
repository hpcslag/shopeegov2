package shopeego
//=======================================================
// ShopGetShopInfoRequest
//=======================================================
type ShopGetShopInfoRequest struct {
}
//=======================================================
// ShopGetShopInfoResponse
//=======================================================
type ShopGetShopInfoResponse struct {
    // shop_name is Name of the shop.
    ShopName string `json:"shop_name,omitempty"`
    // region is Shop's area.
    Region string `json:"region,omitempty"`
    // status is <p>Applicable status: <b>BANNED</b>, <b>FROZEN</b>, <b>NORMAL</b>.</p>
    Status string `json:"status,omitempty"`
    // sip_affi_shops is SIP affiliate shops info list.If you request for SIP primary shop,this field will be returned, if you request for SIP affiliate shop,this field won't be returned
    SipAffiShops []interface{} `json:"sip_affi_shops,omitempty"`
    // is_cb is Use this filed to indicate whether the shop is a cross-border shop.
    IsCb bool `json:"is_cb,omitempty"`
    // is_cnsc is Use this filed to indicate whether the shop is upgraded to CNSC.
    IsCnsc bool `json:"is_cnsc,omitempty"`
    // shop_cbsc is <p>Use this filed to indicate whether this shop is upgraded to cbsc:&nbsp;<b>CNSC/KRSC/UNUPGRADED</b><br /></p>
    ShopCbsc string `json:"shop_cbsc,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // auth_time is The timestamp when the shop was authorized to the partner.
    AuthTime int `json:"auth_time,omitempty"`
    // expire_time is Use this field to indicate the expiration date for shop authorization.
    ExpireTime int `json:"expire_time,omitempty"`
    // is_sip is This filed will return "true" when SIP primary shop or affiliate shop calls
    IsSip bool `json:"is_sip,omitempty"`
}
//=======================================================
// ShopGetProfileRequest
//=======================================================
type ShopGetProfileRequest struct {
}
//=======================================================
// ShopGetProfileResponse
//=======================================================
type ShopGetProfileResponse struct {
    // message is Messages about the error and it's cause if there're errors. It will be just return the updated information if there's no error.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is The information about shop logo, shop name, shop description.
    Response string `json:"response,omitempty"`
    // error is The error key if there's error.
    Error string `json:"error,omitempty"`
}
//=======================================================
// ShopUpdateProfileRequest
//=======================================================
type ShopUpdateProfileRequest struct {
    // shop_name is The new shop name
    ShopName string `json:"shop_name,omitempty"`
    // shop_logo is The new shop logo url. Recommend to use images 
    ShopLogo string `json:"shop_logo,omitempty"`
    // description is The new shop description.
    Description string `json:"description,omitempty"`
}
//=======================================================
// ShopUpdateProfileResponse
//=======================================================
type ShopUpdateProfileResponse struct {
    // message is Messages about the error and it's cause if there're errors. It will be just return the updated information if there's no error.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is If update successfully, the information is about shop logo, shop name, shop description.
    Response string `json:"response,omitempty"`
    // error is The error key if there's error.
    Error string `json:"error,omitempty"`
}
//=======================================================
// ShopGetWarehouseDetailRequest
//=======================================================
type ShopGetWarehouseDetailRequest struct {
}
//=======================================================
// ShopGetWarehouseDetailResponse
//=======================================================
type ShopGetWarehouseDetailResponse struct {
    // request_id is <p>The identifier for an API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
    // error is <p>The error key if there's error.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Messages about the error and it's cause if there're errors. It will be just return the updated information if there's no error.<br /></p>
    Message string `json:"message,omitempty"`
    // response is 
    Response []interface{} `json:"response,omitempty"`
}