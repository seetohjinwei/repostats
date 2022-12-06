package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seetohjinwei/repostats/cli/messages"
	"github.com/seetohjinwei/repostats/models"
)

func initWalk() models.Directory {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(messages.PROMPT_DIRECTORY)
	for scanner.Scan() {
		maybeDir := scanner.Text()
		dir, err := parseDirectory(maybeDir, maybeDir, 0)
		if err == nil {
			return dir
		}
		fmt.Println(err.Error())
	}

	// Should never reach this.
	return models.Directory{}
}

func Start() {
	dir := initWalk()
	fmt.Println(messages.LISTING_TYPES)
	fmt.Println(dir.ListFileTypes())
	fmt.Println(dir.ListOptions())

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(messages.LIST_OPTIONS)

	for {
		fmt.Print(messages.PROMPT_OPTION)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "..":
			fmt.Println("Not yet implemented!")
		case "exit", "bye", "c":
			os.Exit(0)
		default:
			sub, err := dir.SubDirString(input)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			dir = *sub
			fmt.Println(dir.ListOptions())
		}
	}
}
