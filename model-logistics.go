package shopeego

//=======================================================
// Object Raw Type - InfoNeeded
//=======================================================
type InfoNeeded struct {
	// dropoff is Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', choose one to Init. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.Could contain 'slug'.If it contains 'slug', to return selected 3PL partner only for TW C2C sellers to drop-off parcels with.
	Dropoff []string `json:"dropoff,omitempty"`
	// pickup is Could contain 'address_id' and 'pickup_time_id'. Choose one address_id and its corresponding pickup_time_id to Init. If it has empty value, developer should still include "pickup" field in Init API.
	Pickup []string `json:"pickup,omitempty"`
	// non_integrated is Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.
	NonIntegrated []string `json:"non_integrated,omitempty"`
}

//=======================================================
// Object Raw Type - BranchList
//=======================================================
type BranchList struct {
	// branch_id is The identity of logistics branch.
	BranchID int `json:"branch_id,omitempty"`
	// region is The region of specify address.
	Region string `json:"region,omitempty"`
	// state is The state of specify address.
	State string `json:"state,omitempty"`
	// city is The city of specify address.
	City string `json:"city,omitempty"`
	// address is The address description of specify address.
	Address string `json:"address,omitempty"`
	// zipcode is The zipcode of specify address.
	Zipcode string `json:"zipcode,omitempty"`
	// district is The district of specify address.
	District string `json:"district,omitempty"`
	// town is The town of specify address.
	Town string `json:"town,omitempty"`
}

//=======================================================
// Object Raw Type - SlugList
//=======================================================
type SlugList struct {
	// slug is  The identity of slug.
	Slug string `json:"slug,omitempty"`
	// slug_name is  The name of slug.
	SlugName string `json:"slug_name,omitempty"`
}

//=======================================================
// Object Raw Type - Dropoff
//=======================================================
type Dropoff struct {
	// branch_list is List of available dropoff branches info.
	BranchList []BranchList `json:"branch_list"`
	// slug_list is  List of available TW 3PL drop-off partners.
	SlugList []SlugList `json:"slug_list"`
}

//=======================================================
// Object Raw Type - TimeSlotList
//=======================================================
type TimeSlotList struct {
	// date is The date of pickup time. In timestamp.
	Date int `json:"date,omitempty"`
	// time_text is The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text,omitempty"`
	// pickup_time_id is The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id,omitempty"`
}

//=======================================================
// Object Raw Type - AddressList
//=======================================================
type AddressList struct {
	// address_id is The identity of address.
	AddressID int `json:"address_id,omitempty"`
	// region is The region of specify address.
	Region string `json:"region,omitempty"`
	// state is The state of specify address.
	State string `json:"state,omitempty"`
	// city is The city of specify address.
	City string `json:"city,omitempty"`
	// district is The district of specify address.
	District string `json:"district,omitempty"`
	// town is The town of specify address.
	Town string `json:"town,omitempty"`
	// address is The address description of specify address.
	Address string `json:"address,omitempty"`
	// zipcode is The zipcode of specify address.
	Zipcode string `json:"zipcode,omitempty"`
	// address_flag is The flag of shop address, applicable values: default_address, pickup_address, return_address
	AddressFlag []string `json:"address_flag,omitempty"`
	// time_slot_list is List of pickup_time information corresponding to the address_id.
	TimeSlotList []TimeSlotList `json:"time_slot_list"`
}

//=======================================================
// Object Raw Type - Pickup
//=======================================================
type Pickup struct {
	// address_list is List of available pickup address info.
	AddressList []AddressList `json:"address_list"`
}

//=======================================================
// Object Raw Type - LogisticsGetShippingParameter
//=======================================================
type LogisticsGetShippingParameter struct {
	// info_needed is The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init. For logistic_id 80003 and 80004, both Regular and JOB shipping methods are supported. If you choose Regular shipping method, please use "tracking_no" to call Init API. If you choose JOB shipping method, please use "sender_real_name" to call Init API. Note that only one of "tracking_no" and "sender_real_name" can be selected.
	InfoNeeded InfoNeeded `json:"info_needed"`
	// dropoff is Logistics information for dropoff mode order.
	Dropoff Dropoff `json:"dropoff"`
	// pickup is Logistics information for pickup mode order.
	Pickup Pickup `json:"pickup"`
}

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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetShippingParameter `json:"response"`
}

