package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seetohjinwei/repostats/cli/messages"
	"github.com/seetohjinwei/repostats/models"
)

func initWalk() *models.Directory {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(messages.PROMPT_DIRECTORY)
	for scanner.Scan() {
		maybeDir := scanner.Text()
		dir, err := parseDirectory(maybeDir, maybeDir, 0)
		if err == nil {
			return &dir
		}
		fmt.Println(err.Error())
	}

	// Should never reach this.
	return &models.Directory{}
}

func Start() {
	index := 0
	dirs := []*models.Directory{initWalk()}

	fmt.Println(dirs[index].ListEverything())

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(messages.PROMPT_OPTION)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "..":
			if index <= 0 {
				fmt.Println("nothing to go back to")
				continue
			}
			index--
			fmt.Println(dirs[index].ListEverything())
		case "help":
			fmt.Print(messages.LIST_OPTIONS)
		case "exit", "bye", "c":
			os.Exit(0)
		default:
			sub, err := dirs[index].SubDirString(input)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			index++
			if index >= len(dirs) {
				dirs = append(dirs, sub)
			} else {
				dirs[index] = sub
			}
			fmt.Println(dirs[index].ListEverything())
		}
	}
}
