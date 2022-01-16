package unifynil

import "reflect"

type Option func(*Config)

type Config struct {
	sliceToNil   bool
	sliceToEmpty bool
	mapToNil     bool
	mapToEmpty   bool
}

func SliceToNil() Option {
	return func(c *Config) {
		if c.sliceToEmpty {
			panic("both SliceToNil and SliceToEmpty passed")
		}
		c.sliceToNil = true
	}
}

func SliceToEmpty() Option {
	return func(c *Config) {
		if c.sliceToNil {
			panic("both SliceToNil and SliceToEmpty passed")
		}
		c.sliceToEmpty = true
	}
}

func MapToNil() Option {
	return func(c *Config) {
		if c.mapToEmpty {
			panic("both MapToNil and MapToEmpty passed")
		}
		c.mapToNil = true
	}
}

func MapToEmpty() Option {
	return func(c *Config) {
		if c.mapToNil {
			panic("both MapToNil and MapToEmpty passed")
		}
		c.mapToEmpty = true
	}
}

func Unify(objPtr interface{}, opts ...Option) {
	var c Config
	for _, o := range opts {
		o(&c)
	}
	unify(reflect.ValueOf(objPtr), &c)
}

func unify(obj reflect.Value, c *Config) {
	switch obj.Kind() {
	case reflect.Ptr:
		unify(obj.Elem(), c)
	case reflect.Slice:
		if obj.Len() == 0 {
			if c.sliceToNil && !obj.IsNil() {
				obj.Set(reflect.New(obj.Type()).Elem())
			} else if c.sliceToEmpty && obj.IsNil() {
				obj.Set(reflect.MakeSlice(obj.Type(), 0, 0))
			}
		} else {
			for i := 0; i < obj.Len(); i++ {
				unify(obj.Index(i), c)
			}
		}
	case reflect.Array:
		for i := 0; i < obj.Len(); i++ {
			unify(obj.Index(i), c)
		}
	case reflect.Map:
		if obj.Len() == 0 {
			if c.mapToNil && !obj.IsNil() {
				obj.Set(reflect.New(obj.Type()).Elem())
			} else if c.mapToEmpty && obj.IsNil() {
				obj.Set(reflect.MakeMap(obj.Type()))
			}
		} else {
			iter := obj.MapRange()
			clonedValue := reflect.New(obj.Type().Elem()).Elem()
			for iter.Next() {
				// Copy the value, because a value inside a map
				// can not be modified in place.
				clonedValue.Set(iter.Value())
				unify(clonedValue, c)
				obj.SetMapIndex(iter.Key(), clonedValue)
			}
		}
	case reflect.Struct:
		for i := 0; i < obj.NumField(); i++ {
			if obj.Type().Field(i).PkgPath != "" {
				// Not exported.
				continue
			}
			unify(obj.Field(i), c)
		}
	}
}
