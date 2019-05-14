package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ritwik310/mini-db/server"
	"github.com/ritwik310/mini-db/src"
)

var store src.Store

func init() {
	store = src.Store{Persist: false}
	store.Map = make(map[string]interface{})
}

// Start ..
func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdStr, err := reader.ReadString('\n') // Reading Input
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		execCmd(FormatCmd(cmdStr[0 : len(cmdStr)-1])) // Printing response
	}
}

// execCmd executes CRUD command, and handles response
func execCmd(msg string) {
	bs := server.HandleMsg(&store, []byte(msg)) // Executing CRUD
	d, err := src.UnmarshalData(bs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if d["error"] != nil {
		fmt.Println("Error:", d["error"])
		return
	}

	if d["data"] != nil {
		fmt.Println(d["data"])
		return
	}
}

// FormatCmd formats the shell-command into formatted message in protocol format
func FormatCmd(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Join(strings.Fields(str), " ")
	strsl := strings.Split(str, " ")

	switch strsl[0] {
	case "GET":
		return formatNoTyp(strsl)
	case "SET":
		return formatWidTyp(strsl)
	case "DELETE":
		return formatNoTyp(strsl)
	}

	return ""
}

func formatNoTyp(strsl []string) string {
	var fstr string
	for _, s := range strsl {
		fstr += s + "\r\n"
	}

	return fstr
}

func formatWidTyp(strsl []string) string {
	var typ string
	typEl := strsl[len(strsl)-1]

	if len(typEl) >= 5 && typEl[0:2] == "--" {
		strsl = strsl[:len(strsl)-1]

		switch typEl[2:5] {
		case "str":
			typ = "+"
		case "int":
			typ = ":"
		case "bin":
			typ = "$"
		default:
			typ = "+"
		}
	} else {
		typ = "+"
	}

	var fstr string
	for i, s := range strsl {
		if i < 2 {
			fstr += s + "\r\n"
		} else {
			fstr += typ + s + "\r\n"
		}
	}

	return fstr
}
