namespace go videoStorage

struct VideoURLResponse {
    1: required string playURL;
    2: required string coverURL;
    3: required i64 authorID;
}

struct VideoURLRequest {
    1: required i64 videoID;
}

service VideoService {
    VideoURLResponse queryVideoURL(1: VideoURLRequest req)
}
