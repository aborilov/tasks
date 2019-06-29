// Code generated by "enumer -type=Status -text -json -trimprefix=Status -transform=snake -output=enum_status_gen.go"; DO NOT EDIT.

//
package model

import (
	"encoding/json"
	"fmt"
)

const _StatusName = "newin_progresscompletedarchived"

var _StatusIndex = [...]uint8{0, 3, 14, 23, 31}

func (i Status) String() string {
	i -= 1
	if i >= Status(len(_StatusIndex)-1) {
		return fmt.Sprintf("Status(%d)", i+1)
	}
	return _StatusName[_StatusIndex[i]:_StatusIndex[i+1]]
}

var _StatusValues = []Status{1, 2, 3, 4}

var _StatusNameToValueMap = map[string]Status{
	_StatusName[0:3]:   1,
	_StatusName[3:14]:  2,
	_StatusName[14:23]: 3,
	_StatusName[23:31]: 4,
}

// StatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StatusString(s string) (Status, error) {
	if val, ok := _StatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Status values", s)
}

// StatusValues returns all values of the enum
func StatusValues() []Status {
	return _StatusValues
}

// IsAStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Status) IsAStatus() bool {
	for _, v := range _StatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Status
func (i Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Status
func (i *Status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Status should be a string, got %s", data)
	}

	var err error
	*i, err = StatusString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Status
func (i Status) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Status
func (i *Status) UnmarshalText(text []byte) error {
	var err error
	*i, err = StatusString(string(text))
	return err
}
