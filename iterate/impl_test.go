/*
 * Copyright (c) 2021-present Sigma-Soft, Ltd.
 * @author: Nikolay Nikitin
 */

package iterate

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	ITested1 interface {
		Fields(enum func(name string))
	}
	tested1 struct {
		fields []string
	}
)

func (s *tested1) Fields(enum func(name string)) {
	for _, name := range s.fields {
		enum(name)
	}
}

func Test_ForEach(t *testing.T) {
	require := require.New(t)

	t.Run("test ForEach with interface method", func(t *testing.T) {
		var tested ITested1 = &tested1{fields: []string{"a", "b", "c"}}

		line := ""
		ForEach(tested.Fields, func(d string) { line += d })

		require.Equal("abc", line)
	})

	t.Run("test ForEach with structure method", func(t *testing.T) {
		tested := tested1{fields: []string{"a", "b", "c"}}

		line := ""
		ForEach(tested.Fields, func(d string) { line += d })

		require.Equal("abc", line)
	})

	t.Run("test ForEach with naked slice", func(t *testing.T) {
		fields := []string{"a", "b", "c"}

		line := ""
		ForEach(Slice(fields), func(d string) { line += d })

		require.Equal("abc", line)
	})
}

func Test_FindFirst(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirst with interface method", func(t *testing.T) {
		var tested ITested1 = &tested1{fields: []string{"a", "b", "c"}}

		ok, data := FindFirst(tested.Fields, func(data string) bool { return data == "b" })
		require.True(ok)
		require.Equal("b", data)

		ok, data = FindFirst(tested.Fields, func(data string) bool { return data == "impossible" })
		require.False(ok)
		require.Empty(data)
	})

	t.Run("test FindFirst with structure method", func(t *testing.T) {
		tested := tested1{fields: []string{"a", "b", "c"}}

		ok, data := FindFirst(tested.Fields, func(data string) bool { return data == "b" })
		require.True(ok)
		require.Equal("b", data)

		ok, data = FindFirst(tested.Fields, func(data string) bool { return data == "impossible" })
		require.False(ok)
		require.Empty(data)
	})

	t.Run("test FindFirst with naked slice", func(t *testing.T) {
		fields := []string{"a", "b", "c"}

		ok, data := FindFirst(Slice(fields), func(data string) bool { return data == "b" })
		require.True(ok)
		require.Equal("b", data)

		ok, data = FindFirst(Slice(fields), func(data string) bool { return data == "impossible" })
		require.False(ok)
		require.Empty(data)
	})
}

func Test_FindData(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirstData with interface method", func(t *testing.T) {
		var tested ITested1 = &tested1{fields: []string{"a", "b", "c"}}

		ok, idx := FindFirstData(tested.Fields, "b")
		require.True(ok)
		require.Equal(1, idx)

		ok, idx = FindFirstData(tested.Fields, "impossible")
		require.False(ok)
		require.Less(idx, 0)
	})

	t.Run("test FindFirstData with structure method", func(t *testing.T) {
		tested := tested1{fields: []string{"a", "b", "c"}}

		ok, idx := FindFirstData(tested.Fields, "b")
		require.True(ok)
		require.Equal(1, idx)

		ok, idx = FindFirstData(tested.Fields, "impossible")
		require.False(ok)
		require.Less(idx, 0)
	})

	t.Run("test FindFirstData with naked slice", func(t *testing.T) {
		fields := []string{"a", "b", "c"}

		ok, idx := FindFirstData(Slice(fields), "b")
		require.True(ok)
		require.Equal(1, idx)

		ok, idx = FindFirstData(Slice(fields), "impossible")
		require.False(ok)
		require.Less(idx, 0)
	})
}

