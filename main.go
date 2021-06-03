// Package jikipedia 小鸡词典查梗插件
// jikipedia.Register(&bot)
package jikipedia

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mcoo/OPQBot"
	"github.com/mcoo/requests"
)

type SearchResult struct {
	Data []struct {
		ID   int `json:"id"`
		Term struct {
			ID     int    `json:"id"`
			Title  string `json:"title"`
			Status string `json:"status"`
		} `json:"term"`
		Tags []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Status       string `json:"status"`
			Category     string `json:"category"`
			Relationship string `json:"relationship"`
			Leaf         bool   `json:"leaf"`
		} `json:"tags"`
		Content string `json:"content"`
		Author  struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			Role         string `json:"role"`
			AvatarDetail struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"avatar_detail"`
			ScaledAvatarDetail struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"scaled_avatar_detail"`
			BackgroundImageDetail struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"background_image_detail"`
			Badge struct {
				ID              int    `json:"id"`
				Name            string `json:"name"`
				Title           string `json:"title"`
				Description     string `json:"description"`
				FullImageDetail struct {
					Path     string `json:"path"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Viewable bool   `json:"viewable"`
				} `json:"full_image_detail"`
				ScaledImageDetail struct {
					Path     string `json:"path"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Viewable bool   `json:"viewable"`
				} `json:"scaled_image_detail"`
			} `json:"badge"`
			Relationship struct {
				Subscribed    bool `json:"subscribed"`
				SubscribedBy  bool `json:"subscribed_by"`
				Blacklisted   bool `json:"blacklisted"`
				BlacklistedBy bool `json:"blacklisted_by"`
			} `json:"relationship"`
			ShareCount   int `json:"share_count"`
			Recognitions []struct {
				Category    string `json:"category"`
				Description string `json:"description"`
			} `json:"recognitions"`
			Liked        bool `json:"liked"`
			LikeCount    int  `json:"like_count"`
			CommentCount int  `json:"comment_count"`
		} `json:"author"`
		Images []struct {
			Full struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"full"`
			Scaled struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"scaled"`
		} `json:"images"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Plaintext        string    `json:"plaintext"`
		Status           string    `json:"status"`
		StatusMessage    string    `json:"status_message"`
		Shareable        bool      `json:"shareable"`
		ShareCount       int       `json:"share_count"`
		Category         string    `json:"category"`
		Fold             bool      `json:"fold"`
		Liked            bool      `json:"liked"`
		LikeCount        int       `json:"like_count"`
		Disliked         bool      `json:"disliked"`
		DislikeCount     int       `json:"dislike_count"`
		ViewCount        int       `json:"view_count"`
		CommentCount     int       `json:"comment_count"`
		SelectedComments []struct {
			ID             int    `json:"id"`
			DefinitionID   int    `json:"definition_id"`
			EntityID       int    `json:"entity_id"`
			EntityCategory string `json:"entity_category"`
			ReplyToID      int    `json:"reply_to_id"`
			Content        string `json:"content"`
			User           struct {
				ID           int    `json:"id"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Role         string `json:"role"`
				AvatarDetail struct {
					Path     string `json:"path"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Viewable bool   `json:"viewable"`
				} `json:"avatar_detail"`
				ScaledAvatarDetail struct {
					Path     string `json:"path"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Viewable bool   `json:"viewable"`
				} `json:"scaled_avatar_detail"`
				BackgroundImageDetail struct {
					Path     string `json:"path"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Viewable bool   `json:"viewable"`
				} `json:"background_image_detail"`
				Badge struct {
					ID              int    `json:"id"`
					Name            string `json:"name"`
					Title           string `json:"title"`
					Description     string `json:"description"`
					FullImageDetail struct {
						Path     string `json:"path"`
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Viewable bool   `json:"viewable"`
					} `json:"full_image_detail"`
					ScaledImageDetail struct {
						Path     string `json:"path"`
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Viewable bool   `json:"viewable"`
					} `json:"scaled_image_detail"`
				} `json:"badge"`
				Relationship struct {
					Subscribed    bool `json:"subscribed"`
					SubscribedBy  bool `json:"subscribed_by"`
					Blacklisted   bool `json:"blacklisted"`
					BlacklistedBy bool `json:"blacklisted_by"`
				} `json:"relationship"`
				ShareCount   int           `json:"share_count"`
				Recognitions []interface{} `json:"recognitions"`
				Liked        bool          `json:"liked"`
				LikeCount    int           `json:"like_count"`
				CommentCount int           `json:"comment_count"`
			} `json:"user"`
			Status          string        `json:"status"`
			Category        string        `json:"category"`
			CreatedAt       time.Time     `json:"created_at"`
			UpdatedAt       time.Time     `json:"updated_at"`
			Liked           bool          `json:"liked"`
			LikeCount       int           `json:"like_count"`
			Plaintext       string        `json:"plaintext"`
			LinkTerms       []interface{} `json:"link_terms"`
			AtUsers         []interface{} `json:"at_users"`
			FullImageDetail struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"full_image_detail"`
			ScaledImageDetail struct {
				Path     string `json:"path"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Viewable bool   `json:"viewable"`
			} `json:"scaled_image_detail"`
		} `json:"selected_comments"`
		DefinitionCount int `json:"definition_count"`
		AnswerCount     int `json:"answer_count"`
		LinkTerms       []struct {
			ID     int    `json:"id"`
			Title  string `json:"title"`
			Status string `json:"status"`
		} `json:"link_terms"`
		AtUsers         []interface{} `json:"at_users"`
		AlbumCount      int           `json:"album_count"`
		OwnedAlbumCount int           `json:"owned_album_count"`
		Alerts          []interface{} `json:"alerts"`
		References      []interface{} `json:"references"`
		FlipSide        struct {
			DefinitionID int    `json:"definition_id"`
			FlipSideID   int    `json:"flip_side_id"`
			Message      string `json:"message"`
		} `json:"flip_side"`
		Topics []interface{} `json:"topics"`
	} `json:"data"`
	Size        int `json:"size"`
	From        int `json:"from"`
	To          int `json:"to"`
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	LastPage    int `json:"last_page"`
}

