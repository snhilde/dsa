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

	// String to decode.
	decodeMe string

	// Data to encode.
	encodeMe []byte
}

const (
	// This character is not in any of the standard character sets. Decoding with this should fail.
	invalidChar = "∞"
)

var (
	// We will use this data for the standard conversion tests.
	conversionData = []byte{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
		49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
		72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
		95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
		114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
		132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149,
		150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167,
		168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185,
		186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203,
		204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221,
		222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239,
		240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 0,
	}

	// These are the common tests for decoding/encoding using the standard character sets.
	standardTests = []testSet{
		{hconvert.ASCIICharSet(), "ASCII", string([]byte{8, 8, 6, 4, 2, 65, 64, 112, 64, 36, 20, 11, 6, 3, 33, 96, 120, 64, 34, 18, 9, 69, 2, 81, 48, 92, 48, 25, 13, 6, 99, 65, 104, 120, 62, 32, 16, 72, 68, 50, 33, 20, 76, 39, 20, 10, 37, 34, 89, 48, 90, 46, 23, 76, 6, 19, 17, 76, 104, 53, 27, 13, 103, 3, 73, 104, 118, 60, 30, 79, 71, 116, 2, 5, 4, 67, 34, 17, 40, 100, 58, 33, 18, 74, 37, 83, 9, 84, 114, 61, 32, 81, 41, 20, 106, 69, 42, 89, 46, 88, 44, 86, 75, 53, 98, 117, 60, 95, 48, 24, 44, 38, 27, 17, 74, 102, 51, 90, 13, 22, 83, 45, 88, 109, 55, 27, 110, 7, 11, 73, 102, 116, 58, 93, 78, 119, 67, 101, 116, 123, 62, 31, 47, 103, 124, 2, 3, 2, 65, 97, 16, 88, 52, 30, 17, 9, 69, 34, 113, 72, 108, 58, 31, 16, 72, 100, 82, 57, 36, 86, 45, 23, 76, 38, 51, 41, 92, 114, 59, 30, 79, 104, 20, 26, 21, 14, 73, 37, 83, 41, 117, 10, 77, 42, 87, 44, 86, 107, 85, 123, 5, 70, 101, 51, 90, 45, 54, 107, 61, 98, 115, 58, 93, 111, 23, 91, 117, 127, 1, 65, 97, 48, 120, 76, 46, 27, 15, 72, 100, 114, 89, 60, 102, 55, 29, 79, 104, 52, 58, 45, 30, 83, 43, 86, 107, 118, 27, 29, 86, 111, 57, 93, 111, 55, 124, 14, 15, 11, 71, 100, 114, 121, 92, 126, 71, 39, 85, 107, 118, 59, 61, 110, 127, 67, 99, 114, 121, 125, 30, 95, 55, 95, 113, 121, 125, 62, 127, 79, 111, 123, 126, 0}), conversionData},
		{hconvert.Base2CharSet(), "Base2", "1000000100000001100000100000001010000011000000111000010000000100100001010000010110000110000001101000011100000111100010000000100010001001000010011000101000001010100010110000101110001100000011001000110100001101100011100000111010001111000011111001000000010000100100010001000110010010000100101001001100010011100101000001010010010101000101011001011000010110100101110001011110011000000110001001100100011001100110100001101010011011000110111001110000011100100111010001110110011110000111101001111100011111101000000010000010100001001000011010001000100010101000110010001110100100001001001010010100100101101001100010011010100111001001111010100000101000101010010010100110101010001010101010101100101011101011000010110010101101001011011010111000101110101011110010111110110000001100001011000100110001101100100011001010110011001100111011010000110100101101010011010110110110001101101011011100110111101110000011100010111001001110011011101000111010101110110011101110111100001111001011110100111101101111100011111010111111001111111100000001000000110000010100000111000010010000101100001101000011110001000100010011000101010001011100011001000110110001110100011111001000010010001100100101001001110010100100101011001011010010111100110001001100110011010100110111001110010011101100111101001111110100000101000011010001010100011101001001010010110100110101001111010100010101001101010101010101110101100101011011010111010101111101100001011000110110010101100111011010010110101101101101011011110111000101110011011101010111011101111001011110110111110101111111100000011000001110000101100001111000100110001011100011011000111110010001100100111001010110010111100110011001101110011101100111111010000110100011101001011010011110101001101010111010110110101111101100011011001110110101101101111011100110111011101111011011111111000001110000111100010111000111110010011100101111001101110011111101000111010011110101011101011111011001110110111101110111011111111000011110001111100101111001111110100111101011111011011110111111110001111100111111010111110111111110011111101111111101111111100000000", conversionData},
		{hconvert.Base4CharSet(), "Base4", "1000200030010001100120013002000210022002300300031003200330100010101020103011001110112011301200121012201230130013101320133020002010202020302100211021202130220022102220223023002310232023303000301030203030310031103120313032003210322032303300331033203331000100110021003101010111012101310201021102210231030103110321033110011011102110311101111111211131120112111221123113011311132113312001201120212031210121112121213122012211222122312301231123212331300130113021303131013111312131313201321132213231330133113321333200020012002200320102011201220132020202120222023203020312032203321002101210221032110211121122113212021212122212321302131213221332200220122022203221022112212221322202221222222232230223122322233230023012302230323102311231223132320232123222323233023312332233330003001300230033010301130123013302030213022302330303031303230333100310131023103311031113112311331203121312231233130313131323133320032013202320332103211321232133220322132223223323032313232323333003301330233033310331133123313332033213322332333303331333233330000", conversionData},
		{hconvert.Base8CharSet(), "Base8", "100401404012030070200441202606015034074200421102305012426056140310641543407217037100204421062204511423450122250531302645613630061144314641523306716034472166360751743750020241103210425062164411122445514232471172405052224652125254535302625513327056536276601413046154431263147320645523266615533467560342711633507256635674171364755743727717740100602407022054150361042305213431066164371022144511624453132274611463246716235475176405032124351122646517242515252565453327257541306625473226555533670563352735713667657740301605417046134330762144712627463156354772064351323651527266575433166555734673573377016074270762345715637507236535373166756737703617457176475373367761747727677637577577400", conversionData},
		{hconvert.Base10CharSet(), "Base10", "127230350699817268265790210415986277614358943228714554590163297646233167482509497403436886756630001487186648181407241722268275388581709739336294763720580707095447975486798231441569405983848534245902203082296367512170070095694127759745183424077660894124858152066455887596758468715640335978524158618619395289405768079165783862509732465134341096358065068282099783881090630508092928955530424304843770853875114830939508843551345413049522151403140659000805022514228704824208291815239830429675358985653124472252790571328690729279073822386749214169675885314812702159156665077599489100948889358290930410419843114562572189440", conversionData},
		{hconvert.Base16CharSet(), "Base16", "102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F202122232425262728292A2B2C2D2E2F303132333435363738393A3B3C3D3E3F404142434445464748494A4B4C4D4E4F505152535455565758595A5B5C5D5E5F606162636465666768696A6B6C6D6E6F707172737475767778797A7B7C7D7E7F808182838485868788898A8B8C8D8E8F909192939495969798999A9B9C9D9E9FA0A1A2A3A4A5A6A7A8A9AAABACADAEAFB0B1B2B3B4B5B6B7B8B9BABBBCBDBEBFC0C1C2C3C4C5C6C7C8C9CACBCCCDCECFD0D1D2D3D4D5D6D7D8D9DADBDCDDDEDFE0E1E2E3E4E5E6E7E8E9EAEBECEDEEEFF0F1F2F3F4F5F6F7F8F9FAFBFCFDFEFF00", conversionData},
		{hconvert.Base32CharSet(), "Base32", "BAIBQIBIGA4EASCQLBQGQ4DYQCEJBGFAVCYLRQGI2DMOB2HQ7EAQSEIZEEUTCOKBJFIVSYLJOF4YDCMRTGQ2TMNZYHE5DWPB5HY7UAQKCINCEKRSHJBEUUS2MJVHE6UCRKJJVIVKWK5MFSWS3LROV4X3AMFRGGZDFMZTWQ2LKNNWG23TPOBYXE43UOV3HO6DZPJ5XY7L6P6AIDAUDQSCYNB4IRGFIXDENR2HZBEMSSOKJLFUXTCMZVG44TWPJ7IFBUKR2JJNGU6UKTKVLVSW25L5QWGZLHNFVW233RON2XO6L3PV7YDA4FQ6EYXDMPSGJZLF4ZTOOZ7INDUWT2TK5NV6Y3HNNXXG533P6BYPC4PSOLZXH5DU6V27M5XXO74HR6LZ7J5PW674PT6X37T6757YA", conversionData},
		{hconvert.Base36CharSet(), "Base36", "8CELKE9AE4CNZH0NMVTYY4H6L2YN1NSRCDYI12EGMZFWW2RBFKRE94QYN66YPVBDHP99ZGTPC689QFU1G2EO00TDKVUGVB2PJPUOJ50PFRJZF6MSMG6XNQPLTZGP8LO4U56Q38HQBAUBF3G0AFP5EKB74IEV4FZA2IIOF9YQLVZ1CMIU9VYYXMROKJJ56052SI55DIJYWLMMTNTNRCONIIIU184KA87CMND73RG2RBFNBO6CSBL783FTGO90ESSUOY8U4ZHS2I4R4AV15QDPDTN1PFT5NXSQTQ13FTJNNGJ97X9NN99DFLXZVNI9STD17E2OBVMZVIXQOU34ZNI36L1BMXNQCIIDZQA143QKQXZQ6QACSUJ5LAK5HNPJQSOVCK3AX3HYV40", conversionData},
		{hconvert.Base58CharSet(), "Base58", "3hhmTtUMuQ58kC2P5nXzZf6ibAYfJcVkvD3nW7qk9TpwfgKWW8nq6QhytAgaBEPeYdLaZ5oZwZDzKapZMt6RSU1iX2hDEmZeBoggFS29F9nYDQrCC5G4EBCuFoG9YAwtUrkHTwcikEhdc1q9XHL9pHsCEtTaeVc19qiC7N8euWk3DHYtcK8FnfENTJLmVNsPokRewNukGagKDaVctPJmBvYdCQbcqXEm3Tqby7iU1RbZbKaaoDkTx75FhRNP92McgBGkRhvxKk1WdGmVRBFKpDVW9ieEshRMuq2EFUVxtVPub94BdfCb5vnvwQ9xugXGAoPvpgxx69TcfZpdVWFSgavULU6Kq", conversionData},
		{hconvert.Base62CharSet(), "Base62", "Mwh1sSzxfY9hed6AOyH32vhXduSSmqHMCG5xFKbI4vx154miUItspw3CLcPF4Yuym3EyyiECfgvI6ATBIQ67HCgqtbi4tpHZzwSqHUUqUYgaL2gLuPNXkrIOS1GlCz4Qv049wTBccobVntrhoYNzPKK1VJ6wtcVOufu8W6nOvV6kcOpbTevqZeaKGqmTVFT7fmT6Pp5MYSwRAPejKJdrz4OI6OAgFewhJQuEne8RZU073vMFci6TsgTPF2L0Hp8hELPImlLBdygVCIRpo1Ndx4PZ8vD8oVtUAX9vt0S3IDOgJ5nPhwHaewIjg6tGXoKSRGS2x4TF8RERL4EjkgScw72", conversionData},
		{hconvert.Base64CharSet(), "Base64", "BAgMEBQYHCAkKCwwNDg8QERITFBUWFxgZGhscHR4fICEiIyQlJicoKSorLC0uLzAxMjM0NTY3ODk6Ozw9Pj9AQUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVpbXF1eX2BhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ent8fX5/gIGCg4SFhoeIiYqLjI2Oj5CRkpOUlZaXmJmam5ydnp+goaKjpKWmp6ipqqusra6vsLGys7S1tre4ubq7vL2+v8DBwsPExcbHyMnKy8zNzs/Q0dLT1NXW19jZ2tvc3d7f4OHi4+Tl5ufo6err7O3u7/Dx8vP09fb3+Pn6+/z9/v8A", conversionData},
		{hconvert.Base64URLCharSet(), "Base64URL", "BAgMEBQYHCAkKCwwNDg8QERITFBUWFxgZGhscHR4fICEiIyQlJicoKSorLC0uLzAxMjM0NTY3ODk6Ozw9Pj9AQUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVpbXF1eX2BhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ent8fX5_gIGCg4SFhoeIiYqLjI2Oj5CRkpOUlZaXmJmam5ydnp-goaKjpKWmp6ipqqusra6vsLGys7S1tre4ubq7vL2-v8DBwsPExcbHyMnKy8zNzs_Q0dLT1NXW19jZ2tvc3d7f4OHi4-Tl5ufo6err7O3u7_Dx8vP09fb3-Pn6-_z9_v8A", conversionData},
		{hconvert.ASCII85CharSet(), "ASCII85", "$O.5j0s$I$]X^7jK3sSp.h[1\\O?G7Ks,:(@%kdgnR*gj];qs;)GE(#+OR5'RAeg:Fer8**gr8c`dn70ScS;PQ:t1-E8J;99at\\M6SS7Yc&H>+cu5jYf^0(]m_V!D9$E6rdVYRTc^RFfY((U[A?0=9`n9JELX0LXb;`^ZDs[C@#Nno%FppNa4(_3Jb2\\n-WlZRG_kf/#AJl!Ima/[1a7=0F++#\\&E&*M$0)`rL2PTUSOJ^#f@+XZjn\"C)Xl('XQKh(>+.KdUd_B@2h`^%DEoT(Ja?O4:i,>pg?n=n\\Xb4>`h-.a_Y36mh3rJ^rnt((M!", conversionData},
		{hconvert.Z85CharSet(), "Z85", "3Kdk<f%3E3YTZm<Gi%O{d?WgXKuCmG%bp7v4>^*[N9*<Yq}%q8CA72aKNk6Nw!*pB!@n99*@n=-^[mfO=OqLMp$gcAnFqoo:$XIlOOmU=5Dta=#k<U/Zf7Y).R0zo3Al@^RUNP=ZNB/U77QWwufso-[oFAHTfHT+q-ZVz%Wyv2J[]4B{{J:j7.iF+hX[cS(VNC.>/e2wF(0E):eWg:msfBaa2X5A59I3f8-@HhLPQOKFZ2/vaTV<[1y8T(76TMG?7tadG^Q^.xvh?-Z4zA]P7F:uKjp&bt{*u[s[XT+jt-?cd:.Uil)?i@FZ@[$77I0", conversionData},
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

	if b, err := converter.Decode("abc"); b != nil || err == nil {
		t.Error("Failed bad object test for Decode")
	}

	if b, err := converter.DecodeFrom(strings.NewReader("1234")); b != nil || err == nil {
		t.Error("Failed bad object test for DecodeFrom")
	}

	if s, err := converter.Encode([]byte{0x01, 0x02}); s != "" || err == nil {
		t.Error("Failed bad object test for Encode")
	}

	if err := converter.EncodeTo([]byte{0x03, 0x04}, new(strings.Builder)); err == nil {
		t.Error("Failed bad object test for EncodeTo")
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

func TestDecode(t *testing.T) {
	t.Parallel()

	// Test that we can properly decode the data using a converter with each of the standard
	// character sets.
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		b, err := converter.Decode(test.decodeMe)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- Decode failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		if _, err := converter.Decode(invalidChar); err == nil {
			t.Error(test.setName, "- Decode invalid character passed")
		}
	}
}

func TestDecodeFrom(t *testing.T) {
	t.Parallel()

	// Test that we can properly read and decode data from the provided io.Reader using a converter
	// with each of the standard character sets
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		reader := strings.NewReader(test.decodeMe)
		b, err := converter.DecodeFrom(reader)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- DecodeFrom failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		converter.SetDecodeCharSet(test.charSet)
		reader := strings.NewReader(invalidChar)
		if _, err := converter.DecodeFrom(reader); err == nil {
			t.Error(test.setName, "- DecodeFrom invalid character passed")
		}
	}
}

func TestDecodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly decode the data using the standard character sets.
	for _, test := range standardTests {
		b, err := hconvert.DecodeWith(test.decodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if !reflect.DeepEqual(b, test.encodeMe) {
			t.Error(test.setName, "- DecodeWith failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", b)
		}
	}

	// Test that we can't decode characters not in the character set.
	for _, test := range standardTests {
		if _, err := hconvert.DecodeWith(invalidChar, test.charSet); err == nil {
			t.Error(test.setName, "- DecodeWith invalid character passed")
		}
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the data using a converter with each of the standard
	// character sets.
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetEncodeCharSet(test.charSet)
		s, err := converter.Encode(test.encodeMe)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.decodeMe {
			t.Error(test.setName, "- Encode failed")
			t.Log("Expected:", test.decodeMe)
			t.Log("Received:", s)
		}
	}
}

func TestEncodeTo(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode data and write it to the provided io.Writer using a
	// converter with each of the standard character sets
	converter := new(hconvert.Converter)
	for _, test := range standardTests {
		converter.SetEncodeCharSet(test.charSet)
		writer := new(strings.Builder)
		if err := converter.EncodeTo(test.encodeMe, writer); err != nil {
			t.Error(test.setName, "-", err)
		}

		if writer.String() != test.decodeMe {
			t.Error(test.setName, "- EncodeTo failed")
			t.Log("Expected:", test.encodeMe)
			t.Log("Received:", writer.String())
		}
	}
}

func TestEncodeWith(t *testing.T) {
	t.Parallel()

	// Test that we can properly encode the data using the standard character sets.
	for _, test := range standardTests {
		s, err := hconvert.EncodeWith(test.encodeMe, test.charSet)
		if err != nil {
			t.Error(test.setName, "-", err)
		} else if s != test.decodeMe {
			t.Error(test.setName, "- EncodeWith failed")
			t.Log("Expected:", test.decodeMe)
			t.Log("Received:", s)
		}
	}
}
