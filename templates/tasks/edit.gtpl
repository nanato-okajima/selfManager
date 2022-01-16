{{ template "header" .}}
<form action="/task/{{ .ID }}" method="POST">
    <div>
        <label for="name">タスク名</label>
        <input type="text" name="name" id="name" value="{{ .Name }}">
    </div>
    <div>
        <label for="status">ステータス</label>
        <select>
            <option value="0" {{ if eq .Status 0 }} selected {{ end }}>未着手</option>
            <option value="1" {{ if eq .Status 1 }} selected {{ end }}>着手</option>
            <option value="2" {{ if eq .Status 2 }} selected {{ end }}>完了</option>
        </select>
    </div>
    <div>
        <label for="due-datetime">期日</label>
        <input type="date" name="due-datetime" value={{ .DueDatetime }}>
    </div>
    <input type="submit">
</form>
{{ template "footer" .}}
