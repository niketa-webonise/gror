{{ range $key, $value := .Names }}
<a href="/docker/config/{{index $.ID $key}}">NAME: {{ $value}}</a>
<br>
{{end}}