package shopeego
//=======================================================
// MediaSpaceInitVideoUploadRequest
//=======================================================
type MediaSpaceInitVideoUploadRequest struct {
    // file_md5 is md5 of video file
    FileMd5 string `json:"file_md5"`
    // file_size is size of video file, in bytes, maximum is 30MB
    FileSize int `json:"file_size"`
}
//=======================================================
// MediaSpaceInitVideoUploadResponse
//=======================================================
type MediaSpaceInitVideoUploadResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of the API request for error tracking
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
}
//=======================================================
// MediaSpaceUploadVideoPartRequest
//=======================================================
type MediaSpaceUploadVideoPartRequest struct {
    // video_upload_id is The video_upload_id in the response of initiate_video_upload
    VideoUploadID string `json:"video_upload_id"`
    // part_seq is Sequence of the current part, starts from 0
    PartSeq int `json:"part_seq"`
    // content_md5 is md5 of this part
    ContentMd5 string `json:"content_md5"`
}
//=======================================================
// MediaSpaceUploadVideoPartResponse
//=======================================================
type MediaSpaceUploadVideoPartResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
}


//=======================================================
// Object Raw Type - ReportData
//=======================================================
type ReportData struct {
// upload_cost is Time used for uploading the video file via upload_video_part api, in milliseconds. For video upload performance tracking purpose.
UploadCost int `json:"upload_cost,omitempty"`
}
//=======================================================
// MediaSpaceCompleteVideoUploadRequest
//=======================================================
type MediaSpaceCompleteVideoUploadRequest struct {
    // video_upload_id is The ID of this upload session, returned in init_video_upload.
    VideoUploadID string `json:"video_upload_id"`
    // part_seq_list is All uploaded sequence number.
    PartSeqList []int `json:"part_seq_list"`
    // report_data is 
    ReportData ReportData `json:"report_data"`
}
//=======================================================
// MediaSpaceCompleteVideoUploadResponse
//=======================================================
type MediaSpaceCompleteVideoUploadResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier of this upload session, used in following video upload request and item creating and/or updating
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
}
//=======================================================
// MediaSpaceGetVideoUploadResultRequest
//=======================================================
type MediaSpaceGetVideoUploadResultRequest struct {
    // video_upload_id is 
    VideoUploadID string `json:"video_upload_id"`
}
//=======================================================
// MediaSpaceGetVideoUploadResultResponse
//=======================================================
type MediaSpaceGetVideoUploadResultResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
}
//=======================================================
// MediaSpaceCancelVideoUploadRequest
//=======================================================
type MediaSpaceCancelVideoUploadRequest struct {
    // video_upload_id is The ID of this upload session, returned in init_video_upload.
    VideoUploadID string `json:"video_upload_id"`
}
//=======================================================
// MediaSpaceCancelVideoUploadResponse
//=======================================================
type MediaSpaceCancelVideoUploadResponse struct {
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // warning is Warning message.
    Warning string `json:"warning,omitempty"`
}
//=======================================================
// MediaSpaceUploadImageRequest
//=======================================================
type MediaSpaceUploadImageRequest struct {
    // scene is The scene where the picture is used, The value range is normal or desc; normal: we will process the image as a square image, it is recommended to use when uploading item image; desc: we will not process the image, it is recommended to use when uploading the image of extend_description, if you do not upload this field, it will be normal.
    Scene string `json:"scene,omitempty"`
}
//=======================================================
// MediaSpaceUploadImageResponse
//=======================================================
type MediaSpaceUploadImageResponse struct {
    // error is Indicate error type if hit error. Empty if no error happened.
    Error string `json:"error,omitempty"`
    // message is Indicate error details if hit error. Empty if no error happened.
    Message string `json:"message,omitempty"`
    // warning is Indicate warning message you should take care.
    Warning string `json:"warning,omitempty"`
    // request_id is The identifier for an API request for error tracking.
    RequestID string `json:"request_id,omitempty"`
    // response is 
    Response Response `json:"response,omitempty"`
}