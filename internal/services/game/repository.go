package game

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"online-tictactoe/internal/db"
	"online-tictactoe/internal/pusher"
	"time"
)

type GameRepository interface {
	CreateGame(c context.Context, createGame CreateGame, userId string) (Game, error)
	GetGame(c context.Context, universityID string, userId string) (Game, error)
	GetGamesCount(c context.Context) (GetGamesCountResponse, error)
	CreateMove(c context.Context, move Move, userId string) (Game, error)
	ResetGame(c context.Context, gameName string, userId string) (Game, error)
	ClearAllGames(c context.Context) error
}

type repositoryImpl struct {
	collection *mongo.Collection
}

/**
 * NewGameRepository creates a new game repository that will interact with the database
 * @return GameRepository
 */
func NewGameRepository() GameRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_GAME)
	return &repositoryImpl{collection: collection}
}

/**
 * CreateGame creates a new game with the given parameters
 * @param c context.Context
 * @param createGame CreateGame
 * @param userId string
 * @return Game
 * @return error
 */
func (r *repositoryImpl) CreateGame(c context.Context, createGame CreateGame, userId string) (Game, error) {
	// Check if a game with the same name already exists
	var existingGame Game
	err := r.collection.FindOne(c, bson.M{"gameName": createGame.GameName}).Decode(&existingGame)
	if err == nil {
		return Game{}, errors.New("game name already exists! Join the game instead")
	} else if err != mongo.ErrNoDocuments {
		return Game{}, err
	}

	// Create an empty board
	board := make(Board, createGame.BoardSize)
	for i := range board {
		board[i] = make([]Cell, createGame.BoardSize)
		for j := range board[i] {
			board[i][j] = Cell{
				Value:       nil,
				Row:         i,
				Col:         j,
				WinningCell: false,
			}
		}
	}

	game := Game{
		GameName:         createGame.GameName,
		CurrentPlayer:    createGame.CurrentPlayer,
		WinningCondition: createGame.WinningCondition,
		Board:            board,
		IsDraw:           false,
		XPlayer:          &userId,
		XWins:            0,
		OWins:            0,
		Draws:            0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	res, err := r.collection.InsertOne(c, game)
	if err != nil {
		return Game{}, err
	}

	game.ID = res.InsertedID.(primitive.ObjectID)
	return game, nil
}

/**
 * GetGame gets the game with the given parameters
 * @param c context.Context
 * @param gameName string
 * @param userId string
 * @return Game
 * @return error
 */
func (r *repositoryImpl) GetGame(c context.Context, gameName string, userId string) (Game, error) {
	var game Game
	err := r.collection.FindOne(c, bson.M{"gameName": gameName}).Decode(&game)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Game{}, errors.New("game not found")
		}
		return Game{}, err
	}

	// Check the player slots
	if game.XPlayer != nil && game.OPlayer != nil {
		// Both slots are filled, check if the user is one of them
		if *game.XPlayer != userId && *game.OPlayer != userId {
			return Game{}, errors.New("forbidden")
		}
	} else if game.XPlayer != nil && game.OPlayer == nil {
		// xPlayer is filled, oPlayer is empty
		if *game.XPlayer == userId {
			return game, nil
		}
		_, err = r.collection.UpdateOne(
			c,
			bson.M{"gameName": gameName},
			bson.M{"$set": bson.M{"oPlayer": userId}},
		)
		if err != nil {
			return Game{}, err
		}

		// Trigger a Pusher event to notify other players
		pusher.Client.Trigger(pusher.GameChannel, pusher.PlayerJoined, PushPlayerJoin{
			GameName:      game.GameName,
			JoiningPlayer: "O",
			UserId:        userId,
		})
	} else if game.XPlayer == nil && game.OPlayer != nil {
		// oPlayer is filled, xPlayer is empty
		if *game.OPlayer == userId {
			return game, nil
		}
		_, err = r.collection.UpdateOne(
			c,
			bson.M{"gameName": gameName},
			bson.M{"$set": bson.M{"xPlayer": userId}},
		)
		if err != nil {
			return Game{}, err
		}

		// Trigger a Pusher event to notify other players
		pusher.Client.Trigger(pusher.GameChannel, pusher.PlayerJoined, PushPlayerJoin{
			GameName:      game.GameName,
			JoiningPlayer: "X",
			UserId:        userId,
		})
	} else {
		// Both slots are empty, this should ideally not happen in a valid game setup
		return Game{}, errors.New("internal server error")
	}

	// Find the updated game
	err = r.collection.FindOne(c, bson.M{"gameName": gameName}).Decode(&game)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

/**
 * GetGamesCount gets the count of all games
 * @param c context.Context
 * @return GetGamesCountResponse
 * @return error
 */
func (r *repositoryImpl) GetGamesCount(c context.Context) (GetGamesCountResponse, error) {
	count, err := r.collection.CountDocuments(c, bson.M{})
	if err != nil {
		return GetGamesCountResponse{}, err
	}

	return GetGamesCountResponse{Count: count}, nil
}

