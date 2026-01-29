package stats

import "sync"

type Report struct {
    Total   int
    Success int
    Errors  map[int]int
    mu      sync.Mutex
}

func NewReport() Report {
    return Report{
        Errors: make(map[int]int),
    }
}

func (r *Report) AddStatus(code int) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.Total++
    if code == 200 {
        r.Success++
    } else {
        r.Errors[code]++
    }
}

func (r *Report) AddError(code int) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.Total++
    r.Errors[code]++
}
