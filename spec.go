package iso8583

// FieldDescription shows the content required to describe an iso8583 field
type FieldDescription struct {
	ContentType string
	MaxLen      int64
	MinLen      int64
	LenType     string
	Label       string
}

// GetSpecISO8583 returns field specifications for iso8583
func GetSpecISO8583() []FieldDescription {
	iso8583 := make([]FieldDescription, 129)

	iso8583[0] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Message Type Indicator"}
	iso8583[1] = FieldDescription{ContentType: "b", MaxLen: 8, LenType: "fixed", Label: "Bitmap"}
	iso8583[2] = FieldDescription{ContentType: "n", MinLen: 12, MaxLen: 19, LenType: "llvar", Label: "Primary account number (PAN)"}
	iso8583[3] = FieldDescription{ContentType: "n", MaxLen: 6, LenType: "fixed", Label: "Processing code"}
	iso8583[4] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Amount, transaction"}
	iso8583[5] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Amount, settlement"}
	iso8583[6] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Amount, cardholder billing"}
	iso8583[7] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Transmission date & time"}
	iso8583[8] = FieldDescription{ContentType: "n", MaxLen: 8, LenType: "fixed", Label: "Amount, cardholder billing fee"}
	iso8583[9] = FieldDescription{ContentType: "n", MaxLen: 8, LenType: "fixed", Label: "Conversion rate, settlement"}
	iso8583[10] = FieldDescription{ContentType: "n", MaxLen: 8, LenType: "fixed", Label: "Conversion rate, cardholder billing"}
	iso8583[11] = FieldDescription{ContentType: "n", MaxLen: 6, LenType: "fixed", Label: "System trace audit number"}
	iso8583[12] = FieldDescription{ContentType: "n", MaxLen: 6, LenType: "fixed", Label: "Time, local transaction (hhmmss)"}
	iso8583[13] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Date, local transaction (MMDD)"}
	iso8583[14] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Date, expiration"}
	iso8583[15] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Date, settlement"}
	iso8583[16] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Date, conversion"}
	iso8583[17] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Date, capture"}
	iso8583[18] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Merchant type"}
	iso8583[19] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Acquiring institution country code"}
	iso8583[20] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "PAN extended, country code"}
	iso8583[21] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Forwarding institution. country code"}
	iso8583[22] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Point of service entry mode"}
	iso8583[23] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Application PAN sequence number"}
	iso8583[24] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Network International identifier (NII)"}
	iso8583[25] = FieldDescription{ContentType: "n", MaxLen: 2, LenType: "fixed", Label: "Point of service condition code"}
	iso8583[26] = FieldDescription{ContentType: "n", MaxLen: 2, LenType: "fixed", Label: "Point of service capture code"}
	iso8583[27] = FieldDescription{ContentType: "n", MaxLen: 1, LenType: "fixed", Label: "Authorizing identification response length"}
	iso8583[28] = FieldDescription{ContentType: "an", MaxLen: 9, LenType: "fixed", Label: "Amount, transaction fee"}
	iso8583[29] = FieldDescription{ContentType: "an", MaxLen: 9, LenType: "fixed", Label: "Amount, settlement fee"}
	iso8583[30] = FieldDescription{ContentType: "an", MaxLen: 9, LenType: "fixed", Label: "Amount, transaction processing fee"}
	iso8583[31] = FieldDescription{ContentType: "an", MaxLen: 9, LenType: "fixed", Label: "Amount, settlement processing fee"}
	iso8583[32] = FieldDescription{ContentType: "n", MaxLen: 11, LenType: "llvar", Label: "Acquiring institution identification code"}
	iso8583[33] = FieldDescription{ContentType: "n", MaxLen: 11, LenType: "llvar", Label: "Forwarding institution identification code"}
	iso8583[34] = FieldDescription{ContentType: "ns", MaxLen: 28, LenType: "llvar", Label: "Primary account number, extended"}
	iso8583[35] = FieldDescription{ContentType: "z", MaxLen: 37, LenType: "llvar", Label: "Track 2 data"}
	iso8583[36] = FieldDescription{ContentType: "n", MaxLen: 104, LenType: "lllvar", Label: "Track 3 data"}
	iso8583[37] = FieldDescription{ContentType: "an", MaxLen: 12, LenType: "fixed", Label: "Retrieval reference number"}
	iso8583[38] = FieldDescription{ContentType: "an", MaxLen: 6, LenType: "fixed", Label: "Authorization identification response"}
	iso8583[39] = FieldDescription{ContentType: "an", MaxLen: 2, LenType: "fixed", Label: "Response code"}
	iso8583[40] = FieldDescription{ContentType: "an", MaxLen: 3, LenType: "fixed", Label: "Service restriction code"}
	iso8583[41] = FieldDescription{ContentType: "ans", MaxLen: 8, LenType: "fixed", Label: "Card acceptor terminal identification"}
	iso8583[42] = FieldDescription{ContentType: "ans", MaxLen: 15, LenType: "fixed", Label: "Card acceptor identification code"}
	iso8583[43] = FieldDescription{ContentType: "ans", MaxLen: 40, LenType: "fixed", Label: "Card acceptor name/location"}
	iso8583[44] = FieldDescription{ContentType: "an", MaxLen: 25, LenType: "llvar", Label: "Additional response data"}
	iso8583[45] = FieldDescription{ContentType: "an", MaxLen: 76, LenType: "llvar", Label: "Track 1 data"}
	iso8583[46] = FieldDescription{ContentType: "an", MaxLen: 999, LenType: "lllvar", Label: "Additional data - ISO"}
	iso8583[47] = FieldDescription{ContentType: "an", MaxLen: 999, LenType: "lllvar", Label: "Additional data - national"}
	iso8583[48] = FieldDescription{ContentType: "an", MaxLen: 999, LenType: "lllvar", Label: "Additional data - private"}
	iso8583[49] = FieldDescription{ContentType: "an", MaxLen: 3, LenType: "fixed", Label: "Currency code, transaction"}
	iso8583[50] = FieldDescription{ContentType: "an", MaxLen: 3, LenType: "fixed", Label: "Currency code, settlement"}
	iso8583[51] = FieldDescription{ContentType: "an", MaxLen: 3, LenType: "fixed", Label: "Currency code, cardholder billing"}
	iso8583[52] = FieldDescription{ContentType: "b", MaxLen: 8, LenType: "fixed", Label: "Personal identification number data"}
	iso8583[53] = FieldDescription{ContentType: "n", MaxLen: 16, LenType: "fixed", Label: "Security related control information"}
	iso8583[54] = FieldDescription{ContentType: "an", MaxLen: 120, LenType: "lllvar", Label: "Additional amounts"}
	iso8583[55] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved ISO"}
	iso8583[56] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved ISO"}
	iso8583[57] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved national"}
	iso8583[58] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved national"}
	iso8583[59] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved national"}
	iso8583[60] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved national"}
	iso8583[61] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved private"}
	iso8583[62] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved private"}
	iso8583[63] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved private"}
	iso8583[64] = FieldDescription{ContentType: "b", MaxLen: 8, LenType: "fixed", Label: "Message authentication code (MAC)"}
	iso8583[65] = FieldDescription{ContentType: "b", MaxLen: 1, LenType: "fixed", Label: "Bitmap, extended"}
	iso8583[66] = FieldDescription{ContentType: "n", MaxLen: 1, LenType: "fixed", Label: "Settlement code"}
	iso8583[67] = FieldDescription{ContentType: "n", MaxLen: 2, LenType: "fixed", Label: "Extended payment code"}
	iso8583[68] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Receiving institution country code"}
	iso8583[69] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Settlement institution country code"}
	iso8583[70] = FieldDescription{ContentType: "n", MaxLen: 3, LenType: "fixed", Label: "Network management information code"}
	iso8583[71] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Message number"}
	iso8583[72] = FieldDescription{ContentType: "n", MaxLen: 4, LenType: "fixed", Label: "Message number, last"}
	iso8583[73] = FieldDescription{ContentType: "n", MaxLen: 6, LenType: "fixed", Label: "Date, action (YYMMDD)"}
	iso8583[74] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Credits, number"}
	iso8583[75] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Credits, reversal number"}
	iso8583[76] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Debits, number"}
	iso8583[77] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Debits, reversal number"}
	iso8583[78] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Transfer number"}
	iso8583[79] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Transfer, reversal number"}
	iso8583[80] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Inquiries number"}
	iso8583[81] = FieldDescription{ContentType: "n", MaxLen: 10, LenType: "fixed", Label: "Authorizations, number"}
	iso8583[82] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Credits, processing fee amount"}
	iso8583[83] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Credits, transaction fee amount"}
	iso8583[84] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Debits, processing fee amount"}
	iso8583[85] = FieldDescription{ContentType: "n", MaxLen: 12, LenType: "fixed", Label: "Debits, transaction fee amount"}
	iso8583[86] = FieldDescription{ContentType: "n", MaxLen: 16, LenType: "fixed", Label: "Credits, amount"}
	iso8583[87] = FieldDescription{ContentType: "n", MaxLen: 16, LenType: "fixed", Label: "Credits, reversal amount"}
	iso8583[88] = FieldDescription{ContentType: "n", MaxLen: 16, LenType: "fixed", Label: "Debits, amount"}
	iso8583[89] = FieldDescription{ContentType: "n", MaxLen: 16, LenType: "fixed", Label: "Debits, reversal amount"}
	iso8583[90] = FieldDescription{ContentType: "n", MaxLen: 42, LenType: "fixed", Label: "Original data elements"}
	iso8583[91] = FieldDescription{ContentType: "an", MaxLen: 1, LenType: "fixed", Label: "File update code"}
	iso8583[92] = FieldDescription{ContentType: "an", MaxLen: 2, LenType: "fixed", Label: "File security code"}
	iso8583[93] = FieldDescription{ContentType: "an", MaxLen: 5, LenType: "fixed", Label: "Response indicator"}
	iso8583[94] = FieldDescription{ContentType: "an", MaxLen: 7, LenType: "fixed", Label: "Service indicator"}
	iso8583[95] = FieldDescription{ContentType: "an", MaxLen: 42, LenType: "fixed", Label: "Replacement amounts"}
	iso8583[96] = FieldDescription{ContentType: "b", MaxLen: 8, LenType: "fixed", Label: "Message security code"}
	iso8583[97] = FieldDescription{ContentType: "an", MaxLen: 17, LenType: "fixed", Label: "Amount, net settlement"}
	iso8583[98] = FieldDescription{ContentType: "ans", MaxLen: 25, LenType: "fixed", Label: "Payee"}
	iso8583[99] = FieldDescription{ContentType: "n", MaxLen: 11, LenType: "llvar", Label: "Settlement institution identification code"}
	iso8583[100] = FieldDescription{ContentType: "n", MaxLen: 11, LenType: "llvar", Label: "Receiving institution identification code"}
	iso8583[101] = FieldDescription{ContentType: "ans", MaxLen: 17, LenType: "llvar", Label: "File name"}
	iso8583[102] = FieldDescription{ContentType: "ans", MaxLen: 28, LenType: "llvar", Label: "Account identification 1"}
	iso8583[103] = FieldDescription{ContentType: "ans", MaxLen: 28, LenType: "llvar", Label: "Account identification 2"}
	iso8583[104] = FieldDescription{ContentType: "ans", MaxLen: 100, LenType: "lllvar", Label: "Transaction description"}
	iso8583[105] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[106] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[107] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[108] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[109] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[110] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[111] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for ISO use"}
	iso8583[112] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[113] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[114] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[115] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[116] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[117] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[118] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[119] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for national use"}
	iso8583[120] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[121] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[122] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[123] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[124] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[125] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[126] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[127] = FieldDescription{ContentType: "ans", MaxLen: 999, LenType: "lllvar", Label: "Reserved for private use"}
	iso8583[128] = FieldDescription{ContentType: "b", MaxLen: 8, LenType: "fixed", Label: "Message authentication code"}

	return iso8583
}
