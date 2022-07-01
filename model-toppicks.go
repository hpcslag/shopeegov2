package shopeego


//=======================================================
// Object Raw Type - CollectionList
//=======================================================
type CollectionList struct {
// is_activated is whether collection is activated.
IsActivated bool `json:"is_activated,omitempty"`
// item_list is The items of top picks
ItemList []ItemList `json:"item_list"`
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
CollectionList []CollectionList `json:"collection_list"`
}
//=======================================================
// TopPicksGetTopPicksListRequest
//=======================================================
type TopPicksGetTopPicksListRequest struct {
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
// Object Raw Type - TopPicksAddTopPicks
//=======================================================
type TopPicksAddTopPicks struct {
// collection_list is The top picks list in this shop.
CollectionList []CollectionList `json:"collection_list"`
}
//=======================================================
// TopPicksAddTopPicksRequest
//=======================================================
type TopPicksAddTopPicksRequest struct {
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
// Object Raw Type - TopPicksUpdateTopPicks
//=======================================================
type TopPicksUpdateTopPicks struct {
// collection_list is The top picks list in this shop.
CollectionList []CollectionList `json:"collection_list"`
}
//=======================================================
// TopPicksUpdateTopPicksRequest
//=======================================================
type TopPicksUpdateTopPicksRequest struct {
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