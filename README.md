# GoSMSLengthFinder
[![Maintainability](https://api.codeclimate.com/v1/badges/515bfb5c366231a34724/maintainability)](https://codeclimate.com/github/SiddhantAgarwal/GoSMSLengthFinder/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/SiddhantAgarwal/GoSMSLengthFinder)](https://goreportcard.com/report/github.com/SiddhantAgarwal/GoSMSLengthFinder)
[![Coverage Status](https://coveralls.io/repos/github/SiddhantAgarwal/GoSMSLengthFinder/badge.svg?branch=master)](https://coveralls.io/github/SiddhantAgarwal/GoSMSLengthFinder?branch=master)
![Go](https://github.com/SiddhantAgarwal/GoSMSLengthFinder/workflows/Go/badge.svg?branch=master)

A library to Get SMS parts length for given string, it takes into account the charset of message body detecting between GSM7, Unicoded message bodies and giving SMS parts accordingly.

TL;DR

More Info on SMS lengths here
https://www.world-text.com/docs/sms-length.php

**Examples**

Simple GSM7 Char body
```
smsParts := smslength.GetSMSParts("testing")
fmt.Println(smsParts)

Output
1
```

Unicode Char body
```
smsParts := smslength.GetSMSParts("ăZʍ𦞺9r𔡛)(Δݷ֣퐘Řaas􂵜꓄悗償%O򊿟夒Д--2䴛槫𐻵춡󣛃<ꣻئ񯕱񚲪Ʌ뢉񤲔՛􏅛􁐜󯞒؀,خs񚺣򳼙sϞ񴧆:ʣ򆝳ጳ(򶳫蝹򝚿ӂ交柪?٦쳕磏͒qtԜ󣜐҈쏒W㗁ԂD欕D"憬ˆ!Ӹⵗ򏘯𸎫ɝ󁦡?Ӯ]܂ҥ4󘒙򚲣읊Ζ􉵑}󪆓豱շښܾ纒藯b䬢ܫtؕjx񑾖ńֺ<𝚝􉆢ˬîґNڧ󓺻h񃴵󍣧}Ǜ贈#<ᜏ񞹚ނ'򧿲b]˽刺񽞙Ѐ𨦨2ͫ֞􇓄쪧𙅝򳉎")
fmt.Prinln(smsParts)

Output
3
```