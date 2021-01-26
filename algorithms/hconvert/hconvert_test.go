package hconvert_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/snhilde/dsa/algorithms/hconvert"
)

// testSet holds the information for running decoding/encoding tests.
type testSet struct {
	// Character set to use for the conversion.
	charSet hconvert.CharSet

	// Name of the character set.
	setName string

	// Starting and ending data for the conversion process.
	convertFrom string
	convertTo   string
}

const (
	// This character is not in any of the standard character sets. Decoding with this should fail.
	invalidChar = "∞"

	// We will use this data for the standard conversion tests.
	inputText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In ut arcu fermentum, pellentesque leo sit amet, aliquam nunc. Sed maximus tincidunt libero, in volutpat lectus tincidunt quis."
)

var (

	// These are the common tests for decoding/encoding using the standard character sets.
	standardTests = []testSet{
		{hconvert.ASCIICharSet(), "ASCII", inputText, inputText},
		{hconvert.Base2CharSet(), "Base2", inputText, "1001100110111111100101100101110110101000001101001111000011100111110101110110101000001100100110111111011001101111111001001000001110011110100111101000100000110000111011011100101111010001011000100000110001111011111101110111001111001011100011111010011001011110100111010111100100100000110000111001001101001111000011010011110011110001111010011101110110011101000001100101110110011010011110100010111001000001001001110111001000001110101111010001000001100001111001011000111110101010000011001101100101111001011011011100101110111011101001110101110110101011000100000111000011001011101100110110011001011101110111010011001011110011111000111101011100101010000011011001100101110111101000001110011110100111101000100000110000111011011100101111010001011000100000110000111011001101001111000111101011100001110110101000001101110111010111011101100011010111001000001010011110010111001000100000110110111000011111000110100111011011110101111001101000001110100110100111011101100011110100111001001110101110111011101000100000110110011010011100010110010111100101101111010110001000001101001110111001000001110110110111111011001110101111010011100001100001111010001000001101100110010111000111110100111010111100110100000111010011010011101110110001111010011100100111010111011101110100010000011100011110101110100111100110101110"},
		{hconvert.Base4CharSet(), "Base4", inputText, "21212333211211312220031033003213311312220030212333121233321020032132213220200300323130233101120200301323331313033023203322121132213113210200300321031033003103303301322131312131001211312122132202321001021313020032233101001201321120332222003031211321123130232323221311312223010013003023230312121131313103023303320331130222003121211313220032132213220200300323130233101120200300323031033013223201312220031313113131203113020022132113020200312320133012213123311321220032212213131203310321032232323220200312122130112113211233112020031032321001312313323032233103201201322020031212113013310322330310013103103232301322130213113131310100130132232213212232"},
		{hconvert.Base8CharSet(), "Base8", inputText, "1146774545665015170347656650144677315771101636475040607334572130406173756717134372313647274440607115170323636172356635014566323642710111671016572101417130765203154571333456735165665304070313546631356723136370753452033145675016364750406073345721304060731517075341665015672735432710123627104066703706473365715016464735436471165673504066323426274557261015167101666773165723414172101546270764727464072323566172344727356420343656474656"},
		{hconvert.Base10CharSet(), "Base10", inputText, "3200405597824003522505771728586104169333370877702161880377073123898110960433541903784943665014732691140682556014912685675470142953725501153330966390692810975470167645925149984556159922699475752626216066019650933308946469529432950624859701121387850453752975178851307262250538505359131489464589547994472455871352206792025578123923622137470177046375324362980632949120863179247367018291820974"},
		{hconvert.Base16CharSet(), "Base16", inputText, "99BF965DA834F0E7D76A0C9BF66FE4839E9E8830EDCBD1620C7BF773CB8FA65E9D7920C3934F0D3CF1E9DD9D065D9A7A2E4127720EBD1061E58FAA0CD9796DCBBBA75DAB1070CBB3665DDD32F3E3D72A0D9977A0E7A7A20C3B72F458830ECD3C7AE1DA83775DD8D720A797220DB87C69DBD79A0E9A7763D393AEEE883669C59796F58834EE41DB7ECEBD3861E883665C7D3AF341D34EEC7A7275DDD1071EBA79AE"},
		{hconvert.Base32CharSet(), "Base32", inputText, "EZX6LF3KBU6DT5O2QMTP3G7ZEDT2PIQMHNZPIWEDD365Z4XD5GL2OXSIGDSNHQ2PHR5HOZ2BS5TJ5C4QJHOIHL2EDB4WH2UDGZPFW4XO5HLWVRA4GLWNTF3XJS6PR5OKQNTF32BZ5HUIGDW4XULCBQ5TJ4PLQ5VA3XLXMNOIFHS4RA3OD4NHN5PGQOTJ3WHU4TV3XIQNTJYWLZN5MIGTXEDW36Z26TQYPIQNTFY7J26NA5GTXMPJZHLXORA4PLU6NO"},
		{hconvert.Base36CharSet(), "Base36", inputText, "YUJQE5D1IBM5UXSO4EGVJVLT4MYRIOZ3UUBLTXA72UCECVYIO1NJB9BLY2NMCIOPKPU2UK6GFX5PMPMQ5K3XWS1GP4VHL0SC5PXYY9H9OS3CXHO0FESZXHSDMP8ZBJST1BKXZ7P66KUXWSNX2QO5UM95QA16Y1JJ1EC255YBGLF6QY5N4N02SRCMBZXXSO53RI2587XPQPV9Q6JDR5DKMY4P02L1C24NF8VK1DGNL341KO3B1DPXILCQ6"},
		{hconvert.Base58CharSet(), "Base58", inputText, "MdUMk3Xekk5Wygh2h7PgthLbfAMVhZYUL7tPqNUXhDQQaESYuojAFx4tGXqiLW8aYvG6WqtxMXJzPhEEGTgM5qv2dYzS4dhDrwKoCE6RuK5Zm5Twwo9VXn6SXvUTtS3GyE2GVjnYkZPzfXPkLaDpKz767tbi1RhMhkkEP3UrC9RLKETwD7rnpRTFswUhznVdZp1v3Xz2UXNX9sMNxmj39FNv3WLD"},
		{hconvert.Base62CharSet(), "Base62", inputText, "COV60nOF4iZpgoe0ryEXcKGmm81zDNb6FGjwxcMvnoirzQnViOA2aTZXcBGc2y0VmPnRRPrFNrXCpfJgRl7ZEcliv21ZemSsH5CiwJ17mFsg3bCZYUKfsPdtdh4EgVzYiugDDgOFjYOC54LxV9UTCnh1fRtseob64ZX8itRwvnBBCYr4EPKpBLRXLmdk9X7LJyWyv3PTU32nXwq55Z4HZn2YO"},
		{hconvert.Base64CharSet(), "Base64", inputText, "Jm/ll2oNPDn12oMm/Zv5IOenogw7cvRYgx793PLj6ZenXkgw5NPDTzx6d2dBl2aei5BJ3IOvRBh5Y+qDNl5bcu7p12rEHDLs2Zd3TLz49cqDZl3oOenogw7cvRYgw7NPHrh2oN3XdjXIKeXIg24fGnb15oOmndj05Ou7og2acWXlvWINO5B237OvThh6INmXH0680HTTux6cnXd0Qceunmu"},
		{hconvert.Base64URLCharSet(), "Base64URL", inputText, "Jm_ll2oNPDn12oMm_Zv5IOenogw7cvRYgx793PLj6ZenXkgw5NPDTzx6d2dBl2aei5BJ3IOvRBh5Y-qDNl5bcu7p12rEHDLs2Zd3TLz49cqDZl3oOenogw7cvRYgw7NPHrh2oN3XdjXIKeXIg24fGnb15oOmndj05Ou7og2acWXlvWINO5B237OvThh6INmXH0680HTTux6cnXd0Qceunmu"},
		{hconvert.ASCII85CharSet(), "ASCII85", inputText, "JfuFW*C5A+1%G*'8.+GH(`*]iGL5.T^7lm3\"[)F3?4L87A#ke2r`%/bC!GrN;F/GJeoXb#/+Jk`AG[`)iL2)?>Jk;kD:;W4HOAoh&patAg`.>S95prCM;'A8&N>)(f?_9^O(a\\ThK/E9DWu!J*3-<[1/6BNI84,6#R\\k\"L7>Ta1Tgj/+q9MCBcsKa?5.0DF;OEVB?9+Ik"},
		{hconvert.Z85CharSet(), "Z85", inputText, "F/#BS9ykwag4C96ndaCD7-9Y&CHkdPZm()i1W8BiujHnmw2>!h@-4e+y0C@JqBeCF!]T+2eaF>-wCW-8&Hh8utF>q>zpqSjDKw]?5{:$w*-dtOok{@yIq6wn5Jt87/u.oZK7:XP?GeAozS#0F9icrWgelxJEnjbl2NX>1HmtP:gP*<ea}oIyx=%G:ukdfzBqKARxuoaE>"},
	}
)

