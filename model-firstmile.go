package shopeego


//=======================================================
// Object Raw Type - FirstMileGetUnbindOrderListOrderList
//=======================================================
type FirstMileGetUnbindOrderListOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// logistics_status is The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
LogisticsStatus string `json:"logistics_status,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGetUnbindOrderList
//=======================================================
type FirstMileGetUnbindOrderList struct {
// more is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
More bool `json:"more,omitempty"`
// order_list is The result list of order you querying.
OrderList []FirstMileGetUnbindOrderListOrderList `json:"order_list"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// FirstMileGetUnbindOrderListRequest
//=======================================================
type FirstMileGetUnbindOrderListRequest struct {
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 100]
    PageSize int `json:"page_size,omitempty"`
    // response_optional_fields is Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them.  Available values: logistic_status,package_number. 
    ResponseOptionalFields string `json:"response_optional_fields,omitempty"`
}
//=======================================================
// FirstMileGetUnbindOrderListResponse
//=======================================================
type FirstMileGetUnbindOrderListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileGetUnbindOrderList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGetDetailOrderList
//=======================================================
type FirstMileGetDetailOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// sls_tracking_number is The tracking number of SLS for orders/forders.
SlsTrackingNumber string `json:"sls_tracking_number,omitempty"`
// pick_up_done is Use this filed to indicate whether the order has been picked up by carrier.
PickUpDone bool `json:"pick_up_done,omitempty"`
// arrived_transit_warehouse is Use this filed to indicate whether the order has arrived at transit warehouse.
ArrivedTransitWarehouse bool `json:"arrived_transit_warehouse,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGetDetail
//=======================================================
type FirstMileGetDetail struct {
// logistics_channel_id is The identity of logistic channel.
LogisticsChannelID int `json:"logistics_channel_id,omitempty"`
// first_mile_tracking_number is The first-mile tracking number.
FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
// shipment_method is The shipment method for bound orders, should be pickup or dropoff.
ShipmentMethod string `json:"shipment_method,omitempty"`
// status is The logistics status for first-mile tracking number. Status could be: INIT,ORDER_CREATED,PICKED_UP,DELIVERED,ORDER_RECEIVED,CANCELING,CANCELED.
Status string `json:"status,omitempty"`
// declare_date is The specified delivery date.
DeclareDate string `json:"declare_date,omitempty"`
// more is This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
More bool `json:"more,omitempty"`
// order_list is The list of order.
OrderList []FirstMileGetDetailOrderList `json:"order_list"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// FirstMileGetDetailRequest
//=======================================================
type FirstMileGetDetailRequest struct {
    // first_mile_tracking_number is The first mile tracking number.
    FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
}
//=======================================================
// FirstMileGetDetailResponse
//=======================================================
type FirstMileGetDetailResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileGetDetail `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - SellerInfo
//=======================================================
type SellerInfo struct {
// address is The full address of the seller.
Address string `json:"address,omitempty"`
// name is Seller's name for the address.
Name string `json:"name,omitempty"`
// zipcode is Seller's postal code.
Zipcode string `json:"zipcode,omitempty"`
// region is Seller's location.
Region string `json:"region,omitempty"`
// phone is Seller's phone number.
Phone string `json:"phone,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGenerateFirstMileTrackingNumber
//=======================================================
type FirstMileGenerateFirstMileTrackingNumber struct {
// first_mile_tracking_number_list is The list of first mile tracking number that you generate
FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list,omitempty"`
}
//=======================================================
// FirstMileGenerateFirstMileTrackingNumberRequest
//=======================================================
type FirstMileGenerateFirstMileTrackingNumberRequest struct {
    // declare_date is This field is used for seller to specify the declare time.
    DeclareDate string `json:"declare_date"`
    // quantity is The number of first-mile tracking numbers generated. Up to 20 first-mile tracking numbers can be generated for one declaration day.
    Quantity int `json:"quantity,omitempty"`
    // seller_info is This object contains detailed breakdown for the seller address.
    SellerInfo SellerInfo `json:"seller_info"`
}
//=======================================================
// FirstMileGenerateFirstMileTrackingNumberResponse
//=======================================================
type FirstMileGenerateFirstMileTrackingNumberResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileGenerateFirstMileTrackingNumber `json:"response"`
}


