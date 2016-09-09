package goroutinepool

import (
	"testing"
)
// CPU
//	vendor_id       : GenuineIntel
//	cpu family      : 6
//	model           : 58
//	model name      : Intel(R) Core(TM) i5-3210M CPU @ 2.50GHz
//	stepping        : 9
//	microcode       : 0x12
//	cpu MHz         : 2494.429
//	cache size      : 3072 KB

// golang 1.5.4
// Windows
//	BenchmarkGoPool-4                            100          13120751 ns/op
//	BenchmarkGoNoPool-4                          100          13730785 ns/op
//	BenchmarkGoPoolWithReflectWorker-4           100          17240986 ns/op
//	BenchmarkGoPoolWithRefReflectWorker-4        100          16460941 ns/op
//	BenchmarkGoPoolWithClosureWorker-4           100          13070747 ns/op
//	BenchmarkGoPoolWithRefClosureWorker-4        100          14340820 ns/op
// Linux VMware WorkStation
//	BenchmarkGoPool-4                            200           6521967 ns/op
//	BenchmarkGoNoPool-4                          100          11954329 ns/op
//	BenchmarkGoPoolWithReflectWorker-4           100          10892595 ns/op
//	BenchmarkGoPoolWithRefReflectWorker-4        100          12450327 ns/op
//	BenchmarkGoPoolWithClosureWorker-4           200           7015494 ns/op
//	BenchmarkGoPoolWithRefClosureWorker-4        200           7110853 ns/op

// golang 1.7.1
// Windows
//	BenchmarkGoPool-4                            300           5646989 ns/op
//	BenchmarkGoNoPool-4                          200           9830562 ns/op
//	BenchmarkGoPoolWithReflectWorker-4           200           8560489 ns/op
//	BenchmarkGoPoolWithRefReflectWorker-4        200           9175525 ns/op
//	BenchmarkGoPoolWithClosureWorker-4           200           6195354 ns/op
//	BenchmarkGoPoolWithRefClosureWorker-4        200           6360364 ns/op
// Linux VMware WorkStation
//	BenchmarkGoPool-4                            200           6013567 ns/op
//	BenchmarkGoNoPool-4                          100          13270766 ns/op
//	BenchmarkGoPoolWithReflectWorker-4           100          10300905 ns/op
//	BenchmarkGoPoolWithRefReflectWorker-4        100          10912750 ns/op
//	BenchmarkGoPoolWithClosureWorker-4           200           6973391 ns/op
//	BenchmarkGoPoolWithRefClosureWorker-4        200           7172302 ns/op

// CPU
//	vendor_id       : GenuineIntel
//	cpu family      : 6
//	model           : 58
//	model name      : Intel(R) Xeon(R) CPU E3-1240 V2 @ 3.40GHz
//	stepping        : 9
//	microcode       : 0x1b
//	cpu MHz         : 3392.294
//	cache size      : 8192 KB

// Linux VMware ESXi 5.5.0
//	BenchmarkGoPool-4                            500           3815840 ns/op
//	BenchmarkGoNoPool-4                          300           5195865 ns/op
//	BenchmarkGoPoolWithReflectWorker-4           300           4938705 ns/op
//	BenchmarkGoPoolWithRefReflectWorker-4        300           5130110 ns/op
//	BenchmarkGoPoolWithClosureWorker-4           300           4334370 ns/op
//	BenchmarkGoPoolWithRefClosureWorker-4        300           4306747 ns/op


const counts  = 10000

func BenchmarkGoPool(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoPool(counts)
	}
}


func BenchmarkGoNoPool(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoNoPool(counts)
	}
}


func BenchmarkGoPoolWithReflectWorker(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoPoolWithReflectWorker(counts)
	}
}


func BenchmarkGoPoolWithRefReflectWorker(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoPoolWithRefReflectWorker(counts)
	}
}


func BenchmarkGoPoolWithClosureWorker(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoPoolWithClosureWorker(counts)
	}
}

func BenchmarkGoPoolWithRefClosureWorker(b *testing.B){
	for i := 0; i < b.N; i++ {
		GoPoolWithRefClosureWorker(counts)
	}
}