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
	XWins            int                `bson:"xWins" json:"xWins"`
	OWins            int                `bson:"oWins" json:"oWins"`
	Draws            int                `bson:"draws" json:"draws"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateGame struct {
	GameName         string `bson:"gameName" json:"gameName" validate:"required"`
	CurrentPlayer    string `bson:"currentPlayer" json:"currentPlayer" validate:"required,oneof=X O"`
	BoardSize        int    `bson:"boardSize" json:"boardSize" validate:"min=3,max=15"`
	WinningCondition int    `bson:"winningCondition" json:"winningCondition" validate:"min=3,max=15,ltefield=BoardSize"`
}

type Move struct {
	GameName string `bson:"gameName" json:"gameName" validate:"required"`
	Player   string `bson:"player" json:"player" validate:"required,oneof=X O"`
	Row      int    `bson:"row" json:"row" validate:"min=0"`
	Col      int    `bson:"col" json:"col" validate:"min=0"`
}

type PushMoveMade struct {
	GameName      string `bson:"gameName" json:"gameName"`
	CurrentPlayer string `bson:"currentPlayer" json:"currentPlayer"` // "X" or "O"
	Row           int    `bson:"row" json:"row"`
	Col           int    `bson:"col" json:"col"`
}

type PushGameReset struct {
	GameName      string `bson:"gameName" json:"gameName"`
	CurrentPlayer string `bson:"currentPlayer" json:"currentPlayer"`
}

type PushPlayerJoin struct {
	GameName      string `bson:"gameName" json:"gameName"`
	JoiningPlayer string `bson:"joiningPlayer" json:"joiningPlayer"`
	UserId        string `bson:"userId" json:"userId"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type GetGamesCountResponse struct {
	Count int64 `json:"count"`
}
