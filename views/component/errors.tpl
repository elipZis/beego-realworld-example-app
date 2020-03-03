{{ define "errors" }}
    <ul class="error-messages">
        {{range $i, $error := .}}
            <li>{{ $i }}: {{ $error.Message }}</li>
        {{end}}
    </ul>
{{ end }}