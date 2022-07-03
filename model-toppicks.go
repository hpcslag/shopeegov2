package shopeego


//=======================================================
// Object Raw Type - TopPicksGetTopPicksListCollectionItem
//=======================================================
type TopPicksGetTopPicksListCollectionItem struct {
// item_name is The name of item.
ItemName string `json:"item_name,omitempty"`
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// current_price is The price before tax of item.
CurrentPrice float64 `json:"current_price,omitempty,string"`
// inflated_price_of_current_price is The price after tax of item.
InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price,omitempty,string"`
// sales is The sales of  item.
Sales int `json:"sales,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksGetTopPicksListCollection
//=======================================================
type TopPicksGetTopPicksListCollection struct {
// is_activated is whether collection is activated.
IsActivated bool `json:"is_activated,omitempty"`
// item_list is The items of top picks
ItemList []TopPicksGetTopPicksListCollectionItem `json:"item_list"`
// top_picks_id is collection id.
TopPicksID int `json:"top_picks_id,omitempty"`
// name is The title of  top picks.
Name string `json:"name,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksGetTopPicksList
//=======================================================
type TopPicksGetTopPicksList struct {
// collection_list is The top picks list in this shop.
CollectionList []TopPicksGetTopPicksListCollection `json:"collection_list"`
}
//=======================================================
// TopPicksGetTopPicksListRequest
//=======================================================
type TopPicksGetTopPicksListRequest struct {
    V2RequestAuthenticationParams
    

}
//=======================================================
// TopPicksGetTopPicksListResponse
//=======================================================
type TopPicksGetTopPicksListResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response TopPicksGetTopPicksList `json:"response"`
}


//=======================================================
// Object Raw Type - TopPicksAddTopPicksCollectionItem
//=======================================================
type TopPicksAddTopPicksCollectionItem struct {
// item_name is The name of item.
ItemName string `json:"item_name,omitempty"`
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// current_price is The price before tax of item.
CurrentPrice float64 `json:"current_price,omitempty,string"`
// inflated_price_of_current_price is The price after tax of item.
InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price,omitempty,string"`
// sales is The sales of item.
Sales int `json:"sales,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksAddTopPicksCollection
//=======================================================
type TopPicksAddTopPicksCollection struct {
// is_activated is whether collection is activated.
IsActivated bool `json:"is_activated,omitempty"`
// item_list is The items of top picks
ItemList []TopPicksAddTopPicksCollectionItem `json:"item_list"`
// top_picks_id is Collection id.
TopPicksID int `json:"top_picks_id,omitempty"`
// name is The title of top picks.
Name string `json:"name,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksAddTopPicks
//=======================================================
type TopPicksAddTopPicks struct {
// collection_list is The top picks list in this shop.
CollectionList []TopPicksAddTopPicksCollection `json:"collection_list"`
}
//=======================================================
// TopPicksAddTopPicksRequest
//=======================================================
type TopPicksAddTopPicksRequest struct {
    V2RequestAuthenticationParams
    

    // name is 
    Name string `json:"name"`
    // item_id_list is 
    ItemIdList []int `json:"item_id_list"`
    // is_activated is 
    IsActivated bool `json:"is_activated"`
}
//=======================================================
// TopPicksAddTopPicksResponse
//=======================================================
type TopPicksAddTopPicksResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response TopPicksAddTopPicks `json:"response"`
}


//=======================================================
// Object Raw Type - TopPicksUpdateTopPicksCollectionItem
//=======================================================
type TopPicksUpdateTopPicksCollectionItem struct {
// item_name is The name of item.
ItemName string `json:"item_name,omitempty"`
// item_id is The id of item.
ItemID int `json:"item_id,omitempty"`
// current_price is The price before tax of item.
CurrentPrice float64 `json:"current_price,omitempty,string"`
// inflated_price_of_current_price is The price after tax of item.
InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price,omitempty,string"`
// sales is The sales of item.
Sales int `json:"sales,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksUpdateTopPicksCollection
//=======================================================
type TopPicksUpdateTopPicksCollection struct {
// is_activated is whether is activated
IsActivated bool `json:"is_activated,omitempty"`
// item_list is a list of item
ItemList []TopPicksUpdateTopPicksCollectionItem `json:"item_list"`
// top_picks_id is collection id
TopPicksID int `json:"top_picks_id,omitempty"`
// name is collection name
Name string `json:"name,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksUpdateTopPicks
//=======================================================
type TopPicksUpdateTopPicks struct {
// collection_list is The top picks list in this shop.
CollectionList []TopPicksUpdateTopPicksCollection `json:"collection_list"`
}
//=======================================================
// TopPicksUpdateTopPicksRequest
//=======================================================
type TopPicksUpdateTopPicksRequest struct {
    V2RequestAuthenticationParams
    

    // top_picks_id is collection id
    TopPicksID int `json:"top_picks_id"`
    // name is collection name
    Name string `json:"name,omitempty"`
    // item_id_list is a list of item id, and we will cover old item_ids by new_item_ids
    ItemIdList []int `json:"item_id_list,omitempty"`
    // is_activated is if true, we will close other collection and open this collection
    IsActivated bool `json:"is_activated,omitempty"`
}
//=======================================================
// TopPicksUpdateTopPicksResponse
//=======================================================
type TopPicksUpdateTopPicksResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response TopPicksUpdateTopPicks `json:"response,omitempty"`
}


//=======================================================
// Object Raw Type - TopPicksDeleteTopPicks
//=======================================================
type TopPicksDeleteTopPicks struct {
// top_picks_id is collection id
TopPicksID int `json:"top_picks_id,omitempty"`
}
//=======================================================
// TopPicksDeleteTopPicksRequest
//=======================================================
type TopPicksDeleteTopPicksRequest struct {
    V2RequestAuthenticationParams
    

    // top_picks_id is collection id
    TopPicksID int `json:"top_picks_id"`
}
//=======================================================
// TopPicksDeleteTopPicksResponse
//=======================================================
type TopPicksDeleteTopPicksResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is Detail informations you are querying.
    Response TopPicksDeleteTopPicks `json:"response"`
}