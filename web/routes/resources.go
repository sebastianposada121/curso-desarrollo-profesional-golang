package routes

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"
	"web/models"
	"web/utils"

	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
	"github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
)

func Resources(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/resources.html", utils.PATH_TEMPLATE))
	message, status := utils.GetFlashMessage(response, request)
	flash := map[string]string{
		"message": message,
		"status":  status,
	}
	template.Execute(response, flash)
}

/*
Generar pdf
*/
func GeneratePdf(response http.ResponseWriter, request *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	name := strings.Split(time.Now().String(), " ")
	err := pdf.OutputFileAndClose(name[4][6:14] + ".pdf")
	if err != nil {
		panic("Error generate pdf")
	}
	utils.CreateFlashMessage(response, request, "Generate pdf", "success")
	http.Redirect(response, request, "/resources", http.StatusSeeOther)
}

/*
Generar excel
*/
func GenerateExcel(response http.ResponseWriter, request *http.Request) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet1")

	if err != nil {
		fmt.Println(err)
		return
	}

	// construir fila
	f.SetCellValue("Sheet1", "A1", "id")
	f.SetCellValue("Sheet1", "B1", "name")
	f.SetCellValue("Sheet1", "C1", "email")

	for i := 1; i < 10; i++ {
		row := fmt.Sprint(i + 1)
		client := models.Client{
			Id:    i,
			Name:  "user " + fmt.Sprint(i),
			Email: fmt.Sprint(i) + "@email.co",
		}
		f.SetCellValue("Sheet1", "A"+row, client.Id)
		f.SetCellValue("Sheet1", "B"+row, client.Name)
		f.SetCellValue("Sheet1", "C"+row, client.Email)
	}

	f.SetActiveSheet(index)

	// construir documento
	name := strings.Split(time.Now().String(), " ")[4][6:14] + ".xlsx"

	if err := f.SaveAs("assets/excel/" + name); err != nil {
		fmt.Println(err)
	}

	utils.CreateFlashMessage(response, request, "Generate Excel", "success")
	http.Redirect(response, request, "/resources", http.StatusSeeOther)
}

/*
Generar codigo qr apartir de un link
*/
func GenerateQr(response http.ResponseWriter, request *http.Request) {

	template := template.Must(template.ParseFiles("templates/resources.html", utils.PATH_TEMPLATE))

	qr, err := qrcode.Encode("https://www.youtube.com/watch?v=xBw_3DRrzs4&list=OLAK5uy_ku-onqjttX98ePGLE8Xz4zriRvcfQ3uvY&index=2", qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	image := base64.StdEncoding.EncodeToString(qr)

	flash := map[string]string{
		"qr": image,
	}

	utils.CreateFlashMessage(response, request, "Generate Qr", "success")
	template.Execute(response, flash)
}

func GenerateEmail(response http.ResponseWriter, request *http.Request) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "sebaspj67@gmail.com")
	msg.SetHeader("To", "sebaspj67@gmail.com")
	msg.SetHeader("Subject", "Curso golang")
	msg.SetBody("text/html", "<h1>Holiiii</h1>")
	//necesitamos tener una configuracion smtp
	n := gomail.NewDialer("smtp.dreamhost.com", 587, "noreply@agendahoras.cl", "khdwJAXysB")
	err := n.DialAndSend(msg)
	if err != nil {
		panic(err)

	}

	utils.CreateFlashMessage(response, request, "Generate Email", "success")
	http.Redirect(response, request, "/resources", http.StatusSeeOther)
}
