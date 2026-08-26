package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"str-sim/internal/server"
	"str-sim/internal/sim"
)

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "str-sim: "+format+"\n", args...)
	os.Exit(1)
}

func reorderFlags(args []string) []string {
	var flags, pos []string
	i := 0
	for i < len(args) {
		a := args[i]
		if strings.HasPrefix(a, "-") {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") && !strings.Contains(a, "=") {
				flags = append(flags, a, args[i+1])
				i += 2
			} else {
				flags = append(flags, a)
				i++
			}
		} else {
			pos = append(pos, a)
			i++
		}
	}
	return append(flags, pos...)
}

func firstPositional(args []string) int {
	for i, a := range args {
		if !strings.HasPrefix(a, "-") {
			return i
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		runHTTPServer(nil)
		return
	}
	if os.Args[1] == "serve" {
		runHTTPServer(os.Args[2:])
		return
	}
	args := reorderFlags(os.Args[1:])
	idx := firstPositional(args)
	if idx == -1 {
		printUsage()
		os.Exit(1)
	}

	switch args[idx] {
	case "match":
		runMatch(args)
	case "list":
		runList()
	case "help", "-h", "--help":
		printUsage()
	default:
		runScore(args)
	}
}

func runHTTPServer(args []string) {
	addr := ":8080"
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-addr" || args[i] == "--addr" {
			addr = args[i+1]
			break
		}
	}
	cfg := server.Config{Addr: addr}
	fmt.Fprintf(os.Stdout, "str-sim server listening on %s\n", server.FormatAddr(addr))
	if err := server.ListenAndServe(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: str-sim <algo> <a> <b> [-p N]
       str-sim match <algo> <a> <b> <threshold>
       str-sim list

Commands:
  <algo> <a> <b>    Compute similarity between strings a and b
  match             Check if similarity meets threshold (prints true/false)
  list              List all supported algorithms

Algorithms: %s
`, strings.Join(sim.Algorithms(), ", "))
}

func runList() {
	for _, a := range sim.Algorithms() {
		fmt.Println(a)
	}
}

func runScore(args []string) {
	fs := flag.NewFlagSet("score", flag.ContinueOnError)
	prec := fs.Int("p", 6, "decimal precision for the printed score")
	if err := fs.Parse(args); err != nil {
		fatal("%v", err)
	}
	if fs.NArg() < 3 {
		fatal("score requires <algo> <a> <b>")
	}
	algo, a, b := fs.Arg(0), fs.Arg(1), fs.Arg(2)
	s, err := sim.Similarity(a, b, algo)
	if err != nil {
		fatal("%v", err)
	}
	fmt.Printf("%.*f\n", *prec, s)
}

func runMatch(args []string) {
	fs := flag.NewFlagSet("match", flag.ContinueOnError)
	if err := fs.Parse(args); err != nil {
		fatal("%v", err)
	}
	if fs.NArg() < 5 {
		fatal("match requires match <algo> <a> <b> <threshold>")
	}
	algo, a, b, thrStr := fs.Arg(1), fs.Arg(2), fs.Arg(3), fs.Arg(4)
	thr, err := strconv.ParseFloat(thrStr, 64)
	if err != nil {
		fatal("invalid threshold %q: %v", thrStr, err)
	}
	ok, err := sim.Match(a, b, algo, thr)
	if err != nil {
		fatal("%v", err)
	}
	fmt.Println(ok)
}
