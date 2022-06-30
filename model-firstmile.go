package shopeego
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
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
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// FirstMileBindFirstMileTrackingNumberResponse
//=======================================================
type FirstMileBindFirstMileTrackingNumberResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning []interface{} `json:"warning,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}
//=======================================================
// FirstMileUnbindFirstMileTrackingNumberRequest
//=======================================================
type FirstMileUnbindFirstMileTrackingNumberRequest struct {
    // first_mile_tracking_number is The identifier for an API request for error tracking.
    FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
    // order_list is The list of order info you want to unbind from the given first mile tracking number.
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// FirstMileUnbindFirstMileTrackingNumberResponse
//=======================================================
type FirstMileUnbindFirstMileTrackingNumberResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning []interface{} `json:"warning,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}