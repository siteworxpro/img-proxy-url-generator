package printer

import "fmt"

func (p *Printer) LogSuccess(message string) {
	fmt.Println(p.getSuccess().Render("✅  " + message))
}

func (p *Printer) LogWarning(message string) {
	fmt.Println(p.getWarning().Render("‼️   WARNING: " + message))
}

func (p *Printer) LogInfo(message string) {
	fmt.Println(p.getInfo().Render("ℹ️   " + message))
}

func (p *Printer) LogError(message string) {
	fmt.Println(p.getError().Render("❌  ERROR: " + message))
}
