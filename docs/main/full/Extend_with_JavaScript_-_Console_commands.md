# POCKETBASE DOCS|2026-02-25|1 sections

## 1.Extend with JavaScript - Console commands
# Extend with JavaScript - Console commands
You can register custom console commands using
app.rootCmd.addCommand(cmd), where cmd is a
Command instance.
Here is an example:
$app.rootCmd.addCommand(new Command({
use: "hello",
run: (cmd, args) => {
console.log("Hello world!")
To run the command you can execute:
`./pocketbase hello`
Keep in mind that the console commands execute in their own separate app process and run
independently from the main serve command (aka. hook events between different processes
are not shared with one another).
