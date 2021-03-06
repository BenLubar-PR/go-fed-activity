package codegen

import (
	"github.com/dave/jennifer/jen"
)

// join appends a bunch of Go Code together, each on their own line.
func join(s []jen.Code) *jen.Statement {
	r := jen.Empty()
	for i, stmt := range s {
		if i > 0 {
			r.Line()
		}
		r.Add(stmt)
	}
	return r
}

// Struct defines a struct-based type, its functions, and its methods for Go
// code generation.
type Struct struct {
	comment      jen.Code
	name         string
	methods      map[string]*Method
	constructors map[string]*Function
	members      []jen.Code
}

// NewStruct creates a new commented Struct type.
func NewStruct(comment jen.Code,
	name string,
	methods []*Method,
	constructors []*Function,
	members []jen.Code) *Struct {
	s := &Struct{
		comment:      comment,
		name:         name,
		methods:      make(map[string]*Method, len(methods)),
		constructors: make(map[string]*Function, len(constructors)),
		members:      members,
	}
	for _, m := range methods {
		s.methods[m.Name()] = m
	}
	for _, c := range constructors {
		s.constructors[c.Name()] = c
	}
	return s
}

// Definition generates the Go code required to define and implement this
// struct, its methods, and its functions.
func (s *Struct) Definition() jen.Code {
	comment := jen.Empty()
	if s.comment != nil {
		comment = jen.Empty().Add(s.comment).Line()
	}
	def := comment.Type().Id(s.name).Struct(
		join(s.members),
	)
	for _, c := range s.constructors {
		def = def.Line().Line().Add(c.Definition())
	}
	for _, m := range s.methods {
		def = def.Line().Line().Add(m.Definition())
	}
	return def
}

// Method obtains the Go code to be generated for the method with a specific
// name. Panics if no such method exists.
func (s *Struct) Method(name string) *Method {
	return s.methods[name]
}

// Constructors obtains the Go code to be generated for the function with a
// specific name. Panics if no such function exists.
func (s *Struct) Constructors(name string) *Function {
	return s.constructors[name]
}
