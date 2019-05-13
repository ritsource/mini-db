package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ritwik310/mini-db/server"
	"github.com/ritwik310/mini-db/src"
)

// Start ..
func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fstr := formatCmd(cmdStr[0 : len(cmdStr)-1])

		// fmt.Println(fstr)
		execCmd(fstr)
		// cmdString = strings.TrimSuffix(cmdString, "\n")
		// cmd := exec.Command(cmdString)
		// cmd.Stderr = os.Stderr
		// cmd.Stdout = os.Stdout
		// err = cmd.Run()
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// }
	}
}

func execCmd(msg string) {
	bs := server.HandleMsg([]byte(msg))
	d, err := src.UnmarshalData(bs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if d["error"] != nil {
		fmt.Println(err)
		return
	}

	if d["data"] != nil {
		fmt.Println(d["data"])
		return
	}
}

func formatCmd(str string) string {
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
		switch typEl[2:6] {
		// case "str":
		// 	typ = "+"
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
