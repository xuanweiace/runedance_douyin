namespace go videoStorage

struct VideoUploadRequest {
    1:required binary video_data;
    2:required i64 video_id;
    3:required binary cover_data;
}
struct VideoUploadResponse {
    1:required bool result;
    2: string message;
}

service VideoStorageService {
    VideoUploadResponse uploadVideoToDB(1: VideoUploadRequest req)
}
