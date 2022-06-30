package shopeego
//=======================================================
// TopPicksGetTopPicksListRequest
//=======================================================
type TopPicksGetTopPicksListRequest struct {
}
//=======================================================
// TopPicksGetTopPicksListResponse
//=======================================================
type TopPicksGetTopPicksListResponse struct {
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
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
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
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
    // request_id is The identifier for an API request for error tracking. 
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // response is Detail informations you are querying.
    Response Response `json:"response"`
}