
package shopeego
var availablePaths map[string]string = map[string]string{

    "ProductGetCategory": "/api/v2/product/get_category",

    "ProductGetAttributes": "/api/v2/product/get_attributes",

    "ProductGetBrandList": "/api/v2/product/get_brand_list",

    "ProductGetDtsLimit": "/api/v2/product/get_dts_limit",

    "ProductGetItemLimit": "/api/v2/product/get_item_limit",

    "ProductGetItemList": "/api/v2/product/get_item_list",

    "ProductGetItemBaseInfo": "/api/v2/product/get_item_base_info",

    "ProductGetItemExtraInfo": "/api/v2/product/get_item_extra_info",

    "ProductAddItem": "/api/v2/product/add_item",

    "ProductUpdateItem": "/api/v2/product/update_item",

    "ProductDeleteItem": "/api/v2/product/delete_item",

    "ProductInitTierVariation": "/api/v2/product/init_tier_variation",

    "ProductUpdateTierVariation": "/api/v2/product/update_tier_variation",

    "ProductGetModelList": "/api/v2/product/get_model_list",

    "ProductAddModel": "/api/v2/product/add_model",

    "ProductUpdateModel": "/api/v2/product/update_model",

    "ProductDeleteModel": "/api/v2/product/delete_model",

    "ProductSupportSizeChart": "/api/v2/product/support_size_chart",

    "ProductUpdateSizeChart": "/api/v2/product/update_size_chart",

    "ProductUnlistItem": "/api/v2/product/unlist_item",

    "ProductUpdatePrice": "/api/v2/product/update_price",

    "ProductUpdateStock": "/api/v2/product/update_stock",

    "ProductBoostItem": "/api/v2/product/boost_item",

    "ProductGetBoostedList": "/api/v2/product/get_boosted_list",

    "ProductGetItemPromotion": "/api/v2/product/get_item_promotion",

    "ProductUpdateSipItemPrice": "/api/v2/product/update_sip_item_price",

    "ProductSearchItem": "/api/v2/product/search_item",

    "ProductGetComment": "/api/v2/product/get_comment",

    "ProductReplyComment": "/api/v2/product/reply_comment",

    "ProductCategoryRecommend": "/api/v2/product/category_recommend",

    "ProductRegisterBrand": "/api/v2/product/register_brand",

    "ProductGetRecommendAttribute": "/api/v2/product/get_recommend_attribute",

    "GlobalProductGetCategory": "/api/v2/global_product/get_category",

    "GlobalProductGetAttributes": "/api/v2/global_product/get_attributes",

    "GlobalProductGetBrandList": "/api/v2/global_product/get_brand_list",

    "GlobalProductGetGlobalItemLimit": "/api/v2/global_product/get_global_item_limit",

    "GlobalProductGetDtsLimit": "/api/v2/global_product/get_dts_limit",

    "GlobalProductGetGlobalItemList": "/api/v2/global_product/get_global_item_list",

    "GlobalProductGetGlobalItemInfo": "/api/v2/global_product/get_global_item_info",

    "GlobalProductAddGlobalItem": "/api/v2/global_product/add_global_item",

    "GlobalProductUpdateGlobalItem": "/api/v2/global_product/update_global_item",

    "GlobalProductDeleteGlobalItem": "/api/v2/global_product/delete_global_item",

    "GlobalProductInitTierVariation": "/api/v2/global_product/init_tier_variation",

    "GlobalProductUpdateTierVariation": "/api/v2/global_product/update_tier_variation",

    "GlobalProductAddGlobalModel": "/api/v2/global_product/add_global_model",

    "GlobalProductUpdateGlobalModel": "/api/v2/global_product/update_global_model",

    "GlobalProductDeleteGlobalModel": "/api/v2/global_product/delete_global_model",

    "GlobalProductGetGlobalModelList": "/api/v2/global_product/get_global_model_list",

    "GlobalProductSupportSizeChart": "/api/v2/global_product/support_size_chart",

    "GlobalProductUpdateSizeChart": "/api/v2/global_product/update_size_chart",

    "GlobalProductCreatePublishTask": "/api/v2/global_product/create_publish_task",

    "GlobalProductGetPublishableShop": "/api/v2/global_product/get_publishable_shop",

    "GlobalProductGetPublishTaskResult": "/api/v2/global_product/get_publish_task_result",

    "GlobalProductGetPublishedList": "/api/v2/global_product/get_published_list",

    "GlobalProductUpdatePrice": "/api/v2/global_product/update_price",

    "GlobalProductUpdateStock": "/api/v2/global_product/update_stock",

    "GlobalProductSetSyncField": "/api/v2/global_product/set_sync_field",

    "GlobalProductGetGlobalItemID": "/api/v2/global_product/get_global_item_id",

    "GlobalProductCategoryRecommend": "/api/v2/global_product/category_recommend",

    "GlobalProductGetRecommendAttribute": "/api/v2/global_product/get_recommend_attribute",

    "MediaSpaceInitVideoUpload": "/api/v2/media_space/init_video_upload",

    "MediaSpaceUploadVideoPart": "/api/v2/media_space/upload_video_part",

    "MediaSpaceCompleteVideoUpload": "/api/v2/media_space/complete_video_upload",

    "MediaSpaceGetVideoUploadResult": "/api/v2/media_space/get_video_upload_result",

    "MediaSpaceCancelVideoUpload": "/api/v2/media_space/cancel_video_upload",

    "MediaSpaceUploadImage": "/api/v2/media_space/upload_image",

    "ShopGetShopInfo": "/api/v2/shop/get_shop_info",

    "ShopGetProfile": "/api/v2/shop/get_profile",

    "ShopUpdateProfile": "/api/v2/shop/update_profile",

    "ShopGetWarehouseDetail": "/api/v2/shop/get_warehouse_detail",

    "MerchantGetMerchantInfo": "/api/v2/merchant/get_merchant_info",

    "MerchantGetShopListByMerchant": "/api/v2/merchant/get_shop_list_by_merchant",

    "OrderGetOrderList": "/api/v2/order/get_order_list",

    "OrderGetShipmentList": "/api/v2/order/get_shipment_list",

    "OrderGetOrderDetail": "/api/v2/order/get_order_detail",

    "OrderSplitOrder": "/api/v2/order/split_order",

    "OrderUnsplitOrder": "/api/v2/order/unsplit_order",

    "OrderCancelOrder": "/api/v2/order/cancel_order",

    "OrderHandleBuyerCancellation": "/api/v2/order/handle_buyer_cancellation",

    "OrderSetNote": "/api/v2/order/set_note",

    "OrderAddInvoiceData": "/api/v2/order/add_invoice_data",

    "OrderGetPendingBuyerInvoiceOrderList": "/api/v2/order/get_pending_buyer_invoice_order_list",

    "OrderUploadInvoiceDoc": "/api/v2/order/upload_invoice_doc",

    "OrderDownloadInvoiceDoc": "/api/v2/order/download_invoice_doc",

    "OrderGetBuyerInvoiceInfo": "/api/v2/order/get_buyer_invoice_info",

    "LogisticsGetShippingParameter": "/api/v2/logistics/get_shipping_parameter",

    "LogisticsGetTrackingNumber": "/api/v2/logistics/get_tracking_number",

    "LogisticsShipOrder": "/api/v2/logistics/ship_order",

    "LogisticsUpdateShippingOrder": "/api/v2/logistics/update_shipping_order",

    "LogisticsGetShippingDocumentParameter": "/api/v2/logistics/get_shipping_document_parameter",

    "LogisticsCreateShippingDocument": "/api/v2/logistics/create_shipping_document",

    "LogisticsGetShippingDocumentResult": "/api/v2/logistics/get_shipping_document_result",

    "LogisticsDownloadShippingDocument": "/api/v2/logistics/download_shipping_document",

    "LogisticsGetShippingDocumentInfo": "/api/v2/logistics/get_shipping_document_info",

    "LogisticsGetTrackingInfo": "/api/v2/logistics/get_tracking_info",

    "LogisticsGetAddressList": "/api/v2/logistics/get_address_list",

    "LogisticsSetAddressConfig": "/api/v2/logistics/set_address_config",

    "LogisticsDeleteAddress": "/api/v2/logistics/delete_address",

    "LogisticsGetChannelList": "/api/v2/logistics/get_channel_list",

    "LogisticsUpdateChannel": "/api/v2/logistics/update_channel",

    "LogisticsBatchShipOrder": "/api/v2/logistics/batch_ship_order",

    "FirstMileGetUnbindOrderList": "/api/v2/first_mile/get_unbind_order_list",

    "FirstMileGetDetail": "/api/v2/first_mile/get_detail",

    "FirstMileGenerateFirstMileTrackingNumber": "/api/v2/first_mile/generate_first_mile_tracking_number",

    "FirstMileBindFirstMileTrackingNumber": "/api/v2/first_mile/bind_first_mile_tracking_number",

    "FirstMileUnbindFirstMileTrackingNumber": "/api/v2/first_mile/unbind_first_mile_tracking_number",

    "FirstMileGetTrackingNumberList": "/api/v2/first_mile/get_tracking_number_list",

    "FirstMileGetWaybill": "/api/v2/first_mile/get_waybill",

    "FirstMileGetChannelList": "/api/v2/first_mile/get_channel_list",

    "PaymentGetEscrowDetail": "/api/v2/payment/get_escrow_detail",

    "PaymentSetShopInstallmentStatus": "/api/v2/payment/set_shop_installment_status",

    "PaymentGetShopInstallmentStatus": "/api/v2/payment/get_shop_installment_status",

    "PaymentGetPayoutDetail": "/api/v2/payment/get_payout_detail",

    "PaymentSetItemInstallmentStatus": "/api/v2/payment/set_item_installment_status",

    "PaymentGetItemInstallmentStatus": "/api/v2/payment/get_item_installment_status",

    "PaymentGetPaymentMethodList": "/api/v2/payment/get_payment_method_list",

    "PaymentGetWalletTransactionList": "/api/v2/payment/get_wallet_transaction_list",

    "PaymentGetEscrowList": "/api/v2/payment/get_escrow_list",

    "DiscountAddDiscount": "/api/v2/discount/add_discount",

    "DiscountAddDiscountItem": "/api/v2/discount/add_discount_item",

    "DiscountDeleteDiscount": "/api/v2/discount/delete_discount",

    "DiscountDeleteDiscountItem": "/api/v2/discount/delete_discount_item",

    "DiscountGetDiscount": "/api/v2/discount/get_discount",

    "DiscountGetDiscountList": "/api/v2/discount/get_discount_list",

    "DiscountUpdateDiscount": "/api/v2/discount/update_discount",

    "DiscountUpdateDiscountItem": "/api/v2/discount/update_discount_item",

    "DiscountEndDiscount": "/api/v2/discount/end_discount",

    "BundleDealAddBundleDeal": "/api/v2/bundle_deal/add_bundle_deal",

    "BundleDealAddBundleDealItem": "/api/v2/bundle_deal/add_bundle_deal_item",

    "BundleDealGetBundleDealList": "/api/v2/bundle_deal/get_bundle_deal_list",

    "BundleDealGetBundleDeal": "/api/v2/bundle_deal/get_bundle_deal",

    "BundleDealGetBundleDealItem": "/api/v2/bundle_deal/get_bundle_deal_item",

    "BundleDealUpdateBundleDeal": "/api/v2/bundle_deal/update_bundle_deal",

    "BundleDealUpdateBundleDealItem": "/api/v2/bundle_deal/update_bundle_deal_item",

    "BundleDealEndBundleDeal": "/api/v2/bundle_deal/end_bundle_deal",

    "BundleDealDeleteBundleDeal": "/api/v2/bundle_deal/delete_bundle_deal",

    "BundleDealDeleteBundleDealItem": "/api/v2/bundle_deal/delete_bundle_deal_item",

    "AddOnDealAddAddOnDeal": "/api/v2/add_on_deal/add_add_on_deal",

    "AddOnDealAddAddOnDealMainItem": "/api/v2/add_on_deal/add_add_on_deal_main_item",

    "AddOnDealAddAddOnDealSubItem": "/api/v2/add_on_deal/add_add_on_deal_sub_item",

    "AddOnDealDeleteAddOnDeal": "/api/v2/add_on_deal/delete_add_on_deal",

    "AddOnDealDeleteAddOnDealMainItem": "/api/v2/add_on_deal/delete_add_on_deal_main_item",

    "AddOnDealDeleteAddOnDealSubItem": "/api/v2/add_on_deal/delete_add_on_deal_sub_item",

    "AddOnDealGetAddOnDealList": "/api/v2/add_on_deal/get_add_on_deal_list",

    "AddOnDealGetAddOnDeal": "/api/v2/add_on_deal/get_add_on_deal",

    "AddOnDealGetAddOnDealMainItem": "/api/v2/add_on_deal/get_add_on_deal_main_item",

    "AddOnDealGetAddOnDealSubItem": "/api/v2/add_on_deal/get_add_on_deal_sub_item",

    "AddOnDealUpdateAddOnDeal": "/api/v2/add_on_deal/update_add_on_deal",

    "AddOnDealUpdateAddOnDealMainItem": "/api/v2/add_on_deal/update_add_on_deal_main_item",

    "AddOnDealUpdateAddOnDealSubItem": "/api/v2/add_on_deal/update_add_on_deal_sub_item",

    "AddOnDealEndAddOnDeal": "/api/v2/add_on_deal/end_add_on_deal",

    "VoucherAddVoucher": "/api/v2/voucher/add_voucher",

    "VoucherDeleteVoucher": "/api/v2/voucher/delete_voucher",

    "VoucherEndVoucher": "/api/v2/voucher/end_voucher",

    "VoucherUpdateVoucher": "/api/v2/voucher/update_voucher",

    "VoucherGetVoucher": "/api/v2/voucher/get_voucher",

    "VoucherGetVoucherList": "/api/v2/voucher/get_voucher_list",

    "FollowPrizeAddFollowPrize": "/api/v2/follow_prize/add_follow_prize",

    "FollowPrizeDeleteFollowPrize": "/api/v2/follow_prize/delete_follow_prize",

    "FollowPrizeEndFollowPrize": "/api/v2/follow_prize/end_follow_prize",

    "FollowPrizeUpdateFollowPrize": "/api/v2/follow_prize/update_follow_prize",

    "FollowPrizeGetFollowPrizeDetail": "/api/v2/follow_prize/get_follow_prize_detail",

    "FollowPrizeGetFollowPrizeList": "/api/v2/follow_prize/get_follow_prize_list",

    "TopPicksGetTopPicksList": "/api/v2/top_picks/get_top_picks_list",

    "TopPicksAddTopPicks": "/api/v2/top_picks/add_top_picks",

    "TopPicksUpdateTopPicks": "/api/v2/top_picks/update_top_picks",

    "TopPicksDeleteTopPicks": "/api/v2/top_picks/delete_top_picks",

    "ShopCategoryAddShopCategory": "/api/v2/shop_category/add_shop_category",

    "ShopCategoryGetShopCategoryList": "/api/v2/shop_category/get_shop_category_list",

    "ShopCategoryDeleteShopCategory": "/api/v2/shop_category/delete_shop_category",

    "ShopCategoryUpdateShopCategory": "/api/v2/shop_category/update_shop_category",

    "ShopCategoryAddItemList": "/api/v2/shop_category/add_item_list",

    "ShopCategoryGetItemList": "/api/v2/shop_category/get_item_list",

    "ShopCategoryDeleteItemList": "/api/v2/shop_category/delete_item_list",

    "ReturnsGetReturnDetail": "/api/v2/returns/get_return_detail",

    "ReturnsGetReturnList": "/api/v2/returns/get_return_list",

    "ReturnsConfirm": "/api/v2/returns/confirm",

    "ReturnsDispute": "/api/v2/returns/dispute",

    "ReturnsGetAvailableSolutions": "/api/v2/returns/get_available_solutions",

    "ReturnsOffer": "/api/v2/returns/offer",

    "ReturnsAcceptOffer": "/api/v2/returns/accept_offer",

    "AccountHealthShopPerformance": "/api/v2/account_health/shop_performance",

    "AccountHealthShopPenalty": "/api/v2/account_health/shop_penalty",

    "PublicGetShopsByPartner": "/api/v2/public/get_shops_by_partner",

    "PublicGetMerchantsByPartner": "/api/v2/public/get_merchants_by_partner",

    "PublicGetTokenByResendCode": "/api/v2/public/get_token_by_resend_code",

    "PublicGetRefreshTokenByUpgradeCode": "/api/v2/public/get_refresh_token_by_upgrade_code",

    "PushGetPushConfig": "/api/v2/push/get_push_config",

    "PushSetPushConfig": "/api/v2/push/set_push_config",

}
