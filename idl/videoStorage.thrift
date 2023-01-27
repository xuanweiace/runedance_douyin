namespace go videoStorage

struct VideoURLResponse {
    1: required string play_url;
    2: required string cover_url;
    3: required i64 author_id;
}

struct VideoURLRequest {
    1: required i64 video_id;
}

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
    VideoURLResponse queryVideoURL(1: VideoURLRequest req)
    VideoUploadResponse uploadVideoToDB(1: VideoUploadRequest req)
}
