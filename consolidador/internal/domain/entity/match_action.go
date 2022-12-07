package entity

import "github.com/google/uuid"

type GameAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewGameAction(playerID, playerName string, minute int, action string, score int) *GameAction {
	return &GameAction{
		ID:         uuid.New().String(),
		PlayerID:   playerID,
		PlayerName: playerName,
		Minute:     minute,
		Action:     action,
		Score:      score,
	}
}
