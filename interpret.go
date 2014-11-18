/*
	Brainfuck interpreter written in GoLang
	Version V0.1.0
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	tokenInc   = 43
	tokenDec   = 45
	tokenPrev  = 60
	tokenNext  = 62
	tokenOut   = 46
	tokenIn    = 44
	tokenOpen  = 91
	tokenClose = 93
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("brainfuck interpreter, try: go run interpreter.go filename.bf")
		os.Exit(1)
	}

	filename := os.Args[1]
	fileBuff, _ := ioutil.ReadFile(filename)
	interpret(fileBuff)
}

func interpret(fileBuff []byte) {
	bracketStack := make([]int, 1000)
	bracketStackPointer := 0
	memory := make([]int, 1000)
	bufferPos := 0
	pointerPos := 0
	waitNextBracket := false
	waitNextBracketSkipCount := 0

	length := len(fileBuff)
	for bufferPos < length {
		if waitNextBracket {
			if fileBuff[bufferPos] != tokenClose || waitNextBracketSkipCount > 0 {
				if fileBuff[bufferPos] == tokenOpen {
					waitNextBracketSkipCount++
				}
				if fileBuff[bufferPos] == tokenClose {
					waitNextBracketSkipCount--
				}
			} else {
				waitNextBracket = false
			}

			bufferPos++
			continue
		}

		switch fileBuff[bufferPos] {
		case tokenInc:
			memory[pointerPos]++
		case tokenDec:
			memory[pointerPos]--
		case tokenNext:
			pointerPos++
		case tokenPrev:
			pointerPos--
		case tokenIn:
		case tokenOut:
			fmt.Print(string(memory[pointerPos]))
		case tokenOpen:
			if memory[pointerPos] <= 0 {
				waitNextBracket = true
			} else {
				bracketStack[bracketStackPointer] = bufferPos
				bracketStackPointer++
			}
		case tokenClose:
			bracketStackPointer--
			if memory[pointerPos] <= 0 {
				bracketStack[bracketStackPointer] = 0
			} else {
				bufferPos = bracketStack[bracketStackPointer] - 1
				bracketStack[bracketStackPointer] = 0
			}
		}

		bufferPos++
	}
}