//=======================================================
// Object Raw Type - LogisticsGetTrackingNumber
//=======================================================
type LogisticsGetTrackingNumber struct {
	// tracking_number is The tracking number of this order.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// plp_number is The unique identifier for package of BR correios.
	PlpNumber string `json:"plp_number,omitempty"`
	// first_mile_tracking_number is The first mile tracking number of the order. Only for Cross Border Seller
	FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
	// last_mile_tracking_number is The last mile tracking number of the order. Only for Cross Border BR seller.
	LastMileTrackingNumber string `json:"last_mile_tracking_number,omitempty"`
	// hint is Indicate hint information if cannot get some fields under special scenarios. For example, cannot get tracking_number when cvs store is closed.
	Hint string `json:"hint,omitempty"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetTrackingNumber `json:"response,omitempty"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse
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
	// 通用的 Response 回傳參數
	V2UnityResponse
}

//=======================================================
// Object Raw Type - LogisticsGetShippingDocumentParameter
//=======================================================
type LogisticsGetShippingDocumentParameter struct {
	// result_list is The list of the result data.
	ResultList []ResultList `json:"result_list"`
}

//=======================================================
// LogisticsGetShippingDocumentParameterRequest
//=======================================================
type LogisticsGetShippingDocumentParameterRequest struct {
	// order_list is The list of orders you want to get. limit [1,50]
	OrderList []OrderList `json:"order_list"`
}

//=======================================================
// LogisticsGetShippingDocumentParameterResponse
//=======================================================
type LogisticsGetShippingDocumentParameterResponse struct {
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetShippingDocumentParameter `json:"response"`
}

//=======================================================
// Object Raw Type - LogisticsCreateShippingDocument
//=======================================================
type LogisticsCreateShippingDocument struct {
	// result_list is The list of the result data.
	ResultList []ResultList `json:"result_list"`
}

//=======================================================
// LogisticsCreateShippingDocumentRequest
//=======================================================
type LogisticsCreateShippingDocumentRequest struct {
	// order_list is The list of order you want to create shipping document. limit [1, 50]
	OrderList []OrderList `json:"order_list"`
}

//=======================================================
// LogisticsCreateShippingDocumentResponse
//=======================================================
type LogisticsCreateShippingDocumentResponse struct {
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsCreateShippingDocument `json:"response"`
}

//=======================================================
// Object Raw Type - LogisticsGetShippingDocumentResult
//=======================================================
type LogisticsGetShippingDocumentResult struct {
	// result_list is The result data list of the API response.
	ResultList []ResultList `json:"result_list"`
}

//=======================================================
// LogisticsGetShippingDocumentResultRequest
//=======================================================
type LogisticsGetShippingDocumentResultRequest struct {
	// order_list is The list of orders, limit [1,50]
	OrderList []OrderList `json:"order_list"`
}

//=======================================================
// LogisticsGetShippingDocumentResultResponse
//=======================================================
type LogisticsGetShippingDocumentResultResponse struct {
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetShippingDocumentResult `json:"response"`
}

//=======================================================
// LogisticsDownloadShippingDocumentRequest
//=======================================================
type LogisticsDownloadShippingDocumentRequest struct {
	// shipping_document_type is The type of shipping document. Available values: NORMAL_AIR_WAYBILL,THERMAL_AIR_WAYBILL,NORMAL_JOB_AIR_WAYBILL,THERMAL_JOB_AIR_WAYBILL
	ShippingDocumentType string `json:"shipping_document_type,omitempty"`
	// order_list is The list of orders you need to download it's shipping document.
	OrderList []OrderList `json:"order_list"`
}

//=======================================================
// LogisticsDownloadShippingDocumentResponse
//=======================================================
type LogisticsDownloadShippingDocumentResponse struct {
	// 通用的 Response 回傳參數
	V2UnityResponse
}

