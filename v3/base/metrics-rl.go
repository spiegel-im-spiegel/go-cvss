package base

import "strings"

//RemediationLevel is metric type for Temporal Metrics
type RemediationLevel int

//Constant of RemediationLevel result
const (
	RemediationLevelNotDefined RemediationLevel = iota
	RemediationLevelOfficialFix
	RemediationLevelTemporaryFix
	RemediationLevelWorkaround
	RemediationLevelUnavailable
)

var RemediationLevelMap = map[RemediationLevel]string{
	RemediationLevelNotDefined:   "X",
	RemediationLevelOfficialFix:  "O",
	RemediationLevelTemporaryFix: "T",
	RemediationLevelWorkaround:   "W",
	RemediationLevelUnavailable:  "U",
}

var RemediationLevelValueMap = map[RemediationLevel]float64{
	RemediationLevelNotDefined:   1,
	RemediationLevelOfficialFix:  0.95,
	RemediationLevelTemporaryFix: 0.96,
	RemediationLevelWorkaround:   0.97,
	RemediationLevelUnavailable:  1,
}

//GetRemediationLevel returns result of RemediationLevel metric
func GetRemediationLevel(s string) RemediationLevel {
	s = strings.ToUpper(s)
	for k, v := range RemediationLevelMap {
		if s == v {
			return k
		}
	}
	return RemediationLevelNotDefined
}

func (ai RemediationLevel) String() string {
	if s, ok := RemediationLevelMap[ai]; ok {
		return s
	}
	return ""
}

//Value returns value of RemediationLevel metric
func (ai RemediationLevel) Value() float64 {
	if v, ok := RemediationLevelValueMap[ai]; ok {
		return v
	}
	return 1
}

//IsDefined returns false if undefined result value of metric
func (ai RemediationLevel) IsDefined() bool {
	return ai != RemediationLevelNotDefined
}

/* Copyright by Florent Viel, 2020 */
