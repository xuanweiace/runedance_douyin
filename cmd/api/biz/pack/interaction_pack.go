package pack

import (
	"runedance_douyin/cmd/api/biz/model/douyin"
	"runedance_douyin/kitex_gen/interaction"
)

func ConvertVideolist(videoList []*interaction.Video) (vl []*douyin.Video) {
	vl = make([]*douyin.Video, 0)
	for _, video := range videoList {
		author := &douyin.User{
			ID:            video.Author.UserId,
			Name:          video.Author.Username,
			FollowCount:   video.Author.FollowerCount,
			FollowerCount: video.Author.FollowerCount,
			IsFollow:      video.Author.IsFollow,
		}
		vl = append(vl, &douyin.Video{
			ID:            video.Id,
			Author:        author,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	return
}

func ConvertCommentlist(commentList []*interaction.Comment) (cl []*douyin.Comment) {
	cl = make([]*douyin.Comment, 0)
	for _, comment := range commentList {
		user := &douyin.User{
			ID:            comment.User.UserId,
			Name:          comment.User.Username,
			FollowCount:   comment.User.FollowerCount,
			FollowerCount: comment.User.FollowerCount,
			IsFollow:      comment.User.IsFollow,
		}
		cl = append(cl, &douyin.Comment{
			ID:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		})
	}
	return
}
