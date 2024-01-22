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
var outputFileName string
var wComments bool

// main function that is run on cli start
func CliAction(cctx *cli.Context) error {
	files := cctx.StringSlice("files")
	results := make(chan peptideseq.PeptideSeq, len(files)*2)
	// make enough space for all collector and reciever routines
	errChan := make(chan error, len(files)+1)

	wgCollect := new(sync.WaitGroup)
	wgRecieve := new(sync.WaitGroup)

	var outputFunc fastaproc.OutputFunc
	var outFile *os.File
	var err error

	if outputFileName == "" {
		outFile = nil
		// write directly to stdout
		outputFunc = func(ps peptideseq.PeptideSeq) error {
			fmt.Print(ps.Write(wComments))
			return nil
		}
	} else {
		// create file here in order to avoid open/close on
		// every function call
		outFile, err = os.Create(outputFileName)
		if err != nil {
			return err
		}

		// write to file
		outputFunc = func(ps peptideseq.PeptideSeq) error {
			_, e := outFile.WriteString(ps.Write(wComments))
			if e != nil {
				return e
			}
			return nil
		}
	}
	defer outFile.Close()

	// ok to pass file name here, we checked if it exists before
	for _, file := range files {
		filter := fastaproc.NewFilter(wgCollect, results)

		wgCollect.Add(1)
		go filter.FilterWithRange(file, minRange, maxRange, errChan)
	}

	wgRecieve.Add(1)
	go fastaproc.WritePeptideSequences(outputFunc, results, errChan, wgRecieve)

	// wait for collecting to end
	wgCollect.Wait()
	// no more sequences to read, close channel
	close(results)
	// wait for write to file to finish
	wgRecieve.Wait()

	close(errChan)
	err = <-errChan
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
