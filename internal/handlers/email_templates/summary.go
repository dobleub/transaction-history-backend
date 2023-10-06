package email_templates

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/dobleub/transaction-history-backend/internal/models"
)

func GenerateSummaryEmailBody(summary models.SummaryEmailBody) string {
	tmpl, err := template.ParseFiles("internal/handlers/email_templates/summary.html")
	if err != nil {
		log.Println(err)
		return ""
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, summary)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("%s", tpl.String())
}
