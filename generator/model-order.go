package shopeego
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning []string `json:"warning,omitempty"`
}
//=======================================================
// OrderSplitOrderRequest
//=======================================================
type OrderSplitOrderRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
    // package_list is The list of packages that you want to split
    PackageList []interface{} `json:"package_list"`
}
//=======================================================
// OrderSplitOrderResponse
//=======================================================
type OrderSplitOrderResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
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
    ItemList []interface{} `json:"item_list,omitempty"`
}
//=======================================================
// OrderCancelOrderResponse
//=======================================================
type OrderCancelOrderResponse struct {
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
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
    // invoice_doc is a filetype, should parse by http.request

}


//=======================================================
// Object Raw Type - InvoiceDetail
//=======================================================
type InvoiceDetail struct {
// name is <p>Buyer name (only has value when invoice_type is personal)</p>
Name string `json:"name,omitempty"`
// email is <p>Buyer email address&nbsp;(only has value when invoice_type is personal)</p>
Email string `json:"email,omitempty"`
// address is <p>Buyer address in format "Streat &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is personal)<br /></p>
Address string `json:"address,omitempty"`
// company_name is <p>Buyer name&nbsp;(only has value when invoice_type is company)</p>
CompanyName string `json:"company_name,omitempty"`
// company_email is <p>Buyer email address&nbsp;(only has value when invoice_type is company)</p>
CompanyEmail string `json:"company_email,omitempty"`
// company_nip is <p>Corresponds to NIP in PL.&nbsp;(only has value when invoice_type is company)v</p>
CompanyNip string `json:"company_nip,omitempty"`
// company_address is <p>Buyer address in format "Streat &amp; number, city, zipcode, any additional info provided by buyer"&nbsp;(only has value when invoice_type is company)<br /></p>
CompanyAddress string `json:"company_address,omitempty"`
}


//=======================================================
// Object Raw Type - InvoiceInfo
//=======================================================
type InvoiceInfo struct {
// order_sn is Shopee's unique identifier for an order.
OrderSn string `json:"order_sn,omitempty"`
// invoice_type is Type of invoice requested: {1: personal, 2: company}.
InvoiceType string `json:"invoice_type,omitempty"`
// invoice_detail is Invoice info submitted by buyer. Might be masked, e.g. A****b, depending on order status.
InvoiceDetail InvoiceDetail `json:"invoice_detail"`
// error is Error in retrieving the receipt setting of a particular order.
Error string `json:"error,omitempty"`
}


//=======================================================
// Object Raw Type - InvoiceInfoList
//=======================================================
type InvoiceInfoList struct {
// invoice_info is 
InvoiceInfo InvoiceInfo `json:"invoice_info"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoRequest
//=======================================================
type OrderGetBuyerInvoiceInfoRequest struct {
    // queries is 
    Queries []interface{} `json:"queries"`
}
//=======================================================
// OrderGetBuyerInvoiceInfoResponse
//=======================================================
type OrderGetBuyerInvoiceInfoResponse struct {
    // invoice_info_list is 
    InvoiceInfoList InvoiceInfoList `json:"invoice_info_list"`
    // request_id is <p>Request id for debugging purposes</p>
    RequestID string `json:"request_id,omitempty"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
}