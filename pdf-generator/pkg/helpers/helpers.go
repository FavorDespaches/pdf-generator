package helpers

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	PAC_FILEPATH              = "pac.png"
	SEDEX_STANDARD_FILEPATH   = "sedex-standard.png"
	SEDEX_10_FILEPATH         = "sedex-10.png"
	SEDEX_12_FILEPATH         = "sedex-12.png"
	SEDEX_HOJE_FILEPATH       = "sedex-hoje.png"
	MINI_ENVIOS_FILEPATH      = "mini-envios.png"
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
	pdf.DrawPath("D")
}

func addImage(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height float64, keepAspectRatio bool) {
	// Extract the file extension
	extWithDot := filepath.Ext(imagePath)
	if len(extWithDot) <= 1 {
		// Handle the error: file extension is missing or invalid
		log.Fatalf("Invalid file extension for image: %s", imagePath)
		return
	}

	// Extract the file extension
	ext := strings.ToUpper(filepath.Ext(imagePath)[1:])

	// Debugging: Print the imagePath and ImageType
	fmt.Printf("Image path: %s, Image type: %s\n", imagePath, ext)

	options := gofpdf.ImageOptions{
		ReadDpi:   true,
		ImageType: ext,
	}

	if !keepAspectRatio {
		height = 0
	}

	pdf.ImageOptions(imagePath, x, y, width, height, false, options, 0, "")
}

func addBase64ImageToPDF(pdf *gofpdf.Fpdf, base64String string, x, y, w, h float64) error {
	// Decode the base64 string
	imgData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "*.png")
	if err != nil {
		return err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name()) // Clean up

	// Write the decoded image to the temporary file
	_, err = tempFile.Write(imgData)
	if err != nil {
		return err
	}

	// Add the image to the PDF
	pdf.Image(tempFile.Name(), x, y, w, h, false, "", 0, "")

	return nil
}

func findTipoServicoImagemByCodServicoPostagem(codServicoPostagem string) string {
	fmt.Println("codServicoPostagem:", codServicoPostagem)
	var tipoServicoImagem string
	switch codServicoPostagem {
	case "03298", "3298":
		tipoServicoImagem = PAC_FILEPATH
	case "03220", "3220":
		tipoServicoImagem = SEDEX_STANDARD_FILEPATH
	case "04227", "4227":
		tipoServicoImagem = MINI_ENVIOS_FILEPATH
	case "03158", "3158":
		tipoServicoImagem = SEDEX_10_FILEPATH
	case "03140", "3140":
		tipoServicoImagem = SEDEX_12_FILEPATH
	case "03204", "3204":
		tipoServicoImagem = SEDEX_HOJE_FILEPATH
	default:
		panic("Código não implementado!")
	}
	return tipoServicoImagem
}

// ! ===== PRIMEIRA LINHA DA ETIQUETA =====
func DrawFirstRow(pdf *gofpdf.Fpdf, x, y float64, idPlp int, tipoServicoImagem string, dataMatrixBase64String string, local bool) float64 {
	spaceBetween := 12.0 // Space between elements

	// Calculate positions based on the dimensions and space between elements
	tipoServicoX := x
	dataMatrixX := tipoServicoX + tipoServicoSize + spaceBetween
	brandingX := dataMatrixX + dataMatrixSize + spaceBetween

	//! TIPO SERVIÇO LOGO
	var tipoServicoImagemPath string
	if local {
		tipoServicoImagemPath = filepath.Join("../../layers/images", tipoServicoImagem)
	} else {
		tipoServicoImagemPath = filepath.Join("/opt", "bin", "images", tipoServicoImagem)
	}

	keepAspectRatio := false
	ratio := 1.4
	if tipoServicoImagem == MINI_ENVIOS_FILEPATH {
		keepAspectRatio = true
		ratio = 1.0
	}

	addImage(pdf, tipoServicoImagemPath, tipoServicoX, y, ratio*tipoServicoSize, tipoServicoSize, keepAspectRatio)

	idPlpX := tipoServicoX - 0.7
	idPLpY := y + tipoServicoSize + 0.25
	idPlpText := fmt.Sprintf("PLP: %v", idPlp)
	pdf.SetXY(idPlpX, idPLpY)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, 8, idPlpText, "", 0, "LM", false, 0, "")

	//! DATA MATRIX
	// errDataMatrix := generateDataMatrix(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	errDataMatrix := addBase64ImageToPDF(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	if errDataMatrix != nil {
		errDataMatrixString := fmt.Sprintf("Erro generateDataMatrix %s", errDataMatrix.Error())
		panic(errDataMatrixString)
	}

	//! LOGO FAVOR
	var favorLogoImagePath string
	if local {
		favorLogoImagePath = filepath.Join("../../layers/images", "favor-logo.png")
	} else {
		favorLogoImagePath = filepath.Join("/opt", "bin", "images", "favor-logo.png")
	}
	addImage(pdf, favorLogoImagePath, brandingX, y, logoWidth, logoHeight, false)

	nextY := y + dataMatrixSize

	return nextY
}

