package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName = "GO_CLI_REMINDER"
	markValue = "1"
)

func main() {
	fmt.Println("Welcome to Gominder!")
	if len(os.Args) < 3 {
		fmt.Println("Usage: gominder <time> <message>")
		os.Exit(1)
	}

	now = time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if t == nil {
		fmt.Println("Error: ", "Cannot parse time")
		os.Exit(2)
	}
	now.After(t.Time) {
		fmt.Println("Error: ", "Time is in the past")
		os.Exit(3)
	}

	difference := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(difference)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(5)
		}
		fmt.Println("Reminder set for", difference.Round(time.Second)
		os.Exit(0)
	}
}