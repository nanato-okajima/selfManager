{{template "header" .}}
<form action="/task/create" method="POST">
    <div>
        <label for="name">タスク名</label>
        <input type="text" name="name" id="name">
    </div>
    <div>
        <label for="status">ステータス</label>
        <select name="status" id="status">
            <option value="1">未着手</option>
        </select>
    </div>
    <div>
        <label for="due-datetime"></label>
        <input type="date" name="due-datetime">
    </div>
    <input type="submit">
</form>
{{template "footer" .}}
