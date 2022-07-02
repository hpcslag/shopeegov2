package shopeego


//=======================================================
// Object Raw Type - MediaSpaceInitVideoUpload
//=======================================================
type MediaSpaceInitVideoUpload struct {
// video_upload_id is The identifier of this upload session, used in following video upload request and item creating and/or updating
VideoUploadID string `json:"video_upload_id,omitempty"`
}
//=======================================================
// MediaSpaceInitVideoUploadRequest
//=======================================================
type MediaSpaceInitVideoUploadRequest struct {
    V2RequestAuthenticationParams
    

    // file_md5 is md5 of video file
    FileMd5 string `json:"file_md5"`
    // file_size is size of video file, in bytes, maximum is 30MB
    FileSize int `json:"file_size"`
}
//=======================================================
// MediaSpaceInitVideoUploadResponse
//=======================================================
type MediaSpaceInitVideoUploadResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response MediaSpaceInitVideoUpload `json:"response"`
}
//=======================================================
// MediaSpaceUploadVideoPartRequest
//=======================================================
type MediaSpaceUploadVideoPartRequest struct {
    V2RequestAuthenticationParams
    

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
    // 通用的 Response 回傳參數
    V2UnityResponse

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
    V2RequestAuthenticationParams
    

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
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - MediaSpaceGetVideoUploadResultVideoInfoVideoUrl
//=======================================================
type MediaSpaceGetVideoUploadResultVideoInfoVideoUrl struct {
// video_url_region is The region of this video URL.
VideoUrlRegion string `json:"video_url_region,omitempty"`
// video_url is Video playback URL.
VideoUrl string `json:"video_url,omitempty"`
}


//=======================================================
// Object Raw Type - MediaSpaceGetVideoUploadResultVideoInfoThumbnailUrl
//=======================================================
type MediaSpaceGetVideoUploadResultVideoInfoThumbnailUrl struct {
// image_url_region is The region of this image URL.
ImageUrlRegion string `json:"image_url_region,omitempty"`
// image_url is Image display URL.
ImageUrl string `json:"image_url,omitempty"`
}


//=======================================================
// Object Raw Type - MediaSpaceGetVideoUploadResultVideoInfo
//=======================================================
type MediaSpaceGetVideoUploadResultVideoInfo struct {
// video_url_list is Video playback URL list.
VideoUrlList []MediaSpaceGetVideoUploadResultVideoInfoVideoUrl `json:"video_url_list"`
// thumbnail_url_list is Video thumbnail image URL list.
ThumbnailUrlList []MediaSpaceGetVideoUploadResultVideoInfoThumbnailUrl `json:"thumbnail_url_list"`
// duration is Duration of this video, in seconds.
Duration int `json:"duration,omitempty"`
}


//=======================================================
// Object Raw Type - MediaSpaceGetVideoUploadResult
//=======================================================
type MediaSpaceGetVideoUploadResult struct {
// status is Current status of this video upload session. could be: INITIATED(waiting for part uploading and/or the complete_video_upload API call), TRANSCODING(has received all video parts, and is transcoding the video file), SUCCEEDED(transcoding completed, and this upload_id can now be used for item adding/updating), FAILED(this upload failed, see the message filed for some info), CANCELLED(this upload is cancelled)
Status string `json:"status,omitempty"`
// video_info is Transcoded video info, will be present if status is SUCCEEDED.
VideoInfo MediaSpaceGetVideoUploadResultVideoInfo `json:"video_info"`
// message is Detail error message if video uploading/transcoding failed.
Message string `json:"message,omitempty"`
}
//=======================================================
// MediaSpaceGetVideoUploadResultRequest
//=======================================================
type MediaSpaceGetVideoUploadResultRequest struct {
    V2RequestAuthenticationParams
    

    // video_upload_id is 
    VideoUploadID string `json:"video_upload_id"`
}
//=======================================================
// MediaSpaceGetVideoUploadResultResponse
//=======================================================
type MediaSpaceGetVideoUploadResultResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response MediaSpaceGetVideoUploadResult `json:"response"`
}
//=======================================================
// MediaSpaceCancelVideoUploadRequest
//=======================================================
type MediaSpaceCancelVideoUploadRequest struct {
    V2RequestAuthenticationParams
    

    // video_upload_id is The ID of this upload session, returned in init_video_upload.
    VideoUploadID string `json:"video_upload_id"`
}
//=======================================================
// MediaSpaceCancelVideoUploadResponse
//=======================================================
type MediaSpaceCancelVideoUploadResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

}


//=======================================================
// Object Raw Type - MediaSpaceUploadImageImageInfoImageUrl
//=======================================================
type MediaSpaceUploadImageImageInfoImageUrl struct {
// image_url_region is Region of image url
ImageUrlRegion string `json:"image_url_region,omitempty"`
// image_url is image url
ImageUrl string `json:"image_url,omitempty"`
}


//=======================================================
// Object Raw Type - MediaSpaceUploadImageImageInfo
//=======================================================
type MediaSpaceUploadImageImageInfo struct {
// image_id is Id of image 
ImageID string `json:"image_id,omitempty"`
// image_url_list is Image URL of each region
ImageUrlList []MediaSpaceUploadImageImageInfoImageUrl `json:"image_url_list"`
}


//=======================================================
// Object Raw Type - MediaSpaceUploadImage
//=======================================================
type MediaSpaceUploadImage struct {
// image_info is 
ImageInfo MediaSpaceUploadImageImageInfo `json:"image_info"`
}
//=======================================================
// MediaSpaceUploadImageRequest
//=======================================================
type MediaSpaceUploadImageRequest struct {
    V2RequestAuthenticationParams
    

    // scene is The scene where the picture is used, The value range is normal or desc; normal: we will process the image as a square image, it is recommended to use when uploading item image; desc: we will not process the image, it is recommended to use when uploading the image of extend_description, if you do not upload this field, it will be normal.
    Scene string `json:"scene,omitempty"`
}
//=======================================================
// MediaSpaceUploadImageResponse
//=======================================================
type MediaSpaceUploadImageResponse struct {
    // 通用的 Response 回傳參數
    V2UnityResponse

    // response is 
    Response MediaSpaceUploadImage `json:"response,omitempty"`
}