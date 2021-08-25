package data

import "time"

type JsonTime struct {
	time.Time
}

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}

	jt.Time = t
	return nil
}