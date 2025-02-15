/*
Copyright © 2021 windvalley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package log

import (
	"encoding/json"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

// entry ...
type entry struct {
	Logger *Logger
	Data   Fields
}

func newEntry(logger *Logger) *entry {
	return &entry{
		Logger: logger,
		Data:   make(Fields),
	}
}

func (e *entry) print() {
	e.Data["time"] = time.Now().Format(timeFormat)

	entry := ""
	if e.Logger.JSONFormat {
		entryByte, _ := json.Marshal(e.Data)
		entry = string(entryByte)
	} else {
		if len(e.Data) <= 3 {
			for k, v := range e.Data {
				entry += fmt.Sprintf("%v=%v ", k, v)
			}
		} else {
			entry = fmt.Sprintf("%q,%q,%q,\"%s\"",
				e.Data["hostname"],
				e.Data["status"],
				e.Data["time"],
				e.Data["output"],
			)
		}
	}

	fmt.Fprintln(e.Logger.Out, entry)
}

// Debugf ...
func (e *entry) Debugf(format string, args ...interface{}) {
	if !e.Logger.Verbose {
		return
	}

	e.Data["level"] = "debug"

	msg := fmt.Sprintf(format, args...)
	e.Data["msg"] = msg

	e.print()
}

// Infof ...
func (e *entry) Infof(format string, args ...interface{}) {
	e.Data["level"] = "info"

	msg := fmt.Sprintf(format, args...)
	e.Data["msg"] = msg

	e.print()
}

// Warnf ...
func (e *entry) Warnf(format string, args ...interface{}) {
	e.Data["level"] = "warn"

	msg := fmt.Sprintf(format, args...)
	e.Data["msg"] = msg

	e.print()
}

// Errorf ...
func (e *entry) Errorf(format string, args ...interface{}) {
	e.Data["level"] = "error"

	msg := fmt.Sprintf(format, args...)
	e.Data["msg"] = msg

	e.print()
}
