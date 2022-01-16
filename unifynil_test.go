package unifynil

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnify(t *testing.T) {
	type A struct {
		Ll  []A
		Mm  map[int64]A
		Llp []*A
		Mmp map[int64]*A
	}

	type B struct {
		Ii []int
		Ss map[string]string
		A  A
		Ap *A
	}

	t.Run("SliceToNil + MapToNil, already nil", func(t *testing.T) {
		obj := &B{}
		want := &B{}
		Unify(obj, SliceToNil(), MapToNil())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToNil + MapToNil, not nil", func(t *testing.T) {
		obj := &B{
			Ii: []int{},
			Ss: map[string]string{},
			A: A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		want := &B{
			Ap: &A{},
		}
		Unify(obj, SliceToNil(), MapToNil())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToNil + MapToNil, not nil, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll: []A{},
					},
				},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Mm: map[int64]A{
					10: A{},
				},
			},
			Ap: &A{},
		}
		Unify(obj, SliceToNil(), MapToNil())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToNil, not nil, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll: []A{},
					},
				},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Mm: map[int64]A{
					10: A{},
				},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Mm:  map[int64]A{},
				Mmp: map[int64]*A{},
			},
		}
		Unify(obj, SliceToNil())
		require.Equal(t, want, obj)
	})

	t.Run("MapToNil, not nil, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll: []A{},
					},
				},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll: []A{},
					},
				},
				Llp: []*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Llp: []*A{},
			},
		}
		Unify(obj, MapToNil())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToEmpty + MapToEmpty, already empty", func(t *testing.T) {
		obj := &B{
			Ii: []int{},
			Ss: map[string]string{},
			A: A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		want := &B{
			Ii: []int{},
			Ss: map[string]string{},
			A: A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		Unify(obj, SliceToEmpty(), MapToEmpty())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToEmpty + MapToEmpty, not empty", func(t *testing.T) {
		obj := &B{}
		want := &B{
			Ii: []int{},
			Ss: map[string]string{},
			A: A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		Unify(obj, SliceToEmpty(), MapToEmpty())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToEmpty + MapToEmpty, not empty, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Mm: map[int64]A{
					10: A{},
				},
			},
			Ap: &A{},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll:  []A{},
						Mm:  map[int64]A{},
						Llp: []*A{},
						Mmp: map[int64]*A{},
					},
				},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Mm:  map[int64]A{},
				Llp: []*A{},
				Mmp: map[int64]*A{},
			},
		}
		Unify(obj, SliceToEmpty(), MapToEmpty())
		require.Equal(t, want, obj)
	})

	t.Run("SliceToEmpty, not empty, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Mm: map[int64]A{
					10: A{},
				},
			},
			Ap: &A{},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Ll: []A{},
				Mm: map[int64]A{
					10: A{
						Ll:  []A{},
						Llp: []*A{},
					},
				},
				Llp: []*A{},
			},
			Ap: &A{
				Ll:  []A{},
				Llp: []*A{},
			},
		}
		Unify(obj, SliceToEmpty())
		require.Equal(t, want, obj)
	})

	t.Run("MapToEmpty, not empty, some filled", func(t *testing.T) {
		obj := &B{
			Ii: []int{1, 2, 3},
			A: A{
				Mm: map[int64]A{
					10: A{},
				},
			},
			Ap: &A{},
		}
		want := &B{
			Ii: []int{1, 2, 3},
			Ss: map[string]string{},
			A: A{
				Mm: map[int64]A{
					10: A{
						Mm:  map[int64]A{},
						Mmp: map[int64]*A{},
					},
				},
				Mmp: map[int64]*A{},
			},
			Ap: &A{
				Mm:  map[int64]A{},
				Mmp: map[int64]*A{},
			},
		}
		Unify(obj, MapToEmpty())
		require.Equal(t, want, obj)
	})

}

func ExampleUnify() {
	type Response struct {
		Items []int             `json:"items"`
		Users map[string]string `json:"users"`
	}

	res := &Response{} // Leave the slice and the map nil.

	buf, _ := json.Marshal(res)
	fmt.Println(string(buf))

	Unify(res, SliceToEmpty(), MapToEmpty())

	// Now res.Items = []int{} and res.Users = map[string]string{}.

	buf, _ = json.Marshal(res)
	fmt.Println(string(buf))

	// Convert back to nils.
	Unify(res, SliceToNil(), MapToNil())

	// Now res.Items = nil and res.Users = nil.

	buf, _ = json.Marshal(res)
	fmt.Println(string(buf))

	// Output:
	// {"items":null,"users":null}
	// {"items":[],"users":{}}
	// {"items":null,"users":null}
}
