# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with Go - Jobs scheduling
# Extend with Go - Jobs scheduling
If you have tasks that need to be performed periodically, you could setup crontab-like jobs with the
builtin
cron package
Each scheduled job runs in its own goroutine and must have:
-name - identifier for the scheduled job; could be used to replace or remove an existing job
-cron expression like 0 0 * * * (
supports numeric list, steps, ranges or
macros
-handler - the function that will be executed everytime when the job runs
Here is an example:
// main.go
package main
import (
"log"
"github.com/pocketbase/pocketbase"
"github.com/pocketbase/pocketbase/core"
"github.com/pocketbase/pocketbase/tools/cron"
func main() {
app := pocketbase.New()
app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
scheduler := cron.New()
// prints "Hello!" every 2 minutes
scheduler.MustAdd("hello", "*/2 * * * *", func() {
log.Println("Hello!")
scheduler.Start()
return nil
if err := app.Start(); err != nil {
log.Fatal(err)
