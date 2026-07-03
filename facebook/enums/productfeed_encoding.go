package enums

type ProductfeedEncoding string

const (
	ProductfeedEncodingAutodetect ProductfeedEncoding = "AUTODETECT"
	ProductfeedEncodingLatin1     ProductfeedEncoding = "LATIN1"
	ProductfeedEncodingUtf16be    ProductfeedEncoding = "UTF16BE"
	ProductfeedEncodingUtf16le    ProductfeedEncoding = "UTF16LE"
	ProductfeedEncodingUtf32be    ProductfeedEncoding = "UTF32BE"
	ProductfeedEncodingUtf32le    ProductfeedEncoding = "UTF32LE"
	ProductfeedEncodingUTF8       ProductfeedEncoding = "UTF8"
)

func (value ProductfeedEncoding) String() string {
	return string(value)
}
