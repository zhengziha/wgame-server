package main

import (
	"fmt"

	"wgame-server/server/db"
	"wgame-server/server/model"
)

func main() {
	err := db.Init(&db.DBConfig{
		Driver:   "mysql",
		DSN:      "zhengzh:zheng@1013@tcp(139.196.189.119:3116)/w-game-preview?charset=utf8mb4&parseTime=True&loc=Local",
		LogLevel: 2,
	}, nil)
	if err != nil {
		fmt.Printf("Game DB init failed: %v\n", err)
		return
	}
	defer db.Close()

	var characters []model.Characters
	result := db.GORM().Limit(10).Find(&characters)
	if result.Error != nil {
		fmt.Printf("Query characters failed: %v\n", result.Error)
		return
	}

	fmt.Println("=== Game DB - characters (top 10) ===")
	for _, chara := range characters {
		fmt.Printf("id=%d name=%s account_id=%d level=%d map_id=%d x=%d y=%d\n",
			chara.ID, chara.Name, chara.AccountId, chara.Level, chara.MapId, chara.X, chara.Y)
	}

	var mapInfos []model.MapInfo
	result = db.GORM().Limit(10).Find(&mapInfos)
	if result.Error != nil {
		fmt.Printf("Query map_info failed: %v\n", result.Error)
		return
	}

	fmt.Println("\n=== Game DB - map_info (top 10) ===")
	for _, mi := range mapInfos {
		fmt.Printf("id=%d map_id=%d name=%s x=%d y=%d\n",
			mi.ID, mi.MapID, mi.Name, mi.X, mi.Y)
	}
}