//=======================================================
// Object Raw Type - RecipientAddress
//=======================================================
type RecipientAddress struct {
	// name is Recipient's name for the address.
	Name string `json:"name,omitempty"`
	// phone is Recipient's phone number input when order was placed.
	Phone string `json:"phone,omitempty"`
	// town is The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town,omitempty"`
	// district is The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district,omitempty"`
	// city is The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city,omitempty"`
	// state is The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state,omitempty"`
	// region is The two-digit code representing the country of the Recipient.
	Region string `json:"region,omitempty"`
	// zipcode is Recipient's postal code.
	Zipcode string `json:"zipcode,omitempty"`
	// full_address is The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address,omitempty"`
}

//=======================================================
// Object Raw Type - RecipientSortCode
//=======================================================
type RecipientSortCode struct {
	// first_recipient_sort_code is The first-level sort_code of recipient.
	FirstRecipientSortCode string `json:"first_recipient_sort_code,omitempty"`
	// second_recipient_sort_code is The second-level sort_code of recipient.
	SecondRecipientSortCode string `json:"second_recipient_sort_code,omitempty"`
	// third_recipient_sort_code is The third-level sort_code of recipient.
	ThirdRecipientSortCode string `json:"third_recipient_sort_code,omitempty"`
}

//=======================================================
// Object Raw Type - SenderSortCode
//=======================================================
type SenderSortCode struct {
	// first_sender_sort_code is The first-level sort_code of sender.
	FirstSenderSortCode string `json:"first_sender_sort_code,omitempty"`
	// second_sender_sort_code is The second-level sort_code of sender.
	SecondSenderSortCode string `json:"second_sender_sort_code,omitempty"`
	// third_sender_sort_code is The third-level sort_code of sender.
	ThirdSenderSortCode string `json:"third_sender_sort_code,omitempty"`
}

