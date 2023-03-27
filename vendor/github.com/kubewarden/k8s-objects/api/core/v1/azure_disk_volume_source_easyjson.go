// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonC4f0b390DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *AzureDiskVolumeSource) {
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
		case "cachingMode":
			out.CachingMode = string(in.String())
		case "diskName":
			if in.IsNull() {
				in.Skip()
				out.DiskName = nil
			} else {
				if out.DiskName == nil {
					out.DiskName = new(string)
				}
				*out.DiskName = string(in.String())
			}
		case "diskURI":
			if in.IsNull() {
				in.Skip()
				out.DiskURI = nil
			} else {
				if out.DiskURI == nil {
					out.DiskURI = new(string)
				}
				*out.DiskURI = string(in.String())
			}
		case "fsType":
			out.FSType = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
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
func easyjsonC4f0b390EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in AzureDiskVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CachingMode != "" {
		const prefix string = ",\"cachingMode\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.CachingMode))
	}
	{
		const prefix string = ",\"diskName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.DiskName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DiskName))
		}
	}
	{
		const prefix string = ",\"diskURI\":"
		out.RawString(prefix)
		if in.DiskURI == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.DiskURI))
		}
	}
	if in.FSType != "" {
		const prefix string = ",\"fsType\":"
		out.RawString(prefix)
		out.String(string(in.FSType))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AzureDiskVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC4f0b390EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AzureDiskVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC4f0b390EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AzureDiskVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC4f0b390DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AzureDiskVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC4f0b390DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
