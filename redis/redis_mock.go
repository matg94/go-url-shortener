package redis

type RedisConnectionMock struct {
	ReturnErrorGet error
	ReturnErrorSet error
	Key            string
	Value          string
	ReturnValue    string
}

func CreateRedisMock(GetError, SetError error, ReturnValue string) *RedisConnectionMock {
	return &RedisConnectionMock{
		ReturnErrorGet: GetError,
		ReturnErrorSet: SetError,
		ReturnValue:    ReturnValue,
	}
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
