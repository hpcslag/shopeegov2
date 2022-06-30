package shopeego


//=======================================================
// Object Raw Type - TotalData
//=======================================================
type TotalData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - SpamListingData
//=======================================================
type SpamListingData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - CounterfeitIpInfringementData
//=======================================================
type CounterfeitIpInfringementData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - ProhibitedListingData
//=======================================================
type ProhibitedListingData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - SevereListingViolations
//=======================================================
type SevereListingViolations struct {
// total_data is Overall data of severe listing violation
TotalData TotalData `json:"total_data"`
// spam_listing_data is Spam listings contain irrelevant and misleading information used to manipulate search results. This includes: - Duplicate listings (products with no significant difference) - Price spam listings (setting very low / high prices with no intention to sell) - Keyword spam listings (irrelevant keywords in product title) - Attribute spam listings (incorrect product information)
SpamListingData SpamListingData `json:"spam_listing_data"`
// counterfeit_ip_infringement_data is Sellers should only list authentic products. Counterfeit items are products that were made in exact imitation of an existing brand with the intention to deceive or defraud, and may include, but are not limited to: - Products that are fake or replicas of an existing official product - Products that have never been produced by a specific brand - Products that bear such similarities with other products (e.g. a replica of a branded item with or without altered logos) without the authorization of the trademark owner
CounterfeitIpInfringementData CounterfeitIpInfringementData `json:"counterfeit_ip_infringement_data"`
// prohibited_listing_data is Prohibited listings are products that are not allowed on Shopee. This includes: - Listings that do not comply with local regulations or Shopee’s Listing Policy; - Listings that advertise for services / products not intended for sale on Shopee - Listings that intent to direct transaction outside of Shopee - Listings edited to sell a different item from the original listing, to mislead buyers on actual sold count or product review
ProhibitedListingData ProhibitedListingData `json:"prohibited_listing_data"`
}


//=======================================================
// Object Raw Type - PercentPreOrderListingData
//=======================================================
type PercentPreOrderListingData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - NumberOfDaysPreOrderListingExceedsTargetDa
//=======================================================
type NumberOfDaysPreOrderListingExceedsTargetDa struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - PreOrderListing
//=======================================================
type PreOrderListing struct {
// percent_pre_order_listing_data is 
PercentPreOrderListingData PercentPreOrderListingData `json:"percent_pre_order_listing_data"`
// number_of_days_pre_order_listing_exceeds_target_da is The number of days in the last 30 days where pre-order listing precentage exceeds target 
NumberOfDaysPreOrderListingExceedsTargetDa NumberOfDaysPreOrderListingExceedsTargetDa `json:"number_of_days_pre_order_listing_exceeds_target_da"`
}


//=======================================================
// Object Raw Type - OtherListingViolations
//=======================================================
type OtherListingViolations struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - ListingViolations
//=======================================================
type ListingViolations struct {
// severe_listing_violations is Listings are deleted if they constitute a serious violation of Shopee's policies. Sellers who violate our policy will earn penalty points under the Seller Penalty Points system. From 7th May 2018 onwards, sellers with a high number of spam, prohibited and/or counterfeit listings will receive 2 penalty points.
SevereListingViolations SevereListingViolations `json:"severe_listing_violations"`
// pre_order_listing is Live pre-order listings as a percentage of total live listings.
PreOrderListing PreOrderListing `json:"pre_order_listing"`
// other_listing_violations is Listings may also be suspended or deleted if they: - Are listed in the wrong categories - Have poor quality images - Have inaccurate description or insufficient details as required by local regulations - Do not meet Shopee's listings requirements
OtherListingViolations OtherListingViolations `json:"other_listing_violations"`
}


//=======================================================
// Object Raw Type - CancellationRateData
//=======================================================
type CancellationRateData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - ReturnRefundRateData
//=======================================================
type ReturnRefundRateData struct {
// target is 
Target string `json:"target,omitempty"`
// my_shop_performance is 
MyShopPerformance string `json:"my_shop_performance,omitempty"`
// penalty_points is 
PenaltyPoints string `json:"penalty_points,omitempty"`
}


//=======================================================
// Object Raw Type - NonFulfillmentRate
//=======================================================
type NonFulfillmentRate struct {
// total_data is 
TotalData TotalData `json:"total_data"`
// cancellation_rate_data is Cancellation rate is the percentage of orders (out of total orders) cancelled by a seller in the past 7 days. Buyers initiated cancellations are not included in the calculation.
CancellationRateData CancellationRateData `json:"cancellation_rate_data"`
// return_refund_rate_data is Return-refund rate is the percentage of orders (out of total orders) that were requested by buyers for a return-refund in the past 7 days. Tips
ReturnRefundRateData ReturnRefundRateData `json:"return_refund_rate_data"`
}


