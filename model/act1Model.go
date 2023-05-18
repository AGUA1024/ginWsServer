package model

import "sync"

const GANE_TYPE = 1

type Act1Model struct {
	actBaseModel
}

func NewAct1Model(uid int32) ActBaseInterface {
	var act1ModelOnce sync.Once
	var act1Model *Act1Model
	act1ModelOnce.Do(func() {
		act1Model = &Act1Model{
			actBaseModel: actBaseModel{
				uid:   uid,
				actId: GANE_TYPE,
				actInfo: map[string]any{
					"playerId":         -1,        // 玩家id
					"playerList":       [][]int{}, // playerId : cardId
					"curPlayerId":      -1,        // 当前出牌玩家id
					"curPlayerRemTime": -1,        // 当前玩家剩余时间
					"cardPool":         []int{},   // 卡池
				},
			},
		}
	})

	act1Model.Init()
	return act1Model
}
