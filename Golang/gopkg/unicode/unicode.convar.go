/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: unicode
 **Element: unicode.convar
 **Type: const,var
 ------------------------------------------------------------------------------------
 **Definition:
 const (
    MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
    ReplacementChar = '\uFFFD'     // Represents invalid code points.
    MaxASCII        = '\u007F'     // maximum ASCII value.
    MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
 )

 const (
    UpperCase = iota
    LowerCase
    TitleCase
    MaxCase
 )
 Indices into the Delta arrays inside CaseRanges for case mapping.
 CaseRange 中 Delta 数组的下标，以用于写法映射。

 const (
    UpperLower = MaxRune + 1 // (Cannot be a valid delta.)
 )
 If the Delta field of a CaseRange is UpperLower, it means this CaseRange represents a sequence of the form (say) Upper Lower Upper Lower.
 若 CaseRange 的 Delta 字段为 UpperLower 或 LowerUpper，则该 CaseRange 即表示
 （所谓的）“Upper Lower Upper Lower”序列。

 const Version = "10.0.0"
 Version is the Unicode edition from which the tables are derived.
 Version为此表所用的 Unicode 版本。

 Variables
 var (
    Cc     = _Cc // Cc is the set of Unicode characters in category Cc.
    Cf     = _Cf // Cf is the set of Unicode characters in category Cf.
    Co     = _Co // Co is the set of Unicode characters in category Co.
    Cs     = _Cs // Cs is the set of Unicode characters in category Cs.
    Digit  = _Nd // Digit is the set of Unicode characters with the "decimal digit" property.
    Nd     = _Nd // Nd is the set of Unicode characters in category Nd.
    Letter = _L  // Letter/L is the set of Unicode letters, category L.
    L      = _L
    Lm     = _Lm // Lm is the set of Unicode characters in category Lm.
    Lo     = _Lo // Lo is the set of Unicode characters in category Lo.
    Lower  = _Ll // Lower is the set of Unicode lower case letters.
    Ll     = _Ll // Ll is the set of Unicode characters in category Ll.
    Mark   = _M  // Mark/M is the set of Unicode mark characters, category M.
    M      = _M
    Mc     = _Mc // Mc is the set of Unicode characters in category Mc.
    Me     = _Me // Me is the set of Unicode characters in category Me.
    Mn     = _Mn // Mn is the set of Unicode characters in category Mn.
    Nl     = _Nl // Nl is the set of Unicode characters in category Nl.
    No     = _No // No is the set of Unicode characters in category No.
    Number = _N  // Number/N is the set of Unicode number characters, category N.
    N      = _N
    Other  = _C // Other/C is the set of Unicode control and special characters, category C.
    C      = _C
    Pc     = _Pc // Pc is the set of Unicode characters in category Pc.
    Pd     = _Pd // Pd is the set of Unicode characters in category Pd.
    Pe     = _Pe // Pe is the set of Unicode characters in category Pe.
    Pf     = _Pf // Pf is the set of Unicode characters in category Pf.
    Pi     = _Pi // Pi is the set of Unicode characters in category Pi.
    Po     = _Po // Po is the set of Unicode characters in category Po.
    Ps     = _Ps // Ps is the set of Unicode characters in category Ps.
    Punct  = _P  // Punct/P is the set of Unicode punctuation characters, category P.
    P      = _P
    Sc     = _Sc // Sc is the set of Unicode characters in category Sc.
    Sk     = _Sk // Sk is the set of Unicode characters in category Sk.
    Sm     = _Sm // Sm is the set of Unicode characters in category Sm.
    So     = _So // So is the set of Unicode characters in category So.
    Space  = _Z  // Space/Z is the set of Unicode space characters, category Z.
    Z      = _Z
    Symbol = _S // Symbol/S is the set of Unicode symbol characters, category S.
    S      = _S
    Title  = _Lt // Title is the set of Unicode title case letters.
    Lt     = _Lt // Lt is the set of Unicode characters in category Lt.
    Upper  = _Lu // Upper is the set of Unicode upper case letters.
    Lu     = _Lu // Lu is the set of Unicode characters in category Lu.
    Zl     = _Zl // Zl is the set of Unicode characters in category Zl.
    Zp     = _Zp // Zp is the set of Unicode characters in category Zp.
    Zs     = _Zs // Zs is the set of Unicode characters in category Zs.
 )
 These variables have type *RangeTable.

 var (
    Adlam                  = _Adlam                  // Adlam is the set of Unicode characters in script Adlam.
    Ahom                   = _Ahom                   // Ahom is the set of Unicode characters in script Ahom.
    Anatolian_Hieroglyphs  = _Anatolian_Hieroglyphs  // Anatolian_Hieroglyphs is the set of Unicode characters in script Anatolian_Hieroglyphs.
    Arabic                 = _Arabic                 // Arabic is the set of Unicode characters in script Arabic.
    Armenian               = _Armenian               // Armenian is the set of Unicode characters in script Armenian.
    Avestan                = _Avestan                // Avestan is the set of Unicode characters in script Avestan.
    Balinese               = _Balinese               // Balinese is the set of Unicode characters in script Balinese.
    Bamum                  = _Bamum                  // Bamum is the set of Unicode characters in script Bamum.
    Bassa_Vah              = _Bassa_Vah              // Bassa_Vah is the set of Unicode characters in script Bassa_Vah.
    Batak                  = _Batak                  // Batak is the set of Unicode characters in script Batak.
    Bengali                = _Bengali                // Bengali is the set of Unicode characters in script Bengali.
    Bhaiksuki              = _Bhaiksuki              // Bhaiksuki is the set of Unicode characters in script Bhaiksuki.
    Bopomofo               = _Bopomofo               // Bopomofo is the set of Unicode characters in script Bopomofo.
    Brahmi                 = _Brahmi                 // Brahmi is the set of Unicode characters in script Brahmi.
    Braille                = _Braille                // Braille is the set of Unicode characters in script Braille.
    Buginese               = _Buginese               // Buginese is the set of Unicode characters in script Buginese.
    Buhid                  = _Buhid                  // Buhid is the set of Unicode characters in script Buhid.
    Canadian_Aboriginal    = _Canadian_Aboriginal    // Canadian_Aboriginal is the set of Unicode characters in script Canadian_Aboriginal.
    Carian                 = _Carian                 // Carian is the set of Unicode characters in script Carian.
    Caucasian_Albanian     = _Caucasian_Albanian     // Caucasian_Albanian is the set of Unicode characters in script Caucasian_Albanian.
    Chakma                 = _Chakma                 // Chakma is the set of Unicode characters in script Chakma.
    Cham                   = _Cham                   // Cham is the set of Unicode characters in script Cham.
    Cherokee               = _Cherokee               // Cherokee is the set of Unicode characters in script Cherokee.
    Common                 = _Common                 // Common is the set of Unicode characters in script Common.
    Coptic                 = _Coptic                 // Coptic is the set of Unicode characters in script Coptic.
    Cuneiform              = _Cuneiform              // Cuneiform is the set of Unicode characters in script Cuneiform.
    Cypriot                = _Cypriot                // Cypriot is the set of Unicode characters in script Cypriot.
    Cyrillic               = _Cyrillic               // Cyrillic is the set of Unicode characters in script Cyrillic.
    Deseret                = _Deseret                // Deseret is the set of Unicode characters in script Deseret.
    Devanagari             = _Devanagari             // Devanagari is the set of Unicode characters in script Devanagari.
    Duployan               = _Duployan               // Duployan is the set of Unicode characters in script Duployan.
    Egyptian_Hieroglyphs   = _Egyptian_Hieroglyphs   // Egyptian_Hieroglyphs is the set of Unicode characters in script Egyptian_Hieroglyphs.
    Elbasan                = _Elbasan                // Elbasan is the set of Unicode characters in script Elbasan.
    Ethiopic               = _Ethiopic               // Ethiopic is the set of Unicode characters in script Ethiopic.
    Georgian               = _Georgian               // Georgian is the set of Unicode characters in script Georgian.
    Glagolitic             = _Glagolitic             // Glagolitic is the set of Unicode characters in script Glagolitic.
    Gothic                 = _Gothic                 // Gothic is the set of Unicode characters in script Gothic.
    Grantha                = _Grantha                // Grantha is the set of Unicode characters in script Grantha.
    Greek                  = _Greek                  // Greek is the set of Unicode characters in script Greek.
    Gujarati               = _Gujarati               // Gujarati is the set of Unicode characters in script Gujarati.
    Gurmukhi               = _Gurmukhi               // Gurmukhi is the set of Unicode characters in script Gurmukhi.
    Han                    = _Han                    // Han is the set of Unicode characters in script Han.
    Hangul                 = _Hangul                 // Hangul is the set of Unicode characters in script Hangul.
    Hanunoo                = _Hanunoo                // Hanunoo is the set of Unicode characters in script Hanunoo.
    Hatran                 = _Hatran                 // Hatran is the set of Unicode characters in script Hatran.
    Hebrew                 = _Hebrew                 // Hebrew is the set of Unicode characters in script Hebrew.
    Hiragana               = _Hiragana               // Hiragana is the set of Unicode characters in script Hiragana.
    Imperial_Aramaic       = _Imperial_Aramaic       // Imperial_Aramaic is the set of Unicode characters in script Imperial_Aramaic.
    Inherited              = _Inherited              // Inherited is the set of Unicode characters in script Inherited.
    Inscriptional_Pahlavi  = _Inscriptional_Pahlavi  // Inscriptional_Pahlavi is the set of Unicode characters in script Inscriptional_Pahlavi.
    Inscriptional_Parthian = _Inscriptional_Parthian // Inscriptional_Parthian is the set of Unicode characters in script Inscriptional_Parthian.
    Javanese               = _Javanese               // Javanese is the set of Unicode characters in script Javanese.
    Kaithi                 = _Kaithi                 // Kaithi is the set of Unicode characters in script Kaithi.
    Kannada                = _Kannada                // Kannada is the set of Unicode characters in script Kannada.
    Katakana               = _Katakana               // Katakana is the set of Unicode characters in script Katakana.
    Kayah_Li               = _Kayah_Li               // Kayah_Li is the set of Unicode characters in script Kayah_Li.
    Kharoshthi             = _Kharoshthi             // Kharoshthi is the set of Unicode characters in script Kharoshthi.
    Khmer                  = _Khmer                  // Khmer is the set of Unicode characters in script Khmer.
    Khojki                 = _Khojki                 // Khojki is the set of Unicode characters in script Khojki.
    Khudawadi              = _Khudawadi              // Khudawadi is the set of Unicode characters in script Khudawadi.
    Lao                    = _Lao                    // Lao is the set of Unicode characters in script Lao.
    Latin                  = _Latin                  // Latin is the set of Unicode characters in script Latin.
    Lepcha                 = _Lepcha                 // Lepcha is the set of Unicode characters in script Lepcha.
    Limbu                  = _Limbu                  // Limbu is the set of Unicode characters in script Limbu.
    Linear_A               = _Linear_A               // Linear_A is the set of Unicode characters in script Linear_A.
    Linear_B               = _Linear_B               // Linear_B is the set of Unicode characters in script Linear_B.
    Lisu                   = _Lisu                   // Lisu is the set of Unicode characters in script Lisu.
    Lycian                 = _Lycian                 // Lycian is the set of Unicode characters in script Lycian.
    Lydian                 = _Lydian                 // Lydian is the set of Unicode characters in script Lydian.
    Mahajani               = _Mahajani               // Mahajani is the set of Unicode characters in script Mahajani.
    Malayalam              = _Malayalam              // Malayalam is the set of Unicode characters in script Malayalam.
    Mandaic                = _Mandaic                // Mandaic is the set of Unicode characters in script Mandaic.
    Manichaean             = _Manichaean             // Manichaean is the set of Unicode characters in script Manichaean.
    Marchen                = _Marchen                // Marchen is the set of Unicode characters in script Marchen.
    Masaram_Gondi          = _Masaram_Gondi          // Masaram_Gondi is the set of Unicode characters in script Masaram_Gondi.
    Meetei_Mayek           = _Meetei_Mayek           // Meetei_Mayek is the set of Unicode characters in script Meetei_Mayek.
    Mende_Kikakui          = _Mende_Kikakui          // Mende_Kikakui is the set of Unicode characters in script Mende_Kikakui.
    Meroitic_Cursive       = _Meroitic_Cursive       // Meroitic_Cursive is the set of Unicode characters in script Meroitic_Cursive.
    Meroitic_Hieroglyphs   = _Meroitic_Hieroglyphs   // Meroitic_Hieroglyphs is the set of Unicode characters in script Meroitic_Hieroglyphs.
    Miao                   = _Miao                   // Miao is the set of Unicode characters in script Miao.
    Modi                   = _Modi                   // Modi is the set of Unicode characters in script Modi.
    Mongolian              = _Mongolian              // Mongolian is the set of Unicode characters in script Mongolian.
    Mro                    = _Mro                    // Mro is the set of Unicode characters in script Mro.
    Multani                = _Multani                // Multani is the set of Unicode characters in script Multani.
    Myanmar                = _Myanmar                // Myanmar is the set of Unicode characters in script Myanmar.
    Nabataean              = _Nabataean              // Nabataean is the set of Unicode characters in script Nabataean.
    New_Tai_Lue            = _New_Tai_Lue            // New_Tai_Lue is the set of Unicode characters in script New_Tai_Lue.
    Newa                   = _Newa                   // Newa is the set of Unicode characters in script Newa.
    Nko                    = _Nko                    // Nko is the set of Unicode characters in script Nko.
    Nushu                  = _Nushu                  // Nushu is the set of Unicode characters in script Nushu.
    Ogham                  = _Ogham                  // Ogham is the set of Unicode characters in script Ogham.
    Ol_Chiki               = _Ol_Chiki               // Ol_Chiki is the set of Unicode characters in script Ol_Chiki.
    Old_Hungarian          = _Old_Hungarian          // Old_Hungarian is the set of Unicode characters in script Old_Hungarian.
    Old_Italic             = _Old_Italic             // Old_Italic is the set of Unicode characters in script Old_Italic.
    Old_North_Arabian      = _Old_North_Arabian      // Old_North_Arabian is the set of Unicode characters in script Old_North_Arabian.
    Old_Permic             = _Old_Permic             // Old_Permic is the set of Unicode characters in script Old_Permic.
    Old_Persian            = _Old_Persian            // Old_Persian is the set of Unicode characters in script Old_Persian.
    Old_South_Arabian      = _Old_South_Arabian      // Old_South_Arabian is the set of Unicode characters in script Old_South_Arabian.
    Old_Turkic             = _Old_Turkic             // Old_Turkic is the set of Unicode characters in script Old_Turkic.
    Oriya                  = _Oriya                  // Oriya is the set of Unicode characters in script Oriya.
    Osage                  = _Osage                  // Osage is the set of Unicode characters in script Osage.
    Osmanya                = _Osmanya                // Osmanya is the set of Unicode characters in script Osmanya.
    Pahawh_Hmong           = _Pahawh_Hmong           // Pahawh_Hmong is the set of Unicode characters in script Pahawh_Hmong.
    Palmyrene              = _Palmyrene              // Palmyrene is the set of Unicode characters in script Palmyrene.
    Pau_Cin_Hau            = _Pau_Cin_Hau            // Pau_Cin_Hau is the set of Unicode characters in script Pau_Cin_Hau.
    Phags_Pa               = _Phags_Pa               // Phags_Pa is the set of Unicode characters in script Phags_Pa.
    Phoenician             = _Phoenician             // Phoenician is the set of Unicode characters in script Phoenician.
    Psalter_Pahlavi        = _Psalter_Pahlavi        // Psalter_Pahlavi is the set of Unicode characters in script Psalter_Pahlavi.
    Rejang                 = _Rejang                 // Rejang is the set of Unicode characters in script Rejang.
    Runic                  = _Runic                  // Runic is the set of Unicode characters in script Runic.
    Samaritan              = _Samaritan              // Samaritan is the set of Unicode characters in script Samaritan.
    Saurashtra             = _Saurashtra             // Saurashtra is the set of Unicode characters in script Saurashtra.
    Sharada                = _Sharada                // Sharada is the set of Unicode characters in script Sharada.
    Shavian                = _Shavian                // Shavian is the set of Unicode characters in script Shavian.
    Siddham                = _Siddham                // Siddham is the set of Unicode characters in script Siddham.
    SignWriting            = _SignWriting            // SignWriting is the set of Unicode characters in script SignWriting.
    Sinhala                = _Sinhala                // Sinhala is the set of Unicode characters in script Sinhala.
    Sora_Sompeng           = _Sora_Sompeng           // Sora_Sompeng is the set of Unicode characters in script Sora_Sompeng.
    Soyombo                = _Soyombo                // Soyombo is the set of Unicode characters in script Soyombo.
    Sundanese              = _Sundanese              // Sundanese is the set of Unicode characters in script Sundanese.
    Syloti_Nagri           = _Syloti_Nagri           // Syloti_Nagri is the set of Unicode characters in script Syloti_Nagri.
    Syriac                 = _Syriac                 // Syriac is the set of Unicode characters in script Syriac.
    Tagalog                = _Tagalog                // Tagalog is the set of Unicode characters in script Tagalog.
    Tagbanwa               = _Tagbanwa               // Tagbanwa is the set of Unicode characters in script Tagbanwa.
    Tai_Le                 = _Tai_Le                 // Tai_Le is the set of Unicode characters in script Tai_Le.
    Tai_Tham               = _Tai_Tham               // Tai_Tham is the set of Unicode characters in script Tai_Tham.
    Tai_Viet               = _Tai_Viet               // Tai_Viet is the set of Unicode characters in script Tai_Viet.
    Takri                  = _Takri                  // Takri is the set of Unicode characters in script Takri.
    Tamil                  = _Tamil                  // Tamil is the set of Unicode characters in script Tamil.
    Tangut                 = _Tangut                 // Tangut is the set of Unicode characters in script Tangut.
    Telugu                 = _Telugu                 // Telugu is the set of Unicode characters in script Telugu.
    Thaana                 = _Thaana                 // Thaana is the set of Unicode characters in script Thaana.
    Thai                   = _Thai                   // Thai is the set of Unicode characters in script Thai.
    Tibetan                = _Tibetan                // Tibetan is the set of Unicode characters in script Tibetan.
    Tifinagh               = _Tifinagh               // Tifinagh is the set of Unicode characters in script Tifinagh.
    Tirhuta                = _Tirhuta                // Tirhuta is the set of Unicode characters in script Tirhuta.
    Ugaritic               = _Ugaritic               // Ugaritic is the set of Unicode characters in script Ugaritic.
    Vai                    = _Vai                    // Vai is the set of Unicode characters in script Vai.
    Warang_Citi            = _Warang_Citi            // Warang_Citi is the set of Unicode characters in script Warang_Citi.
    Yi                     = _Yi                     // Yi is the set of Unicode characters in script Yi.
    Zanabazar_Square       = _Zanabazar_Square       // Zanabazar_Square is the set of Unicode characters in script Zanabazar_Square.
 )
 These variables have type *RangeTable.

 var (
    ASCII_Hex_Digit                    = _ASCII_Hex_Digit                    // ASCII_Hex_Digit is the set of Unicode characters with property ASCII_Hex_Digit.
    Bidi_Control                       = _Bidi_Control                       // Bidi_Control is the set of Unicode characters with property Bidi_Control.
    Dash                               = _Dash                               // Dash is the set of Unicode characters with property Dash.
    Deprecated                         = _Deprecated                         // Deprecated is the set of Unicode characters with property Deprecated.
    Diacritic                          = _Diacritic                          // Diacritic is the set of Unicode characters with property Diacritic.
    Extender                           = _Extender                           // Extender is the set of Unicode characters with property Extender.
    Hex_Digit                          = _Hex_Digit                          // Hex_Digit is the set of Unicode characters with property Hex_Digit.
    Hyphen                             = _Hyphen                             // Hyphen is the set of Unicode characters with property Hyphen.
    IDS_Binary_Operator                = _IDS_Binary_Operator                // IDS_Binary_Operator is the set of Unicode characters with property IDS_Binary_Operator.
    IDS_Trinary_Operator               = _IDS_Trinary_Operator               // IDS_Trinary_Operator is the set of Unicode characters with property IDS_Trinary_Operator.
    Ideographic                        = _Ideographic                        // Ideographic is the set of Unicode characters with property Ideographic.
    Join_Control                       = _Join_Control                       // Join_Control is the set of Unicode characters with property Join_Control.
    Logical_Order_Exception            = _Logical_Order_Exception            // Logical_Order_Exception is the set of Unicode characters with property Logical_Order_Exception.
    Noncharacter_Code_Point            = _Noncharacter_Code_Point            // Noncharacter_Code_Point is the set of Unicode characters with property Noncharacter_Code_Point.
    Other_Alphabetic                   = _Other_Alphabetic                   // Other_Alphabetic is the set of Unicode characters with property Other_Alphabetic.
    Other_Default_Ignorable_Code_Point = _Other_Default_Ignorable_Code_Point // Other_Default_Ignorable_Code_Point is the set of Unicode characters with property Other_Default_Ignorable_Code_Point.
    Other_Grapheme_Extend              = _Other_Grapheme_Extend              // Other_Grapheme_Extend is the set of Unicode characters with property Other_Grapheme_Extend.
    Other_ID_Continue                  = _Other_ID_Continue                  // Other_ID_Continue is the set of Unicode characters with property Other_ID_Continue.
    Other_ID_Start                     = _Other_ID_Start                     // Other_ID_Start is the set of Unicode characters with property Other_ID_Start.
    Other_Lowercase                    = _Other_Lowercase                    // Other_Lowercase is the set of Unicode characters with property Other_Lowercase.
    Other_Math                         = _Other_Math                         // Other_Math is the set of Unicode characters with property Other_Math.
    Other_Uppercase                    = _Other_Uppercase                    // Other_Uppercase is the set of Unicode characters with property Other_Uppercase.
    Pattern_Syntax                     = _Pattern_Syntax                     // Pattern_Syntax is the set of Unicode characters with property Pattern_Syntax.
    Pattern_White_Space                = _Pattern_White_Space                // Pattern_White_Space is the set of Unicode characters with property Pattern_White_Space.
    Prepended_Concatenation_Mark       = _Prepended_Concatenation_Mark       // Prepended_Concatenation_Mark is the set of Unicode characters with property Prepended_Concatenation_Mark.
    Quotation_Mark                     = _Quotation_Mark                     // Quotation_Mark is the set of Unicode characters with property Quotation_Mark.
    Radical                            = _Radical                            // Radical is the set of Unicode characters with property Radical.
    Regional_Indicator                 = _Regional_Indicator                 // Regional_Indicator is the set of Unicode characters with property Regional_Indicator.
    STerm                              = _Sentence_Terminal                  // STerm is an alias for Sentence_Terminal.
    Sentence_Terminal                  = _Sentence_Terminal                  // Sentence_Terminal is the set of Unicode characters with property Sentence_Terminal.
    Soft_Dotted                        = _Soft_Dotted                        // Soft_Dotted is the set of Unicode characters with property Soft_Dotted.
    Terminal_Punctuation               = _Terminal_Punctuation               // Terminal_Punctuation is the set of Unicode characters with property Terminal_Punctuation.
    Unified_Ideograph                  = _Unified_Ideograph                  // Unified_Ideograph is the set of Unicode characters with property Unified_Ideograph.
    Variation_Selector                 = _Variation_Selector                 // Variation_Selector is the set of Unicode characters with property Variation_Selector.
    White_Space                        = _White_Space                        // White_Space is the set of Unicode characters with property White_Space.
 )
 These variables have type *RangeTable.

 var CaseRanges = _CaseRanges
 CaseRanges is the table describing case mappings for all letters with non-self mappings.

 var Categories = map[string]*RangeTable{
    "C":  C,
    "Cc": Cc,
    "Cf": Cf,
    "Co": Co,
    "Cs": Cs,
    "L":  L,
    "Ll": Ll,
    "Lm": Lm,
    "Lo": Lo,
    "Lt": Lt,
    "Lu": Lu,
    "M":  M,
    "Mc": Mc,
    "Me": Me,
    "Mn": Mn,
    "N":  N,
    "Nd": Nd,
    "Nl": Nl,
    "No": No,
    "P":  P,
    "Pc": Pc,
    "Pd": Pd,
    "Pe": Pe,
    "Pf": Pf,
    "Pi": Pi,
    "Po": Po,
    "Ps": Ps,
    "S":  S,
    "Sc": Sc,
    "Sk": Sk,
    "Sm": Sm,
    "So": So,
    "Z":  Z,
    "Zl": Zl,
    "Zp": Zp,
    "Zs": Zs,
 }
 Categories is the set of Unicode category tables.

 var FoldCategory = map[string]*RangeTable{
    "L":  foldL,
    "Ll": foldLl,
    "Lt": foldLt,
    "Lu": foldLu,
    "M":  foldM,
    "Mn": foldMn,
 }
 FoldCategory maps a category name to a table of code points outside the category that are equivalent under simple case folding to code points inside the category. If there is no entry for a category name, there are no such points.

 var FoldScript = map[string]*RangeTable{
    "Common":    foldCommon,
    "Greek":     foldGreek,
    "Inherited": foldInherited,
 }
 FoldScript maps a script name to a table of code points outside the script that are equivalent under simple case folding to code points inside the script. If there is no entry for a script name, there are no such points.

 var GraphicRanges = []*RangeTable{
    L, M, N, P, S, Zs,
 }
 GraphicRanges defines the set of graphic characters according to Unicode.

 var PrintRanges = []*RangeTable{
    L, M, N, P, S,
 }
 PrintRanges defines the set of printable characters according to Go. ASCII space, U+0020, is handled separately.

 var Properties = map[string]*RangeTable{
    "ASCII_Hex_Digit":                    ASCII_Hex_Digit,
    "Bidi_Control":                       Bidi_Control,
    "Dash":                               Dash,
    "Deprecated":                         Deprecated,
    "Diacritic":                          Diacritic,
    "Extender":                           Extender,
    "Hex_Digit":                          Hex_Digit,
    "Hyphen":                             Hyphen,
    "IDS_Binary_Operator":                IDS_Binary_Operator,
    "IDS_Trinary_Operator":               IDS_Trinary_Operator,
    "Ideographic":                        Ideographic,
    "Join_Control":                       Join_Control,
    "Logical_Order_Exception":            Logical_Order_Exception,
    "Noncharacter_Code_Point":            Noncharacter_Code_Point,
    "Other_Alphabetic":                   Other_Alphabetic,
    "Other_Default_Ignorable_Code_Point": Other_Default_Ignorable_Code_Point,
    "Other_Grapheme_Extend":              Other_Grapheme_Extend,
    "Other_ID_Continue":                  Other_ID_Continue,
    "Other_ID_Start":                     Other_ID_Start,
    "Other_Lowercase":                    Other_Lowercase,
    "Other_Math":                         Other_Math,
    "Other_Uppercase":                    Other_Uppercase,
    "Pattern_Syntax":                     Pattern_Syntax,
    "Pattern_White_Space":                Pattern_White_Space,
    "Prepended_Concatenation_Mark":       Prepended_Concatenation_Mark,
    "Quotation_Mark":                     Quotation_Mark,
    "Radical":                            Radical,
    "Regional_Indicator":                 Regional_Indicator,
    "Sentence_Terminal":                  Sentence_Terminal,
    "STerm":                              Sentence_Terminal,
    "Soft_Dotted":                        Soft_Dotted,
    "Terminal_Punctuation":               Terminal_Punctuation,
    "Unified_Ideograph":                  Unified_Ideograph,
    "Variation_Selector":                 Variation_Selector,
    "White_Space":                        White_Space,
 }
 Properties is the set of Unicode property tables.

 var Scripts = map[string]*RangeTable{
	//141 elements not displayed

 }
 Scripts is the set of Unicode script tables.
 ------------------------------------------------------------------------------------
 **Description:
 ------------------------------------------------------------------------------------
 **要点总结:
 1. go语言的所有代码都是UTF8的，所以如果我们在程序中的字符串都是utf8编码的，
 但是我们的单个字符（单引号扩起来的）却是unicode的;
 2. const和var是unicode包的核心，需要仔细阅读源码。
*************************************************************************************/
package main

import (
	"fmt"
	"unicode"
)

func main() {

	// constant with mixed type runes
	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}
}
