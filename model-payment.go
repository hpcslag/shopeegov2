package shopeego
//=======================================================
// PaymentGetEscrowDetailRequest
//=======================================================
type PaymentGetEscrowDetailRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSn string `json:"order_sn"`
}
//=======================================================
// PaymentGetEscrowDetailResponse
//=======================================================
type PaymentGetEscrowDetailResponse struct {
    // response is 
    Response Response `json:"response"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is Error message
    Message string `json:"message,omitempty"`
}
//=======================================================
// PaymentSetShopInstallmentStatusRequest
//=======================================================
type PaymentSetShopInstallmentStatusRequest struct {
    // installment_status is 
    InstallmentStatus int `json:"installment_status"`
}
//=======================================================
// PaymentSetShopInstallmentStatusResponse
//=======================================================
type PaymentSetShopInstallmentStatusResponse struct {
    // request_id is 
    RequestID string `json:"request_id,omitempty"`
    // error is 
    Error string `json:"error,omitempty"`
    // message is 
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// PaymentGetShopInstallmentStatusRequest
//=======================================================
type PaymentGetShopInstallmentStatusRequest struct {
}
//=======================================================
// PaymentGetShopInstallmentStatusResponse
//=======================================================
type PaymentGetShopInstallmentStatusResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is Error message
    Message string `json:"message,omitempty"`
    // request_id is Request Id
    RequestID string `json:"request_id,omitempty"`
    // response is The business content of the response.
    Response Response `json:"response"`
}
//=======================================================
// PaymentGetPayoutDetailRequest
//=======================================================
type PaymentGetPayoutDetailRequest struct {
    // page_size is Number of pages returned  max:100
    PageSize int `json:"page_size"`
    // page_no is The page number  min:1  default:1
    PageNo int `json:"page_no"`
    // payout_time_from is Strat time
    PayoutTimeFrom int `json:"payout_time_from"`
    // payout_time_to is End time
    PayoutTimeTo int `json:"payout_time_to"`
}
//=======================================================
// PaymentGetPayoutDetailResponse
//=======================================================
type PaymentGetPayoutDetailResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is Error message
    Message string `json:"message,omitempty"`
    // request_id is The unique id for request.
    RequestID string `json:"request_id,omitempty"`
    // response is The business content of the response
    Response Response `json:"response"`
}
//=======================================================
// PaymentSetItemInstallmentStatusRequest
//=======================================================
type PaymentSetItemInstallmentStatusRequest struct {
    // item_id_list is The id array of the item, Max :100
    ItemIdList []int `json:"item_id_list"`
    // tenure_list is Staged array, TH must be [3,6,10], [] means closed
    TenureList []int `json:"tenure_list"`
    // participate_plan_ahora is Only applicable and required for local AR sellers. 
    ParticipatePlanAhora bool `json:"participate_plan_ahora,omitempty"`
}
//=======================================================
// PaymentSetItemInstallmentStatusResponse
//=======================================================
type PaymentSetItemInstallmentStatusResponse struct {
    // error is Error code
    Error string `json:"error,omitempty"`
    // message is Error message
    Message string `json:"message,omitempty"`
    // request_id is Request id
    RequestID string `json:"request_id,omitempty"`
    // response is The business content of the response
    Response Response `json:"response,omitempty"`
}
//=======================================================
// PaymentGetItemInstallmentStatusRequest
//=======================================================
type PaymentGetItemInstallmentStatusRequest struct {
    // item_id_list is Item id array, Max :100
    ItemIdList []int `json:"item_id_list"`
}
//=======================================================
// PaymentGetItemInstallmentStatusResponse
//=======================================================
type PaymentGetItemInstallmentStatusResponse struct {
    // error is Error Code
    Error string `json:"error,omitempty"`
    // message is Error message
    Message string `json:"message,omitempty"`
    // request_id is Request id
    RequestID string `json:"request_id,omitempty"`
    // response is The business content of the response
    Response Response `json:"response"`
}
//=======================================================
// PaymentGetPaymentMethodListRequest
//=======================================================
type PaymentGetPaymentMethodListRequest struct {
}
//=======================================================
// PaymentGetPaymentMethodListResponse
//=======================================================
type PaymentGetPaymentMethodListResponse struct {
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error message
    Message string `json:"message,omitempty"`
    // request_id is Unique id for request
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response []interface{} `json:"response,omitempty"`
}
//=======================================================
// PaymentGetWalletTransactionListRequest
//=======================================================
type PaymentGetWalletTransactionListRequest struct {
    // page_no is Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
    PageNo int `json:"page_no"`
    // page_size is If many transactions are available to retrieve, you may need to call GetTransactionList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
    PageSize int `json:"page_size"`
    // create_time_from is The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
    CreateTimeFrom int `json:"create_time_from"`
    // create_time_to is The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
    CreateTimeTo int `json:"create_time_to"`
    // wallet_type is This field indicates the wallet type.
    WalletType string `json:"wallet_type,omitempty"`
    // transaction_type is <p>This field indicates the transaction type. More detail can refer Data Definition to check.</p>
    TransactionType string `json:"transaction_type,omitempty"`
}
//=======================================================
// PaymentGetWalletTransactionListResponse
//=======================================================
type PaymentGetWalletTransactionListResponse struct {
    // response is 
    Response Response `json:"response,omitempty"`
    // request_id is 
    RequestID string `json:"request_id,omitempty"`
    // message is 
    Message string `json:"message,omitempty"`
    // error is 
    Error string `json:"error,omitempty"`
}
//=======================================================
// PaymentGetEscrowListRequest
//=======================================================
type PaymentGetEscrowListRequest struct {
    // release_time_from is Query start time
    ReleaseTimeFrom int `json:"release_time_from"`
    // release_time_to is Query end time
    ReleaseTimeTo int `json:"release_time_to"`
    // page_size is Number of pages returned  max:100  default:40
    PageSize int `json:"page_size,omitempty"`
    // page_no is The page number  min:1  default:1
    PageNo int `json:"page_no,omitempty"`
}
//=======================================================
// PaymentGetEscrowListResponse
//=======================================================
type PaymentGetEscrowListResponse struct {
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error message
    Message string `json:"message,omitempty"`
    // request_id is The request id
    RequestID string `json:"request_id,omitempty"`
    // response is The business content of the response
    Response Response `json:"response,omitempty"`
}