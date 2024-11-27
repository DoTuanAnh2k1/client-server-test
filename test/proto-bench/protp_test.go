package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	jsonPkg "proto-bench/json"

	pb "proto-bench/proto"

	"google.golang.org/protobuf/proto"
)

var (
	pbTest = pb.BenchmarkMessage{
		// Primitive types
		Field_1:  123,
		Field_2:  456789,
		Field_3:  789,
		Field_4:  123456,
		Field_5:  -123,
		Field_6:  -456789,
		Field_7:  987,
		Field_8:  654321,
		Field_9:  42,
		Field_10: 84,
		Field_11: 3.14,
		Field_12: 2.718281828,
		Field_13: true,
		Field_14: "Hello, World!",

		// Enum field
		Field_15: pb.TestEnum_OPTION_ONE,

		// Nested message
		Field_16: &pb.NestedMessage{
			NestedId:     1,
			NestedName:   "Nested Example",
			NestedValues: []float32{1.1, 2.2, 3.3},
		},

		// Repeated fields
		Field_17: []int32{1, 2, 3, 4, 5},
		Field_18: []string{"one", "two", "three"},

		// Map field
		Field_19: map[string]int32{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},

		// Additional fields
		Field_20: "Sample Text",
		Field_21: 100,
		Field_22: 200,
		Field_23: 300,
		Field_24: 400,
		Field_25: 500,
		Field_26: 600,
		Field_27: false,
		Field_28: 7.77,
		Field_29: 9.99,
		Field_30: 700,
		Field_31: 800,
		Field_32: -100,
		Field_33: -200,
		Field_34: "Another Sample Text",
		Field_35: []byte{0x01, 0x02, 0x03},
		Field_36: []*pb.NestedMessage{
			{NestedId: 2, NestedName: "Nested1", NestedValues: []float32{4.4, 5.5}},
			{NestedId: 3, NestedName: "Nested2", NestedValues: []float32{6.6, 7.7}},
		},
		Field_37: map[int32]*pb.NestedMessage{
			1: {NestedId: 4, NestedName: "NestedMap1", NestedValues: []float32{8.8}},
			2: {NestedId: 5, NestedName: "NestedMap2", NestedValues: []float32{9.9}},
		},
		Field_38: -300,
		Field_39: -400,
		Field_40: 900,
		Field_41: 1000,
		Field_42: 1100,
		Field_43: 1200,
		Field_44: 1300,
		Field_45: 1400,
		Field_46: 1500,
		Field_47: 1600,
		Field_48: 1.23,
		Field_49: 4.56,
		Field_50: true,
		Field_51: "Field51 Example",
		Field_52: "Field52 Example",
		Field_53: "Field53 Example",
		Field_54: "Field54 Example",
		Field_55: "Field55 Example",
		Field_56: 2100,
		Field_57: 2200,
		Field_58: 2300,
		Field_59: 2400,
		Field_60: 2500,
		Field_61: 2600,
		Field_62: false,
		Field_63: 5.67,
		Field_64: 8.90,
		Field_65: 2700,
		Field_66: 2800,
		Field_67: 2900,
		Field_68: 3000,
		Field_69: "Field69 Example",
		Field_70: []byte{0x10, 0x20, 0x30},
		Field_71: pb.TestEnum_OPTION_TWO,
		Field_72: &pb.NestedMessage{NestedId: 6, NestedName: "NestedField72", NestedValues: []float32{10.1, 11.2}},

		Field_73: []int32{101, 102, 103},
		Field_74: []string{"alpha", "beta", "gamma"},
		Field_75: []bool{true, false, true},
		Field_76: []float32{12.12, 13.13},
		Field_77: []float64{14.14, 15.15},
		Field_78: [][]byte{{0x01}, {0x02}},
		Field_79: map[string]*pb.NestedMessage{
			"mapKey1": {NestedId: 7, NestedName: "MapNested1", NestedValues: []float32{16.16}},
			"mapKey2": {NestedId: 8, NestedName: "MapNested2", NestedValues: []float32{17.17}},
		},
		Field_80: map[int32]string{
			101: "MapValue1",
			102: "MapValue2",
		},

		Field_81: "Another Example",
		Field_82: 3100,
		Field_83: 3200,
		Field_84: 3300,
		Field_85: 3400,
		Field_86: 3500,
		Field_87: 3600,
		Field_88: 3700,
		Field_89: 3800,
		Field_90: 3900,
		Field_91: 4000,
		Field_92: 3.33,
		Field_93: 6.66,
		Field_94: true,
		Field_95: "Field95 Example",

		Field_96: []string{"x", "y", "z"},
		Field_97: []bool{false, true},
		Field_98: []float64{18.18, 19.19},

		Field_99:  &pb.NestedMessage{NestedId: 9, NestedName: "Nested99", NestedValues: []float32{20.2, 21.21}},
		Field_100: pb.TestEnum_OPTION_THREE,
	}

	easyJsonTest = jsonPkg.BenchmarkMessage{
		// Primitive types
		Field1:  123,
		Field2:  456789,
		Field3:  789,
		Field4:  123456,
		Field5:  -123,
		Field6:  -456789,
		Field7:  987,
		Field8:  654321,
		Field9:  42,
		Field10: 84,
		Field11: 3.14,
		Field12: 2.718281828,
		Field13: true,
		Field14: "Hello, World!",

		// Enum field
		Field15: jsonPkg.TestEnum_OPTION_ONE,

		// Nested message
		Field16: jsonPkg.NestedMessage{
			NestedID:     1,
			NestedName:   "Nested Example",
			NestedValues: []float32{1.1, 2.2, 3.3},
		},

		// Repeated fields
		Field17: []int32{1, 2, 3, 4, 5},
		Field18: []string{"one", "two", "three"},

		// Map field
		Field19: map[string]int32{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},

		// Additional fields
		Field20: "Sample Text",
		Field21: 100,
		Field22: 200,
		Field23: 300,
		Field24: 400,
		Field25: 500,
		Field26: 600,
		Field27: false,
		Field28: 7.77,
		Field29: 9.99,
		Field30: 700,
		Field31: 800,
		Field32: -100,
		Field33: -200,
		Field34: "Another Sample Text",
		Field35: []byte{0x01, 0x02, 0x03},
		Field36: []jsonPkg.NestedMessage{
			{NestedID: 2, NestedName: "Nested1", NestedValues: []float32{4.4, 5.5}},
			{NestedID: 3, NestedName: "Nested2", NestedValues: []float32{6.6, 7.7}},
		},
		Field37: map[int32]jsonPkg.NestedMessage{
			1: {NestedID: 4, NestedName: "NestedMap1", NestedValues: []float32{8.8}},
			2: {NestedID: 5, NestedName: "NestedMap2", NestedValues: []float32{9.9}},
		},

		Field38: -300,
		Field39: -400,
		Field40: 900,
		Field41: 1000,
		Field42: 1100,
		Field43: 1200,
		Field44: 1300,
		Field45: 1400,
		Field46: 1500,
		Field47: 1600,
		Field48: 1.23,
		Field49: 4.56,
		Field50: true,
		Field51: "Field51 Example",
		Field52: "Field52 Example",
		Field53: "Field53 Example",
		Field54: "Field54 Example",
		Field55: "Field55 Example",
		Field56: 2100,
		Field57: 2200,
		Field58: 2300,
		Field59: 2400,
		Field60: 2500,
		Field61: 2600,
		Field62: false,
		Field63: 5.67,
		Field64: 8.90,
		Field65: 2700,
		Field66: 2800,
		Field67: 2900,
		Field68: 3000,
		Field69: "Field69 Example",
		Field70: []byte{0x10, 0x20, 0x30},
		Field71: jsonPkg.TestEnum_OPTION_TWO,
		Field72: jsonPkg.NestedMessage{NestedID: 6, NestedName: "NestedField72", NestedValues: []float32{10.1, 11.2}},

		Field73: []int32{101, 102, 103},
		Field74: []string{"alpha", "beta", "gamma"},
		Field75: []bool{true, false, true},
		Field76: []float32{12.12, 13.13},
		Field77: []float64{14.14, 15.15},
		Field78: [][]byte{{0x01}, {0x02}},
		Field79: map[string]jsonPkg.NestedMessage{
			"mapKey1": {NestedID: 7, NestedName: "MapNested1", NestedValues: []float32{16.16}},
			"mapKey2": {NestedID: 8, NestedName: "MapNested2", NestedValues: []float32{17.17}},
		},
		Field80: map[int32]string{
			101: "MapValue1",
			102: "MapValue2",
		},

		Field81: "Another Example",
		Field82: 3100,
		Field83: 3200,
		Field84: 3300,
		Field85: 3400,
		Field86: 3500,
		Field87: 3600,
		Field88: 3700,
		Field89: 3800,
		Field90: 3900,
		Field91: 4000,
		Field92: 3.33,
		Field93: 6.66,
		Field94: true,
		Field95: "Field95 Example",

		Field96: []string{"x", "y", "z"},
		Field97: []bool{false, true},
		Field98: []float64{18.18, 19.19},

		Field99:  jsonPkg.NestedMessage{NestedID: 9, NestedName: "Nested99", NestedValues: []float32{20.2, 21.21}},
		Field100: jsonPkg.TestEnum_OPTION_THREE,

		// Timestamps
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

func TestDataAllocationsLarge(_ *testing.T) {
	fmt.Printf("Large ---------\n")
	bs := pbTest
	ej := easyJsonTest
	j, _ := json.Marshal(&bs)
	p, _ := proto.Marshal(&bs)
	e, _ := ej.MarshalJSON()

	printInfo(j, "json")
	printInfo(p, "protobuf")
	printInfo(e, "easyjson")
	fmt.Printf("\n")
}

func BenchmarkJSONMarshal(b *testing.B) {
	data := pbTest

	b.ResetTimer()

	b.Run("Json", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&data)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufMarshal(b *testing.B) {
	data := pbTest

	b.ResetTimer()

	b.Run("Proto", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&data)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkEasyjsonMarshal(b *testing.B) {
	data := easyJsonTest
	b.ResetTimer()

	b.Run("EasyJson", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.MarshalJSON()
		}
	})
	fmt.Printf("\n")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	data := pbTest
	dataD, _ := json.Marshal(&data)
	var dataf pb.BenchmarkMessage

	b.ResetTimer()

	b.Run("Json", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(dataD, &dataf)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	data := pbTest
	dataD, _ := proto.Marshal(&data)
	var dataf pb.BenchmarkMessage

	b.ResetTimer()

	b.Run("Proto", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(dataD, &dataf)
		}
	})
}

func BenchmarkEasyjsonUnmarshal(b *testing.B) {
	data, _ := easyJsonTest.MarshalJSON()

	b.ResetTimer()

	b.Run("Proto", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			easyJsonTest.UnmarshalJSON(data)
		}
	})
}

func printInfo(d []byte, ser string) {
	used := len(d)
	allocated := cap(d)
	fmt.Printf("Type: %s \t\tData size: %d \t\tTotal Allocated: %d \t\t Used/Allocated: %.2f%%\n", ser, used, allocated, percentUsed(used, allocated)*100)
}

func percentUsed(used, allocated int) float32 {
	return float32(used) / float32(allocated)
}