//=======================================================
// Object Raw Type - ThirdPartyLogisticInfo
//=======================================================
type ThirdPartyLogisticInfo struct {
	// service_description is Use this field to indicate the order category.
	ServiceDescription string `json:"service_description,omitempty"`
	// barcode is The manufacturer barcode.
	Barcode string `json:"barcode,omitempty"`
	// purchase_time is The purchase_time of the store.
	PurchaseTime string `json:"purchase_time,omitempty"`
	// return_time is The return_time of the store.
	ReturnTime string `json:"return_time,omitempty"`
	// manufacturers_name is The name of manufacturers.
	ManufacturersName string `json:"manufacturers_name,omitempty"`
	// manufacturers_website is The website of manufacturers.
	ManufacturersWebsite string `json:"manufacturers_website,omitempty"`
	// recipient_area is The identification of recipient area.
	RecipientArea string `json:"recipient_area,omitempty"`
	// route_step is The route code of the waybill.
	RouteStep string `json:"route_step,omitempty"`
	// suda5_code is The tally code of the waybill.
	Suda5Code string `json:"suda5_code,omitempty"`
	// large_logistics_id is The code of large logistics.
	LargeLogisticsID string `json:"large_logistics_id,omitempty"`
	// parent_id is The parent code of the waybill.
	ParentID string `json:"parent_id,omitempty"`
	// return_cycle is Use this field to indicate the return cycle.
	ReturnCycle string `json:"return_cycle,omitempty"`
	// return_mode is Use this field to indicate the return mode.
	ReturnMode string `json:"return_mode,omitempty"`
	// prompt is The reminder of stork work.
	Prompt string `json:"prompt,omitempty"`
	// order_sn is Shopee's unique identifier for an order.
	OrderSn string `json:"order_sn,omitempty"`
	// qrcode is The QR code of the waybill.
	Qrcode string `json:"qrcode,omitempty"`
	// ec_supplier_name is The supplier name of channel.
	EcSupplierName string `json:"ec_supplier_name,omitempty"`
	// ec_bar_code16 is Use this field to indicate the first barcode.
	EcBarCode16 string `json:"ec_bar_code16,omitempty"`
	// equipment_id is The device code.
	EquipmentID string `json:"equipment_id,omitempty"`
	// eshop_id is The child code for B2C Family-mart.
	EshopID string `json:"eshop_id,omitempty"`
	// ec_bar_code9 is Use this field to indicate the pick barcode.
	EcBarCode9 string `json:"ec_bar_code9,omitempty"`
	// pelican_tracking_no is The tracking number of Shopee Delivery.
	PelicanTrackingNo string `json:"pelican_tracking_no,omitempty"`
	// print_date is The date of printing the wayBill.
	PrintDate string `json:"print_date,omitempty"`
	// pzip is The sort code of the order.
	Pzip string `json:"pzip,omitempty"`
	// pzip_c is The barcode of the sort code.
	PzipC string `json:"pzip_c,omitempty"`
	// deliver_area_txt is The code of the delivery area.
	DeliverAreaTxt string `json:"deliver_area_txt,omitempty"`
	// deliver_date_ymd is Expected delivery date of the order.
	DeliverDateYmd string `json:"deliver_date_ymd,omitempty"`
	// sd_driver_code is Lorry driver code of the order.
	SdDriverCode string `json:"sd_driver_code,omitempty"`
	// md_driver_code is Motorcycle driver code of the order.
	MdDriverCode string `json:"md_driver_code,omitempty"`
	// putorder_stackzone_code is Stacking area of the order.
	PutorderStackzoneCode string `json:"putorder_stackzone_code,omitempty"`
	// customer_code is  The cutomer code of Shopee.
	CustomerCode string `json:"customer_code,omitempty"`
	// deliver_router is Use this field to indicate the delivery router.
	DeliverRouter string `json:"deliver_router,omitempty"`
	// store_type is Use this field to indicate the store type.
	StoreType string `json:"store_type,omitempty"`
	// pick_router is Use this field to indicate the pick router.
	PickRouter string `json:"pick_router,omitempty"`
	// barcode_dc is The main logistic barcode of the waybill.
	BarcodeDc string `json:"barcode_dc,omitempty"`
	// ec_order_number is Use this field to indicate the logistics order number.
	EcOrderNumber string `json:"ec_order_number,omitempty"`
	// barcode_pr is The sorting barcode of the waybill.
	BarcodePr string `json:"barcode_pr,omitempty"`
	// first_pick_barcode is The first pick barcode of the waybill.
	FirstPickBarcode string `json:"first_pick_barcode,omitempty"`
	// second_pick_barcode is The second pick barcode of the waybill.
	SecondPickBarcode string `json:"second_pick_barcode,omitempty"`
	// is_cod_bool is Use this field to indicate the service type.
	IsCodBool string `json:"is_cod_bool,omitempty"`
	// receiver_name is Use this field to indicate the receiver name.
	ReceiverName string `json:"receiver_name,omitempty"`
	// rcv_store_name is Use this field to indicate the receiver store name.
	RcvStoreName string `json:"rcv_store_name,omitempty"`
	// branch_code is Use this field indicates destination service point code.
	BranchCode string `json:"branch_code,omitempty"`
	// branch_name is Use this field indicates destination service point name.
	BranchName string `json:"branch_name,omitempty"`
	// last_third_digits_recipient_phone is Use this field indicates buyer phone number (last 3 digits).
	LastThirdDigitsRecipientPhone string `json:"last_third_digits_recipient_phone,omitempty"`
	// last_third_digits_sender_phone is Use this field indicates seller phone number (last 3 digits).
	LastThirdDigitsSenderPhone string `json:"last_third_digits_sender_phone,omitempty"`
}

//=======================================================
// Object Raw Type - ReturnSortCode
//=======================================================
type ReturnSortCode struct {
	// return_first_sort_code is The first-level sort code for 3PL doing RTS.
	ReturnFirstSortCode string `json:"return_first_sort_code,omitempty"`
}

//=======================================================
// Object Raw Type - SpxReceiveStation
//=======================================================
type SpxReceiveStation struct {
	// spx_first_receive_station is The first pickup station.
	SpxFirstReceiveStation string `json:"spx_first_receive_station,omitempty"`
}

