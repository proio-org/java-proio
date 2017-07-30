package eicio

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEventPushGet(t *testing.T) {
	buffer := &bytes.Buffer{}

	writer := NewWriter(buffer)

	event0Out := NewEvent()

	MC := &MCParticleCollection{}
	MC.Particles = append(MC.Particles, &MCParticle{})
	MC.Particles = append(MC.Particles, &MCParticle{})
	event0Out.AddCollection(MC, "MCParticles")

	simTrack := &SimTrackerHitCollection{}
	simTrack.Hits = append(simTrack.Hits, &SimTrackerHit{})
	simTrack.Hits = append(simTrack.Hits, &SimTrackerHit{})
	event0Out.AddCollection(simTrack, "TrackerHits")

	writer.PushEvent(event0Out)

	event1Out := NewEvent()

	simTrack = &SimTrackerHitCollection{}
	simTrack.Hits = append(simTrack.Hits, &SimTrackerHit{})
	simTrack.Hits = append(simTrack.Hits, &SimTrackerHit{})
	event1Out.AddCollection(simTrack, "TrackerHits")

	writer.PushEvent(event1Out)

	reader := NewReader(buffer)

	event0In, err := reader.GetEvent()
	if err != nil {
		t.Error("Event 0 failed to Get")
	}
	if !reflect.DeepEqual(event0Out, event0In) {
		t.Error("Event 0 corrupted")
	}

	event1In, err := reader.GetEvent()
	if err != nil {
		t.Error("Event 1 failed to Get")
	}
	if !reflect.DeepEqual(event1Out, event1In) {
		t.Error("Event 1 corrupted")
	}
}
