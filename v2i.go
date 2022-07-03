package shopeego

type V2I interface {
	ProductGetCategory(*ProductGetCategoryRequest) (*ProductGetCategory, error)

	ProductGetAttributes(*ProductGetAttributesRequest) (*ProductGetAttributes, error)

	ProductGetBrandList(*ProductGetBrandListRequest) (*ProductGetBrandList, error)

	ProductGetDtsLimit(*ProductGetDtsLimitRequest) (*ProductGetDtsLimit, error)

	ProductGetItemLimit(*ProductGetItemLimitRequest) (*ProductGetItemLimit, error)

	ProductGetItemList(*ProductGetItemListRequest) (*ProductGetItemList, error)

	ProductGetItemBaseInfo(*ProductGetItemBaseInfoRequest) (*ProductGetItemBaseInfo, error)

	ProductGetItemExtraInfo(*ProductGetItemExtraInfoRequest) (*ProductGetItemExtraInfo, error)

	ProductAddItem(*ProductAddItemRequest) (*ProductAddItem, error)

	ProductUpdateItem(*ProductUpdateItemRequest) (*ProductUpdateItem, error)

	ProductDeleteItem(*ProductDeleteItemRequest) error

	ProductInitTierVariation(*ProductInitTierVariationRequest) (*ProductInitTierVariation, error)

	ProductUpdateTierVariation(*ProductUpdateTierVariationRequest) error

	ProductGetModelList(*ProductGetModelListRequest) (*ProductGetModelList, error)

	ProductAddModel(*ProductAddModelRequest) (*ProductAddModel, error)

	ProductUpdateModel(*ProductUpdateModelRequest) error

	ProductDeleteModel(*ProductDeleteModelRequest) error

	ProductSupportSizeChart(*ProductSupportSizeChartRequest) (*ProductSupportSizeChart, error)

	ProductUpdateSizeChart(*ProductUpdateSizeChartRequest) error

	ProductUnlistItem(*ProductUnlistItemRequest) (*ProductUnlistItem, error)

	ProductUpdatePrice(*ProductUpdatePriceRequest) (*ProductUpdatePrice, error)

	ProductUpdateStock(*ProductUpdateStockRequest) (*ProductUpdateStock, error)

	ProductBoostItem(*ProductBoostItemRequest) (*ProductBoostItem, error)

	ProductGetBoostedList(*ProductGetBoostedListRequest) (*ProductGetBoostedList, error)

	ProductGetItemPromotion(*ProductGetItemPromotionRequest) (*ProductGetItemPromotion, error)

	ProductUpdateSipItemPrice(*ProductUpdateSipItemPriceRequest) error

	ProductSearchItem(*ProductSearchItemRequest) (*ProductSearchItem, error)

	ProductGetComment(*ProductGetCommentRequest) (*ProductGetComment, error)

	ProductReplyComment(*ProductReplyCommentRequest) (*ProductReplyComment, error)

	ProductCategoryRecommend(*ProductCategoryRecommendRequest) (*ProductCategoryRecommend, error)

	ProductRegisterBrand(*ProductRegisterBrandRequest) (*ProductRegisterBrand, error)

	ProductGetRecommendAttribute(*ProductGetRecommendAttributeRequest) (*ProductGetRecommendAttribute, error)

	GlobalProductGetCategory(*GlobalProductGetCategoryRequest) (*GlobalProductGetCategory, error)

	GlobalProductGetAttributes(*GlobalProductGetAttributesRequest) (*GlobalProductGetAttributes, error)

	GlobalProductGetBrandList(*GlobalProductGetBrandListRequest) (*GlobalProductGetBrandList, error)

	GlobalProductGetGlobalItemLimit(*GlobalProductGetGlobalItemLimitRequest) (*GlobalProductGetGlobalItemLimit, error)

	GlobalProductGetDtsLimit(*GlobalProductGetDtsLimitRequest) (*GlobalProductGetDtsLimit, error)

	GlobalProductGetGlobalItemList(*GlobalProductGetGlobalItemListRequest) (*GlobalProductGetGlobalItemList, error)

	GlobalProductGetGlobalItemInfo(*GlobalProductGetGlobalItemInfoRequest) (*GlobalProductGetGlobalItemInfo, error)

	GlobalProductAddGlobalItem(*GlobalProductAddGlobalItemRequest) (*GlobalProductAddGlobalItem, error)

	GlobalProductUpdateGlobalItem(*GlobalProductUpdateGlobalItemRequest) (*GlobalProductUpdateGlobalItem, error)

	GlobalProductDeleteGlobalItem(*GlobalProductDeleteGlobalItemRequest) (*GlobalProductDeleteGlobalItem, error)

	GlobalProductInitTierVariation(*GlobalProductInitTierVariationRequest) error

	GlobalProductUpdateTierVariation(*GlobalProductUpdateTierVariationRequest) error

	GlobalProductAddGlobalModel(*GlobalProductAddGlobalModelRequest) error

	GlobalProductUpdateGlobalModel(*GlobalProductUpdateGlobalModelRequest) error

	GlobalProductDeleteGlobalModel(*GlobalProductDeleteGlobalModelRequest) (*GlobalProductDeleteGlobalModel, error)

	GlobalProductGetGlobalModelList(*GlobalProductGetGlobalModelListRequest) (*GlobalProductGetGlobalModelList, error)

	GlobalProductSupportSizeChart(*GlobalProductSupportSizeChartRequest) (*GlobalProductSupportSizeChart, error)

	GlobalProductUpdateSizeChart(*GlobalProductUpdateSizeChartRequest) error

	GlobalProductCreatePublishTask(*GlobalProductCreatePublishTaskRequest) (*GlobalProductCreatePublishTask, error)

	GlobalProductGetPublishableShop(*GlobalProductGetPublishableShopRequest) (*GlobalProductGetPublishableShop, error)

	GlobalProductGetPublishTaskResult(*GlobalProductGetPublishTaskResultRequest) (*GlobalProductGetPublishTaskResult, error)

	GlobalProductGetPublishedList(*GlobalProductGetPublishedListRequest) (*GlobalProductGetPublishedList, error)

	GlobalProductUpdatePrice(*GlobalProductUpdatePriceRequest) error

	GlobalProductUpdateStock(*GlobalProductUpdateStockRequest) error

	GlobalProductSetSyncField(*GlobalProductSetSyncFieldRequest) error

	GlobalProductGetGlobalItemID(*GlobalProductGetGlobalItemIDRequest) (*GlobalProductGetGlobalItemID, error)

	GlobalProductCategoryRecommend(*GlobalProductCategoryRecommendRequest) (*GlobalProductCategoryRecommend, error)

	GlobalProductGetRecommendAttribute(*GlobalProductGetRecommendAttributeRequest) (*GlobalProductGetRecommendAttribute, error)

	MediaSpaceInitVideoUpload(*MediaSpaceInitVideoUploadRequest) (*MediaSpaceInitVideoUpload, error)

	MediaSpaceUploadVideoPart(*MediaSpaceUploadVideoPartRequest) error

	MediaSpaceCompleteVideoUpload(*MediaSpaceCompleteVideoUploadRequest) error

	MediaSpaceGetVideoUploadResult(*MediaSpaceGetVideoUploadResultRequest) (*MediaSpaceGetVideoUploadResult, error)

	MediaSpaceCancelVideoUpload(*MediaSpaceCancelVideoUploadRequest) error

	MediaSpaceUploadImage(*MediaSpaceUploadImageRequest) (*MediaSpaceUploadImage, error)

	ShopGetShopInfo(*ShopGetShopInfoRequest) error

	ShopGetProfile(*ShopGetProfileRequest) (*ShopGetProfile, error)

	ShopUpdateProfile(*ShopUpdateProfileRequest) (*ShopUpdateProfile, error)

	ShopGetWarehouseDetail(*ShopGetWarehouseDetailRequest) ([]ShopGetWarehouseDetail, error)

	MerchantGetMerchantInfo(*MerchantGetMerchantInfoRequest) error

	MerchantGetShopListByMerchant(*MerchantGetShopListByMerchantRequest) error

	OrderGetOrderList(*OrderGetOrderListRequest) (*OrderGetOrderList, error)

	OrderGetShipmentList(*OrderGetShipmentListRequest) (*OrderGetShipmentList, error)

	OrderGetOrderDetail(*OrderGetOrderDetailRequest) (*OrderGetOrderDetail, error)

	OrderSplitOrder(*OrderSplitOrderRequest) (*OrderSplitOrder, error)

	OrderUnsplitOrder(*OrderUnsplitOrderRequest) error

	OrderCancelOrder(*OrderCancelOrderRequest) (*OrderCancelOrder, error)

	OrderHandleBuyerCancellation(*OrderHandleBuyerCancellationRequest) (*OrderHandleBuyerCancellation, error)

	OrderSetNote(*OrderSetNoteRequest) error

	OrderAddInvoiceData(*OrderAddInvoiceDataRequest) error

	OrderGetPendingBuyerInvoiceOrderList(*OrderGetPendingBuyerInvoiceOrderListRequest) (*OrderGetPendingBuyerInvoiceOrderList, error)

	OrderUploadInvoiceDoc(*OrderUploadInvoiceDocRequest) error

	OrderDownloadInvoiceDoc(*OrderDownloadInvoiceDocRequest) error

	OrderGetBuyerInvoiceInfo(*OrderGetBuyerInvoiceInfoRequest) error

	LogisticsGetShippingParameter(*LogisticsGetShippingParameterRequest) (*LogisticsGetShippingParameter, error)

	LogisticsGetTrackingNumber(*LogisticsGetTrackingNumberRequest) (*LogisticsGetTrackingNumber, error)

	LogisticsShipOrder(*LogisticsShipOrderRequest) error

	LogisticsUpdateShippingOrder(*LogisticsUpdateShippingOrderRequest) error

	LogisticsGetShippingDocumentParameter(*LogisticsGetShippingDocumentParameterRequest) (*LogisticsGetShippingDocumentParameter, error)

	LogisticsCreateShippingDocument(*LogisticsCreateShippingDocumentRequest) (*LogisticsCreateShippingDocument, error)

	LogisticsGetShippingDocumentResult(*LogisticsGetShippingDocumentResultRequest) (*LogisticsGetShippingDocumentResult, error)

	LogisticsDownloadShippingDocument(saveFilePath string) func(req *LogisticsDownloadShippingDocumentRequest) (err error)

	LogisticsGetShippingDocumentInfo(*LogisticsGetShippingDocumentInfoRequest) (*LogisticsGetShippingDocumentInfo, error)

	LogisticsGetTrackingInfo(*LogisticsGetTrackingInfoRequest) (*LogisticsGetTrackingInfo, error)

	LogisticsGetAddressList(*LogisticsGetAddressListRequest) (*LogisticsGetAddressList, error)

	LogisticsSetAddressConfig(*LogisticsSetAddressConfigRequest) error

	LogisticsDeleteAddress(*LogisticsDeleteAddressRequest) error

	LogisticsGetChannelList(*LogisticsGetChannelListRequest) (*LogisticsGetChannelList, error)

	LogisticsUpdateChannel(*LogisticsUpdateChannelRequest) (*LogisticsUpdateChannel, error)

	LogisticsBatchShipOrder(*LogisticsBatchShipOrderRequest) (*LogisticsBatchShipOrder, error)

	FirstMileGetUnbindOrderList(*FirstMileGetUnbindOrderListRequest) (*FirstMileGetUnbindOrderList, error)

	FirstMileGetDetail(*FirstMileGetDetailRequest) (*FirstMileGetDetail, error)

	FirstMileGenerateFirstMileTrackingNumber(*FirstMileGenerateFirstMileTrackingNumberRequest) (*FirstMileGenerateFirstMileTrackingNumber, error)

	FirstMileBindFirstMileTrackingNumber(*FirstMileBindFirstMileTrackingNumberRequest) (*FirstMileBindFirstMileTrackingNumber, error)

	FirstMileUnbindFirstMileTrackingNumber(*FirstMileUnbindFirstMileTrackingNumberRequest) (*FirstMileUnbindFirstMileTrackingNumber, error)

	FirstMileGetTrackingNumberList(*FirstMileGetTrackingNumberListRequest) (*FirstMileGetTrackingNumberList, error)

	FirstMileGetWaybill(*FirstMileGetWaybillRequest) error

	FirstMileGetChannelList(*FirstMileGetChannelListRequest) (*FirstMileGetChannelList, error)

	PaymentGetEscrowDetail(*PaymentGetEscrowDetailRequest) (*PaymentGetEscrowDetail, error)

	PaymentSetShopInstallmentStatus(*PaymentSetShopInstallmentStatusRequest) (*PaymentSetShopInstallmentStatus, error)

	PaymentGetShopInstallmentStatus(*PaymentGetShopInstallmentStatusRequest) (*PaymentGetShopInstallmentStatus, error)

	PaymentGetPayoutDetail(*PaymentGetPayoutDetailRequest) (*PaymentGetPayoutDetail, error)

	PaymentSetItemInstallmentStatus(*PaymentSetItemInstallmentStatusRequest) (*PaymentSetItemInstallmentStatus, error)

	PaymentGetItemInstallmentStatus(*PaymentGetItemInstallmentStatusRequest) (*PaymentGetItemInstallmentStatus, error)

	PaymentGetPaymentMethodList(*PaymentGetPaymentMethodListRequest) ([]PaymentGetPaymentMethodList, error)

	PaymentGetWalletTransactionList(*PaymentGetWalletTransactionListRequest) (*PaymentGetWalletTransactionList, error)

	PaymentGetEscrowList(*PaymentGetEscrowListRequest) (*PaymentGetEscrowList, error)

	DiscountAddDiscount(*DiscountAddDiscountRequest) (*DiscountAddDiscount, error)

	DiscountAddDiscountItem(*DiscountAddDiscountItemRequest) (*DiscountAddDiscountItem, error)

	DiscountDeleteDiscount(*DiscountDeleteDiscountRequest) (*DiscountDeleteDiscount, error)

	DiscountDeleteDiscountItem(*DiscountDeleteDiscountItemRequest) (*DiscountDeleteDiscountItem, error)

	DiscountGetDiscount(*DiscountGetDiscountRequest) (*DiscountGetDiscount, error)

	DiscountGetDiscountList(*DiscountGetDiscountListRequest) (*DiscountGetDiscountList, error)

	DiscountUpdateDiscount(*DiscountUpdateDiscountRequest) (*DiscountUpdateDiscount, error)

	DiscountUpdateDiscountItem(*DiscountUpdateDiscountItemRequest) (*DiscountUpdateDiscountItem, error)

	DiscountEndDiscount(*DiscountEndDiscountRequest) (*DiscountEndDiscount, error)

	BundleDealAddBundleDeal(*BundleDealAddBundleDealRequest) (*BundleDealAddBundleDeal, error)

	BundleDealAddBundleDealItem(*BundleDealAddBundleDealItemRequest) (*BundleDealAddBundleDealItem, error)

	BundleDealGetBundleDealList(*BundleDealGetBundleDealListRequest) (*BundleDealGetBundleDealList, error)

	BundleDealGetBundleDeal(*BundleDealGetBundleDealRequest) (*BundleDealGetBundleDeal, error)

	BundleDealGetBundleDealItem(*BundleDealGetBundleDealItemRequest) (*BundleDealGetBundleDealItem, error)

	BundleDealUpdateBundleDeal(*BundleDealUpdateBundleDealRequest) (*BundleDealUpdateBundleDeal, error)

	BundleDealUpdateBundleDealItem(*BundleDealUpdateBundleDealItemRequest) (*BundleDealUpdateBundleDealItem, error)

	BundleDealEndBundleDeal(*BundleDealEndBundleDealRequest) (*BundleDealEndBundleDeal, error)

	BundleDealDeleteBundleDeal(*BundleDealDeleteBundleDealRequest) (*BundleDealDeleteBundleDeal, error)

	BundleDealDeleteBundleDealItem(*BundleDealDeleteBundleDealItemRequest) (*BundleDealDeleteBundleDealItem, error)

	AddOnDealAddAddOnDeal(*AddOnDealAddAddOnDealRequest) (*AddOnDealAddAddOnDeal, error)

	AddOnDealAddAddOnDealMainItem(*AddOnDealAddAddOnDealMainItemRequest) (*AddOnDealAddAddOnDealMainItem, error)

	AddOnDealAddAddOnDealSubItem(*AddOnDealAddAddOnDealSubItemRequest) (*AddOnDealAddAddOnDealSubItem, error)

	AddOnDealDeleteAddOnDeal(*AddOnDealDeleteAddOnDealRequest) (*AddOnDealDeleteAddOnDeal, error)

	AddOnDealDeleteAddOnDealMainItem(*AddOnDealDeleteAddOnDealMainItemRequest) (*AddOnDealDeleteAddOnDealMainItem, error)

	AddOnDealDeleteAddOnDealSubItem(*AddOnDealDeleteAddOnDealSubItemRequest) (*AddOnDealDeleteAddOnDealSubItem, error)

	AddOnDealGetAddOnDealList(*AddOnDealGetAddOnDealListRequest) (*AddOnDealGetAddOnDealList, error)

	AddOnDealGetAddOnDeal(*AddOnDealGetAddOnDealRequest) (*AddOnDealGetAddOnDeal, error)

	AddOnDealGetAddOnDealMainItem(*AddOnDealGetAddOnDealMainItemRequest) (*AddOnDealGetAddOnDealMainItem, error)

	AddOnDealGetAddOnDealSubItem(*AddOnDealGetAddOnDealSubItemRequest) (*AddOnDealGetAddOnDealSubItem, error)

	AddOnDealUpdateAddOnDeal(*AddOnDealUpdateAddOnDealRequest) (*AddOnDealUpdateAddOnDeal, error)

	AddOnDealUpdateAddOnDealMainItem(*AddOnDealUpdateAddOnDealMainItemRequest) (*AddOnDealUpdateAddOnDealMainItem, error)

	AddOnDealUpdateAddOnDealSubItem(*AddOnDealUpdateAddOnDealSubItemRequest) (*AddOnDealUpdateAddOnDealSubItem, error)

	AddOnDealEndAddOnDeal(*AddOnDealEndAddOnDealRequest) (*AddOnDealEndAddOnDeal, error)

	VoucherAddVoucher(*VoucherAddVoucherRequest) (*VoucherAddVoucher, error)

	VoucherDeleteVoucher(*VoucherDeleteVoucherRequest) (*VoucherDeleteVoucher, error)

	VoucherEndVoucher(*VoucherEndVoucherRequest) (*VoucherEndVoucher, error)

	VoucherUpdateVoucher(*VoucherUpdateVoucherRequest) (*VoucherUpdateVoucher, error)

	VoucherGetVoucher(*VoucherGetVoucherRequest) (*VoucherGetVoucher, error)

	VoucherGetVoucherList(*VoucherGetVoucherListRequest) (*VoucherGetVoucherList, error)

	FollowPrizeAddFollowPrize(*FollowPrizeAddFollowPrizeRequest) (*FollowPrizeAddFollowPrize, error)

	FollowPrizeDeleteFollowPrize(*FollowPrizeDeleteFollowPrizeRequest) (*FollowPrizeDeleteFollowPrize, error)

	FollowPrizeEndFollowPrize(*FollowPrizeEndFollowPrizeRequest) (*FollowPrizeEndFollowPrize, error)

	FollowPrizeUpdateFollowPrize(*FollowPrizeUpdateFollowPrizeRequest) (*FollowPrizeUpdateFollowPrize, error)

	FollowPrizeGetFollowPrizeDetail(*FollowPrizeGetFollowPrizeDetailRequest) (*FollowPrizeGetFollowPrizeDetail, error)

	FollowPrizeGetFollowPrizeList(*FollowPrizeGetFollowPrizeListRequest) (*FollowPrizeGetFollowPrizeList, error)

	TopPicksGetTopPicksList(*TopPicksGetTopPicksListRequest) (*TopPicksGetTopPicksList, error)

	TopPicksAddTopPicks(*TopPicksAddTopPicksRequest) (*TopPicksAddTopPicks, error)

	TopPicksUpdateTopPicks(*TopPicksUpdateTopPicksRequest) (*TopPicksUpdateTopPicks, error)

	TopPicksDeleteTopPicks(*TopPicksDeleteTopPicksRequest) (*TopPicksDeleteTopPicks, error)

	ShopCategoryAddShopCategory(*ShopCategoryAddShopCategoryRequest) (*ShopCategoryAddShopCategory, error)

	ShopCategoryGetShopCategoryList(*ShopCategoryGetShopCategoryListRequest) (*ShopCategoryGetShopCategoryList, error)

	ShopCategoryDeleteShopCategory(*ShopCategoryDeleteShopCategoryRequest) (*ShopCategoryDeleteShopCategory, error)

	ShopCategoryUpdateShopCategory(*ShopCategoryUpdateShopCategoryRequest) (*ShopCategoryUpdateShopCategory, error)

	ShopCategoryAddItemList(*ShopCategoryAddItemListRequest) (*ShopCategoryAddItemList, error)

	ShopCategoryGetItemList(*ShopCategoryGetItemListRequest) (*ShopCategoryGetItemList, error)

	ShopCategoryDeleteItemList(*ShopCategoryDeleteItemListRequest) (*ShopCategoryDeleteItemList, error)

	ReturnsGetReturnDetail(*ReturnsGetReturnDetailRequest) (*ReturnsGetReturnDetail, error)

	ReturnsGetReturnList(*ReturnsGetReturnListRequest) ([]ReturnsGetReturnList, error)

	ReturnsConfirm(*ReturnsConfirmRequest) (*ReturnsConfirm, error)

	ReturnsDispute(*ReturnsDisputeRequest) (*ReturnsDispute, error)

	ReturnsGetAvailableSolutions(*ReturnsGetAvailableSolutionsRequest) (*ReturnsGetAvailableSolutions, error)

	ReturnsOffer(*ReturnsOfferRequest) (*ReturnsOffer, error)

	ReturnsAcceptOffer(*ReturnsAcceptOfferRequest) (*ReturnsAcceptOffer, error)

	AccountHealthShopPerformance(*AccountHealthShopPerformanceRequest) error

	AccountHealthShopPenalty(*AccountHealthShopPenaltyRequest) error

	PublicGetShopsByPartner(*PublicGetShopsByPartnerRequest) error

	PublicGetMerchantsByPartner(*PublicGetMerchantsByPartnerRequest) error

	PublicGetTokenByResendCode(*PublicGetTokenByResendCodeRequest) error

	PublicGetRefreshTokenByUpgradeCode(*PublicGetRefreshTokenByUpgradeCodeRequest) (*PublicGetRefreshTokenByUpgradeCode, error)

	PushGetPushConfig(*PushGetPushConfigRequest) error

	PushSetPushConfig(*PushSetPushConfigRequest) error
}
