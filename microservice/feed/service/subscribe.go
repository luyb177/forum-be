package service

// import (
// 	"encoding/json"
// 	"github.com/Muxi-X/forum-be/microservice/feed/dao"
// 	logger "github.com/Muxi-X/forum-be/log"
// 	"github.com/Muxi-X/forum-be/model"
// )
//
// // SubServiceRun ... 写入feed数据
// func SubServiceRun() {
// 	var feed dao.FeedModel
//
// 	ch := model.PubSubClient.Self.Channel()
// 	for msg := range ch {
// 		logger.Info("received")
//
// 		if err := json.Unmarshal([]byte(msg.Payload), &feed); err != nil {
// 			panic(err)
// 		}
//
// 		if err := feed.Create(); err != nil {
// 			logger.Error(err.Error())
// 		}
// 	}
// }
