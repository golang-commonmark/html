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
	"html"
	"testing"
)

var escapeBenchmarkStrings = []string{
	"",
	"a b c d e f g h i j k l m n o p q r s t u v w x y z ",
	`<a href="http://google.com?q=&">google.com</a>`,
}

func BenchmarkEscapeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range escapeBenchmarkStrings {
			EscapeString(s)
		}
	}
}

func BenchmarkEscapeStringStdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range escapeBenchmarkStrings {
			html.EscapeString(s)
		}
	}
}
