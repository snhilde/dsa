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
	// These are the common tests for decoding/encoding using the standard character sets.
	data = []byte{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93,
		94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
		113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
		131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148,
		149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
		167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184,
		185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202,
		203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220,
		221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238,
		239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255,
	}
	standardTests = []testSet{
		{hconvert.ASCIICharSet(), "ASCII", string([]byte{4, 4, 3, 2, 1, 32, 96, 56, 32, 18, 10, 5, 67, 1, 80, 112, 60, 32, 17, 9, 4, 98, 65, 40, 88, 46, 24, 12, 70, 67, 49, 96, 116, 60, 31, 16, 8, 36, 34, 25, 16, 74, 38, 19, 74, 5, 18, 81, 44, 88, 45, 23, 11, 102, 3, 9, 72, 102, 52, 26, 77, 70, 115, 65, 100, 116, 59, 30, 15, 39, 99, 122, 1, 2, 66, 33, 81, 8, 84, 50, 29, 16, 73, 37, 18, 105, 68, 106, 57, 30, 80, 40, 84, 74, 53, 34, 85, 44, 87, 44, 22, 43, 37, 90, 113, 58, 94, 47, 88, 12, 22, 19, 13, 72, 101, 51, 25, 109, 6, 75, 41, 86, 108, 54, 91, 77, 119, 3, 69, 100, 115, 58, 29, 46, 103, 59, 97, 114, 122, 61, 95, 15, 87, 115, 126, 1, 1, 65, 32, 112, 72, 44, 26, 15, 8, 68, 98, 81, 56, 100, 54, 29, 15, 72, 36, 50, 41, 28, 82, 43, 22, 75, 102, 19, 25, 84, 110, 57, 29, 79, 39, 116, 10, 13, 10, 71, 36, 82, 105, 84, 122, 69, 38, 85, 43, 86, 43, 53, 106, 125, 66, 99, 50, 89, 109, 22, 91, 53, 94, 113, 57, 93, 46, 119, 75, 109, 122, 127, 64, 96, 112, 88, 60, 38, 23, 13, 71, 100, 50, 57, 44, 94, 51, 27, 78, 103, 116, 26, 29, 22, 79, 41, 85, 107, 53, 123, 13, 78, 107, 55, 92, 110, 119, 91, 126, 7, 7, 69, 99, 114, 57, 60, 110, 63, 35, 83, 106, 117, 123, 29, 94, 119, 63, 97, 113, 121, 60, 126, 79, 47, 91, 111, 120, 124, 126, 95, 63, 103, 119, 125, 127}), data},
		{hconvert.Base2CharSet(), "Base2", "10000001000000011000001000000010100000110000001110000100000001001000010100000101100001100000011010000111000001111000100000001000100010010000100110001010000010101000101100001011100011000000110010001101000011011000111000001110100011110000111110010000000100001001000100010001100100100001001010010011000100111001010000010100100101010001010110010110000101101001011100010111100110000001100010011001000110011001101000011010100110110001101110011100000111001001110100011101100111100001111010011111000111111010000000100000101000010010000110100010001000101010001100100011101001000010010010100101001001011010011000100110101001110010011110101000001010001010100100101001101010100010101010101011001010111010110000101100101011010010110110101110001011101010111100101111101100000011000010110001001100011011001000110010101100110011001110110100001101001011010100110101101101100011011010110111001101111011100000111000101110010011100110111010001110101011101100111011101111000011110010111101001111011011111000111110101111110011111111000000010000001100000101000001110000100100001011000011010000111100010001000100110001010100010111000110010001101100011101000111110010000100100011001001010010011100101001001010110010110100101111001100010011001100110101001101110011100100111011001111010011111101000001010000110100010101000111010010010100101101001101010011110101000101010011010101010101011101011001010110110101110101011111011000010110001101100101011001110110100101101011011011010110111101110001011100110111010101110111011110010111101101111101011111111000000110000011100001011000011110001001100010111000110110001111100100011001001110010101100101111001100110011011100111011001111110100001101000111010010110100111101010011010101110101101101011111011000110110011101101011011011110111001101110111011110110111111110000011100001111000101110001111100100111001011110011011100111111010001110100111101010111010111110110011101101111011101110111111110000111100011111001011110011111101001111010111110110111101111111100011111001111110101111101111111100111111011111111011111111", data},
		{hconvert.Base4CharSet(), "Base4", "100020003001000110012001300200021002200230030003100320033010001010102010301100111011201130120012101220123013001310132013302000201020202030210021102120213022002210222022302300231023202330300030103020303031003110312031303200321032203230330033103320333100010011002100310101011101210131020102110221023103010311032103311001101110211031110111111121113112011211122112311301131113211331200120112021203121012111212121312201221122212231230123112321233130013011302130313101311131213131320132113221323133013311332133320002001200220032010201120122013202020212022202320302031203220332100210121022103211021112112211321202121212221232130213121322133220022012202220322102211221222132220222122222223223022312232223323002301230223032310231123122313232023212322232323302331233223333000300130023003301030113012301330203021302230233030303130323033310031013102310331103111311231133120312131223123313031313132313332003201320232033210321132123213322032213222322332303231323232333300330133023303331033113312331333203321332233233330333133323333", data},
		{hconvert.Base8CharSet(), "Base8", "201003010024060160401102405414032070170401042204612025054134300621503307016436076200411042144411223047120244521262605513427460142310631503246615634071164354741723707720040502206421052144351022245113230465162365012124451524252531272605453226656135274575403026114331062546316641513246555433267157340705623467216535473570362751733707657637700201405016044130320742104612427062154350762044311223451126264571423146515634473172375012064250722245515236505232525353126656537302615453166455333267561346725673627557537700603413036114270661744311625457146334731764150722647523256555373066355333671567366776034170561744713633477216475272766355735677607437136375172766757743717657577477377377", data},
		{hconvert.Base10CharSet(), "Base10", "496993557421161204163243009437446396931089621987166228867825381430598310478552724232175338893085943309322844458622037977610450736647303669282401420783518387091593654245305591568630492124408336898055480790220185594414336311305186561504622750303362867675227156509593310924837768420470062416109994603982012849241281559241343212928642441931019907648691672976952280785510275422238003732540719940795979897949667308357456420122443019724695903918518199221894619196205878219563639903280587615919371037707517469737463169252698161246382118698239117850296427010987117809205722959373004300581599055823946915702512166260047615", data},
		{hconvert.Base16CharSet(), "Base16", "102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F202122232425262728292A2B2C2D2E2F303132333435363738393A3B3C3D3E3F404142434445464748494A4B4C4D4E4F505152535455565758595A5B5C5D5E5F606162636465666768696A6B6C6D6E6F707172737475767778797A7B7C7D7E7F808182838485868788898A8B8C8D8E8F909192939495969798999A9B9C9D9E9FA0A1A2A3A4A5A6A7A8A9AAABACADAEAFB0B1B2B3B4B5B6B7B8B9BABBBCBDBEBFC0C1C2C3C4C5C6C7C8C9CACBCCCDCECFD0D1D2D3D4D5D6D7D8D9DADBDCDDDEDFE0E1E2E3E4E5E6E7E8E9EAEBECEDEEEFF0F1F2F3F4F5F6F7F8F9FAFBFCFDFEFF", data},
		{hconvert.Base32CharSet(), "Base32", "EBAGBAFAYDQQCIKBMGA2DQPCAIREEYUCULBOGAZDINRYHI6D4QCCIRDEQSSMJZIFEVCWLBNFYXTAMJSGM2DKNRXHA4TUOZ4HU7D6QCBIJBUIRKGI5EESSSLJRGU4T2QKFJFGVCVKZLVQWK2LNOF2XS7MBQWEY3EMVTGO2DJNJVWY3LON5YHC4TTOR2XM53YPF5HW7D5PZ7YBAMCQOCILBUHRCEYVC4MRWHI7EERSKJZJFMWS6MJTGU3TSOZ5H5AUGRKHJFFU2T2RKNKVOWK3LVPWCY3FM5UWW3LPOFZXK53ZPN6X7AMDQWDYTC4NR6IZHFMXTGNZ3H5BUOS2PKNLVWX3DM5VW643XPN7YHB4LR6JZPG47UOT2XL5TW6537Q6HZPH5HV6337R6P27P6P37X7", data},
		{hconvert.Base36CharSet(), "Base36", "168SWOI6IUZJ4FBWKNLNH695ZL88V65QCFGNWRWEPQCXB9DYSMLUOWQAHVT3R9GSC1V47SSXDIVJDA3NTTL6R044PZZ7ZWHTGU2MKOW5TS28X2MBWENH3WFZ4S1SARSPFHLRAKVQRGPMZB66SGTZ2LZBOTL7R28WCQ8925C747B44L60VRK3SCRIN4ZVNWN7PDSUKGO6LGJHU1NUWJ7YT1H9UJPE3OS17ONSK7SP4YSMYTU568DO2TQETWNRMBXB2DTD8KQORCOAKAIZLM9SVR8AXE1ACXFURSZ11NUBRHIGHFD64YHMP99UCVZR944N8CO01O4X64CMBD8BE0HQBM2ZY5UWE4UPLC4SA50XAJEL4BKKXB1KH21PISNA37EQWPBPQ11YPR", data},
		{hconvert.Base58CharSet(), "Base58", "cWB5HCBdLjAuqGGReWE3R3CguuwSjw6RHn39s2yuDRTS5NsBgNiFpWgAnEx6VQi8csexkgYw3mdYrMHr8x9i7aEwP8kZ7vccXWqKDvGv3u1GxFKPuAkn8JCPPGDMf3vMMnbzm6Nh9zh1gcNsMvH3ZNLmP5fSG6DGbbi2tuwMWPthr4boWwCxf7ewSgNQeacyozhKDDQQ1qL5fQFUW52QKUZDZ5fw3KXNQJMcNTcaB723LchjeKun7MuGW5qyCBZYzA1KjofN1gYBV3NqyhQJ3Ns746GNuf9N2pQPmHz4xpnSrrfCvy6TVVz5d4PdrjeshsWQwpZsZGzvbdAdN8MKV5QsBDY", data},
		{hconvert.Base62CharSet(), "Base62", "DF6jwBonlJbe2WufDkWP7RUL88ITCkWFU1BY71YArwYLlvIW0Vebv345Mcne8n331AuNRFVT841EkUZ46p270jAKVpEF9a4g3oRHSPF8m0bOFwi5F2Ds6wOGW8UeRnrK0RK3oRMcqwhuZlr78nv4eTZ2jnrwmITpz7CKuPKdPVSiBBmIEOseIG2v8sOHkkoL9aOgB1QRC8tlDb8MYP1NCYvBMKKwccRjJfpLwTl0ogo5IK6nGa10PIxNzTaHSUU7V0IBlbZa97Nn8F2olfSRYMKf7vPAnxzLIyXMC10CrHjACZGHJnQUcvfIlgnBDw7oQtlfqHVFx7V1jU0b9nxCH9", data},
		{hconvert.Base64CharSet(), "Base64", "QIDBAUGBwgJCgsMDQ4PEBESExQVFhcYGRobHB0eHyAhIiMkJSYnKCkqKywtLi8wMTIzNDU2Nzg5Ojs8PT4/QEFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaW1xdXl9gYWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXp7fH1+f4CBgoOEhYaHiImKi4yNjo+QkZKTlJWWl5iZmpucnZ6foKGio6SlpqeoqaqrrK2ur7CxsrO0tba3uLm6u7y9vr/AwcLDxMXGx8jJysvMzc7P0NHS09TV1tfY2drb3N3e3+Dh4uPk5ebn6Onq6+zt7u/w8fLz9PX29/j5+vv8/f7/", data},
		{hconvert.Base64URLCharSet(), "Base64URL", "QIDBAUGBwgJCgsMDQ4PEBESExQVFhcYGRobHB0eHyAhIiMkJSYnKCkqKywtLi8wMTIzNDU2Nzg5Ojs8PT4_QEFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaW1xdXl9gYWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXp7fH1-f4CBgoOEhYaHiImKi4yNjo-QkZKTlJWWl5iZmpucnZ6foKGio6SlpqeoqaqrrK2ur7CxsrO0tba3uLm6u7y9vr_AwcLDxMXGx8jJysvMzc7P0NHS09TV1tfY2drb3N3e3-Dh4uPk5ebn6Onq6-zt7u_w8fLz9PX29_j5-vv8_f7_", data},
		{hconvert.ASCII85CharSet(), "ASCII85", "\"/uDIm[G=mT>d.Q;:lQGQE(@F\\Os)Hgb/sHN*lt;Dm;M&DLJ88B>5o_8)]G]41O6BeAVJk<3\"o8L,@pP'KpO!b'$@>sh.@sSdVraq3>bYdlJ317Y&jX?88-R]s>?;5^h]#NA&ht6*Fd9SgS\\4@pHT]_0b3ar_oR&DFaW75-`k%K,o]5m>Bcf(kc50!m=T>L^9B*o:9R7`$tg[[K&A+]NkiDmA)tf(!h&\\dtXq`pP?'flTeYmC`W`%XRL%8T.K?AJ2%#$fXtpSp-?.j96^54GkN_k4!)RLq=/&U5PAE*\\kr9q_Z?)\\i>gHK=,8OINM!", data},
		{hconvert.Z85CharSet(), "Z85", "1e#zE)WCs)Pt^dMqp(MCMA7vBXK%8D*+e%DJ9($qz)qI5zHFnnxtk].n8YCYjgKlx!wRF>ri1]nHbv{L6G{K0+63vt%?dv%O^R@:}it+U^(FigmU5<TunncNY%tuqkZ?Y2Jw5?$l9B^oO*OXjv{DPY.f+i:@.]N5zB:Smkc->4Gb]Yk)tx=/7>=kf0)sPtHZox9]poNm-3$*WWG5waYJ>&z)w8$/70?5X^$T}-{Lu6/(P!U)y-S-4TNH4nPdGuwFh423/T${O{cud<olZkjC>J.>j08NH}se5QkLwA9X>@o}.Vu8X&t*DGsbnKEJI0", data},
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
