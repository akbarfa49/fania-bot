package tiktok

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/RambIing/urlValues"
	"github.com/rs/zerolog/log"
)

type FollowingLive struct {
	Data []struct {
		AnchorRelationType int         `json:"anchor_relation_type"`
		Data               WebcastData `json:"data"`
		// DebugInfo     string `json:"debug_info"`
		// Deprecated1   []any  `json:"deprecated1"`
		// Deprecated2   string `json:"deprecated2"`
		// Deprecated6   []any  `json:"deprecated6"`
		// DrawerGameTag string `json:"drawer_game_tag"`
		// FlareInfo     struct {
		// 	IsFlare bool   `json:"is_flare"`
		// 	TaskID  string `json:"task_id"`
		// } `json:"flare_info"`
		// IsFresh         bool   `json:"is_fresh"`
		// IsPseudoLiving  bool   `json:"is_pseudo_living"`
		// IsRecommendCard bool   `json:"is_recommend_card"`
		LiveReason string `json:"live_reason"`
		Rid        string `json:"rid"`
		Type       int    `json:"type"`
	} `json:"data"`
	Extra struct {
		Banner struct {
			Banners     []any  `json:"banners"`
			BannersType int    `json:"banners_type"`
			SwitchType  int    `json:"switch_type"`
			Title       string `json:"title"`
			Total       int    `json:"total"`
		} `json:"banner"`
		Cost        int    `json:"cost"`
		HasMore     bool   `json:"has_more"`
		HashtagText string `json:"hashtag_text"`
		IsBackup    int    `json:"is_backup"`
		LogPb       struct {
			ImprID    string `json:"impr_id"`
			SessionID int64  `json:"session_id"`
		} `json:"log_pb"`
		MaxTime            int64  `json:"max_time"`
		MinTime            int    `json:"min_time"`
		NoResultReason     string `json:"no_result_reason"`
		Now                int64  `json:"now"`
		ShowFollowingPopup bool   `json:"show_following_popup"`
		Style              int    `json:"style"`
		Total              int    `json:"total"`
		UnreadExtra        string `json:"unread_extra"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

func (t *Client) CheckFollowingLive() (*FollowingLive, []byte, error) {

	req := t.r.R()
	urVal := url.Values{}
	urVal.Add("WebIdLastTime", strconv.FormatInt(time.Now().Add(-24*time.Hour).Unix(), 10))
	urVal.Add("aid", "1988")
	urVal.Set("app_language", "en")
	urVal.Set("app_name", "tiktok_web")
	urVal.Set("browser_language", "en-US")
	urVal.Set("browser_name", "Mozilla")
	urVal.Set("browser_online", "true")
	urVal.Set("browser_platform", "Win32")
	urVal.Set("browser_version", "5.0 (Windows)")
	urVal.Set("channel", "tiktok_web")
	urVal.Set("cookie_enabled", "true")
	urVal.Set("device_id", t.deviceID)
	urVal.Set("device_platform", "web_pc")
	urVal.Set("from_page", "following")
	urVal.Set("os", "windows")
	urVal.Set("priority_region", "ID")
	urVal.Set("referer", "https://www.tiktok.com/live")
	urVal.Set("region", "ID")
	urVal.Set("req_from", "tiktok_message_webapp_following")
	urVal.Set("root_referer", "https://www.tiktok.com/live")
	urVal.Set("screen_height", "1080")
	urVal.Set("screen_width", "1920")
	urVal.Set("tz_name", "Asia/Bangkok")
	urVal.Set("webcast_language", "en")
	urVal.Set("X-Bogus", "DFSzsIVOLx2ANtORtL9/QKxfrmfE")
	urVal.Set("_signature", "_02B4Z6wo00001YY33YgAAIDAxQdE6kxnr52GNdkAAAR1b6")
	res, err := req.Get("https://webcast.tiktok.com/webcast/feed/?" + urVal.Encode())
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode() > 400 {
		return nil, res.Body(), errors.New(res.Status())
	}
	out := &FollowingLive{}
	err = json.Unmarshal(res.Body(), out)
	if err != nil {
		return nil, res.Body(), err
	}
	return out, nil, nil
}

type WebcastData struct {
	AnchorABMap struct {
	} `json:"AnchorABMap"`
	AdjustDisplayOrder    int `json:"adjust_display_order"`
	AdminEcShowPermission struct {
	} `json:"admin_ec_show_permission"`
	AdminUserIds  []any `json:"admin_user_ids"`
	AgeRestricted struct {
		AgeInterval int  `json:"AgeInterval"`
		Restricted  bool `json:"restricted"`
		Source      int  `json:"source"`
	} `json:"age_restricted"`
	AllowPreviewTime  int `json:"allow_preview_time"`
	AnchorLiveProInfo struct {
		BannerStarlingKey      string `json:"banner_starling_key"`
		GamerBannerStarlingKey string `json:"gamer_banner_starling_key"`
		IsLivePro              bool   `json:"is_live_pro"`
		LiveProType            int    `json:"live_pro_type"`
		ShowBanner             bool   `json:"show_banner"`
	} `json:"anchor_live_pro_info"`
	AnchorScheduledTimeText  string `json:"anchor_scheduled_time_text"`
	AnchorShareText          string `json:"anchor_share_text"`
	AnchorTabType            int    `json:"anchor_tab_type"`
	AnsweringQuestionContent string `json:"answering_question_content"`
	AppID                    int    `json:"app_id"`
	AudioMute                int    `json:"audio_mute"`
	AutoCover                int    `json:"auto_cover"`
	BaLeadsGenInfo           struct {
		LeadsGenModel      string `json:"leads_gen_model"`
		LeadsGenPermission bool   `json:"leads_gen_permission"`
	} `json:"ba_leads_gen_info"`
	BaLinkInfo struct {
		BaLinkData       string `json:"ba_link_data"`
		BaLinkPermission int    `json:"ba_link_permission"`
	} `json:"ba_link_info"`
	BcToggleInfo struct {
		BcToggleShowInterval int    `json:"bc_toggle_show_interval"`
		BcToggleText         string `json:"bc_toggle_text"`
		EcomBcToggle         int    `json:"ecom_bc_toggle"`
	} `json:"bc_toggle_info"`
	BlurredCover struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"blurred_cover"`
	BookEndTime              int    `json:"book_end_time"`
	BookTime                 int    `json:"book_time"`
	BusinessLive             int    `json:"business_live"`
	ChallengeInfo            string `json:"challenge_info"`
	ClientVersion            int    `json:"client_version"`
	CommentHasTextEmojiEmote int    `json:"comment_has_text_emoji_emote"`
	CommentNameMode          int    `json:"comment_name_mode"`
	CommerceInfo             struct {
		CommercePermission       int    `json:"commerce_permission"`
		OecLiveEnterRoomInitData string `json:"oec_live_enter_room_init_data"`
		ProductNum               int    `json:"product_num"`
		UseAsyncLoad             bool   `json:"use_async_load"`
		UseNewPromotion          int    `json:"use_new_promotion"`
	} `json:"commerce_info"`
	CommercialContentToggle struct {
		OpenCommercialContentToggle bool `json:"open_commercial_content_toggle"`
		PromoteMyself               bool `json:"promote_myself"`
		PromoteThirdParty           bool `json:"promote_third_party"`
	} `json:"commercial_content_toggle"`
	CommonLabelList string `json:"common_label_list"`
	ContentTag      string `json:"content_tag"`
	Cover           struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"cover"`
	CoverType     int    `json:"cover_type"`
	CppVersion    int    `json:"cpp_version"`
	CreateTime    int    `json:"create_time"`
	DecoList      []any  `json:"deco_list"`
	Deprecated10  string `json:"deprecated10"`
	Deprecated11  string `json:"deprecated11"`
	Deprecated12  string `json:"deprecated12"`
	Deprecated13  string `json:"deprecated13"`
	Deprecated14  int    `json:"deprecated14"`
	Deprecated15  int    `json:"deprecated15"`
	Deprecated16  int    `json:"deprecated16"`
	Deprecated17  []any  `json:"deprecated17"`
	Deprecated18  int    `json:"deprecated18"`
	Deprecated19  string `json:"deprecated19"`
	Deprecated195 bool   `json:"deprecated195"`
	Deprecated2   string `json:"deprecated2"`
	Deprecated20  int    `json:"deprecated20"`
	Deprecated21  bool   `json:"deprecated21"`
	Deprecated22  int    `json:"deprecated22"`
	Deprecated23  string `json:"deprecated23"`
	Deprecated24  int    `json:"deprecated24"`
	Deprecated26  string `json:"deprecated26"`
	Deprecated28  string `json:"deprecated28"`
	Deprecated3   struct {
	} `json:"deprecated3"`
	Deprecated30          string `json:"deprecated30"`
	Deprecated31          bool   `json:"deprecated31"`
	Deprecated32          string `json:"deprecated32"`
	Deprecated35          int    `json:"deprecated35"`
	Deprecated36          int    `json:"deprecated36"`
	Deprecated39          string `json:"deprecated39"`
	Deprecated4           int    `json:"deprecated4"`
	Deprecated41          int    `json:"deprecated41"`
	Deprecated43          bool   `json:"deprecated43"`
	Deprecated44          int    `json:"deprecated44"`
	Deprecated5           bool   `json:"deprecated5"`
	Deprecated6           string `json:"deprecated6"`
	Deprecated7           int    `json:"deprecated7"`
	Deprecated8           string `json:"deprecated8"`
	Deprecated9           string `json:"deprecated9"`
	DisablePreloadStream  bool   `json:"disable_preload_stream"`
	DisablePreviewSubOnly int    `json:"disable_preview_sub_only"`
	DisableScreenRecord   bool   `json:"disable_screen_record"`
	DrawerTabPosition     string `json:"drawer_tab_position"`
	DropCommentGroup      int    `json:"drop_comment_group"`
	DropsInfo             struct {
		DropsListEntrance        bool `json:"drops_list_entrance"`
		EarliestGiftExpireTs     int  `json:"earliest_gift_expire_ts"`
		ShowClaimDropsGiftNotice bool `json:"show_claim_drops_gift_notice"`
	} `json:"drops_info"`
	EffectFrameUploadDemotion   int   `json:"effect_frame_upload_demotion"`
	EffectInfo                  []any `json:"effect_info"`
	EnableOptimizeSensitiveWord bool  `json:"enable_optimize_sensitive_word"`
	EnableServerDrop            int   `json:"enable_server_drop"`
	EnableStreamEncryption      bool  `json:"enable_stream_encryption"`
	ExistedCommerceGoods        bool  `json:"existed_commerce_goods"`
	FansclubMsgStyle            int   `json:"fansclub_msg_style"`
	FeedRoomLabel               struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"feed_room_label"`
	FeedRoomLabels []any `json:"feed_room_labels"`
	FilterMsgRules []struct {
		Name   string `json:"name"`
		Random struct {
			Percentage int `json:"percentage"`
		} `json:"random"`
		Rule int `json:"rule"`
	} `json:"filter_msg_rules"`
	FinishReason   int    `json:"finish_reason"`
	FinishTime     int    `json:"finish_time"`
	FinishURL      string `json:"finish_url"`
	FinishURLV2    string `json:"finish_url_v2"`
	FollowMsgStyle int    `json:"follow_msg_style"`
	ForumExtraData string `json:"forum_extra_data"`
	GameDemo       int    `json:"game_demo"`
	GameTag        []struct {
		BundleID     string `json:"bundle_id"`
		FullName     string `json:"full_name"`
		GameCategory []any  `json:"game_category"`
		HashtagID    []int  `json:"hashtag_id"`
		HashtagList  []any  `json:"hashtag_list"`
		ID           int    `json:"id"`
		IsNewGame    bool   `json:"is_new_game"`
		Landscape    int    `json:"landscape"`
		PackageName  string `json:"package_name"`
		ShortName    string `json:"short_name"`
		ShowName     string `json:"show_name"`
	} `json:"game_tag"`
	GameTagDetail struct {
		DisplayName             string `json:"display_name"`
		GameTagID               int    `json:"game_tag_id"`
		GameTagName             string `json:"game_tag_name"`
		PreviewGameMomentEnable bool   `json:"preview_game_moment_enable"`
		StarlingKey             string `json:"starling_key"`
	} `json:"game_tag_detail"`
	GiftMsgStyle          int  `json:"gift_msg_style"`
	GiftPollVoteEnabled   bool `json:"gift_poll_vote_enabled"`
	GroupSource           int  `json:"group_source"`
	HasCommerceGoods      bool `json:"has_commerce_goods"`
	HasMoreHistoryComment bool `json:"has_more_history_comment"`
	HasUsedMusic          bool `json:"has_used_music"`
	Hashtag               struct {
		ID    int `json:"id"`
		Image struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"image"`
		Namespace int    `json:"namespace"`
		Title     string `json:"title"`
	} `json:"hashtag"`
	HaveWishlist         bool     `json:"have_wishlist"`
	HistoryCommentCursor string   `json:"history_comment_cursor"`
	HistoryCommentList   []any    `json:"history_comment_list"`
	HotSentenceInfo      string   `json:"hot_sentence_info"`
	ID                   int64    `json:"id"`
	IDStr                string   `json:"id_str"`
	IdcRegion            string   `json:"idc_region"`
	Indicators           []string `json:"indicators"`
	InteractionQuestion  struct {
		HasLightningStrengthen bool `json:"has_lightning_strengthen"`
		HasQuickAnswer         bool `json:"has_quick_answer"`
		HasRecommend           bool `json:"has_recommend"`
		QuestionAndAnswerEntry int  `json:"question_and_answer_entry"`
	} `json:"interaction_question"`
	InteractionQuestionVersion int    `json:"interaction_question_version"`
	Introduction               string `json:"introduction"`
	IsGatedRoom                bool   `json:"is_gated_room"`
	IsReplay                   bool   `json:"is_replay"`
	IsShowUserCardSwitch       bool   `json:"is_show_user_card_switch"`
	KaraokeInfo                struct {
		DisplayKaraoke     bool `json:"display_karaoke"`
		KaraokeLyricStatus bool `json:"karaoke_lyric_status"`
		KaraokeStatus      bool `json:"karaoke_status"`
	} `json:"karaoke_info"`
	LastPingTime int `json:"last_ping_time"`
	Layout       int `json:"layout"`
	LikeCount    int `json:"like_count"`
	LikeEffect   struct {
		EffectCnt        int `json:"effect_cnt"`
		EffectIntervalMs int `json:"effect_interval_ms"`
		Level            int `json:"level"`
		Version          int `json:"version"`
	} `json:"like_effect"`
	LikeIconInfo struct {
		DefaultIcons []any `json:"default_icons"`
		Icons        []any `json:"icons"`
		IconsSelf    []any `json:"icons_self"`
	} `json:"like_icon_info"`
	LikeInfo struct {
		ClickCnt int `json:"click_cnt"`
		ShowCnt  int `json:"show_cnt"`
	} `json:"like_info"`
	LinkMic struct {
		AudienceIDList []any `json:"audience_id_list"`
		BattleScores   []any `json:"battle_scores"`
		BattleSettings struct {
			BattleID    int    `json:"battle_id"`
			ChannelID   int    `json:"channel_id"`
			Duration    int    `json:"duration"`
			Finished    int    `json:"finished"`
			MatchType   int    `json:"match_type"`
			StartTime   int    `json:"start_time"`
			StartTimeMs int    `json:"start_time_ms"`
			Theme       string `json:"theme"`
		} `json:"battle_settings"`
		ChannelID   int `json:"channel_id"`
		ChannelInfo struct {
			Dimension      int `json:"dimension"`
			GroupChannelID int `json:"group_channel_id"`
			InnerChannelID int `json:"inner_channel_id"`
			Layout         int `json:"layout"`
			Vendor         int `json:"vendor"`
		} `json:"channel_info"`
		FollowedCount  int   `json:"followed_count"`
		LinkedUserList []any `json:"linked_user_list"`
		MultiLiveEnum  int   `json:"multi_live_enum"`
		RivalAnchorID  int   `json:"rival_anchor_id"`
		ShowUserList   []any `json:"show_user_list"`
	} `json:"link_mic"`
	LinkerMap struct {
	} `json:"linker_map"`
	LinkmicLayout       int    `json:"linkmic_layout"`
	LiteUserNotVisible  bool   `json:"lite_user_not_visible"`
	LiteUserVisible     bool   `json:"lite_user_visible"`
	LiveDistribution    []any  `json:"live_distribution"`
	LiveID              int    `json:"live_id"`
	LiveReason          string `json:"live_reason"`
	LiveRoomMode        int    `json:"live_room_mode"`
	LiveSubOnly         int    `json:"live_sub_only"`
	LiveSubOnlyUseMusic int    `json:"live_sub_only_use_music"`
	LiveTypeAudio       bool   `json:"live_type_audio"`
	LiveTypeLinkmic     bool   `json:"live_type_linkmic"`
	LiveTypeNormal      bool   `json:"live_type_normal"`
	LiveTypeSandbox     bool   `json:"live_type_sandbox"`
	LiveTypeScreenshot  bool   `json:"live_type_screenshot"`
	LiveTypeSocialLive  bool   `json:"live_type_social_live"`
	LiveTypeThirdParty  bool   `json:"live_type_third_party"`
	LivingRoomAttrs     struct {
		AdminFlag   int    `json:"admin_flag"`
		Rank        int    `json:"rank"`
		RoomID      int64  `json:"room_id"`
		RoomIDStr   string `json:"room_id_str"`
		SilenceFlag int    `json:"silence_flag"`
	} `json:"living_room_attrs"`
	LotteryFinishTime int    `json:"lottery_finish_time"`
	MaxPreviewTime    int    `json:"max_preview_time"`
	MosaicStatus      int    `json:"mosaic_status"`
	MultiStreamID     int    `json:"multi_stream_id"`
	MultiStreamIDStr  string `json:"multi_stream_id_str"`
	MultiStreamScene  int    `json:"multi_stream_scene"`
	MultiStreamSource int    `json:"multi_stream_source"`
	NetMode           int    `json:"net_mode"`
	OsType            int    `json:"os_type"`
	Owner             struct {
		AllowFindByContacts                 bool `json:"allow_find_by_contacts"`
		AllowOthersDownloadVideo            bool `json:"allow_others_download_video"`
		AllowOthersDownloadWhenSharingVideo bool `json:"allow_others_download_when_sharing_video"`
		AllowShareShowProfile               bool `json:"allow_share_show_profile"`
		AllowShowInGossip                   bool `json:"allow_show_in_gossip"`
		AllowShowMyAction                   bool `json:"allow_show_my_action"`
		AllowStrangeComment                 bool `json:"allow_strange_comment"`
		AllowUnfollowerComment              bool `json:"allow_unfollower_comment"`
		AllowUseLinkmic                     bool `json:"allow_use_linkmic"`
		AvatarLarge                         struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"avatar_large"`
		AvatarMedium struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"avatar_medium"`
		AvatarThumb struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"avatar_thumb"`
		BadgeImageList []any `json:"badge_image_list"`
		BadgeList      []struct {
			OpenWebURL string `json:"OpenWebURL"`
			Combine    struct {
				Background struct {
					BackgroundColorCode string `json:"background_color_code"`
					BorderColorCode     string `json:"border_color_code"`
					Image               struct {
						AvgColor   string `json:"avg_color"`
						Height     int    `json:"height"`
						ImageType  int    `json:"image_type"`
						IsAnimated bool   `json:"is_animated"`
						OpenWebURL string `json:"open_web_url"`
						URI        string `json:"uri"`
						URLList    []any  `json:"url_list"`
						Width      int    `json:"width"`
					} `json:"image"`
				} `json:"background"`
				BackgroundAutoMirrored bool `json:"background_auto_mirrored"`
				BackgroundDarkMode     struct {
					BackgroundColorCode string `json:"background_color_code"`
					BorderColorCode     string `json:"border_color_code"`
					Image               struct {
						AvgColor   string `json:"avg_color"`
						Height     int    `json:"height"`
						ImageType  int    `json:"image_type"`
						IsAnimated bool   `json:"is_animated"`
						OpenWebURL string `json:"open_web_url"`
						URI        string `json:"uri"`
						URLList    []any  `json:"url_list"`
						Width      int    `json:"width"`
					} `json:"image"`
				} `json:"background_dark_mode"`
				DisplayType int `json:"display_type"`
				FontStyle   struct {
					BorderColor string `json:"border_color"`
					FontColor   string `json:"font_color"`
					FontSize    int    `json:"font_size"`
					FontWidth   int    `json:"font_width"`
				} `json:"font_style"`
				Icon struct {
					AvgColor   string   `json:"avg_color"`
					Height     int      `json:"height"`
					ImageType  int      `json:"image_type"`
					IsAnimated bool     `json:"is_animated"`
					OpenWebURL string   `json:"open_web_url"`
					URI        string   `json:"uri"`
					URLList    []string `json:"url_list"`
					Width      int      `json:"width"`
				} `json:"icon"`
				IconAutoMirrored    bool `json:"icon_auto_mirrored"`
				MultiGuestShowStyle int  `json:"multi_guest_show_style"`
				Padding             struct {
					BadgeWidth            int  `json:"badge_width"`
					HorizontalPaddingRule int  `json:"horizontal_padding_rule"`
					IconBottomPadding     int  `json:"icon_bottom_padding"`
					IconTopPadding        int  `json:"icon_top_padding"`
					LeftPadding           int  `json:"left_padding"`
					MiddlePadding         int  `json:"middle_padding"`
					RightPadding          int  `json:"right_padding"`
					UseSpecific           bool `json:"use_specific"`
					VerticalPaddingRule   int  `json:"vertical_padding_rule"`
				} `json:"padding"`
				PaddingNewFont struct {
					BadgeWidth            int  `json:"badge_width"`
					HorizontalPaddingRule int  `json:"horizontal_padding_rule"`
					IconBottomPadding     int  `json:"icon_bottom_padding"`
					IconTopPadding        int  `json:"icon_top_padding"`
					LeftPadding           int  `json:"left_padding"`
					MiddlePadding         int  `json:"middle_padding"`
					RightPadding          int  `json:"right_padding"`
					UseSpecific           bool `json:"use_specific"`
					VerticalPaddingRule   int  `json:"vertical_padding_rule"`
				} `json:"padding_new_font"`
				PersonalCardShowStyle int `json:"personal_card_show_style"`
				ProfileCardPanel      struct {
					BadgeTextPosition int `json:"badge_text_position"`
					ProfileContent    struct {
						IconList   []any `json:"icon_list"`
						UseContent bool  `json:"use_content"`
					} `json:"profile_content"`
					ProjectionConfig struct {
						Icon struct {
							AvgColor   string `json:"avg_color"`
							Height     int    `json:"height"`
							ImageType  int    `json:"image_type"`
							IsAnimated bool   `json:"is_animated"`
							OpenWebURL string `json:"open_web_url"`
							URI        string `json:"uri"`
							URLList    []any  `json:"url_list"`
							Width      int    `json:"width"`
						} `json:"icon"`
						UseProjection bool `json:"use_projection"`
					} `json:"projection_config"`
					UseNewProfileCardStyle bool `json:"use_new_profile_card_style"`
				} `json:"profile_card_panel"`
				PublicScreenShowStyle           int    `json:"public_screen_show_style"`
				RanklistOnlineAudienceShowStyle int    `json:"ranklist_online_audience_show_style"`
				Str                             string `json:"str"`
			} `json:"combine"`
			Display           bool `json:"display"`
			DisplayStatus     int  `json:"display_status"`
			DisplayType       int  `json:"display_type"`
			ExhibitionType    int  `json:"exhibition_type"`
			GreyedByClient    int  `json:"greyed_by_client"`
			IsCustomized      bool `json:"is_customized"`
			Position          int  `json:"position"`
			PriorityType      int  `json:"priority_type"`
			PrivilegeLogExtra struct {
				DataVersion      string `json:"data_version"`
				Level            string `json:"level"`
				PrivilegeID      string `json:"privilege_id"`
				PrivilegeOrderID string `json:"privilege_order_id"`
				PrivilegeVersion string `json:"privilege_version"`
			} `json:"privilege_log_extra"`
			SceneType int `json:"scene_type"`
		} `json:"badge_list"`
		BgImgURL                 string `json:"bg_img_url"`
		BioDescription           string `json:"bio_description"`
		BlockStatus              int    `json:"block_status"`
		BorderList               []any  `json:"border_list"`
		CommentRestrict          int    `json:"comment_restrict"`
		CommerceWebcastConfigIds []any  `json:"commerce_webcast_config_ids"`
		Constellation            string `json:"constellation"`
		CreateTime               int    `json:"create_time"`
		Deprecated1              int    `json:"deprecated1"`
		Deprecated12             int    `json:"deprecated12"`
		Deprecated13             int    `json:"deprecated13"`
		Deprecated15             int    `json:"deprecated15"`
		Deprecated16             bool   `json:"deprecated16"`
		Deprecated17             bool   `json:"deprecated17"`
		Deprecated18             string `json:"deprecated18"`
		Deprecated19             bool   `json:"deprecated19"`
		Deprecated2              int    `json:"deprecated2"`
		Deprecated21             int    `json:"deprecated21"`
		Deprecated28             bool   `json:"deprecated28"`
		Deprecated29             string `json:"deprecated29"`
		Deprecated3              int    `json:"deprecated3"`
		Deprecated4              int    `json:"deprecated4"`
		Deprecated5              string `json:"deprecated5"`
		Deprecated6              int    `json:"deprecated6"`
		Deprecated7              string `json:"deprecated7"`
		Deprecated8              int    `json:"deprecated8"`
		DisableIchat             int    `json:"disable_ichat"`
		DisplayID                string `json:"display_id"`
		EnableIchatImg           int    `json:"enable_ichat_img"`
		Exp                      int    `json:"exp"`
		FanTicketCount           int    `json:"fan_ticket_count"`
		FansClubInfo             struct {
			Badge struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"badge"`
			FansClubName string `json:"fans_club_name"`
			FansCount    int    `json:"fans_count"`
			FansLevel    int    `json:"fans_level"`
			FansScore    int    `json:"fans_score"`
			IsSleeping   bool   `json:"is_sleeping"`
		} `json:"fans_club_info"`
		FoldStrangerChat bool `json:"fold_stranger_chat"`
		FollowInfo       struct {
			FollowStatus   int `json:"follow_status"`
			FollowerCount  int `json:"follower_count"`
			FollowingCount int `json:"following_count"`
			PushStatus     int `json:"push_status"`
		} `json:"follow_info"`
		FollowStatus        int     `json:"follow_status"`
		IchatRestrictType   int     `json:"ichat_restrict_type"`
		ID                  int64   `json:"id"`
		IDStr               string  `json:"id_str"`
		IsAnchorMarked      bool    `json:"is_anchor_marked"`
		IsBlock             bool    `json:"is_block"`
		IsFollower          bool    `json:"is_follower"`
		IsFollowing         bool    `json:"is_following"`
		IsSubscribe         bool    `json:"is_subscribe"`
		LinkMicStats        int     `json:"link_mic_stats"`
		MediaBadgeImageList []any   `json:"media_badge_image_list"`
		MintTypeLabel       []int64 `json:"mint_type_label"`
		ModifyTime          int     `json:"modify_time"`
		NeedProfileGuide    bool    `json:"need_profile_guide"`
		NewRealTimeIcons    []any   `json:"new_real_time_icons"`
		Nickname            string  `json:"nickname"`
		OwnRoom             struct {
			RoomIds    []int64  `json:"room_ids"`
			RoomIdsStr []string `json:"room_ids_str"`
		} `json:"own_room"`
		PayGrade struct {
			Deprecated20       int    `json:"deprecated20"`
			Deprecated22       int    `json:"deprecated22"`
			Deprecated23       int    `json:"deprecated23"`
			Deprecated24       int    `json:"deprecated24"`
			Deprecated25       int    `json:"deprecated25"`
			Deprecated26       int    `json:"deprecated26"`
			GradeBanner        string `json:"grade_banner"`
			GradeDescribe      string `json:"grade_describe"`
			GradeIconList      []any  `json:"grade_icon_list"`
			Level              int    `json:"level"`
			Name               string `json:"name"`
			NextName           string `json:"next_name"`
			NextPrivileges     string `json:"next_privileges"`
			Score              int    `json:"score"`
			ScreenChatType     int    `json:"screen_chat_type"`
			UpgradeNeedConsume int    `json:"upgrade_need_consume"`
		} `json:"pay_grade"`
		PayScore           int    `json:"pay_score"`
		PayScores          int    `json:"pay_scores"`
		PushCommentStatus  bool   `json:"push_comment_status"`
		PushDigg           bool   `json:"push_digg"`
		PushFollow         bool   `json:"push_follow"`
		PushFriendAction   bool   `json:"push_friend_action"`
		PushIchat          bool   `json:"push_ichat"`
		PushStatus         bool   `json:"push_status"`
		PushVideoPost      bool   `json:"push_video_post"`
		PushVideoRecommend bool   `json:"push_video_recommend"`
		RealTimeIcons      []any  `json:"real_time_icons"`
		ScmLabel           string `json:"scm_label"`
		SecUID             string `json:"sec_uid"`
		Secret             int    `json:"secret"`
		ShareQrcodeURI     string `json:"share_qrcode_uri"`
		SpecialID          string `json:"special_id"`
		Status             int    `json:"status"`
		SubscribeInfo      struct {
			AnchorGiftSubAuth bool `json:"anchor_gift_sub_auth"`
			Badge             struct {
				IsCustomized bool `json:"is_customized"`
			} `json:"badge"`
			EnableSubscription   bool   `json:"enable_subscription"`
			IsInGracePeriod      bool   `json:"is_in_grace_period"`
			IsSubscribe          bool   `json:"is_subscribe"`
			IsSubscribedToAnchor bool   `json:"is_subscribed_to_anchor"`
			PackageID            string `json:"package_id"`
			Qualification        bool   `json:"qualification"`
			Status               int    `json:"status"`
			SubEndTime           int    `json:"sub_end_time"`
			SubscriberCount      int    `json:"subscriber_count"`
			TimerDetail          struct {
				AnchorID             int    `json:"anchor_id"`
				AnchorSideTitle      string `json:"anchor_side_title"`
				AntidirtStatus       int    `json:"antidirt_status"`
				AuditStatus          int    `json:"audit_status"`
				LastPauseTimestampS  int64  `json:"last_pause_timestamp_s"`
				RemainingTimeS       int    `json:"remaining_time_s"`
				ScreenH              int    `json:"screen_h"`
				ScreenW              int    `json:"screen_w"`
				StartCountdownTimeS  int    `json:"start_countdown_time_s"`
				StartTimestampS      int64  `json:"start_timestamp_s"`
				StickerX             int    `json:"sticker_x"`
				StickerY             int    `json:"sticker_y"`
				SubCount             int    `json:"sub_count"`
				TimeIncreaseCapS     int    `json:"time_increase_cap_s"`
				TimeIncreasePerSubS  int    `json:"time_increase_per_sub_s"`
				TimeIncreaseReachCap bool   `json:"time_increase_reach_cap"`
				TimerID              int    `json:"timer_id"`
				TimerStatus          int    `json:"timer_status"`
				TimestampS           int    `json:"timestamp_s"`
				TotalPauseTimeS      int    `json:"total_pause_time_s"`
				TotalTimeS           int    `json:"total_time_s"`
				UserSideTitle        string `json:"user_side_title"`
			} `json:"timer_detail"`
			UserGiftSubAuth bool `json:"user_gift_sub_auth"`
		} `json:"subscribe_info"`
		TicketCount       int   `json:"ticket_count"`
		TopFans           []any `json:"top_fans"`
		TopVipNo          int   `json:"top_vip_no"`
		UpcomingEventList []any `json:"upcoming_event_list"`
		UserAttr          struct {
			AdminPermissions struct {
			} `json:"admin_permissions"`
			HasVotingFunction bool `json:"has_voting_function"`
			IsAdmin           bool `json:"is_admin"`
			IsChannelAdmin    bool `json:"is_channel_admin"`
			IsMuted           bool `json:"is_muted"`
			IsSuperAdmin      bool `json:"is_super_admin"`
			MuteDuration      int  `json:"mute_duration"`
		} `json:"user_attr"`
		UserRole                    int    `json:"user_role"`
		Verified                    bool   `json:"verified"`
		VerifiedContent             string `json:"verified_content"`
		VerifiedReason              string `json:"verified_reason"`
		WithCarManagementPermission bool   `json:"with_car_management_permission"`
		WithCommercePermission      bool   `json:"with_commerce_permission"`
		WithFusionShopEntry         bool   `json:"with_fusion_shop_entry"`
	} `json:"owner"`
	OwnerDeviceID    int64  `json:"owner_device_id"`
	OwnerDeviceIDStr string `json:"owner_device_id_str"`
	OwnerUserID      int64  `json:"owner_user_id"`
	OwnerUserIDStr   string `json:"owner_user_id_str"`
	PaidContentInfo  struct {
		PaidContentLiveData   string `json:"paid_content_live_data"`
		PaidContentPermission bool   `json:"paid_content_permission"`
	} `json:"paid_content_info"`
	PaidEvent struct {
		EventID  int `json:"event_id"`
		PaidType int `json:"paid_type"`
	} `json:"paid_event"`
	PartnershipInfo struct {
		PartnershipRoom  bool   `json:"partnership_room"`
		PromotingDropsID string `json:"promoting_drops_id"`
		PromotingGameID  string `json:"promoting_game_id"`
		PromotingRoom    bool   `json:"promoting_room"`
		PromotingTaskID  string `json:"promoting_task_id"`
		ShowTaskID       string `json:"show_task_id"`
		TaskIDList       []any  `json:"task_id_list"`
	} `json:"partnership_info"`
	PicoLiveType int `json:"pico_live_type"`
	PinInfo      struct {
		DisplayDuration int  `json:"display_duration"`
		HasPin          bool `json:"has_pin"`
	} `json:"pin_info"`
	PollConf struct {
		GiftPollLimit struct {
			CurrentPollCount int `json:"current_poll_count"`
			MaxPollCount     int `json:"max_poll_count"`
		} `json:"gift_poll_limit"`
		UseNewGiftPoll bool `json:"use_new_gift_poll"`
	} `json:"poll_conf"`
	PollingStarComment bool `json:"polling_star_comment"`
	PreEnterTime       int  `json:"pre_enter_time"`
	PreviewFlowTag     int  `json:"preview_flow_tag"`
	QuotaConfig        struct {
	} `json:"quota_config"`
	RankCommentGroups    []string `json:"rank_comment_groups"`
	RanklistAudienceType int      `json:"ranklist_audience_type"`
	RectangleCoverImg    struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"rectangle_cover_img"`
	RegionalRestricted struct {
		BlockList []any `json:"block_list"`
	} `json:"regional_restricted"`
	RelationTag     string `json:"relation_tag"`
	Replay          bool   `json:"replay"`
	Reposted        bool   `json:"reposted"`
	RoomAuditStatus int    `json:"room_audit_status"`
	RoomAuth        struct {
		Banner                   int  `json:"Banner"`
		BroadcastMessage         int  `json:"BroadcastMessage"`
		Chat                     bool `json:"Chat"`
		ChatL2                   bool `json:"ChatL2"`
		ChatSubOnly              bool `json:"ChatSubOnly"`
		CommercePermission       int  `json:"CommercePermission"`
		CommunityFlagged         bool `json:"CommunityFlagged"`
		CommunityFlaggedReview   bool `json:"CommunityFlaggedReview"`
		CustomizableGiftPoll     int  `json:"CustomizableGiftPoll"`
		CustomizablePoll         int  `json:"CustomizablePoll"`
		Danmaku                  bool `json:"Danmaku"`
		Digg                     bool `json:"Digg"`
		DonationSticker          int  `json:"DonationSticker"`
		EmotePoll                int  `json:"EmotePoll"`
		EnableFansLevel          bool `json:"EnableFansLevel"`
		EventPromotion           int  `json:"EventPromotion"`
		Explore                  bool `json:"Explore"`
		GameRankingSwitch        int  `json:"GameRankingSwitch"`
		Gift                     bool `json:"Gift"`
		GiftAnchorMt             int  `json:"GiftAnchorMt"`
		GiftPoll                 int  `json:"GiftPoll"`
		GoldenEnvelope           int  `json:"GoldenEnvelope"`
		GoldenEnvelopeActivity   int  `json:"GoldenEnvelopeActivity"`
		InteractionQuestion      bool `json:"InteractionQuestion"`
		Landscape                int  `json:"Landscape"`
		LandscapeChat            int  `json:"LandscapeChat"`
		LuckMoney                bool `json:"LuckMoney"`
		MultiEnableReserve       bool `json:"MultiEnableReserve"`
		Pictionary               int  `json:"Pictionary"`
		PictionaryBubble         int  `json:"PictionaryBubble"`
		PictionaryPermission     int  `json:"PictionaryPermission"`
		Poll                     int  `json:"Poll"`
		Promote                  bool `json:"Promote"`
		PromoteOther             int  `json:"PromoteOther"`
		Props                    bool `json:"Props"`
		PublicScreen             int  `json:"PublicScreen"`
		QuickChat                int  `json:"QuickChat"`
		Rank                     int  `json:"Rank"`
		RankingChangeAlterSwitch int  `json:"RankingChangeAlterSwitch"`
		RoomContributor          bool `json:"RoomContributor"`
		SecretRoom               int  `json:"SecretRoom"`
		Share                    bool `json:"Share"`
		ShareEffect              int  `json:"ShareEffect"`
		ShoppingRanking          int  `json:"ShoppingRanking"`
		SpamComments             bool `json:"SpamComments"`
		UserCard                 bool `json:"UserCard"`
		UserCount                int  `json:"UserCount"`
		Viewers                  bool `json:"Viewers"`
		AnchorLevelPermission    struct {
			InteractionQuestion int `json:"InteractionQuestion"`
			Beauty              int `json:"beauty"`
			CommentFilter       int `json:"comment_filter"`
			CommentSetting      int `json:"comment_setting"`
			CustomizablePoll    int `json:"customizable_poll"`
			DonationSticker     int `json:"donation_sticker"`
			Effects             int `json:"effects"`
			Flip                int `json:"flip"`
			FullScreenMode      int `json:"full_screen_mode"`
			GoodyBag            int `json:"goody_bag"`
			HearYourOwnVoice    int `json:"hear_your_own_voice"`
			Karaoke             int `json:"karaoke"`
			LiveBackground      int `json:"live_background"`
			LiveCenter          int `json:"live_center"`
			LiveIntro           int `json:"live_intro"`
			Mirror              int `json:"mirror"`
			ModeratorSetting    int `json:"moderator_setting"`
			PauseLive           int `json:"pause_live"`
			Pictionary          int `json:"pictionary"`
			Pin                 int `json:"pin"`
			PlayTogether        int `json:"play_together"`
			Poll                int `json:"poll"`
			Portal              int `json:"portal"`
			Promote             int `json:"promote"`
			Share               int `json:"share"`
			Sticker             int `json:"sticker"`
			Topic               int `json:"topic"`
			TreasureBox         int `json:"treasure_box"`
			ViewerRankList      int `json:"viewer_rank_list"`
			VoiceEffect         int `json:"voice_effect"`
		} `json:"anchor_level_permission"`
		CommentTrayStatus           int   `json:"comment_tray_status"`
		CreditEntranceForAudience   bool  `json:"credit_entrance_for_audience"`
		Deprecated1                 bool  `json:"deprecated1"`
		Deprecated118               []any `json:"deprecated118"`
		Deprecated119               int   `json:"deprecated119"`
		Deprecated2                 int   `json:"deprecated2"`
		Deprecated3                 int   `json:"deprecated3"`
		Deprecated4                 int   `json:"deprecated4"`
		Deprecated5                 int   `json:"deprecated5"`
		Deprecated6                 int   `json:"deprecated6"`
		Deprecated7                 int   `json:"deprecated7"`
		Deprecated8                 int   `json:"deprecated8"`
		Deprecated9                 int   `json:"deprecated9"`
		GameGuessPermission         bool  `json:"game_guess_permission"`
		GuessEntranceForHost        bool  `json:"guess_entrance_for_host"`
		ShowCreditWidget            bool  `json:"show_credit_widget"`
		StarCommentPermissionSwitch struct {
			OffReason string `json:"OffReason"`
			Status    int    `json:"status"`
		} `json:"star_comment_permission_switch"`
		TransactionHistory int  `json:"transaction_history"`
		UseUserPv          bool `json:"use_user_pv"`
	} `json:"room_auth"`
	RoomCreateAbParam string `json:"room_create_ab_param"`
	RoomLayout        int    `json:"room_layout"`
	RoomPcu           int    `json:"room_pcu"`
	RoomStickerList   []any  `json:"room_sticker_list"`
	RoomTabs          []any  `json:"room_tabs"`
	RoomTag           int    `json:"room_tag"`
	RtcAppID          string `json:"rtc_app_id"`
	ScrollConfig      string `json:"scroll_config"`
	SearchID          int    `json:"search_id"`
	ShareMsgStyle     int    `json:"share_msg_style"`
	ShareShowTime     struct {
		ShowTimeOnEnter int `json:"show_time_on_enter"`
		ShowTimeOnShare int `json:"show_time_on_share"`
	} `json:"share_show_time"`
	ShareURL                string `json:"share_url"`
	ShortTitle              string `json:"short_title"`
	ShortTouchItems         []any  `json:"short_touch_items"`
	ShowStarCommentEntrance bool   `json:"show_star_comment_entrance"`
	SocialInteraction       struct {
		Cohost struct {
			LinkedUsers           []any `json:"linked_users"`
			MultiCohostPermission bool  `json:"multi_cohost_permission"`
		} `json:"cohost"`
		LinkmicSceneLinker struct {
			Num2 int64 `json:"2"`
		} `json:"linkmic_scene_linker"`
		MultiLive struct {
			AnchorSettingInfo struct {
				LastLayoutSettings []any `json:"last_layout_settings"`
			} `json:"anchor_setting_info"`
			AudienceSharedInviteePanelType int `json:"audience_shared_invitee_panel_type"`
			HostGifterLinkmicEnum          int `json:"host_gifter_linkmic_enum"`
			HostMultiGuestDevMode          int `json:"host_multi_guest_dev_mode"`
			LinkmicServiceVersion          int `json:"linkmic_service_version"`
			RoomMultiGuestLinkmicInfo      struct {
				LinkmicRoomCreateAbParam string `json:"linkmic_room_create_ab_param"`
				PackErrCode              int    `json:"pack_err_code"`
			} `json:"room_multi_guest_linkmic_info"`
			TryOpenMultiGuestWhenCreateRoom bool `json:"try_open_multi_guest_when_create_room"`
			UserSettings                    struct {
				ApplierSortGiftScoreThreshold       int `json:"applier_sort_gift_score_threshold"`
				ApplierSortSetting                  int `json:"applier_sort_setting"`
				MultiGuestAllowRequestFromFollowers int `json:"multi_guest_allow_request_from_followers"`
				MultiGuestAllowRequestFromFriends   int `json:"multi_guest_allow_request_from_friends"`
				MultiGuestAllowRequestFromOthers    int `json:"multi_guest_allow_request_from_others"`
				MultiLiveApplyPermission            int `json:"multi_live_apply_permission"`
			} `json:"user_settings"`
			ViewerGifterLinkmicEnum int `json:"viewer_gifter_linkmic_enum"`
		} `json:"multi_live"`
	} `json:"social_interaction"`
	SquareCoverImg struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"square_cover_img"`
	StarCommentConfig struct {
		DisplayLock              bool `json:"display_lock"`
		GrantGroup               int  `json:"grant_group"`
		GrantLevel               int  `json:"grant_level"`
		StarCommentQualification bool `json:"star_comment_qualification"`
		StarCommentSwitch        bool `json:"star_comment_switch"`
	} `json:"star_comment_config"`
	StartTime int `json:"start_time"`
	Stats     struct {
		Deprecated1          int    `json:"deprecated1"`
		Deprecated2          string `json:"deprecated2"`
		DiggCount            int    `json:"digg_count"`
		EnterCount           int    `json:"enter_count"`
		FanTicket            int    `json:"fan_ticket"`
		FollowCount          int    `json:"follow_count"`
		GiftUvCount          int    `json:"gift_uv_count"`
		ID                   int64  `json:"id"`
		IDStr                string `json:"id_str"`
		LikeCount            int    `json:"like_count"`
		ReplayFanTicket      int    `json:"replay_fan_ticket"`
		ReplayViewers        int    `json:"replay_viewers"`
		ShareCount           int    `json:"share_count"`
		TotalUser            int    `json:"total_user"`
		TotalUserDesp        string `json:"total_user_desp"`
		UserCountComposition struct {
			Deprecated1 int     `json:"deprecated1"`
			MyFollow    float64 `json:"my_follow"`
			Other       float64 `json:"other"`
			VideoDetail int     `json:"video_detail"`
		} `json:"user_count_composition"`
		Watermelon int `json:"watermelon"`
	} `json:"stats"`
	Status       int    `json:"status"`
	StickerList  []any  `json:"sticker_list"`
	StreamID     int64  `json:"stream_id"`
	StreamIDStr  string `json:"stream_id_str"`
	StreamStatus int    `json:"stream_status"`
	StreamURL    struct {
		CandidateResolution []string `json:"candidate_resolution"`
		CompletePushUrls    []any    `json:"complete_push_urls"`
		DefaultResolution   string   `json:"default_resolution"`
		Extra               struct {
			AnchorInteractProfile   int  `json:"anchor_interact_profile"`
			AudienceInteractProfile int  `json:"audience_interact_profile"`
			BframeEnable            bool `json:"bframe_enable"`
			BitrateAdaptStrategy    int  `json:"bitrate_adapt_strategy"`
			Bytevc1Enable           bool `json:"bytevc1_enable"`
			DefaultBitrate          int  `json:"default_bitrate"`
			Deprecated1             bool `json:"deprecated1"`
			Fps                     int  `json:"fps"`
			GopSec                  int  `json:"gop_sec"`
			HardwareEncode          bool `json:"hardware_encode"`
			Height                  int  `json:"height"`
			MaxBitrate              int  `json:"max_bitrate"`
			MinBitrate              int  `json:"min_bitrate"`
			Roi                     bool `json:"roi"`
			SwRoi                   bool `json:"sw_roi"`
			VideoProfile            int  `json:"video_profile"`
			Width                   int  `json:"width"`
		} `json:"extra"`
		FlvPullURL struct {
			Hd1 string `json:"HD1"`
			Sd1 string `json:"SD1"`
			Sd2 string `json:"SD2"`
		} `json:"flv_pull_url"`
		FlvPullURLParams struct {
			Hd1 string `json:"HD1"`
			Sd1 string `json:"SD1"`
			Sd2 string `json:"SD2"`
		} `json:"flv_pull_url_params"`
		HlsPullURL    string `json:"hls_pull_url"`
		HlsPullURLMap struct {
		} `json:"hls_pull_url_map"`
		HlsPullURLParams string `json:"hls_pull_url_params"`
		ID               int64  `json:"id"`
		IDStr            string `json:"id_str"`
		LiveCoreSdkData  struct {
			PullData struct {
				Options struct {
					DefaultQuality struct {
						IconType   int    `json:"icon_type"`
						Level      int    `json:"level"`
						Name       string `json:"name"`
						Resolution string `json:"resolution"`
						SdkKey     string `json:"sdk_key"`
						VCodec     string `json:"v_codec"`
					} `json:"default_quality"`
					Qualities []struct {
						IconType   int    `json:"icon_type"`
						Level      int    `json:"level"`
						Name       string `json:"name"`
						Resolution string `json:"resolution"`
						SdkKey     string `json:"sdk_key"`
						VCodec     string `json:"v_codec"`
					} `json:"qualities"`
					ShowQualityButton bool `json:"show_quality_button"`
				} `json:"options"`
				StreamData string `json:"stream_data"`
			} `json:"pull_data"`
		} `json:"live_core_sdk_data"`
		Provider       int    `json:"provider"`
		PushResolution string `json:"push_resolution"`
		PushUrls       []any  `json:"push_urls"`
		ResolutionName struct {
			Auto             string `json:"AUTO"`
			FullHd1          string `json:"FULL_HD1"`
			Hd1              string `json:"HD1"`
			Origion          string `json:"ORIGION"`
			Sd1              string `json:"SD1"`
			Sd2              string `json:"SD2"`
			PmMtVideo1080P60 string `json:"pm_mt_video_1080p60"`
			PmMtVideo720P60  string `json:"pm_mt_video_720p60"`
		} `json:"resolution_name"`
		RtmpPullURL       string `json:"rtmp_pull_url"`
		RtmpPullURLParams string `json:"rtmp_pull_url_params"`
		RtmpPushURL       string `json:"rtmp_push_url"`
		RtmpPushURLParams string `json:"rtmp_push_url_params"`
		StreamAppID       int    `json:"stream_app_id"`
		StreamControlType int    `json:"stream_control_type"`
		StreamDelayMs     int    `json:"stream_delay_ms"`
		VrType            int    `json:"vr_type"`
	} `json:"stream_url"`
	SupportQuiz int    `json:"support_quiz"`
	Title       string `json:"title"`
	TopFans     []struct {
		FanTicket int `json:"fan_ticket"`
		User      struct {
			AllowFindByContacts                 bool `json:"allow_find_by_contacts"`
			AllowOthersDownloadVideo            bool `json:"allow_others_download_video"`
			AllowOthersDownloadWhenSharingVideo bool `json:"allow_others_download_when_sharing_video"`
			AllowShareShowProfile               bool `json:"allow_share_show_profile"`
			AllowShowInGossip                   bool `json:"allow_show_in_gossip"`
			AllowShowMyAction                   bool `json:"allow_show_my_action"`
			AllowStrangeComment                 bool `json:"allow_strange_comment"`
			AllowUnfollowerComment              bool `json:"allow_unfollower_comment"`
			AllowUseLinkmic                     bool `json:"allow_use_linkmic"`
			AvatarLarge                         struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"avatar_large"`
			AvatarMedium struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"avatar_medium"`
			AvatarThumb struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"avatar_thumb"`
			BadgeImageList []any `json:"badge_image_list"`
			BadgeList      []struct {
				OpenWebURL string `json:"OpenWebURL"`
				Combine    struct {
					Background struct {
						BackgroundColorCode string `json:"background_color_code"`
						BorderColorCode     string `json:"border_color_code"`
						Image               struct {
							AvgColor   string `json:"avg_color"`
							Height     int    `json:"height"`
							ImageType  int    `json:"image_type"`
							IsAnimated bool   `json:"is_animated"`
							OpenWebURL string `json:"open_web_url"`
							URI        string `json:"uri"`
							URLList    []any  `json:"url_list"`
							Width      int    `json:"width"`
						} `json:"image"`
					} `json:"background"`
					BackgroundAutoMirrored bool `json:"background_auto_mirrored"`
					BackgroundDarkMode     struct {
						BackgroundColorCode string `json:"background_color_code"`
						BorderColorCode     string `json:"border_color_code"`
						Image               struct {
							AvgColor   string `json:"avg_color"`
							Height     int    `json:"height"`
							ImageType  int    `json:"image_type"`
							IsAnimated bool   `json:"is_animated"`
							OpenWebURL string `json:"open_web_url"`
							URI        string `json:"uri"`
							URLList    []any  `json:"url_list"`
							Width      int    `json:"width"`
						} `json:"image"`
					} `json:"background_dark_mode"`
					DisplayType int `json:"display_type"`
					FontStyle   struct {
						BorderColor string `json:"border_color"`
						FontColor   string `json:"font_color"`
						FontSize    int    `json:"font_size"`
						FontWidth   int    `json:"font_width"`
					} `json:"font_style"`
					Icon struct {
						AvgColor   string   `json:"avg_color"`
						Height     int      `json:"height"`
						ImageType  int      `json:"image_type"`
						IsAnimated bool     `json:"is_animated"`
						OpenWebURL string   `json:"open_web_url"`
						URI        string   `json:"uri"`
						URLList    []string `json:"url_list"`
						Width      int      `json:"width"`
					} `json:"icon"`
					IconAutoMirrored    bool `json:"icon_auto_mirrored"`
					MultiGuestShowStyle int  `json:"multi_guest_show_style"`
					Padding             struct {
						BadgeWidth            int  `json:"badge_width"`
						HorizontalPaddingRule int  `json:"horizontal_padding_rule"`
						IconBottomPadding     int  `json:"icon_bottom_padding"`
						IconTopPadding        int  `json:"icon_top_padding"`
						LeftPadding           int  `json:"left_padding"`
						MiddlePadding         int  `json:"middle_padding"`
						RightPadding          int  `json:"right_padding"`
						UseSpecific           bool `json:"use_specific"`
						VerticalPaddingRule   int  `json:"vertical_padding_rule"`
					} `json:"padding"`
					PaddingNewFont struct {
						BadgeWidth            int  `json:"badge_width"`
						HorizontalPaddingRule int  `json:"horizontal_padding_rule"`
						IconBottomPadding     int  `json:"icon_bottom_padding"`
						IconTopPadding        int  `json:"icon_top_padding"`
						LeftPadding           int  `json:"left_padding"`
						MiddlePadding         int  `json:"middle_padding"`
						RightPadding          int  `json:"right_padding"`
						UseSpecific           bool `json:"use_specific"`
						VerticalPaddingRule   int  `json:"vertical_padding_rule"`
					} `json:"padding_new_font"`
					PersonalCardShowStyle int `json:"personal_card_show_style"`
					ProfileCardPanel      struct {
						BadgeTextPosition int `json:"badge_text_position"`
						ProfileContent    struct {
							IconList   []any `json:"icon_list"`
							UseContent bool  `json:"use_content"`
						} `json:"profile_content"`
						ProjectionConfig struct {
							Icon struct {
								AvgColor   string `json:"avg_color"`
								Height     int    `json:"height"`
								ImageType  int    `json:"image_type"`
								IsAnimated bool   `json:"is_animated"`
								OpenWebURL string `json:"open_web_url"`
								URI        string `json:"uri"`
								URLList    []any  `json:"url_list"`
								Width      int    `json:"width"`
							} `json:"icon"`
							UseProjection bool `json:"use_projection"`
						} `json:"projection_config"`
						UseNewProfileCardStyle bool `json:"use_new_profile_card_style"`
					} `json:"profile_card_panel"`
					PublicScreenShowStyle           int    `json:"public_screen_show_style"`
					RanklistOnlineAudienceShowStyle int    `json:"ranklist_online_audience_show_style"`
					Str                             string `json:"str"`
				} `json:"combine"`
				Display           bool `json:"display"`
				DisplayStatus     int  `json:"display_status"`
				DisplayType       int  `json:"display_type"`
				ExhibitionType    int  `json:"exhibition_type"`
				GreyedByClient    int  `json:"greyed_by_client"`
				IsCustomized      bool `json:"is_customized"`
				Position          int  `json:"position"`
				PriorityType      int  `json:"priority_type"`
				PrivilegeLogExtra struct {
					DataVersion      string `json:"data_version"`
					Level            string `json:"level"`
					PrivilegeID      string `json:"privilege_id"`
					PrivilegeOrderID string `json:"privilege_order_id"`
					PrivilegeVersion string `json:"privilege_version"`
				} `json:"privilege_log_extra"`
				SceneType int `json:"scene_type"`
			} `json:"badge_list"`
			BgImgURL                 string `json:"bg_img_url"`
			BioDescription           string `json:"bio_description"`
			BlockStatus              int    `json:"block_status"`
			BorderList               []any  `json:"border_list"`
			CommentRestrict          int    `json:"comment_restrict"`
			CommerceWebcastConfigIds []any  `json:"commerce_webcast_config_ids"`
			Constellation            string `json:"constellation"`
			CreateTime               int    `json:"create_time"`
			Deprecated1              int    `json:"deprecated1"`
			Deprecated12             int    `json:"deprecated12"`
			Deprecated13             int    `json:"deprecated13"`
			Deprecated15             int    `json:"deprecated15"`
			Deprecated16             bool   `json:"deprecated16"`
			Deprecated17             bool   `json:"deprecated17"`
			Deprecated18             string `json:"deprecated18"`
			Deprecated19             bool   `json:"deprecated19"`
			Deprecated2              int    `json:"deprecated2"`
			Deprecated21             int    `json:"deprecated21"`
			Deprecated28             bool   `json:"deprecated28"`
			Deprecated29             string `json:"deprecated29"`
			Deprecated3              int    `json:"deprecated3"`
			Deprecated4              int    `json:"deprecated4"`
			Deprecated5              string `json:"deprecated5"`
			Deprecated6              int    `json:"deprecated6"`
			Deprecated7              string `json:"deprecated7"`
			Deprecated8              int    `json:"deprecated8"`
			DisableIchat             int    `json:"disable_ichat"`
			DisplayID                string `json:"display_id"`
			EnableIchatImg           int    `json:"enable_ichat_img"`
			Exp                      int    `json:"exp"`
			FanTicketCount           int    `json:"fan_ticket_count"`
			FoldStrangerChat         bool   `json:"fold_stranger_chat"`
			FollowInfo               struct {
				FollowStatus   int `json:"follow_status"`
				FollowerCount  int `json:"follower_count"`
				FollowingCount int `json:"following_count"`
				PushStatus     int `json:"push_status"`
			} `json:"follow_info"`
			FollowStatus        int    `json:"follow_status"`
			IchatRestrictType   int    `json:"ichat_restrict_type"`
			ID                  int64  `json:"id"`
			IDStr               string `json:"id_str"`
			IsAnchorMarked      bool   `json:"is_anchor_marked"`
			IsBlock             bool   `json:"is_block"`
			IsFollower          bool   `json:"is_follower"`
			IsFollowing         bool   `json:"is_following"`
			IsSubscribe         bool   `json:"is_subscribe"`
			LinkMicStats        int    `json:"link_mic_stats"`
			MediaBadgeImageList []any  `json:"media_badge_image_list"`
			MintTypeLabel       []any  `json:"mint_type_label"`
			ModifyTime          int    `json:"modify_time"`
			NeedProfileGuide    bool   `json:"need_profile_guide"`
			NewRealTimeIcons    []any  `json:"new_real_time_icons"`
			Nickname            string `json:"nickname"`
			PayGrade            struct {
				Deprecated20       int    `json:"deprecated20"`
				Deprecated22       int    `json:"deprecated22"`
				Deprecated23       int    `json:"deprecated23"`
				Deprecated24       int    `json:"deprecated24"`
				Deprecated25       int    `json:"deprecated25"`
				Deprecated26       int    `json:"deprecated26"`
				GradeBanner        string `json:"grade_banner"`
				GradeDescribe      string `json:"grade_describe"`
				GradeIconList      []any  `json:"grade_icon_list"`
				Level              int    `json:"level"`
				Name               string `json:"name"`
				NextName           string `json:"next_name"`
				NextPrivileges     string `json:"next_privileges"`
				Score              int    `json:"score"`
				ScreenChatType     int    `json:"screen_chat_type"`
				UpgradeNeedConsume int    `json:"upgrade_need_consume"`
			} `json:"pay_grade"`
			PayScore           int    `json:"pay_score"`
			PayScores          int    `json:"pay_scores"`
			PushCommentStatus  bool   `json:"push_comment_status"`
			PushDigg           bool   `json:"push_digg"`
			PushFollow         bool   `json:"push_follow"`
			PushFriendAction   bool   `json:"push_friend_action"`
			PushIchat          bool   `json:"push_ichat"`
			PushStatus         bool   `json:"push_status"`
			PushVideoPost      bool   `json:"push_video_post"`
			PushVideoRecommend bool   `json:"push_video_recommend"`
			RealTimeIcons      []any  `json:"real_time_icons"`
			ScmLabel           string `json:"scm_label"`
			SecUID             string `json:"sec_uid"`
			Secret             int    `json:"secret"`
			ShareQrcodeURI     string `json:"share_qrcode_uri"`
			SpecialID          string `json:"special_id"`
			Status             int    `json:"status"`
			TicketCount        int    `json:"ticket_count"`
			TopFans            []any  `json:"top_fans"`
			TopVipNo           int    `json:"top_vip_no"`
			UpcomingEventList  []any  `json:"upcoming_event_list"`
			UserAttr           struct {
				AdminPermissions struct {
				} `json:"admin_permissions"`
				HasVotingFunction bool `json:"has_voting_function"`
				IsAdmin           bool `json:"is_admin"`
				IsChannelAdmin    bool `json:"is_channel_admin"`
				IsMuted           bool `json:"is_muted"`
				IsSuperAdmin      bool `json:"is_super_admin"`
				MuteDuration      int  `json:"mute_duration"`
			} `json:"user_attr"`
			UserRole                    int    `json:"user_role"`
			Verified                    bool   `json:"verified"`
			VerifiedContent             string `json:"verified_content"`
			VerifiedReason              string `json:"verified_reason"`
			WithCarManagementPermission bool   `json:"with_car_management_permission"`
			WithCommercePermission      bool   `json:"with_commerce_permission"`
			WithFusionShopEntry         bool   `json:"with_fusion_shop_entry"`
		} `json:"user"`
	} `json:"top_fans"`
	UseFilter         bool   `json:"use_filter"`
	UserCount         int    `json:"user_count"`
	UserShareText     string `json:"user_share_text"`
	VideoFeedTag      string `json:"video_feed_tag"`
	WebcastCommentTcs int    `json:"webcast_comment_tcs"`
	WebcastSdkVersion int    `json:"webcast_sdk_version"`
	WithDrawSomething bool   `json:"with_draw_something"`
	WithKtv           bool   `json:"with_ktv"`
	WithLinkmic       bool   `json:"with_linkmic"`
}

type WebcastPreloadRoom struct {
	Data       WebcastData `json:"data"`
	StatusCode int         `json:"status_code"`
	Extra      struct {
		Reason string `json:"reason"`
	}
}

const ReasonNoWebcast = "delegate_empty"

func (t *Client) GetWebcastPreloadRoom(ctx context.Context, userID string) (WebcastPreloadRoom, []byte, error) {
	values := urlValues.Values{}

	msToken := ""
	for _, v := range t.r.Cookies {
		if v.Name == "msToken" {
			msToken = v.Value
			break
		}
	}

	// Inserting query parameters manually
	values.Add("WebIdLastTime", strconv.Itoa(int(time.Now().Add(-24*time.Hour).Unix())))
	values.Add("aid", "1988")
	values.Add("app_language", "en")
	values.Add("app_name", "tiktok_web")
	values.Add("browser_language", "en-US")
	values.Add("browser_name", "Mozilla")
	values.Add("browser_online", "true")
	values.Add("browser_platform", "Win32")
	values.Add("browser_version", "5.0 (Windows)")
	values.Add("channel", "tiktok_web")
	values.Add("cookie_enabled", "true")
	values.Add("device_id", t.deviceID)
	values.Add("device_platform", "web_pc")
	values.Add("focus_state", "true")
	values.Add("from_page", "user")
	values.Add("history_len", "4")
	values.Add("is_fullscreen", "false")
	values.Add("is_page_visible", "true")
	values.Add("os", "windows")
	values.Add("owner_user_id", userID)
	values.Add("priority_region", "")
	values.Add("referer", "")
	values.Add("region", "ID")
	values.Add("scene", "webapp_profile_preview")
	values.Add("screen_height", "1080")
	values.Add("screen_width", "1920")
	values.Add("tz_name", "Asia/Bangkok")
	values.Add("webcast_language", "en")
	values.Add("msToken", msToken)
	// Print the url.Values
	urlQuery := values.EncodeWithOrder()
	bogus, err := GenerateBogus(urlQuery, t.userAgent)
	out := WebcastPreloadRoom{}
	if err != nil {
		if t.debug {
			log.Err(err).Msg("failed to generate bogus")
		}
		return out, nil, err
	}
	values.Add("X-Bogus", bogus)
	req := t.r.R()
	uri := "https://webcast.tiktok.com/webcast/room/preload_room/?" + values.EncodeWithOrder()
	res, err := req.Get(uri)
	if err != nil {
		return out, nil, err
	}
	if res.StatusCode() > 400 {
		return out, res.Body(), errors.New(res.Status())
	}
	// fmt.Println(string(res.Body()))
	err = json.Unmarshal(res.Body(), &out)
	if err != nil {
		return out, res.Body(), err
	}
	if err := ParseStatusCode(out.StatusCode); err != nil {
		return out, nil, err
	}
	return out, nil, nil

}

type WebcastRoom struct {
	Data       WebcastData `json:"data"`
	StatusCode int         `json:"status_code"`
	Extra      struct {
		Reason string `json:"reason"`
	}
}

func (t *Client) WebcastEnterRoom(ctx context.Context, roomID string) (WebcastRoom, []byte, error) {
	values := urlValues.Values{}
	msToken := ""
	for _, v := range t.r.Cookies {
		if v.Name == "msToken" {
			msToken = v.Value
			break
		}
	}

	// Inserting query parameters manually
	values.Add("aid", "1988")
	values.Add("app_language", "en")
	values.Add("app_name", "tiktok_web")
	values.Add("browser_language", "en-US")
	values.Add("browser_name", "Mozilla")
	values.Add("browser_online", "true")
	values.Add("browser_platform", "Win32")
	values.Add("browser_version", "5.0 (Windows)")
	values.Add("channel", "tiktok_web")
	values.Add("cookie_enabled", "true")
	values.Add("device_id", t.deviceID)
	values.Add("device_platform", "web_pc")
	values.Add("device_type", "web_h264")
	values.Add("focus_state", "true")
	values.Add("from_page", "user")
	values.Add("history_len", "2")
	values.Add("is_fullscreen", "false")
	values.Add("is_page_visible", "false")
	values.Add("os", "windows")
	values.Add("region", "ID")
	values.Add("screen_height", "1080")
	values.Add("screen_width", "1920")
	values.Add("tz_name", "Asia/Bangkok")
	values.Add("webcast_language", "en")
	values.Add("room_id", roomID)
	if msToken != "" {
		values.Add("msToken", msToken)
	}
	// Print the url.Values
	urlQuery := values.EncodeWithOrder()
	bogus, err := GenerateBogus(urlQuery, t.userAgent)
	out := WebcastRoom{}
	if err != nil {
		if t.debug {
			log.Err(err).Msg("failed to generate bogus")
		}
		return out, nil, err
	}
	values.Add("X-Bogus", bogus)
	req := t.r.R()
	uri := "https://webcast.tiktok.com/webcast/room/enter/?" + values.EncodeWithOrder()
	// fmt.Println(uri)
	res, err := req.Get(uri)
	if err != nil {
		return out, nil, err
	}
	if res.StatusCode() > 400 {
		return out, res.Body(), errors.New(res.Status())
	}
	// fmt.Println(string(res.Body()))
	err = json.Unmarshal(res.Body(), &out)
	if err != nil {
		return out, res.Body(), err
	}
	if err := ParseStatusCode(out.StatusCode); err != nil {
		// fmt.Println(string(res.Body()))
		return out, nil, err
	}
	return out, nil, nil

}
