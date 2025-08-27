package tiktok

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	client := New("msToken=iyi6rRnafEjq85H7T_N7RQGyniM1ThW1GNNlNJUEdPGnlOijrfjpryPr0UjWD4Q4K1_7TyvK8sbOL5XmhlYHt9a7jZstYIzK9KEl3PUGzKECraOUkDkCFlOdZKMkIa3s47Odxs8=", "7338245547936581121")
	userDetail, _, err := client.GetUserDetailByUniqueID(context.Background(), "qiyy.1")

	assert.Empty(t, err)
	assert.NotEmpty(t, userDetail)
	log.Println(userDetail.UserInfo.User.Nickname, userDetail.UserInfo.User.UniqueID, userDetail.UserInfo.User.SecUID, userDetail.UserInfo.User.ID)
	fmt.Println("roomid", userDetail.UserInfo.User.RoomID)
	// fmt.Println(string(b))

}
