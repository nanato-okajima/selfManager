{{template "header" .}}
    <a href="/task/create">タスク登録</a>
    <h1>タスク一覧</h1>
    {{ range . }}
        <a href="/task/{{ .ID }}">{{ .Name }}</a>
        <form action="/task/delete/{{ .ID }}" method="POST">
            <input type="submit" value="削除">
        </form>
    {{ end }}
    <p></p>
{{template "footer" .}}
