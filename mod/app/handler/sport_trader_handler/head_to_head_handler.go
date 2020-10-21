package sport_trader_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/sport_trader_case"
)

type (
	ProvideHeadToHeads echo.HandlerFunc

	HeadToHeadParam struct {
		Team1ID string `json:"team1_id"`
		Team2ID string `json:"team2_id"`
	}
)

func (p ProvideHeadToHeads) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideHeadToHeadsHandler(uc usecase.HeadToHeads) ProvideHeadToHeads {
	return func(c echo.Context) error {
		param := HeadToHeadParam{}
		if err := c.Bind(&param); err != nil {
			return err
		}
		headToHeads, err := uc(c.Request().Context(), param.Team1ID, param.Team2ID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, headToHeads)
	}
}
