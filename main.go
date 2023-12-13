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
	