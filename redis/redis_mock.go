package redis

type RedisConnectionMock struct {
	ReturnErrorGet error
	ReturnErrorSet error
	ReturnErrorDel error
	Key            string
	Value          string
	ReturnValue    string
}

func (r *RedisConnectionMock) GET(key string) (string, error) {
	r.Key = key
	if r.ReturnErrorGet != nil {
		return "", r.ReturnErrorGet
	}
	return r.ReturnValue, nil
}

func (r *RedisConnectionMock) SET(key, value string) error {
	r.Key = key
	r.Value = value
	if r.ReturnErrorSet != nil {
		return r.ReturnErrorSet
	}
	return nil
}
