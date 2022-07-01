
package shopeego

// 統一的 API 回應介面
type V2UnityResponse struct {
	// error is Indicate error type if hit error. Empty if no error happened.
	Error string `json:"error,omitempty"`
	// message is Indicate error details if hit error. Empty if no error happened.
	Message string `json:"message,omitempty"`
	// request_id is The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// Warning message.
	Warning string `json:"warning,omitempty"`
}


func (s *ShopeeClient) ProductGetCategory(req *ProductGetCategoryRequest) (resp *ProductGetCategory, err error) {
	b, err := s.post("ProductGetCategory", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetAttributes(req *ProductGetAttributesRequest) (resp *ProductGetAttributes, err error) {
	b, err := s.post("ProductGetAttributes", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetAttributeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetBrandList(req *ProductGetBrandListRequest) (resp *ProductGetBrandList, err error) {
	b, err := s.post("ProductGetBrandList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetBrandResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetDtsLimit(req *ProductGetDtsLimitRequest) (resp *ProductGetDtsLimit, err error) {
	b, err := s.post("ProductGetDtsLimit", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetDtsLimitResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetItemLimit(req *ProductGetItemLimitRequest) (resp *ProductGetItemLimit, err error) {
	b, err := s.post("ProductGetItemLimit", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetItemLimitResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetItemList(req *ProductGetItemListRequest) (resp *ProductGetItemList, err error) {
	b, err := s.post("ProductGetItemList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetItemBaseInfo(req *ProductGetItemBaseInfoRequest) (resp *ProductGetItemBaseInfo, err error) {
	b, err := s.post("ProductGetItemBaseInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetItemBaseInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetItemExtraInfo(req *ProductGetItemExtraInfoRequest) (resp *ProductGetItemExtraInfo, err error) {
	b, err := s.post("ProductGetItemExtraInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetItemExtraInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductAddItem(req *ProductAddItemRequest) (resp *ProductAddItem, err error) {
	b, err := s.post("ProductAddItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductAddItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateItem(req *ProductUpdateItemRequest) (resp *ProductUpdateItem, err error) {
	b, err := s.post("ProductUpdateItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductDeleteItem(req *ProductDeleteItemRequest) (err error) {
	b, err := s.post("ProductDeleteItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductDeleteItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductInitTierVariation(req *ProductInitTierVariationRequest) (resp *ProductInitTierVariation, err error) {
	b, err := s.post("ProductInitTierVariation", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductInitTierVariationResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateTierVariation(req *ProductUpdateTierVariationRequest) (err error) {
	b, err := s.post("ProductUpdateTierVariation", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateTierVariationResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductGetModelList(req *ProductGetModelListRequest) (resp *ProductGetModelList, err error) {
	b, err := s.post("ProductGetModelList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductAddModel(req *ProductAddModelRequest) (resp *ProductAddModel, err error) {
	b, err := s.post("ProductAddModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductAddModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateModel(req *ProductUpdateModelRequest) (err error) {
	b, err := s.post("ProductUpdateModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductDeleteModel(req *ProductDeleteModelRequest) (err error) {
	b, err := s.post("ProductDeleteModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductDeleteModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductSupportSizeChart(req *ProductSupportSizeChartRequest) (resp *ProductSupportSizeChart, err error) {
	b, err := s.post("ProductSupportSizeChart", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductSupportSizeChartResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateSizeChart(req *ProductUpdateSizeChartRequest) (err error) {
	b, err := s.post("ProductUpdateSizeChart", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateSizeChartResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductUnlistItem(req *ProductUnlistItemRequest) (resp *ProductUnlistItem, err error) {
	b, err := s.post("ProductUnlistItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUnlistItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdatePrice(req *ProductUpdatePriceRequest) (resp *ProductUpdatePrice, err error) {
	b, err := s.post("ProductUpdatePrice", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdatePriceResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateStock(req *ProductUpdateStockRequest) (resp *ProductUpdateStock, err error) {
	b, err := s.post("ProductUpdateStock", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateStockResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductBoostItem(req *ProductBoostItemRequest) (resp *ProductBoostItem, err error) {
	b, err := s.post("ProductBoostItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductBoostItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetBoostedList(req *ProductGetBoostedListRequest) (resp *ProductGetBoostedList, err error) {
	b, err := s.post("ProductGetBoostedList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetBoostedResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetItemPromotion(req *ProductGetItemPromotionRequest) (resp *ProductGetItemPromotion, err error) {
	b, err := s.post("ProductGetItemPromotion", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetItemPromotionResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductUpdateSipItemPrice(req *ProductUpdateSipItemPriceRequest) (err error) {
	b, err := s.post("ProductUpdateSipItemPrice", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductUpdateSipItemPriceResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ProductSearchItem(req *ProductSearchItemRequest) (resp *ProductSearchItem, err error) {
	b, err := s.post("ProductSearchItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductSearchItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetComment(req *ProductGetCommentRequest) (resp *ProductGetComment, err error) {
	b, err := s.post("ProductGetComment", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetCommentResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductReplyComment(req *ProductReplyCommentRequest) (resp *ProductReplyComment, err error) {
	b, err := s.post("ProductReplyComment", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductReplyCommentResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductCategoryRecommend(req *ProductCategoryRecommendRequest) (resp *ProductCategoryRecommend, err error) {
	b, err := s.post("ProductCategoryRecommend", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductCategoryRecommendResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductRegisterBrand(req *ProductRegisterBrandRequest) (resp *ProductRegisterBrand, err error) {
	b, err := s.post("ProductRegisterBrand", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductRegisterBrandResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ProductGetRecommendAttribute(req *ProductGetRecommendAttributeRequest) (resp *ProductGetRecommendAttribute, err error) {
	b, err := s.post("ProductGetRecommendAttribute", req)
	if err != nil {
		return
	}

	var wrappedResponse *ProductGetRecommendAttributeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetCategory(req *GlobalProductGetCategoryRequest) (resp *GlobalProductGetCategory, err error) {
	b, err := s.post("GlobalProductGetCategory", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetAttributes(req *GlobalProductGetAttributesRequest) (resp *GlobalProductGetAttributes, err error) {
	b, err := s.post("GlobalProductGetAttributes", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetAttributeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetBrandList(req *GlobalProductGetBrandListRequest) (resp *GlobalProductGetBrandList, err error) {
	b, err := s.post("GlobalProductGetBrandList", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetBrandResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetGlobalItemLimit(req *GlobalProductGetGlobalItemLimitRequest) (resp *GlobalProductGetGlobalItemLimit, err error) {
	b, err := s.post("GlobalProductGetGlobalItemLimit", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetGlobalItemLimitResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetDtsLimit(req *GlobalProductGetDtsLimitRequest) (resp *GlobalProductGetDtsLimit, err error) {
	b, err := s.post("GlobalProductGetDtsLimit", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetDtsLimitResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetGlobalItemList(req *GlobalProductGetGlobalItemListRequest) (resp *GlobalProductGetGlobalItemList, err error) {
	b, err := s.post("GlobalProductGetGlobalItemList", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetGlobalItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetGlobalItemInfo(req *GlobalProductGetGlobalItemInfoRequest) (resp *GlobalProductGetGlobalItemInfo, err error) {
	b, err := s.post("GlobalProductGetGlobalItemInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetGlobalItemInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductAddGlobalItem(req *GlobalProductAddGlobalItemRequest) (resp *GlobalProductAddGlobalItem, err error) {
	b, err := s.post("GlobalProductAddGlobalItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductAddGlobalItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductUpdateGlobalItem(req *GlobalProductUpdateGlobalItemRequest) (resp *GlobalProductUpdateGlobalItem, err error) {
	b, err := s.post("GlobalProductUpdateGlobalItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdateGlobalItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductDeleteGlobalItem(req *GlobalProductDeleteGlobalItemRequest) (resp *GlobalProductDeleteGlobalItem, err error) {
	b, err := s.post("GlobalProductDeleteGlobalItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductDeleteGlobalItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductInitTierVariation(req *GlobalProductInitTierVariationRequest) (err error) {
	b, err := s.post("GlobalProductInitTierVariation", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductInitTierVariationResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductUpdateTierVariation(req *GlobalProductUpdateTierVariationRequest) (err error) {
	b, err := s.post("GlobalProductUpdateTierVariation", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdateTierVariationResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductAddGlobalModel(req *GlobalProductAddGlobalModelRequest) (err error) {
	b, err := s.post("GlobalProductAddGlobalModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductAddGlobalModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductUpdateGlobalModel(req *GlobalProductUpdateGlobalModelRequest) (err error) {
	b, err := s.post("GlobalProductUpdateGlobalModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdateGlobalModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductDeleteGlobalModel(req *GlobalProductDeleteGlobalModelRequest) (resp *GlobalProductDeleteGlobalModel, err error) {
	b, err := s.post("GlobalProductDeleteGlobalModel", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductDeleteGlobalModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetGlobalModelList(req *GlobalProductGetGlobalModelListRequest) (resp *GlobalProductGetGlobalModelList, err error) {
	b, err := s.post("GlobalProductGetGlobalModelList", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetGlobalModelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductSupportSizeChart(req *GlobalProductSupportSizeChartRequest) (resp *GlobalProductSupportSizeChart, err error) {
	b, err := s.post("GlobalProductSupportSizeChart", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductSupportSizeChartResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductUpdateSizeChart(req *GlobalProductUpdateSizeChartRequest) (err error) {
	b, err := s.post("GlobalProductUpdateSizeChart", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdateSizeChartResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductCreatePublishTask(req *GlobalProductCreatePublishTaskRequest) (resp *GlobalProductCreatePublishTask, err error) {
	b, err := s.post("GlobalProductCreatePublishTask", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductCreatePublishTaskResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetPublishableShop(req *GlobalProductGetPublishableShopRequest) (resp *GlobalProductGetPublishableShop, err error) {
	b, err := s.post("GlobalProductGetPublishableShop", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetPublishableShopResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetPublishTaskResult(req *GlobalProductGetPublishTaskResultRequest) (resp *GlobalProductGetPublishTaskResult, err error) {
	b, err := s.post("GlobalProductGetPublishTaskResult", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetPublishTaskResultResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetPublishedList(req *GlobalProductGetPublishedListRequest) (resp *GlobalProductGetPublishedList, err error) {
	b, err := s.post("GlobalProductGetPublishedList", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetPublishedResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductUpdatePrice(req *GlobalProductUpdatePriceRequest) (err error) {
	b, err := s.post("GlobalProductUpdatePrice", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdatePriceResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductUpdateStock(req *GlobalProductUpdateStockRequest) (err error) {
	b, err := s.post("GlobalProductUpdateStock", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductUpdateStockResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductSetSyncField(req *GlobalProductSetSyncFieldRequest) (err error) {
	b, err := s.post("GlobalProductSetSyncField", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductSetSyncFieldResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) GlobalProductGetGlobalItemID(req *GlobalProductGetGlobalItemIDRequest) (resp *GlobalProductGetGlobalItemID, err error) {
	b, err := s.post("GlobalProductGetGlobalItemID", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetGlobalItemIDResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductCategoryRecommend(req *GlobalProductCategoryRecommendRequest) (resp *GlobalProductCategoryRecommend, err error) {
	b, err := s.post("GlobalProductCategoryRecommend", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductCategoryRecommendResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) GlobalProductGetRecommendAttribute(req *GlobalProductGetRecommendAttributeRequest) (resp *GlobalProductGetRecommendAttribute, err error) {
	b, err := s.post("GlobalProductGetRecommendAttribute", req)
	if err != nil {
		return
	}

	var wrappedResponse *GlobalProductGetRecommendAttributeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) MediaSpaceInitVideoUpload(req *MediaSpaceInitVideoUploadRequest) (resp *MediaSpaceInitVideoUpload, err error) {
	b, err := s.post("MediaSpaceInitVideoUpload", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceInitVideoUploadResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) MediaSpaceUploadVideoPart(req *MediaSpaceUploadVideoPartRequest) (err error) {
	b, err := s.post("MediaSpaceUploadVideoPart", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceUploadVideoPartResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) MediaSpaceCompleteVideoUpload(req *MediaSpaceCompleteVideoUploadRequest) (err error) {
	b, err := s.post("MediaSpaceCompleteVideoUpload", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceCompleteVideoUploadResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) MediaSpaceGetVideoUploadResult(req *MediaSpaceGetVideoUploadResultRequest) (resp *MediaSpaceGetVideoUploadResult, err error) {
	b, err := s.post("MediaSpaceGetVideoUploadResult", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceGetVideoUploadResultResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) MediaSpaceCancelVideoUpload(req *MediaSpaceCancelVideoUploadRequest) (err error) {
	b, err := s.post("MediaSpaceCancelVideoUpload", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceCancelVideoUploadResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) MediaSpaceUploadImage(req *MediaSpaceUploadImageRequest) (resp *MediaSpaceUploadImage, err error) {
	b, err := s.post("MediaSpaceUploadImage", req)
	if err != nil {
		return
	}

	var wrappedResponse *MediaSpaceUploadImageResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopGetShopInfo(req *ShopGetShopInfoRequest) (err error) {
	b, err := s.post("ShopGetShopInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopGetShopInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) ShopGetProfile(req *ShopGetProfileRequest) (resp *ShopGetProfile, err error) {
	b, err := s.post("ShopGetProfile", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopGetProfileResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopUpdateProfile(req *ShopUpdateProfileRequest) (resp *ShopUpdateProfile, err error) {
	b, err := s.post("ShopUpdateProfile", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopUpdateProfileResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopGetWarehouseDetail(req *ShopGetWarehouseDetailRequest) (resp []ShopGetWarehouseDetail, err error) {
	b, err := s.post("ShopGetWarehouseDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopGetWarehouseDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = wrappedResponse.Response
	return
}


func (s *ShopeeClient) MerchantGetMerchantInfo(req *MerchantGetMerchantInfoRequest) (err error) {
	b, err := s.post("MerchantGetMerchantInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *MerchantGetMerchantInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) MerchantGetShopListByMerchant(req *MerchantGetShopListByMerchantRequest) (err error) {
	b, err := s.post("MerchantGetShopListByMerchant", req)
	if err != nil {
		return
	}

	var wrappedResponse *MerchantGetShopListByMerchantResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderGetOrderList(req *OrderGetOrderListRequest) (resp *OrderGetOrderList, err error) {
	b, err := s.post("OrderGetOrderList", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderGetOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderGetShipmentList(req *OrderGetShipmentListRequest) (resp *OrderGetShipmentList, err error) {
	b, err := s.post("OrderGetShipmentList", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderGetShipmentResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderGetOrderDetail(req *OrderGetOrderDetailRequest) (resp *OrderGetOrderDetail, err error) {
	b, err := s.post("OrderGetOrderDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderGetOrderDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderSplitOrder(req *OrderSplitOrderRequest) (resp *OrderSplitOrder, err error) {
	b, err := s.post("OrderSplitOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderSplitOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderUnsplitOrder(req *OrderUnsplitOrderRequest) (err error) {
	b, err := s.post("OrderUnsplitOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderUnsplitOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderCancelOrder(req *OrderCancelOrderRequest) (resp *OrderCancelOrder, err error) {
	b, err := s.post("OrderCancelOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderCancelOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderHandleBuyerCancellation(req *OrderHandleBuyerCancellationRequest) (resp *OrderHandleBuyerCancellation, err error) {
	b, err := s.post("OrderHandleBuyerCancellation", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderHandleBuyerCancellationResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderSetNote(req *OrderSetNoteRequest) (err error) {
	b, err := s.post("OrderSetNote", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderSetNoteResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderAddInvoiceData(req *OrderAddInvoiceDataRequest) (err error) {
	b, err := s.post("OrderAddInvoiceData", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderAddInvoiceDataResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderGetPendingBuyerInvoiceOrderList(req *OrderGetPendingBuyerInvoiceOrderListRequest) (resp *OrderGetPendingBuyerInvoiceOrderList, err error) {
	b, err := s.post("OrderGetPendingBuyerInvoiceOrderList", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderGetPendingBuyerInvoiceOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) OrderUploadInvoiceDoc(req *OrderUploadInvoiceDocRequest) (err error) {
	b, err := s.post("OrderUploadInvoiceDoc", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderUploadInvoiceDocResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderDownloadInvoiceDoc(req *OrderDownloadInvoiceDocRequest) (err error) {
	b, err := s.post("OrderDownloadInvoiceDoc", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderDownloadInvoiceDocResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) OrderGetBuyerInvoiceInfo(req *OrderGetBuyerInvoiceInfoRequest) (err error) {
	b, err := s.post("OrderGetBuyerInvoiceInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *OrderGetBuyerInvoiceInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsGetShippingParameter(req *LogisticsGetShippingParameterRequest) (resp *LogisticsGetShippingParameter, err error) {
	b, err := s.post("LogisticsGetShippingParameter", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetShippingParameterResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsGetTrackingNumber(req *LogisticsGetTrackingNumberRequest) (resp *LogisticsGetTrackingNumber, err error) {
	b, err := s.post("LogisticsGetTrackingNumber", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetTrackingNumberResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsShipOrder(req *LogisticsShipOrderRequest) (err error) {
	b, err := s.post("LogisticsShipOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsShipOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsUpdateShippingOrder(req *LogisticsUpdateShippingOrderRequest) (err error) {
	b, err := s.post("LogisticsUpdateShippingOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsUpdateShippingOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsGetShippingDocumentParameter(req *LogisticsGetShippingDocumentParameterRequest) (resp *LogisticsGetShippingDocumentParameter, err error) {
	b, err := s.post("LogisticsGetShippingDocumentParameter", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetShippingDocumentParameterResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsCreateShippingDocument(req *LogisticsCreateShippingDocumentRequest) (resp *LogisticsCreateShippingDocument, err error) {
	b, err := s.post("LogisticsCreateShippingDocument", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsCreateShippingDocumentResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsGetShippingDocumentResult(req *LogisticsGetShippingDocumentResultRequest) (resp *LogisticsGetShippingDocumentResult, err error) {
	b, err := s.post("LogisticsGetShippingDocumentResult", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetShippingDocumentResultResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsDownloadShippingDocument(req *LogisticsDownloadShippingDocumentRequest) (err error) {
	b, err := s.post("LogisticsDownloadShippingDocument", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsDownloadShippingDocumentResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsGetShippingDocumentInfo(req *LogisticsGetShippingDocumentInfoRequest) (resp *LogisticsGetShippingDocumentInfo, err error) {
	b, err := s.post("LogisticsGetShippingDocumentInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetShippingDocumentInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsGetTrackingInfo(req *LogisticsGetTrackingInfoRequest) (resp *LogisticsGetTrackingInfo, err error) {
	b, err := s.post("LogisticsGetTrackingInfo", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetTrackingInfoResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsGetAddressList(req *LogisticsGetAddressListRequest) (resp *LogisticsGetAddressList, err error) {
	b, err := s.post("LogisticsGetAddressList", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetAddressResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsSetAddressConfig(req *LogisticsSetAddressConfigRequest) (err error) {
	b, err := s.post("LogisticsSetAddressConfig", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsSetAddressConfigResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsDeleteAddress(req *LogisticsDeleteAddressRequest) (err error) {
	b, err := s.post("LogisticsDeleteAddress", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsDeleteAddressResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) LogisticsGetChannelList(req *LogisticsGetChannelListRequest) (resp *LogisticsGetChannelList, err error) {
	b, err := s.post("LogisticsGetChannelList", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsGetChannelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsUpdateChannel(req *LogisticsUpdateChannelRequest) (resp *LogisticsUpdateChannel, err error) {
	b, err := s.post("LogisticsUpdateChannel", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsUpdateChannelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) LogisticsBatchShipOrder(req *LogisticsBatchShipOrderRequest) (resp *LogisticsBatchShipOrder, err error) {
	b, err := s.post("LogisticsBatchShipOrder", req)
	if err != nil {
		return
	}

	var wrappedResponse *LogisticsBatchShipOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileGetUnbindOrderList(req *FirstMileGetUnbindOrderListRequest) (resp *FirstMileGetUnbindOrderList, err error) {
	b, err := s.post("FirstMileGetUnbindOrderList", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGetUnbindOrderResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileGetDetail(req *FirstMileGetDetailRequest) (resp *FirstMileGetDetail, err error) {
	b, err := s.post("FirstMileGetDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGetDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileGenerateFirstMileTrackingNumber(req *FirstMileGenerateFirstMileTrackingNumberRequest) (resp *FirstMileGenerateFirstMileTrackingNumber, err error) {
	b, err := s.post("FirstMileGenerateFirstMileTrackingNumber", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGenerateFirstMileTrackingNumberResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileBindFirstMileTrackingNumber(req *FirstMileBindFirstMileTrackingNumberRequest) (resp *FirstMileBindFirstMileTrackingNumber, err error) {
	b, err := s.post("FirstMileBindFirstMileTrackingNumber", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileBindFirstMileTrackingNumberResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileUnbindFirstMileTrackingNumber(req *FirstMileUnbindFirstMileTrackingNumberRequest) (resp *FirstMileUnbindFirstMileTrackingNumber, err error) {
	b, err := s.post("FirstMileUnbindFirstMileTrackingNumber", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileUnbindFirstMileTrackingNumberResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileGetTrackingNumberList(req *FirstMileGetTrackingNumberListRequest) (resp *FirstMileGetTrackingNumberList, err error) {
	b, err := s.post("FirstMileGetTrackingNumberList", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGetTrackingNumberResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FirstMileGetWaybill(req *FirstMileGetWaybillRequest) (err error) {
	b, err := s.post("FirstMileGetWaybill", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGetWaybillResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) FirstMileGetChannelList(req *FirstMileGetChannelListRequest) (resp *FirstMileGetChannelList, err error) {
	b, err := s.post("FirstMileGetChannelList", req)
	if err != nil {
		return
	}

	var wrappedResponse *FirstMileGetChannelResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetEscrowDetail(req *PaymentGetEscrowDetailRequest) (resp *PaymentGetEscrowDetail, err error) {
	b, err := s.post("PaymentGetEscrowDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetEscrowDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentSetShopInstallmentStatus(req *PaymentSetShopInstallmentStatusRequest) (resp *PaymentSetShopInstallmentStatus, err error) {
	b, err := s.post("PaymentSetShopInstallmentStatus", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentSetShopInstallmentStatuResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetShopInstallmentStatus(req *PaymentGetShopInstallmentStatusRequest) (resp *PaymentGetShopInstallmentStatus, err error) {
	b, err := s.post("PaymentGetShopInstallmentStatus", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetShopInstallmentStatuResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetPayoutDetail(req *PaymentGetPayoutDetailRequest) (resp *PaymentGetPayoutDetail, err error) {
	b, err := s.post("PaymentGetPayoutDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetPayoutDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentSetItemInstallmentStatus(req *PaymentSetItemInstallmentStatusRequest) (resp *PaymentSetItemInstallmentStatus, err error) {
	b, err := s.post("PaymentSetItemInstallmentStatus", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentSetItemInstallmentStatuResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetItemInstallmentStatus(req *PaymentGetItemInstallmentStatusRequest) (resp *PaymentGetItemInstallmentStatus, err error) {
	b, err := s.post("PaymentGetItemInstallmentStatus", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetItemInstallmentStatuResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetPaymentMethodList(req *PaymentGetPaymentMethodListRequest) (resp []PaymentGetPaymentMethodList, err error) {
	b, err := s.post("PaymentGetPaymentMethodList", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetPaymentMethodResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetWalletTransactionList(req *PaymentGetWalletTransactionListRequest) (resp *PaymentGetWalletTransactionList, err error) {
	b, err := s.post("PaymentGetWalletTransactionList", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetWalletTransactionResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PaymentGetEscrowList(req *PaymentGetEscrowListRequest) (resp *PaymentGetEscrowList, err error) {
	b, err := s.post("PaymentGetEscrowList", req)
	if err != nil {
		return
	}

	var wrappedResponse *PaymentGetEscrowResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountAddDiscount(req *DiscountAddDiscountRequest) (resp *DiscountAddDiscount, err error) {
	b, err := s.post("DiscountAddDiscount", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountAddDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountAddDiscountItem(req *DiscountAddDiscountItemRequest) (resp *DiscountAddDiscountItem, err error) {
	b, err := s.post("DiscountAddDiscountItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountAddDiscountItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountDeleteDiscount(req *DiscountDeleteDiscountRequest) (resp *DiscountDeleteDiscount, err error) {
	b, err := s.post("DiscountDeleteDiscount", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountDeleteDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountDeleteDiscountItem(req *DiscountDeleteDiscountItemRequest) (resp *DiscountDeleteDiscountItem, err error) {
	b, err := s.post("DiscountDeleteDiscountItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountDeleteDiscountItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountGetDiscount(req *DiscountGetDiscountRequest) (resp *DiscountGetDiscount, err error) {
	b, err := s.post("DiscountGetDiscount", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountGetDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountGetDiscountList(req *DiscountGetDiscountListRequest) (resp *DiscountGetDiscountList, err error) {
	b, err := s.post("DiscountGetDiscountList", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountGetDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountUpdateDiscount(req *DiscountUpdateDiscountRequest) (resp *DiscountUpdateDiscount, err error) {
	b, err := s.post("DiscountUpdateDiscount", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountUpdateDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountUpdateDiscountItem(req *DiscountUpdateDiscountItemRequest) (resp *DiscountUpdateDiscountItem, err error) {
	b, err := s.post("DiscountUpdateDiscountItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountUpdateDiscountItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) DiscountEndDiscount(req *DiscountEndDiscountRequest) (resp *DiscountEndDiscount, err error) {
	b, err := s.post("DiscountEndDiscount", req)
	if err != nil {
		return
	}

	var wrappedResponse *DiscountEndDiscountResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealAddBundleDeal(req *BundleDealAddBundleDealRequest) (resp *BundleDealAddBundleDeal, err error) {
	b, err := s.post("BundleDealAddBundleDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealAddBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealAddBundleDealItem(req *BundleDealAddBundleDealItemRequest) (resp *BundleDealAddBundleDealItem, err error) {
	b, err := s.post("BundleDealAddBundleDealItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealAddBundleDealItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealGetBundleDealList(req *BundleDealGetBundleDealListRequest) (resp *BundleDealGetBundleDealList, err error) {
	b, err := s.post("BundleDealGetBundleDealList", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealGetBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealGetBundleDeal(req *BundleDealGetBundleDealRequest) (resp *BundleDealGetBundleDeal, err error) {
	b, err := s.post("BundleDealGetBundleDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealGetBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealGetBundleDealItem(req *BundleDealGetBundleDealItemRequest) (resp *BundleDealGetBundleDealItem, err error) {
	b, err := s.post("BundleDealGetBundleDealItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealGetBundleDealItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealUpdateBundleDeal(req *BundleDealUpdateBundleDealRequest) (resp *BundleDealUpdateBundleDeal, err error) {
	b, err := s.post("BundleDealUpdateBundleDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealUpdateBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealUpdateBundleDealItem(req *BundleDealUpdateBundleDealItemRequest) (resp *BundleDealUpdateBundleDealItem, err error) {
	b, err := s.post("BundleDealUpdateBundleDealItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealUpdateBundleDealItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealEndBundleDeal(req *BundleDealEndBundleDealRequest) (resp *BundleDealEndBundleDeal, err error) {
	b, err := s.post("BundleDealEndBundleDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealEndBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealDeleteBundleDeal(req *BundleDealDeleteBundleDealRequest) (resp *BundleDealDeleteBundleDeal, err error) {
	b, err := s.post("BundleDealDeleteBundleDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealDeleteBundleDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) BundleDealDeleteBundleDealItem(req *BundleDealDeleteBundleDealItemRequest) (resp *BundleDealDeleteBundleDealItem, err error) {
	b, err := s.post("BundleDealDeleteBundleDealItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *BundleDealDeleteBundleDealItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealAddAddOnDeal(req *AddOnDealAddAddOnDealRequest) (resp *AddOnDealAddAddOnDeal, err error) {
	b, err := s.post("AddOnDealAddAddOnDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealAddAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealAddAddOnDealMainItem(req *AddOnDealAddAddOnDealMainItemRequest) (resp *AddOnDealAddAddOnDealMainItem, err error) {
	b, err := s.post("AddOnDealAddAddOnDealMainItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealAddAddOnDealMainItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealAddAddOnDealSubItem(req *AddOnDealAddAddOnDealSubItemRequest) (resp *AddOnDealAddAddOnDealSubItem, err error) {
	b, err := s.post("AddOnDealAddAddOnDealSubItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealAddAddOnDealSubItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealDeleteAddOnDeal(req *AddOnDealDeleteAddOnDealRequest) (resp *AddOnDealDeleteAddOnDeal, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealDeleteAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealDeleteAddOnDealMainItem(req *AddOnDealDeleteAddOnDealMainItemRequest) (resp *AddOnDealDeleteAddOnDealMainItem, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDealMainItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealDeleteAddOnDealMainItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealDeleteAddOnDealSubItem(req *AddOnDealDeleteAddOnDealSubItemRequest) (resp *AddOnDealDeleteAddOnDealSubItem, err error) {
	b, err := s.post("AddOnDealDeleteAddOnDealSubItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealDeleteAddOnDealSubItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealGetAddOnDealList(req *AddOnDealGetAddOnDealListRequest) (resp *AddOnDealGetAddOnDealList, err error) {
	b, err := s.post("AddOnDealGetAddOnDealList", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealGetAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealGetAddOnDeal(req *AddOnDealGetAddOnDealRequest) (resp *AddOnDealGetAddOnDeal, err error) {
	b, err := s.post("AddOnDealGetAddOnDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealGetAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealGetAddOnDealMainItem(req *AddOnDealGetAddOnDealMainItemRequest) (resp *AddOnDealGetAddOnDealMainItem, err error) {
	b, err := s.post("AddOnDealGetAddOnDealMainItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealGetAddOnDealMainItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealGetAddOnDealSubItem(req *AddOnDealGetAddOnDealSubItemRequest) (resp *AddOnDealGetAddOnDealSubItem, err error) {
	b, err := s.post("AddOnDealGetAddOnDealSubItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealGetAddOnDealSubItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealUpdateAddOnDeal(req *AddOnDealUpdateAddOnDealRequest) (resp *AddOnDealUpdateAddOnDeal, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealUpdateAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealUpdateAddOnDealMainItem(req *AddOnDealUpdateAddOnDealMainItemRequest) (resp *AddOnDealUpdateAddOnDealMainItem, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDealMainItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealUpdateAddOnDealMainItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealUpdateAddOnDealSubItem(req *AddOnDealUpdateAddOnDealSubItemRequest) (resp *AddOnDealUpdateAddOnDealSubItem, err error) {
	b, err := s.post("AddOnDealUpdateAddOnDealSubItem", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealUpdateAddOnDealSubItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AddOnDealEndAddOnDeal(req *AddOnDealEndAddOnDealRequest) (resp *AddOnDealEndAddOnDeal, err error) {
	b, err := s.post("AddOnDealEndAddOnDeal", req)
	if err != nil {
		return
	}

	var wrappedResponse *AddOnDealEndAddOnDealResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherAddVoucher(req *VoucherAddVoucherRequest) (resp *VoucherAddVoucher, err error) {
	b, err := s.post("VoucherAddVoucher", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherAddVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherDeleteVoucher(req *VoucherDeleteVoucherRequest) (resp *VoucherDeleteVoucher, err error) {
	b, err := s.post("VoucherDeleteVoucher", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherDeleteVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherEndVoucher(req *VoucherEndVoucherRequest) (resp *VoucherEndVoucher, err error) {
	b, err := s.post("VoucherEndVoucher", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherEndVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherUpdateVoucher(req *VoucherUpdateVoucherRequest) (resp *VoucherUpdateVoucher, err error) {
	b, err := s.post("VoucherUpdateVoucher", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherUpdateVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherGetVoucher(req *VoucherGetVoucherRequest) (resp *VoucherGetVoucher, err error) {
	b, err := s.post("VoucherGetVoucher", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherGetVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) VoucherGetVoucherList(req *VoucherGetVoucherListRequest) (resp *VoucherGetVoucherList, err error) {
	b, err := s.post("VoucherGetVoucherList", req)
	if err != nil {
		return
	}

	var wrappedResponse *VoucherGetVoucherResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeAddFollowPrize(req *FollowPrizeAddFollowPrizeRequest) (resp *FollowPrizeAddFollowPrize, err error) {
	b, err := s.post("FollowPrizeAddFollowPrize", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeAddFollowPrizeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeDeleteFollowPrize(req *FollowPrizeDeleteFollowPrizeRequest) (resp *FollowPrizeDeleteFollowPrize, err error) {
	b, err := s.post("FollowPrizeDeleteFollowPrize", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeDeleteFollowPrizeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeEndFollowPrize(req *FollowPrizeEndFollowPrizeRequest) (resp *FollowPrizeEndFollowPrize, err error) {
	b, err := s.post("FollowPrizeEndFollowPrize", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeEndFollowPrizeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeUpdateFollowPrize(req *FollowPrizeUpdateFollowPrizeRequest) (resp *FollowPrizeUpdateFollowPrize, err error) {
	b, err := s.post("FollowPrizeUpdateFollowPrize", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeUpdateFollowPrizeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeGetFollowPrizeDetail(req *FollowPrizeGetFollowPrizeDetailRequest) (resp *FollowPrizeGetFollowPrizeDetail, err error) {
	b, err := s.post("FollowPrizeGetFollowPrizeDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeGetFollowPrizeDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) FollowPrizeGetFollowPrizeList(req *FollowPrizeGetFollowPrizeListRequest) (resp *FollowPrizeGetFollowPrizeList, err error) {
	b, err := s.post("FollowPrizeGetFollowPrizeList", req)
	if err != nil {
		return
	}

	var wrappedResponse *FollowPrizeGetFollowPrizeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) TopPicksGetTopPicksList(req *TopPicksGetTopPicksListRequest) (resp *TopPicksGetTopPicksList, err error) {
	b, err := s.post("TopPicksGetTopPicksList", req)
	if err != nil {
		return
	}

	var wrappedResponse *TopPicksGetTopPicksResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) TopPicksAddTopPicks(req *TopPicksAddTopPicksRequest) (resp *TopPicksAddTopPicks, err error) {
	b, err := s.post("TopPicksAddTopPicks", req)
	if err != nil {
		return
	}

	var wrappedResponse *TopPicksAddTopPickResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) TopPicksUpdateTopPicks(req *TopPicksUpdateTopPicksRequest) (resp *TopPicksUpdateTopPicks, err error) {
	b, err := s.post("TopPicksUpdateTopPicks", req)
	if err != nil {
		return
	}

	var wrappedResponse *TopPicksUpdateTopPickResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) TopPicksDeleteTopPicks(req *TopPicksDeleteTopPicksRequest) (resp *TopPicksDeleteTopPicks, err error) {
	b, err := s.post("TopPicksDeleteTopPicks", req)
	if err != nil {
		return
	}

	var wrappedResponse *TopPicksDeleteTopPickResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryAddShopCategory(req *ShopCategoryAddShopCategoryRequest) (resp *ShopCategoryAddShopCategory, err error) {
	b, err := s.post("ShopCategoryAddShopCategory", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryAddShopCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryGetShopCategoryList(req *ShopCategoryGetShopCategoryListRequest) (resp *ShopCategoryGetShopCategoryList, err error) {
	b, err := s.post("ShopCategoryGetShopCategoryList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryGetShopCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryDeleteShopCategory(req *ShopCategoryDeleteShopCategoryRequest) (resp *ShopCategoryDeleteShopCategory, err error) {
	b, err := s.post("ShopCategoryDeleteShopCategory", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryDeleteShopCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryUpdateShopCategory(req *ShopCategoryUpdateShopCategoryRequest) (resp *ShopCategoryUpdateShopCategory, err error) {
	b, err := s.post("ShopCategoryUpdateShopCategory", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryUpdateShopCategoryResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryAddItemList(req *ShopCategoryAddItemListRequest) (resp *ShopCategoryAddItemList, err error) {
	b, err := s.post("ShopCategoryAddItemList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryAddItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryGetItemList(req *ShopCategoryGetItemListRequest) (resp *ShopCategoryGetItemList, err error) {
	b, err := s.post("ShopCategoryGetItemList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryGetItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ShopCategoryDeleteItemList(req *ShopCategoryDeleteItemListRequest) (resp *ShopCategoryDeleteItemList, err error) {
	b, err := s.post("ShopCategoryDeleteItemList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ShopCategoryDeleteItemResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsGetReturnDetail(req *ReturnsGetReturnDetailRequest) (resp *ReturnsGetReturnDetail, err error) {
	b, err := s.post("ReturnsGetReturnDetail", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsGetReturnDetailResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsGetReturnList(req *ReturnsGetReturnListRequest) (resp []ReturnsGetReturnList, err error) {
	b, err := s.post("ReturnsGetReturnList", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsGetReturnResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsConfirm(req *ReturnsConfirmRequest) (resp *ReturnsConfirm, err error) {
	b, err := s.post("ReturnsConfirm", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsConfirmResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsDispute(req *ReturnsDisputeRequest) (resp *ReturnsDispute, err error) {
	b, err := s.post("ReturnsDispute", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsDisputeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsGetAvailableSolutions(req *ReturnsGetAvailableSolutionsRequest) (resp *ReturnsGetAvailableSolutions, err error) {
	b, err := s.post("ReturnsGetAvailableSolutions", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsGetAvailableSolutionResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsOffer(req *ReturnsOfferRequest) (resp *ReturnsOffer, err error) {
	b, err := s.post("ReturnsOffer", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsOfferResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) ReturnsAcceptOffer(req *ReturnsAcceptOfferRequest) (resp *ReturnsAcceptOffer, err error) {
	b, err := s.post("ReturnsAcceptOffer", req)
	if err != nil {
		return
	}

	var wrappedResponse *ReturnsAcceptOfferResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) AccountHealthShopPerformance(req *AccountHealthShopPerformanceRequest) (err error) {
	b, err := s.post("AccountHealthShopPerformance", req)
	if err != nil {
		return
	}

	var wrappedResponse *AccountHealthShopPerformanceResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) AccountHealthShopPenalty(req *AccountHealthShopPenaltyRequest) (err error) {
	b, err := s.post("AccountHealthShopPenalty", req)
	if err != nil {
		return
	}

	var wrappedResponse *AccountHealthShopPenaltyResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) PublicGetShopsByPartner(req *PublicGetShopsByPartnerRequest) (err error) {
	b, err := s.post("PublicGetShopsByPartner", req)
	if err != nil {
		return
	}

	var wrappedResponse *PublicGetShopsByPartnerResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) PublicGetMerchantsByPartner(req *PublicGetMerchantsByPartnerRequest) (err error) {
	b, err := s.post("PublicGetMerchantsByPartner", req)
	if err != nil {
		return
	}

	var wrappedResponse *PublicGetMerchantsByPartnerResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) PublicGetTokenByResendCode(req *PublicGetTokenByResendCodeRequest) (err error) {
	b, err := s.post("PublicGetTokenByResendCode", req)
	if err != nil {
		return
	}

	var wrappedResponse *PublicGetTokenByResendCodeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) PublicGetRefreshTokenByUpgradeCode(req *PublicGetRefreshTokenByUpgradeCodeRequest) (resp *PublicGetRefreshTokenByUpgradeCode, err error) {
	b, err := s.post("PublicGetRefreshTokenByUpgradeCode", req)
	if err != nil {
		return
	}

	var wrappedResponse *PublicGetRefreshTokenByUpgradeCodeResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	resp = &wrappedResponse.Response
	return
}


func (s *ShopeeClient) PushGetPushConfig(req *PushGetPushConfigRequest) (err error) {
	b, err := s.post("PushGetPushConfig", req)
	if err != nil {
		return
	}

	var wrappedResponse *PushGetPushConfigResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}


func (s *ShopeeClient) PushSetPushConfig(req *PushSetPushConfigRequest) (err error) {
	b, err := s.post("PushSetPushConfig", req)
	if err != nil {
		return
	}

	var wrappedResponse *PushSetPushConfigResponse
	err = json.Unmarshal(b, &wrappedResponse)
	if err != nil {
		return
	}

    if wrappedResponse.Error != "" {
		err = errors.New(wrappedResponse.Error)
		return
	}

	if wrappedResponse.Warning != "" {
		log.Printf("[Warning From SHOPEE]" + wrappedResponse.Warning + "\n")
		return
	}

	return
}
