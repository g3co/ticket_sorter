package ticket_sorter

import (
	"github.com/g3co/ticket_sorter/parser"
)

type TicketSort struct {
	parser parser.Parser
}

type ITicketSort interface {
	Process(cards []string) ([]string, error)
}

func NewTicketSorter(parser parser.Parser) TicketSort {
	return TicketSort{parser: parser}
}
