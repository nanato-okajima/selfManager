{{ template "header" .}}
<form action="/task/{{ .ID }}" method="POST">
    <div>
        <label for="name">タスク名</label>
        <input type="text" name="name" id="name" value="{{ .Name }}">
    </div>
    <div>
        <label for="status">ステータス</label>
        <select>
            <option value="2">完了</option>
        </select>
    </div>
    <div>
        <label for="due-datetime">期日</label>
        <input type="date" name="due-datetime">
    </div>
    <input type="submit">
</form>
{{ template "footer" .}}
