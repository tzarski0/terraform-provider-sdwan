{{- if hasReference .Attributes}}
# Expected import identifier with the format: "{{getProfileParcelName .}}_id{{range .Attributes}}{{if .Reference}},{{.TfName}}{{end}}{{end}}"
terraform import sdwan_{{getProfileParcelName .}}.example "f6b2c44c-693c-4763-b010-895aa3d236bd{{range .Attributes}}{{if .Reference}},{{.Example}}{{end}}{{end}}"
{{- else}}
terraform import sdwan_{{getProfileParcelName .}}.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
{{- end}}