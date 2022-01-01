{{template "header" .}}
    <a href="/task/create">タスク登録</a>
    <h1>タスク一覧</h1>
    {{ range . }}
        <a href="/task/{{ .ID }}">{{ .Name }}</a>
    {{ end }}
    <p></p>
{{template "footer" .}}
