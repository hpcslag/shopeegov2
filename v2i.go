
package shopeego

type V2I interface {

    ProductGetCategory(*ProductGetCategoryRequest) (*ProductGetCategoryResponse, error)

    ProductGetAttributes(*ProductGetAttributesRequest) (*ProductGetAttributesResponse, error)

    ProductGetBrandList(*ProductGetBrandListRequest) (*ProductGetBrandListResponse, error)

    ProductGetDtsLimit(*ProductGetDtsLimitRequest) (*ProductGetDtsLimitResponse, error)

    ProductGetItemLimit(*ProductGetItemLimitRequest) (*ProductGetItemLimitResponse, error)

    ProductGetItemList(*ProductGetItemListRequest) (*ProductGetItemListResponse, error)

    ProductGetItemBaseInfo(*ProductGetItemBaseInfoRequest) (*ProductGetItemBaseInfoResponse, error)

    ProductGetItemExtraInfo(*ProductGetItemExtraInfoRequest) (*ProductGetItemExtraInfoResponse, error)

    ProductAddItem(*ProductAddItemRequest) (*ProductAddItemResponse, error)

    ProductUpdateItem(*ProductUpdateItemRequest) (*ProductUpdateItemResponse, error)

    ProductDeleteItem(*ProductDeleteItemRequest) (*ProductDeleteItemResponse, error)

    ProductInitTierVariation(*ProductInitTierVariationRequest) (*ProductInitTierVariationResponse, error)

    ProductUpdateTierVariation(*ProductUpdateTierVariationRequest) (*ProductUpdateTierVariationResponse, error)

    ProductGetModelList(*ProductGetModelListRequest) (*ProductGetModelListResponse, error)

    ProductAddModel(*ProductAddModelRequest) (*ProductAddModelResponse, error)

    ProductUpdateModel(*ProductUpdateModelRequest) (*ProductUpdateModelResponse, error)

    ProductDeleteModel(*ProductDeleteModelRequest) (*ProductDeleteModelResponse, error)

    ProductSupportSizeChart(*ProductSupportSizeChartRequest) (*ProductSupportSizeChartResponse, error)

    ProductUpdateSizeChart(*ProductUpdateSizeChartRequest) (*ProductUpdateSizeChartResponse, error)

    ProductUnlistItem(*ProductUnlistItemRequest) (*ProductUnlistItemResponse, error)

    ProductUpdatePrice(*ProductUpdatePriceRequest) (*ProductUpdatePriceResponse, error)

    ProductUpdateStock(*ProductUpdateStockRequest) (*ProductUpdateStockResponse, error)

    ProductBoostItem(*ProductBoostItemRequest) (*ProductBoostItemResponse, error)

    ProductGetBoostedList(*ProductGetBoostedListRequest) (*ProductGetBoostedListResponse, error)

    ProductGetItemPromotion(*ProductGetItemPromotionRequest) (*ProductGetItemPromotionResponse, error)

    ProductUpdateSipItemPrice(*ProductUpdateSipItemPriceRequest) (*ProductUpdateSipItemPriceResponse, error)

    ProductSearchItem(*ProductSearchItemRequest) (*ProductSearchItemResponse, error)

    ProductGetComment(*ProductGetCommentRequest) (*ProductGetCommentResponse, error)

    ProductReplyComment(*ProductReplyCommentRequest) (*ProductReplyCommentResponse, error)

    ProductCategoryRecommend(*ProductCategoryRecommendRequest) (*ProductCategoryRecommendResponse, error)

    ProductRegisterBrand(*ProductRegisterBrandRequest) (*ProductRegisterBrandResponse, error)

    ProductGetRecommendAttribute(*ProductGetRecommendAttributeRequest) (*ProductGetRecommendAttributeResponse, error)

    GlobalProductGetCategory(*GlobalProductGetCategoryRequest) (*GlobalProductGetCategoryResponse, error)

    GlobalProductGetAttributes(*GlobalProductGetAttributesRequest) (*GlobalProductGetAttributesResponse, error)

    GlobalProductGetBrandList(*GlobalProductGetBrandListRequest) (*GlobalProductGetBrandListResponse, error)

    GlobalProductGetGlobalItemLimit(*GlobalProductGetGlobalItemLimitRequest) (*GlobalProductGetGlobalItemLimitResponse, error)

    GlobalProductGetDtsLimit(*GlobalProductGetDtsLimitRequest) (*GlobalProductGetDtsLimitResponse, error)

    GlobalProductGetGlobalItemList(*GlobalProductGetGlobalItemListRequest) (*GlobalProductGetGlobalItemListResponse, error)

    GlobalProductGetGlobalItemInfo(*GlobalProductGetGlobalItemInfoRequest) (*GlobalProductGetGlobalItemInfoResponse, error)

    GlobalProductAddGlobalItem(*GlobalProductAddGlobalItemRequest) (*GlobalProductAddGlobalItemResponse, error)

    GlobalProductUpdateGlobalItem(*GlobalProductUpdateGlobalItemRequest) (*GlobalProductUpdateGlobalItemResponse, error)

    GlobalProductDeleteGlobalItem(*GlobalProductDeleteGlobalItemRequest) (*GlobalProductDeleteGlobalItemResponse, error)

    GlobalProductInitTierVariation(*GlobalProductInitTierVariationRequest) (*GlobalProductInitTierVariationResponse, error)

    GlobalProductUpdateTierVariation(*GlobalProductUpdateTierVariationRequest) (*GlobalProductUpdateTierVariationResponse, error)

    GlobalProductAddGlobalModel(*GlobalProductAddGlobalModelRequest) (*GlobalProductAddGlobalModelResponse, error)

    GlobalProductUpdateGlobalModel(*GlobalProductUpdateGlobalModelRequest) (*GlobalProductUpdateGlobalModelResponse, error)

    GlobalProductDeleteGlobalModel(*GlobalProductDeleteGlobalModelRequest) (*GlobalProductDeleteGlobalModelResponse, error)

    GlobalProductGetGlobalModelList(*GlobalProductGetGlobalModelListRequest) (*GlobalProductGetGlobalModelListResponse, error)

    GlobalProductSupportSizeChart(*GlobalProductSupportSizeChartRequest) (*GlobalProductSupportSizeChartResponse, error)

    GlobalProductUpdateSizeChart(*GlobalProductUpdateSizeChartRequest) (*GlobalProductUpdateSizeChartResponse, error)

    GlobalProductCreatePublishTask(*GlobalProductCreatePublishTaskRequest) (*GlobalProductCreatePublishTaskResponse, error)

    GlobalProductGetPublishableShop(*GlobalProductGetPublishableShopRequest) (*GlobalProductGetPublishableShopResponse, error)

    GlobalProductGetPublishTaskResult(*GlobalProductGetPublishTaskResultRequest) (*GlobalProductGetPublishTaskResultResponse, error)

    GlobalProductGetPublishedList(*GlobalProductGetPublishedListRequest) (*GlobalProductGetPublishedListResponse, error)

    GlobalProductUpdatePrice(*GlobalProductUpdatePriceRequest) (*GlobalProductUpdatePriceResponse, error)

    GlobalProductUpdateStock(*GlobalProductUpdateStockRequest) (*GlobalProductUpdateStockResponse, error)

    GlobalProductSetSyncField(*GlobalProductSetSyncFieldRequest) (*GlobalProductSetSyncFieldResponse, error)

    GlobalProductGetGlobalItemID(*GlobalProductGetGlobalItemIDRequest) (*GlobalProductGetGlobalItemIDResponse, error)

    GlobalProductCategoryRecommend(*GlobalProductCategoryRecommendRequest) (*GlobalProductCategoryRecommendResponse, error)

    GlobalProductGetRecommendAttribute(*GlobalProductGetRecommendAttributeRequest) (*GlobalProductGetRecommendAttributeResponse, error)

    MediaSpaceInitVideoUpload(*MediaSpaceInitVideoUploadRequest) (*MediaSpaceInitVideoUploadResponse, error)

    MediaSpaceUploadVideoPart(*MediaSpaceUploadVideoPartRequest) (*MediaSpaceUploadVideoPartResponse, error)

    MediaSpaceCompleteVideoUpload(*MediaSpaceCompleteVideoUploadRequest) (*MediaSpaceCompleteVideoUploadResponse, error)

    MediaSpaceGetVideoUploadResult(*MediaSpaceGetVideoUploadResultRequest) (*MediaSpaceGetVideoUploadResultResponse, error)

    MediaSpaceCancelVideoUpload(*MediaSpaceCancelVideoUploadRequest) (*MediaSpaceCancelVideoUploadResponse, error)

    MediaSpaceUploadImage(*MediaSpaceUploadImageRequest) (*MediaSpaceUploadImageResponse, error)

    ShopGetShopInfo(*ShopGetShopInfoRequest) (*ShopGetShopInfoResponse, error)

    ShopGetProfile(*ShopGetProfileRequest) (*ShopGetProfileResponse, error)

    ShopUpdateProfile(*ShopUpdateProfileRequest) (*ShopUpdateProfileResponse, error)

    ShopGetWarehouseDetail(*ShopGetWarehouseDetailRequest) (*ShopGetWarehouseDetailResponse, error)

    MerchantGetMerchantInfo(*MerchantGetMerchantInfoRequest) (*MerchantGetMerchantInfoResponse, error)

    MerchantGetShopListByMerchant(*MerchantGetShopListByMerchantRequest) (*MerchantGetShopListByMerchantResponse, error)

    OrderGetOrderList(*OrderGetOrderListRequest) (*OrderGetOrderListResponse, error)

    OrderGetShipmentList(*OrderGetShipmentListRequest) (*OrderGetShipmentListResponse, error)

    OrderGetOrderDetail(*OrderGetOrderDetailRequest) (*OrderGetOrderDetailResponse, error)

    OrderSplitOrder(*OrderSplitOrderRequest) (*OrderSplitOrderResponse, error)

    OrderUnsplitOrder(*OrderUnsplitOrderRequest) (*OrderUnsplitOrderResponse, error)

    OrderCancelOrder(*OrderCancelOrderRequest) (*OrderCancelOrderResponse, error)

    OrderHandleBuyerCancellation(*OrderHandleBuyerCancellationRequest) (*OrderHandleBuyerCancellationResponse, error)

    OrderSetNote(*OrderSetNoteRequest) (*OrderSetNoteResponse, error)

    OrderAddInvoiceData(*OrderAddInvoiceDataRequest) (*OrderAddInvoiceDataResponse, error)

    OrderGetPendingBuyerInvoiceOrderList(*OrderGetPendingBuyerInvoiceOrderListRequest) (*OrderGetPendingBuyerInvoiceOrderListResponse, error)

    OrderUploadInvoiceDoc(*OrderUploadInvoiceDocRequest) (*OrderUploadInvoiceDocResponse, error)

    OrderDownloadInvoiceDoc(*OrderDownloadInvoiceDocRequest) (*OrderDownloadInvoiceDocResponse, error)

    OrderGetBuyerInvoiceInfo(*OrderGetBuyerInvoiceInfoRequest) (*OrderGetBuyerInvoiceInfoResponse, error)

    LogisticsGetShippingParameter(*LogisticsGetShippingParameterRequest) (*LogisticsGetShippingParameterResponse, error)

    LogisticsGetTrackingNumber(*LogisticsGetTrackingNumberRequest) (*LogisticsGetTrackingNumberResponse, error)

    LogisticsShipOrder(*LogisticsShipOrderRequest) (*LogisticsShipOrderResponse, error)

    LogisticsUpdateShippingOrder(*LogisticsUpdateShippingOrderRequest) (*LogisticsUpdateShippingOrderResponse, error)

    LogisticsGetShippingDocumentParameter(*LogisticsGetShippingDocumentParameterRequest) (*LogisticsGetShippingDocumentParameterResponse, error)

    LogisticsCreateShippingDocument(*LogisticsCreateShippingDocumentRequest) (*LogisticsCreateShippingDocumentResponse, error)

    LogisticsGetShippingDocumentResult(*LogisticsGetShippingDocumentResultRequest) (*LogisticsGetShippingDocumentResultResponse, error)

    LogisticsDownloadShippingDocument(*LogisticsDownloadShippingDocumentRequest) (*LogisticsDownloadShippingDocumentResponse, error)

    LogisticsGetShippingDocumentInfo(*LogisticsGetShippingDocumentInfoRequest) (*LogisticsGetShippingDocumentInfoResponse, error)

    LogisticsGetTrackingInfo(*LogisticsGetTrackingInfoRequest) (*LogisticsGetTrackingInfoResponse, error)

    LogisticsGetAddressList(*LogisticsGetAddressListRequest) (*LogisticsGetAddressListResponse, error)

    LogisticsSetAddressConfig(*LogisticsSetAddressConfigRequest) (*LogisticsSetAddressConfigResponse, error)

    LogisticsDeleteAddress(*LogisticsDeleteAddressRequest) (*LogisticsDeleteAddressResponse, error)

    LogisticsGetChannelList(*LogisticsGetChannelListRequest) (*LogisticsGetChannelListResponse, error)

    LogisticsUpdateChannel(*LogisticsUpdateChannelRequest) (*LogisticsUpdateChannelResponse, error)

    LogisticsBatchShipOrder(*LogisticsBatchShipOrderRequest) (*LogisticsBatchShipOrderResponse, error)

    FirstMileGetUnbindOrderList(*FirstMileGetUnbindOrderListRequest) (*FirstMileGetUnbindOrderListResponse, error)

    FirstMileGetDetail(*FirstMileGetDetailRequest) (*FirstMileGetDetailResponse, error)

    FirstMileGenerateFirstMileTrackingNumber(*FirstMileGenerateFirstMileTrackingNumberRequest) (*FirstMileGenerateFirstMileTrackingNumberResponse, error)

    FirstMileBindFirstMileTrackingNumber(*FirstMileBindFirstMileTrackingNumberRequest) (*FirstMileBindFirstMileTrackingNumberResponse, error)

    FirstMileUnbindFirstMileTrackingNumber(*FirstMileUnbindFirstMileTrackingNumberRequest) (*FirstMileUnbindFirstMileTrackingNumberResponse, error)

    FirstMileGetTrackingNumberList(*FirstMileGetTrackingNumberListRequest) (*FirstMileGetTrackingNumberListResponse, error)

    FirstMileGetWaybill(*FirstMileGetWaybillRequest) (*FirstMileGetWaybillResponse, error)

    FirstMileGetChannelList(*FirstMileGetChannelListRequest) (*FirstMileGetChannelListResponse, error)

    PaymentGetEscrowDetail(*PaymentGetEscrowDetailRequest) (*PaymentGetEscrowDetailResponse, error)

    PaymentSetShopInstallmentStatus(*PaymentSetShopInstallmentStatusRequest) (*PaymentSetShopInstallmentStatusResponse, error)

    PaymentGetShopInstallmentStatus(*PaymentGetShopInstallmentStatusRequest) (*PaymentGetShopInstallmentStatusResponse, error)

    PaymentGetPayoutDetail(*PaymentGetPayoutDetailRequest) (*PaymentGetPayoutDetailResponse, error)

    PaymentSetItemInstallmentStatus(*PaymentSetItemInstallmentStatusRequest) (*PaymentSetItemInstallmentStatusResponse, error)

    PaymentGetItemInstallmentStatus(*PaymentGetItemInstallmentStatusRequest) (*PaymentGetItemInstallmentStatusResponse, error)

    PaymentGetPaymentMethodList(*PaymentGetPaymentMethodListRequest) (*PaymentGetPaymentMethodListResponse, error)

    PaymentGetWalletTransactionList(*PaymentGetWalletTransactionListRequest) (*PaymentGetWalletTransactionListResponse, error)

    PaymentGetEscrowList(*PaymentGetEscrowListRequest) (*PaymentGetEscrowListResponse, error)

    DiscountAddDiscount(*DiscountAddDiscountRequest) (*DiscountAddDiscountResponse, error)

    DiscountAddDiscountItem(*DiscountAddDiscountItemRequest) (*DiscountAddDiscountItemResponse, error)

    DiscountDeleteDiscount(*DiscountDeleteDiscountRequest) (*DiscountDeleteDiscountResponse, error)

    DiscountDeleteDiscountItem(*DiscountDeleteDiscountItemRequest) (*DiscountDeleteDiscountItemResponse, error)

    DiscountGetDiscount(*DiscountGetDiscountRequest) (*DiscountGetDiscountResponse, error)

    DiscountGetDiscountList(*DiscountGetDiscountListRequest) (*DiscountGetDiscountListResponse, error)

    DiscountUpdateDiscount(*DiscountUpdateDiscountRequest) (*DiscountUpdateDiscountResponse, error)

    DiscountUpdateDiscountItem(*DiscountUpdateDiscountItemRequest) (*DiscountUpdateDiscountItemResponse, error)

    DiscountEndDiscount(*DiscountEndDiscountRequest) (*DiscountEndDiscountResponse, error)

    BundleDealAddBundleDeal(*BundleDealAddBundleDealRequest) (*BundleDealAddBundleDealResponse, error)

    BundleDealAddBundleDealItem(*BundleDealAddBundleDealItemRequest) (*BundleDealAddBundleDealItemResponse, error)

    BundleDealGetBundleDealList(*BundleDealGetBundleDealListRequest) (*BundleDealGetBundleDealListResponse, error)

    BundleDealGetBundleDeal(*BundleDealGetBundleDealRequest) (*BundleDealGetBundleDealResponse, error)

    BundleDealGetBundleDealItem(*BundleDealGetBundleDealItemRequest) (*BundleDealGetBundleDealItemResponse, error)

    BundleDealUpdateBundleDeal(*BundleDealUpdateBundleDealRequest) (*BundleDealUpdateBundleDealResponse, error)

    BundleDealUpdateBundleDealItem(*BundleDealUpdateBundleDealItemRequest) (*BundleDealUpdateBundleDealItemResponse, error)

    BundleDealEndBundleDeal(*BundleDealEndBundleDealRequest) (*BundleDealEndBundleDealResponse, error)

    BundleDealDeleteBundleDeal(*BundleDealDeleteBundleDealRequest) (*BundleDealDeleteBundleDealResponse, error)

    BundleDealDeleteBundleDealItem(*BundleDealDeleteBundleDealItemRequest) (*BundleDealDeleteBundleDealItemResponse, error)

    AddOnDealAddAddOnDeal(*AddOnDealAddAddOnDealRequest) (*AddOnDealAddAddOnDealResponse, error)

    AddOnDealAddAddOnDealMainItem(*AddOnDealAddAddOnDealMainItemRequest) (*AddOnDealAddAddOnDealMainItemResponse, error)

    AddOnDealAddAddOnDealSubItem(*AddOnDealAddAddOnDealSubItemRequest) (*AddOnDealAddAddOnDealSubItemResponse, error)

    AddOnDealDeleteAddOnDeal(*AddOnDealDeleteAddOnDealRequest) (*AddOnDealDeleteAddOnDealResponse, error)

    AddOnDealDeleteAddOnDealMainItem(*AddOnDealDeleteAddOnDealMainItemRequest) (*AddOnDealDeleteAddOnDealMainItemResponse, error)

    AddOnDealDeleteAddOnDealSubItem(*AddOnDealDeleteAddOnDealSubItemRequest) (*AddOnDealDeleteAddOnDealSubItemResponse, error)

    AddOnDealGetAddOnDealList(*AddOnDealGetAddOnDealListRequest) (*AddOnDealGetAddOnDealListResponse, error)

    AddOnDealGetAddOnDeal(*AddOnDealGetAddOnDealRequest) (*AddOnDealGetAddOnDealResponse, error)

    AddOnDealGetAddOnDealMainItem(*AddOnDealGetAddOnDealMainItemRequest) (*AddOnDealGetAddOnDealMainItemResponse, error)

    AddOnDealGetAddOnDealSubItem(*AddOnDealGetAddOnDealSubItemRequest) (*AddOnDealGetAddOnDealSubItemResponse, error)

    AddOnDealUpdateAddOnDeal(*AddOnDealUpdateAddOnDealRequest) (*AddOnDealUpdateAddOnDealResponse, error)

    AddOnDealUpdateAddOnDealMainItem(*AddOnDealUpdateAddOnDealMainItemRequest) (*AddOnDealUpdateAddOnDealMainItemResponse, error)

    AddOnDealUpdateAddOnDealSubItem(*AddOnDealUpdateAddOnDealSubItemRequest) (*AddOnDealUpdateAddOnDealSubItemResponse, error)

    AddOnDealEndAddOnDeal(*AddOnDealEndAddOnDealRequest) (*AddOnDealEndAddOnDealResponse, error)

    VoucherAddVoucher(*VoucherAddVoucherRequest) (*VoucherAddVoucherResponse, error)

    VoucherDeleteVoucher(*VoucherDeleteVoucherRequest) (*VoucherDeleteVoucherResponse, error)

    VoucherEndVoucher(*VoucherEndVoucherRequest) (*VoucherEndVoucherResponse, error)

    VoucherUpdateVoucher(*VoucherUpdateVoucherRequest) (*VoucherUpdateVoucherResponse, error)

    VoucherGetVoucher(*VoucherGetVoucherRequest) (*VoucherGetVoucherResponse, error)

    VoucherGetVoucherList(*VoucherGetVoucherListRequest) (*VoucherGetVoucherListResponse, error)

    FollowPrizeAddFollowPrize(*FollowPrizeAddFollowPrizeRequest) (*FollowPrizeAddFollowPrizeResponse, error)

    FollowPrizeDeleteFollowPrize(*FollowPrizeDeleteFollowPrizeRequest) (*FollowPrizeDeleteFollowPrizeResponse, error)

    FollowPrizeEndFollowPrize(*FollowPrizeEndFollowPrizeRequest) (*FollowPrizeEndFollowPrizeResponse, error)

    FollowPrizeUpdateFollowPrize(*FollowPrizeUpdateFollowPrizeRequest) (*FollowPrizeUpdateFollowPrizeResponse, error)

    FollowPrizeGetFollowPrizeDetail(*FollowPrizeGetFollowPrizeDetailRequest) (*FollowPrizeGetFollowPrizeDetailResponse, error)

    FollowPrizeGetFollowPrizeList(*FollowPrizeGetFollowPrizeListRequest) (*FollowPrizeGetFollowPrizeListResponse, error)

    TopPicksGetTopPicksList(*TopPicksGetTopPicksListRequest) (*TopPicksGetTopPicksListResponse, error)

    TopPicksAddTopPicks(*TopPicksAddTopPicksRequest) (*TopPicksAddTopPicksResponse, error)

    TopPicksUpdateTopPicks(*TopPicksUpdateTopPicksRequest) (*TopPicksUpdateTopPicksResponse, error)

    TopPicksDeleteTopPicks(*TopPicksDeleteTopPicksRequest) (*TopPicksDeleteTopPicksResponse, error)

    ShopCategoryAddShopCategory(*ShopCategoryAddShopCategoryRequest) (*ShopCategoryAddShopCategoryResponse, error)

    ShopCategoryGetShopCategoryList(*ShopCategoryGetShopCategoryListRequest) (*ShopCategoryGetShopCategoryListResponse, error)

    ShopCategoryDeleteShopCategory(*ShopCategoryDeleteShopCategoryRequest) (*ShopCategoryDeleteShopCategoryResponse, error)

    ShopCategoryUpdateShopCategory(*ShopCategoryUpdateShopCategoryRequest) (*ShopCategoryUpdateShopCategoryResponse, error)

    ShopCategoryAddItemList(*ShopCategoryAddItemListRequest) (*ShopCategoryAddItemListResponse, error)

    ShopCategoryGetItemList(*ShopCategoryGetItemListRequest) (*ShopCategoryGetItemListResponse, error)

    ShopCategoryDeleteItemList(*ShopCategoryDeleteItemListRequest) (*ShopCategoryDeleteItemListResponse, error)

    ReturnsGetReturnDetail(*ReturnsGetReturnDetailRequest) (*ReturnsGetReturnDetailResponse, error)

    ReturnsGetReturnList(*ReturnsGetReturnListRequest) (*ReturnsGetReturnListResponse, error)

    ReturnsConfirm(*ReturnsConfirmRequest) (*ReturnsConfirmResponse, error)

    ReturnsDispute(*ReturnsDisputeRequest) (*ReturnsDisputeResponse, error)

    ReturnsGetAvailableSolutions(*ReturnsGetAvailableSolutionsRequest) (*ReturnsGetAvailableSolutionsResponse, error)

    ReturnsOffer(*ReturnsOfferRequest) (*ReturnsOfferResponse, error)

    ReturnsAcceptOffer(*ReturnsAcceptOfferRequest) (*ReturnsAcceptOfferResponse, error)

    AccountHealthShopPerformance(*AccountHealthShopPerformanceRequest) (*AccountHealthShopPerformanceResponse, error)

    AccountHealthShopPenalty(*AccountHealthShopPenaltyRequest) (*AccountHealthShopPenaltyResponse, error)

    PublicGetShopsByPartner(*PublicGetShopsByPartnerRequest) (*PublicGetShopsByPartnerResponse, error)

    PublicGetMerchantsByPartner(*PublicGetMerchantsByPartnerRequest) (*PublicGetMerchantsByPartnerResponse, error)

    PublicGetTokenByResendCode(*PublicGetTokenByResendCodeRequest) (*PublicGetTokenByResendCodeResponse, error)

    PublicGetRefreshTokenByUpgradeCode(*PublicGetRefreshTokenByUpgradeCodeRequest) (*PublicGetRefreshTokenByUpgradeCodeResponse, error)

    PushGetPushConfig(*PushGetPushConfigRequest) (*PushGetPushConfigResponse, error)

    PushSetPushConfig(*PushSetPushConfigRequest) (*PushSetPushConfigResponse, error)

}
