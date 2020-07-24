package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go_BackEnd/graph/generated"
	"Go_BackEnd/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	video = model.Video{
		UserID:         input.UserID,
		Title:          input.Title,
		URL:            input.URL,
		Description:    input.Description,
		Category:       input.Category,
		Location:       input.Location,
		Views:          input.Views,
		Day:            input.Day,
		Month:          input.Month,
		Year:           input.Year,
		Thumbnail:      input.Thumbnail,
		Likes:          input.Likes,
		Dislikes:       input.Dislikes,
		AgeRestriction: input.AgeRestriction,
		Privacy:        input.Privacy,
		Premium:        input.Premium,
		Length:         input.Length,
		Time:           input.Time,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		return nil, errors.New("failed to inset new video!")
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("failed to get video")
	}

	video.UserID = input.UserID
	video.Title = input.Title
	video.URL = input.URL
	video.Description = input.Description
	video.Category = input.Category
	video.Location = input.Location
	video.Views = input.Views
	video.Day = input.Day
	video.Month = input.Month
	video.Year = input.Year
	video.Thumbnail = input.Thumbnail
	video.Likes = input.Likes
	video.Dislikes = input.Dislikes
	video.AgeRestriction = input.AgeRestriction
	video.Privacy = input.Privacy
	video.Premium = input.Premium
	video.Length = input.Length
	video.Time = input.Time

	_, error := r.DB.Model(&video).Where("id = ?", id).Update()

	if error != nil {
		return nil, errors.New("failed to update the video!")
	}

	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id string, input *model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	var user model.User

	user = model.User{
		Name:            input.Name,
		Email:           input.Email,
		Premium:         input.Premium,
		ImgURL:          input.ImgURL,
		Subscribers:     input.Subscribers,
		LikedVideo:      input.LikedVideo,
		DislikedVideo:   input.DislikedVideo,
		LikedComment:    input.LikedComment,
		DislikedComment: input.DislikedComment,
	}
	_, error := r.DB.Model(&user).Insert()

	if error != nil {
		return nil, errors.New("failed to insert")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("failed to get user")
	}

	user.Name = input.Name
	user.Email = input.Email
	user.ImgURL = input.ImgURL
	user.Subscribers = input.Subscribers
	user.Premium = input.Premium
	user.LikedVideo = input.LikedVideo
	user.DislikedVideo = input.DislikedVideo
	user.LikedComment = input.LikedComment
	user.DislikedComment = input.DislikedComment

	_, error := r.DB.Model(&user).Where("id = ?", id).Update()

	if error != nil {
		return nil, errors.New("failed to update user")
	}

	return &user, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	var comment model.Comment

	comment = model.Comment{
		UserID:      input.UserID,
		VideoID:     input.VideoID,
		Likes:       input.Likes,
		Dislikes:    input.Dislikes,
		Description: input.Description,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, errors.New("failed to add a new comment!")
	}

	return &comment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input *model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateReply(ctx context.Context, input *model.NewReply) (*model.Reply, error) {
	var reply model.Reply

	reply = model.Reply{
		CommentID:   input.CommentID,
		UserID:      input.UserID,
		Likes:       input.Likes,
		Dislikes:    input.Dislikes,
		Description: input.Description,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
	}

	_, err := r.DB.Model(&reply).Insert()

	if err != nil {
		return nil, errors.New("failed to reply!")
	}

	return &reply, nil
}

func (r *mutationResolver) UpdateReply(ctx context.Context, id string, input *model.NewReply) (*model.Reply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteReply(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateChannel(ctx context.Context, input *model.NewChannel) (*model.Channel, error) {
	var channel model.Channel

	channel = model.Channel{
		UserID:        input.UserID,
		BackgroundURL: input.BackgroundURL,
		Description:   input.Description,
		JoinDate:      input.JoinDate,
		Links:         input.Links,
	}

	_, err := r.DB.Model(&channel).Insert()

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *mutationResolver) UpdateChannel(ctx context.Context, id string, input *model.NewChannel) (*model.Channel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	var post model.Post

	post = model.Post{
		ChannelID:   input.ChannelID,
		Description: input.Description,
		Picture:     input.Picture,
		Date:        input.Date,
		Likes:       input.Likes,
		Dislikes:    input.Dislikes,
	}

	_, err := r.DB.Model(&post).Insert()

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input *model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	error := r.DB.Model(&videos).Order("id").Select()

	if error != nil {
		return nil, errors.New("failed to get all videos!")
	}

	return videos, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	error := r.DB.Model(&users).Select()

	if error != nil {
		return nil, errors.New("failed to get all users")
	}

	return users, error
}

func (r *queryResolver) Replies(ctx context.Context, commentID string) ([]*model.Reply, error) {
	var replies []*model.Reply

	err := r.DB.Model(&replies).Where("comment_id = ?", commentID).Select()

	if err != nil {
		return nil, errors.New("failed to reply a comment!")
	}

	return replies, nil
}

func (r *queryResolver) Channel(ctx context.Context, id string) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *queryResolver) Posts(ctx context.Context, channelID int) ([]*model.Post, error) {
	var posts []*model.Post

	err := r.DB.Model(&posts).Where("channel_id = ?", channelID).Select()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	error := r.DB.Model(&user).Where("id = ?", id).First()

	if error != nil {
		return nil, errors.New("failed to get this user")
	}

	return &user, nil
}

func (r *queryResolver) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	error := r.DB.Model(&user).Where("email = ?", email).First()

	if error != nil {
		return nil, errors.New("failed to get this user by email")
	}

	return &user, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("failed to get this video")
	}

	return &video, nil
}

func (r *queryResolver) GetAllComments(ctx context.Context, videoID int) ([]*model.Comment, error) {
	var comments []*model.Comment

	error := r.DB.Model(&comments).Where("video_id = ?", videoID).Select()

	if error != nil {
		return nil, errors.New("failed to get all comments")
	}

	return comments, nil
}

func (r *queryResolver) GetChannelByUser(ctx context.Context, userID int) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("user_id = ?", userID).First()

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *queryResolver) GetAllVideosByTitle(ctx context.Context, title string) ([]*model.Video, error) {
	var videos []*model.Video

	_, err := r.DB.Query(&videos, "SELECT * FROM videos WHERE LOWER(title) like LOWER(?)", title)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
