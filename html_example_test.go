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
	"bytes"
	"fmt"
)

func ExampleEscapeString() {
	fmt.Println(EscapeString(`<a href="https://www.google.com/search?q=something&ie=utf-8">Google Search</a>`))

	// Output:
	// &lt;a href=&quot;https://www.google.com/search?q=something&amp;ie=utf-8&quot;&gt;Google Search&lt;/a&gt;
}

func ExampleWriteEscapedString() {
	var buf bytes.Buffer
	WriteEscapedString(&buf, `<a href="https://www.google.com/search?q=something&ie=utf-8">Google Search</a>`)
	fmt.Println(buf.String())

	// Output:
	// &lt;a href=&quot;https://www.google.com/search?q=something&amp;ie=utf-8&quot;&gt;Google Search&lt;/a&gt;
}

func ExampleParseEntity() {
	fmt.Println(ParseEntity("&quot;"))
	fmt.Println(ParseEntity("&#x22;"))
	fmt.Println(ParseEntity("&#34;"))
	fmt.Println(ParseEntity("&#x00000022;"))
	fmt.Println(ParseEntity("&#00000034;"))
	fmt.Println(ParseEntity("&#x000000022;"))
	fmt.Println(ParseEntity("&#000000034;"))
	fmt.Println(ParseEntity("&#98765432;"))

	// Output:
	// " 6
	// " 6
	// " 5
	// " 12
	// " 11
	//  0
	//  0
	// � 11
}

func ExampleReplaceEntities() {
	fmt.Println(ReplaceEntities("&nbsp; &amp; &copy; &AElig; &Dcaron; &frac34; &HilbertSpace; &DifferentialD; &ClockwiseContourIntegral; &#35; &#1234; &#992; &#98765432; &#X22; &#XD06; &#xcab; &nbsp &x; &#; &#x; &ThisIsWayTooLongToBeAnEntityIsntIt; &hi?; &copy &MadeUpEntity;"))

	// Output:
	//   & © Æ Ď ¾ ℋ ⅆ ∲ # Ӓ Ϡ � " ആ ಫ &nbsp &x; &#; &#x; &ThisIsWayTooLongToBeAnEntityIsntIt; &hi?; &copy &MadeUpEntity;
}