// Register 绑定查梗功能, 作用于群聊消息
func Register(bot *OPQBot.BotManager) {
	bot.AddEvent(
		OPQBot.EventNameOnGroupMessage,
		func(botQQ int64, packet *OPQBot.GroupMsgPack) {
			if packet.FromUserID == botQQ {
				return
			}
			matchStrings := regexp.MustCompile(
				`[查|问|这|这个]{0,}(.*?)[是|叫|又是]{0,}[啥|什么|啥子]{1,}梗`,
			).FindStringSubmatch(packet.Content)
			if len(matchStrings) > 1 {
				word := matchStrings[1]
				if resp, err := requests.PostJson("https://api.jikipedia.com/go/search_definitions", map[string]interface{}{"page": 1, "phrase": word}); err == nil {
					var result *SearchResult
					if err = resp.Json(&result); err == nil {
						if len(result.Data) > 0 {
							entry := result.Data[0]
							var tagNames []string
							for _, tag := range entry.Tags {
								tagNames = append(tagNames, tag.Name)
							}
							msg := fmt.Sprintf(
								"是在查梗吗？%s\n\n标题：%s\n\n标签：%s\n\n正文: %s",
								word,
								entry.Term.Title,
								strings.Join(tagNames, "、"),
								entry.Plaintext,
							)
							if len(entry.Images) > 0 {
								bot.Send(
									OPQBot.SendMsgPack{
										SendToType: OPQBot.SendToTypeGroup,
										ToUserUid:  packet.FromGroupID,
										Content: OPQBot.SendTypePicMsgByUrlContent{
											Content: msg,
											PicUrl:  entry.Images[0].Full.Path,
										},
									},
								)
							} else {
								bot.SendGroupTextMsg(packet.FromGroupID, msg)
							}
						}
					}
				}
			}
		},
	)
}

