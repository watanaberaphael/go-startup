package task

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"net/http"
	"time"
)

type Task struct {
	Name        string
	Description string
	CreatedOn   time.Time
}

var taskListTemplate = template.Must(template.New("taskList").Parse(constTaskListTemplate))

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save", save)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Task").Order("-CreatedOn")

	var tasks []Task

	_, err := q.GetAll(c, &tasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := taskListTemplate.Execute(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, taskForm)
}

func save(w http.ResponseWriter, r *http.Request) {
	task := Task{Name: r.FormValue("taskname"), Description: r.FormValue("description"), CreatedOn: time.Now()}

	c := appengine.NewContext(r)

	taskKey := datastore.NewIncompleteKey(c, "Task", nil)

	_, err := datastore.Put(c, taskKey, &task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
