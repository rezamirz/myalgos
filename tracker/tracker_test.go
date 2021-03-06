package tracker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TRACKER_SIZE = 1000

func TestSimpleDynamicTracker(t *testing.T) {
	trackerSize := uint64(TRACKER_SIZE)
	tracker := NewTracker(trackerSize, DynamicTracker, 0)

	for i := uint64(0); i < tracker.size; i++ {
		err := tracker.Track(i)
		assert.Equal(t, nil, err)

		nextLowcontig, _ := tracker.NextLowcontig()
		assert.Equal(t, i+1, nextLowcontig)

		rc := tracker.IsTracked(i)
		assert.Equal(t, true, rc)
	}
}

func TestSimpleFixedTracker(t *testing.T) {
	trackerSize := uint64(TRACKER_SIZE)
	tracker := NewTracker(trackerSize, FixedTracker, 0)

	for i := uint64(0); i < tracker.size; i++ {
		err := tracker.Track(i)
		assert.Equal(t, nil, err)

		nextLowcontig, err := tracker.NextLowcontig()
		if i < tracker.size-1 {
			assert.Equal(t, i+1, nextLowcontig)
		}

		rc := tracker.IsTracked(i)
		assert.Equal(t, true, rc)
	}

	err := tracker.Untrack(5)
	assert.Equal(t, nil, err)

	nextLowcontig, err := tracker.NextLowcontig()
	assert.Equal(t, uint64(5), nextLowcontig)

	rc := tracker.IsTracked(5)
	assert.Equal(t, false, rc)

}


func TestTracker(t *testing.T) {

	fmt.Printf("This test takes more than 10 sec ...\n")

	for j := 1; j <= 10; j++ {
		trackerSize := uint64(j * TRACKER_SIZE)
		tracker := NewTracker(trackerSize, DynamicTracker, 0)

		for i := uint64(0); i < tracker.size; i++ {
			err := tracker.Track(i)
			assert.Equal(t, nil, err)

			n, err := tracker.Next(i)
			assert.Equal(t, i+1, n)
			assert.Equal(t, nil, err)

			nextLowcontig, _ := tracker.NextLowcontig()
			assert.Equal(t, i+1, nextLowcontig)

			rc := tracker.IsTracked(i)
			assert.Equal(t, true, rc)
		}
	}

	for j := 1; j <= 100; j++ {
		trackerSize := uint64(j * TRACKER_SIZE)
		nl := uint64(j * 19)

		tracker := NewTracker(trackerSize, DynamicTracker, nl)

		for i := nl; i < tracker.size; i++ {
			_ = tracker.Track(i)

			n, _ := tracker.Next(i)
			assert.Equal(t, i+1, n)

			nextLowcontig, _ := tracker.NextLowcontig()
			assert.Equal(t, i+1, nextLowcontig)

			_ = tracker.IsTracked(i)
		}
	}

	for j := 1; j <= 100; j++ {
		trackerSize := uint64(j * TRACKER_SIZE)
		nl := uint64(j*19 + j)

		tracker := NewTracker(trackerSize, DynamicTracker, nl)

		for i := nl; i < trackerSize; i++ {
			_ = tracker.Track(i)

			n, _ := tracker.Next(i)
			assert.Equal(t, i+1, n)

			nextLowcontig, _ := tracker.NextLowcontig()
			assert.Equal(t, i+1, nextLowcontig)

			_ = tracker.IsTracked(i)
		}
	}

}

func TestTrackerLowcontig(t *testing.T) {
	trackerSize := uint64(TRACKER_SIZE)
	tracker := NewTracker(trackerSize, DynamicTracker, 0)

	err := tracker.Track(10)
	assert.Equal(t, nil, err)
	nextLowcontig, _ := tracker.NextLowcontig()
	assert.Equal(t, uint64(0), nextLowcontig)

	err = tracker.Track(11)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(0), nextLowcontig)

	err = tracker.Track(2)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(0), nextLowcontig)

	err = tracker.Track(1)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(0), nextLowcontig)

	err = tracker.Track(0)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(3), nextLowcontig)

	err = tracker.Track(3)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(4), nextLowcontig)

	err = tracker.Track(6)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(4), nextLowcontig)

	err = tracker.Track(4)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(5), nextLowcontig)

	err = tracker.Track(5)
	assert.Equal(t, nil, err)
	nextLowcontig, _ = tracker.NextLowcontig()
	assert.Equal(t, uint64(7), nextLowcontig)
}