func TestNewConverter(t *testing.T) {
	t.Parallel()

	// We should be able to create a Converter with any combination of valid or invalid character
	// sets.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	if converter := hconvert.NewConverter(dec, enc); !reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create empty converter")
	}

	dec = hconvert.CharSet{}
	enc = hconvert.Base16CharSet()
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only encoding character set")
	}

	dec = hconvert.Base32CharSet()
	enc = hconvert.CharSet{}
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with only decoding character set")
	}

	dec = hconvert.Base32CharSet()
	enc = hconvert.Base16CharSet()
	if converter := hconvert.NewConverter(dec, enc); reflect.DeepEqual(converter, hconvert.Converter{}) {
		t.Error("Failed to create converter with both encoding and decoding character sets")
	}
}

func TestBad(t *testing.T) {
	t.Parallel()

	var converter *hconvert.Converter

	if err := converter.SetDecodeCharSet(hconvert.Base10CharSet()); err == nil {
		t.Error("Failed bad object test for SetDecodeCharSet")
	}

	if err := converter.SetEncodeCharSet(hconvert.Base10CharSet()); err == nil {
		t.Error("Failed bad object test for SetEncodeCharSet")
	}

	if s, err := converter.Convert("abc"); s != "" || err == nil {
		t.Error("Failed bad object test for Convert")
	}

	if s, err := converter.ConvertFrom(strings.NewReader("abc")); s != "" || err == nil {
		t.Error("Failed bad object test for ConvertFrom")
	}
}

