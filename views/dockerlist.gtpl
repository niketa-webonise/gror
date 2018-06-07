{{ range $key, $value := .Names }}
<a href="/docker/config/{{index $.Id $key}}">NAME: {{ $value}}</a>
<br>
{{end}}