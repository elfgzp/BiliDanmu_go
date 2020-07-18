package models

import (
	"fmt"
	"github.com/google/logger"
)

// Pool`s fields map CMD value
type Pool struct {
	UserMsg         chan string
	UserGift        chan string
	UserEnter       chan string
	UserGuard       chan string
	MsgUncompressed chan string
	UserEntry       chan string
	msgHandler PoolMsgHandler
}

func NewPool(msgHandler PoolMsgHandler) *Pool {
	return &Pool{
		UserMsg:         make(chan string, 10),
		UserGift:        make(chan string, 10),
		UserEnter:       make(chan string, 10),
		MsgUncompressed: make(chan string, 10),
		msgHandler: msgHandler,
	}
}

func (pool *Pool) Handle() {
	for {
		select {
		case uc := <-pool.MsgUncompressed:
			// 目前只处理未压缩数据的关注数变化信息
			if cmd := json.Get([]byte(uc), "cmd").ToString(); CMD(cmd) == CMDRoomRealTimeMessageUpdate {
				fans := json.Get([]byte(uc), "data", "fans").ToInt()
				fmt.Println("当前房间关注数变动：", fans)
			}
		case src := <-pool.UserMsg:
			m := NewDanmu()
			m.GetDanmuMsg([]byte(src))
			fmt.Printf("%d-%s | %d-%s: %s\n", m.MedalLevel, m.MedalName, m.Ulevel, m.Uname, m.Text)
			if pool.msgHandler == nil {
				continue
			}
			if err := pool.msgHandler.HandleUserMsg(m.MedalLevel, m.MedalName, m.Ulevel, m.Uname, m.Text); err != nil {
				logger.Errorf("pool.msgHandler.HandleUserMsg error %s", err.Error())
			}
		case src := <-pool.UserGift:
			g := NewGift()
			g.GetGiftMsg([]byte(src))
			fmt.Printf("%s %s 价值 %d 的 %s\n", g.Uname, g.Action, g.Price, g.GiftName)
			if pool.msgHandler == nil {
				continue
			}

			if err := pool.msgHandler.HandleUserGift(g.Uname, g.Action, g.Num, g.Price, g.GiftName); err != nil {
				logger.Errorf("pool.msgHandler.HandleUserGift error %s", err.Error())
			}
		case src := <-pool.UserEnter:
			name := json.Get([]byte(src), "data", "uname").ToString()
			fmt.Printf("欢迎VIP %s 进入直播间", name)
			if pool.msgHandler == nil {
				continue
			}
			if err := pool.msgHandler.HandleUserEnter(name); err != nil  {
				logger.Errorf("pool.msgHandler.HandleUserEnter error %s", err.Error())
			}
		case src := <-pool.UserGuard:
			name := json.Get([]byte(src), "data", "username").ToString()
			fmt.Printf("欢迎房管 %s 进入直播间", name)
			if pool.msgHandler == nil {
				continue
			}
			if err := pool.msgHandler.HandleUserGuard(name); err != nil {
				logger.Errorf("pool.msgHandler.HandleUserGuard error %s", err.Error())
			}
		case src := <-pool.UserEntry:
			cw := json.Get([]byte(src), "data", "copy_writing").ToString()
			fmt.Printf("%s", cw)
			if err := pool.msgHandler.HandleUserEntry(cw); err != nil {
				logger.Errorf("pool.msgHandler.HandleUserEntry error %s", err.Error())
			}
		}
	}
}

type PoolMsgHandler interface {
	HandleUserMsg(medalLevel uint32, medalName string, userLevel uint32, username, text string) error
	HandleUserGift(username string, action string, num, price int, giftName string) error
	HandleUserEnter(username string) error
	HandleUserGuard(username string) error
	HandleUserEntry(copyWriting string) error
}