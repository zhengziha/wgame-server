package system

import (
	"math/rand"
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/msg/auth"
	system_msg "wgame-server/server/msg/system"
	"wgame-server/server/network/handler"
	"wgame-server/server/network/socket"
	"wgame-server/server/session"
)

const (
	ECHO_INTERVAL_LOWER_MS = 30
)

func init() {
	handler.Register(4274, "CmdEcho", CmdEchoHandler)
}

// CmdEchoHandler 处理 CMD_ECHO (cmd=4274)
// 心跳检测，包含加速器检测功能
func CmdEchoHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	currentTime, _ := reader.ReadInt()
	peerTime, _ := reader.ReadInt()

	socketCtx, ok := ctx.(*socket.SocketCmdContext)
	if !ok {
		return nil
	}

	sess := socketCtx.Session()
	speedCheck(sess, currentTime, peerTime, ctx)

	return nil
}

func speedCheck(sess *session.GameSession, currentTime, peerTime int32, ctx context.MyCmdContext) {
	if sess == nil {
		return
	}

	chara := sess.GetChara()
	if chara != nil {
		log.Info("[echo] 心跳检测: %s, currentTime=%d, peerTime=%d", chara.Name, currentTime, peerTime)
	} else {
		log.Info("[echo] 心跳检测(未登录): currentTime=%d, peerTime=%d", currentTime, peerTime)
	}

	sess.SetHeartEcho(session.NowMilli())
	lastEcho := sess.EchoTime()
	sess.SetEchoTime(currentTime)

	newPeerTime := peerTime + int32(rand.Intn(1401)+100)

	ctx.Write(&system_msg.MsgReplyEcho{A: newPeerTime})

	sess.SetOnline(true)

	if chara == nil {
		return
	}

	isAddSpeed := false

	detTime := int(currentTime - lastEcho)
	val := getAddSpeedTime()

	if detTime < val {
		log.Error("[echo] %s 心跳echo时间错误: detTime=%d, currentTime=%d, lastEcho=%d, val=%d",
			chara.Name, detTime, currentTime, lastEcho, val)
		isAddSpeed = true
	} else {
		sess.EchoWindow().Add(session.NowMilli())
		speedTimeInterval := getSpeedTimeInterval()

		if sess.EchoWindow().Size() >= speedTimeInterval {
			count := sess.EchoWindow().CountIntervalsBetween(ECHO_INTERVAL_LOWER_MS, val)
			threshold := sess.EchoWindow().Size() / 2

			if count >= threshold {
				log.Error("[echo] %s 心跳请求时间错误: 异常间隔次数=%d, 容错阈值=%d",
					chara.Name, count, threshold)
				isAddSpeed = true
			}

			sess.EchoWindow().Reset()
		}
	}

	if isAddSpeed {
		log.Error("[echo] %s 被系统检测到使用了加速器", chara.Name)
		ctx.Write(&auth.MsgKickOff{Msg: "系统检测到你开了#R加速器(或使用了辅助)#n，被强制下线！#R如若检测次数超限,将面临永久封号！"})
		ctx.Disconnect()
	}
}

func getAddSpeedTime() int {
	return 9500
}

func getSpeedTimeInterval() int {
	return 10
}
