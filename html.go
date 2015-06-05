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
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/opennota/byteutil"
)

const runeError = string(utf8.RuneError)

var htmlEscapeReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	`"`, "&quot;",
)

func EscapeString(s string) string {
	return htmlEscapeReplacer.Replace(s)
}

func WriteEscapedString(w io.Writer, s string) error {
	_, err := htmlEscapeReplacer.WriteString(w, s)
	return err
}

func isValidEntityCode(c int64) bool {
	switch {
	case !utf8.ValidRune(rune(c)):
		return false

	// never used
	case c >= 0xfdd0 && c <= 0xfdef:
		return false
	case c&0xffff == 0xffff || c&0xffff == 0xfffe:
		return false
	// control codes
	case c >= 0x00 && c <= 0x08:
		return false
	case c == 0x0b:
		return false
	case c >= 0x0e && c <= 0x1f:
		return false
	case c >= 0x7f && c <= 0x9f:
		return false
	}

	return true
}

func ParseEntity(s string) (string, int) {
	st := 0
	var n int

	for i := 1; i < len(s); i++ {
		b := s[i]

		switch st {
		case 0: // initial state
			switch {
			case b == '#':
				st = 1
			case byteutil.IsLetter(b):
				n = 1
				st = 2
			default:
				return "", 0
			}

		case 1: // &#
			switch {
			case b == 'x' || b == 'X':
				st = 3
			case byteutil.IsDigit(b):
				n = 1
				st = 4
			default:
				return "", 0
			}

		case 2: // &q
			switch {
			case byteutil.IsAlphaNum(b):
				n++
				if n > 31 {
					return "", 0
				}
			case b == ';':
				if e, ok := entities[s[i-n:i]]; ok {
					return e, i + 1
				}
				return "", 0
			default:
				return "", 0
			}

		case 3: // &#x
			switch {
			case byteutil.IsHexDigit(b):
				n = 1
				st = 5
			default:
				return "", 0
			}

		case 4: // &#0
			switch {
			case byteutil.IsDigit(b):
				n++
				if n > 8 {
					return "", 0
				}
			case b == ';':
				c, _ := strconv.ParseInt(s[i-n:i], 10, 32)
				if !isValidEntityCode(c) {
					return runeError, i + 1
				}
				return string(rune(c)), i + 1
			default:
				return "", 0
			}

		case 5: // &#x0
			switch {
			case byteutil.IsHexDigit(b):
				n++
				if n > 8 {
					return "", 0
				}
			case b == ';':
				c, err := strconv.ParseInt(s[i-n:i], 16, 32)
				if err != nil {
					return runeError, i + 1
				}
				if !isValidEntityCode(c) {
					return runeError, i + 1
				}
				return string(rune(c)), i + 1
			default:
				return "", 0
			}
		}
	}

	return "", 0
}

func ReplaceEntities(s string) string {
	i := strings.IndexByte(s, '&')
	if i < 0 {
		return s
	}

	anyChanges := false
	var entityStr string
	var entityLen int
	for i < len(s) {
		if s[i] == '&' {
			entityStr, entityLen = ParseEntity(s[i:])
			if entityLen > 0 {
				anyChanges = true
				break
			}
		}
		i++
	}

	if !anyChanges {
		return s
	}

	buf := make([]byte, len(s)-entityLen+len(entityStr))
	copy(buf[:i], s)
	n := copy(buf[i:], entityStr)
	j := i + n
	i += entityLen
	for i < len(s) {
		b := s[i]
		if b == '&' {
			entityStr, entityLen = ParseEntity(s[i:])
			if entityLen > 0 {
				n = copy(buf[j:], entityStr)
				j += n
				i += entityLen
				continue
			}
		}

		buf[j] = b
		j++
		i++
	}

	return string(buf[:j])
}
