package storage

type StorageMock struct {
	data map[string]interface{}
}

func (s *StorageMock) GetValue(key string) interface{} {
	if v, ok := s.data[key]; ok {
		return v
	}
	return nil
}
