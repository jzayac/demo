package airline

import (
	"testing"
)

func TestAirlinesLen(t *testing.T) {
	c := airline{}

	size := c.Len()
	expect := 0

	if size != expect {
		t.Fatalf("Expect Len to be %d but get %d", expect, size)
	}
}

func TestAirlinesGetList(t *testing.T) {
	cases := []struct {
		inst      *airline
		expectLen int
	}{
		{
			inst: &airline{
				"1": "1",
			},
			expectLen: 1,
		},
		{
			inst: &airline{
				"1": "1",
				"2": "2",
			},
			expectLen: 2,
		},
	}

	for _, c := range cases {
		list := c.inst.GetList()
		size := len(list)
		if size != c.expectLen {
			t.Fatalf("Expect Len to be %d but get %d", c.expectLen, size)
		}
	}
}

// TODO: many onstances with gorutine

// func TestAirlinesLen(t *testing.T) {
// 	cities := Airlines
// }

// func TestIsFile(t *testing.T) {
// 	cities := Airlines
// 	cases := []struct {
// 		index  string
// 		value  string
// 		expect string
// 	}{
// 		{
// 			index:  "123",
// 			path:   "Onanovo",
// 			expect: "Onanovo",
// 		},
// 	}

// 	for _, c := range cases {
// 		exc := IsFile(c.index)
// 		if exc != c.expect {
// 			t.Fatalf("Expect to not be %t with file %s", exc, c.path)
// 		}
// 	}
// }

// func TestGetDirName(t *testing.T) {
// 	cases := []struct {
// 		path   string
// 		expect string
// 	}{
// 		{
// 			path:   "/movies/test.rit",
// 			expect: "movies",
// 		},
// 		{
// 			path:   "/movies/test/sub/",
// 			expect: "sub",
// 		},
// 		{
// 			path:   "/movies/test/sub/more",
// 			expect: "sub",
// 		},
// 		{
// 			path:   "/",
// 			expect: "/",
// 		},
// 	}

// 	for _, c := range cases {
// 		name := GetDirName(c.path)
// 		if name != c.expect {
// 			t.Fatalf("Expected name is %s but get %s", c.expect, name)
// 		}
// 	}
// }

// func TestGetParentDirName(t *testing.T) {
// 	cases := []struct {
// 		path   string
// 		expect string
// 	}{
// 		{
// 			path:   "/movies/test.rit",
// 			expect: "/",
// 		},
// 		{
// 			path:   "/movies/test/sub/",
// 			expect: "test",
// 		},
// 		{
// 			path:   "/movies/test/sub/more",
// 			expect: "test",
// 		},
// 	}

// 	for _, c := range cases {
// 		name := GetParentDirName(c.path)
// 		if name != c.expect {
// 			t.Fatalf("Expected name is %s but get %s", c.expect, name)
// 		}
// 	}
// }

// func TestGetParentDirPath(t *testing.T) {
// 	cases := []struct {
// 		path   string
// 		expect string
// 	}{
// 		{
// 			path:   "/movies/test/asdasd.123",
// 			expect: "/movies/",
// 		},
// 		{
// 			path:   "/movies/test/sub",
// 			expect: "/movies/",
// 		},
// 		{
// 			path:   "/series/test/",
// 			expect: "/series/",
// 		},
// 		{
// 			path:   "/series/test",
// 			expect: "/",
// 		},
// 		{
// 			path:   "/movies/",
// 			expect: "/",
// 		},
// 		{
// 			path:   "/",
// 			expect: "/",
// 		},
// 		{
// 			path:   "/movies/test/sub/",
// 			expect: "/movies/test/",
// 		},
// 	}

// 	for _, c := range cases {
// 		name := GetParentDirPath(c.path)
// 		if name != c.expect {
// 			t.Fatalf("Expected name is %s but get %s", c.expect, name)
// 		}
// 	}
// }

// func TestGetDirPath(t *testing.T) {
// 	cases := []struct {
// 		path   string
// 		expect string
// 	}{
// 		{
// 			path:   "/movies/test/asdasd.123",
// 			expect: "/movies/test/",
// 		},
// 		{
// 			path:   "/movies/test/sub",
// 			expect: "/movies/test/",
// 		},
// 		{
// 			path:   "/movies/test/",
// 			expect: "/movies/test/",
// 		},
// 		{
// 			path:   "/",
// 			expect: "/",
// 		},
// 	}

// 	for _, c := range cases {
// 		exc := GetDirPath(c.path)
// 		if exc != c.expect {
// 			t.Fatalf("Expect path to be %s but get %s", exc, c.expect)
// 		}
// 	}
// }
