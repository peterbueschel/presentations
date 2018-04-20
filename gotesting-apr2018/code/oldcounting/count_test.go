/*
testTable := []struct {
	name    string
	input   string
	want    string
	wantErr bool
}{
	 {"neutral1",    "yes",                               `Hlja'`,  false},
	 {"neutral2",    "no",                                `ghobe'`, false},
	{"friendly1",   "Hello",                             `nuqneH`, false},
	{"friendly2",   "Well done!",                        `majQa'`, false},
	{"unfriendly1", "Your mother has a smooth forehead", ``,       true},
}

*/

package translator

import (
	"fmt"
	"testing"
)

func TestToOldKlingonNumber_Errors(t *testing.T) {
	tests := []struct {
		name    string
		input   uint8
		want    string
		wantErr bool
	}{
		{"less than 1", 0, `NaKN`, true},     // isolated: maybe parallel testing // HL
		{"greater than 9", 10, `NaKN`, true}, // HL
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Subtests // HL
			_, err := ToOldKlingonNumber(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToOldKlingonNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestToOldKlingonNumber(t *testing.T) {
	tests := []struct {
		input   uint8
		want    string
		wantErr bool
	}{ //                               ┌──────── cha' ? input > 6 : ''
		{0, `NaKN`, true},           // │ ┌────── 'wejje' ? input > 3 : ''
		{1, `wa'`, false},           // 0x3+1──── ring(input % 3) // HL
		{2, `cha'`, false},          // 0x3+2 // HL
		{3, `wej`, false},           // 0x3+3 // HL
		{4, `wejjewa'`, false},      // 1x3+1 // HL
		{5, `wejjecha'`, false},     // 1x3+2 // HL
		{6, `wejjewej`, false},      // 1x3+3 // HL
		{7, `cha'wejjewa'`, false},  // 2x3+1 // HL
		{8, `cha'wejjecha'`, false}, // 2x3+2 // HL
		{9, `cha'wejjewej`, false},  // 2x3+3 // HL
		{10, `NaKN`, true},          //
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			got, err := ToOldKlingonNumber(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToOldKlingonNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToOldKlingonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Benchmark_ToOldKlingonNumber(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		ToOldKlingonNumber(9)
//	}
//}
