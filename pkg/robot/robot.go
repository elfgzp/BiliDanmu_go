package robot

import (
	"biliDanMu/models"
	"fmt"
	"time"
)

type Robot struct {

}

func (r *Robot) Run(roomid uint32) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var err error

	if roomid == 0 {
		roomid = 234024
	}

	// 兼容房间号短 ID
	if roomid >= 100 && roomid < 1000 {
		roomid, err = models.GetRealRoomID(int(roomid))
		if err != nil {
			return err
		}

	}

	c, err := models.NewClient(roomid, r)
	if err != nil {
		return err
	}
	if err = c.Start(); err != nil {
		return err
	}

	time.Sleep(time.Minute * 3)
	return nil
}

func (r *Robot) HandleUserMsg(medalLevel uint32, medalName string, userLevel uint32, username, text string) error {
	return nil
}
func (r *Robot) HandleUserGift(username string, action string, num, price int, giftName string) error {
	return nil
}
func (r *Robot)  HandleUserEnter(username string) error {
	return nil
}
func (r *Robot)  HandleUserGuard(username string) error {
	return nil
}
func (r *Robot)  HandleUserEntry(copyWriting string) error {
	return nil
}