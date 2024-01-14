package util_test

import (
	"testing"
	"time"

	"application-gaspar/internal/util"
)

func TestTimeOverlap(t *testing.T) {
	f1, _ := time.Parse(time.DateOnly, "2023-12-16")
	t1, _ := time.Parse(time.DateOnly, "2023-12-21")
	f2, _ := time.Parse(time.DateOnly, "2023-12-22")
	t2, _ := time.Parse(time.DateOnly, "2023-12-25")
	if ret := util.TimeOverlap(f1, t1, f2, t2); ret {
		t.Fatalf("expected false, but got %v", ret)
	}

	f1, _ = time.Parse(time.DateOnly, "2023-12-16")
	t1, _ = time.Parse(time.DateOnly, "2023-12-21")
	f2, _ = time.Parse(time.DateOnly, "2023-12-11")
	t2, _ = time.Parse(time.DateOnly, "2023-12-15")
	if ret := util.TimeOverlap(f1, t1, f2, t2); ret {
		t.Fatalf("expected false, but got %v", ret)
	}

	f1, _ = time.Parse(time.DateOnly, "2023-12-16")
	t1, _ = time.Parse(time.DateOnly, "2023-12-21")
	f2, _ = time.Parse(time.DateOnly, "2023-12-17")
	t2, _ = time.Parse(time.DateOnly, "2023-12-22")
	if ret := util.TimeOverlap(f1, t1, f2, t2); !ret {
		t.Fatalf("expected true, but got %v", ret)
	}

	f1, _ = time.Parse(time.DateOnly, "2023-12-16")
	t1, _ = time.Parse(time.DateOnly, "2023-12-21")
	f2, _ = time.Parse(time.DateOnly, "2023-12-15")
	t2, _ = time.Parse(time.DateOnly, "2023-12-22")
	if ret := util.TimeOverlap(f1, t1, f2, t2); !ret {
		t.Fatalf("expected true, but got %v", ret)
	}

	f1, _ = time.Parse(time.DateOnly, "2023-12-16")
	t1, _ = time.Parse(time.DateOnly, "2023-12-21")
	f2, _ = time.Parse(time.DateOnly, "2023-12-15")
	t2, _ = time.Parse(time.DateOnly, "2023-12-17")
	if ret := util.TimeOverlap(f1, t1, f2, t2); !ret {
		t.Fatalf("expected true, but got %v", ret)
	}

	f1, _ = time.Parse(time.DateOnly, "2023-12-16")
	t1, _ = time.Parse(time.DateOnly, "2023-12-21")
	f2, _ = time.Parse(time.DateOnly, "2023-12-16")
	t2, _ = time.Parse(time.DateOnly, "2023-12-21")
	if ret := util.TimeOverlap(f1, t1, f2, t2); !ret {
		t.Fatalf("expected true, but got %v", ret)
	}
}

func TestValidateEmail(t *testing.T) {
	if ret := util.ValidateEmail("some@email.com"); !ret {
		t.Fatalf("expected true, but got %v", ret)
	}

	if ret := util.ValidateEmail("some_text"); ret {
		t.Fatalf("expected false, but got %v", ret)
	}

	if ret := util.ValidateEmail(""); ret {
		t.Fatalf("expected false, but got %v", ret)
	}
}
