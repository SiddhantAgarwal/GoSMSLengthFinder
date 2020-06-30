package smslength

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSMSParts(t *testing.T) {
	testCases := []struct {
		name             string
		expectedSMSParts int64
		stringToTest     string
	}{
		{
			name:             "Simple ASCII message body",
			stringToTest:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			expectedSMSParts: 1,
		},
		{
			name:             "Simple ASCII message body, multipart",
			stringToTest:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			expectedSMSParts: 3,
		},
		{
			name:             "Extended ASCII message body, multipart",
			stringToTest:     "Lorem ipsum dolor sit amet, ~~~~~~~~~~~~~~~~~~~~~consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			expectedSMSParts: 4,
		},
		{
			name:             "Unicode message body",
			stringToTest:     `ăZʍ𦞺9r𔡛)(Δݷ֣퐘`,
			expectedSMSParts: 1,
		},
		{
			name:             "Unicode message body, multipart",
			stringToTest:     `ăZʍ𦞺9r𔡛)(Δݷ֣퐘Řaas􂵜꓄悗償%O򊿟夒Д--2䴛槫𐻵춡󣛃<ꣻئ񯕱񚲪Ʌ뢉񤲔՛􏅛􁐜󯞒؀,خs񚺣򳼙sϞ񴧆:ʣ򆝳ጳ(򶳫蝹򝚿ӂ交柪?٦쳕磏͒qtԜ󣜐҈쏒W㗁ԂD欕D"憬ˆ!Ӹⵗ򏘯𸎫ɝ󁦡?Ӯ]܂ҥ4󘒙򚲣읊Ζ􉵑}󪆓豱շښܾ纒藯b䬢ܫtؕjx񑾖ńֺ<𝚝􉆢ˬîґNڧ󓺻h񃴵󍣧}Ǜ贈#<ᜏ񞹚ނ'򧿲b]˽刺񽞙Ѐ𨦨2ͫ֞􇓄쪧𙅝򳉎`,
			expectedSMSParts: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSMSParts, GetSMSParts(tc.stringToTest))
		})
	}
}