//=======================================================
// Object Raw Type - FirstMileBindFirstMileTrackingNumberOrderList
//=======================================================
type FirstMileBindFirstMileTrackingNumberOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// fail_error is Indicate error type if one element hit error.
FailError string `json:"fail_error,omitempty"`
// fail_message is Indicate error details if one element hit error.
FailMessage string `json:"fail_message,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileBindFirstMileTrackingNumber
//=======================================================
type FirstMileBindFirstMileTrackingNumber struct {
// first_mile_tracking_number is The first mile tracking number
FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
// order_list is The list of orders.
OrderList []FirstMileBindFirstMileTrackingNumberOrderList `json:"order_list"`
}
//=======================================================
// FirstMileBindFirstMileTrackingNumberRequest
//=======================================================
type FirstMileBindFirstMileTrackingNumberRequest struct {
    // first_mile_tracking_number is The first-mile tracking number.
    FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
    // shipment_method is <p>The shipment method for bound orders, should be pickup, dropoff or self-deliver.</p>
    ShipmentMethod string `json:"shipment_method"`
    // region is Use this field to specify the region you want to ship parcel.
    Region string `json:"region"`
    // logistics_channel_id is The identity of first-mile logistic channel
    LogisticsChannelID int `json:"logistics_channel_id"`
    // volume is The volume of the parcel.
    Volume float64 `json:"volume,omitempty"`
    // weight is The weight of the parcel.
    Weight float64 `json:"weight,omitempty"`
    // width is The width of the parcel.
    Width float64 `json:"width,omitempty"`
    // length is The length of the parcel.
    Length float64 `json:"length,omitempty"`
    // height is The height of the parcel.
    Height float64 `json:"height,omitempty"`
    // order_list is The set of ordersn. You can specify up to 50 ordersns in this call.one fm_tn maximum number of total bind orders is 10000.
    OrderList []OrderList `json:"order_list"`
}
//=======================================================
// FirstMileBindFirstMileTrackingNumberResponse
//=======================================================
type FirstMileBindFirstMileTrackingNumberResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileBindFirstMileTrackingNumber `json:"response"`
}


//=======================================================
// Object Raw Type - FirstMileUnbindFirstMileTrackingNumberOrderList
//=======================================================
type FirstMileUnbindFirstMileTrackingNumberOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// fail_error is Indicate error type if one element hit error.
FailError string `json:"fail_error,omitempty"`
// fail_message is Indicate error details if one element hit error.
FailMessage string `json:"fail_message,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileUnbindFirstMileTrackingNumber
//=======================================================
type FirstMileUnbindFirstMileTrackingNumber struct {
// first_mile_tracking_number is The first mile tracking number.
FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
// order_list is The binding result list of each order.
OrderList []FirstMileUnbindFirstMileTrackingNumberOrderList `json:"order_list"`
}
//=======================================================
// FirstMileUnbindFirstMileTrackingNumberRequest
//=======================================================
type FirstMileUnbindFirstMileTrackingNumberRequest struct {
    // first_mile_tracking_number is The identifier for an API request for error tracking.
    FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
    // order_list is The list of order info you want to unbind from the given first mile tracking number.
    OrderList []OrderList `json:"order_list"`
}
//=======================================================
// FirstMileUnbindFirstMileTrackingNumberResponse
//=======================================================
type FirstMileUnbindFirstMileTrackingNumberResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileUnbindFirstMileTrackingNumber `json:"response"`
}


//=======================================================
// Object Raw Type - FirstMileGetTrackingNumberListFirstMileTrackingNumberList
//=======================================================
type FirstMileGetTrackingNumberListFirstMileTrackingNumberList struct {
// first_mile_tracking_number is The specified delivery date.	
FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
// status is The logistics status for bound orders.	
Status string `json:"status,omitempty"`
// declare_date is The first-mile tracking number.
DeclareDate string `json:"declare_date,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGetTrackingNumberList
//=======================================================
type FirstMileGetTrackingNumberList struct {
// more is This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
More bool `json:"more,omitempty"`
// first_mile_tracking_number_list is The first-mile tracking number.
FirstMileTrackingNumberList []FirstMileGetTrackingNumberListFirstMileTrackingNumberList `json:"first_mile_tracking_number_list"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// FirstMileGetTrackingNumberListRequest
//=======================================================
type FirstMileGetTrackingNumberListRequest struct {
    // from_date is The start time of declare_date.
    FromDate string `json:"from_date"`
    // to_date is The end time of declare_date.
    ToDate string `json:"to_date"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data. limit [1, 50]
    PageSize int `json:"page_size,omitempty"`
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
}
//=======================================================
// FirstMileGetTrackingNumberListResponse
//=======================================================
type FirstMileGetTrackingNumberListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileGetTrackingNumberList `json:"response,omitempty"`
}
//=======================================================
// FirstMileGetWaybillRequest
//=======================================================
type FirstMileGetWaybillRequest struct {
    // first_mile_tracking_number_list is The first mile tracking number that you want to print waybill.limit [1, 50]
    FirstMileTrackingNumberList []string `json:"first_mile_tracking_number_list"`
}
//=======================================================
// FirstMileGetWaybillResponse
//=======================================================
type FirstMileGetWaybillResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - FirstMileGetChannelListLogisticsChannelList
//=======================================================
type FirstMileGetChannelListLogisticsChannelList struct {
// logistics_channel_id is The identity of logistic channel.
LogisticsChannelID int `json:"logistics_channel_id,omitempty"`
// logistics_channel_name is The name of logistic channel.
LogisticsChannelName string `json:"logistics_channel_name,omitempty"`
// shipment_method is <p>The shipment method for bound orders.Available values: pickup, dropoff, self_deliver.</p>
ShipmentMethod string `json:"shipment_method,omitempty"`
}


//=======================================================
// Object Raw Type - FirstMileGetChannelList
//=======================================================
type FirstMileGetChannelList struct {
// logistics_channel_list is 
LogisticsChannelList []FirstMileGetChannelListLogisticsChannelList `json:"logistics_channel_list"`
}
//=======================================================
// FirstMileGetChannelListRequest
//=======================================================
type FirstMileGetChannelListRequest struct {
    // region is Use this field to specify the region you want to ship parcel. Available value: CN
    Region string `json:"region"`
}
//=======================================================
// FirstMileGetChannelListResponse
//=======================================================
type FirstMileGetChannelListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response FirstMileGetChannelList `json:"response"`
}