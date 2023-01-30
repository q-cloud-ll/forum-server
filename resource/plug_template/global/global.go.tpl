package global

{{- if .HasGlobal }}

import "forum-server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}