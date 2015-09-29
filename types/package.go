// This file was generated by nomdl/codegen and then had references to this package (types) removed by hand. The $type field of Package was also manually set to the TypeRef that describes a Package directly.

package types

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

// Package

// PackageDef defines a Package object, which represents a Noms type package.
// Dependencies is a set of refs that point to other type packages required to resolve all the types in this pacakge.
// NamedTypes is a lookup table for types defined in this package. These should all be EnumKind or StructKind. When traversing the definition of a given type, you may run into a TypeRef that IsUnresolved(). In that case, look it up by name in the NamedTypes of the appropriate package.
type PackageDef struct {
	Dependencies SetOfRefOfPackageDef
	NamedTypes   MapOfStringToTypeRefDef
}

type Package struct {
	m Map
}

func NewPackage() Package {
	return Package{NewMap(
		NewString("$name"), NewString("Package"),
		NewString("$type"), __typeRefOfPackage(),
		NewString("Dependencies"), NewSet(),
		NewString("NamedTypes"), NewMap(),
	)}
}

func (def PackageDef) New() Package {
	return Package{
		NewMap(
			NewString("$name"), NewString("Package"),
			NewString("$type"), __typeRefOfPackage(),
			NewString("Dependencies"), def.Dependencies.New().NomsValue(),
			NewString("NamedTypes"), def.NamedTypes.New().NomsValue(),
		)}
}

func (self Package) Def() PackageDef {
	return PackageDef{
		SetOfRefOfPackageFromVal(self.m.Get(NewString("Dependencies"))).Def(),
		MapOfStringToTypeRefFromVal(self.m.Get(NewString("NamedTypes"))).Def(),
	}
}

// Creates and returns a Noms Value that describes Package.
func __typeRefOfPackage() TypeRef {
	return MakeStructTypeRef("Package",
		[]Field{
			Field{"Dependencies", MakeCompoundTypeRef("", SetKind, MakeCompoundTypeRef("", RefKind, MakeTypeRef("Package", ref.Ref{}))), false},
			Field{"NamedTypes", MakeCompoundTypeRef("", MapKind, MakePrimitiveTypeRef(StringKind), MakePrimitiveTypeRef(TypeRefKind)), false},
		},
		Choices{})

}

func PackageFromVal(val Value) Package {
	// TODO: Validate here
	return Package{val.(Map)}
}

func (self Package) NomsValue() Value {
	return self.m
}

func (self Package) Equals(other Package) bool {
	return self.m.Equals(other.m)
}

func (self Package) Ref() ref.Ref {
	return self.m.Ref()
}

func (self Package) Type() TypeRef {
	return self.m.Get(NewString("$type")).(TypeRef)
}

func (self Package) Dependencies() SetOfRefOfPackage {
	return SetOfRefOfPackageFromVal(self.m.Get(NewString("Dependencies")))
}

func (self Package) SetDependencies(val SetOfRefOfPackage) Package {
	return Package{self.m.Set(NewString("Dependencies"), val.NomsValue())}
}

func (self Package) NamedTypes() MapOfStringToTypeRef {
	return MapOfStringToTypeRefFromVal(self.m.Get(NewString("NamedTypes")))
}

func (self Package) SetTypes(val MapOfStringToTypeRef) Package {
	return Package{self.m.Set(NewString("NamedTypes"), val.NomsValue())}
}

// SetOfRefOfPackage

type SetOfRefOfPackage struct {
	s Set
}

type SetOfRefOfPackageDef map[ref.Ref]bool

func NewSetOfRefOfPackage() SetOfRefOfPackage {
	return SetOfRefOfPackage{NewSet()}
}

func (def SetOfRefOfPackageDef) New() SetOfRefOfPackage {
	l := make([]Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = Ref{R: d}
		i++
	}
	return SetOfRefOfPackage{NewSet(l...)}
}

func (s SetOfRefOfPackage) Def() SetOfRefOfPackageDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v Value) bool {
		def[v.Ref()] = true
		return false
	})
	return def
}

func SetOfRefOfPackageFromVal(p Value) SetOfRefOfPackage {
	return SetOfRefOfPackage{p.(Set)}
}

func (s SetOfRefOfPackage) NomsValue() Value {
	return s.s
}

func (s SetOfRefOfPackage) Equals(p SetOfRefOfPackage) bool {
	return s.s.Equals(p.s)
}

func (s SetOfRefOfPackage) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfRefOfPackage) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfPackage) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfPackage) Has(p RefOfPackage) bool {
	return s.s.Has(p.NomsValue())
}

type SetOfRefOfPackageIterCallback func(p RefOfPackage) (stop bool)

func (s SetOfRefOfPackage) Iter(cb SetOfRefOfPackageIterCallback) {
	s.s.Iter(func(v Value) bool {
		return cb(RefOfPackageFromVal(v))
	})
}

type SetOfRefOfPackageIterAllCallback func(p RefOfPackage)

func (s SetOfRefOfPackage) IterAll(cb SetOfRefOfPackageIterAllCallback) {
	s.s.IterAll(func(v Value) {
		cb(RefOfPackageFromVal(v))
	})
}

