package global

{{- if .HasGlobal }}

import "forum/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}