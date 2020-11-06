package sampler

import (
	"fmt"
	"sync/atomic"
)

type SamplingRateSampler struct {
	samplingRate int
	counter      int32
}

func NewSamplingRateSampler(samplingRate int) *SamplingRateSampler {
	if samplingRate <= 0 {
		panic(fmt.Sprint("Invalid samplingRate ", samplingRate))
	}
	return &SamplingRateSampler{samplingRate: samplingRate}
}

func (sr *SamplingRateSampler) isSampling() bool {
	samplingCount := atomic.AddInt32(&sr.counter, 1)
	isSampling := samplingCount % int32(sr.samplingRate)
	return isSampling == 0
}
