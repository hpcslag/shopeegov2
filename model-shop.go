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
    // 通用的 Response 回傳參數
    V2UnityResponse

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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The information about shop logo, shop name, shop description.
    Response string `json:"response,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is If update successfully, the information is about shop logo, shop name, shop description.
    Response string `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ShopGetWarehouseDetail
//=======================================================
type ShopGetWarehouseDetail struct {
// warehouse_id is <p>Warehouse address identifier. It should be unique for every warehouse address<br /></p>
WarehouseID int `json:"warehouse_id,omitempty"`
// warehouse_name is <p>The warehouse name filled in when creating the warehouse address</p>
WarehouseName string `json:"warehouse_name,omitempty"`
// location_id is <p>Location identifier for stocks. Different location_ids represent that your addresses are in different item stocks</p>
LocationID string `json:"location_id,omitempty"`
// country is <p>Country of your warehouse address</p>
Country string `json:"country,omitempty"`
// state is <p>State of your warehouse address<br /></p>
State string `json:"state,omitempty"`
// city is <p>City of your warehouse address<br /></p>
City string `json:"city,omitempty"`
// district is <p>Distinct of your warehouse address<br /></p>
District string `json:"district,omitempty"`
// town is <p>Town of your warehouse address<br /></p>
Town string `json:"town,omitempty"`
// address is <p>Detail address of your warehouse address</p>
Address string `json:"address,omitempty"`
// zipcode is <p>Zipcode of your warehouse address<br /></p>
Zipcode string `json:"zipcode,omitempty"`
// state_code is <p>State code of your warehouse address<br /></p>
StateCode string `json:"state_code,omitempty"`
// holiday_mode_state is <p>The holiday mode state of your address.<br />0: not in holiday mode</p><p>1: holiday mode active</p><p>2: holiday mode is turning of</p><p>3: holiday mode is turning on</p>
HolidayModeState int `json:"holiday_mode_state,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ShopGetWarehouseDetail `json:"response,omitempty"`
}