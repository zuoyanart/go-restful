// Package provides ability to print colored text on stdout
package clicolor

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"github.com/shiena/ansicolor"
	"runtime"
)

const (
	startSeq = "\033["
	endSeq   = "\033[0m"
)

type Printer struct {
	text string
}

var colors map[string]string

// Will be used as first argument of fmt.Fprintln
// so with this you are able to redirect output stream
var Out io.Writer

var (
	// Used to find groups that need to be colored
	colorGroupRE *regexp.Regexp = regexp.MustCompile(`(\{\w*\}[^{}]+)`)

	// Used to extract color part of `{color} text`
	colorPartRE *regexp.Regexp = regexp.MustCompile(`{(\w*)}`)

	// Used to extract text part of `{color} text` 
	textPartRE *regexp.Regexp = regexp.MustCompile(`^{\w*}(.*)`)	
)

// Printing colored text in one choosed color
func (p *Printer) In(color string) {
	p.text = "{" + color + "}" + p.text
	p.InFormat()
}

// Printing colored text based on user format
// User needs to provide one of supported colors
// Example: `this{red} is red{blue} and this is blue. {default} Now is default`
func (p *Printer) InFormat() {
	// find all text groups which need to be colored
	matches := colorGroupRE.FindAllStringSubmatch(p.text, -1)

	for _, value := range matches {
		// extract color from `{color} some text`
		color := colorPartRE.FindStringSubmatch(value[0])[1]
		colorcode := getColor(color)

		// extract text from `{color} some text`
		text := textPartRE.FindStringSubmatch(value[0])[1]

		// format string for ANSI/VT1000 terminal
		clifmt := startSeq + colorcode + text + endSeq
		p.text = strings.Replace(p.text, value[0], clifmt, -1)
	}

	fmt.Fprintln(Out, p.text)
}

// Initialization of supported colors
func init() {
	colors = make(map[string]string)
	colors["black"] = "30m"
	colors["red"] = "31m"
	colors["green"] = "32m"
	colors["yellow"] = "33m"
	colors["blue"] = "34m"
	colors["magenta"] = "35m"
	colors["cyan"] = "36m"
	colors["white"] = "37m"
	colors["default"] = "39m"

	if runtime.GOOS == "windows" {
    	Out = ansicolor.NewAnsiColorWriter(os.Stdout)
	} else {
		Out = os.Stdout
	}
}

// Finding color code, and returning default one if requested doesn't exit
func getColor(color string) string {
	var colorcode string

	if value, ok := colors[color]; ok {
		colorcode = value
	} else {
		colorcode = colors["default"]
	}

	return colorcode
}

// Provide text for printing
func Print(color string) *Printer {
	return &Printer{color}
}