// ! ===== PRIMEIRA LINHA DA ETIQUETA =====
func DrawSmallFirstRow(pdf *gofpdf.Fpdf, x, y float64, idPlp int, tipoServicoImagem string, dataMatrixBase64String string, local bool) float64 {
	spaceBetween := 12.0 // Space between elements

	// Calculate positions based on the dimensions and space between elements
	tipoServicoX := x

	//! TIPO SERVIÇO LOGO
	var tipoServicoImagemPath string
	if local {
		tipoServicoImagemPath = filepath.Join("../../layers/images", tipoServicoImagem)
	} else {
		tipoServicoImagemPath = filepath.Join("/opt", "bin", "images", tipoServicoImagem)
	}

	keepAspectRatio := false
	ratio := 1.4
	if tipoServicoImagem == MINI_ENVIOS_FILEPATH {
		keepAspectRatio = true
		ratio = 1.0
	}
	addImage(pdf, tipoServicoImagemPath, tipoServicoX, y, ratio*tipoServicoSize, tipoServicoSize, keepAspectRatio)

	idPlpX := tipoServicoX - 0.7
	idPLpY := y + tipoServicoSize + 0.25
	idPlpText := fmt.Sprintf("PLP: %v", idPlp)
	pdf.SetXY(idPlpX, idPLpY)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, 8, idPlpText, "", 0, "LM", false, 0, "")

	//! DATA MATRIX
	// errDataMatrix := generateDataMatrix(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	dataMatrixX := tipoServicoX + tipoServicoSize + 1.3*spaceBetween
	errDataMatrix := addBase64ImageToPDF(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	if errDataMatrix != nil {
		errDataMatrixString := fmt.Sprintf("Erro generateDataMatrix %s", errDataMatrix.Error())
		panic(errDataMatrixString)
	}

	//! LOGO FAVOR
	var favorLogoImagePath string
	brandingX := dataMatrixX + dataMatrixSize + 0.7*spaceBetween
	if local {
		favorLogoImagePath = filepath.Join("../../layers/images", "favor-logo.png")
	} else {
		favorLogoImagePath = filepath.Join("/opt", "bin", "images", "favor-logo.png")
	}
	addImage(pdf, favorLogoImagePath, brandingX, y, logoWidth, logoHeight, false)

	nextY := y + dataMatrixSize

	return nextY
}

func DrawSmallDelimiter(pdf *gofpdf.Fpdf, x, y float64) {
	pdf.SetDrawColor(0, 0, 0)
	//! LINHA HORIZONTAL
	DrawDottedLine(pdf, x+3.5, y+4.75, x+3.5+98, y+4.75)
	DrawDottedLine(pdf, x+3.5+98, y+4.75, x+3.5+98, y+4.75+139)
	DrawDottedLine(pdf, x+3.5+98, y+4.75+139, x+3.5, y+4.75+139)
	DrawDottedLine(pdf, x+3.5, y+4.75+139, x+3.5, y+4.75)
}

//-----------------------------------------------------------------

