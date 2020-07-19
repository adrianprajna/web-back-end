// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID          string `json:"id"`
	UserID      int    `json:"user_id"`
	VideoID     int    `json:"video_id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Description string `json:"description"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
}

type Playlist struct {
	ID          string  `json:"id"`
	UserID      int     `json:"user_id"`
	Name        *string `json:"name"`
	Privacy     string  `json:"privacy"`
	Description string  `json:"description"`
	Views       int     `json:"views"`
	Day         int     `json:"day"`
	Month       int     `json:"month"`
	Year        int     `json:"year"`
}

type Reply struct {
	ID          string `json:"id"`
	CommentID   int    `json:"comment_id"`
	UserID      int    `json:"user_id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Description string `json:"description"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Premium     bool   `json:"premium"`
	ImgURL      string `json:"img_url"`
	Subscribers int    `json:"subscribers"`
}

type Video struct {
	ID             string `json:"id"`
	UserID         int    `json:"user_id"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	Location       string `json:"location"`
	Views          int    `json:"views"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Thumbnail      string `json:"thumbnail"`
	Likes          int    `json:"likes"`
	Dislikes       int    `json:"dislikes"`
	AgeRestriction bool   `json:"age_restriction"`
	Privacy        string `json:"privacy"`
	Premium        bool   `json:"premium"`
}

type NewComment struct {
	UserID      int    `json:"user_id"`
	VideoID     int    `json:"video_id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Description string `json:"description"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
}

type NewPlaylist struct {
	UserID      int     `json:"user_id"`
	Name        *string `json:"name"`
	Privacy     string  `json:"privacy"`
	Description string  `json:"description"`
	Views       int     `json:"views"`
	Day         int     `json:"day"`
	Month       int     `json:"month"`
	Year        int     `json:"year"`
}

type NewReply struct {
	CommentID   int    `json:"comment_id"`
	UserID      int    `json:"user_id"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Description string `json:"description"`
	Day         int    `json:"day"`
	Month       int    `json:"month"`
	Year        int    `json:"year"`
}

type NewUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Premium     bool   `json:"premium"`
	ImgURL      string `json:"img_url"`
	Subscribers int    `json:"subscribers"`
}

type NewVideo struct {
	UserID         int    `json:"user_id"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	Location       string `json:"location"`
	Views          int    `json:"views"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Thumbnail      string `json:"thumbnail"`
	Likes          int    `json:"likes"`
	Dislikes       int    `json:"dislikes"`
	AgeRestriction bool   `json:"age_restriction"`
	Privacy        string `json:"privacy"`
	Premium        bool   `json:"premium"`
}
