package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/go-fed/activity/tools/exp/codegen"
	"github.com/go-fed/activity/tools/exp/props"
	"github.com/go-fed/activity/tools/exp/types"
)

func main() {
	serializeIRIFn := *codegen.NewFunction(
		"test",
		"SerializeIRI",
		[]jen.Code{jen.Id("u").Op("*").Qual("url", "URL")},
		[]jen.Code{jen.List(jen.Interface(), jen.Error())},
		[]jen.Code{jen.Empty()})
	deserializeIRIFn := *codegen.NewFunction(
		"test",
		"DeserializeIRI",
		[]jen.Code{jen.Interface()},
		[]jen.Code{jen.List(jen.Op("*").Qual("url", "URL"), jen.Bool(), jen.Error())},
		[]jen.Code{jen.Empty()})
	lessIRIFn := *codegen.NewFunction(
		"test",
		"LessIRI",
		[]jen.Code{jen.Id("lhs").Id("rhs").Op("*").Qual("url", "URL")},
		[]jen.Code{jen.Bool()},
		[]jen.Code{jen.Empty()})
	serializeIntFn := *codegen.NewFunction(
		"test",
		"SerializeInt",
		[]jen.Code{jen.Id("i").Op("*").Int()},
		[]jen.Code{jen.List(jen.Interface(), jen.Error())},
		[]jen.Code{jen.Empty()})
	deserializeIntFn := *codegen.NewFunction(
		"test",
		"DeserializeInt",
		[]jen.Code{jen.Interface()},
		[]jen.Code{jen.List(jen.Int(), jen.Bool(), jen.Error())},
		[]jen.Code{jen.Empty()})
	lessIntFn := *codegen.NewFunction(
		"test",
		"LessInt",
		[]jen.Code{jen.Id("lhs").Id("rhs").Int()},
		[]jen.Code{jen.Bool()},
		[]jen.Code{jen.Empty()})
	x := props.NewFunctionalPropertyGenerator(
		"test",
		props.Identifier{
			LowerName: "testFunctional",
			CamelName: "TestFunctional",
		},
		[]props.Kind{
			{
				Name: props.Identifier{
					LowerName: "iri",
					CamelName: "IRI",
				},
				ConcreteKind:          "*url.URL",
				Nilable:               true,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIRIFn,
				DeserializeFn:         deserializeIRIFn,
				LessFn:                lessIRIFn,
			},
		},
		true)
	y := props.NewFunctionalPropertyGenerator(
		"test",
		props.Identifier{
			LowerName: "testFunctionalNonnil",
			CamelName: "TestFunctionalNonil",
		},
		[]props.Kind{
			{
				Name: props.Identifier{
					LowerName: "number",
					CamelName: "Number",
				},
				ConcreteKind:          "int",
				Nilable:               false,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIntFn,
				DeserializeFn:         deserializeIntFn,
				LessFn:                lessIntFn,
			},
		},
		true)
	z := props.NewFunctionalPropertyGenerator(
		"test",
		props.Identifier{
			LowerName: "testFunctionalMultiType",
			CamelName: "TestFunctionalMultiType",
		},
		[]props.Kind{
			{
				Name: props.Identifier{
					LowerName: "iri",
					CamelName: "IRI",
				},
				ConcreteKind:          "*url.URL",
				Nilable:               true,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIRIFn,
				DeserializeFn:         deserializeIRIFn,
				LessFn:                lessIRIFn,
			},
			{
				Name: props.Identifier{
					LowerName: "number",
					CamelName: "Number",
				},
				ConcreteKind:          "int",
				Nilable:               false,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIntFn,
				DeserializeFn:         deserializeIntFn,
				LessFn:                lessIntFn,
			},
		},
		true)
	zz := props.NewNonFunctionalPropertyGenerator(
		"test",
		props.Identifier{
			LowerName: "testNonFunctionalMultiType",
			CamelName: "TestNonFunctionalMultiType",
		},
		[]props.Kind{
			{
				Name: props.Identifier{
					LowerName: "iri",
					CamelName: "IRI",
				},
				ConcreteKind:          "*url.URL",
				Nilable:               true,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIRIFn,
				DeserializeFn:         deserializeIRIFn,
				LessFn:                lessIRIFn,
			},
			{
				Name: props.Identifier{
					LowerName: "number",
					CamelName: "Number",
				},
				ConcreteKind:          "int",
				Nilable:               false,
				HasNaturalLanguageMap: false,
				SerializeFn:           serializeIntFn,
				DeserializeFn:         deserializeIntFn,
				LessFn:                lessIntFn,
			},
		},
		true)
	t1, err := types.NewTypeGenerator("test", "TestType", "TestType is a test type", []types.Property{x, y, z, zz}, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", x.Definition().Definition())
	fmt.Printf("%#v\n\n", y.Definition().Definition())
	fmt.Printf("%#v\n\n", z.Definition().Definition())
	s, t := zz.Definitions()
	fmt.Printf("%#v\n\n%#v\n\n", s.Definition(), t.Definition())
	fmt.Printf("%#v\n\n", types.TypeInterface("test").Definition())
	fmt.Printf("%#v\n\n", t1.Definition().Definition())
}
