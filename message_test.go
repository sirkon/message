package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplit(t *testing.T) {
	data := split("\033[1m%s lol\033[0m \033[31mabadfdf")
	require.Equal(t,
		[]string{
			"\033[1m",
			"%s lol",
			"\033[0m",
			" ",
			"\033[31m",
			"abadfdf",
		},
		data)

	res := printer(cyan, "\033[1m%s lol\033[0m \033[31mabadfdf")
	require.Equal(t, string(cyan)+"\033[1m%s lol\033[0m"+string(cyan)+" \033[31mabadfdf\033[0m", res)
}
