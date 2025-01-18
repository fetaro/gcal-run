package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major  int
	minor  int
	bugfix int
}

func (v *Version) IsNewer(other *Version) bool {
	return v.major*10000+v.minor*100+v.bugfix > other.major*10000+other.minor*100+other.bugfix
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.bugfix)
}

func ParseVersionStr(versionStr string) (*Version, error) {
	// versionStrから先頭のvを取り除く
	versionStr = strings.TrimPrefix(versionStr, "v")
	// 末尾から改行を取り除く
	versionStr = strings.TrimSuffix(versionStr, "\n")
	v := strings.Split(versionStr, ".")
	if len(v) != 3 {
		return nil, fmt.Errorf("len(v) != 3: %s", versionStr)
	}
	major, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %d, err: %v", major, err)
	}
	minor, err := strconv.Atoi(v[1])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %d, err: %v", minor, err)
	}
	bugfix, err := strconv.Atoi(v[2])
	if err != nil {
		return nil, fmt.Errorf("invalid bugfix version: %d, err: %v", bugfix, err)
	}
	return &Version{
		major:  major,
		minor:  minor,
		bugfix: bugfix,
	}, nil
}