//=======================================================
// Object Raw Type - ShippingDocumentInfo
//=======================================================
type ShippingDocumentInfo struct {
	// logistics_channel_id is The identity of logistic channel.
	LogisticsChannelID int `json:"logistics_channel_id,omitempty"`
	// shipping_carrier is The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// service_code is Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code,omitempty"`
	// first_mile_name is Only work for cross-border order.The name of the carrier ships cross countries.
	FirstMileName string `json:"first_mile_name,omitempty"`
	// last_mile_name is Only work for cross-border order.The name of the carrier delivers the parcels in local country.
	LastMileName string `json:"last_mile_name,omitempty"`
	// goods_to_declare is Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare,omitempty"`
	// tracking_number is The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// zone is The zone of this order.
	Zone string `json:"zone,omitempty"`
	// lane_code is Only work for cross-border order. The string use for waybill printing. The format is "S - country_code and lane_number". For example, S-TH01, S-TH02
	LaneCode string `json:"lane_code,omitempty"`
	// warehouse_address is Only work for cross-border order in some special shop. The address info of the warehouse.
	WarehouseAddress string `json:"warehouse_address,omitempty"`
	// warehouse_id is Only work for cross-border order in some special shop. The ID of the warehouse.
	WarehouseID string `json:"warehouse_id,omitempty"`
	// recipient_address is This object contains detailed breakdown for the recipient address.
	RecipientAddress RecipientAddress `json:"recipient_address"`
	// cod is This value indicates whether the order was a COD (cash on delivery) order.
	Cod bool `json:"cod,omitempty"`
	// recipient_sort_code is The sort_code of recipient.
	RecipientSortCode RecipientSortCode `json:"recipient_sort_code"`
	// sender_sort_code is The sort_code of sender.
	SenderSortCode SenderSortCode `json:"sender_sort_code"`
	// third_party_logistic_info is Only used for local TW sellers.
	ThirdPartyLogisticInfo ThirdPartyLogisticInfo `json:"third_party_logistic_info"`
	// buyer_cpf_id is Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.
	BuyerCpfID string `json:"buyer_cpf_id,omitempty"`
	// shopee_tracking_number is First mile tracking NO. for CrossBoard BR seller can be used to self-design CB Brazil AWB.
	ShopeeTrackingNumber string `json:"shopee_tracking_number,omitempty"`
	// last_mile_tracking_number is The last-mile tracking number. Only for Cross Board BR seller.
	LastMileTrackingNumber string `json:"last_mile_tracking_number,omitempty"`
	// pickup_hub is The name of pickup hub.
	PickupHub string `json:"pickup_hub,omitempty"`
	// delivery_hub is The name of delivery hub.
	DeliveryHub string `json:"delivery_hub,omitempty"`
	// ec_order_no is The name of ec order.
	EcOrderNo string `json:"ec_order_no,omitempty"`
	// create_date_ymd_sl is The date of create shipment order.
	CreateDateYmdSl string `json:"create_date_ymd_sl,omitempty"`
	// manufacturers_name is The name of manufacturer.
	ManufacturersName string `json:"manufacturers_name,omitempty"`
	// manufacturers_website is The website of manufacturer.
	ManufacturersWebsite string `json:"manufacturers_website,omitempty"`
	// is_lm_dg_bool is Use this field to indicate order contains dangerous goods or not.
	IsLmDgBool string `json:"is_lm_dg_bool,omitempty"`
	// preferred_delivery_option is Use this field to indicate delivery address is residential or office address.
	PreferredDeliveryOption string `json:"preferred_delivery_option,omitempty"`
	// return_sort_code is The sort code for 3PL doing RTS.
	ReturnSortCode ReturnSortCode `json:"return_sort_code"`
	// spx_sub_district is The sub-district of recipient's address.
	SpxSubDistrict string `json:"spx_sub_district,omitempty"`
	// recipient_addr is Full address of recipient.
	RecipientAddr string `json:"recipient_addr,omitempty"`
	// deliver_area is Zone name.
	DeliverArea string `json:"deliver_area,omitempty"`
	// spx_receive_station is The spx receive station.
	SpxReceiveStation SpxReceiveStation `json:"spx_receive_station"`
}

