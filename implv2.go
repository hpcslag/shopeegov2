package shopeego

func (s *ShopeeClient) ProductGetCategory(req *ProductGetCategoryRequest) (resp *ProductGetCategoryResponse, err error) {
	b, err := s.post("ProductGetCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetAttributes(req *ProductGetAttributesRequest) (resp *ProductGetAttributesResponse, err error) {
	b, err := s.post("ProductGetAttributes", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetBrandList(req *ProductGetBrandListRequest) (resp *ProductGetBrandListResponse, err error) {
	b, err := s.post("ProductGetBrandList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetDtsLimit(req *ProductGetDtsLimitRequest) (resp *ProductGetDtsLimitResponse, err error) {
	b, err := s.post("ProductGetDtsLimit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetItemLimit(req *ProductGetItemLimitRequest) (resp *ProductGetItemLimitResponse, err error) {
	b, err := s.post("ProductGetItemLimit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetItemList(req *ProductGetItemListRequest) (resp *ProductGetItemListResponse, err error) {
	b, err := s.post("ProductGetItemList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetItemBaseInfo(req *ProductGetItemBaseInfoRequest) (resp *ProductGetItemBaseInfoResponse, err error) {
	b, err := s.post("ProductGetItemBaseInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetItemExtraInfo(req *ProductGetItemExtraInfoRequest) (resp *ProductGetItemExtraInfoResponse, err error) {
	b, err := s.post("ProductGetItemExtraInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductAddItem(req *ProductAddItemRequest) (resp *ProductAddItemResponse, err error) {
	b, err := s.post("ProductAddItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateItem(req *ProductUpdateItemRequest) (resp *ProductUpdateItemResponse, err error) {
	b, err := s.post("ProductUpdateItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductDeleteItem(req *ProductDeleteItemRequest) (resp *ProductDeleteItemResponse, err error) {
	b, err := s.post("ProductDeleteItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductInitTierVariation(req *ProductInitTierVariationRequest) (resp *ProductInitTierVariationResponse, err error) {
	b, err := s.post("ProductInitTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateTierVariation(req *ProductUpdateTierVariationRequest) (resp *ProductUpdateTierVariationResponse, err error) {
	b, err := s.post("ProductUpdateTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetModelList(req *ProductGetModelListRequest) (resp *ProductGetModelListResponse, err error) {
	b, err := s.post("ProductGetModelList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductAddModel(req *ProductAddModelRequest) (resp *ProductAddModelResponse, err error) {
	b, err := s.post("ProductAddModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateModel(req *ProductUpdateModelRequest) (resp *ProductUpdateModelResponse, err error) {
	b, err := s.post("ProductUpdateModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductDeleteModel(req *ProductDeleteModelRequest) (resp *ProductDeleteModelResponse, err error) {
	b, err := s.post("ProductDeleteModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductSupportSizeChart(req *ProductSupportSizeChartRequest) (resp *ProductSupportSizeChartResponse, err error) {
	b, err := s.post("ProductSupportSizeChart", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateSizeChart(req *ProductUpdateSizeChartRequest) (resp *ProductUpdateSizeChartResponse, err error) {
	b, err := s.post("ProductUpdateSizeChart", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUnlistItem(req *ProductUnlistItemRequest) (resp *ProductUnlistItemResponse, err error) {
	b, err := s.post("ProductUnlistItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdatePrice(req *ProductUpdatePriceRequest) (resp *ProductUpdatePriceResponse, err error) {
	b, err := s.post("ProductUpdatePrice", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateStock(req *ProductUpdateStockRequest) (resp *ProductUpdateStockResponse, err error) {
	b, err := s.post("ProductUpdateStock", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductBoostItem(req *ProductBoostItemRequest) (resp *ProductBoostItemResponse, err error) {
	b, err := s.post("ProductBoostItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetBoostedList(req *ProductGetBoostedListRequest) (resp *ProductGetBoostedListResponse, err error) {
	b, err := s.post("ProductGetBoostedList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetItemPromotion(req *ProductGetItemPromotionRequest) (resp *ProductGetItemPromotionResponse, err error) {
	b, err := s.post("ProductGetItemPromotion", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductUpdateSipItemPrice(req *ProductUpdateSipItemPriceRequest) (resp *ProductUpdateSipItemPriceResponse, err error) {
	b, err := s.post("ProductUpdateSipItemPrice", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductSearchItem(req *ProductSearchItemRequest) (resp *ProductSearchItemResponse, err error) {
	b, err := s.post("ProductSearchItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetComment(req *ProductGetCommentRequest) (resp *ProductGetCommentResponse, err error) {
	b, err := s.post("ProductGetComment", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductReplyComment(req *ProductReplyCommentRequest) (resp *ProductReplyCommentResponse, err error) {
	b, err := s.post("ProductReplyComment", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductCategoryRecommend(req *ProductCategoryRecommendRequest) (resp *ProductCategoryRecommendResponse, err error) {
	b, err := s.post("ProductCategoryRecommend", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductRegisterBrand(req *ProductRegisterBrandRequest) (resp *ProductRegisterBrandResponse, err error) {
	b, err := s.post("ProductRegisterBrand", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ProductGetRecommendAttribute(req *ProductGetRecommendAttributeRequest) (resp *ProductGetRecommendAttributeResponse, err error) {
	b, err := s.post("ProductGetRecommendAttribute", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetCategory(req *GlobalProductGetCategoryRequest) (resp *GlobalProductGetCategoryResponse, err error) {
	b, err := s.post("GlobalProductGetCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetAttributes(req *GlobalProductGetAttributesRequest) (resp *GlobalProductGetAttributesResponse, err error) {
	b, err := s.post("GlobalProductGetAttributes", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetBrandList(req *GlobalProductGetBrandListRequest) (resp *GlobalProductGetBrandListResponse, err error) {
	b, err := s.post("GlobalProductGetBrandList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetGlobalItemLimit(req *GlobalProductGetGlobalItemLimitRequest) (resp *GlobalProductGetGlobalItemLimitResponse, err error) {
	b, err := s.post("GlobalProductGetGlobalItemLimit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetDtsLimit(req *GlobalProductGetDtsLimitRequest) (resp *GlobalProductGetDtsLimitResponse, err error) {
	b, err := s.post("GlobalProductGetDtsLimit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetGlobalItemList(req *GlobalProductGetGlobalItemListRequest) (resp *GlobalProductGetGlobalItemListResponse, err error) {
	b, err := s.post("GlobalProductGetGlobalItemList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetGlobalItemInfo(req *GlobalProductGetGlobalItemInfoRequest) (resp *GlobalProductGetGlobalItemInfoResponse, err error) {
	b, err := s.post("GlobalProductGetGlobalItemInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductAddGlobalItem(req *GlobalProductAddGlobalItemRequest) (resp *GlobalProductAddGlobalItemResponse, err error) {
	b, err := s.post("GlobalProductAddGlobalItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdateGlobalItem(req *GlobalProductUpdateGlobalItemRequest) (resp *GlobalProductUpdateGlobalItemResponse, err error) {
	b, err := s.post("GlobalProductUpdateGlobalItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductDeleteGlobalItem(req *GlobalProductDeleteGlobalItemRequest) (resp *GlobalProductDeleteGlobalItemResponse, err error) {
	b, err := s.post("GlobalProductDeleteGlobalItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductInitTierVariation(req *GlobalProductInitTierVariationRequest) (resp *GlobalProductInitTierVariationResponse, err error) {
	b, err := s.post("GlobalProductInitTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdateTierVariation(req *GlobalProductUpdateTierVariationRequest) (resp *GlobalProductUpdateTierVariationResponse, err error) {
	b, err := s.post("GlobalProductUpdateTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductAddGlobalModel(req *GlobalProductAddGlobalModelRequest) (resp *GlobalProductAddGlobalModelResponse, err error) {
	b, err := s.post("GlobalProductAddGlobalModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdateGlobalModel(req *GlobalProductUpdateGlobalModelRequest) (resp *GlobalProductUpdateGlobalModelResponse, err error) {
	b, err := s.post("GlobalProductUpdateGlobalModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductDeleteGlobalModel(req *GlobalProductDeleteGlobalModelRequest) (resp *GlobalProductDeleteGlobalModelResponse, err error) {
	b, err := s.post("GlobalProductDeleteGlobalModel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetGlobalModelList(req *GlobalProductGetGlobalModelListRequest) (resp *GlobalProductGetGlobalModelListResponse, err error) {
	b, err := s.post("GlobalProductGetGlobalModelList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductSupportSizeChart(req *GlobalProductSupportSizeChartRequest) (resp *GlobalProductSupportSizeChartResponse, err error) {
	b, err := s.post("GlobalProductSupportSizeChart", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdateSizeChart(req *GlobalProductUpdateSizeChartRequest) (resp *GlobalProductUpdateSizeChartResponse, err error) {
	b, err := s.post("GlobalProductUpdateSizeChart", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductCreatePublishTask(req *GlobalProductCreatePublishTaskRequest) (resp *GlobalProductCreatePublishTaskResponse, err error) {
	b, err := s.post("GlobalProductCreatePublishTask", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetPublishableShop(req *GlobalProductGetPublishableShopRequest) (resp *GlobalProductGetPublishableShopResponse, err error) {
	b, err := s.post("GlobalProductGetPublishableShop", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetPublishTaskResult(req *GlobalProductGetPublishTaskResultRequest) (resp *GlobalProductGetPublishTaskResultResponse, err error) {
	b, err := s.post("GlobalProductGetPublishTaskResult", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetPublishedList(req *GlobalProductGetPublishedListRequest) (resp *GlobalProductGetPublishedListResponse, err error) {
	b, err := s.post("GlobalProductGetPublishedList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdatePrice(req *GlobalProductUpdatePriceRequest) (resp *GlobalProductUpdatePriceResponse, err error) {
	b, err := s.post("GlobalProductUpdatePrice", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductUpdateStock(req *GlobalProductUpdateStockRequest) (resp *GlobalProductUpdateStockResponse, err error) {
	b, err := s.post("GlobalProductUpdateStock", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductSetSyncField(req *GlobalProductSetSyncFieldRequest) (resp *GlobalProductSetSyncFieldResponse, err error) {
	b, err := s.post("GlobalProductSetSyncField", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetGlobalItemID(req *GlobalProductGetGlobalItemIDRequest) (resp *GlobalProductGetGlobalItemIDResponse, err error) {
	b, err := s.post("GlobalProductGetGlobalItemID", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductCategoryRecommend(req *GlobalProductCategoryRecommendRequest) (resp *GlobalProductCategoryRecommendResponse, err error) {
	b, err := s.post("GlobalProductCategoryRecommend", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) GlobalProductGetRecommendAttribute(req *GlobalProductGetRecommendAttributeRequest) (resp *GlobalProductGetRecommendAttributeResponse, err error) {
	b, err := s.post("GlobalProductGetRecommendAttribute", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceInitVideoUpload(req *MediaSpaceInitVideoUploadRequest) (resp *MediaSpaceInitVideoUploadResponse, err error) {
	b, err := s.post("MediaSpaceInitVideoUpload", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceUploadVideoPart(req *MediaSpaceUploadVideoPartRequest) (resp *MediaSpaceUploadVideoPartResponse, err error) {
	b, err := s.post("MediaSpaceUploadVideoPart", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceCompleteVideoUpload(req *MediaSpaceCompleteVideoUploadRequest) (resp *MediaSpaceCompleteVideoUploadResponse, err error) {
	b, err := s.post("MediaSpaceCompleteVideoUpload", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceGetVideoUploadResult(req *MediaSpaceGetVideoUploadResultRequest) (resp *MediaSpaceGetVideoUploadResultResponse, err error) {
	b, err := s.post("MediaSpaceGetVideoUploadResult", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceCancelVideoUpload(req *MediaSpaceCancelVideoUploadRequest) (resp *MediaSpaceCancelVideoUploadResponse, err error) {
	b, err := s.post("MediaSpaceCancelVideoUpload", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MediaSpaceUploadImage(req *MediaSpaceUploadImageRequest) (resp *MediaSpaceUploadImageResponse, err error) {
	b, err := s.post("MediaSpaceUploadImage", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopGetShopInfo(req *ShopGetShopInfoRequest) (resp *ShopGetShopInfoResponse, err error) {
	b, err := s.post("ShopGetShopInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopGetProfile(req *ShopGetProfileRequest) (resp *ShopGetProfileResponse, err error) {
	b, err := s.post("ShopGetProfile", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopUpdateProfile(req *ShopUpdateProfileRequest) (resp *ShopUpdateProfileResponse, err error) {
	b, err := s.post("ShopUpdateProfile", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopGetWarehouseDetail(req *ShopGetWarehouseDetailRequest) (resp *ShopGetWarehouseDetailResponse, err error) {
	b, err := s.post("ShopGetWarehouseDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MerchantGetMerchantInfo(req *MerchantGetMerchantInfoRequest) (resp *MerchantGetMerchantInfoResponse, err error) {
	b, err := s.post("MerchantGetMerchantInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) MerchantGetShopListByMerchant(req *MerchantGetShopListByMerchantRequest) (resp *MerchantGetShopListByMerchantResponse, err error) {
	b, err := s.post("MerchantGetShopListByMerchant", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderGetOrderList(req *OrderGetOrderListRequest) (resp *OrderGetOrderListResponse, err error) {
	b, err := s.post("OrderGetOrderList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderGetShipmentList(req *OrderGetShipmentListRequest) (resp *OrderGetShipmentListResponse, err error) {
	b, err := s.post("OrderGetShipmentList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderGetOrderDetail(req *OrderGetOrderDetailRequest) (resp *OrderGetOrderDetailResponse, err error) {
	b, err := s.post("OrderGetOrderDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderSplitOrder(req *OrderSplitOrderRequest) (resp *OrderSplitOrderResponse, err error) {
	b, err := s.post("OrderSplitOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderUnsplitOrder(req *OrderUnsplitOrderRequest) (resp *OrderUnsplitOrderResponse, err error) {
	b, err := s.post("OrderUnsplitOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderCancelOrder(req *OrderCancelOrderRequest) (resp *OrderCancelOrderResponse, err error) {
	b, err := s.post("OrderCancelOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderHandleBuyerCancellation(req *OrderHandleBuyerCancellationRequest) (resp *OrderHandleBuyerCancellationResponse, err error) {
	b, err := s.post("OrderHandleBuyerCancellation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderSetNote(req *OrderSetNoteRequest) (resp *OrderSetNoteResponse, err error) {
	b, err := s.post("OrderSetNote", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderAddInvoiceData(req *OrderAddInvoiceDataRequest) (resp *OrderAddInvoiceDataResponse, err error) {
	b, err := s.post("OrderAddInvoiceData", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderGetPendingBuyerInvoiceOrderList(req *OrderGetPendingBuyerInvoiceOrderListRequest) (resp *OrderGetPendingBuyerInvoiceOrderListResponse, err error) {
	b, err := s.post("OrderGetPendingBuyerInvoiceOrderList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderUploadInvoiceDoc(req *OrderUploadInvoiceDocRequest) (resp *OrderUploadInvoiceDocResponse, err error) {
	b, err := s.post("OrderUploadInvoiceDoc", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderDownloadInvoiceDoc(req *OrderDownloadInvoiceDocRequest) (resp *OrderDownloadInvoiceDocResponse, err error) {
	b, err := s.post("OrderDownloadInvoiceDoc", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) OrderGetBuyerInvoiceInfo(req *OrderGetBuyerInvoiceInfoRequest) (resp *OrderGetBuyerInvoiceInfoResponse, err error) {
	b, err := s.post("OrderGetBuyerInvoiceInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetShippingParameter(req *LogisticsGetShippingParameterRequest) (resp *LogisticsGetShippingParameterResponse, err error) {
	b, err := s.post("LogisticsGetShippingParameter", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetTrackingNumber(req *LogisticsGetTrackingNumberRequest) (resp *LogisticsGetTrackingNumberResponse, err error) {
	b, err := s.post("LogisticsGetTrackingNumber", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsShipOrder(req *LogisticsShipOrderRequest) (resp *LogisticsShipOrderResponse, err error) {
	b, err := s.post("LogisticsShipOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsUpdateShippingOrder(req *LogisticsUpdateShippingOrderRequest) (resp *LogisticsUpdateShippingOrderResponse, err error) {
	b, err := s.post("LogisticsUpdateShippingOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetShippingDocumentParameter(req *LogisticsGetShippingDocumentParameterRequest) (resp *LogisticsGetShippingDocumentParameterResponse, err error) {
	b, err := s.post("LogisticsGetShippingDocumentParameter", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsCreateShippingDocument(req *LogisticsCreateShippingDocumentRequest) (resp *LogisticsCreateShippingDocumentResponse, err error) {
	b, err := s.post("LogisticsCreateShippingDocument", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetShippingDocumentResult(req *LogisticsGetShippingDocumentResultRequest) (resp *LogisticsGetShippingDocumentResultResponse, err error) {
	b, err := s.post("LogisticsGetShippingDocumentResult", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsDownloadShippingDocument(req *LogisticsDownloadShippingDocumentRequest) (resp *LogisticsDownloadShippingDocumentResponse, err error) {
	b, err := s.post("LogisticsDownloadShippingDocument", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetShippingDocumentInfo(req *LogisticsGetShippingDocumentInfoRequest) (resp *LogisticsGetShippingDocumentInfoResponse, err error) {
	b, err := s.post("LogisticsGetShippingDocumentInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetTrackingInfo(req *LogisticsGetTrackingInfoRequest) (resp *LogisticsGetTrackingInfoResponse, err error) {
	b, err := s.post("LogisticsGetTrackingInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetAddressList(req *LogisticsGetAddressListRequest) (resp *LogisticsGetAddressListResponse, err error) {
	b, err := s.post("LogisticsGetAddressList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsSetAddressConfig(req *LogisticsSetAddressConfigRequest) (resp *LogisticsSetAddressConfigResponse, err error) {
	b, err := s.post("LogisticsSetAddressConfig", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsDeleteAddress(req *LogisticsDeleteAddressRequest) (resp *LogisticsDeleteAddressResponse, err error) {
	b, err := s.post("LogisticsDeleteAddress", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsGetChannelList(req *LogisticsGetChannelListRequest) (resp *LogisticsGetChannelListResponse, err error) {
	b, err := s.post("LogisticsGetChannelList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsUpdateChannel(req *LogisticsUpdateChannelRequest) (resp *LogisticsUpdateChannelResponse, err error) {
	b, err := s.post("LogisticsUpdateChannel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) LogisticsBatchShipOrder(req *LogisticsBatchShipOrderRequest) (resp *LogisticsBatchShipOrderResponse, err error) {
	b, err := s.post("LogisticsBatchShipOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGetUnbindOrderList(req *FirstMileGetUnbindOrderListRequest) (resp *FirstMileGetUnbindOrderListResponse, err error) {
	b, err := s.post("FirstMileGetUnbindOrderList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGetDetail(req *FirstMileGetDetailRequest) (resp *FirstMileGetDetailResponse, err error) {
	b, err := s.post("FirstMileGetDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGenerateFirstMileTrackingNumber(req *FirstMileGenerateFirstMileTrackingNumberRequest) (resp *FirstMileGenerateFirstMileTrackingNumberResponse, err error) {
	b, err := s.post("FirstMileGenerateFirstMileTrackingNumber", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileBindFirstMileTrackingNumber(req *FirstMileBindFirstMileTrackingNumberRequest) (resp *FirstMileBindFirstMileTrackingNumberResponse, err error) {
	b, err := s.post("FirstMileBindFirstMileTrackingNumber", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileUnbindFirstMileTrackingNumber(req *FirstMileUnbindFirstMileTrackingNumberRequest) (resp *FirstMileUnbindFirstMileTrackingNumberResponse, err error) {
	b, err := s.post("FirstMileUnbindFirstMileTrackingNumber", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGetTrackingNumberList(req *FirstMileGetTrackingNumberListRequest) (resp *FirstMileGetTrackingNumberListResponse, err error) {
	b, err := s.post("FirstMileGetTrackingNumberList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGetWaybill(req *FirstMileGetWaybillRequest) (resp *FirstMileGetWaybillResponse, err error) {
	b, err := s.post("FirstMileGetWaybill", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FirstMileGetChannelList(req *FirstMileGetChannelListRequest) (resp *FirstMileGetChannelListResponse, err error) {
	b, err := s.post("FirstMileGetChannelList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetEscrowDetail(req *PaymentGetEscrowDetailRequest) (resp *PaymentGetEscrowDetailResponse, err error) {
	b, err := s.post("PaymentGetEscrowDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentSetShopInstallmentStatus(req *PaymentSetShopInstallmentStatusRequest) (resp *PaymentSetShopInstallmentStatusResponse, err error) {
	b, err := s.post("PaymentSetShopInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetShopInstallmentStatus(req *PaymentGetShopInstallmentStatusRequest) (resp *PaymentGetShopInstallmentStatusResponse, err error) {
	b, err := s.post("PaymentGetShopInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetPayoutDetail(req *PaymentGetPayoutDetailRequest) (resp *PaymentGetPayoutDetailResponse, err error) {
	b, err := s.post("PaymentGetPayoutDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentSetItemInstallmentStatus(req *PaymentSetItemInstallmentStatusRequest) (resp *PaymentSetItemInstallmentStatusResponse, err error) {
	b, err := s.post("PaymentSetItemInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetItemInstallmentStatus(req *PaymentGetItemInstallmentStatusRequest) (resp *PaymentGetItemInstallmentStatusResponse, err error) {
	b, err := s.post("PaymentGetItemInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetPaymentMethodList(req *PaymentGetPaymentMethodListRequest) (resp *PaymentGetPaymentMethodListResponse, err error) {
	b, err := s.post("PaymentGetPaymentMethodList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetWalletTransactionList(req *PaymentGetWalletTransactionListRequest) (resp *PaymentGetWalletTransactionListResponse, err error) {
	b, err := s.post("PaymentGetWalletTransactionList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PaymentGetEscrowList(req *PaymentGetEscrowListRequest) (resp *PaymentGetEscrowListResponse, err error) {
	b, err := s.post("PaymentGetEscrowList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountAddDiscount(req *DiscountAddDiscountRequest) (resp *DiscountAddDiscountResponse, err error) {
	b, err := s.post("DiscountAddDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountAddDiscountItem(req *DiscountAddDiscountItemRequest) (resp *DiscountAddDiscountItemResponse, err error) {
	b, err := s.post("DiscountAddDiscountItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountDeleteDiscount(req *DiscountDeleteDiscountRequest) (resp *DiscountDeleteDiscountResponse, err error) {
	b, err := s.post("DiscountDeleteDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountDeleteDiscountItem(req *DiscountDeleteDiscountItemRequest) (resp *DiscountDeleteDiscountItemResponse, err error) {
	b, err := s.post("DiscountDeleteDiscountItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountGetDiscount(req *DiscountGetDiscountRequest) (resp *DiscountGetDiscountResponse, err error) {
	b, err := s.post("DiscountGetDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountGetDiscountList(req *DiscountGetDiscountListRequest) (resp *DiscountGetDiscountListResponse, err error) {
	b, err := s.post("DiscountGetDiscountList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountUpdateDiscount(req *DiscountUpdateDiscountRequest) (resp *DiscountUpdateDiscountResponse, err error) {
	b, err := s.post("DiscountUpdateDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountUpdateDiscountItem(req *DiscountUpdateDiscountItemRequest) (resp *DiscountUpdateDiscountItemResponse, err error) {
	b, err := s.post("DiscountUpdateDiscountItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) DiscountEndDiscount(req *DiscountEndDiscountRequest) (resp *DiscountEndDiscountResponse, err error) {
	b, err := s.post("DiscountEndDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealAddBundleDeal(req *BundleDealAddBundleDealRequest) (resp *BundleDealAddBundleDealResponse, err error) {
	b, err := s.post("BundleDealAddBundleDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealAddBundleDealItem(req *BundleDealAddBundleDealItemRequest) (resp *BundleDealAddBundleDealItemResponse, err error) {
	b, err := s.post("BundleDealAddBundleDealItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealGetBundleDealList(req *BundleDealGetBundleDealListRequest) (resp *BundleDealGetBundleDealListResponse, err error) {
	b, err := s.post("BundleDealGetBundleDealList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealGetBundleDeal(req *BundleDealGetBundleDealRequest) (resp *BundleDealGetBundleDealResponse, err error) {
	b, err := s.post("BundleDealGetBundleDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealGetBundleDealItem(req *BundleDealGetBundleDealItemRequest) (resp *BundleDealGetBundleDealItemResponse, err error) {
	b, err := s.post("BundleDealGetBundleDealItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealUpdateBundleDeal(req *BundleDealUpdateBundleDealRequest) (resp *BundleDealUpdateBundleDealResponse, err error) {
	b, err := s.post("BundleDealUpdateBundleDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealUpdateBundleDealItem(req *BundleDealUpdateBundleDealItemRequest) (resp *BundleDealUpdateBundleDealItemResponse, err error) {
	b, err := s.post("BundleDealUpdateBundleDealItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealEndBundleDeal(req *BundleDealEndBundleDealRequest) (resp *BundleDealEndBundleDealResponse, err error) {
	b, err := s.post("BundleDealEndBundleDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealDeleteBundleDeal(req *BundleDealDeleteBundleDealRequest) (resp *BundleDealDeleteBundleDealResponse, err error) {
	b, err := s.post("BundleDealDeleteBundleDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) BundleDealDeleteBundleDealItem(req *BundleDealDeleteBundleDealItemRequest) (resp *BundleDealDeleteBundleDealItemResponse, err error) {
	b, err := s.post("BundleDealDeleteBundleDealItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealAddAddOnDeal(req *AddOnDealAddAddOnDealRequest) (resp *AddOnDealAddAddOnDealResponse, err error) {
	b, err := s.post("AddOnDealAddAddOnDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealAddAddOnDealMainItem(req *AddOnDealAddAddOnDealMainItemRequest) (resp *AddOnDealAddAddOnDealMainItemResponse, err error) {
	b, err := s.post("AddOnDealAddAddOnDealMainItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealAddAddOnDealSubItem(req *AddOnDealAddAddOnDealSubItemRequest) (resp *AddOnDealAddAddOnDealSubItemResponse, err error) {
	b, err := s.post("AddOnDealAddAddOnDealSubItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealDeleteAddOnDeal(req *AddOnDealDeleteAddOnDealRequest) (resp *AddOnDealDeleteAddOnDealResponse, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealDeleteAddOnDealMainItem(req *AddOnDealDeleteAddOnDealMainItemRequest) (resp *AddOnDealDeleteAddOnDealMainItemResponse, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDealMainItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealDeleteAddOnDealSubItem(req *AddOnDealDeleteAddOnDealSubItemRequest) (resp *AddOnDealDeleteAddOnDealSubItemResponse, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDealSubItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealGetAddOnDealList(req *AddOnDealGetAddOnDealListRequest) (resp *AddOnDealGetAddOnDealListResponse, err error) {
	b, err := s.post("AddOnDealGetAddOnDealList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealGetAddOnDeal(req *AddOnDealGetAddOnDealRequest) (resp *AddOnDealGetAddOnDealResponse, err error) {
	b, err := s.post("AddOnDealGetAddOnDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealGetAddOnDealMainItem(req *AddOnDealGetAddOnDealMainItemRequest) (resp *AddOnDealGetAddOnDealMainItemResponse, err error) {
	b, err := s.post("AddOnDealGetAddOnDealMainItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealGetAddOnDealSubItem(req *AddOnDealGetAddOnDealSubItemRequest) (resp *AddOnDealGetAddOnDealSubItemResponse, err error) {
	b, err := s.post("AddOnDealGetAddOnDealSubItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealUpdateAddOnDeal(req *AddOnDealUpdateAddOnDealRequest) (resp *AddOnDealUpdateAddOnDealResponse, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealUpdateAddOnDealMainItem(req *AddOnDealUpdateAddOnDealMainItemRequest) (resp *AddOnDealUpdateAddOnDealMainItemResponse, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDealMainItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealUpdateAddOnDealSubItem(req *AddOnDealUpdateAddOnDealSubItemRequest) (resp *AddOnDealUpdateAddOnDealSubItemResponse, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDealSubItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AddOnDealEndAddOnDeal(req *AddOnDealEndAddOnDealRequest) (resp *AddOnDealEndAddOnDealResponse, err error) {
	b, err := s.post("AddOnDealEndAddOnDeal", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherAddVoucher(req *VoucherAddVoucherRequest) (resp *VoucherAddVoucherResponse, err error) {
	b, err := s.post("VoucherAddVoucher", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherDeleteVoucher(req *VoucherDeleteVoucherRequest) (resp *VoucherDeleteVoucherResponse, err error) {
	b, err := s.post("VoucherDeleteVoucher", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherEndVoucher(req *VoucherEndVoucherRequest) (resp *VoucherEndVoucherResponse, err error) {
	b, err := s.post("VoucherEndVoucher", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherUpdateVoucher(req *VoucherUpdateVoucherRequest) (resp *VoucherUpdateVoucherResponse, err error) {
	b, err := s.post("VoucherUpdateVoucher", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherGetVoucher(req *VoucherGetVoucherRequest) (resp *VoucherGetVoucherResponse, err error) {
	b, err := s.post("VoucherGetVoucher", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) VoucherGetVoucherList(req *VoucherGetVoucherListRequest) (resp *VoucherGetVoucherListResponse, err error) {
	b, err := s.post("VoucherGetVoucherList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeAddFollowPrize(req *FollowPrizeAddFollowPrizeRequest) (resp *FollowPrizeAddFollowPrizeResponse, err error) {
	b, err := s.post("FollowPrizeAddFollowPrize", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeDeleteFollowPrize(req *FollowPrizeDeleteFollowPrizeRequest) (resp *FollowPrizeDeleteFollowPrizeResponse, err error) {
	b, err := s.post("FollowPrizeDeleteFollowPrize", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeEndFollowPrize(req *FollowPrizeEndFollowPrizeRequest) (resp *FollowPrizeEndFollowPrizeResponse, err error) {
	b, err := s.post("FollowPrizeEndFollowPrize", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeUpdateFollowPrize(req *FollowPrizeUpdateFollowPrizeRequest) (resp *FollowPrizeUpdateFollowPrizeResponse, err error) {
	b, err := s.post("FollowPrizeUpdateFollowPrize", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeGetFollowPrizeDetail(req *FollowPrizeGetFollowPrizeDetailRequest) (resp *FollowPrizeGetFollowPrizeDetailResponse, err error) {
	b, err := s.post("FollowPrizeGetFollowPrizeDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) FollowPrizeGetFollowPrizeList(req *FollowPrizeGetFollowPrizeListRequest) (resp *FollowPrizeGetFollowPrizeListResponse, err error) {
	b, err := s.post("FollowPrizeGetFollowPrizeList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) TopPicksGetTopPicksList(req *TopPicksGetTopPicksListRequest) (resp *TopPicksGetTopPicksListResponse, err error) {
	b, err := s.post("TopPicksGetTopPicksList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) TopPicksAddTopPicks(req *TopPicksAddTopPicksRequest) (resp *TopPicksAddTopPicksResponse, err error) {
	b, err := s.post("TopPicksAddTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) TopPicksUpdateTopPicks(req *TopPicksUpdateTopPicksRequest) (resp *TopPicksUpdateTopPicksResponse, err error) {
	b, err := s.post("TopPicksUpdateTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) TopPicksDeleteTopPicks(req *TopPicksDeleteTopPicksRequest) (resp *TopPicksDeleteTopPicksResponse, err error) {
	b, err := s.post("TopPicksDeleteTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryAddShopCategory(req *ShopCategoryAddShopCategoryRequest) (resp *ShopCategoryAddShopCategoryResponse, err error) {
	b, err := s.post("ShopCategoryAddShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryGetShopCategoryList(req *ShopCategoryGetShopCategoryListRequest) (resp *ShopCategoryGetShopCategoryListResponse, err error) {
	b, err := s.post("ShopCategoryGetShopCategoryList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryDeleteShopCategory(req *ShopCategoryDeleteShopCategoryRequest) (resp *ShopCategoryDeleteShopCategoryResponse, err error) {
	b, err := s.post("ShopCategoryDeleteShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryUpdateShopCategory(req *ShopCategoryUpdateShopCategoryRequest) (resp *ShopCategoryUpdateShopCategoryResponse, err error) {
	b, err := s.post("ShopCategoryUpdateShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryAddItemList(req *ShopCategoryAddItemListRequest) (resp *ShopCategoryAddItemListResponse, err error) {
	b, err := s.post("ShopCategoryAddItemList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryGetItemList(req *ShopCategoryGetItemListRequest) (resp *ShopCategoryGetItemListResponse, err error) {
	b, err := s.post("ShopCategoryGetItemList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ShopCategoryDeleteItemList(req *ShopCategoryDeleteItemListRequest) (resp *ShopCategoryDeleteItemListResponse, err error) {
	b, err := s.post("ShopCategoryDeleteItemList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsGetReturnDetail(req *ReturnsGetReturnDetailRequest) (resp *ReturnsGetReturnDetailResponse, err error) {
	b, err := s.post("ReturnsGetReturnDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsGetReturnList(req *ReturnsGetReturnListRequest) (resp *ReturnsGetReturnListResponse, err error) {
	b, err := s.post("ReturnsGetReturnList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsConfirm(req *ReturnsConfirmRequest) (resp *ReturnsConfirmResponse, err error) {
	b, err := s.post("ReturnsConfirm", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsDispute(req *ReturnsDisputeRequest) (resp *ReturnsDisputeResponse, err error) {
	b, err := s.post("ReturnsDispute", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsGetAvailableSolutions(req *ReturnsGetAvailableSolutionsRequest) (resp *ReturnsGetAvailableSolutionsResponse, err error) {
	b, err := s.post("ReturnsGetAvailableSolutions", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsOffer(req *ReturnsOfferRequest) (resp *ReturnsOfferResponse, err error) {
	b, err := s.post("ReturnsOffer", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) ReturnsAcceptOffer(req *ReturnsAcceptOfferRequest) (resp *ReturnsAcceptOfferResponse, err error) {
	b, err := s.post("ReturnsAcceptOffer", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AccountHealthShopPerformance(req *AccountHealthShopPerformanceRequest) (resp *AccountHealthShopPerformanceResponse, err error) {
	b, err := s.post("AccountHealthShopPerformance", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) AccountHealthShopPenalty(req *AccountHealthShopPenaltyRequest) (resp *AccountHealthShopPenaltyResponse, err error) {
	b, err := s.post("AccountHealthShopPenalty", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PublicGetShopsByPartner(req *PublicGetShopsByPartnerRequest) (resp *PublicGetShopsByPartnerResponse, err error) {
	b, err := s.post("PublicGetShopsByPartner", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PublicGetMerchantsByPartner(req *PublicGetMerchantsByPartnerRequest) (resp *PublicGetMerchantsByPartnerResponse, err error) {
	b, err := s.post("PublicGetMerchantsByPartner", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PublicGetTokenByResendCode(req *PublicGetTokenByResendCodeRequest) (resp *PublicGetTokenByResendCodeResponse, err error) {
	b, err := s.post("PublicGetTokenByResendCode", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PublicGetRefreshTokenByUpgradeCode(req *PublicGetRefreshTokenByUpgradeCodeRequest) (resp *PublicGetRefreshTokenByUpgradeCodeResponse, err error) {
	b, err := s.post("PublicGetRefreshTokenByUpgradeCode", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PushGetPushConfig(req *PushGetPushConfigRequest) (resp *PushGetPushConfigResponse, err error) {
	b, err := s.post("PushGetPushConfig", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

func (s *ShopeeClient) PushSetPushConfig(req *PushSetPushConfigRequest) (resp *PushSetPushConfigResponse, err error) {
	b, err := s.post("PushSetPushConfig", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}
