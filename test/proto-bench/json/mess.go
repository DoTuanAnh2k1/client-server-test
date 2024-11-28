package json

import (
	_ "github.com/mailru/easyjson"
	_ "github.com/mailru/easyjson/gen"
)
// Enum example
type TestEnum int32

const (
	TestEnum_UNKNOWN      TestEnum = 0
	TestEnum_OPTION_ONE   TestEnum = 1
	TestEnum_OPTION_TWO   TestEnum = 2
	TestEnum_OPTION_THREE TestEnum = 3
)

// Nested message example
type NestedMessage struct {
	NestedID     int32     `json:"1"`
	NestedName   string    `json:"2"`
	NestedValues []float32 `json:"3"`
}

// Main struct with 100+ fields
// easyjson:json
type BenchmarkMessage struct {
	// Primitive types
	Field1  int32   `json:"1"`
	Field2  int64   `json:"2"`
	Field3  uint32  `json:"3"`
	Field4  uint64  `json:"4"`
	Field5  int32   `json:"5"`
	Field6  int64   `json:"6"`
	Field7  uint32  `json:"7"`
	Field8  uint64  `json:"8"`
	Field9  int32   `json:"9"`
	Field10 int64   `json:"10"`
	Field11 float32 `json:"11"`
	Field12 float64 `json:"12"`
	Field13 bool    `json:"13"`
	Field14 string  `json:"14"`

	// Enum field
	Field15 TestEnum `json:"15"`

	// Nested message
	Field16 NestedMessage `json:"16"`

	// Repeated fields
	Field17 []int32  `json:"17"`
	Field18 []string `json:"18"`

	// Map field
	Field19 map[string]int32 `json:"19"`

	// Additional fields
	Field20 string                  `json:"20"`
	Field21 int32                   `json:"21"`
	Field22 int64                   `json:"22"`
	Field23 uint32                  `json:"23"`
	Field24 uint64                  `json:"24"`
	Field25 int32                   `json:"25"`
	Field26 int64                   `json:"26"`
	Field27 bool                    `json:"27"`
	Field28 float32                 `json:"28"`
	Field29 float64                 `json:"29"`
	Field30 uint32                  `json:"30"`
	Field31 uint64                  `json:"31"`
	Field32 int32                   `json:"32"`
	Field33 int64                   `json:"33"`
	Field34 string                  `json:"34"`
	Field35 []byte                  `json:"35"`
	Field36 []NestedMessage         `json:"36"`
	Field37 map[int32]NestedMessage `json:"37"`

	// Remaining fields up to 100+
	Field38 int32         `json:"38"`
	Field39 int64         `json:"39"`
	Field40 uint32        `json:"40"`
	Field41 uint64        `json:"41"`
	Field42 int32         `json:"42"`
	Field43 int64         `json:"43"`
	Field44 uint32        `json:"44"`
	Field45 uint64        `json:"45"`
	Field46 int32         `json:"46"`
	Field47 int64         `json:"47"`
	Field48 float32       `json:"48"`
	Field49 float64       `json:"49"`
	Field50 bool          `json:"50"`
	Field51 string        `json:"51"`
	Field52 string        `json:"52"`
	Field53 string        `json:"53"`
	Field54 string        `json:"54"`
	Field55 string        `json:"55"`
	Field56 int32         `json:"56"`
	Field57 int64         `json:"57"`
	Field58 uint32        `json:"58"`
	Field59 uint64        `json:"59"`
	Field60 int32         `json:"60"`
	Field61 int64         `json:"61"`
	Field62 bool          `json:"62"`
	Field63 float32       `json:"63"`
	Field64 float64       `json:"64"`
	Field65 uint32        `json:"65"`
	Field66 uint64        `json:"66"`
	Field67 int32         `json:"67"`
	Field68 int64         `json:"68"`
	Field69 string        `json:"69"`
	Field70 []byte        `json:"70"`
	Field71 TestEnum      `json:"71"`
	Field72 NestedMessage `json:"72"`

	Field73 []int32                  `json:"73"`
	Field74 []string                 `json:"74"`
	Field75 []bool                   `json:"75"`
	Field76 []float32                `json:"76"`
	Field77 []float64                `json:"77"`
	Field78 [][]byte                 `json:"78"`
	Field79 map[string]NestedMessage `json:"79"`
	Field80 map[int32]string         `json:"80"`

	Field81 string  `json:"81"`
	Field82 int32   `json:"82"`
	Field83 int64   `json:"83"`
	Field84 uint32  `json:"84"`
	Field85 uint64  `json:"85"`
	Field86 int32   `json:"86"`
	Field87 int64   `json:"87"`
	Field88 uint32  `json:"88"`
	Field89 uint64  `json:"89"`
	Field90 int32   `json:"90"`
	Field91 int64   `json:"91"`
	Field92 float32 `json:"92"`
	Field93 float64 `json:"93"`
	Field94 bool    `json:"94"`
	Field95 string  `json:"95"`

	Field96 []string  `json:"96"`
	Field97 []bool    `json:"97"`
	Field98 []float64 `json:"98"`

	Field99  NestedMessage `json:"99"`
	Field100 TestEnum      `json:"100"`
}