//=======================================================
// Object Raw Type - LogisticsGetShippingDocumentInfo
//=======================================================
type LogisticsGetShippingDocumentInfo struct {
	// shipping_document_info is The shipping document info of this order.
	ShippingDocumentInfo ShippingDocumentInfo `json:"shipping_document_info"`
	// cod_amount is Use this field to indicate cod amount.
	CodAmount string `json:"cod_amount,omitempty"`
	// order_weight is Use this field to indicate order weight when calculate the shipping fee. The unit of weigh is gram.
	OrderWeight string `json:"order_weight,omitempty"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetShippingDocumentInfo `json:"response,omitempty"`
}

//=======================================================
// Object Raw Type - TrackingInfo
//=======================================================
type TrackingInfo struct {
	// update_time is The time when logistics info has been updated.
	UpdateTime int `json:"update_time,omitempty"`
	// description is The description of order logistics tracking info.
	Description string `json:"description,omitempty"`
	// logistics_status is The Shopee logistics status for the order. Applicable values: See Data Definition- TrackingLogisticsStatus.
	LogisticsStatus string `json:"logistics_status,omitempty"`
}

//=======================================================
// Object Raw Type - LogisticsGetTrackingInfo
//=======================================================
type LogisticsGetTrackingInfo struct {
	// order_sn is Shopee's unique identifier for an order.
	OrderSn string `json:"order_sn,omitempty"`
	// package_number is Shopee's unique identifier for the package under an order.
	PackageNumber string `json:"package_number,omitempty"`
	// logistics_status is The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticsStatus string `json:"logistics_status,omitempty"`
	// tracking_info is The tracking info of the order.
	TrackingInfo []TrackingInfo `json:"tracking_info"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetTrackingInfo `json:"response,omitempty"`
}

//=======================================================
// Object Raw Type - LogisticsGetAddressList
//=======================================================
type LogisticsGetAddressList struct {
	// show_pickup_address is Show pickup address or not.
	ShowPickupAddress bool `json:"show_pickup_address,omitempty"`
	// address_list is The address list of you shop
	AddressList []AddressList `json:"address_list"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetAddressList `json:"response,omitempty"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse
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
	// 通用的 Response 回傳參數
	V2UnityResponse
}

//=======================================================
// Object Raw Type - SizeList
//=======================================================
type SizeList struct {
	// size_id is The identity of size.
	SizeID string `json:"size_id,omitempty"`
	// name is The name of size.
	Name string `json:"name,omitempty"`
	// default_price is The pre-defined shipping fee for the specific size.
	DefaultPrice float64 `json:"default_price,omitempty"`
}

//=======================================================
// Object Raw Type - WeightLimit
//=======================================================
type WeightLimit struct {
	// item_max_weight is The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMaxWeight float64 `json:"item_max_weight,omitempty"`
	// item_min_weight is The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
	ItemMinWeight float64 `json:"item_min_weight,omitempty"`
}

//=======================================================
// Object Raw Type - ItemMaxDimension
//=======================================================
type ItemMaxDimension struct {
	// height is The max height limit.
	Height float64 `json:"height,omitempty"`
	// width is The max width limit.
	Width float64 `json:"width,omitempty"`
	// length is The max length limit.
	Length float64 `json:"length,omitempty"`
	// unit is The unit for the limit.
	Unit string `json:"unit,omitempty"`
	// dimension_sum is The sum of the item's dimension
	DimensionSum float64 `json:"dimension_sum,omitempty"`
}

//=======================================================
// Object Raw Type - VolumeLimit
//=======================================================
type VolumeLimit struct {
	// item_max_volume is The max volume for an item on this logistic channel.If the value is 0 or null, that means there is no limit for the item weight.
	ItemMaxVolume float64 `json:"item_max_volume,omitempty"`
	// item_min_volume is The min volume for an item on this logistic channel. If the value is 0 or null, that means there is no limit for the item weight.
	ItemMinVolume float64 `json:"item_min_volume,omitempty"`
}