//=======================================================
// Object Raw Type - PreparationTime
//=======================================================
type PreparationTime struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - LateShipmentRate
//=======================================================
type LateShipmentRate struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - Fulfillment
//=======================================================
type Fulfillment struct {
// non_fulfillment_rate is Non-fulfilment rate is the percentage of orders (out of total orders) that were either cancelled or returned in the past 7 days. Only orders cancelled by sellers are taken into consideration when computing non-fulfilment rate. Non-fulfilment rate is also the sum of your cancellation rate and return-refund rate.
NonFulfillmentRate NonFulfillmentRate `json:"non_fulfillment_rate"`
// preparation_time is Preparation time is the number of days it takes a seller to prepare and ship out an order.
PreparationTime PreparationTime `json:"preparation_time"`
// late_shipment_rate is Late shipment rate is the percentage of orders (out of total orders) that were shipped late in the past 7 days. You should maintain your late shipment rate at a healthy level of <2%. If your late shipment rate exceeds %, you will receive a penalty point under the Seller Penalty Points system. One additional penalty point will be issued to seller who exceeds % and has a very high number of late shipped orders (>=50 orders).
LateShipmentRate LateShipmentRate `json:"late_shipment_rate"`
}


//=======================================================
// Object Raw Type - ResponseRate
//=======================================================
type ResponseRate struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - ResponseTime
//=======================================================
type ResponseTime struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - CustomerService
//=======================================================
type CustomerService struct {
// response_rate is Chat response rate is the percentage of new chats and offers (out of total) that a seller responds to within 12 hours of receiving them. Auto replies are not included in the chat response rate computation.
ResponseRate ResponseRate `json:"response_rate"`
// response_time is Chat response time is the average time it takes a seller to respond to a buyer's chat message.
ResponseTime ResponseTime `json:"response_time"`
}


//=======================================================
// Object Raw Type - OverallReviewingRate
//=======================================================
type OverallReviewingRate struct {
// total_data is 
TotalData TotalData `json:"total_data"`
}


//=======================================================
// Object Raw Type - CustomerSatisfaction
//=======================================================
type CustomerSatisfaction struct {
// overall_reviewing_rate is Overall review rating is the average of all order ratings submitted by your buyers.
OverallReviewingRate OverallReviewingRate `json:"overall_reviewing_rate"`
}
//=======================================================
// AccountHealthShopPerformanceRequest
//=======================================================
type AccountHealthShopPerformanceRequest struct {
}
//=======================================================
// AccountHealthShopPerformanceResponse
//=======================================================
type AccountHealthShopPerformanceResponse struct {
    // overall_performance is 1-Excellent,  2-Good,  3-Improvement needed,  4-Poor
    OverallPerformance int `json:"overall_performance,omitempty"`
    // listing_violations is The listing related violations
    ListingViolations ListingViolations `json:"listing_violations"`
    // fulfillment is 
    Fulfillment Fulfillment `json:"fulfillment"`
    // customer_service is 
    CustomerService CustomerService `json:"customer_service"`
    // customer_satisfaction is 
    CustomerSatisfaction CustomerSatisfaction `json:"customer_satisfaction"`
}


//=======================================================
// Object Raw Type - PenaltyPoints
//=======================================================
type PenaltyPoints struct {
// overall_penalty_points is The overall penalty points.
OverallPenaltyPoints int `json:"overall_penalty_points,omitempty"`
// non_fulfillment_rate is The penalty points caused by non-fulfilment orders.
NonFulfillmentRate int `json:"non_fulfillment_rate,omitempty"`
// late_shipment_rate is The penalty points caused by late shipment orders.
LateShipmentRate int `json:"late_shipment_rate,omitempty"`
// listing_violations is The penalty points caused by listing violations.
ListingViolations int `json:"listing_violations,omitempty"`
// others is 
Others int `json:"others,omitempty"`
}
//=======================================================
// AccountHealthShopPenaltyRequest
//=======================================================
type AccountHealthShopPenaltyRequest struct {
}
//=======================================================
// AccountHealthShopPenaltyResponse
//=======================================================
type AccountHealthShopPenaltyResponse struct {
    // penalty_points is Points accumulated will remain on record till the end of a quarter. This will be reset on the first Monday of each quarter
    PenaltyPoints PenaltyPoints `json:"penalty_points"`
    // ongoing_punishment is 
    OngoingPunishment []interface{} `json:"ongoing_punishment,omitempty"`
}