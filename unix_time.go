package atomicasset

import (
	"encoding/json"
	"strconv"
	"time"
)

// UnixTime is a simple wrapper to handle unix timestamps in json data.
type UnixTime int64

func (ts *UnixTime) UnmarshalJSON(b []byte) error {
	var i int64

	// "borrowed" from "gopkg.in/guregu/null.v4" abit.
	if err := json.Unmarshal(b, &i); err != nil {

		// If unmarshal to int64 fails, we assume that its a numeric string.
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}

		// Then we need to parse the string into int64
		i, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
	}

	*ts = UnixTime(i)
	return nil
}

func (ts UnixTime) Time() time.Time {
	v := int64(ts)
	return time.Unix(v/1000, v%1000).UTC()
}
