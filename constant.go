package task

const taskForm = `
<!DOCTYPE html>
<html lang="en">
<body>
<form action="/save" method="post">
<p>Task Name: <input type="text" name="taskname" ></p>
<p>Description: <input type="text" name="description" ></p>
<p><input type="submit" value=""Submit></p>
</form>
</body>
</html>`

const taskTemplateHTML = `
<!DOCTYPE html>
<html lang="en">
<body>
<p>New Task has been created:</p>
<div>Task: {{.Name}}</div>
<div>Description: {{.Description}}</div>
</body>
</html>`

const constTaskListTemplate = `
<!DOCTYPE html>
<html lang="en">
<body>
<p>Task list:</p>
{{range .}}
	<p>{{.Name}} - {{.Description}}</p>
{{end}}
<p><a href="/create">Create task</a>
</body>
</html>`
