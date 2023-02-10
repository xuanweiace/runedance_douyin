namespace go videoStorage

struct VideoStorageUploadRequest {
    1:optional binary video_data;
    2:required i64 video_id;
    3:optional binary cover_data;
}
service VideoStorageService {
    string uploadVideoToDB(1: VideoStorageUploadRequest req)
}