func DrawSecondRow(pdf *gofpdf.Fpdf, x, y float64, peso float64) float64 {
	spaceBetween := 12.0
	lineHeight := 6.0

	pedidoTextX := x - 0.7
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

func DrawSmallSecondRow(pdf *gofpdf.Fpdf, x, y float64, peso float64) float64 {
	spaceBetween := 12.0
	lineHeight := 6.0

	pedidoTextX := x - 0.7
	pdf.SetXY(pedidoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, "Pedido: 0", "", 0, "L", false, 0, "")

	nfTextX := x + tipoServicoSize + 1.3*spaceBetween - 0.7
	pdf.SetXY(nfTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, "NF: 0", "", 0, "L", false, 0, "")

	pesoTextX := nfTextX + dataMatrixSize + 0.7*spaceBetween - 0.7
	pesoText := fmt.Sprintf("Peso (g): %v", peso)
	pdf.SetXY(pesoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, pesoText, "", 0, "L", false, 0, "")

	nextY := y + 4
	return nextY
}

//-----------------------------------------------------------------

// ! ===== CÓDIGO DE RASTREIO =====
func DrawTrackingCode(pdf *gofpdf.Fpdf, x, y float64, trackingCode string) float64 {
	pdf.SetFont("Arial", "B", 12)

	textWidth := pdf.GetStringWidth(trackingCode)
	var startX = (labelWidth / 2) - (textWidth / 2)

	if x != 0.0 {
		startX += labelWidth
	}

	pdf.SetXY(startX, y)
	pdf.CellFormat(textWidth, 10, trackingCode, "", 0, "C", false, 0, "")

	return y + 8
}

// ! ========== BARRA DE CÓDIGO MAIOR ==========
func DrawBarcode(pdf *gofpdf.Fpdf, x, y float64, barcodeBase64String string) float64 {
	centerX := x + (labelWidth / 2) - (barcodeWidth / 2)
	// generateBarcode128(pdf, barcodeBase64String, centerX, y, barcodeWidth, barcodeHeight)
	errBarcode := addBase64ImageToPDF(pdf, barcodeBase64String, centerX, y, barcodeWidth, barcodeHeight)

	if errBarcode != nil {
		errBarcodeString := fmt.Sprintf("Erro DrawBarcode addBase64ImageToPDF %s", errBarcode.Error())
		panic(errBarcodeString)
	}

	return y + barcodeHeight
}

// ! ========== ASSINATURAS ==========
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

func DrawSmallRecebedorAssinaturaDocumentoLines(pdf *gofpdf.Fpdf, x, y float64) float64 {
	const RECEBEDOR = "Recebedor: "
	const ASSINATURA = "Assinatura: "
	const DOCUMENTO = "Documento: "

	pdf.SetFont("Arial", "", 8)

	lineHeight := 5.0

	//! RECEBEDOR
	recebedorX := x
	recebedorY := y + lineHeight

	recebedorLineXStart := recebedorX + pdf.GetStringWidth(RECEBEDOR)
	recebedorLineXEnd := x + labelWidth - 10 - 4

	pdf.Text(recebedorX, recebedorY, RECEBEDOR)
	pdf.Line(recebedorLineXStart, recebedorY, recebedorLineXEnd, recebedorY)

	//! ASSINATURA
	assinaturaX := x
	assinaturaY := recebedorY + lineHeight

	assinaturaLineXStart := assinaturaX + pdf.GetStringWidth(ASSINATURA)
	assinaturaLineXEnd := x + labelWidth/2 - 6

	pdf.Text(assinaturaX, assinaturaY, ASSINATURA)
	pdf.Line(assinaturaLineXStart, assinaturaY, assinaturaLineXEnd, assinaturaY)

	//! DOCUMENTO
	documentoX := assinaturaLineXEnd + 1
	documentoY := assinaturaY

	documentoLineXStart := documentoX + pdf.GetStringWidth(DOCUMENTO)
	documentoLineXEnd := x + labelWidth - 14
	pdf.Text(documentoX, documentoY, DOCUMENTO)
	pdf.Line(documentoLineXStart, documentoY, documentoLineXEnd, documentoY)

	nextY := documentoY + 4

	return nextY
}

//-----------------------------------------------------------------

// ! ========== DIVISOR DESTINATÁRIO ==========
func DrawDestinatarioCorreiosLogoDivisor(pdf *gofpdf.Fpdf, x, y float64, local bool) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
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
	pdf.Text(destinatarioTextX, destinatarioTextY, translator(DESTINATARIO))
	pdf.SetTextColor(0, 0, 0)

	//!TODO: Adicionar o logo dos correios
	widthHeightRatio := 4781.0 / 958.0
	imageWidth := 20.0
	imageHeight := imageWidth / widthHeightRatio

	var correiosLogoImagePath string
	if local {
		correiosLogoImagePath = filepath.Join("../../layers/images", "correios-logo.png")
	} else {
		correiosLogoImagePath = filepath.Join("/opt", "bin", "images", "correios-logo.png")
	}
	addImage(pdf, correiosLogoImagePath, x+labelWidth-22, y+1, imageWidth, imageHeight, false)

	return y + 8.0
}

