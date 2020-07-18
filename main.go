package main

import (
	"biliDanMu/pkg/robot"
	"os"
	"strconv"
)

func main() {
	r := robot.Robot{}
	roomid, _ := strconv.Atoi(os.Getenv("ROOMID"))
	r.Run(uint32(roomid))
}
