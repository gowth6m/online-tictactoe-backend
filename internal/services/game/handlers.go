package game

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-tictactoe/internal/api"
)

type GameHandler struct {
	Repo GameRepository
}

func NewGameHandler(repo GameRepository) *GameHandler {
	return &GameHandler{Repo: repo}
}

// @Summary Create a new game
// @Description Create a new game with the given parameters
// @Tags game
// @Accept json
// @Produce json
// @Param createGame body CreateGame true "Game details"
// @Success 201 {object} Game
// @Failure 400 {object} api.ResponseData
// @Failure 500 {object} api.ResponseData
// @Router /game/create [post]
func (h *GameHandler) CreateGame(c *gin.Context) {
	var createGame CreateGame
	if err := c.ShouldBindJSON(&createGame); err != nil {
		api.Error(c, http.StatusBadRequest, "Invalid request format or parameters", nil)
		return
	}

	if field, err := ValidateCreateGame(createGame); err != nil {
		fieldErrors := []api.FieldError{
			{
				Field:   &field,
				Message: err.Error(),
			},
		}
		api.Error(c, http.StatusBadRequest, err.Error(), &fieldErrors)
		return
	}

	game, err, fieldErrors := h.Repo.CreateGame(c, createGame, c.GetString("userId"))
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), &fieldErrors)
		return
	}

	api.Success(c, http.StatusCreated, "Game created successfully", game)
}

// @Summary Get a game
// @Description Get the game with the given parameters
// @Tags game
// @Accept json
// @Produce json
// @Param gameName query string true "Game name"
// @Success 200 {object} Game
// @Failure 400 {object} api.ResponseData
// @Failure 404 {object} api.ResponseData
// @Failure 500 {object} api.ResponseData
// @Router /game/{gameName} [get]
func (h *GameHandler) GetGame(c *gin.Context) {

	gameName := c.Param("gameName")
	if gameName == "" {
		api.Error(c, http.StatusBadRequest, "Invalid request format or parameters", nil)
		return
	}

	game, err := h.Repo.GetGame(c, gameName, c.GetString("userId"))
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Game retrieved successfully", game)
}

// @Summary Get number of games in the database
// @Description Get the number of games in the database
// @Tags game
// @Accept json
// @Produce json
// @Success 200 {object} GetGamesCountResponse
// @Failure 500 {object} api.ResponseData
// @Router /game/all/count [get]
func (h *GameHandler) GetAllGamesCount(c *gin.Context) {

	count, err := h.Repo.GetGamesCount(c)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Games count retrieved successfully", count)
}

// @Summary Post a move
// @Description Post a move to the game with the given parameters
// @Tags game
// @Accept json
// @Produce json
// @Param createMove body Move true "Move details"
// @Success 200 {object} Game
// @Failure 400 {object} api.ResponseData
// @Failure 404 {object} api.ResponseData
// @Failure 500 {object} api.ResponseData
// @Router /game/move [post]
func (h *GameHandler) CreateMove(c *gin.Context) {
	var move Move
	if err := c.ShouldBindJSON(&move); err != nil {
		api.Error(c, http.StatusBadRequest, "Invalid request format or parameters", nil)
		return
	}

	if field, err := ValidateMove(move); err != nil {
		fieldErrors := []api.FieldError{
			{
				Field:   &field,
				Message: err.Error(),
			},
		}
		api.Error(c, http.StatusBadRequest, err.Error(), &fieldErrors)
		return
	}

	game, err := h.Repo.CreateMove(c, move, c.GetString("userId"))
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Move made successfully", game)
}

// @Summary Reset a game given the game name
// @Description Reset a game with the given parameters
// @Tags game
// @Accept json
// @Produce json
// @Param gameName query string true "Game name"
// @Success 200 {object} Game
// @Failure 400 {object} api.ResponseData
// @Failure 404 {object} api.ResponseData
// @Failure 500 {object} api.ResponseData
// @Router /game/{gameName}/reset [post]
func (h *GameHandler) ResetGame(c *gin.Context) {

	gameName := c.Param("gameName")
	if gameName == "" {
		api.Error(c, http.StatusBadRequest, "Invalid request format or parameters", nil)
		return
	}

	game, err := h.Repo.ResetGame(c, gameName, c.GetString("userId"))
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Game reset successfully", game)
}

// @Summary Clears all games in the database
// @Description Clears all games in the database (use with caution) - only for Cron Jobs
// @Tags game
// @Accept json
// @Produce json
// @Success 200 {object} api.ResponseData
// @Failure 500 {object} api.ResponseData
// @Router /game/all/clear [post]
func (h *GameHandler) ClearAllGames(c *gin.Context) {

	err := h.Repo.ClearAllGames(c)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	api.Success(c, http.StatusOK, "Games cleared successfully", nil)
}
