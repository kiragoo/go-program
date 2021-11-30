package nonil

type Nonil struct {
	Name  string `json:"name"`
	Value *int32 `json:"value"`
}

func NonilFn(name string, value int32) Nonil {

	return Nonil{
		Name:  name,
		Value: &value,
	}
}
