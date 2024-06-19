package game

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ OBJECTS ------------------------------------------------
// ---------------------------------------------------------------------------------------------------
type XorO string

const (
	PlayerX XorO = "X"
	PlayerO XorO = "O"
)

type Cell struct {
	Value       *string `bson:"value" json:"value"` // "X", "O", or null
	Row         int     `bson:"row" json:"row"`
	Col         int     `bson:"col" json:"col"`
	WinningCell bool    `bson:"winningCell" json:"winningCell"`
}

type Row struct {
	Data []Cell `bson:"data" json:"data"`
}

type Board [][]Cell

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Game struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	GameName         string             `bson:"gameName" json:"gameName"`
	Winner           *string            `bson:"winner,omitempty" json:"winner,omitempty"`
	CurrentPlayer    string             `bson:"currentPlayer" json:"currentPlayer"`
	WinningCondition int                `bson:"winningCondition" json:"winningCondition"`
	Board            Board              `bson:"board" json:"board"`
	IsDraw           bool               `bson:"isDraw" json:"isDraw"`
	XPlayer          *string            `bson:"xPlayer,omitempty" json:"xPlayer,omitempty"`
	OPlayer          *string            `bson:"oPlayer,omitempty" json:"oPlayer,omitempty"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateGame struct {
	GameName         string `bson:"gameName" json:"gameName"`
	CurrentPlayer    string `bson:"currentPlayer" json:"currentPlayer"`
	BoardSize        int    `bson:"boardSize" json:"boardSize"`
	WinningCondition int    `bson:"winningCondition" json:"winningCondition"`
}

type Move struct {
	GameName string `bson:"gameName" json:"gameName"`
	Player   string `bson:"player" json:"player"` // "X" or "O"
	Row      int    `bson:"row" json:"row"`
	Col      int    `bson:"col" json:"col"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type GetGamesCountResponse struct {
	Count int64 `json:"count"`
}