func Test_FindFirstError(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirstError with interface method", func(t *testing.T) {
		var tested ITested1 = &tested1{fields: []string{"a", "b", "c"}}

		data, err := FindFirstError(tested.Fields, func(data string) error {
			if data == "b" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", data)

		data, err = FindFirstError(tested.Fields, func(data string) error {
			if data == "impossible" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(data)
	})

	t.Run("test FindFirstError with structure method", func(t *testing.T) {
		tested := tested1{fields: []string{"a", "b", "c"}}

		data, err := FindFirstError(tested.Fields, func(data string) error {
			if data == "b" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", data)

		data, err = FindFirstError(tested.Fields, func(data string) error {
			if data == "impossible" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(data)
	})

	t.Run("test FindFirstError with naked slice", func(t *testing.T) {
		fields := []string{"a", "b", "c"}

		data, err := FindFirstError(Slice(fields), func(data string) error {
			if data == "b" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", data)

		data, err = FindFirstError(Slice(fields), func(data string) error {
			if data == "impossible" {
				return fmt.Errorf("error at %v", data)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(data)
	})
}

type (
	ITested2 interface {
		Fields(enum func(name string, data int))
	}
	tested2 struct {
		fields map[string]int
	}
)

func (s *tested2) Fields(enum func(name string, data int)) {
	for name, data := range s.fields {
		enum(name, data)
	}
}

func Test_ForEachMap(t *testing.T) {
	require := require.New(t)

	t.Run("test ForEachMap with interface method", func(t *testing.T) {
		var tested ITested2 = &tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		keys, values := "", 0
		ForEachMap(tested.Fields, func(k string, v int) { keys += k; values += v })

		require.Equal(3, len(keys))
		require.Contains(keys, "a")
		require.Contains(keys, "b")
		require.Contains(keys, "c")

		require.Equal(1+2+3, values)
	})

	t.Run("test ForEachMap with structure method", func(t *testing.T) {
		tested := tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		keys, values := "", 0
		ForEachMap(tested.Fields, func(k string, v int) { keys += k; values += v })

		require.Equal(3, len(keys))
		require.Contains(keys, "a")
		require.Contains(keys, "b")
		require.Contains(keys, "c")

		require.Equal(1+2+3, values)
	})

	t.Run("test ForEachMap with naked map", func(t *testing.T) {
		tested := map[string]int{"a": 1, "b": 2, "c": 3}

		keys, values := "", 0
		ForEachMap(Map(tested), func(k string, v int) { keys += k; values += v })

		require.Equal(3, len(keys))
		require.Contains(keys, "a")
		require.Contains(keys, "b")
		require.Contains(keys, "c")

		require.Equal(1+2+3, values)
	})
}

func Test_FindFirstMap(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirstMap with interface method", func(t *testing.T) {
		var tested ITested2 = &tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		ok, key, value := FindFirstMap(tested.Fields, func(k string, _ int) bool { return k == "b" })
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMap(tested.Fields, func(k string, _ int) bool { return k == "impossible" })
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMap with structure method", func(t *testing.T) {
		tested := tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		ok, key, value := FindFirstMap(tested.Fields, func(k string, _ int) bool { return k == "b" })
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMap(tested.Fields, func(k string, _ int) bool { return k == "impossible" })
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMap with naked map", func(t *testing.T) {
		fields := map[string]int{"a": 1, "b": 2, "c": 3}

		ok, key, value := FindFirstMap(Map(fields), func(k string, _ int) bool { return k == "b" })
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMap(Map(fields), func(k string, _ int) bool { return k == "impossible" })
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})
}

func Test_FindFirstMapKey(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirstMapKey with interface method", func(t *testing.T) {
		var tested ITested2 = &tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		ok, key, value := FindFirstMapKey(tested.Fields, "b")
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMapKey(tested.Fields, "impossible")
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMapKey with structure method", func(t *testing.T) {
		tested := tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		ok, key, value := FindFirstMapKey(tested.Fields, "b")
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMapKey(tested.Fields, "impossible")
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMapKey with naked map", func(t *testing.T) {
		fields := map[string]int{"a": 1, "b": 2, "c": 3}

		ok, key, value := FindFirstMapKey(Map(fields), "b")
		require.True(ok)
		require.Equal("b", key)
		require.Equal(2, value)

		ok, key, value = FindFirstMapKey(Map(fields), "impossible")
		require.False(ok)
		require.Empty(key)
		require.Empty(value)
	})
}

func Test_FindFirstMapError(t *testing.T) {
	require := require.New(t)

	t.Run("test FindFirstMapError with interface method", func(t *testing.T) {
		var tested ITested2 = &tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		key, value, err := FindFirstMapError(tested.Fields, func(k string, v int) error {
			if k == "b" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", key)
		require.Equal(2, value)

		key, value, err = FindFirstMapError(tested.Fields, func(k string, v int) error {
			if k == "impossible" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMapError with structure method", func(t *testing.T) {
		tested := tested2{fields: map[string]int{"a": 1, "b": 2, "c": 3}}

		key, value, err := FindFirstMapError(tested.Fields, func(k string, v int) error {
			if k == "b" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", key)
		require.Equal(2, value)

		key, value, err = FindFirstMapError(tested.Fields, func(k string, v int) error {
			if k == "impossible" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(key)
		require.Empty(value)
	})

	t.Run("test FindFirstMapError with naked map", func(t *testing.T) {
		tested := map[string]int{"a": 1, "b": 2, "c": 3}

		key, value, err := FindFirstMapError(Map(tested), func(k string, v int) error {
			if k == "b" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.Error(err)
		require.Equal("b", key)
		require.Equal(2, value)

		key, value, err = FindFirstMapError(Map(tested), func(k string, v int) error {
			if k == "impossible" {
				return fmt.Errorf("error at %v: %v", k, v)
			}
			return nil
		})
		require.NoError(err)
		require.Empty(key)
		require.Empty(value)
	})
}

func TestForEachError(t *testing.T) {
	require := require.New(t)
	t.Run("no error", func(t *testing.T) {
		err := ForEachError(func(enum func(str string)) {
			enum("1")
		}, func(s string) error {
			require.Equal("1", s)
			return nil
		})
		require.NoError(err)
	})

	t.Run("error", func(t *testing.T) {
		testErr := errors.New("test error")
		expected := ""
		err := ForEachError(func(enum func(str string)) {
			enum("1")
			enum("2")
			enum("3")
			enum("4")
		}, func(s string) error {
			if s == "3" {
				return testErr
			}
			expected += s
			return nil
		})
		require.ErrorIs(err, testErr)
		require.Equal("12", expected)
	})
}
