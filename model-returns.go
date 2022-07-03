package shopeego


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailUser
//=======================================================
type ReturnsGetReturnDetailUser struct {
// username is Buyer's nickname.
Username string `json:"username,omitempty"`
// email is Buyer's email.
Email string `json:"email,omitempty"`
// portrait is Buyer's portrait.
Portrait string `json:"portrait,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailItem
//=======================================================
type ReturnsGetReturnDetailItem struct {
// model_id is Shopee's unique identifier for a variation of an item.
ModelID int `json:"model_id,omitempty"`
// name is Name of item in local language.
Name string `json:"name,omitempty"`
// images is Image URLs of item.
Images []string `json:"images,omitempty"`
// amount is Amount of this item.
Amount int `json:"amount,omitempty"`
// item_price is The price of item.
ItemPrice float64 `json:"item_price,omitempty,string"`
// is_add_on_deal is To indicate if this item belongs to an addon deal.
IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
// is_main_item is To indicate if this item is main item or sub item. True means main item, false means sub item.
IsMainItem bool `json:"is_main_item,omitempty"`
// add_on_deal_id is The unique identity of an addon deal.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// item_sku is The sku of item.
ItemSku string `json:"item_sku,omitempty"`
// variation_sku is the variation sku of item
VariationSku string `json:"variation_sku,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailActivityItem
//=======================================================
type ReturnsGetReturnDetailActivityItem struct {
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// variation_id is Shopee's unique identifier for a variation of an item.
VariationID int `json:"variation_id,omitempty"`
// quantity_purchased is item's quantity purchase
QuantityPurchased int `json:"quantity_purchased,omitempty"`
// original_price is item's origin price
OriginalPrice string `json:"original_price,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailActivity
//=======================================================
type ReturnsGetReturnDetailActivity struct {
// activity_id is The id of activity.
ActivityID int `json:"activity_id,omitempty"`
// activity_type is The type of activity.
ActivityType string `json:"activity_type,omitempty"`
// original_price is activity's origin price
OriginalPrice string `json:"original_price,omitempty"`
// discounted_price is activity's discount price
DiscountedPrice string `json:"discounted_price,omitempty"`
// items is 
Items []ReturnsGetReturnDetailActivityItem `json:"items"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailSellerProof
//=======================================================
type ReturnsGetReturnDetailSellerProof struct {
// seller_proof_status is To indicate whether the seller needs to provide evidence when the return status is RT4 and RT8. See "Data Definition - SellerProofStatus"
SellerProofStatus string `json:"seller_proof_status,omitempty"`
// seller_evidence_deadline is To indicate the deadline for submitting the evidence
SellerEvidenceDeadline int `json:"seller_evidence_deadline,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailSellerCompensation
//=======================================================
type ReturnsGetReturnDetailSellerCompensation struct {
// seller_compensation_status is To indicate whether the seller is eligible for raising a compensation request. See "Data Definition - SellerCompensationStatus"
SellerCompensationStatus string `json:"seller_compensation_status,omitempty"`
// seller_compensation_due_date is To indicate the deadline for requesting the compensation
SellerCompensationDueDate int `json:"seller_compensation_due_date,omitempty"`
// compensation_amount is To indicate the compensation amount that the agent decided
CompensationAmount float64 `json:"compensation_amount,omitempty,string"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailNegotiation
//=======================================================
type ReturnsGetReturnDetailNegotiation struct {
// negotiation_status is To indicate whether the seller can negotiate with the buyer. See "Data Definition - NegotiationStatus"
NegotiationStatus string `json:"negotiation_status,omitempty"`
// latest_solution is To indicate what is the offer solution. See "Data Definition - ReturnSolution"
LatestSolution string `json:"latest_solution,omitempty"`
// latest_offer_amount is To indicate the refund amount in the latest offer solution
LatestOfferAmount float64 `json:"latest_offer_amount,omitempty,string"`
// latest_offer_creator is To indicate which party made the latest offer
LatestOfferCreator string `json:"latest_offer_creator,omitempty"`
// counter_limit is To indicate the remaining counter limit
CounterLimit int `json:"counter_limit,omitempty"`
// offer_due_date is To indicate offer_due_date
OfferDueDate int `json:"offer_due_date,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetailReturnPickupAddres
//=======================================================
type ReturnsGetReturnDetailReturnPickupAddres struct {
// address is 
Address string `json:"address,omitempty"`
// name is 
Name string `json:"name,omitempty"`
// phone is 
Phone string `json:"phone,omitempty"`
// town is 
Town string `json:"town,omitempty"`
// district is 
District string `json:"district,omitempty"`
// city is 
City string `json:"city,omitempty"`
// state is 
State string `json:"state,omitempty"`
// region is 
Region string `json:"region,omitempty"`
// zipcode is 
Zipcode string `json:"zipcode,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnDetail
//=======================================================
type ReturnsGetReturnDetail struct {
// image is Image URLs of return.
Image []string `json:"image,omitempty"`
// reason is Reason for return product. Applicable values: See Data Definition- ReturnReason.
Reason string `json:"reason,omitempty"`
// text_reason is Reason that buyer provide.
TextReason string `json:"text_reason,omitempty"`
// return_sn is The serial number of return.
ReturnSN int `json:"return_sn,omitempty"`
// refund_amount is Amount of the refund.
RefundAmount float64 `json:"refund_amount,omitempty,string"`
// currency is Currency of the return.
Currency string `json:"currency,omitempty"`
// create_time is The time of return create.
CreateTime int `json:"create_time,omitempty"`
// update_time is The time of modify return.
UpdateTime int `json:"update_time,omitempty"`
// status is Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
Status string `json:"status,omitempty"`
// due_date is The last time seller deal with this return.
DueDate int `json:"due_date,omitempty"`
// tracking_number is The tracking number assigned by the shipping carrier for item shipment.
TrackingNumber string `json:"tracking_number,omitempty"`
// dispute_reason is The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
DisputeReason []string `json:"dispute_reason,omitempty"`
// dispute_text_reason is The reason that seller provide. While the return has been disputed, this field is useful.
DisputeTextReason []string `json:"dispute_text_reason,omitempty"`
// needs_logistics is Items to be sent back to seller. Can be either integrated/non-integrated.
NeedsLogistics bool `json:"needs_logistics,omitempty"`
// amount_before_discount is Order price before discount.
AmountBeforeDiscount float64 `json:"amount_before_discount,omitempty,string"`
// user is 
User ReturnsGetReturnDetailUser `json:"user"`
// item is 
Item []ReturnsGetReturnDetailItem `json:"item"`
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// return_ship_due_date is The due date for buyer to ship order.
ReturnShipDueDate int `json:"return_ship_due_date,omitempty"`
// return_seller_due_date is The due date for seller to deal with this return when buyer have shipped order.
ReturnSellerDueDate int `json:"return_seller_due_date,omitempty"`
// activity is 
Activity []ReturnsGetReturnDetailActivity `json:"activity"`
// seller_proof is 
SellerProof ReturnsGetReturnDetailSellerProof `json:"seller_proof"`
// seller_compensation is 
SellerCompensation ReturnsGetReturnDetailSellerCompensation `json:"seller_compensation"`
// negotiation is 
Negotiation ReturnsGetReturnDetailNegotiation `json:"negotiation"`
// logistics_status is To indicate the reverse logistic status. See "Data Definition - LogisticsStatus"
LogisticsStatus string `json:"logistics_status,omitempty"`
// return_pickup_address is To indicate the buyer's pickup address
ReturnPickupAddress ReturnsGetReturnDetailReturnPickupAddres `json:"return_pickup_address"`
}
//=======================================================
// ReturnsGetReturnDetailRequest
//=======================================================
type ReturnsGetReturnDetailRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
}
//=======================================================
// ReturnsGetReturnDetailResponse
//=======================================================
type ReturnsGetReturnDetailResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Amount of the refund.
    Response ReturnsGetReturnDetail `json:"response"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnListReturnUser
//=======================================================
type ReturnsGetReturnListReturnUser struct {
// username is Buyer's nickname.
Username string `json:"username,omitempty"`
// email is Buyer's email.
Email string `json:"email,omitempty"`
// portrait is Buyer's portrait.
Portrait string `json:"portrait,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnListReturnItem
//=======================================================
type ReturnsGetReturnListReturnItem struct {
// model_id is Shopee's unique identifier for a variation of an item.
ModelID int `json:"model_id,omitempty"`
// name is Name of item in local language.
Name string `json:"name,omitempty"`
// images is Image URLs of item.
Images []string `json:"images,omitempty"`
// amount is Amount of this item.
Amount int `json:"amount,omitempty"`
// item_price is The price of item.
ItemPrice float64 `json:"item_price,omitempty,string"`
// is_add_on_deal is To indicate if this item belongs to an addon deal.
IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
// is_main_item is To indicate if this item is main item or sub item. True means main item, false means sub item.
IsMainItem bool `json:"is_main_item,omitempty"`
// add_on_deal_id is The unique identity of an addon deal.
AddOnDealID int `json:"add_on_deal_id,omitempty"`
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// item_sku is The sku of item.
ItemSku string `json:"item_sku,omitempty"`
// variation_sku is The variation sku of item
VariationSku string `json:"variation_sku,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnListReturn
//=======================================================
type ReturnsGetReturnListReturn struct {
// image is Image URLs of return.
Image []string `json:"image,omitempty"`
// reason is Reason for return product. Applicable values: See Data Definition- ReturnReason.
Reason string `json:"reason,omitempty"`
// text_reason is Reason that buyer provide.
TextReason string `json:"text_reason,omitempty"`
// return_sn is The serial number of return.
ReturnSN int `json:"return_sn,omitempty"`
// refund_amount is Amount of the refund.
RefundAmount float64 `json:"refund_amount,omitempty,string"`
// currency is Currency of the return.
Currency string `json:"currency,omitempty"`
// create_time is The time of return create.
CreateTime int `json:"create_time,omitempty"`
// update_time is The time of modify return.
UpdateTime int `json:"update_time,omitempty"`
// status is Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
Status string `json:"status,omitempty"`
// due_date is The last time seller deal with this return.
DueDate int `json:"due_date,omitempty"`
// tracking_number is The tracking number assigned by the shipping carrier for item shipment.
TrackingNumber string `json:"tracking_number,omitempty"`
// dispute_reason is The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
DisputeReason []string `json:"dispute_reason,omitempty"`
// dispute_text_reason is The reason that seller provide. While the return has been disputed, this field is useful.
DisputeTextReason []string `json:"dispute_text_reason,omitempty"`
// needs_logistics is Items to be sent back to seller. Can be either integrated/non-integrated.
NeedsLogistics bool `json:"needs_logistics,omitempty"`
// amount_before_discount is Order price before discount.
AmountBeforeDiscount float64 `json:"amount_before_discount,omitempty,string"`
// user is 
User ReturnsGetReturnListReturnUser `json:"user"`
// item is 
Item []ReturnsGetReturnListReturnItem `json:"item"`
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// return_ship_due_date is The due date for buyer to ship order.
ReturnShipDueDate int `json:"return_ship_due_date,omitempty"`
// return_seller_due_date is The due date for seller to deal with this return when buyer have shipped order.
ReturnSellerDueDate int `json:"return_seller_due_date,omitempty"`
// negotiation_status is Counter status. See "Data Definition - NegotiationStatus"
NegotiationStatus string `json:"negotiation_status,omitempty"`
// seller_proof_status is Proof status. See "Data Definition - SellerProofStatus"
SellerProofStatus string `json:"seller_proof_status,omitempty"`
// seller_compensation_status is Compensation status. See "Data Definition - SellerCompensationStatus"
SellerCompensationStatus string `json:"seller_compensation_status,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsGetReturnList
//=======================================================
type ReturnsGetReturnList struct {
// more is Whether has next page
More bool `json:"more,omitempty"`
// return is 
Return []ReturnsGetReturnListReturn `json:"return"`
}
//=======================================================
// ReturnsGetReturnListRequest
//=======================================================
type ReturnsGetReturnListRequest struct {
    V2RequestAuthenticationParams
    

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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Amount of the refund.
    Response []ReturnsGetReturnList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsConfirm
//=======================================================
type ReturnsConfirm struct {
// return_sn is The identifier for an API request for error tracking
ReturnSN string `json:"return_sn,omitempty"`
}
//=======================================================
// ReturnsConfirmRequest
//=======================================================
type ReturnsConfirmRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
}
//=======================================================
// ReturnsConfirmResponse
//=======================================================
type ReturnsConfirmResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ReturnsConfirm `json:"response"`
}


//=======================================================
// Object Raw Type - ReturnsDispute
//=======================================================
type ReturnsDispute struct {
// return_sn is The identifier for an API request for error tracking
ReturnSN string `json:"return_sn,omitempty"`
}
//=======================================================
// ReturnsDisputeRequest
//=======================================================
type ReturnsDisputeRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ReturnsDispute `json:"response"`
}


//=======================================================
// Object Raw Type - ReturnsGetAvailableSolutionsOfferReturnRefund
//=======================================================
type ReturnsGetAvailableSolutionsOfferReturnRefund struct {
// eligibility is To indicate whether ReturnRefund solution is available for sellers to select
Eligibility bool `json:"eligibility,omitempty"`
// refund_amount_adjustable is To indicate whether refund is adjustable for ReturnRefund solution
RefundAmountAdjustable bool `json:"refund_amount_adjustable,omitempty"`
// max_refund_amount is The max refund amount for ReturnRefund solution 
MaxRefundAmount float64 `json:"max_refund_amount,omitempty,string"`
}


//=======================================================
// Object Raw Type - ReturnsGetAvailableSolutionsOfferRefund
//=======================================================
type ReturnsGetAvailableSolutionsOfferRefund struct {
// eligibility is To indicate whether Refund solution is available for sellers to select
Eligibility bool `json:"eligibility,omitempty"`
// refund_amount_adjustable is To indicate whether refund is adjustable for Refund solution
RefundAmountAdjustable bool `json:"refund_amount_adjustable,omitempty"`
// max_refund_amount is The max refund amount for ReturnRefund solution 
MaxRefundAmount float64 `json:"max_refund_amount,omitempty,string"`
}


//=======================================================
// Object Raw Type - ReturnsGetAvailableSolutions
//=======================================================
type ReturnsGetAvailableSolutions struct {
// return_sn is 
ReturnSN string `json:"return_sn,omitempty"`
// offer_return_refund is 
OfferReturnRefund ReturnsGetAvailableSolutionsOfferReturnRefund `json:"offer_return_refund"`
// offer_refund is 
OfferRefund ReturnsGetAvailableSolutionsOfferRefund `json:"offer_refund"`
}
//=======================================================
// ReturnsGetAvailableSolutionsRequest
//=======================================================
type ReturnsGetAvailableSolutionsRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
}
//=======================================================
// ReturnsGetAvailableSolutionsResponse
//=======================================================
type ReturnsGetAvailableSolutionsResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ReturnsGetAvailableSolutions `json:"response"`
}


