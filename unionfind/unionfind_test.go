package unionfind

import "testing"

func TestJoin(t *testing.T) {
	finder := Newfinder(6)
	err := finder.Join(0, 4)
	if err != nil {
		t.Error(err)
	}
	a, _ := finder.GetValue(0)
	b, _ := finder.GetValue(4)

	if a != b {
		t.Errorf("Values do not match: %v, %v", a, b)
	}
	err = finder.Join(2, 3, 5)
	if err != nil {
		t.Error(err)
	}
	a, _ = finder.GetValue(2)
	b, _ = finder.GetValue(3)
	c, _ := finder.GetValue(5)
	if a != b || a != c || b != c {
		t.Errorf("Values do not match: %v, %v, %v", a, b, c)
	}
}

func TestJoinEager(t *testing.T) {
	finder := Newfinder(10)
	finder.Join(0, 2, 3)
	finder.JoinEager(1, 2)
	for i := 0; i < 4; i++ {
		res, _ := finder.GetValue(i)
		if res != 1 {
			t.Errorf("Error: incorrect value. Wanted %v, got %v", 1, res)
		}
	}
}

func TestConnected(t *testing.T) {
	finder := Newfinder(2)
	finder.Join(0, 1)
	if !finder.Connected(0, 1) {
		t.Error("Error: numbers should be connected")
	}
	a, _ := finder.GetValue(0)
	b, _ := finder.GetValue(1)
	if a != b {
		t.Error("Well clearly the connected function has some problems")
	}
}