/**
 * CreateMove creates a move in the game with the given parameters
 * @param c context.Context
 * @param move Move
 * @param userId string
 * @return Game
 * @return error
 */
func (r *repositoryImpl) CreateMove(c context.Context, move Move, userId string) (Game, error) {
	var game Game
	err := r.collection.FindOne(c, bson.M{"gameName": move.GameName}).Decode(&game)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Game{}, errors.New("game not found")
		}
		return Game{}, err
	}

	// Validate the player
	if (move.Player != "X" && move.Player != "O") || (move.Player == "X" && game.XPlayer == nil) || (move.Player == "O" && game.OPlayer == nil) {
		return Game{}, errors.New("invalid player")
	}

	if game.CurrentPlayer != move.Player {
		return Game{}, errors.New("incorrect player turn in payload")
	}

	if (move.Player == "X" && *game.XPlayer != userId) || (move.Player == "O" && *game.OPlayer != userId) {
		return Game{}, errors.New("forbidden")
	}

	if game.Winner != nil {
		return Game{}, errors.New("game already won")
	}

	if game.Board[move.Row][move.Col].Value != nil {
		return Game{}, errors.New("cell already occupied")
	}

	// Validate the move
	if move.Row < 0 || move.Row >= len(game.Board) || move.Col < 0 || move.Col >= len(game.Board[0]) {
		return Game{}, errors.New("invalid move")
	}

	// Update the board
	game.Board[move.Row][move.Col].Value = &move.Player
	game.UpdatedAt = time.Now()

	// Check if the game is won
	winningPath := IsWinner(game.Board, move.Player, game.WinningCondition)
	if winningPath != nil {
		game.Winner = &move.Player
		if *game.Winner == "X" {
			game.XWins++
		} else {
			game.OWins++
		}
		for _, cell := range winningPath {
			game.Board[cell.Row][cell.Col].WinningCell = true
		}
	}

	// Check if the game is a draw
	if game.Winner == nil && IsGameDraw(game.Board) {
		game.IsDraw = true
		game.Draws++
	}

	// Update the current player
	if game.CurrentPlayer == "X" {
		game.CurrentPlayer = "O"
	} else {
		game.CurrentPlayer = "X"
	}

	// Update the game in the database
	_, err = r.collection.UpdateOne(
		c,
		bson.M{"gameName": move.GameName},
		bson.M{
			"$set": bson.M{
				"board":         game.Board,
				"winner":        game.Winner,
				"currentPlayer": game.CurrentPlayer,
				"isDraw":        game.IsDraw,
				"xWins":         game.XWins,
				"oWins":         game.OWins,
				"draws":         game.Draws,
				"updatedAt":     game.UpdatedAt,
			},
		},
	)
	if err != nil {
		return Game{}, err
	}

	// Trigger a Pusher event to notify other players
	err = pusher.Client.Trigger(pusher.GameChannel, pusher.MoveMade, PushMoveMade{
		GameName:      game.GameName,
		CurrentPlayer: game.CurrentPlayer,
		Row:           move.Row,
		Col:           move.Col,
	})
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

/**
 * ResetGame resets the game with the given parameters
 * @param c context.Context
 * @param gameName string
 * @return Game
 * @return error
 */
func (r *repositoryImpl) ResetGame(c context.Context, gameName string, userId string) (Game, error) {

	var game Game
	err := r.collection.FindOne(c, bson.M{"gameName": gameName}).Decode(&game)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Game{}, errors.New("game not found")
		}
		return Game{}, err
	}

	// Check if the user is a player in the game
	if game.XPlayer != nil && *game.XPlayer != userId && game.OPlayer != nil && *game.OPlayer != userId {
		return Game{}, errors.New("forbidden")
	}

	// Reset the game
	board := make(Board, len(game.Board))
	for i := range board {
		board[i] = make([]Cell, len(game.Board[i]))
		for j := range board[i] {
			board[i][j] = Cell{
				Value:       nil,
				Row:         i,
				Col:         j,
				WinningCell: false,
			}
		}
	}

	game.Board = board
	game.Winner = nil
	game.CurrentPlayer = "X"
	game.IsDraw = false
	game.UpdatedAt = time.Now()

	// Update the game in the database
	_, err = r.collection.UpdateOne(
		c,
		bson.M{"gameName": gameName},
		bson.M{
			"$set": bson.M{
				"board":         game.Board,
				"winner":        game.Winner,
				"currentPlayer": game.CurrentPlayer,
				"isDraw":        game.IsDraw,
				"updatedAt":     game.UpdatedAt,
			},
		},
	)
	if err != nil {
		return Game{}, err
	}

	// Trigger a Pusher event to notify other players
	err = pusher.Client.Trigger(pusher.GameChannel, pusher.GameReset, PushGameReset{
		GameName:      game.GameName,
		CurrentPlayer: game.CurrentPlayer,
	})
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

/**
 * ClearGamesDB clears the games database
 * @param c context.Context
 * @return error
 */
func (r *repositoryImpl) ClearAllGames(c context.Context) error {
	_, err := r.collection.DeleteMany(c, bson.M{})
	if err != nil {
		return err
	}
	return nil
}
