package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type gamesTodayModel struct {
	gamesTable table.Model
}

func (m gamesTodayModel) Init() tea.Cmd {
	return nil
}

func (m gamesTodayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return initialModel(), nil
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.gamesTable.SelectedRow()[1]),
			)
		}
	}
	m.gamesTable, cmd = m.gamesTable.Update(msg)
	return m, cmd

}

func (m gamesTodayModel) View() string {
	return baseStyle.Render(m.gamesTable.View()) + "\n"
}

func getGamesData() GamesResponse {
	url := "https://api-nba-v1.p.rapidapi.com/games?date=2024-04-14"

	req, _ := http.NewRequest("GET", url, nil)

	apiKey, exists := os.LookupEnv("X_RAPIDAPI_KEY")

	if exists {
		req.Header.Add("X-RapidAPI-Key", apiKey)
	}

	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	var response GamesResponse
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return response
}

func populateGameTable(data GamesResponse) table.Model {
	columns := []table.Column{
		{Title: "Home", Width: 15},
		{Title: "Away", Width: 15},
		{Title: "Score", Width: 12},
		{Title: "Quarter", Width: 7},
		{Title: "Status", Width: 12},
		{Title: "Start Time", Width: 20},
	}
	rows := make([]table.Row, 0)

	for _, game := range data.Response {
		newRow := []string{game.Teams.Home.Nickname,
			game.Teams.Visitors.Nickname,
			fmt.Sprintf("%s : %s", strconv.Itoa(game.Scores.Home.Points), strconv.Itoa(game.Scores.Visitors.Points)),
			strconv.Itoa(game.Periods.Current),
			game.Status.Long,
			game.Date.Start.String(),
		}
		rows = append(rows, newRow)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(16),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}
func newGamesModel() gamesTodayModel {
	data := getGamesData()
	return gamesTodayModel{populateGameTable(data)}
}
