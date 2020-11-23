package apifields

type Field struct {
	Key   string
	Value interface{}
}

func String(key string, value string) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}
