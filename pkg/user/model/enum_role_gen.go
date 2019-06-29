// Code generated by "enumer -type=Role -text -json -trimprefix=Role -transform=snake -output=enum_role_gen.go"; DO NOT EDIT.

//
package model

import (
	"encoding/json"
	"fmt"
)

const _RoleName = "adminuser"

var _RoleIndex = [...]uint8{0, 5, 9}

func (i Role) String() string {
	i -= 1
	if i >= Role(len(_RoleIndex)-1) {
		return fmt.Sprintf("Role(%d)", i+1)
	}
	return _RoleName[_RoleIndex[i]:_RoleIndex[i+1]]
}

var _RoleValues = []Role{1, 2}

var _RoleNameToValueMap = map[string]Role{
	_RoleName[0:5]: 1,
	_RoleName[5:9]: 2,
}

// RoleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoleString(s string) (Role, error) {
	if val, ok := _RoleNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Role values", s)
}

// RoleValues returns all values of the enum
func RoleValues() []Role {
	return _RoleValues
}

// IsARole returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Role) IsARole() bool {
	for _, v := range _RoleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Role
func (i Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Role
func (i *Role) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Role should be a string, got %s", data)
	}

	var err error
	*i, err = RoleString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Role
func (i Role) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Role
func (i *Role) UnmarshalText(text []byte) error {
	var err error
	*i, err = RoleString(string(text))
	return err
}
