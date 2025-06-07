package pgn

const (
	WHITE_QUEEN_MOVE_ZERO = ".."
	FIRST_WHITE_SQUARE    = "a6"
	FIRST_BLACK_SQUARE    = "f8"
	END_OF_GAME_DRAW      = "1/2-1/2"

	// 11-digit, zero-padded Event number
	PGN_HEADER_TEMPLATE = `[Event "%011d"]
[Site "?"]
[Date "????.??.??"]
[Round "?"]
[White "?"]
[Black "?"]
[Result "1/2-1/2"]
[WhiteELO "?"]
[BlackELO "?"]
[SetUp "1"]
`
)
