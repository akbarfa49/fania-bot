package tiktok

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/RambIing/urlValues"
	"github.com/rs/zerolog/log"
)

var ErrUserDoesntExist = errors.New("UserNotExist")
var ErrTiktokInternalError = errors.New("TiktokInternalError")

func ParseStatusCode(i int) error {
	switch i {
	case 0:
		return nil
	case 10202, 10221:
		return ErrUserDoesntExist
	case 10002:
		return ErrTiktokInternalError

	default:
		fmt.Println(i)
		return errors.New("unknown error")
	}
}

type UserDetail struct {
	Extra struct {
		FatalItemIds []any  `json:"fatal_item_ids"`
		Logid        string `json:"logid"`
		Now          int64  `json:"now"`
	} `json:"extra"`
	LogPb struct {
		ImprID string `json:"impr_id"`
	} `json:"log_pb"`
	ShareMeta struct {
		Desc  string `json:"desc"`
		Title string `json:"title"`
	} `json:"shareMeta"`
	StatusCode  int    `json:"statusCode"`
	StatusCode0 int    `json:"status_code"`
	StatusMsg   string `json:"status_msg"`
	UserInfo    struct {
		Stats struct {
			DiggCount      int `json:"diggCount"`
			FollowerCount  int `json:"followerCount"`
			FollowingCount int `json:"followingCount"`
			FriendCount    int `json:"friendCount"`
			Heart          int `json:"heart"`
			HeartCount     int `json:"heartCount"`
			VideoCount     int `json:"videoCount"`
		} `json:"stats"`
		User struct {
			AvatarLarger string `json:"avatarLarger"`
			AvatarMedium string `json:"avatarMedium"`
			AvatarThumb  string `json:"avatarThumb"`
			BioLink      struct {
				Link string `json:"link"`
				Risk int    `json:"risk"`
			} `json:"bioLink"`
			CanExpPlaylist   bool `json:"canExpPlaylist"`
			CommentSetting   int  `json:"commentSetting"`
			CommerceUserInfo struct {
				Category       string `json:"category"`
				CategoryButton bool   `json:"categoryButton"`
				CommerceUser   bool   `json:"commerceUser"`
			} `json:"commerceUserInfo"`
			DownloadSetting        int    `json:"downloadSetting"`
			DuetSetting            int    `json:"duetSetting"`
			FollowingVisibility    int    `json:"followingVisibility"`
			Ftc                    bool   `json:"ftc"`
			ID                     string `json:"id"`
			IsADVirtual            bool   `json:"isADVirtual"`
			IsEmbedBanned          bool   `json:"isEmbedBanned"`
			NickNameModifyTime     int    `json:"nickNameModifyTime"`
			Nickname               string `json:"nickname"`
			OpenFavorite           bool   `json:"openFavorite"`
			PrivateAccount         bool   `json:"privateAccount"`
			ProfileEmbedPermission int    `json:"profileEmbedPermission"`
			ProfileTab             struct {
				ShowPlayListTab bool `json:"showPlayListTab"`
			} `json:"profileTab"`
			Relation      int    `json:"relation"`
			RoomID        string `json:"roomId"`
			SecUID        string `json:"secUid"`
			Secret        bool   `json:"secret"`
			Signature     string `json:"signature"`
			StitchSetting int    `json:"stitchSetting"`
			TtSeller      bool   `json:"ttSeller"`
			UniqueID      string `json:"uniqueId"`
			Verified      bool   `json:"verified"`
		} `json:"user"`
	} `json:"userInfo"`
}

func (t *Client) GetUserDetailByUniqueID(ctx context.Context, id string) (UserDetail, []byte, error) {
	values := urlValues.Values{}

	// Inserting query parameters manually
	msToken := ""
	for _, v := range t.r.Cookies {
		if v.Name == "msToken" {
			msToken = v.Value
			break
		}
	}

	// Inserting query parameters manually
	values.Add("WebIdLastTime", strconv.FormatInt(time.Now().Unix(), 10))
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
	values.Add("device_platform", "mobile")
	values.Add("focus_state", "true")
	values.Add("from_page", "user")
	values.Add("history_len", "1")
	values.Add("is_fullscreen", "false")
	values.Add("is_page_visible", "true")
	values.Add("language", "en")
	values.Add("os", "windows")
	values.Add("priority_region", "ID")
	values.Add("referer", "")
	values.Add("region", "ID")
	values.Add("screen_height", "1080")
	values.Add("screen_width", "1920")
	values.Add("tz_name", "Asia/Bangkok")
	values.Add("uniqueId", id)
	values.Add("webcast_language", "en")
	values.Add("msToken", msToken)
	// Print the url.Values
	urlQuery := values.EncodeWithOrder()
	bogus, err := GenerateBogus(urlQuery, t.userAgent)
	out := UserDetail{}
	if err != nil {
		if t.debug {
			log.Err(err).Msg("failed to generate bogus")
		}
		return out, nil, err
	}
	values.Add("X-Bogus", bogus)
	req := t.r.R()

	uri := "https://www.tiktok.com/api/user/detail/?" + values.EncodeWithOrder()
	res, err := req.Get(uri)
	if err != nil {
		return out, nil, err
	}
	if res.StatusCode() > 400 {
		return out, res.Body(), errors.New(res.Status())
	}
	err = json.Unmarshal(res.Body(), &out)
	if err != nil {
		return out, res.Body(), err
	}
	if err := ParseStatusCode(out.StatusCode); err != nil {
		return out, nil, err
	}
	return out, res.Body(), nil

}

// multi user
// https://www.tiktok.com/api/im/multi_user/?WebIdLastTime=1708568447&aid=1988&app_language=en&app_name=tiktok_web&browser_language=en-US&browser_name=Mozilla&browser_online=true&browser_platform=Win32&browser_version=5.0%20%28Windows%29&channel=tiktok_web&cookie_enabled=true&device_id=7338245547936581121&device_platform=web_pc&focus_state=false&from_page=fyp&history_len=2&is_fullscreen=false&is_page_visible=false&os=windows&priority_region=ID&referer=&region=ID&screen_height=1080&screen_width=1920&tz_name=Asia%2FBangkok&uids=6753936573073916930%2C6783899462454035458%2C7061501189221663770&webcast_language=en&msToken=_Hde4lbEnyV73Evj6mER2tPMS_NlXE1hy1vkRpTP5j3OE6grzoHIL5IoP_mLEq9V7-woLZqd-x32-xqZzfXzGhcBzi-6RxqEIGA906iH0GmY-gMJP-uws4P1MXc-B1y7LCf9&X-Bogus=DFSzswSLQMsANJj1tLCEKz9WcBrt

// func ( t *Client) GetUserPostList(ctx context.Context)
