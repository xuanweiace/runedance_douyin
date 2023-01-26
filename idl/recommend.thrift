namespace go recommend

struct RecommendResponse {
    1: required list<i64> Recommended;
}

service RecommendService {
    RecommendResponse getRecommended()
}