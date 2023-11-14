package helpers

import (
	"fmt"

	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/jung-kurt/gofpdf"
)

const (
	pageWidth                 = 210.0 // A4 width in mm
	pageHeight                = 297.0 // A4 height in mm
	labelWidth                = 210.0 / 2
	labelHeight               = 297.0 / 2
	logoWidth                 = 25.0
	logoHeight                = 25.0
	dataMatrixSize            = 25.0
	tipoServicoSize           = 20.0
	barcodeWidth              = 80.0
	barcodeHeight             = 18.0
	destinatarioBarCodeWidth  = 40.0
	destinatarioBarCodeHeight = 18.0
	defaultLineWidth          = 0.3
)

// ! ===== DIVISOR DO PAPEL A4 EM 4 PARTES IDÊNTICAS =====
func DrawDottedLines(pdf *gofpdf.Fpdf) {
	midX := pageWidth / 2
	midY := pageHeight / 2

	pdf.SetDrawColor(0, 0, 0)
	pdf.SetDashPattern([]float64{3, 2}, 0)

	//! LINHA PONTILHADA VERTICAL
	DrawDottedLine(pdf, midX, 0, midX, pageHeight)

	//! LINHA PONTILHADA HORIZONTAL
	DrawDottedLine(pdf, 0, midY, pageWidth, midY)

	pdf.SetDashPattern([]float64{}, 0)
}
func DrawDottedLine(pdf *gofpdf.Fpdf, x1, y1, x2, y2 float64) {
	pdf.MoveTo(x1, y1)
	pdf.LineTo(x2, y2)
	pdf.DrawPath("D") // "D" for draw (stroke) the path
}
func addImage(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height float64) {
	options := gofpdf.ImageOptions{
		ReadDpi:   true, // Use the DPI embedded in the image file
		ImageType: "",   // An empty string means gofpdf will automatically detect the type
	}

	// The "ImageOptions" function automatically detects the image type based on the file extension
	pdf.ImageOptions(imagePath, x, y, width, height, false, options, 0, "")
}

// ! ===== PRIMEIRA LINHA DA ETIQUETA =====
func DrawFirstRow(pdf *gofpdf.Fpdf, x, y float64, idPlp int) float64 {
	spaceBetween := 12.0 // Space between elements

	// Calculate positions based on the dimensions and space between elements
	tipoServicoX := x
	dataMatrixX := tipoServicoX + tipoServicoSize + spaceBetween
	brandingX := dataMatrixX + dataMatrixSize + spaceBetween

	//! TIPO SERVIÇO LOGO
	//pdf.SetFillColor(255, 0, 0)
	//pdf.Rect(tipoServicoX, y, tipoServicoSize, tipoServicoSize, "F")
	addImage(pdf, "images/sedex-expresso.png", tipoServicoX, y, tipoServicoSize, tipoServicoSize)

	idPlpX := tipoServicoX - 0.7
	idPLpY := y + tipoServicoSize + 0.25
	idPlpText := fmt.Sprintf("PLP: %v", idPlp)
	pdf.SetXY(idPlpX, idPLpY)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, 8, idPlpText, "", 0, "LM", false, 0, "")

	//! DATA MATRIX
	pdf.SetFillColor(0, 0, 0) // Black
	pdf.Rect(dataMatrixX, y, dataMatrixSize, dataMatrixSize, "F")

	//! LOGO FAVOR
	//pdf.SetFillColor(200, 200, 200) // Light gray
	//pdf.Rect(brandingX, y, logoWidth, logoHeight, "F")
	addImage(pdf, "images/favor-logo.png", brandingX, y, logoWidth, logoHeight)
	//pdf.SetFillColor(0, 0, 0)

	nextY := y + dataMatrixSize

	return nextY
}

func DrawSecondRow(pdf *gofpdf.Fpdf, x, y float64, peso float64) float64 {
	spaceBetween := 12.0
	lineHeight := 6.0

	pedidoTextX := x - 0.7 // Remove 0.7 pois os textos possuem um pequeno padding
	pdf.SetXY(pedidoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, "Pedido: 0", "", 0, "L", false, 0, "")

	nfTextX := x + tipoServicoSize + spaceBetween - 0.7
	pdf.SetXY(nfTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, "NF: 0", "", 0, "L", false, 0, "")

	pesoTextX := nfTextX + dataMatrixSize + spaceBetween - 0.7
	pesoText := fmt.Sprintf("Peso (g): %v", peso)
	pdf.SetXY(pesoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, pesoText, "", 0, "L", false, 0, "")

	nextY := y + 4
	return nextY
}

