package pepcli

// represents all logic regarding cli control for
// peptide analyser program

import (
	"fmt"
	"os"
	"peptide-analyse/fastaproc"
	peptideseq "peptide-analyse/peptide-seq"
	"sync"

	"github.com/urfave/cli/v2"
)

var minRange, maxRange float64

// main function that is run on cli start
func CliAction(cctx *cli.Context) error {
	files := cctx.StringSlice("files")
	results := make(chan peptideseq.PeptideSeq, len(files)*2)
	errChan := make(chan error, len(files))

	wgCollect := new(sync.WaitGroup)
	wgRecieve := new(sync.WaitGroup)

	// ok to pass file name here, we checked if it exists before
	for _, file := range files {
		filter := fastaproc.NewFilter(wgCollect, results)

		wgCollect.Add(1)
		go filter.FilterWithRange(file, minRange, maxRange, errChan)
	}

	wgRecieve.Add(1)
	go fastaproc.WritePeptideSequencesToFile("out.fasta", results, wgRecieve)

	// wait for collecting to end
	wgCollect.Wait()
	// no more sequences to read, close channel
	close(results)
	// wait for write to file to finish
	wgRecieve.Wait()

	close(errChan)
	err := <-errChan
	if err != nil {
		return err
	}

	return nil
}

func CheckBefore(cctx *cli.Context) error {
	if minRange > maxRange {
		return fmt.Errorf("Invalid range given: start %f - stop %f", minRange, maxRange)
	}

	files := cctx.StringSlice("files")
	if len(files) == 0 {
		return fmt.Errorf("no files specified")
	}

	// check if files exist and is not a directory
	for _, f := range files {
		info, err := os.Stat(f)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s is a directory, not a file", info.Name())
		}
	}

	return nil
}
