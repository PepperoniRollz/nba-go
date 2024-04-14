package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func getStandingsData(year, conference string) StandingsResponse {

	url := fmt.Sprintf("https://api-nba-v1.p.rapidapi.com/standings?league=standard&season=%s&conference=%s", year, conference)

	req, _ := http.NewRequest("GET", url, nil)

	apiKey, exists := os.LookupEnv("X_RAPIDAPI_KEY")
	if exists {

		req.Header.Add("X-RapidAPI-Key", apiKey)
	}
	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	var response StandingsResponse
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	return response
}

type standingsModel struct {
	westStandings table.Model
	eastStandings table.Model
}

func newStandingsModel(conference string) standingsModel {
	data := getStandingsData("2023", conference)
	return standingsModel{populateTables(data), table.Model{}}
}

func (m standingsModel) Init() tea.Cmd { return nil }

func (m standingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				tea.Printf("Let's go to %s!", m.westStandings.SelectedRow()[1]),
			)
		}
	}
	m.westStandings, cmd = m.westStandings.Update(msg)
	return m, cmd
}

func (m standingsModel) View() string {
	return baseStyle.Render(m.westStandings.View()) + "\n"
}

func populateTables(data StandingsResponse) table.Model {

	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "Team", Width: 25},
		{Title: "W", Width: 4},
		{Title: "L", Width: 4},
		{Title: "Pct", Width: 7},
		{Title: "GB", Width: 4},

		{Title: "Home", Width: 4},
		{Title: "Away", Width: 4},
		{Title: "L10", Width: 4},
		{Title: "Streak", Width: 6},
	}

	rows := make([]table.Row, 0)
	trows := make([][]string, 0)
	for _, team := range data.Response {
		newRow := []string{
			strconv.Itoa(team.Conference.Rank),
			team.Team.Name,
			strconv.Itoa(team.Win.Total),
			strconv.Itoa(team.Loss.Total),
			team.Win.Percentage,
			team.GamesBehind,
			strconv.Itoa(team.Win.Home),
			strconv.Itoa(team.Win.Away),
			strconv.Itoa(team.Win.LastTen),
			strconv.Itoa(team.Streak),
		}
		trows = append(trows, newRow)
	}
	sort.Slice(trows, func(i, j int) bool {
		num, err := strconv.Atoi(trows[i][0])
		num2, err2 := strconv.Atoi(trows[j][0])
		fmt.Println(num, num2)
		if err == nil && err2 == nil {
			return num < num2
		} else {
			return false
		}
	})

	for _, r := range trows {
		rows = append(rows, r)
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