func TestSetDecodeCharSet(t *testing.T) {
	t.Parallel()

	// Test that you can't set an empty set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.CharSet{}); err == nil {
		t.Error("Passed setting empty decoding character set")
	}

	// Test setting a decode char set when one is not already set.
	dec = hconvert.CharSet{}
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base16CharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base16CharSet()) {
		t.Error("Failed to set decoding character set when one is not already set")
	}

	// Test setting a decode char set when one is already set.
	dec = hconvert.Base2CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base64URLCharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base64URLCharSet()) {
		t.Error("Failed to set decoding character set when one is already set")
	}

	// Test setting a decode char set when both are set.
	dec = hconvert.Base10CharSet()
	enc = hconvert.Base4CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetDecodeCharSet(hconvert.Base10CharSet()); err != nil {
		t.Error(err)
	} else if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base10CharSet()) {
		t.Error("Failed to set decoding character set when both are set")
	}
}

func TestSetEncodeCharSet(t *testing.T) {
	t.Parallel()

	// Test that you can't set an empty set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.CharSet{}); err == nil {
		t.Error("Passed setting empty encoding character set")
	}

	// Test setting an encode char set when one is not already set.
	dec = hconvert.CharSet{}
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base58CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base58CharSet()) {
		t.Error("Failed to set encoding character set when one is not already set")
	}

	// Test setting an encode char set when one is already set.
	dec = hconvert.CharSet{}
	enc = hconvert.ASCIICharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base10CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base10CharSet()) {
		t.Error("Failed to set encoding character set when one is already set")
	}

	// Test setting an encode char set when both are set.
	dec = hconvert.Base4CharSet()
	enc = hconvert.Base10CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if err := converter.SetEncodeCharSet(hconvert.Base64CharSet()); err != nil {
		t.Error(err)
	} else if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base64CharSet()) {
		t.Error("Failed to set encoding character set when both are set")
	}
}

func TestDecodeCharSet(t *testing.T) {
	t.Parallel()

	// Test getting the decoding character set when nothing is set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve decoding character set when nothing is set")
	}

	// Test getting the decoding character set when only that is set.
	dec = hconvert.Base2CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base2CharSet()) {
		t.Error("Failed to retrieve decoding character set when only that is set")
	}

	// Test getting the decoding character set when only the encoding one is set.
	dec = hconvert.CharSet{}
	enc = hconvert.Base4CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve decoding character set when only the encoding one is set")
	}

	// Test getting the decoding character set when both are set.
	dec = hconvert.Base10CharSet()
	enc = hconvert.Base8CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if decSet := converter.DecodeCharSet(); !reflect.DeepEqual(decSet, hconvert.Base10CharSet()) {
		t.Error("Failed to retrieve decoding character set when both are set")
	}
}

func TestEncodeCharSet(t *testing.T) {
	t.Parallel()

	// Test getting the encoding character set when nothing is set.
	dec := hconvert.CharSet{}
	enc := hconvert.CharSet{}
	converter := hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve encoding character set when nothing is set")
	}

	// Test getting the encoding character set when only that is set.
	dec = hconvert.CharSet{}
	enc = hconvert.Base58CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base58CharSet()) {
		t.Error("Failed to retrieve encoding character set when only that is set")
	}

	// Test getting the encoding character set when only the encoding one is set.
	dec = hconvert.Base64CharSet()
	enc = hconvert.CharSet{}
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.CharSet{}) {
		t.Error("Failed to retrieve encoding character set when only the encoding one is set")
	}

	// Test getting the encoding character set when both are set.
	dec = hconvert.Base32CharSet()
	enc = hconvert.Base2CharSet()
	converter = hconvert.NewConverter(dec, enc)
	if encSet := converter.EncodeCharSet(); !reflect.DeepEqual(encSet, hconvert.Base2CharSet()) {
		t.Error("Failed to retrieve encoding character set when both are set")
	}
}

func TestConvert(t *testing.T) {
	// Test that we can properly convert the data using a converter with each of the standard
	// character sets.
	converter := hconvert.NewConverter(hconvert.ASCIICharSet(), hconvert.CharSet{})
	for _, test := range standardTests {
		converter.SetEncodeCharSet(test.charSet)
		s, err := converter.Convert(test.convertFrom)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.convertTo {
			t.Error(test.setName, "- Conversion failed")
			t.Log("Expected:", test.convertTo)
			t.Log("Received:", s)
		}
	}
}

func TestConvertFrom(t *testing.T) {
}