type SetOfRefOfPackageFilterCallback func(p RefOfPackage) (keep bool)

func (s SetOfRefOfPackage) Filter(cb SetOfRefOfPackageFilterCallback) SetOfRefOfPackage {
	ns := NewSetOfRefOfPackage()
	s.IterAll(func(v RefOfPackage) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfRefOfPackage) Insert(p ...RefOfPackage) SetOfRefOfPackage {
	return SetOfRefOfPackage{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfRefOfPackage) Remove(p ...RefOfPackage) SetOfRefOfPackage {
	return SetOfRefOfPackage{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfRefOfPackage) Union(others ...SetOfRefOfPackage) SetOfRefOfPackage {
	return SetOfRefOfPackage{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfRefOfPackage) Subtract(others ...SetOfRefOfPackage) SetOfRefOfPackage {
	return SetOfRefOfPackage{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfRefOfPackage) Any() RefOfPackage {
	return RefOfPackageFromVal(s.s.Any())
}

func (s SetOfRefOfPackage) fromStructSlice(p []SetOfRefOfPackage) []Set {
	r := make([]Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfPackage) fromElemSlice(p []RefOfPackage) []Value {
	r := make([]Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

// RefOfPackage

type RefOfPackage struct {
	r ref.Ref
}

func NewRefOfPackage(r ref.Ref) RefOfPackage {
	return RefOfPackage{r}
}

func (r RefOfPackage) Ref() ref.Ref {
	return r.r
}

func (r RefOfPackage) Equals(other RefOfPackage) bool {
	return r.Ref() == other.Ref()
}

func (r RefOfPackage) NomsValue() Value {
	return Ref{R: r.r}
}

func RefOfPackageFromVal(p Value) RefOfPackage {
	return RefOfPackage{p.(Ref).Ref()}
}

func (r RefOfPackage) GetValue(cs chunks.ChunkSource) Package {
	return PackageFromVal(ReadValue(r.r, cs))
}

func (r RefOfPackage) SetValue(val Package, cs chunks.ChunkSink) RefOfPackage {
	ref := WriteValue(val.NomsValue(), cs)
	return RefOfPackage{ref}
}

// MapOfStringToTypeRef

type MapOfStringToTypeRef struct {
	m Map
}

type MapOfStringToTypeRefDef map[string]TypeRef

func NewMapOfStringToTypeRef() MapOfStringToTypeRef {
	return MapOfStringToTypeRef{NewMap()}
}

func (def MapOfStringToTypeRefDef) New() MapOfStringToTypeRef {
	kv := make([]Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, NewString(k), v)
	}
	return MapOfStringToTypeRef{NewMap(kv...)}
}

func (self MapOfStringToTypeRef) Def() MapOfStringToTypeRefDef {
	def := make(map[string]TypeRef)
	self.m.Iter(func(k, v Value) bool {
		def[k.(String).String()] = v.(TypeRef)
		return false
	})
	return def
}

func MapOfStringToTypeRefFromVal(p Value) MapOfStringToTypeRef {
	// TODO: Validate here
	return MapOfStringToTypeRef{p.(Map)}
}

func (m MapOfStringToTypeRef) NomsValue() Value {
	return m.m
}

func (m MapOfStringToTypeRef) Equals(p MapOfStringToTypeRef) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToTypeRef) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToTypeRef) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToTypeRef) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToTypeRef) Has(p string) bool {
	return m.m.Has(NewString(p))
}

func (m MapOfStringToTypeRef) Get(p string) TypeRef {
	return m.m.Get(NewString(p)).(TypeRef)
}

func (m MapOfStringToTypeRef) Set(k string, v TypeRef) MapOfStringToTypeRef {
	return MapOfStringToTypeRef{m.m.Set(NewString(k), v)}
}

// TODO: Implement SetM?

func (m MapOfStringToTypeRef) Remove(p string) MapOfStringToTypeRef {
	return MapOfStringToTypeRef{m.m.Remove(NewString(p))}
}

type MapOfStringToTypeRefIterCallback func(k string, v TypeRef) (stop bool)

func (m MapOfStringToTypeRef) Iter(cb MapOfStringToTypeRefIterCallback) {
	m.m.Iter(func(k, v Value) bool {
		return cb(k.(String).String(), v.(TypeRef))
	})
}

type MapOfStringToTypeRefIterAllCallback func(k string, v TypeRef)

func (m MapOfStringToTypeRef) IterAll(cb MapOfStringToTypeRefIterAllCallback) {
	m.m.IterAll(func(k, v Value) {
		cb(k.(String).String(), v.(TypeRef))
	})
}

type MapOfStringToTypeRefFilterCallback func(k string, v TypeRef) (keep bool)

func (m MapOfStringToTypeRef) Filter(cb MapOfStringToTypeRefFilterCallback) MapOfStringToTypeRef {
	nm := NewMapOfStringToTypeRef()
	m.IterAll(func(k string, v TypeRef) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}