func DrawTrackingCode(pdf *gofpdf.Fpdf, x, y float64, trackingCode string) float64 {
	pdf.SetFont("Arial", "B", 12) // Set the desired font and size

	textWidth := pdf.GetStringWidth(trackingCode)   // Get the width of the tracking code text
	var startX = (labelWidth / 2) - (textWidth / 2) // Calculate the starting X position

	if x == 0.0 {
		startX += labelWidth
	}

	pdf.SetXY(startX, y)                                                  // Move to the starting position
	pdf.CellFormat(textWidth, 10, trackingCode, "", 0, "C", false, 0, "") // Draw the cell with the tracking code centered

	return y + 8
}

func DrawBarcodePlaceholder(pdf *gofpdf.Fpdf, x, y float64) float64 {
	// Calculate the X coordinate to center the barcode
	centerX := x + (labelWidth / 2) - (barcodeWidth / 2)

	// Draw the barcode with 'D' for draw (stroke) the path
	pdf.Rect(centerX, y, barcodeWidth, barcodeHeight, "D")

	return y + barcodeHeight
}

func DrawRecebedorAssinaturaDocumentoLines(pdf *gofpdf.Fpdf, x, y float64) float64 {
	const RECEBEDOR = "Recebedor: "
	const ASSINATURA = "Assinatura: "
	const DOCUMENTO = "Documento: "

	pdf.SetFont("Arial", "", 10)

	lineHeight := 6.0

	//! RECEBEDOR
	recebedorX := x
	recebedorY := y + lineHeight

	recebedorLineXStart := recebedorX + pdf.GetStringWidth(RECEBEDOR)
	recebedorLineXEnd := x + labelWidth - 10

	pdf.Text(recebedorX, recebedorY, RECEBEDOR)
	pdf.Line(recebedorLineXStart, recebedorY, recebedorLineXEnd, recebedorY)

	//! ASSINATURA
	assinaturaX := x
	assinaturaY := recebedorY + lineHeight

	assinaturaLineXStart := assinaturaX + pdf.GetStringWidth(ASSINATURA)
	assinaturaLineXEnd := x + labelWidth/2

	pdf.Text(assinaturaX, assinaturaY, ASSINATURA)
	pdf.Line(assinaturaLineXStart, assinaturaY, assinaturaLineXEnd, assinaturaY)

	//! DOCUMENTO
	documentoX := assinaturaLineXEnd + 1
	documentoY := assinaturaY

	documentoLineXStart := documentoX + pdf.GetStringWidth(DOCUMENTO)
	documentoLineXEnd := x + labelWidth - 10
	pdf.Text(documentoX, documentoY, DOCUMENTO)
	pdf.Line(documentoLineXStart, documentoY, documentoLineXEnd, documentoY)

	nextY := documentoY + 4

	return nextY
}

func DrawDestinatarioCorreiosLogoDivisor(pdf *gofpdf.Fpdf, x, y float64) float64 {
	fmt.Println("DrawDestinatarioDivisor")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	const DESTINATARIO = "DESTINATÁRIO   "
	destinatarioTextWidth := pdf.GetStringWidth(DESTINATARIO) + 10
	lineHeight := 8.0
	fontSize := 12.0

	pdf.SetFont("Arial", "B", fontSize)

	pdf.SetLineWidth(0.5)
	pdf.Line(x, y, x+labelWidth, y)
	pdf.SetLineWidth(defaultLineWidth)

	//! DESENHA O RETANGULO COM FUNDO PRETO
	pdf.SetFillColor(0, 0, 0)
	pdf.Rect(x, y, destinatarioTextWidth, lineHeight, "F")

	destinatarioTextX := x + 1
	destinatarioTextY := y + (lineHeight / 2) + (pdf.PointConvert(fontSize) / 2) - 0.5
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(destinatarioTextX, destinatarioTextY, DESTINATARIO)
	pdf.SetTextColor(0, 0, 0)

	//!TODO: Adicionar o logo dos correios
	widthHeightRatio := 4781.0 / 958.0
	imageWidth := 20.0
	imageHeight := imageWidth / widthHeightRatio
	addImage(pdf, "images/correios-logo.png", x+labelWidth-22, y+1, imageWidth, imageHeight)

	return y + 8.0
}

