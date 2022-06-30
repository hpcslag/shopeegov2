package shopeego
//=======================================================
// PushGetPushConfigRequest
//=======================================================
type PushGetPushConfigRequest struct {
}
//=======================================================
// PushGetPushConfigResponse
//=======================================================
type PushGetPushConfigResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - PushConfig
//=======================================================
type PushConfig struct {
// order_status is 0 stands for off and 1 stands for on.
OrderStatus int `json:"order_status,omitempty"`
// order_tracking_no is 0 stands for off and 1 stands for on.
OrderTrackingNo int `json:"order_tracking_no,omitempty"`
// shop_update is 0 stands for off and 1 stands for on.
ShopUpdate int `json:"shop_update,omitempty"`
// banned_item is 0 stands for off and 1 stands for on.
BannedItem int `json:"banned_item,omitempty"`
// item_promotion is 0 stands for off and 1 stands for on.
ItemPromotion int `json:"item_promotion,omitempty"`
// reserved_stock_change is 0 stands for off and 1 stands for on.
ReservedStockChange int `json:"reserved_stock_change,omitempty"`
// brand_register_result is 0 stands for off and 1 stands for on.
BrandRegisterResult int `json:"brand_register_result,omitempty"`
// promotion_update is 0 stands for off and 1 stands for on.
PromotionUpdate int `json:"promotion_update,omitempty"`
// webchat_update is 0 stands for off and 1 stands for on.
WebchatUpdate int `json:"webchat_update,omitempty"`
// open_api_authorization_expiry is 0 stands for off and 1 stands for on.
OpenApiAuthorizationExpiry int `json:"open_api_authorization_expiry,omitempty"`
// shop_authorization_canceled_push is <p>0 stands for off and 1 stands for on.<br /></p>
ShopAuthorizationCanceledPush int `json:"shop_authorization_canceled_push,omitempty"`
// shop_authorization_push is <p>0 stands for off and 1 stands for on.<br /></p>
ShopAuthorizationPush int `json:"shop_authorization_push,omitempty"`
}
//=======================================================
// PushSetPushConfigRequest
//=======================================================
type PushSetPushConfigRequest struct {
    // callback_url is The callback url of push mechanism.
    CallbackUrl string `json:"callback_url,omitempty"`
    // push_config is Detail configuration of push mechanism.
    PushConfig PushConfig `json:"push_config,omitempty"`
    // blocked_shop_id is Use this filed to set shops that need to be blocked.
    BlockedShopID []int `json:"blocked_shop_id,omitempty"`
}
//=======================================================
// PushSetPushConfigResponse
//=======================================================
type PushSetPushConfigResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}