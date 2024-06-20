package game

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateGame validates the CreateGame struct
func ValidateCreateGame(game CreateGame) error {
	err := validate.Struct(game)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "GameName":
				return errors.New("game name is required")
			case "CurrentPlayer":
				return errors.New("current player is required and must be 'X' or 'O'")
			case "BoardSize":
				return errors.New("board size is required and must be between 3 and 15")
			case "WinningCondition":
				return errors.New("winning condition is required, must be between 3 and 15, and less than or equal to board size")
			}
		}
	}
	return nil
}

// ValidateMove validates the Move struct
func ValidateMove(move Move) error {
	err := validate.Struct(move)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "GameName":
				return errors.New("game name is required")
			case "Player":
				return errors.New("player is required and must be 'X' or 'O'")
			case "Row":
				return errors.New("row is required and must be 0 or greater")
			case "Col":
				return errors.New("column is required and must be 0 or greater")
			}
		}
	}
	return nil
}
