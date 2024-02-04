package emitter

import (
	"os"
)

// Emitter keeps track of the generated code and outputs it.
type Emitter struct {
	fullPath string
	header   string
	code     string
}

func NewEmitter(fullPath string) *Emitter {
	return &Emitter{fullPath: fullPath}
}

func (e *Emitter) Emit(code string) {
	e.code += code
}

func (e *Emitter) EmitLine(code string) {
	e.code += code + "\n"
}

func (e *Emitter) HeaderLine(code string) {
	e.header += code + "\n"
}

func (e *Emitter) WriteFile() error {
	file, err := os.Create(e.fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := e.header + e.code
	_, err = file.WriteString(content)
	return err
}
