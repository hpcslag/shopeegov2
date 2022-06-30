package shopeego


//=======================================================
// Object Raw Type - OrderList
//=======================================================
type OrderList struct {
// order_sn is  Shopee's unique identifier for an order.	
OrderSn string `json:"order_sn,omitempty"`
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
OrderList OrderList `json:"order_list"`
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
// Object Raw Type - OrderGetShipmentList
//=======================================================
type OrderGetShipmentList struct {
// order_list is The list of  shipment orders
OrderList OrderList `json:"order_list"`
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
// Object Raw Type - OrderGetOrderDetail
//=======================================================
type OrderGetOrderDetail struct {
// order_list is The list of orders.
OrderList OrderList `json:"order_list"`
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
// Object Raw Type - PackageList
//=======================================================
type PackageList struct {
// item_list is The list of items under the same package.
ItemList ItemList `json:"item_list"`
}


//=======================================================
// Object Raw Type - OrderSplitOrder
//=======================================================
type OrderSplitOrder struct {
// order_sn is Shopee's unique identifier for an order.
OrderSn string `json:"order_sn,omitempty"`
// package_list is The list of package under this order you have split.
PackageList PackageList `json:"package_list"`
}
//=======================================================
// OrderSplitOrderRequest
//=======================================================
type OrderSplitOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_list is The list of packages that you want to split
    PackageList PackageList `json:"package_list"`
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
    OrderSn string `json:"order_sn"`
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
    OrderSn string `json:"order_sn"`
    // cancel_reason is The reason seller want to cancel this order. Applicable values: OUT_OF_STOCK, CUSTOMER_REQUEST, UNDELIVERABLE_AREA, COD_NOT_SUPPORTED.
    CancelReason string `json:"cancel_reason"`
    // item_list is Required when cancel_reason is OUT_OF_STOCK. 
    ItemList ItemList `json:"item_list,omitempty"`
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
    OrderSn string `json:"order_sn"`
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
    OrderSn string `json:"order_sn"`
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
    OrderSn string `json:"order_sn"`
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
// Object Raw Type - OrderGetPendingBuyerInvoiceOrderList
//=======================================================
type OrderGetPendingBuyerInvoiceOrderList struct {
// more is This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
More bool `json:"more,omitempty"`
// next_cursor is If more is true, you should pass the next_cursor in the next request as cursor. The value of next_cursor will be empty string when more is false.
NextCursor string `json:"next_cursor,omitempty"`
// order_list is 
OrderList OrderList `json:"order_list"`
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
    OrderSn string `json:"order_sn"`
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
    OrderSn string `json:"order_sn"`
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
OrderSn string `json:"order_sn,omitempty"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoRequest
//=======================================================
type OrderGetBuyerInvoiceInfoRequest struct {
    // queries is 
    Queries Queries `json:"queries"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoResponse
//=======================================================
type OrderGetBuyerInvoiceInfoResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}