func DrawSmallDestinatarioCorreiosLogoDivisor(pdf *gofpdf.Fpdf, x, y float64, local bool) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	const DESTINATARIO = "DESTINATÁRIO   "
	destinatarioTextWidth := pdf.GetStringWidth(DESTINATARIO) + 10
	lineHeight := 8.0
	fontSize := 12.0

	pdf.SetFont("Arial", "B", fontSize)

	pdf.SetLineWidth(0.5)
	pdf.Line(x, y, x+labelWidth-7, y)
	pdf.SetLineWidth(defaultLineWidth)

	//! DESENHA O RETANGULO COM FUNDO PRETO
	pdf.SetFillColor(0, 0, 0)
	pdf.Rect(x, y, destinatarioTextWidth, lineHeight, "F")

	destinatarioTextX := x + 1
	destinatarioTextY := y + (lineHeight / 2) + (pdf.PointConvert(fontSize) / 2) - 0.5
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(destinatarioTextX, destinatarioTextY, translator(DESTINATARIO))
	pdf.SetTextColor(0, 0, 0)

	widthHeightRatio := 4781.0 / 958.0
	imageWidth := 20.0
	imageHeight := imageWidth / widthHeightRatio

	var correiosLogoImagePath string
	if local {
		correiosLogoImagePath = filepath.Join("../../layers/images", "correios-logo.png")
	} else {
		correiosLogoImagePath = filepath.Join("/opt", "bin", "images", "correios-logo.png")
	}
	addImage(pdf, correiosLogoImagePath, x+labelWidth-22-7, y+1, imageWidth, imageHeight, false)

	return y + 8.0
}

//-----------------------------------------------------------------

func buildLogradouroDestinatarioString(destinatario types.Destinatario) string {
	var hasNumerodestinatario = destinatario.NumeroEndDestinatario != ""

	var logradouroDestinatarioString string

	logradouroDestinatarioString += destinatario.LogradouroDestinatario

	if hasNumerodestinatario {
		logradouroDestinatarioString += ", "
		logradouroDestinatarioString += destinatario.NumeroEndDestinatario
	}

	return logradouroDestinatarioString
}

func buildComplementoBairroDestinatarioString(destinatario types.Destinatario, nacional types.Nacional) string {
	var hasComplemento = destinatario.ComplementoDestinatario != ""
	var hasBairro = nacional.BairroDestinatario != ""

	var complementoBairroDestinatarioString string

	if hasComplemento {
		complementoBairroDestinatarioString += destinatario.ComplementoDestinatario
	}
	if hasComplemento && hasBairro {
		complementoBairroDestinatarioString += ", "
	}
	if hasBairro {
		complementoBairroDestinatarioString += nacional.BairroDestinatario
	}

	return complementoBairroDestinatarioString
}

func buildCepDestinatarioString(nacional types.Nacional) string {
	var formattedCEP string
	cep := nacional.CepDestinatario
	lenCEP := len(nacional.CepDestinatario)

	if lenCEP != 8 {
		if lenCEP == 7 {
			formattedCEP += "0"
			formattedCEP += nacional.CepDestinatario
		} else {
			panic("CEP INVÁLIDO")
		}
	} else {
		formattedCEP = cep[:5] + "-" + cep[5:]
	}

	return formattedCEP
}

func buildCidadeUfDestinatarioString(nacional types.Nacional) string {
	cidadeUfDestinatarioString := fmt.Sprintf("%s / %s", nacional.CidadeDestinatario, nacional.UfDestinatario)
	return cidadeUfDestinatarioString
}

