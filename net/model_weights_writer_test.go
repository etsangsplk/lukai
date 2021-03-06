package net

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"errors"
	"io"
	"reflect"
	"sync"
	"testing"

	"github.com/luk-ai/lukai/protobuf/aggregatorpb"
	"github.com/luk-ai/lukai/testutil"
	"github.com/luk-ai/lukai/units"
)

func TestModelWeightsWriter(t *testing.T) {

	chunkChan := make(chan *aggregatorpb.ModelWeightChunk)

	var expect bytes.Buffer
	w := NewModelWeightsWriter(func(chunk aggregatorpb.ModelWeightChunk) error {
		chunkChan <- &chunk
		return nil
	})
	r := ReadModelWeights(func() (*aggregatorpb.ModelWeightChunk, error) {
		return <-chunkChan, nil
	})

	var buf bytes.Buffer
	mu := struct {
		sync.Mutex

		done bool
	}{}
	go func() {
		if _, err := buf.ReadFrom(r); err != nil {
			t.Fatal(err)
		}
		mu.Lock()
		defer mu.Unlock()
		mu.done = true
	}()

	mw := io.MultiWriter(w, &expect)

	gw := gzip.NewWriter(mw)

	if _, err := gw.Write([]byte(`first part of the message`)); err != nil {
		t.Fatal(err)
	}
	msg := make([]byte, 1*units.MB)
	if _, err := rand.Read(msg); err != nil {
		t.Fatal(err)
	}
	if _, err := gw.Write(msg); err != nil {
		t.Fatal(err)
	}
	if _, err := gw.Write([]byte(`middle part of the message`)); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		msg = make([]byte, 1537*units.B)
		if _, err := rand.Read(msg); err != nil {
			t.Fatal(err)
		}
		if _, err := gw.Write(msg); err != nil {
			t.Fatal(err)
		}
	}

	msg = make([]byte, 1*units.MB)
	if _, err := rand.Read(msg); err != nil {
		t.Fatal(err)
	}

	if _, err := gw.Write(msg); err != nil {
		t.Fatal(err)
	}
	if _, err := gw.Write([]byte(`last part of the message`)); err != nil {
		t.Fatal(err)
	}

	if err := gw.Close(); err != nil {
		t.Fatal(err)
	}

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	testutil.SucceedsSoon(t, func() error {
		mu.Lock()
		defer mu.Unlock()

		if !mu.done {
			return errors.New("ReadFrom hasn't finished")
		}
		return nil
	})

	want := expect.Bytes()
	out := buf.Bytes()
	if !reflect.DeepEqual(want, out) {
		t.Fatalf("got '%s' (len %d); want '%s' (len %d)", out, len(out), want, len(want))
	}
}
