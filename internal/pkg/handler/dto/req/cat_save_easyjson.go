// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package req

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB21bb8b0DecodeAwesomeProjectInternalPkgHandlerDtoReq(in *jlexer.Lexer, out *CatSave) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "age":
			out.Age = int(in.Int())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB21bb8b0EncodeAwesomeProjectInternalPkgHandlerDtoReq(out *jwriter.Writer, in CatSave) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CatSave) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB21bb8b0EncodeAwesomeProjectInternalPkgHandlerDtoReq(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CatSave) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB21bb8b0EncodeAwesomeProjectInternalPkgHandlerDtoReq(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CatSave) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB21bb8b0DecodeAwesomeProjectInternalPkgHandlerDtoReq(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CatSave) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB21bb8b0DecodeAwesomeProjectInternalPkgHandlerDtoReq(l, v)
}