func DrawDadosDestinatario(pdf *gofpdf.Fpdf, x, y float64, destinatario types.Destinatario, nacional types.Nacional) float64 {
	fontSize := 9.0
	lineHeight := 4.0
	pdf.SetFont("Arial", "", fontSize)

	nomeDestinatarioX := x
	nomeDestinatarioY := y + lineHeight
	nomeDestinatarioText := "Nelson Mendes Jr."
	pdf.Text(nomeDestinatarioX, nomeDestinatarioY, nomeDestinatarioText)

	logradouroDestinatarioX := x
	logradouroDestinatarioY := nomeDestinatarioY + lineHeight
	logradouroDestinatarioText := "Rua João Moreira da Costa, 12"
	pdf.Text(logradouroDestinatarioX, logradouroDestinatarioY, logradouroDestinatarioText)

	complementoBairroDestinatarioX := x
	complementoBairroDestinatarioY := logradouroDestinatarioY + lineHeight
	complementoBairroDestinatarioText := "Docvalle, Vila Resende"
	pdf.Text(complementoBairroDestinatarioX, complementoBairroDestinatarioY, complementoBairroDestinatarioText)

	cepCidadeUfDestinatarioX := x
	cepCidadeUfDestinatarioY := complementoBairroDestinatarioY + lineHeight*1.3
	cepCidadeUfDestinatarioText := "12282-220 Caçapava/SP"
	pdf.Text(cepCidadeUfDestinatarioX, cepCidadeUfDestinatarioY, cepCidadeUfDestinatarioText)

	nextY := cepCidadeUfDestinatarioY + lineHeight

	return nextY
}

func DrawDestinatarioBarCodePlaceholder(pdf *gofpdf.Fpdf, x, y float64) float64 {
	pdf.Rect(x, y, destinatarioBarCodeWidth, destinatarioBarCodeHeight, "D")
	return y + destinatarioBarCodeHeight
}

func DrawSeparadorRemetente(pdf *gofpdf.Fpdf, x, y float64) float64 {
	paddingTop := 6.0
	paddingBottom := 4.0

	pdf.SetLineWidth(0.5)
	pdf.Line(x, y+paddingTop, x+labelWidth, y+paddingTop)
	pdf.SetLineWidth(defaultLineWidth)

	nextY := paddingTop + y + paddingBottom
	return nextY
}

func DrawDadosRemetente(pdf *gofpdf.Fpdf, x, y float64, remetente types.Remetente) float64 {
	fontSize := 8.5
	lineHeight := 3.5
	pdf.SetFont("Arial", "", fontSize)

	nomeRemetenteX := x
	nomeRemetenteY := y
	nomeRemetenteText := "Nelson Mendes Jr."
	pdf.Text(nomeRemetenteX, nomeRemetenteY, nomeRemetenteText)

	logradouroRemetenteX := x
	logradouroRemetenteY := nomeRemetenteY + lineHeight
	logradouroRemetenteText := "Rua João Moreira da Costa, 12"
	pdf.Text(logradouroRemetenteX, logradouroRemetenteY, logradouroRemetenteText)

	complementoBairroRemetenteX := x
	complementoBairroRemetenteY := logradouroRemetenteY + lineHeight
	complementoBairroRemetenteText := "Docvalle, Vila Resende"
	pdf.Text(complementoBairroRemetenteX, complementoBairroRemetenteY, complementoBairroRemetenteText)

	cepCidadeUfRemetenteX := x
	cepCidadeUfRemetenteY := complementoBairroRemetenteY + lineHeight*1.3
	cepCidadeUfRemetenteText := "12282-220 Caçapava/SP"
	pdf.Text(cepCidadeUfRemetenteX, cepCidadeUfRemetenteY, cepCidadeUfRemetenteText)

	nextY := cepCidadeUfRemetenteY + lineHeight

	return nextY
}

// !
func FormatTrackingCode(code string) string {
	if len(code) != 13 {
		return code // Or handle the error according to your needs
	}

	return code[:2] + " " + code[2:5] + " " + code[5:8] + " " + code[8:11] + " " + code[11:]
}
