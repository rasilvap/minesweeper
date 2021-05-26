package docs

import (
	"minesweeper-API/models/dto"
)

// swagger:route GET /games/{id} games getGameEndpoint
// Play a movement.
// responses:
//   200: gameResponseResponseWrapper
//	 404: notfound Not Found
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response gameResponseResponseWrapper
type GameResponseResponseWrapper struct {
	// in:body
	Body dto.GetGameResponse
}

// Game attributes to set
// swagger:parameters getGameEndpoint
type GameRequestParamsWrapper struct {

	// The game id
	// required:true
	// in:path
	GameID int64 `json:"id"`
}

// swagger:route POST /games games createGameEndpoint
// Create a game.
// responses:
//   200: gameSimpleResponseResponse
//	 400: badrequest Missing or invalid attributes in body
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response gameSimpleResponseResponse
type gameSimpleResponseResponseWrapper struct {
	// in:body
	Body dto.CreateGameResponse
}

// Game attributes to set
// swagger:parameters createGameEndpoint
type gameParamsWrapper struct {
	// The game's attributes to set.
	// in:body
	Body dto.CreateGameRequest
}

// swagger:route POST /games/{id}/play games playGameEndpoint
// Play a movement.
// responses:
//   200: playResponseResponseWrapper
//	 400: badrequest Missing or invalid attributes in body
//	 404: notfound Not Found
//   500: internal Internal Server Errors
// Game created OK.
// swagger:response playResponseResponseWrapper
type PlayResponseResponseWrapper struct {
	// in:body
	Body dto.PlayResponse
}

// Play attributes to set
// swagger:parameters playGameEndpoint
type PlayRequestParamsWrapper struct {
	// The game's attributes to set.
	// in:body
	Body dto.PlayRequest

	// The game id
	// required:true
	// in:path
	GameID int64 `json:"id"`
}
