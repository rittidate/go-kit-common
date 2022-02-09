package message

import (
	"bytes"
	"encoding/json"
)

type Language int

const (
	Thai Language = iota
	English
	Chinese
	Unknown
)

func (p Language) String() string {
	return LanguagesValue[p]
}

func (p Language) FromString(s string) Language {
	return LanguagesName[s]
}

var LanguagesValue = map[Language]string{
	Thai:    "TH",
	English: "EN",
	Chinese: "ZH",
	Unknown: "UNKNOWN",
}

var LanguagesName = map[string]Language{
	"TH":      Thai,
	"EN":      English,
	"ZH":      Chinese,
	"UNKNOWN": Unknown,
}

func (p Language) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(LanguagesValue[p])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (p *Language) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*p = LanguagesName[s]
	return nil
}
