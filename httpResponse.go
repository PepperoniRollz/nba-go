package main

import "time"

type StandingsResponse struct {
	Get        string `json:"get"`
	Parameters struct {
		League     string `json:"league"`
		Conference string `json:"conference"`
		Season     string `json:"season"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Response []struct {
		League string `json:"league"`
		Season int    `json:"season"`
		Team   struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Nickname string `json:"nickname"`
			Code     string `json:"code"`
			Logo     string `json:"logo"`
		} `json:"team"`
		Conference struct {
			Name string `json:"name"`
			Rank int    `json:"rank"`
			Win  int    `json:"win"`
			Loss int    `json:"loss"`
		} `json:"conference"`
		Division struct {
			Name        string `json:"name"`
			Rank        int    `json:"rank"`
			Win         int    `json:"win"`
			Loss        int    `json:"loss"`
			GamesBehind string `json:"gamesBehind"`
		} `json:"division"`
		Win struct {
			Home       int    `json:"home"`
			Away       int    `json:"away"`
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
			LastTen    int    `json:"lastTen"`
		} `json:"win"`
		Loss struct {
			Home       int    `json:"home"`
			Away       int    `json:"away"`
			Total      int    `json:"total"`
			Percentage string `json:"percentage"`
			LastTen    int    `json:"lastTen"`
		} `json:"loss"`
		GamesBehind      string      `json:"gamesBehind"`
		Streak           int         `json:"streak"`
		WinStreak        bool        `json:"winStreak"`
		TieBreakerPoints interface{} `json:"tieBreakerPoints"`
	} `json:"response"`
}

type GamesResponse struct {
	Get        string `json:"get"`
	Parameters struct {
		Date string `json:"date"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Response []struct {
		ID     int    `json:"id"`
		League string `json:"league"`
		Season int    `json:"season"`
		Date   struct {
			Start    time.Time   `json:"start"`
			End      interface{} `json:"end"`
			Duration interface{} `json:"duration"`
		} `json:"date"`
		Stage  int `json:"stage"`
		Status struct {
			Clock    interface{} `json:"clock"`
			Halftime bool        `json:"halftime"`
			Short    int         `json:"short"`
			Long     string      `json:"long"`
		} `json:"status"`
		Periods struct {
			Current     int  `json:"current"`
			Total       int  `json:"total"`
			EndOfPeriod bool `json:"endOfPeriod"`
		} `json:"periods"`
		Arena struct {
			Name    string      `json:"name"`
			City    string      `json:"city"`
			State   string      `json:"state"`
			Country interface{} `json:"country"`
		} `json:"arena"`
		Teams struct {
			Visitors struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"visitors"`
			Home struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"home"`
		} `json:"teams"`
		Scores struct {
			Visitors struct {
				Win    int `json:"win"`
				Loss   int `json:"loss"`
				Series struct {
					Win  int `json:"win"`
					Loss int `json:"loss"`
				} `json:"series"`
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"visitors"`
			Home struct {
				Win    int `json:"win"`
				Loss   int `json:"loss"`
				Series struct {
					Win  int `json:"win"`
					Loss int `json:"loss"`
				} `json:"series"`
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"home"`
		} `json:"scores"`
		Officials   []interface{} `json:"officials"`
		TimesTied   interface{}   `json:"timesTied"`
		LeadChanges interface{}   `json:"leadChanges"`
		Nugget      interface{}   `json:"nugget"`
	} `json:"response"`
}
