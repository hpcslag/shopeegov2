package shopeego

import (
	"io/fs"
	"io/ioutil"
)

// these function replace original function
func (s *ShopeeClient) LogisticsDownloadShippingDocument(saveFilePath string) func(req *LogisticsDownloadShippingDocumentRequest) (err error) {
	return func(req *LogisticsDownloadShippingDocumentRequest) (err error) {

		b, err := s.postDownloadFile("LogisticsDownloadShippingDocument", req)
		if err != nil {
			return nil
		}

		err = ioutil.WriteFile(saveFilePath, b, fs.ModePerm|fs.ModeAppend)

		return nil
	}
}

//=======================================================
// LogisticsShipOrderRequest
//=======================================================
type LogisticsShipOrderRequest struct {
	V2RequestAuthenticationParams

	// order_sn is Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	PackageNumber string `json:"package_number,omitempty"`
	// pickup is Required parameter ONLY if get_shipping_parameter returns "pickup" under "info_needed". Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup *Pickup `json:"pickup,omitempty"`
	// dropoff is Required parameter ONLY if get_shipping_parameter returns "dropoff" under "info_needed". Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	Dropoff *Dropoff `json:"dropoff,omitempty"`
	// non_integrated is Optional parameter when get_shipping_parameter returns "non-integrated" under "info_needed".
	NonIntegrated *NonIntegrated `json:"non_integrated,omitempty"`
}

//=======================================================
// LogisticsBatchShipOrderRequest
//=======================================================
type LogisticsBatchShipOrderRequest struct {
	V2RequestAuthenticationParams

	// order_list is The list of order.
	OrderList []OrderList `json:"order_list"`
	// pickup is Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup *Pickup `json:"pickup,omitempty"`
	// dropoff is Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	Dropoff *Dropoff `json:"dropoff,omitempty"`
	// non_integrated is Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
	NonIntegrated *NonIntegrated `json:"non_integrated,omitempty"`
}

//=======================================================
// LogisticsCreateShippingDocumentRequest
//=======================================================
type LogisticsCreateShippingDocumentRequest struct {
	V2RequestAuthenticationParams

	// order_list is The list of order you want to create shipping document. limit [1, 50]
	OrderList []LogisticsCreateShippingDocumentOrderList `json:"order_list"`
}

type LogisticsCreateShippingDocumentOrderList struct {
	// order_sn is Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn,omitempty"`
	// package_number is Shopee's unique identifier for the package under an order. You should't fill the field with empty string when there is't a package number.
	PackageNumber string `json:"package_number,omitempty"`
	// The tracking number of order. Required except for the channel allow print before arrange shipment.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL,NORMAL_JOB_AIR_WAYBILL,THERMAL_JOB_AIR_WAYBILL
	ShippingDocumentType string `json:"shipping_document_type,omitempty"`
}
