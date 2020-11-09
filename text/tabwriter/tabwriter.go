package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Tabwriter() {
	w := tabwriter.NewWriter(os.Stdout, 8, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\ta\tb")
	fmt.Fprintln(w, "c\td\tf")
	fmt.Fprintln(w, "a\ts\td\tf\tsdf")
	w.Flush()
}

func main() {
	Tabwriter()
}
