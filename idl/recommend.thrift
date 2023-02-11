namespace go recommend

struct RecommendResponse {
    1: required list<i64> Recommended;
}

service RecommendService {
    RecommendResponse getRecommended(1: i64 user)
}