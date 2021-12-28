package day7

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day7", "day-7-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	circuit := NewCircuit()
	circuit.ApplyInstructions(lines)
	ans := int(*circuit.Wire("a").Value)
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 1 (%dms): Value \"a\" = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	circuit := NewCircuit()
	circuit.ApplyInstructions(lines)
	aVal := *circuit.Wire("a").Value
	circuit.WireValue("b", aVal)
	ans := int(*circuit.Wire("a").Value)
	end := time.Now().UnixMilli()
	log.Printf("Day 7, Part 2 (%dms): Value \"a\" = %d", end-start, ans)
	return ans
}

func NewCircuit() *Circuit {
	c := Circuit{}
	c.wires = make(map[string]*Wire, 10)
	c.gates = make(map[string][]*Gate, 10)
	return &c
}

type Circuit struct {
	wires map[string]*Wire
	gates map[string][]*Gate
}

func (c *Circuit) Wire(name string) *Wire {
	if _, ok := c.wires[name]; !ok {
		wire := &Wire{}
		c.wires[name] = wire
	}
	return c.wires[name]
}

func (c *Circuit) WireValue(name string, value uint16) {
	// set value on wire (if changed)
	wire := c.Wire(name)
	if wire.Value != nil && *wire.Value == value {
		return
	}
	wire.Value = &value

	// evaluate any gates associated with wire
	if gates, ok := c.gates[name]; ok {
		for _, gate := range gates {
			gate.Evaluate(c)
		}
	}
}

func (c *Circuit) ApplyInstructions(instructions []string) {
	for _, s := range instructions {
		c.ApplyInstruction(s)
	}
}

func (c *Circuit) ApplyInstruction(s string) {
	pieces := strings.Split(s, " ")

	if pieces[1] == "AND" || pieces[1] == "OR" {
		g := &Gate{
			gateType: pieces[1],
			input1:   c.toWire(pieces[0]),
			input2:   c.toWire(pieces[2]),
			output:   pieces[4],
		}
		c.addGateMapping(pieces[0], g)
		c.addGateMapping(pieces[2], g)
		g.Evaluate(c)
		return
	}

	if pieces[1] == "LSHIFT" || pieces[1] == "RSHIFT" {
		g := &Gate{
			gateType: pieces[1],
			input1:   c.toWire(pieces[0]),
			shiftAmt: utils.Number(pieces[2]),
			output:   pieces[4],
		}
		c.addGateMapping(pieces[0], g)
		g.Evaluate(c)
		return
	}

	if pieces[0] == "NOT" {
		g := &Gate{
			gateType: pieces[0],
			input1:   c.toWire(pieces[1]),
			output:   pieces[3],
		}
		c.addGateMapping(pieces[1], g)
		g.Evaluate(c)
		return
	}

	num, err := strconv.Atoi(pieces[0])
	if err != nil {
		g := &Gate{
			gateType: pieces[1],
			input1:   c.toWire(pieces[0]),
			output:   pieces[2],
		}
		c.addGateMapping(pieces[0], g)
		g.Evaluate(c)
		return
	}
	c.WireValue(pieces[2], uint16(num))
}

func (c *Circuit) toWire(name string) *Wire {
	num, err := strconv.Atoi(name)
	if err == nil {
		value := uint16(num)
		return &Wire{Value: &value}
	}
	return c.Wire(name)
}

func (c *Circuit) addGateMapping(name string, gate *Gate) {
	gates, ok := c.gates[name]
	if !ok {
		gates = make([]*Gate, 0, 5)
	}
	c.gates[name] = append(gates, gate)
}

type Wire struct {
	Value *uint16
}

type Gate struct {
	gateType string
	input1   *Wire
	input2   *Wire
	shiftAmt int
	output   string
}

func (g *Gate) Evaluate(c *Circuit) {
	switch g.gateType {
	case "AND":
		if g.input1.Value != nil && g.input2.Value != nil {
			c.WireValue(g.output, *g.input1.Value&*g.input2.Value)
		}
	case "OR":
		if g.input1.Value != nil && g.input2.Value != nil {
			c.WireValue(g.output, *g.input1.Value|*g.input2.Value)
		}
	case "NOT":
		if g.input1.Value != nil {
			c.WireValue(g.output, ^*g.input1.Value)
		}
	case "LSHIFT":
		if g.input1.Value != nil {
			c.WireValue(g.output, *g.input1.Value<<g.shiftAmt)
		}
	case "RSHIFT":
		if g.input1.Value != nil {
			c.WireValue(g.output, *g.input1.Value>>g.shiftAmt)
		}
	default:
		if g.input1.Value != nil {
			c.WireValue(g.output, *g.input1.Value)
		}
	}
}
