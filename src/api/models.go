package main

import (
	"time"
)

type DatasetSummaryVo struct {
	Title	string `form:"title" json:"title" binding:"required"`
	Desc	string `form:"desc" json:"desc" binding:"required"`
	Api		string `form:"api" json:"api" binding:"required"`
}

type YoutubeInfoVo struct {
	FileName	string `form:"filename" json:"filename" binding:"required"`
	Url 		string `form:"url" json:"url" binding:"required"`
}

type LprItem struct {
    Tag 		[]LprTag	`json:"tag"`
    Filename 	string 		`json:"filename"`
}
type LprTag struct {
    Confidences  []int    `json:"confidences"`
    ObjectHeight int      `json:"objectHeight"`
    ObjectPicY   int      `json:"objectPicY"`
    ObjectTypes  []string `json:"objectTypes"`
    ObjectWidth  int      `json:"objectWidth"`
	ObjectPicX   int      `json:"objectPicX"`
	PlateNumber  string   `json:"plateNumber"`
}

type YoloItem struct {
    Tag 		[]YoloTag 	`json:"tag"`
    Filename 	string 		`json:"filename"`
}
type YoloTag struct {
    Confidences  []int    `json:"confidences"`
    ObjectHeight int      `json:"objectHeight"`
    ObjectPicY   int      `json:"objectPicY"`
    ObjectTypes  []string `json:"objectTypes"`
    ObjectWidth  int      `json:"objectWidth"`
    ObjectPicX   int      `json:"objectPicX"`
}

type SubtitleTagDataSetVo struct {
	Desc string 		`json:"desc"`
	Data []SubtitleTag  `json:"data"`
}

type SubtitleDataSetVo struct {
	Desc string		`json:"desc"`
	Data []Subtitle	`json:"data"`
}

type SubtitleTag struct {
	Id  		int 	`json:"id"`
	TagName 	string 	`json:"tagName"`
	Thumbnail	string 	`json:"thumbnail"`
}

type Subtitle struct {
	Id  		int 		`json:"id"`
	Title 		string 	`json:"title"`
	Url 		string 		`json:"url"`
	Thumbnail	string `json:"thumbnail"`
}

type TrainTwOrgDataSetVo struct {
	Title	string 			`json:"title"`
	Desc 	string			`json:"desc"`
	Data 	[]TrainTwOrgVo	`json:"data"`
}

type TrainTwOrgVo struct {
	CarAccidentID  	string `json:"carAccidentID"`
	Title 			string `json:"title"`
	YoutubeId 		string `json:"youtubeId"`
	Url 			string `json:"url"`
	Thumbnail 		string `json:"thumbnail"`
}

type TrainTwTagVo struct {
	Id 			int `json:"id"`
	YoutubeId  	string `json:"youtubeId"`
	Object 		string `json:object"`
	Filename 	string `json:"filename"`
	XNum 		int `json:"x_num"`
	YNum 		int `json:"y_num"`
	Width 		int `json:"width"`
	Height 		int `json:"height"`
}

type VideoVo struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle         string `json:"channelTitle"`
			CategoryID           string `json:"categoryId"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
	} `json:"items"`
}

type SearchVo struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
		} `json:"snippet"`
	} `json:"items"`
}