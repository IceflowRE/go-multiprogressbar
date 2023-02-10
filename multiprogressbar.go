package multiprogressbar

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"sync"
)

type MultiProgressBar struct {
	curLine int
	bars    []*progressbar.ProgressBar
	guard   sync.Mutex
	output  io.Writer
}

func New() *MultiProgressBar {
	return &MultiProgressBar{
		curLine: 0,
		bars:    []*progressbar.ProgressBar{},
		guard:   sync.Mutex{},
		output:  os.Stdout,
	}
}

// Add a progress bar. This will change the writer of the progress bar.
// Do not change the writer afterwards!
// Not thread safe.
// Return the passed progress bar.
func (mpb *MultiProgressBar) Add(pBar *progressbar.ProgressBar) *progressbar.ProgressBar {
	progressbar.OptionSetWriter(&multiProgressBarWriter{
		MultiProgressBar: mpb,
		idx:              len(mpb.bars),
	})(pBar)
	mpb.bars = append(mpb.bars, pBar)
	return pBar
}

// Get returns the progressbar.Progressbar with the given index.
// Will panic if the index does not exist.
func (mpb *MultiProgressBar) Get(idx int) *progressbar.ProgressBar {
	return mpb.bars[idx]
}

func (mpb *MultiProgressBar) BarCount() int {
	return len(mpb.bars)
}

// RenderBlank calls RenderBlank on all progress bars.
// If an error is thrown, RenderBlank might not be called on all bars.
func (mpb *MultiProgressBar) RenderBlank() error {
	for _, pbar := range mpb.bars {
		err := pbar.RenderBlank()
		if err != nil {
			return err
		}
	}
}

// Finish calls Finish on all progress bars.
// If an error is thrown, Finish might not be called on all bars.
// This will also call End.
func (mpb *MultiProgressBar) Finish() error {
	for _, pbar := range mpb.bars {
		err := pbar.Finish()
		if err != nil {
			return err
		}
	}
	return mpb.End()
}

// End Move cursor to the end of the progress bars.
// Not thread safe.
func (mpb *MultiProgressBar) End() error {
	_, err := mpb.move(len(mpb.bars), mpb.output)
	return err
}

// Move cursor to the beginning of the current progressbar.
func (mpb *MultiProgressBar) move(id int, writer io.Writer) (int, error) {
	bias := mpb.curLine - id
	mpb.curLine = id
	if bias > 0 {
		// move up
		return fmt.Fprintf(writer, "\r\033[%dA", bias)
	} else if bias < 0 {
		// move down
		return fmt.Fprintf(writer, "\r\033[%dB", -bias)
	}
	return 0, nil
}

// personal note about Option. I do not like this pattern, but i will use it nonetheless, to stay close to schollz/progressbar.

// Option is the type all options need to adhere to
type Option func(p *MultiProgressBar)

// OptionSetWriter sets the output writer.
// Unknown behaviour if called while using the multi progress bar.
func OptionSetWriter(writer io.Writer) Option {
	return func(mpb *MultiProgressBar) {
		mpb.output = writer
	}
}

// io.Writer wrapper to know which progressbar wants to write.
type multiProgressBarWriter struct {
	*MultiProgressBar
	idx int
}

func (lw *multiProgressBarWriter) Write(p []byte) (n int, err error) {
	lw.guard.Lock()
	defer lw.guard.Unlock()
	n, err = lw.move(lw.idx, lw.output)
	if err != nil {
		return n, err
	}
	return lw.output.Write(p)
}
