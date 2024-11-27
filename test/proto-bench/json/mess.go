package json

import "time"

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
	NestedID     int32     `json:"nested_id"`
	NestedName   string    `json:"nested_name"`
	NestedValues []float32 `json:"nested_values"`
}

// Main struct with 100+ fields
// easyjson:json
type BenchmarkMessage struct {
	// Primitive types
	Field1  int32   `json:"field_1"`
	Field2  int64   `json:"field_2"`
	Field3  uint32  `json:"field_3"`
	Field4  uint64  `json:"field_4"`
	Field5  int32   `json:"field_5"`
	Field6  int64   `json:"field_6"`
	Field7  uint32  `json:"field_7"`
	Field8  uint64  `json:"field_8"`
	Field9  int32   `json:"field_9"`
	Field10 int64   `json:"field_10"`
	Field11 float32 `json:"field_11"`
	Field12 float64 `json:"field_12"`
	Field13 bool    `json:"field_13"`
	Field14 string  `json:"field_14"`

	// Enum field
	Field15 TestEnum `json:"field_15"`

	// Nested message
	Field16 NestedMessage `json:"field_16"`

	// Repeated fields
	Field17 []int32  `json:"field_17"`
	Field18 []string `json:"field_18"`

	// Map field
	Field19 map[string]int32 `json:"field_19"`

	// Additional fields
	Field20 string                  `json:"field_20"`
	Field21 int32                   `json:"field_21"`
	Field22 int64                   `json:"field_22"`
	Field23 uint32                  `json:"field_23"`
	Field24 uint64                  `json:"field_24"`
	Field25 int32                   `json:"field_25"`
	Field26 int64                   `json:"field_26"`
	Field27 bool                    `json:"field_27"`
	Field28 float32                 `json:"field_28"`
	Field29 float64                 `json:"field_29"`
	Field30 uint32                  `json:"field_30"`
	Field31 uint64                  `json:"field_31"`
	Field32 int32                   `json:"field_32"`
	Field33 int64                   `json:"field_33"`
	Field34 string                  `json:"field_34"`
	Field35 []byte                  `json:"field_35"`
	Field36 []NestedMessage         `json:"field_36"`
	Field37 map[int32]NestedMessage `json:"field_37"`

	// Remaining fields up to 100+
	Field38 int32         `json:"field_38"`
	Field39 int64         `json:"field_39"`
	Field40 uint32        `json:"field_40"`
	Field41 uint64        `json:"field_41"`
	Field42 int32         `json:"field_42"`
	Field43 int64         `json:"field_43"`
	Field44 uint32        `json:"field_44"`
	Field45 uint64        `json:"field_45"`
	Field46 int32         `json:"field_46"`
	Field47 int64         `json:"field_47"`
	Field48 float32       `json:"field_48"`
	Field49 float64       `json:"field_49"`
	Field50 bool          `json:"field_50"`
	Field51 string        `json:"field_51"`
	Field52 string        `json:"field_52"`
	Field53 string        `json:"field_53"`
	Field54 string        `json:"field_54"`
	Field55 string        `json:"field_55"`
	Field56 int32         `json:"field_56"`
	Field57 int64         `json:"field_57"`
	Field58 uint32        `json:"field_58"`
	Field59 uint64        `json:"field_59"`
	Field60 int32         `json:"field_60"`
	Field61 int64         `json:"field_61"`
	Field62 bool          `json:"field_62"`
	Field63 float32       `json:"field_63"`
	Field64 float64       `json:"field_64"`
	Field65 uint32        `json:"field_65"`
	Field66 uint64        `json:"field_66"`
	Field67 int32         `json:"field_67"`
	Field68 int64         `json:"field_68"`
	Field69 string        `json:"field_69"`
	Field70 []byte        `json:"field_70"`
	Field71 TestEnum      `json:"field_71"`
	Field72 NestedMessage `json:"field_72"`

	Field73 []int32                  `json:"field_73"`
	Field74 []string                 `json:"field_74"`
	Field75 []bool                   `json:"field_75"`
	Field76 []float32                `json:"field_76"`
	Field77 []float64                `json:"field_77"`
	Field78 [][]byte                 `json:"field_78"`
	Field79 map[string]NestedMessage `json:"field_79"`
	Field80 map[int32]string         `json:"field_80"`

	Field81 string  `json:"field_81"`
	Field82 int32   `json:"field_82"`
	Field83 int64   `json:"field_83"`
	Field84 uint32  `json:"field_84"`
	Field85 uint64  `json:"field_85"`
	Field86 int32   `json:"field_86"`
	Field87 int64   `json:"field_87"`
	Field88 uint32  `json:"field_88"`
	Field89 uint64  `json:"field_89"`
	Field90 int32   `json:"field_90"`
	Field91 int64   `json:"field_91"`
	Field92 float32 `json:"field_92"`
	Field93 float64 `json:"field_93"`
	Field94 bool    `json:"field_94"`
	Field95 string  `json:"field_95"`

	Field96 []string  `json:"field_96"`
	Field97 []bool    `json:"field_97"`
	Field98 []float64 `json:"field_98"`

	Field99  NestedMessage `json:"field_99"`
	Field100 TestEnum      `json:"field_100"`

	// Optional timestamps for additional benchmark scenarios
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
