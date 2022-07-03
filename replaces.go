package shopeego

var replaces []string = []string{
	"volume",
	"weight",
	"width",
	"length",
	"height",
	"price",
	"refund_amount",
	"amount_before_discount",
	"target",
	"my",
	"unit_price",
	"shipping_fee",
	"original_price",
	"shipping_fee",
	"estimated_shipping_fee",
	"unit_price",
	"rating_star",
	"package_length",
	"package_width",
	"package_height",
	"promotion_price",
	"variation_promotion_price",
	"item_promotion_price",
	"variation_original_price",
	"item_original_price",
	"actual_shipping_cost",
	"total_amount",
	"escrow_amount",
	"coin",
	"voucher",
	"voucher_seller",
	"seller_rebate",
	"shipping_fee_rebate",
	"commission_fee",
	"voucher_code",
	"voucher_name",
	"crooss_border_tax",
	"credit_card_transaction_fee",
	"service_fee",
	"buyer_shopee_kredit",
	"seller_coin_cash_back",
	"final_shipping_fee",
	"seller_return_refund_amount",
	"credit_card_promotion",
	"discounted_price",
	"discount_from_coin",
	"discount_from_voucher",
	"discount_from_voucher_seller",
	"item_price",
	"variation_price",
	"payout_amount",
	"deal_price",
	"item_max_weight",
	"item_min_weight",
	"withdraw_id",
	"root_withdrawal_id",
	"transaction_fee",
	"current_balance",
	"amount",
	"buyer_total_amount",
	"seller_discount",
	"shopee_discount",
	"voucher_from_seller",
	"voucher_from_shopee",
	"coins",
	"buyer_paid_shipping_fee",
	"buyer_transaction_fee",
	"cross_border_tax",
	"payment_promotion",
	"commission_fee",
	"service_fee",
	"seller_transaction_fee",
	"seller_lost_compensation",
	"seller_coin_cash_back",
	"escrow_tax",
	"final_shipping_fee",
	"actual_shipping_fee",
	"shopee_shipping_rebate",
	"shipping_fee_discount_from_3pl",
	"seller_shipping_discount",
	"estimated_shipping_fee",
	"drc_adjustable_refund",
	"cost_of_goods_sold",
	"seller_return_refund",
	"original_cost_of_goods_sold",
	"original_shopee_discount",
	"inflated_price",
	"inflated_original_price",
	"sip_item_price",
	"purchase_min_spend",
	"sub_item_input_price",
	"purchase_min_spend",
	"purchase_min_spend",
	"promo_price",
	"purchase_min_spend",
	"purchase_min_spend",
	"sub_item_input_price",
	"discount_value",
	"fix_price",
	"discount_value",
	"fix_price",
	"discount_value",
	"fix_price",
	"discount_value",
	"fix_price",
	"discount_value",
	"fix_price",
	"model_original_price",
	"model_promotion_price",
	"model_inflated_price_of_original_price",
	"model_inflated_price_of_promotion_price",
	"item_promotion_price",
	"item_original_price",
	"item_inflated_price_of_original_price",
	"item_inflated_price_of_promotion_price",
	"volume",
	"weight",
	"width",
	"length",
	"height",
	"min_spend",
	"discount_amount",
	"max_price",
	"min_limit",
	"max_limit",
	"description_image_aspect_ratio_min",
	"description_image_aspect_ratio_max",
	"text_length_multiplier",
	"original_price",
	"sip_item_price",
	"weight",
	"original_price",
	"weight",
	"weight",
	"original_price",
	"original_price",
	"original_price",
	"shipping_fee",
	"original_price",
	"default_price",
	"item_max_weight",
	"item_min_weight",
	"height",
	"width",
	"length",
	"dimension_sum",
	"item_max_volume",
	"item_min_volume",
	"model_original_price",
	"model_discounted_price",
	"weight",
	"total_value",
	"products_total_value",
	"total_amount",
	"estimated_shipping_fee",
	"actual_shipping_fee",
	"reverse_shipping_fee",
	"total_value",
	"products_total_value",
	"original_price",
	"discounted_price",
	"seller_discount",
	"shopee_discount",
	"discount_from_coin",
	"discount_from_voucher_shopee",
	"discount_from_voucher_seller",
	"escrow_amount",
	"buyer_total_amount",
	"original_price",
	"seller_discount",
	"shopee_discount",
	"voucher_from_seller",
	"voucher_from_shopee",
	"coins",
	"buyer_paid_shipping_fee",
	"buyer_transaction_fee",
	"cross_border_tax",
	"payment_promotion",
	"commission_fee",
	"service_fee",
	"seller_transaction_fee",
	"seller_lost_compensation",
	"seller_coin_cash_back",
	"escrow_tax",
	"final_shipping_fee",
	"actual_shipping_fee",
	"shopee_shipping_rebate",
	"shipping_fee_discount_from_3pl",
	"seller_shipping_discount",
	"estimated_shipping_fee",
	"drc_adjustable_refund",
	"cost_of_goods_sold",
	"original_cost_of_goods_sold",
	"original_shopee_discount",
	"seller_return_refund",
	"escrow_amount_pri",
	"buyer_total_amount_pri",
	"original_price_pri",
	"seller_return_refund_pri",
	"commission_fee_pri",
	"service_fee_pri",
	"drc_adjustable_refund_pri",
	"exchange_rate",
	"reverse_shipping_fee",
	"final_product_protection",
	"credit_card_promotion",
	"credit_card_transaction_fee",
	"final_product_vat_tax",
	"from_amount",
	"payout_amount",
	"escrow_amount",
	"adjustment_amount",
	"amount",
	"current_balance",
	"transaction_fee",
	"payout_amount",
	"min_limit",
	"max_limit",
	"description_image_aspect_ratio_min",
	"description_image_aspect_ratio_max",
	"original_price",
	"current_price",
	"inflated_price_of_original_price",
	"inflated_price_of_current_price",
	"sip_item_price",
	"shipping_fee",
	"estimated_shipping_fee",
	"unit_price",
	"inflated_price_of_unit_price",
	"rating_star",
	"shipping_fee",
	"unit_price",
	"current_price",
	"original_price",
	"shipping_fee",
	"unit_price",
	"original_price",
	"weight",
	"estimated_shipping_fee",
	"weight",
	"weight",
	"original_price",
	"original_price",
	"current_price",
	"original_price",
	"inflated_price_of_original_price",
	"inflated_price_of_current_price",
	"sip_item_price",
	"original_price",
	"original_price",
	"original_price",
	"original_price",
	"promotion_price",
	"sip_item_price",
	"item_price",
	"compensation_amount",
	"latest_offer_amount",
	"refund_amount",
	"amount_before_discount",
	"item_price",
	"refund_amount",
	"amount_before_discount",
	"max_refund_amount",
	"max_refund_amount",
	"proposed_adjusted_refund_amount",
	"current_price",
	"inflated_price_of_current_price",
	"current_price",
	"inflated_price_of_current_price",
	"current_price",
	"inflated_price_of_current_price",
	"min_basket_price",
	"discount_amount",
	"max_price",
	"min_basket_price",
	"discount_amount",
	"max_price",
	"min_basket_price",
	"max_price",
	"discount_amount",
	"discount_amount",
	"target",
	"my",
	"price",
	"unit_price",
	"shipping_fee",
	"price",
	"original_price",
	"shipping_fee",
	"estimated_shipping_fee",
	"unit_price",
	"price",
	"weight",
	"original_price",
	"rating_star",
	"price",
	"price",
	"original_price",
	"price",
	"original_price",
	"inflated_price",
	"inflated_original_price",
	"sip_item_price",
	"shipping_fee",
	"estimated_shipping_fee",
	"unit_price",
	"price",
	"weight",
	"original_price",
	"rating_star",
	"package_length",
	"package_width",
	"package_height",
	"unit_price",
	"shipping_fee",
	"price",
	"original_pirce",
	"shipping_fee",
	"estimated_shipping_fee",
	"unit_price",
	"price",
	"weight",
	"original_price",
	"rating_star",
	"price",
	"price",
	"price",
	"price",
	"promotion_price",
	"variation_promotion_price",
	"item_promotion_price",
	"variation_promotion_price",
	"item_promotion_price",
	"variation_original_price",
	"variation_promotion_price",
	"item_original_price",
	"item_promotion_price",
	"variation_promotion_price",
	"item_original_price",
	"variation_original_price",
	"variation_discounted_price",
	"weight",
	"estimated_shipping_fee",
	"actual_shipping_cost",
	"total_amount",
	"escrow_amount",
	"escrow_tax",
	"total_amount",
	"coin",
	"voucher",
	"voucher_seller",
	"seller_rebate",
	"actual_shipping_cost",
	"shipping_fee_rebate",
	"commission_fee",
	"voucher_code",
	"voucher_name",
	"escrow_amount",
	"crooss_border_tax",
	"credit_card_transaction_fee",
	"service_fee",
	"buyer_shopee_kredit",
	"seller_coin_cash_back",
	"final_shipping_fee",
	"seller_return_refund_amount",
	"credit_card_promotion",
	"seller_transaction_fee",
	"buyer_transaction_fee",
	"original_price",
	"original_price",
	"discounted_price",
	"original_price",
	"discounted_price",
	"discount_from_coin",
	"discount_from_voucher",
	"discount_from_voucher_seller",
	"seller_rebate",
	"deal_price",
	"credit_card_promotion",
	"item_price",
	"variation_price",
	"payout_amount",
	"item_max_weight",
	"item_min_weight",
	"height",
	"width",
	"length",
	"item_price",
	"refund_amount",
	"amount_before_discount",
	"item_price",
	"item_price",
	"item_price",
	"item_price",
	"amount",
	"current_balance",
	"transaction_fee",
	"withdraw_id",
	"root_withdrawal_id",
	"escrow_amount",
	"buyer_total_amount",
	"original_price",
	"seller_discount",
	"shopee_discount",
	"voucher_from_seller",
	"voucher_from_shopee",
	"coins",
	"buyer_paid_shipping_fee",
	"buyer_transaction_fee",
	"cross_border_tax",
	"payment_promotion",
	"commission_fee",
	"service_fee",
	"seller_transaction_fee",
	"seller_lost_compensation",
	"seller_coin_cash_back",
	"escrow_tax",
	"final_shipping_fee",
	"actual_shipping_fee",
	"shopee_shipping_rebate",
	"shipping_fee_discount_from_3pl",
	"seller_shipping_discount",
	"estimated_shipping_fee",
	"drc_adjustable_refund",
	"cost_of_goods_sold",
	"original_cost_of_goods_sold",
	"original_shopee_discount",
	"seller_return_refund",
}