// ! ========== DADOS DO DESTINATÁRIO ==========
func DrawDadosDestinatario(pdf *gofpdf.Fpdf, x, y float64, destinatario types.Destinatario, nacional types.Nacional) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	fontSize := 9.0
	lineHeight := 4.0
	pdf.SetFont("Arial", "", fontSize)

	nomeDestinatarioX := x
	nomeDestinatarioY := y + lineHeight
	nomeDestinatarioText := translator(destinatario.NomeDestinatario)
	pdf.Text(nomeDestinatarioX, nomeDestinatarioY, nomeDestinatarioText)

	logradouroDestinatarioX := x
	logradouroDestinatarioY := nomeDestinatarioY + lineHeight
	logradouroDestinatarioString := buildLogradouroDestinatarioString(destinatario)
	logradouroDestinatarioText := translator(logradouroDestinatarioString)
	pdf.Text(logradouroDestinatarioX, logradouroDestinatarioY, logradouroDestinatarioText)

	complementoBairroDestinatarioX := x
	complementoBairroDestinatarioY := logradouroDestinatarioY + lineHeight
	complementoBairroDestinatarioString := buildComplementoBairroDestinatarioString(destinatario, nacional)
	complementoBairroDestinatarioText := translator(complementoBairroDestinatarioString)
	pdf.Text(complementoBairroDestinatarioX, complementoBairroDestinatarioY, complementoBairroDestinatarioText)

	cepCidadeUfDestinatarioY := complementoBairroDestinatarioY + lineHeight*1.3

	pdf.SetFont("Arial", "B", fontSize)
	cepDestinatarioX := x
	cepDestinatarioString := buildCepDestinatarioString(nacional)
	cepDestinatarioText := translator(cepDestinatarioString)
	pdf.Text(cepDestinatarioX, cepCidadeUfDestinatarioY, cepDestinatarioText)

	pdf.SetFont("Arial", "", fontSize)
	cidadeUfDestinatarioPaddingLeft := 1.5
	cidadeUfDestinatarioX := x + pdf.GetStringWidth(cepDestinatarioString) + cidadeUfDestinatarioPaddingLeft
	cidadeUfDestinatarioString := buildCidadeUfDestinatarioString(nacional)
	cidadeUfDestinatarioText := translator(cidadeUfDestinatarioString)
	pdf.Text(cidadeUfDestinatarioX, cepCidadeUfDestinatarioY, cidadeUfDestinatarioText)

	nextY := cepCidadeUfDestinatarioY + lineHeight

	return nextY
}

// ! ========== BARRA DE CÓDIGO DESTINATÁRIO ==========
func DrawDestinatarioBarCode(pdf *gofpdf.Fpdf, x, y float64, destinatarioBarcodeBase64String string) float64 {
	// generateBarcode128(pdf, destinatarioBarcodeBase64String, x, y, destinatarioBarCodeWidth, destinatarioBarCodeHeight)
	errDestinatarioBarCode := addBase64ImageToPDF(pdf, destinatarioBarcodeBase64String, x, y, destinatarioBarCodeWidth, destinatarioBarCodeHeight)

	if errDestinatarioBarCode != nil {
		errDestinatarioBarCodeString := fmt.Sprintf("Erro DrawDestinatarioBarCode generateBarcode128 %s", errDestinatarioBarCode.Error())
		panic(errDestinatarioBarCodeString)
	}

	return y + destinatarioBarCodeHeight
}

// ! ========== SEPARADOR REMETENTE ==========
func DrawSeparadorRemetente(pdf *gofpdf.Fpdf, x, y float64) float64 {
	paddingTop := 6.0
	paddingBottom := 4.0

	pdf.SetLineWidth(0.5)
	pdf.Line(x, y+paddingTop, x+labelWidth, y+paddingTop)
	pdf.SetLineWidth(defaultLineWidth)

	nextY := paddingTop + y + paddingBottom
	return nextY
}

// ! ========== SEPARADOR REMETENTE ==========
func DrawSmallSeparadorRemetente(pdf *gofpdf.Fpdf, x, y float64) float64 {
	paddingTop := 2.0
	paddingBottom := 4.0

	pdf.SetLineWidth(0.5)
	pdf.Line(x+3.5, y+paddingTop, x+labelWidth-3.5, y+paddingTop)
	pdf.SetLineWidth(defaultLineWidth)

	nextY := paddingTop + y + paddingBottom
	return nextY
}

//----------------------------------------------------

func buildLogradouroRemetenteString(remetente types.Remetente) string {
	var hasNumeroRemetente = remetente.NumeroRemetente != ""

	var logradouroRemetenteString string

	logradouroRemetenteString += remetente.LogradouroRemetente

	if hasNumeroRemetente {
		logradouroRemetenteString += ", "
		logradouroRemetenteString += remetente.NumeroRemetente
	}

	return logradouroRemetenteString
}

