namespace go videoProcess

struct VideoProcessUploadRequest {
    1:required i64 author_id;
    2:required string title;
    3:required binary video_data;
}
struct VideoProcessUploadResponse {
    1:required bool result;
    2: string message;
}
struct VideoInfoResponse {
    1:required i64 video_id;
    2:required i64 author_id;
    3:required string storage_id;
    4:required i32 favorite_count;
    5:required i32 comment_count;
    6:required string title;
}
struct VideoInfoRequest {
    1:required i64 video_id;
}
struct VideoListResponse {
    1: required list<i64> published;
}
service VideoProcessService {
    VideoInfoResponse getVideoInfo (1:VideoInfoRequest req)
    VideoProcessUploadResponse uploadVideo (1:VideoProcessUploadRequest req)
    VideoListResponse getVideoList (1:i64 author_id)
}

