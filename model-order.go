package shopeego


//=======================================================
// Object Raw Type - OrderGetOrderListOrderList
//=======================================================
type OrderGetOrderListOrderList struct {
// order_sn is  Shopee's unique identifier for an order.	
OrderSN string `json:"order_sn,omitempty"`
// order_status is The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED
OrderStatus string `json:"order_status,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderList
//=======================================================
type OrderGetOrderList struct {
// more is This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
More bool `json:"more,omitempty"`
// order_list is 
OrderList []OrderGetOrderListOrderList `json:"order_list"`
// next_cursor is If  more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// OrderGetOrderListRequest
//=======================================================
type OrderGetOrderListRequest struct {
    // time_range_field is The kind of time_from and time_to. Available value: create_time, update_time.
    TimeRangeField string `json:"time_range_field"`
    // time_from is The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days. 
    TimeFrom int `json:"time_from"`
    // time_to is The time_from and time_to fields specify a date range for retrieving orders (based on the time_range_field). The time_from field is the starting date range. The maximum date range that may be specified with the time_from and time_to fields is 15 days.
    TimeTo int `json:"time_to"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
    PageSize int `json:"page_size"`
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
    // order_status is The order_status filter for retriveing orders and each one only every request. Available value: UNPAID/READY_TO_SHIP/PROCESSED/SHIPPED/COMPLETED/IN_CANCEL/CANCELLED/INVOICE_PENDING
    OrderStatus string `json:"order_status,omitempty"`
    // response_optional_fields is Optional fields in response. Available value: order_status.
    ResponseOptionalFields string `json:"response_optional_fields,omitempty"`
}
//=======================================================
// OrderGetOrderListResponse
//=======================================================
type OrderGetOrderListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderGetOrderList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetShipmentListOrderList
//=======================================================
type OrderGetShipmentListOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_number is Shopee's unique identifier for the package under an order
PackageNumber string `json:"package_number,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetShipmentList
//=======================================================
type OrderGetShipmentList struct {
// order_list is The list of  shipment orders
OrderList []OrderGetShipmentListOrderList `json:"order_list"`
// more is This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
More bool `json:"more,omitempty"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
}
//=======================================================
// OrderGetShipmentListRequest
//=======================================================
type OrderGetShipmentListRequest struct {
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
    PageSize int `json:"page_size"`
}
//=======================================================
// OrderGetShipmentListResponse
//=======================================================
type OrderGetShipmentListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderGetShipmentList `json:"response"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListRecipientAddress
//=======================================================
type OrderGetOrderDetailOrderListRecipientAddress struct {
// name is Recipient's name for the address.	
Name string `json:"name,omitempty"`
// phone is Recipient's phone number input when order was placed.	
Phone string `json:"phone,omitempty"`
// town is  The town of the recipient's address. Whether there is a town will depend on the region and/or country.	
Town string `json:"town,omitempty"`
// district is The district of the recipient's address. Whether there is a district will depend on the region and/or country.	
District string `json:"district,omitempty"`
// city is The city of the recipient's address. Whether there is a city will depend on the region and/or country.	
City string `json:"city,omitempty"`
// state is The state/province of the recipient's address. Whether there is a state/province will depend on the region and/or country.	
State string `json:"state,omitempty"`
// region is The two-digit code representing the region of the Recipient.	
Region string `json:"region,omitempty"`
// zipcode is Recipient's postal code.	
Zipcode string `json:"zipcode,omitempty"`
// full_address is The full address of the recipient, including country, state, even street, and etc.	
FullAddress string `json:"full_address,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListItemListImageInfo
//=======================================================
type OrderGetOrderDetailOrderListItemListImageInfo struct {
// image_url is The image url of the product. Default to be variation image, if the model does not have a variation image, will use an item main image instead.
ImageUrl string `json:"image_url,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListItemList
//=======================================================
type OrderGetOrderDetailOrderListItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// item_name is The name of the item.
ItemName string `json:"item_name,omitempty"`
// item_sku is  A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.	
ItemSku string `json:"item_sku,omitempty"`
// model_id is ID of the model that belongs to the same item.	
ModelID int `json:"model_id,omitempty"`
// model_name is Name of the model that belongs to the same item. A seller can offer models of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate model. Each model can have a different quantity and price.	
ModelName string `json:"model_name,omitempty"`
// model_sku is A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are models of one item in Shopee Listings.	
ModelSku string `json:"model_sku,omitempty"`
// model_quantity_purchased is The number of identical items purchased at the same time by the same buyer from one listing/item.	
ModelQuantityPurchased int `json:"model_quantity_purchased,omitempty"`
// model_original_price is The original price of the item in the listing currency.	
ModelOriginalPrice float64 `json:"model_original_price,omitempty"`
// model_discounted_price is The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of model_original_price. In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/model level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.	
ModelDiscountedPrice float64 `json:"model_discounted_price,omitempty"`
// wholesale is This value indicates whether buyer buy the order item in wholesale price.	
Wholesale bool `json:"wholesale,omitempty"`
// weight is The weight of the item	
Weight float64 `json:"weight,omitempty"`
// add_on_deal is To indicate if this item belongs to an addon deal.	
AddOnDeal bool `json:"add_on_deal,omitempty"`
// main_item is To indicate if this item is main item or sub item. True means main item, false means sub item.	
MainItem bool `json:"main_item,omitempty"`
// add_on_deal_id is A unique ID to distinguish groups of items in Cart, and Order. (e.g. AddOnDeal)	
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// promotion_type is <p>Available type：product_promotion, flash_sale, group_by, bundle_deal, add_on_deal_main, add_on_deal_sub, add_on_free_gift_main,&nbsp;add_on_free_gift_sub</p>
PromotionType string `json:"promotion_type,omitempty"`
// promotion_id is The ID of the promotion.	
PromotionID int `json:"promotion_id,omitempty"`
// order_item_id is The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.
OrderItemID int `json:"order_item_id,omitempty"`
// promotion_group_id is The identify of product promotion.
PromotionGroupID int `json:"promotion_group_id,omitempty"`
// image_info is  Image info of the product.
ImageInfo OrderGetOrderDetailOrderListItemListImageInfo `json:"image_info"`
// product_location_id is The list of warehouse IDs of the item.
ProductLocationID []string `json:"product_location_id,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListPackageListItemList
//=======================================================
type OrderGetOrderDetailOrderListPackageListItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// model_id is Shopee's unique identifier for a model.
ModelID int `json:"model_id,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListPackageList
//=======================================================
type OrderGetOrderDetailOrderListPackageList struct {
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// logistics_status is The Shopee logistics status for the order. Applicable values: See Data Definition-LogisticsStatus.
LogisticsStatus string `json:"logistics_status,omitempty"`
// shipping_carrier is The logistics service provider that the buyer selected for the order to deliver items.
ShippingCarrier string `json:"shipping_carrier,omitempty"`
// item_list is The lis of items.
ItemList []OrderGetOrderDetailOrderListPackageListItemList `json:"item_list"`
// parcel_chargeable_weight_gram is <p>For CB shop, display weight used to calculate actual_shipping_fee for this parcel.<br /></p>
ParcelChargeableWeightGram int `json:"parcel_chargeable_weight_gram,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderListInvoiceData
//=======================================================
type OrderGetOrderDetailOrderListInvoiceData struct {
// number is The number of the invoice. The number should be 9 digits. pt: número da NF-e.
Number string `json:"number,omitempty"`
// series_number is The series number of the invoice. The series number should be 3 digits. pt: série da NF-e.
SeriesNumber string `json:"series_number,omitempty"`
// access_key is The access key of the invoice. The access key should be 44 digits. pt: chave de acesso da NF-e.
AccessKey string `json:"access_key,omitempty"`
// issue_date is The issue date of the invoice. The issue date should be later than the order pay date. pt: data de emissão da NF-e.
IssueDate int `json:"issue_date,omitempty"`
// total_value is The total value of the invoice. pt: valor total da NF-e (R$).
TotalValue float64 `json:"total_value,omitempty"`
// products_total_value is The products total value of the invoice. pt: valor total dos produtos (R$) da NF-e.
ProductsTotalValue float64 `json:"products_total_value,omitempty"`
// tax_code is The tax code for the invoice. The tax code should be 4 digits. pt: Código Fiscal de Operações e Prestações (CFOP) predominante na NF-e. 
TaxCode string `json:"tax_code,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetailOrderList
//=======================================================
type OrderGetOrderDetailOrderList struct {
// order_sn is Return by default. Shopee's unique identifier for an order.	
OrderSN string `json:"order_sn,omitempty"`
// region is Return by default. The two-digit code representing the region where the order was made.	
Region string `json:"region,omitempty"`
// currency is Return by default. The three-digit code representing the currency unit for which the order was paid.	
Currency string `json:"currency,omitempty"`
// cod is Return by default. This value indicates whether the order was a COD (cash on delivery) order.	
Cod bool `json:"cod,omitempty"`
// total_amount is The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable. This value will only return after the buyer has completed payment for the order.	
TotalAmount float64 `json:"total_amount,omitempty"`
// order_status is Return by default. Enumerated type that defines the current status of the order.	
OrderStatus string `json:"order_status,omitempty"`
// shipping_carrier is The logistics service provider that the buyer selected for the order to deliver items.	
ShippingCarrier string `json:"shipping_carrier,omitempty"`
// payment_method is The payment method that the buyer selected to pay for the order.
PaymentMethod string `json:"payment_method,omitempty"`
// estimated_shipping_fee is The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard.	
EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty"`
// message_to_seller is Return by default. Message to seller.	
MessageToSeller string `json:"message_to_seller,omitempty"`
// create_time is Return by default. Timestamp that indicates the date and time that the order was created.	
CreateTime int `json:"create_time,omitempty"`
// update_time is Return by default. Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.	
UpdateTime int `json:"update_time,omitempty"`
// days_to_ship is Return by default. Shipping preparation time set by the seller when listing item on Shopee.	
DaysToShip int `json:"days_to_ship,omitempty"`
// ship_by_date is Return by default. The deadline to ship out the parcel.	
ShipByDate int `json:"ship_by_date,omitempty"`
// buyer_user_id is The user id of buyer of this order	
BuyerUserID int `json:"buyer_user_id,omitempty"`
// buyer_username is The name of buyer	
BuyerUsername string `json:"buyer_username,omitempty"`
// recipient_address is This object contains detailed breakdown for the recipient address.	
RecipientAddress OrderGetOrderDetailOrderListRecipientAddress `json:"recipient_address"`
// actual_shipping_fee is The actual shipping fee of the order if available from external logistics partners.	
ActualShippingFee float64 `json:"actual_shipping_fee,omitempty"`
// goods_to_declare is Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.	
GoodsToDeclare bool `json:"goods_to_declare,omitempty"`
// note is The note seller made for own reference.	
Note string `json:"note,omitempty"`
// note_update_time is Update time for the note.	
NoteUpdateTime int `json:"note_update_time,omitempty"`
// item_list is This object contains the detailed breakdown for the result of this API call.	
ItemList []OrderGetOrderDetailOrderListItemList `json:"item_list"`
// pay_time is The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.	
PayTime int `json:"pay_time,omitempty"`
// dropshipper is For Indonesia orders only. The name of the dropshipper.	
Dropshipper string `json:"dropshipper,omitempty"`
// dropshipper_phone is The phone number of dropshipper, could be empty.
DropshipperPhone string `json:"dropshipper_phone,omitempty"`
// split_up is <p>To indicate whether this order is split to fullfil order(forder) level. Call&nbsp;v2.order.split_order if it's "true".</p>
SplitUp bool `json:"split_up,omitempty"`
// buyer_cancel_reason is Cancel reason from buyer, could be empty.
BuyerCancelReason string `json:"buyer_cancel_reason,omitempty"`
// cancel_by is Could be one of buyer, seller, system or Ops.	
CancelBy string `json:"cancel_by,omitempty"`
// cancel_reason is Use this field to get reason for buyer, seller, and system cancellation.	
CancelReason string `json:"cancel_reason,omitempty"`
// actual_shipping_fee_confirmed is Use this filed to judge whether the actual_shipping_fee is confirmed.	
ActualShippingFeeConfirmed bool `json:"actual_shipping_fee_confirmed,omitempty"`
// buyer_cpf_id is Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.	
BuyerCpfID string `json:"buyer_cpf_id,omitempty"`
// fulfillment_flag is Use this field to indicate the order is fulfilled by shopee or seller. Applicable values: fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.	
FulfillmentFlag string `json:"fulfillment_flag,omitempty"`
// pickup_done_time is The timestamp when pickup is done.
PickupDoneTime int `json:"pickup_done_time,omitempty"`
// package_list is The list of package under an order
PackageList []OrderGetOrderDetailOrderListPackageList `json:"package_list"`
// invoice_data is The invoice data of the order. pt: Nota Fiscal eletrônica (NF-e) do pedido.
InvoiceData OrderGetOrderDetailOrderListInvoiceData `json:"invoice_data"`
// checkout_shipping_carrier is For non masking order, the logistics service provider that the buyer selected for the order to deliver items. For masking order, the logistics service type that the buyer selected for the order to deliver items.
CheckoutShippingCarrier string `json:"checkout_shipping_carrier,omitempty"`
// reverse_shipping_fee is Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
ReverseShippingFee float64 `json:"reverse_shipping_fee,omitempty"`
// order_chargeable_weight_gram is <p>For CB shop, display weight used to calculate actual_shipping_fee for this order.<br /></p>
OrderChargeableWeightGram int `json:"order_chargeable_weight_gram,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetOrderDetail
//=======================================================
type OrderGetOrderDetail struct {
// order_list is The list of orders.
OrderList []OrderGetOrderDetailOrderList `json:"order_list"`
}
//=======================================================
// OrderGetOrderDetailRequest
//=======================================================
type OrderGetOrderDetailRequest struct {
    // order_sn_list is The set of order_sn. If there are multiple order_sn, you need to use English comma to connect them. limit [1,50]
    OrderSnList []string `json:"order_sn_list"`
    // response_optional_fields is <p>Indicate response fields you want to get. Please select from the below response parameters. If you input an object field, all the params under it will be included automatically in the response. If there are multiple response fields you want to get, you need to use English comma to connect them.  Available values: buyer_user_id,buyer_username,estimated_shipping_fee,recipient_address,actual_shipping_fee ,goods_to_declare,note,note_update_time,item_list,pay_time,dropshipper,dropshipper_phone,split_up,buyer_cancel_reason,cancel_by,cancel_reason,actual_shipping_fee_confirmed,buyer_cpf_id,fulfillment_flag,pickup_done_time,package_list,shipping_carrier,payment_method,total_amount,buyer_username,invoice_data, checkout_shipping_carrier, reverse_shipping_fee, order_chargeable_weight_gram etc.</p>
    ResponseOptionalFields []string `json:"response_optional_fields,omitempty"`
}
//=======================================================
// OrderGetOrderDetailResponse
//=======================================================
type OrderGetOrderDetailResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderGetOrderDetail `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - PackageListItemList
//=======================================================
type PackageListItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// model_id is Shopee's unique identifier for a model of an item.
ModelID int `json:"model_id,omitempty"`
// order_item_id is The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.
OrderItemID int `json:"order_item_id,omitempty"`
// promotion_group_id is The identify of product promotion. For items in one same add on deal promotion, the promotion_group_id should share the same id. For items not in add on deal promotion, the promotion_group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.
PromotionGroupID int `json:"promotion_group_id,omitempty"`
}


//=======================================================
// Object Raw Type - PackageList
//=======================================================
type PackageList struct {
// item_list is The list of items under the same package.
ItemList []PackageListItemList `json:"item_list"`
}


//=======================================================
// Object Raw Type - OrderSplitOrderPackageListItemList
//=======================================================
type OrderSplitOrderPackageListItemList struct {
// item_id is Shopee's unique identifier for an item.
ItemID int `json:"item_id,omitempty"`
// model_id is Shopee's unique identifier for a model.
ModelID int `json:"model_id,omitempty"`
// order_item_id is The identify of order item. For items in one same bundle deal promotion, the order_item_id should share the same id, such as 1,2. For items not in bundle deal promotion, the order_item_id should be the same as item_id.
OrderItemID int `json:"order_item_id,omitempty"`
// promotion_group_id is The identify of product promotion. For items in one same add on deal promotion, the promotion_group_id should share the same id. For items not in add on deal promotion, the promotion_group_id should be 0. And the data is from group_id of shopee.orders.GetOrderDetails.
PromotionGroupID int `json:"promotion_group_id,omitempty"`
}


//=======================================================
// Object Raw Type - OrderSplitOrderPackageList
//=======================================================
type OrderSplitOrderPackageList struct {
// package_number is Shopee's unique identifier for the package under an order.
PackageNumber string `json:"package_number,omitempty"`
// item_list is The list of items under this package.
ItemList []OrderSplitOrderPackageListItemList `json:"item_list"`
}


//=======================================================
// Object Raw Type - OrderSplitOrder
//=======================================================
type OrderSplitOrder struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// package_list is The list of package under this order you have split.
PackageList []OrderSplitOrderPackageList `json:"package_list"`
}
//=======================================================
// OrderSplitOrderRequest
//=======================================================
type OrderSplitOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
    // package_list is The list of packages that you want to split
    PackageList []PackageList `json:"package_list"`
}
//=======================================================
// OrderSplitOrderResponse
//=======================================================
type OrderSplitOrderResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderSplitOrder `json:"response"`
}
//=======================================================
// OrderUnsplitOrderRequest
//=======================================================
type OrderUnsplitOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
}
//=======================================================
// OrderUnsplitOrderResponse
//=======================================================
type OrderUnsplitOrderResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - OrderCancelOrder
//=======================================================
type OrderCancelOrder struct {
// update_time is The time when the order is updated.
UpdateTime int `json:"update_time,omitempty"`
}
//=======================================================
// OrderCancelOrderRequest
//=======================================================
type OrderCancelOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
    // cancel_reason is The reason seller want to cancel this order. Applicable values: OUT_OF_STOCK, CUSTOMER_REQUEST, UNDELIVERABLE_AREA, COD_NOT_SUPPORTED.
    CancelReason string `json:"cancel_reason"`
    // item_list is Required when cancel_reason is OUT_OF_STOCK. 
    ItemList []ItemList `json:"item_list,omitempty"`
}
//=======================================================
// OrderCancelOrderResponse
//=======================================================
type OrderCancelOrderResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderCancelOrder `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - OrderHandleBuyerCancellation
//=======================================================
type OrderHandleBuyerCancellation struct {
// update_time is The time when the order is updated.
UpdateTime int `json:"update_time,omitempty"`
}
//=======================================================
// OrderHandleBuyerCancellationRequest
//=======================================================
type OrderHandleBuyerCancellationRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
    // operation is The operation you want to handle.Avaiable value: ACCEPT, REJECT
    Operation string `json:"operation"`
}
//=======================================================
// OrderHandleBuyerCancellationResponse
//=======================================================
type OrderHandleBuyerCancellationResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderHandleBuyerCancellation `json:"response"`
}
//=======================================================
// OrderSetNoteRequest
//=======================================================
type OrderSetNoteRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
    // note is The note seller add for reference.
    Note string `json:"note"`
}
//=======================================================
// OrderSetNoteResponse
//=======================================================
type OrderSetNoteResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - InvoiceData
//=======================================================
type InvoiceData struct {
// number is  The number of the invoice. The number should be 9 digits. pt: número da NF-e.
Number string `json:"number,omitempty"`
// series_number is The series number of the invoice. The series number should be 3 digits. pt: série da NF-e.
SeriesNumber string `json:"series_number,omitempty"`
// access_key is The access key of the invoice. The access key should be 44 digits. pt: chave de acesso da NF-e.
AccessKey string `json:"access_key,omitempty"`
// issue_date is The issue date of the invoice. The issue date should be later than the order pay date. pt: data de emissão da NF-e.
IssueDate int `json:"issue_date,omitempty"`
// total_value is The total value of the invoice. pt: valor total da NF-e (R$).
TotalValue float64 `json:"total_value,omitempty"`
// products_total_value is The products total value of the invoice. pt: valor total dos produtos (R$) da NF-e.
ProductsTotalValue float64 `json:"products_total_value,omitempty"`
// tax_code is The tax code for the invoice. The tax code should be 4 digits. pt: Código Fiscal de Operações e Prestações (CFOP) predominante na NF-e. 
TaxCode string `json:"tax_code,omitempty"`
}
//=======================================================
// OrderAddInvoiceDataRequest
//=======================================================
type OrderAddInvoiceDataRequest struct {
    // order_sn is Shopee's unique identifier for an order.	
    OrderSN string `json:"order_sn"`
    // invoice_data is The invoice data of the order. pt: Nota Fiscal eletrônica (NF-e) do pedido.
    InvoiceData InvoiceData `json:"invoice_data"`
}
//=======================================================
// OrderAddInvoiceDataResponse
//=======================================================
type OrderAddInvoiceDataResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - OrderGetPendingBuyerInvoiceOrderListOrderList
//=======================================================
type OrderGetPendingBuyerInvoiceOrderListOrderList struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
}


//=======================================================
// Object Raw Type - OrderGetPendingBuyerInvoiceOrderList
//=======================================================
type OrderGetPendingBuyerInvoiceOrderList struct {
// more is This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
More bool `json:"more,omitempty"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
// order_list is 
OrderList []OrderGetPendingBuyerInvoiceOrderListOrderList `json:"order_list"`
}
//=======================================================
// OrderGetPendingBuyerInvoiceOrderListRequest
//=======================================================
type OrderGetPendingBuyerInvoiceOrderListRequest struct {
    // cursor is Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.
    Cursor string `json:"cursor,omitempty"`
    // page_size is Each result set is returned as a page of entries. Use the "page_size" filters to control the maximum number of entries to retrieve per page (i.e., per call). This integer value is used to specify the maximum number of entries to return in a single "page" of data.The limit of page_size if between 1 and 100.
    PageSize int `json:"page_size"`
}
//=======================================================
// OrderGetPendingBuyerInvoiceOrderListResponse
//=======================================================
type OrderGetPendingBuyerInvoiceOrderListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response OrderGetPendingBuyerInvoiceOrderList `json:"response"`
}
//=======================================================
// OrderUploadInvoiceDocRequest
//=======================================================
type OrderUploadInvoiceDocRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
    // file_type is <p>the type of invoice file. 1:pdf 2.jpeg 3.png.&nbsp;For CO only accepts 1:pdf</p>
    FileType string `json:"file_type"`
}
//=======================================================
// OrderUploadInvoiceDocResponse
//=======================================================
type OrderUploadInvoiceDocResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}
//=======================================================
// OrderDownloadInvoiceDocRequest
//=======================================================
type OrderDownloadInvoiceDocRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
}
//=======================================================
// OrderDownloadInvoiceDocResponse
//=======================================================
type OrderDownloadInvoiceDocResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - Queries
//=======================================================
type Queries struct {
// order_sn is <p>Shopee's unique identifier for an order.Limit 50.<br /></p>
OrderSN string `json:"order_sn,omitempty"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoRequest
//=======================================================
type OrderGetBuyerInvoiceInfoRequest struct {
    // queries is 
    Queries []Queries `json:"queries"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoResponse
//=======================================================
type OrderGetBuyerInvoiceInfoResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}