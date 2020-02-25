/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: rand
 **Element: rand.Rand
 **Type: struct
 ------------------------------------------------------------------------------------
 **Definition:
 // A Rand is a source of random numbers.
 type Rand struct {
    // contains filtered or unexported fields
	src Source
	s64 Source64 // non-nil if src is source64

	// readVal contains remainder of 63-bit integer used for bytes
	// generation during most recent Read call.
	// It is saved so next Read call can start where the previous
	// one finished.
	readVal int64
	// readPos indicates the number of low-order bytes of readVal
	// that are still valid.
	readPos int8
 }
 func (r *Rand) ExpFloat64() float64
 func (r *Rand) Float32() float32
 func (r *Rand) Float64() float64
 func (r *Rand) Int() int
 func (r *Rand) Int31() int32
 func (r *Rand) Int31n(n int32) int32
 func (r *Rand) Int63() int64
 func (r *Rand) Int63n(n int64) int64
 func (r *Rand) Intn(n int) int
 func (r *Rand) NormFloat64() float64
 func (r *Rand) Perm(n int) []int
 func (r *Rand) Read(p []byte) (n int, err error)
 func (r *Rand) Seed(seed int64)
 func (r *Rand) Shuffle(n int, swap func(i, j int))
 func (r *Rand) Uint32() uint32
 func (r *Rand) Uint64() uint64
 Uint64 returns a pseudo-random 64-bit value as a uint64.
 ------------------------------------------------------------------------------------
 **Description:
 A Rand is a source of random numbers.
 ExpFloat64 returns an exponentially distributed float64 in the range
 (0, +math.MaxFloat64] with an exponential distribution whose rate parameter (lambda)
 is 1 and whose mean is 1/lambda (1). To produce a distribution with a different rate
 parameter, callers can adjust the output using:
 	sample = ExpFloat64() / desiredRateParameter
 Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
 Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
 Int returns a non-negative pseudo-random int.
 Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
 Int31n returns, as an int32, a non-negative pseudo-random number in [0,n). It panics
 if n <= 0.
 Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
 Int63n returns, as an int64, a non-negative pseudo-random number in [0,n). It panics
 if n <= 0.
 Intn returns, as an int, a non-negative pseudo-random number in [0,n). It panics if
 n <= 0.
 NormFloat64 returns a normally distributed float64 in the range -math.MaxFloat64
 through +math.MaxFloat64 inclusive, with standard normal distribution (mean = 0,
 stddev = 1). To produce a different normal distribution, callers can adjust the
 output using:
 	sample = NormFloat64() * desiredStdDev + desiredMean
 Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
 Read generates len(p) random bytes and writes them into p. It always returns len(p)
 and a nil error. Read should not be called concurrently with any other Rand method.
 Seed uses the provided seed value to initialize the generator to a deterministic state.
 Seed should not be called concurrently with any other Rand method.
 Shuffle pseudo-randomizes the order of elements. n is the number of elements. Shuffle
 panics if n < 0. swap swaps the elements with indexes i and j.
 Uint32 returns a pseudo-random 32-bit value as a uint32.
 Uint64 returns a pseudo-random 64-bit value as a uint64.
 ------------------------------------------------------------------------------------
 **要点总结:
 0. Rand结构体是随机数的源；
 1.
*************************************************************************************/
package main

func main() {
}
