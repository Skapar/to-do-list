package model

type Task struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    ActiveAt string `json:"activeAt"`
    Status   string `json:"status"`
}
