package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/suzuki/go-cmnt-eol-lint/src/env"
	"github.com/suzuki/go-cmnt-eol-lint/src/formatter"
	"github.com/suzuki/go-cmnt-eol-lint/src/linter"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	ctx = env.WithConfig(ctx)

	eg, ctx := errgroup.WithContext(ctx)

	var results []*linter.Result
	paths := os.Args[1:]
	if len(paths) < 1 {
		usage()
		os.Exit(1)
	}

	for _, path := range paths {
		path := path // https://golang.org/doc/faq#closures_and_goroutines

		st, err := os.Stat(path)
		if st.IsDir() {
			continue
		}
		if err != nil {
			printError(err)
		}

		eg.Go(func() error {
			r, err := doLint(ctx, path)
			if err != nil {
				return err
			}

			results = append(results, r...)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		printError(err)
		os.Exit(1)
	}

	formatter := formatter.NewDefaultFormatter()
	output, err := formatter.Format(results)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	if output != "" {
		fmt.Print(output)
		os.Exit(1)
	}

	// There is no problem.
	os.Exit(0)
}

func doLint(ctx context.Context, path string) ([]*linter.Result, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	linter := linter.New(env.GetConfig(ctx))
	results, err := linter.LintFile(path)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func usage() {
	fmt.Printf("%s filepath [filepath...]\n", path.Base(os.Args[0]))
}

func printError(err error) {
	fmt.Printf("Error: %s", err)
}
