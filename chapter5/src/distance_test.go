package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance_CalculateDistance(t *testing.T) {
	oxford := CreateCity(51.45, 1.15)
	turin := CreateCity(45.04, 7.42)
	ml, km := CalculateHarvesineDistance(oxford, turin)
	if ml < 0.0 || km < 0.0 {
		t.Errorf("The value cant be less than zero")
	}
}

func TestDistance_CalculateDistance_ManualBetweenPoints(t *testing.T) {
	oxford := CreateCity(51.45, 1.15)
	turin := CreateCity(45.04, 7.42)
	calculateOxford, _ := oxford.CalculateHarvesineDistanceManual(turin)
	calculateTuron, _ := turin.CalculateHarvesineDistanceManual((oxford))
	assert.Equal(t, calculateOxford, calculateTuron, "The comparison is the same")
}

func TestDistance_FullTest_Athens_Amsterdam(t *testing.T) {
	type args struct {
		athens    City
		amsterdam City
	}
	tests := map[string]struct {
		args        args
		want        float64
		expectedErr string
	}{
		"success": {args: args{athens: City{37.983972, 23.727806}, amsterdam: City{52.366667, 4.9}}, want: 2163.2310285824487},
		"failure": {args: args{athens: City{0.0, 0.0}, amsterdam: City{0.0, 0.0}}, expectedErr: "The value cant be zero"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, error := tt.args.athens.CalculateHarvesineDistanceManual(&tt.args.amsterdam)
			if tt.expectedErr != "" {
				assert.EqualError(t, error, tt.expectedErr)
				assert.Equal(t, 0.0, got)
			} else {
				assert.NoError(t, error)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDistance_FullTest_Amsterdam_Berlin(t *testing.T) {
	type args struct {
		amsterdam City
		berlin    City
	}
	tests := map[string]struct {
		args        args
		want        float64
		expectedErr string
	}{
		"success": {args: args{amsterdam: City{52.366667, 4.9}, berlin: City{52.516667, 13.388889}}, want: 575.2949643958796},
		"failure": {args: args{amsterdam: City{0.0, 0.0}, berlin: City{0.0, 0.0}}, expectedErr: "The value cant be zero"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, error := tt.args.amsterdam.CalculateHarvesineDistanceManual(&tt.args.berlin)
			if tt.expectedErr != "" {
				assert.EqualError(t, error, tt.expectedErr)
				assert.Equal(t, 0.0, got)
			} else {
				assert.NoError(t, error)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDistance_FullTest_Berlin_Athens(t *testing.T) {
	type args struct {
		berlin City
		athens City
	}
	tests := map[string]struct {
		args        args
		want        float64
		expectedErr string
	}{
		"success": {args: args{berlin: City{52.516667, 13.388889}, athens: City{37.983972, 23.727806}}, want: 1803.1087879059255},
		"failure": {args: args{berlin: City{0.0, 0.0}, athens: City{0.0, 0.0}}, expectedErr: "The value cant be zero"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, error := tt.args.berlin.CalculateHarvesineDistanceManual(&tt.args.athens)
			if tt.expectedErr != "" {
				assert.EqualError(t, error, tt.expectedErr)
				assert.Equal(t, 0.0, got)
			} else {
				assert.NoError(t, error)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

var resu float64
var err error

func BenchMark_CalculateManualDistance(b *testing.B) {

	var res1 float64
	var err1 error
	amsterdam := CreateCity(52.366667, 4.9)
	berlin := CreateCity(52.516667, 13.388889)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res1, err1 = amsterdam.CalculateHarvesineDistanceManual(berlin)
	}
	resu = res1
	err = err1
}
