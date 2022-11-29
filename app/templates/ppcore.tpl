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
{{ if eq .status "firing"}}🔥 {{ .labels.alertname }} {{ else }}✅ {{ .labels.alertname }} {{ end }}
{{- if .annotations.summary }}
Annotations: {{ .annotations.summary }}
{{- end -}}
{{/*{{ keys .labels | sortAlpha }} */}}
Labels:{{ range $key, $value := .labels }}{{ if  ne $key "alertname" }}
    {{- if has $key $exclude_labels }}
    {{- else }}
    {{ $key }}: {{ $value }}{{ end }}{{ end }}
    {{- end }}
{{ end }}
