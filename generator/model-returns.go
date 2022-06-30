package shopeego
//=======================================================
// ReturnsGetReturnDetailRequest
//=======================================================
type ReturnsGetReturnDetailRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
}
//=======================================================
// ReturnsGetReturnDetailResponse
//=======================================================
type ReturnsGetReturnDetailResponse struct {
    // request_id is  The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error description
    Message string `json:"message,omitempty"`
    // response is Amount of the refund.
    Response Response `json:"response"`
}
//=======================================================
// ReturnsGetReturnListRequest
//=======================================================
type ReturnsGetReturnListRequest struct {
    // page_no is Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
    PageNo int `json:"page_no"`
    // page_size is if many items are available to retrieve, you may need to call GetReturnList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
    PageSize int `json:"page_size"`
    // create_time_from is The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
    CreateTimeFrom int `json:"create_time_from,omitempty"`
    // create_time_to is The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
    CreateTimeTo int `json:"create_time_to,omitempty"`
    // status is This is for filtering return request by return status. See "Data Definition - ReturnStatus"
    Status string `json:"status,omitempty"`
    // negotiation_status is This is for filtering return request by counter status. See "Data Definition - NegotiationStatus"
    NegotiationStatus string `json:"negotiation_status,omitempty"`
    // seller_proof_status is This is for filtering return request by proof status. See "Data Definition - SellerProofStatus"
    SellerProofStatus string `json:"seller_proof_status,omitempty"`
    // seller_compensation_status is This is for filtering return request by compensation status. See "Data Definition - SellerCompensationStatus"
    SellerCompensationStatus string `json:"seller_compensation_status,omitempty"`
}
//=======================================================
// ReturnsGetReturnListResponse
//=======================================================
type ReturnsGetReturnListResponse struct {
    // request_id is  The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error description
    Message string `json:"message,omitempty"`
    // response is Amount of the refund.
    Response []interface{} `json:"response,omitempty"`
}
//=======================================================
// ReturnsConfirmRequest
//=======================================================
type ReturnsConfirmRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
}
//=======================================================
// ReturnsConfirmResponse
//=======================================================
type ReturnsConfirmResponse struct {
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error description
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ReturnsDisputeRequest
//=======================================================
type ReturnsDisputeRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
    // email is 
    Email string `json:"email"`
    // dispute_reason is 
    DisputeReason string `json:"dispute_reason"`
    // dispute_text_reason is 
    DisputeTextReason string `json:"dispute_text_reason"`
    // image is 
    Image []string `json:"image"`
}
//=======================================================
// ReturnsDisputeResponse
//=======================================================
type ReturnsDisputeResponse struct {
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is error code
    Error string `json:"error,omitempty"`
    // message is error description
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ReturnsGetAvailableSolutionsRequest
//=======================================================
type ReturnsGetAvailableSolutionsRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
}
//=======================================================
// ReturnsGetAvailableSolutionsResponse
//=======================================================
type ReturnsGetAvailableSolutionsResponse struct {
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The error code
    Error string `json:"error,omitempty"`
    // message is The error description
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}
//=======================================================
// ReturnsOfferRequest
//=======================================================
type ReturnsOfferRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
    // proposed_solution is The new solution to be offered. See "Data Definition - ReturnSolution"
    ProposedSolution string `json:"proposed_solution"`
    // proposed_adjusted_refund_amount is The new refund amount to be offered
    ProposedAdjustedRefundAmount float64 `json:"proposed_adjusted_refund_amount,omitempty"`
}
//=======================================================
// ReturnsOfferResponse
//=======================================================
type ReturnsOfferResponse struct {
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The error code
    Error string `json:"error,omitempty"`
    // message is The error description
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}
//=======================================================
// ReturnsAcceptOfferRequest
//=======================================================
type ReturnsAcceptOfferRequest struct {
    // return_sn is The serial number of return.
    ReturnSn string `json:"return_sn"`
}
//=======================================================
// ReturnsAcceptOfferResponse
//=======================================================
type ReturnsAcceptOfferResponse struct {
    // request_id is The identifier for an API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // error is The error code
    Error string `json:"error,omitempty"`
    // message is The error description
    Message string `json:"message,omitempty"`
    // response is 
    Response Response `json:"response"`
}