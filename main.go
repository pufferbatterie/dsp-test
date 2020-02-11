package main

import (
	"fmt"
	"github.com/mjibson/go-dsp/wav"
	"github.com/snuffpuppet/spectre/spectral"
	"os"
)

func max(arr []float64) (maxVal float64, maxIndex int) {
	var max float64 = 0
	var imax = 0
	for i, v := range arr {
		if v > max {
			max = v
			imax = i
		}
	}
	return max, imax
}

func main() {
	filePath := "1kHz_44100Hz_16bit_05sec.wav"
	//	filePath := "100Hz_44100Hz_16bit_05sec.wav"

	f, _ := os.Open(filePath)
	wav, _ := wav.New(f)
	floats, _ := wav.ReadFloats(wav.Samples)

	b := make([]float64, wav.Samples)
	for i, v := range floats {
		b[i] = float64(v)
	}

	pxx, freqs := spectral.Simple(b, int(wav.SampleRate))
	_, imax := max(pxx)

	fmt.Printf("max-2: val=%.2f \t f=%.2f\n", pxx[imax-2], freqs[imax-2]) //max-2: val=1787.00       f=999.42
	fmt.Printf("max-1: val=%.2f \t f=%.2f\n", pxx[imax-1], freqs[imax-1]) //max-1: val=3752.55       f=999.62
	fmt.Printf("max  : val=%.2f \t f=%.2f\n", pxx[imax], freqs[imax])     //max  : val=37622.17      f=999.82
	fmt.Printf("max+1: val=%.2f \t f=%.2f\n", pxx[imax+1], freqs[imax+1]) //max+1: val=3128.90       f=1000.02
	fmt.Printf("max+2: val=%.2f \t f=%.2f\n", pxx[imax+2], freqs[imax+2]) //max+2: val=1632.45       f=1000.22

}
