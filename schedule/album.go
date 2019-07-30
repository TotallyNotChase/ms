package schedule

// Album is a record to be scheduled. Includes whether it's been listened
// or rated.
type Album struct {
	Name         string `json:"name"`
	FirstListen  bool   `json:"first_listen"`
	SecondListen bool   `json:"second_listen"`
	ThirdListen  bool   `json:"third_listen"`
	Rated        bool   `json:"rated"`
}
