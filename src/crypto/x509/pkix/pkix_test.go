package pkix

import (
	"reflect"
	"testing"
)

var (
	oidDomainComponent = []int{0, 9, 2342, 19200300, 100, 1, 25}
	oidUserId          = []int{0, 9, 2342, 19200300, 100, 1, 1}
)

var nameTestData = []RDNSequence{
	{ // #0
		{{Type: oidCountry, Value: "EN"}},
		{{Type: oidProvince, Value: "StateOrProvinceName"}},
		{{Type: oidLocality, Value: "LocalityName"}},
		{{Type: oidOrganization, Value: "OrganizationName"}},
		{{Type: oidOrganizationalUnit, Value: "OrganizationalUnitName"}},
		{{Type: oidCommonName, Value: "CommonName"}},
	},
	{ // #1
		{{Type: oidCountry, Value: "CN"}},
		{{Type: oidProvince, Value: "StateOrProvinceName"}},
		{{Type: oidLocality, Value: "LocalityName"}},
		{{Type: oidStreetAddress, Value: "StreetAddress"}},
		{{Type: oidCommonName, Value: "CommonName"}},
	},
	{ // #2 see rfc4514 section-4 Examples
		{{Type: oidDomainComponent, Value: "net"}},
		{{Type: oidDomainComponent, Value: "example"}},
		{{Type: oidUserId, Value: "jsmith"}},
	},
	{ // #3 see rfc4514 section-4 Examples
		{{Type: oidDomainComponent, Value: "net"}},
		{{Type: oidDomainComponent, Value: "example"}},
		{ // multi-valued RDN
			{Type: oidOrganizationalUnit, Value: "Sales"},
			{Type: oidCommonName, Value: "J.  Smith"},
		},
	},
	{ // #4 see rfc4514 section-4 Examples
		{{Type: oidDomainComponent, Value: "net"}},
		{{Type: oidDomainComponent, Value: "example"}},
		{{Type: oidCommonName, Value: `James "Jim" Smith, III`}},
	},
	{ // #5 see rfc4514 section-4 Examples
		{{Type: oidDomainComponent, Value: "net"}},
		{{Type: oidDomainComponent, Value: "example"}},
		{{Type: oidCommonName, Value: `Before\x0dAfter`}},
	},
	{ // #6 Chinese
		{{Type: oidCountry, Value: "CN"}},
		{{Type: oidProvince, Value: "浙江"}},
		{{Type: oidLocality, Value: "杭州"}},
		{{Type: oidOrganization, Value: "公司或单位"}},
		{{Type: oidOrganizationalUnit, Value: "部门或科室"}},
		{{Type: oidCommonName, Value: "使用人"}},
	},
	{ // #7 RDN Sequence
		{{Type: oidCountry, Value: "US"}},
		{{Type: oidOrganization, Value: "Company Group"}},
		{{Type: oidOrganization, Value: "Parent Company"}},
		{{Type: oidOrganization, Value: "SubCompany"}},
		{{Type: oidOrganizationalUnit, Value: "Department"}},
		{{Type: oidOrganizationalUnit, Value: "SubDepartment"}},
		{{Type: oidCommonName, Value: "Name"}},
	},
}

func TestName(t *testing.T) {
	for i, test := range nameTestData {
		var name Name
		name.FillFromRDNSequence(&test)
		ret := name.ToRDNSequence()
		if !reflect.DeepEqual(ret, test) {
			t.Errorf("#%d: Failed", i)
		}
	}
}