func buildComplementoBairroRemetenteString(remetente types.Remetente) string {
	var hasComplemento = remetente.ComplementoRemetente != ""
	var hasBairro = remetente.BairroRemetente != ""

	var complementoBairroRemetenteString string

	if hasComplemento {
		complementoBairroRemetenteString += remetente.ComplementoRemetente
	}
	if hasComplemento && hasBairro {
		complementoBairroRemetenteString += ", "
	}
	if hasBairro {
		complementoBairroRemetenteString += remetente.BairroRemetente
	}

	return complementoBairroRemetenteString
}

func buildCepRemetenteString(remetente types.Remetente) string {
	var formattedCEP string
	cep := remetente.CepRemetente
	lenCEP := len(remetente.CepRemetente)

	if lenCEP != 8 {
		if lenCEP == 7 {
			formattedCEP += "0"
			formattedCEP += remetente.CepRemetente
		} else {
			panic("CEP INVÁLIDO")
		}
	} else {
		formattedCEP = cep[:5] + "-" + cep[5:]
	}
	return formattedCEP
}

func buildCidadeUfRemetenteString(remetente types.Remetente) string {
	cidadeUfRemetenteString := fmt.Sprintf("%s / %s", remetente.CidadeRemetente, remetente.UfRemetente)
	return cidadeUfRemetenteString
}

// ! ========== DADOS DO REMETENTE ==========
func DrawDadosRemetente(pdf *gofpdf.Fpdf, x, y float64, remetente types.Remetente) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	fontSize := 8.5
	lineHeight := 3.5
	pdf.SetFont("Arial", "", fontSize)

	nomeRemetenteY := y

	pdf.SetFont("Arial", "B", fontSize)
	remetenteX := x
	remetenteText := translator("Remetente: ")
	pdf.Text(remetenteX, nomeRemetenteY, remetenteText)

	pdf.SetFont("Arial", "", fontSize)
	nomeRemetentePaddingLeft := 1.0
	nomeRemetenteX := x + pdf.GetStringWidth("Remetente: ") + nomeRemetentePaddingLeft
	nomeRemetenteText := translator(remetente.NomeRemetente)
	pdf.Text(nomeRemetenteX, nomeRemetenteY, nomeRemetenteText)

	logradouroRemetenteX := x
	logradouroRemetenteY := nomeRemetenteY + lineHeight
	logradouroRemetenteString := buildLogradouroRemetenteString(remetente)
	logradouroRemetenteText := translator(logradouroRemetenteString)
	pdf.Text(logradouroRemetenteX, logradouroRemetenteY, logradouroRemetenteText)

	complementoBairroRemetenteX := x
	complementoBairroRemetenteY := logradouroRemetenteY + lineHeight
	complementoBairroRemetenteString := buildComplementoBairroRemetenteString(remetente)
	complementoBairroRemetenteText := translator(complementoBairroRemetenteString)
	pdf.Text(complementoBairroRemetenteX, complementoBairroRemetenteY, complementoBairroRemetenteText)

	cepCidadeUfRemetenteY := complementoBairroRemetenteY + lineHeight*1.3

	pdf.SetFont("Arial", "B", fontSize)
	cepRemetenteX := x
	cepRemetenteString := buildCepRemetenteString(remetente)
	cepRemetenteText := translator(cepRemetenteString)
	pdf.Text(cepRemetenteX, cepCidadeUfRemetenteY, cepRemetenteText)

	pdf.SetFont("Arial", "", fontSize)
	cidadeUfRemetentePaddingLeft := 1.5
	cidadeUfRemetenteX := x + pdf.GetStringWidth(cepRemetenteString) + cidadeUfRemetentePaddingLeft
	cidadeUfRemetenteString := buildCidadeUfRemetenteString(remetente)
	cidadeUfRemetenteText := translator(cidadeUfRemetenteString)
	pdf.Text(cidadeUfRemetenteX, cepCidadeUfRemetenteY, cidadeUfRemetenteText)

	nextY := cepCidadeUfRemetenteY + lineHeight

	return nextY
}

// ! ========== FORMATADOR DO CÓDIGO DE RASTREIO ==========
func FormatTrackingCode(code string) string {
	if len(code) != 13 {
		return code
	}

	return code[:2] + " " + code[2:5] + " " + code[5:8] + " " + code[8:11] + " " + code[11:]
}
