package multierror

import (
	"testing"
)

func waitNotNil() error {
	g := new(Group)
	g.Go(func() error { return nil })
	return g.Wait()
}

func waitNil() error {
	g := new(Group)
	g.Go(func() error { return nil })
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

// This test will fail with the current implementation of Group.Wait
func TestWait_NilNotNil(t *testing.T) {
	err1 := waitNotNil()
	if err1 != nil {
		t.Fatalf("err1 should be nil; actual %v", err1)
	}

	err2 := waitNil()
	if err2 != nil {
		t.Fatalf("err2 should be nil; actual %v", err2)
	}
}
