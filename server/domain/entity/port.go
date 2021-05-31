package entity

type Port struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

func (f *Port) BeforeSave() {

}

func (f *Port) Prepare() {
}

func (f *Port) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	return errorMessages
}