//=======================================================
// Object Raw Type - LogisticsChannelList
//=======================================================
type LogisticsChannelList struct {
	// logistics_channel_id is The identity of logistic channel.
	LogisticsChannelID int `json:"logistics_channel_id,omitempty"`
	// preferred is Whether to make this logistic channel preferred. Indonestia logistics channel are not applicable.
	Preferred bool `json:"preferred,omitempty"`
	// logistics_channel_name is The name of logistic channel.
	LogisticsChannelName string `json:"logistics_channel_name,omitempty"`
	// cod_enabled is This is to indicate whether this logistic channel supports COD
	CodEnabled bool `json:"cod_enabled,omitempty"`
	// enabled is Whether this logistic channel is enabled on shop level.
	Enabled bool `json:"enabled,omitempty"`
	// fee_type is See define FeeType, related to FeeType Value
	FeeType string `json:"fee_type,omitempty"`
	// size_list is Only for fee_type is SIZE_SELECTION
	SizeList []SizeList `json:"size_list"`
	// weight_limit is The weight limit for this logistic channel.
	WeightLimit WeightLimit `json:"weight_limit"`
	// item_max_dimension is The dimension limit for this logistic channel.
	ItemMaxDimension ItemMaxDimension `json:"item_max_dimension"`
	// volume_limit is The limit of item volume.
	VolumeLimit VolumeLimit `json:"volume_limit"`
	// logistics_description is For checkout channels, this field indicates its corresponding fulfillment channels.
	LogisticsDescription string `json:"logistics_description,omitempty"`
	// force_enable is Indicates whether the logistic channel is force enabled on Shop Level. If true, sellers cannot close this channel.
	ForceEnable bool `json:"force_enable,omitempty"`
	// mask_channel_id is Indicate the parent logistic channel ID. If it’s 0, it indicates the channel is a checkout(masked) channel; if it’s not 0, indicate the channel is a fulfillment channel and has a checkout channel(checkout channel’s channel_id equals this mask_channel_id) on top of it. Multiple channels may share the same mask_channel_id.
	MaskChannelID int `json:"mask_channel_id,omitempty"`
}

//=======================================================
// Object Raw Type - LogisticsGetChannelList
//=======================================================
type LogisticsGetChannelList struct {
	// logistics_channel_list is The list of logistics channel.
	LogisticsChannelList []LogisticsChannelList `json:"logistics_channel_list"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsGetChannelList `json:"response"`
}

//=======================================================
// Object Raw Type - UnsupportWarehouse
//=======================================================
type UnsupportWarehouse struct {
	// warehouse_id is <p>Unsupported warehouse ID<br /></p>
	WarehouseID int `json:"warehouse_id,omitempty"`
	// warehouse_name is <p>Unsupported warehouse name<br /></p>
	WarehouseName string `json:"warehouse_name,omitempty"`
}

//=======================================================
// Object Raw Type - UpdatedChannels
//=======================================================
type UpdatedChannels struct {
	// channel_id is <p>Logistics channel ID<br /></p>
	ChannelID int `json:"channel_id,omitempty"`
	// channel_display_name is <p>Logistics channel name<br /></p>
	ChannelDisplayName string `json:"channel_display_name,omitempty"`
	// unsupport_warehouse is <p>List details of unsupported warehouses<br /></p>
	UnsupportWarehouse []UnsupportWarehouse `json:"unsupport_warehouse"`
}

//=======================================================
// Object Raw Type - LogisticsUpdateChannel
//=======================================================
type LogisticsUpdateChannel struct {
	// shop_id is Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id,omitempty"`
	// enabled is Whether this logistic channel is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// preferred is Whether this logistic channel is preferred.
	Preferred bool `json:"preferred,omitempty"`
	// cod_enabled is Whether COD is enabled for this channel.
	CodEnabled bool `json:"cod_enabled,omitempty"`
	// logistics_channel_id is The identity of logistic channel.
	LogisticsChannelID int `json:"logistics_channel_id,omitempty"`
	// updated_channels is <p>List of channels that are updated in the operation (inclusive of dependent logistics channels)<br /></p>
	UpdatedChannels []UpdatedChannels `json:"updated_channels"`
	// is_multi_warehouse is
	IsMultiWarehouse bool `json:"is_multi_warehouse,omitempty"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is Detail informations you are querying.
	Response LogisticsUpdateChannel `json:"response,omitempty"`
}

//=======================================================
// Object Raw Type - LogisticsBatchShipOrder
//=======================================================
type LogisticsBatchShipOrder struct {
	// result_list is
	ResultList []ResultList `json:"result_list"`
}

//=======================================================
// LogisticsBatchShipOrderRequest
//=======================================================
type LogisticsBatchShipOrderRequest struct {
	// order_list is The list of order.
	OrderList []OrderList `json:"order_list"`
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
	// 通用的 Response 回傳參數
	V2UnityResponse

	// response is
	Response LogisticsBatchShipOrder `json:"response,omitempty"`
}
