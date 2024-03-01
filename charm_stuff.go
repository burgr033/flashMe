package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"math/rand"
	"strconv"
	"time"
)

type RenderingOption int

type model struct {
	cards        []flashCard
	currentIndex int
	showAnswer   bool
}

const (
	TitleOption    RenderingOption = iota
	QuestionOption RenderingOption = iota
	AnswerOption   RenderingOption = iota
)

var (
	dialogBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 1).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
)

func renderHelpBar(option RenderingOption, m *model) string {
	var messages string
	questionCount := strconv.Itoa(m.currentIndex + 1)
	questionMax := strconv.Itoa(len(m.cards))
	switch option {
	case TitleOption:
		messages = "[🡒]: start • [q]: quit"
	case AnswerOption:
		messages = "[🡒]: next • [r]: repeat question • [q]: quit • Cards:" + questionCount + " / " + questionMax
	default:
		messages = "[🡒]: next • [space]: reveal • [q]: quit • Cards:" + questionCount + " / " + questionMax
	}
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	helpBarContent := helpStyle.Copy().Width(64).Render(messages)
	bar := lipgloss.JoinHorizontal(lipgloss.Center,
		helpBarContent,
	)
	return bar
}

func renderCard(content string, lineHeight *int, option RenderingOption) string {
	var contentR string
	switch option {
	case AnswerOption:
		contentR = lipgloss.NewStyle().Width(64).Height(*lineHeight).AlignVertical(lipgloss.Center).AlignHorizontal(lipgloss.Center).Italic(true).Render(content)
	default:
		contentR = lipgloss.NewStyle().Width(64).Height(*lineHeight).AlignVertical(lipgloss.Center).AlignHorizontal(lipgloss.Center).Bold(true).Render(content)
	}
	ui := lipgloss.JoinVertical(0, contentR)
	dialog := lipgloss.Place(1, 1,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
	)
	return dialog
}

func addCardToPoolAgain(m *model) {
	m.cards = append(m.cards, m.cards[m.currentIndex])
}

func (m model) View() string {
	maxLines := calulateMaxLines(m.cards)
	if m.currentIndex == -1 {
		return renderCard("Welcome", &maxLines, TitleOption) + "\n" +
			renderHelpBar(TitleOption, &m)
	}
	if m.currentIndex >= len(m.cards) {
		return renderCard("Congratulations!", &maxLines, TitleOption) + "\n" +
			renderHelpBar(TitleOption, &m)
	}
	card := m.cards[m.currentIndex]
	if !m.showAnswer {
		return renderCard(card.question, &maxLines, QuestionOption) + "\n" +
			renderHelpBar(QuestionOption, &m)

	} else {

		return renderCard("Answer: "+card.answer, &maxLines, AnswerOption) + "\n" +
			renderHelpBar(AnswerOption, &m)
	}
}

func initialModel() model {
	return model{
		cards:        shuffleFlashCards(getDummyDataSlice()),
		currentIndex: -1,
		showAnswer:   false,
	}
}
func (m model) Init() tea.Cmd {
	rand.Seed(time.Now().UnixNano())
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case " ":
			m.showAnswer = !m.showAnswer
		case "r":
			if m.showAnswer {
				addCardToPoolAgain(&m)
				m.currentIndex = m.currentIndex + 1
				m.showAnswer = false
			}
		case "right":
			m.currentIndex = m.currentIndex + 1
			m.showAnswer = false
		}
	}
	return m, nil
}
