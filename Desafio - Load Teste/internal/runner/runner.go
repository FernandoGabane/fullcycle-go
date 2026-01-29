package runner

import (
    "go-load-tester/internal/stats"
    "net/http"
    "sync"
)

func Run(url string, total int, concurrency int) stats.Report {
    client := http.Client{}
    report := stats.NewReport()

    jobs := make(chan struct{}, total)
    wg := sync.WaitGroup{}

    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for range jobs {
                resp, err := client.Get(url)
                if err != nil {
                    report.AddError(0)
                    continue
                }
                report.AddStatus(resp.StatusCode)
                resp.Body.Close()
            }
        }()
    }

    for i := 0; i < total; i++ {
        jobs <- struct{}{}
    }
    close(jobs)

    wg.Wait()
    return report
}
