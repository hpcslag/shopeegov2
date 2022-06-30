package shopeego
//=======================================================
// FollowPrizeAddFollowPrizeRequest
//=======================================================
type FollowPrizeAddFollowPrizeRequest struct {
    // follow_prize_name is <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
    FollowPrizeName string `json:"follow_prize_name"`
    // start_time is <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
    StartTime int `json:"start_time"`
    // end_time is <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
    EndTime int `json:"end_time"`
    // usage_quantity is <p>Please enter a value between 1 and 200000.</p>
    UsageQuantity int `json:"usage_quantity"`
    // min_spend is <p>The minimum spend required for using this follow prize.<br /></p>
    MinSpend float64 `json:"min_spend"`
    // reward_type is <p>The reward type of the follow prize.The available values are:1:discount---fix amount,2:discount---by percentage,3:coin cash back.<br /></p>
    RewardType int `json:"reward_type"`
    // discount_amount is <p>The discount amount set for this particular follow prize.Only fill in when you are creating a fix amount follow prize.<br /></p>
    DiscountAmount float64 `json:"discount_amount,omitempty"`
    // percentage is <p>The discount percentage set for this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.Discount percentage (reward_type ==2) or Percentage of coins cash back (reward_type==3).<br /></p>
    Percentage int `json:"percentage,omitempty"`
    // max_price is <p>The max amount of discount/value a user can enjoy by using this particular follow prize. Only fill in when you are creating a discount percentage follow prize or coins cashback follow prize.<br /></p>
    MaxPrice float64 `json:"max_price,omitempty"`
}
//=======================================================
// FollowPrizeAddFollowPrizeResponse
//=======================================================
type FollowPrizeAddFollowPrizeResponse struct {
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// FollowPrizeDeleteFollowPrizeRequest
//=======================================================
type FollowPrizeDeleteFollowPrizeRequest struct {
    // campaign_id is <p>The unique identifier for the created follow prize.<br /></p>
    CampaignID int `json:"campaign_id"`
}
//=======================================================
// FollowPrizeDeleteFollowPrizeResponse
//=======================================================
type FollowPrizeDeleteFollowPrizeResponse struct {
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// FollowPrizeEndFollowPrizeRequest
//=======================================================
type FollowPrizeEndFollowPrizeRequest struct {
    // campaign_id is <p>The unique identifier for the created follow prize.<br /></p>
    CampaignID int `json:"campaign_id"`
}
//=======================================================
// FollowPrizeEndFollowPrizeResponse
//=======================================================
type FollowPrizeEndFollowPrizeResponse struct {
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// FollowPrizeUpdateFollowPrizeRequest
//=======================================================
type FollowPrizeUpdateFollowPrizeRequest struct {
    // follow_prize_name is <p>The name of the follow prize,The follow prize name length max limit is 20.<br /></p>
    FollowPrizeName string `json:"follow_prize_name,omitempty"`
    // campaign_id is <p>The unique identifier for the created follow prize.<br /></p>
    CampaignID int `json:"campaign_id"`
    // start_time is <p>The timing from when the follow prize is valid,the start time later than the current time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
    StartTime int `json:"start_time,omitempty"`
    // end_time is <p>The timing until when the follow prize is still valid,the end time must be greater than the start time by at least 1 day and end time cannot exceed 3 months after start time.If the start time and end time passed in by the seller overlap with other upcoming/ongoing activities, it will prompt "Another Follow Prize voucher already exists during this time period, please set another period."<br /></p>
    EndTime int `json:"end_time,omitempty"`
    // usage_quantity is <p>Please enter a value between 1 and 200000.</p>
    UsageQuantity int `json:"usage_quantity,omitempty"`
    // min_spend is <p>The minimum spend required for using this follow prize.<br /></p>
    MinSpend int `json:"min_spend,omitempty"`
}
//=======================================================
// FollowPrizeUpdateFollowPrizeResponse
//=======================================================
type FollowPrizeUpdateFollowPrizeResponse struct {
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response,omitempty"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error details if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// FollowPrizeGetFollowPrizeDetailRequest
//=======================================================
type FollowPrizeGetFollowPrizeDetailRequest struct {
    // campaign_id is <p>The unique identifier for the created follow prize.<br /></p>
    CampaignID int `json:"campaign_id,omitempty"`
}
//=======================================================
// FollowPrizeGetFollowPrizeDetailResponse
//=======================================================
type FollowPrizeGetFollowPrizeDetailResponse struct {
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response,omitempty"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}
//=======================================================
// FollowPrizeGetFollowPrizeListRequest
//=======================================================
type FollowPrizeGetFollowPrizeListRequest struct {
    // page_no is <p>Specifies the page number of data to return in the current call. Default to be 1.<br /></p>
    PageNo int `json:"page_no,omitempty"`
    // page_size is <p>Use the 'page_size' filters to control the maximum number of entries to retrieve per page (i.e., per call). Default to be 20 and allowed input is from 1- 100.<br /></p>
    PageSize int `json:"page_size,omitempty"`
    // status is <p>The status filter for retrieving follow prize list. Available value: upcoming/ongoing/expired/all.<br /></p>
    Status string `json:"status"`
}
//=======================================================
// FollowPrizeGetFollowPrizeListResponse
//=======================================================
type FollowPrizeGetFollowPrizeListResponse struct {
    // response is <p>Detailed informations you are querying.<br /></p>
    Response Response `json:"response"`
    // error is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Error string `json:"error,omitempty"`
    // message is <p>Indicate error type if hit error. Empty if no error happened.<br /></p>
    Message string `json:"message,omitempty"`
    // request_id is <p>The identifier of the API request for error tracking.<br /></p>
    RequestID string `json:"request_id,omitempty"`
}