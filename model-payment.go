package shopeego


//=======================================================
// Object Raw Type - PaymentGetEscrowDetailOrderIncomeItem
//=======================================================
type PaymentGetEscrowDetailOrderIncomeItem struct {
// item_id is ID of item
ItemID int `json:"item_id,omitempty"`
// item_name is Name of item
ItemName string `json:"item_name,omitempty"`
// item_sku is A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
ItemSku string `json:"item_sku,omitempty"`
// model_id is ID of the model that belongs to the same item.
ModelID int `json:"model_id,omitempty"`
// model_name is Name of the model that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
ModelName string `json:"model_name,omitempty"`
// model_sku is A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
ModelSku string `json:"model_sku,omitempty"`
// original_price is The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
OriginalPrice float64 `json:"original_price,omitempty"`
// discounted_price is The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price. 
DiscountedPrice float64 `json:"discounted_price,omitempty"`
// seller_discount is <p>The discount provided by seller for this item<br /></p>
SellerDiscount float64 `json:"seller_discount,omitempty"`
// shopee_discount is <p>The discount provided by Shopee for this item<br /></p>
ShopeeDiscount float64 `json:"shopee_discount,omitempty"`
// discount_from_coin is  The offset of this item when the buyer consumed Shopee Coins upon checkout. In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0.
DiscountFromCoin float64 `json:"discount_from_coin,omitempty"`
// discount_from_voucher_shopee is The offset of this item when the buyer use Shopee voucher. In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0.
DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee,omitempty"`
// discount_from_voucher_seller is The offset of this item when the buyer use seller-specific voucher. In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0.
DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller,omitempty"`
// activity_type is The type of the item, default is "". If the item is a bundle item the value is "bundle_deal", and if a add on deal item, the value is "add_on_deal"
ActivityType string `json:"activity_type,omitempty"`
// activity_id is If bundle_deal the is id of bundle deal, if add_on_deal this is id of add on deal.
ActivityID int `json:"activity_id,omitempty"`
// is_main_item is Meaning a main or sub item for add_on_deal.
IsMainItem bool `json:"is_main_item,omitempty"`
// quantity_purchased is This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.
QuantityPurchased int `json:"quantity_purchased,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetEscrowDetailOrderIncome
//=======================================================
type PaymentGetEscrowDetailOrderIncome struct {
// escrow_amount is <p>The total amount that the seller is expected to receive for the order and will change before order completed.&nbsp;</p><p>For non cb sip affiliate shop: escrow_amount=buyer_total_amount+shopee_discount+voucher_from_shopee+coins+payment_promotion-buyer_transaction_fee-cross_border_tax-commission_fee-service_fee-seller_transaction_fee-seller_coin_cash_back-escrow_tax-final_product_vat_tax-drc_adjustable_refund+final_shipping_fee(could be postitive/negtive).&nbsp;</p><p>For cb sip affiliate shop:&nbsp;</p><p>escrow_amount=sum of all Asku's settlement price - service_fee - commission_fee -seller_return_refund - drc_adjustable_refund.</p>
EscrowAmount float64 `json:"escrow_amount,omitempty"`
// buyer_total_amount is The total amount that paid by buyer.
BuyerTotalAmount float64 `json:"buyer_total_amount,omitempty"`
// original_price is The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
OriginalPrice float64 `json:"original_price,omitempty"`
// seller_discount is Final sum of each item seller discount of a specific order. (Only display for non cb sip affiliate shop. )
SellerDiscount float64 `json:"seller_discount,omitempty"`
// shopee_discount is Final sum of each item Shopee discount of a specific order. This amount will rebate to seller. (Only display for non cb sip affiliate order. )
ShopeeDiscount float64 `json:"shopee_discount,omitempty"`
// voucher_from_seller is Final value of voucher provided by Seller for the order. (Only display for non cb sip affiliate shop. )
VoucherFromSeller float64 `json:"voucher_from_seller,omitempty"`
// voucher_from_shopee is Final value of voucher provided by Shopee for the order. (Only display for non cb sip affiliate shop. )
VoucherFromShopee float64 `json:"voucher_from_shopee,omitempty"`
// coins is This value indicates the total amount offset when the buyer consumed Shopee Coins upon checkout. (Only display for non cb sip affiliate shop. )
Coins float64 `json:"coins,omitempty"`
// buyer_paid_shipping_fee is The shipping fee paid by buyer. (Only display for non cb sip affiliate shop. )
BuyerPaidShippingFee float64 `json:"buyer_paid_shipping_fee,omitempty"`
// buyer_transaction_fee is Tansaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. )
BuyerTransactionFee float64 `json:"buyer_transaction_fee,omitempty"`
// cross_border_tax is Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund. (Only display for non cb sip affiliate shop. )
CrossBorderTax float64 `json:"cross_border_tax,omitempty"`
// payment_promotion is The amount offset via payment promotion. (Only display for non cb sip affiliate shop. )
PaymentPromotion float64 `json:"payment_promotion,omitempty"`
// commission_fee is The commission fee charged by Shopee platform if applicable.
CommissionFee float64 `json:"commission_fee,omitempty"`
// service_fee is Amount charged by Shopee to seller for additional services.
ServiceFee float64 `json:"service_fee,omitempty"`
// seller_transaction_fee is Tansaction fee paid by seller for the order. (Only display for non cb sip affiliate shop. )
SellerTransactionFee float64 `json:"seller_transaction_fee,omitempty"`
// seller_lost_compensation is Compensation to seller in case of lost parcel. (Only display for non cb sip affiliate shop. )
SellerLostCompensation float64 `json:"seller_lost_compensation,omitempty"`
// seller_coin_cash_back is Value of coins provided by Seller for purchasing with his or her store for the order. (Only display for non cb sip affiliate shop. )
SellerCoinCashBack float64 `json:"seller_coin_cash_back,omitempty"`
// escrow_tax is Cross-border tax imposed by the Indonesian government on sellers. (Only display for non cb sip affiliate shop. )
EscrowTax float64 `json:"escrow_tax,omitempty"`
// final_shipping_fee is Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. (Only display for non cb sip affiliate shop. )
FinalShippingFee float64 `json:"final_shipping_fee,omitempty"`
// actual_shipping_fee is The final shipping cost of order and it is positive. For Non-integrated logistics channel is 0. (Only display for non cb sip affiliate shop. )
ActualShippingFee float64 `json:"actual_shipping_fee,omitempty"`
// order_chargeable_weight is <p>For CB shop, display weight used to calculate actual_shipping_fee for this order.<br /></p>
OrderChargeableWeight int `json:"order_chargeable_weight,omitempty"`
// shopee_shipping_rebate is The platform shipping subsidy to the seller. (Only display for non cb sip affiliate shop. )
ShopeeShippingRebate float64 `json:"shopee_shipping_rebate,omitempty"`
// shipping_fee_discount_from_3pl is The discount of shipping fee from 3PL. Currently only applicable to ID. (Only display for non cb sip affiliate shop. )
ShippingFeeDiscountFrom3Pl float64 `json:"shipping_fee_discount_from_3pl,omitempty"`
// seller_shipping_discount is The shipping discount defined by seller. (Only display for non cb sip affiliate shop. )
SellerShippingDiscount float64 `json:"seller_shipping_discount,omitempty"`
// estimated_shipping_fee is The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard. (Only display for non cb sip affiliate shop. )
EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty"`
// seller_voucher_code is The list of voucher code provided by seller. (Only display for non cb sip affiliate shop. )
SellerVoucherCode []string `json:"seller_voucher_code,omitempty"`
// drc_adjustable_refund is The adjustable refund amount from Shopee Dispute Resolution Center.
DrcAdjustableRefund float64 `json:"drc_adjustable_refund,omitempty"`
// cost_of_goods_sold is Final amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
CostOfGoodsSold float64 `json:"cost_of_goods_sold,omitempty"`
// original_cost_of_goods_sold is Amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
OriginalCostOfGoodsSold float64 `json:"original_cost_of_goods_sold,omitempty"`
// original_shopee_discount is Sum of each item Shopee discount of a specific order. (Only display for non cb sip affiliate shop. )
OriginalShopeeDiscount float64 `json:"original_shopee_discount,omitempty"`
// seller_return_refund is Amount returned to Seller in the event of Partial Return.
SellerReturnRefund float64 `json:"seller_return_refund,omitempty"`
// items is This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
Items []PaymentGetEscrowDetailOrderIncomeItem `json:"items"`
// escrow_amount_pri is The total amount in the prim currency that the seller is expected to receive for the order and will change before order completed . escrow_amount_pri=original_price_pri-seller_return_refund_pri-commission_fee_pri-service_fee_pri-drc_adjustable_refund_pri. (Only display for non cb sip order.)
EscrowAmountPri float64 `json:"escrow_amount_pri,omitempty"`
// buyer_total_amount_pri is The total amount that paid by buyer in the primary currency. (Only display for cb sip affiliate order. )
BuyerTotalAmountPri float64 `json:"buyer_total_amount_pri,omitempty"`
// original_price_pri is The original price of the item before ANY promotion/discount in the primary currency. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate order. )
OriginalPricePri float64 `json:"original_price_pri,omitempty"`
// seller_return_refund_pri is Amount returned to Seller in the event of Partial Return in the primary currency. (Only display for cb sip affiliate shop. )
SellerReturnRefundPri float64 `json:"seller_return_refund_pri,omitempty"`
// commission_fee_pri is The commission fee charged by Shopee platform if applicable in the primary currency. (Only display for cb sip affiliate shop. )
CommissionFeePri float64 `json:"commission_fee_pri,omitempty"`
// service_fee_pri is Amount charged by Shopee to seller for additional services in the primary currency. (Only display for cb sip affiliate shop. )
ServiceFeePri float64 `json:"service_fee_pri,omitempty"`
// drc_adjustable_refund_pri is The adjustable refund amount from Shopee Dispute Resolution Center in the primary currency. (Only display for cb sip affiliate shop. )
DrcAdjustableRefundPri float64 `json:"drc_adjustable_refund_pri,omitempty"`
// pri_currency is The currency of the country where the shop that real seller operates. (Only display for cb sip affiliate shop. )
PriCurrency string `json:"pri_currency,omitempty"`
// aff_currency is The currency of the country where shop opened in. (Only display for cb sip affiliate shop. )
AffCurrency string `json:"aff_currency,omitempty"`
// exchange_rate is Exchange rate from primary shop currency to affiliate shop currency.
ExchangeRate float64 `json:"exchange_rate,omitempty"`
// reverse_shipping_fee is Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
ReverseShippingFee float64 `json:"reverse_shipping_fee,omitempty"`
// final_product_protection is The total amount of product protection purchased during placing an order. (Only display for cb normal and cb sip primary shop)
FinalProductProtection float64 `json:"final_product_protection,omitempty"`
// credit_card_promotion is This value indicate the offset via credit card promotion.
CreditCardPromotion float64 `json:"credit_card_promotion,omitempty"`
// credit_card_transaction_fee is This value indicate the credit card transaction fee.
CreditCardTransactionFee float64 `json:"credit_card_transaction_fee,omitempty"`
// final_product_vat_tax is Value-added Tax is required for online purchases based on EU Value-added Tax regulations . (Only display for non cb sip affiliate shop. )
FinalProductVatTax float64 `json:"final_product_vat_tax,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetEscrowDetail
//=======================================================
type PaymentGetEscrowDetail struct {
// order_sn is  Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// buyer_user_name is The username of buyer.
BuyerUserName string `json:"buyer_user_name,omitempty"`
// return_order_sn_list is The list of the serial number of return.
ReturnOrderSnList []string `json:"return_order_sn_list,omitempty"`
// order_income is 
OrderIncome PaymentGetEscrowDetailOrderIncome `json:"order_income"`
}
//=======================================================
// PaymentGetEscrowDetailRequest
//=======================================================
type PaymentGetEscrowDetailRequest struct {
    // order_sn is Shopee's unique identifier for an order.
    OrderSN string `json:"order_sn"`
}
//=======================================================
// PaymentGetEscrowDetailResponse
//=======================================================
type PaymentGetEscrowDetailResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response PaymentGetEscrowDetail `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentSetShopInstallmentStatus
//=======================================================
type PaymentSetShopInstallmentStatus struct {
// installment_status is 
InstallmentStatus int `json:"installment_status,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response PaymentSetShopInstallmentStatus `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentGetShopInstallmentStatus
//=======================================================
type PaymentGetShopInstallmentStatus struct {
// installment_status is The installment status for the shop
InstallmentStatus int `json:"installment_status,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The business content of the response.
    Response PaymentGetShopInstallmentStatus `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentGetPayoutDetailPayoutPayoutInfo
//=======================================================
type PaymentGetPayoutDetailPayoutPayoutInfo struct {
// from_currency is The settlement currency of orders.
FromCurrency string `json:"from_currency,omitempty"`
// payout_currency is The actual currency of payout.
PayoutCurrency string `json:"payout_currency,omitempty"`
// from_amount is The settlement amount.
FromAmount float64 `json:"from_amount,omitempty"`
// payout_amount is The actual amount of payout.
PayoutAmount float64 `json:"payout_amount,omitempty"`
// exchange_rate is The exchange rate.
ExchangeRate string `json:"exchange_rate,omitempty"`
// payout_time is The time of payout.
PayoutTime int `json:"payout_time,omitempty"`
// pay_service is The service provider of seller. Available value: payoneer, pingpong, lianlian.
PayService string `json:"pay_service,omitempty"`
// payee_id is Seller's account to receive the payout.
PayeeID string `json:"payee_id,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetPayoutDetailPayoutEscrow
//=======================================================
type PaymentGetPayoutDetailPayoutEscrow struct {
// escrow_amount is The total amount that the seller is expected to receive for the order.
EscrowAmount float64 `json:"escrow_amount,omitempty"`
// currency is The currency used for calculating escrow amount.
Currency string `json:"currency,omitempty"`
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetPayoutDetailPayoutOfflineAdjustment
//=======================================================
type PaymentGetPayoutDetailPayoutOfflineAdjustment struct {
// adjustment_amount is The amount of offline adjustments.
AdjustmentAmount float64 `json:"adjustment_amount,omitempty"`
// module is The reason for offline adjustment.
Module string `json:"module,omitempty"`
// remark is The remark for the reason.
Remark string `json:"remark,omitempty"`
// scenario is The scenario of adjustment.
Scenario string `json:"scenario,omitempty"`
// adjustment_level is Dimension of offline adjustment. Available value: shop, order.
AdjustmentLevel string `json:"adjustment_level,omitempty"`
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetPayoutDetailPayout
//=======================================================
type PaymentGetPayoutDetailPayout struct {
// payout_info is The information of payout.
PayoutInfo PaymentGetPayoutDetailPayoutPayoutInfo `json:"payout_info"`
// escrow_list is 
EscrowList []PaymentGetPayoutDetailPayoutEscrow `json:"escrow_list"`
// offline_adjustment_list is The list of offline adjustments.
OfflineAdjustmentList []PaymentGetPayoutDetailPayoutOfflineAdjustment `json:"offline_adjustment_list"`
}


//=======================================================
// Object Raw Type - PaymentGetPayoutDetail
//=======================================================
type PaymentGetPayoutDetail struct {
// more is 
More bool `json:"more,omitempty"`
// payout_list is 
PayoutList []PaymentGetPayoutDetailPayout `json:"payout_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The business content of the response
    Response PaymentGetPayoutDetail `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentSetItemInstallmentStatusItemInstallment
//=======================================================
type PaymentSetItemInstallmentStatusItemInstallment struct {
// item_id is Item unique id
ItemID int `json:"item_id,omitempty"`
// tenure_list is The tenures of item support installment. [] represents with no installment
TenureList []int `json:"tenure_list,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentSetItemInstallmentStatusItemPlanAhora
//=======================================================
type PaymentSetItemInstallmentStatusItemPlanAhora struct {
// item_id is Only applicable for local AR sellers.
ItemID int `json:"item_id,omitempty"`
// participate_plan_ahor is Only applicable for local AR sellers.
ParticipatePlanAhor bool `json:"participate_plan_ahor,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentSetItemInstallmentStatus
//=======================================================
type PaymentSetItemInstallmentStatus struct {
// item_installment_list is 
ItemInstallmentList []PaymentSetItemInstallmentStatusItemInstallment `json:"item_installment_list"`
// item_plan_ahora_list is Only applicable for local AR sellers.
ItemPlanAhoraList []PaymentSetItemInstallmentStatusItemPlanAhora `json:"item_plan_ahora_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The business content of the response
    Response PaymentSetItemInstallmentStatus `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetItemInstallmentStatusItemInstallment
//=======================================================
type PaymentGetItemInstallmentStatusItemInstallment struct {
// item_id is Item unique id
ItemID int `json:"item_id,omitempty"`
// tenure_list is The tenures of item support installment. [] represents with no installment
TenureList []int `json:"tenure_list,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetItemInstallmentStatusItemPlanAhora
//=======================================================
type PaymentGetItemInstallmentStatusItemPlanAhora struct {
// item_id is Only applicable for local AR sellers.
ItemID int `json:"item_id,omitempty"`
// participate_plan_ahora is Only applicable for local AR sellers.
ParticipatePlanAhora bool `json:"participate_plan_ahora,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetItemInstallmentStatus
//=======================================================
type PaymentGetItemInstallmentStatus struct {
// item_installment_list is 
ItemInstallmentList []PaymentGetItemInstallmentStatusItemInstallment `json:"item_installment_list"`
// item_plan_ahora_list is Only applicable for local AR sellers.
ItemPlanAhoraList []PaymentGetItemInstallmentStatusItemPlanAhora `json:"item_plan_ahora_list"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The business content of the response
    Response PaymentGetItemInstallmentStatus `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentGetPaymentMethodList
//=======================================================
type PaymentGetPaymentMethodList struct {
// payment_method is 
PaymentMethod []string `json:"payment_method,omitempty"`
// region is 
Region string `json:"region,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response []PaymentGetPaymentMethodList `json:"response"`
}


//=======================================================
// Object Raw Type - PaymentGetWalletTransactionListTransactionPayOrder
//=======================================================
type PaymentGetWalletTransactionListTransactionPayOrder struct {
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// shop_name is Name of the shop.
ShopName string `json:"shop_name,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetWalletTransactionListTransaction
//=======================================================
type PaymentGetWalletTransactionListTransaction struct {
// status is The status of the transaction，available values: FAILED,COMPLETED,PENDING,INITIAL.
Status string `json:"status,omitempty"`
// transaction_type is The type of transaction.
TransactionType string `json:"transaction_type,omitempty"`
// amount is The amount of transaction.
Amount float64 `json:"amount,omitempty"`
// current_balance is The current balance of this account.
CurrentBalance float64 `json:"current_balance,omitempty"`
// create_time is The create time of the transaction.
CreateTime int `json:"create_time,omitempty"`
// order_sn is Shopee's unique identifier for an order.
OrderSN string `json:"order_sn,omitempty"`
// refund_sn is The serial number of return.
RefundSN string `json:"refund_sn,omitempty"`
// withdrawal_type is The type of withdrawal.
WithdrawalType string `json:"withdrawal_type,omitempty"`
// transaction_fee is This field indicates the transaction fee.
TransactionFee float64 `json:"transaction_fee,omitempty"`
// description is The detailed description of TOPUP SUCCESS and TOPUP FAILED.
Description string `json:"description,omitempty"`
// buyer_name is The name of buyer.
BuyerName string `json:"buyer_name,omitempty"`
// pay_order_list is 
PayOrderList []PaymentGetWalletTransactionListTransactionPayOrder `json:"pay_order_list"`
// shop_name is Name of the shop.
ShopName string `json:"shop_name,omitempty"`
// withdraw_id is Withdraw ID when transaction type is withdraw_created, withdrawal_completed, withdrawal_cancelled.
WithdrawID int `json:"withdraw_id,omitempty"`
// reason is The reason for ADJUSTMENT_ADD and ADJUSTMENT_MINUS.
Reason string `json:"reason,omitempty"`
// root_withdrawal_id is Use this field to indicate the event where a withdrawal is split into several withdrawals due to the withdrawal limit.
RootWithdrawalID int `json:"root_withdrawal_id,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetWalletTransactionList
//=======================================================
type PaymentGetWalletTransactionList struct {
// transaction_list is 
TransactionList []PaymentGetWalletTransactionListTransaction `json:"transaction_list"`
// more is 
More bool `json:"more,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response PaymentGetWalletTransactionList `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetEscrowListEscrow
//=======================================================
type PaymentGetEscrowListEscrow struct {
// order_sn is 
OrderSN string `json:"order_sn,omitempty"`
// payout_amount is The settlement amount
PayoutAmount float64 `json:"payout_amount,omitempty"`
// escrow_release_time is The release time
EscrowReleaseTime int `json:"escrow_release_time,omitempty"`
}


//=======================================================
// Object Raw Type - PaymentGetEscrowList
//=======================================================
type PaymentGetEscrowList struct {
// escrow_list is 
EscrowList []PaymentGetEscrowListEscrow `json:"escrow_list"`
// more is 
More bool `json:"more,omitempty"`
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
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is The business content of the response
    Response PaymentGetEscrowList `json:"response,omitempty"`
}