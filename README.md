# tg-sender

This is a yet another service that get you ability to send messages from various systems to telegram

## Why yet another?

Because I want single service to manage this from different sources. So you cloud send messages from Alertmanager, Rundeck or any your script by simple POST request with json data.

## Usage

You need to add config.yaml

```yaml
rundeck:
  chatid: -1001635818539
  token: BDKSDFL:21412411241241212412412412
  template: rundeck.tpl
```

and template `template/rundeck.tpl`

```
Job:      {{ .execution.job.name }}
Project:  {{ .execution.project }}
Link:     {{ .execution.href }}
Status:   {{ .execution.status }}
```

### exclude in template

Also you cloud exclude some key from message.
In this example for alertmanager template excludes unneccesary labels from result message:

```
{{- $exclude_labels := list
        "send_to_beru"
        "kubernetes_name"
        "kubernetes_namespace"
        "kubernetes_node"
        "role"
        "job"
        "app"
        "app_kubernetes_io_managed_by"
        "app_kubernetes_io_instance"
        "app_kubernetes_io_name"
        "helm_sh_chart"
        "alertname"
        "alerting"
        "server"
        "severity"
        "prometheus"
        "service"
        "container"
        "condition"
        "endpoint"
        "namespace"
        "pod"
-}}

{{- range .alerts }}
{{ if eq .status "firing"}}🔥 {{ else }}✅ {{ end }}{{ .annotations.description }}
{{- if .annotations.summary }}
{{ .annotations.summary }}
{{- end -}}
{{- if .labels }}
{{- range $key, $value := .labels }}{{ if  ne $key "alertname" }}
    {{- if has $key $exclude_labels }}
    {{- else }}
    {{ $key }}: {{ $value }}{{ end }}{{ end }}
    {{- end }}
{{- end }}
{{- end }}
```

You could make another rules for your purposes, in templates user [sprig](http://masterminds.github.io/sprig/) library

## for lint used
```
gofmt -w main.go
```
```
go install github.com/mgechev/revive@latest
revive
```

