package shopeego
//=======================================================
// LogisticsGetShippingParameterRequest
//=======================================================
type LogisticsGetShippingParameterRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
}
//=======================================================
// LogisticsGetShippingParameterResponse
//=======================================================
type LogisticsGetShippingParameterResponse struct {
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
// LogisticsGetTrackingNumberRequest
//=======================================================
type LogisticsGetTrackingNumberRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there isn't a package number.
    PackageNumber string `json:"package_number,omitempty"`
    // response_optional_fields is Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them. Available values: plp_number, first_mile_tracking_number,last_mile_tracking_number
    ResponseOptionalFields string `json:"response_optional_fields,omitempty"`
}
//=======================================================
// LogisticsGetTrackingNumberResponse
//=======================================================
type LogisticsGetTrackingNumberResponse struct {
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
// Object Raw Type - Pickup
//=======================================================
type Pickup struct {
// address_id is The identity of address. Retrieved from v2.logistics.get_shipping_parameter.
AddressID int `json:"address_id,omitempty"`
// pickup_time_id is The pickup time id. Retrieved from v2.logistics.get_shipping_parameter.
PickupTimeID string `json:"pickup_time_id,omitempty"`
// tracking_number is Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.	
TrackingNumber string `json:"tracking_number,omitempty"`
}


//=======================================================
// Object Raw Type - Dropoff
//=======================================================
type Dropoff struct {
// branch_id is The identity of branch. 
BranchID int `json:"branch_id,omitempty"`
// sender_real_name is The real name of sender.	
SenderRealName string `json:"sender_real_name,omitempty"`
// tracking_number is Need input this field when "tracking_number" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.	
TrackingNumber string `json:"tracking_number,omitempty"`
// slug is  The selected 3PL partner to drop-off parcels with.
Slug string `json:"slug,omitempty"`
}


//=======================================================
// Object Raw Type - NonIntegrated
//=======================================================
type NonIntegrated struct {
// tracking_number is Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.	
TrackingNumber string `json:"tracking_number,omitempty"`
}
//=======================================================
// LogisticsShipOrderRequest
//=======================================================
type LogisticsShipOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
    PackageNumber string `json:"package_number,omitempty"`
    // pickup is Required parameter ONLY if get_shipping_parameter returns "pickup" under "info_needed". Developer should still include "pickup" field in the call even if "pickup" has empty value.	
    Pickup Pickup `json:"pickup,omitempty"`
    // dropoff is Required parameter ONLY if get_shipping_parameter returns "dropoff" under "info_needed". Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
    Dropoff Dropoff `json:"dropoff,omitempty"`
    // non_integrated is Optional parameter when get_shipping_parameter returns "non-integrated" under "info_needed".	
    NonIntegrated NonIntegrated `json:"non_integrated,omitempty"`
}
//=======================================================
// LogisticsShipOrderResponse
//=======================================================
type LogisticsShipOrderResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// LogisticsUpdateShippingOrderRequest
//=======================================================
type LogisticsUpdateShippingOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
    PackageNumber string `json:"package_number,omitempty"`
    // pickup is Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
    Pickup Pickup `json:"pickup"`
}
//=======================================================
// LogisticsUpdateShippingOrderResponse
//=======================================================
type LogisticsUpdateShippingOrderResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// LogisticsGetShippingDocumentParameterRequest
//=======================================================
type LogisticsGetShippingDocumentParameterRequest struct {
    // order_list is The list of orders you want to get. limit [1,50]
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// LogisticsGetShippingDocumentParameterResponse
//=======================================================
type LogisticsGetShippingDocumentParameterResponse struct {
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
// LogisticsCreateShippingDocumentRequest
//=======================================================
type LogisticsCreateShippingDocumentRequest struct {
    // order_list is The list of order you want to create shipping document. limit [1, 50]
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// LogisticsCreateShippingDocumentResponse
//=======================================================
type LogisticsCreateShippingDocumentResponse struct {
    // request_id is The identifier of the API request for error tracking.
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
// LogisticsGetShippingDocumentResultRequest
//=======================================================
type LogisticsGetShippingDocumentResultRequest struct {
    // order_list is The list of orders, limit [1,50]
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// LogisticsGetShippingDocumentResultResponse
//=======================================================
type LogisticsGetShippingDocumentResultResponse struct {
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
// LogisticsDownloadShippingDocumentRequest
//=======================================================
type LogisticsDownloadShippingDocumentRequest struct {
    // shipping_document_type is The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL,NORMAL_JOB_AIR_WAYBILL,THERMAL_JOB_AIR_WAYBILL
    ShippingDocumentType string `json:"shipping_document_type,omitempty"`
    // order_list is The list of orders you need to download it's shipping document.
    OrderList []interface{} `json:"order_list"`
}
//=======================================================
// LogisticsDownloadShippingDocumentResponse
//=======================================================
type LogisticsDownloadShippingDocumentResponse struct {
}
//=======================================================
// LogisticsGetShippingDocumentInfoRequest
//=======================================================
type LogisticsGetShippingDocumentInfoRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_number is Shopee's unique identifier for the package under an order. You shouldn't fill the field with empty string when there isn't a package number.
    PackageNumber string `json:"package_number,omitempty"`
}
//=======================================================
// LogisticsGetShippingDocumentInfoResponse
//=======================================================
type LogisticsGetShippingDocumentInfoResponse struct {
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
// LogisticsGetTrackingInfoRequest
//=======================================================
type LogisticsGetTrackingInfoRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
    PackageNumber string `json:"package_number,omitempty"`
}
//=======================================================
// LogisticsGetTrackingInfoResponse
//=======================================================
type LogisticsGetTrackingInfoResponse struct {
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
// LogisticsGetAddressListRequest
//=======================================================
type LogisticsGetAddressListRequest struct {
}
//=======================================================
// LogisticsGetAddressListResponse
//=======================================================
type LogisticsGetAddressListResponse struct {
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
// Object Raw Type - AddressTypeConfig
//=======================================================
type AddressTypeConfig struct {
// address_id is The identifier id of the address.
AddressID int `json:"address_id,omitempty"`
// address_type is The type of addres. Available values: DEFAULT_ADDRESS, PICKUP_ADDRESS, RETURN_ADDRESS
AddressType []string `json:"address_type,omitempty"`
}
//=======================================================
// LogisticsSetAddressConfigRequest
//=======================================================
type LogisticsSetAddressConfigRequest struct {
    // show_pickup_address is Definite show pickup address or not.
    ShowPickupAddress bool `json:"show_pickup_address,omitempty"`
    // address_type_config is The config of your shop addres.
    AddressTypeConfig AddressTypeConfig `json:"address_type_config,omitempty"`
}
//=======================================================
// LogisticsSetAddressConfigResponse
//=======================================================
type LogisticsSetAddressConfigResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// LogisticsDeleteAddressRequest
//=======================================================
type LogisticsDeleteAddressRequest struct {
    // address_id is The identity of address you want to delete.
    AddressID int `json:"address_id"`
}
//=======================================================
// LogisticsDeleteAddressResponse
//=======================================================
type LogisticsDeleteAddressResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
}
//=======================================================
// LogisticsGetChannelListRequest
//=======================================================
type LogisticsGetChannelListRequest struct {
}
//=======================================================
// LogisticsGetChannelListResponse
//=======================================================
type LogisticsGetChannelListResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// LogisticsUpdateChannelRequest
//=======================================================
type LogisticsUpdateChannelRequest struct {
    // logistics_channel_id is The identity of logistic channel.
    LogisticsChannelID int `json:"logistics_channel_id"`
    // enabled is Whether to enable this logistic channel.
    Enabled bool `json:"enabled,omitempty"`
    // preferred is Whether to make this logistic channel preferred. Indonestia logistics channel are not applicable.
    Preferred bool `json:"preferred,omitempty"`
    // cod_enabled is Whether to enable COD for this logistic channel. Only COD supported channels are applicable.
    CodEnabled bool `json:"cod_enabled,omitempty"`
}
//=======================================================
// LogisticsUpdateChannelResponse
//=======================================================
type LogisticsUpdateChannelResponse struct {
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
}
//=======================================================
// LogisticsBatchShipOrderRequest
//=======================================================
type LogisticsBatchShipOrderRequest struct {
    // order_list is The list of order.
    OrderList []interface{} `json:"order_list"`
    // pickup is Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.	
    Pickup Pickup `json:"pickup,omitempty"`
    // dropoff is Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
    Dropoff Dropoff `json:"dropoff,omitempty"`
    // non_integrated is Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".	
    NonIntegrated NonIntegrated `json:"non_integrated,omitempty"`
}
//=======================================================
// LogisticsBatchShipOrderResponse
//=======================================================
type LogisticsBatchShipOrderResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning []interface{} `json:"warning,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}