//=======================================================
// Object Raw Type - ReturnsOffer
//=======================================================
type ReturnsOffer struct {
// return_sn is 
ReturnSN string `json:"return_sn,omitempty"`
}
//=======================================================
// ReturnsOfferRequest
//=======================================================
type ReturnsOfferRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
    // proposed_solution is The new solution to be offered. See "Data Definition - ReturnSolution"
    ProposedSolution string `json:"proposed_solution"`
    // proposed_adjusted_refund_amount is The new refund amount to be offered
    ProposedAdjustedRefundAmount float64 `json:"proposed_adjusted_refund_amount,omitempty,string"`
}
//=======================================================
// ReturnsOfferResponse
//=======================================================
type ReturnsOfferResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ReturnsOffer `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnsAcceptOffer
//=======================================================
type ReturnsAcceptOffer struct {
// return_sn is 
ReturnSN string `json:"return_sn,omitempty"`
}
//=======================================================
// ReturnsAcceptOfferRequest
//=======================================================
type ReturnsAcceptOfferRequest struct {
    V2RequestAuthenticationParams
    

    // return_sn is The serial number of return.
    ReturnSN string `json:"return_sn"`
}
//=======================================================
// ReturnsAcceptOfferResponse
//=======================================================
type ReturnsAcceptOfferResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response ReturnsAcceptOffer `json:"response"`
}