// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package html

import (
	"testing"
)

var parseEntityTests = []struct {
	in     string
	length int
	str    string
}{
	{in: "&;"},
	{in: "&"},
	{in: "&@;"},
	{in: "&@"},
	{in: "&#;"},
	{in: "&#"},
	{in: "&##;"},
	{in: "&##"},
	{in: "&#0"},
	{in: "&#09"},
	{in: "&#0a;"},
	{in: "&#0a"},
	{in: "&#34;", length: 5, str: `"`},
	{in: "&#98765432;", length: 11, str: RuneError},
	{in: "&#9;", length: 4, str: "\t"},
	{in: "&ffffffffffffffffuuuuuuuuuuuuuuuu;"},
	{in: "&ffffffffffffffffuuuuuuuuuuuuuuuu"},
	{in: "&q;"},
	{in: "&q"},
	{in: "&q#;"},
	{in: "&q#"},
	{in: "&Q"},
	{in: "&q0;"},
	{in: "&q0"},
	{in: "&qu;"},
	{in: "&qu"},
	{in: "&quot;", length: 6, str: `"`},
	{in: "&#x;"},
	{in: "&#x"},
	{in: "&#X;"},
	{in: "&#X"},
	{in: "&#x0"},
	{in: "&#x0@;"},
	{in: "&#x0@"},
	{in: "&#X0"},
	{in: "&#x000000001;"},
	{in: "&#x000000001"},
	{in: "&#x09"},
	{in: "&#x09;", length: 6, str: "\t"},
	{in: "&#X09;", length: 6, str: "\t"},
	{in: "&#x0a;", length: 6, str: "\n"},
	{in: "&#x0A;", length: 6, str: "\n"},
	{in: "&#X0a;", length: 6, str: "\n"},
	{in: "&#X0A;", length: 6, str: "\n"},
	{in: "&#x22;", length: 6, str: `"`},
	{in: "&#x7fffffff;", length: 12, str: RuneError},
	{in: "&#xa"},
	{in: "&#xA"},
	{in: "&#Xa"},
	{in: "&#XA"},
	{in: "&#xffffffff"},
	{in: "&#xfffffffff"},
	{in: "&#xffffffff;", length: 12, str: RuneError},
	{in: "&#xG;"},
	{in: "&#xG"},
	{in: "&#XG;"},
	{in: "&#XG"},
}

var replaceEntitiesTests = []struct {
	in  string
	out string // empty means == in
}{
	{"", ""},
	{"a b c d e f g h i j k l m n o p q r s t u v w x y z", ""},
	{"&", ""},
	{"&nbsp; &amp; &copy; &quot;", `  & © "`},
	{"&#x22; &#X22; &#34;", `" " "`},
	{"&x; &#; &#x; &ThisIsWayTooLongToBeAnEntityIsntIt; &hi?; &copy &MadeUpEntity;", ""},
}

func TestParseEntity(t *testing.T) {
	for _, tc := range parseEntityTests {
		str, length := ParseEntity(tc.in)
		if str != tc.str || length != tc.length {
			t.Errorf("ParseEntity(%q): want %d,%q, got %d,%q", tc.in,
				tc.length, tc.str, length, str)
		}
	}
}

func TestReplaceEntities(t *testing.T) {
	for _, tc := range replaceEntitiesTests {
		got := ReplaceEntities(tc.in)
		want := tc.out
		if want == "" {
			want = tc.in
		}
		if got != want {
			t.Errorf("ReplaceEntities(%q): want %q, got %q", tc.in, want, got)
		}
	}
}
