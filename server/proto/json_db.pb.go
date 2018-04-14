// Code generated by protoc-gen-go. DO NOT EDIT.
// source: json_db.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	json_db.proto
	swapi.proto

It has these top-level messages:
	DB
	Film
	GetFilmRequest
	GetFilmResponse
	ListFilmsRequest
	ListFilmsResponse
	Vehicle
	GetVehicleRequest
	GetVehicleResponse
	ListVehiclesRequest
	ListVehiclesResponse
	Starship
	GetStarshipRequest
	GetStarshipResponse
	ListStarshipsRequest
	ListStarshipsResponse
	Species
	GetSpeciesRequest
	GetSpeciesResponse
	ListSpeciesRequest
	ListSpeciesResponse
	Planet
	GetPlanetRequest
	GetPlanetResponse
	ListPlanetsRequest
	ListPlanetsResponse
	Person
	GetPersonRequest
	GetPersonResponse
	ListPeopleRequest
	ListPeopleResponse
	StarshipAction
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type DB struct {
	Films     []*Film     `protobuf:"bytes,1,rep,name=films" json:"films,omitempty"`
	People    []*Person   `protobuf:"bytes,2,rep,name=people" json:"people,omitempty"`
	Planets   []*Planet   `protobuf:"bytes,3,rep,name=planets" json:"planets,omitempty"`
	Species   []*Species  `protobuf:"bytes,4,rep,name=species" json:"species,omitempty"`
	Starships []*Starship `protobuf:"bytes,5,rep,name=starships" json:"starships,omitempty"`
	Vehicles  []*Vehicle  `protobuf:"bytes,6,rep,name=vehicles" json:"vehicles,omitempty"`
}

func (m *DB) Reset()                    { *m = DB{} }
func (m *DB) String() string            { return proto1.CompactTextString(m) }
func (*DB) ProtoMessage()               {}
func (*DB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DB) GetFilms() []*Film {
	if m != nil {
		return m.Films
	}
	return nil
}

func (m *DB) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func (m *DB) GetPlanets() []*Planet {
	if m != nil {
		return m.Planets
	}
	return nil
}

func (m *DB) GetSpecies() []*Species {
	if m != nil {
		return m.Species
	}
	return nil
}

func (m *DB) GetStarships() []*Starship {
	if m != nil {
		return m.Starships
	}
	return nil
}

func (m *DB) GetVehicles() []*Vehicle {
	if m != nil {
		return m.Vehicles
	}
	return nil
}

func init() {
	proto1.RegisterType((*DB)(nil), "swapi.v1.DB")
}

func init() { proto1.RegisterFile("json_db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0xbf, 0x4a, 0xc7, 0x30,
	0x10, 0xc0, 0x71, 0x6c, 0xed, 0x1f, 0xaf, 0x28, 0x7a, 0x53, 0x70, 0x12, 0x71, 0x28, 0x8a, 0xc1,
	0x3f, 0x6f, 0x50, 0xc4, 0x59, 0x22, 0x38, 0xb8, 0x48, 0x5b, 0x4f, 0x1a, 0x49, 0x9b, 0xd0, 0x2b,
	0xf5, 0x2d, 0x7c, 0x66, 0x21, 0x69, 0xad, 0xf8, 0x9b, 0x12, 0xee, 0xfb, 0xe1, 0x86, 0x83, 0xc3,
	0x4f, 0xb6, 0xc3, 0xdb, 0x7b, 0x23, 0xdd, 0x68, 0x27, 0x8b, 0x39, 0x7f, 0xd5, 0x4e, 0xcb, 0xf9,
	0xf6, 0xb4, 0x08, 0x3f, 0x3f, 0x3e, 0xff, 0x8e, 0x20, 0x7a, 0xa8, 0xf0, 0x02, 0x92, 0x0f, 0x6d,
	0x7a, 0x16, 0x7b, 0x67, 0x71, 0x59, 0xdc, 0x1d, 0xc9, 0x55, 0xcb, 0x47, 0x6d, 0x7a, 0x15, 0x22,
	0x96, 0x90, 0x3a, 0xb2, 0xce, 0x90, 0x88, 0x3c, 0x3b, 0xde, 0xd8, 0x13, 0x8d, 0x6c, 0x07, 0xb5,
	0x74, 0xbc, 0x84, 0xcc, 0x99, 0x7a, 0xa0, 0x89, 0x45, 0xbc, 0x43, 0x7d, 0x50, 0x2b, 0xc0, 0x2b,
	0xc8, 0xd8, 0x51, 0xab, 0x89, 0xc5, 0xbe, 0xb7, 0x27, 0x9b, 0x7d, 0x0e, 0x41, 0xad, 0x02, 0x6f,
	0xe0, 0x80, 0xa7, 0x7a, 0xe4, 0x4e, 0x3b, 0x16, 0x89, 0xe7, 0xf8, 0x87, 0x2f, 0x49, 0x6d, 0x08,
	0xaf, 0x21, 0x9f, 0xa9, 0xd3, 0xad, 0x21, 0x16, 0xe9, 0xff, 0xfd, 0x2f, 0xa1, 0xa8, 0x5f, 0x52,
	0x65, 0xaf, 0x89, 0xbf, 0x4c, 0x93, 0xfa, 0xe7, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x78,
	0xc0, 0xa0, 0x48, 0x01, 0x00, 0x00,
}
