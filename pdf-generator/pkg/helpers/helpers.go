package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/jung-kurt/gofpdf"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/datamatrix"
)

const (
	pageWidth                 = 210.0 // A4 width in mm
	pageHeight                = 297.0 // A4 height in mm
	labelWidth                = 98
	labelHeight               = 138
	cartaWidth                = 97.0
	cartaHeight               = 36.0
	DPI                       = 300
	paddingRight              = 12.0
	logoWidth                 = 25.0
	logoHeight                = 25.0
	dataMatrixSize            = 25.0
	tipoServicoSize           = 20.0
	barcodeWidth              = 80.0
	barcodeHeight             = 18.0
	dadosDestinatarioHeight   = 42.4
	destinatarioBarCodeWidth  = 40.0
	destinatarioBarCodeHeight = 18.0
	paddingDestinatario       = 3.0
	defaultLineWidth          = 0.3
	PAC_FILEPATH              = "pac.png"
	SEDEX_STANDARD_FILEPATH   = "sedex-standard.png"
	SEDEX_10_FILEPATH         = "sedex-10.png"
	SEDEX_12_FILEPATH         = "sedex-12.png"
	SEDEX_HOJE_FILEPATH       = "sedex-hoje.png"
	MINI_ENVIOS_FILEPATH      = "mini-envios.png"
	CARTA_SIMPLES_FILEPATH    = "carta-simples.png"
)

func DrawDelimiter(pdf *gofpdf.Fpdf, x, y float64) {
	labelTopY := y + 8
	labelBottomY := labelTopY + labelHeight
	labelLeftX := x + 5
	labelRightX := labelLeftX + labelWidth

	if x == pageWidth/2 {
		labelLeftX = x + 2.5
		labelRightX = labelLeftX + labelWidth
	}
	if y == pageHeight/2 {
		labelTopY = y + 2
		labelBottomY = labelTopY + labelHeight
	}

	pdf.Line(labelLeftX, labelTopY, labelRightX, labelTopY)
	pdf.Line(labelRightX, labelTopY, labelRightX, labelBottomY)
	pdf.Line(labelRightX, labelBottomY, labelLeftX, labelBottomY)
	pdf.Line(labelLeftX, labelBottomY, labelLeftX, labelTopY)
}

func addImage(pdf *gofpdf.Fpdf, imagePath string, x, y, width, height float64, keepAspectRatio bool) error {
	// Check if file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fmt.Errorf("image file not found: %s", imagePath)
	}

	var imageOption gofpdf.ImageOptions
	if keepAspectRatio {
		imageOption = gofpdf.ImageOptions{ImageType: "", ReadDpi: false, AllowNegativePosition: false}
	} else {
		imageOption = gofpdf.ImageOptions{ImageType: "", ReadDpi: true, AllowNegativePosition: false}
	}

	pdf.ImageOptions(imagePath, x, y, width, height, false, imageOption, 0, "")
	return nil
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
	// fmt.Println("codServicoPostagem:", codServicoPostagem)
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
	case "80160":
		tipoServicoImagem = CARTA_SIMPLES_FILEPATH
	default:
		log.Fatalf("CÓDIGO %s NÃO IMPLEMENTADO", codServicoPostagem)
		panic("Código não implementado!")
	}
	return tipoServicoImagem
}

// ! ===== PRIMEIRA LINHA DA ETIQUETA =====
func DrawFirstRow(pdf *gofpdf.Fpdf, x, y float64, idPrePostagem string, tipoServicoImagem string, dataMatrixBase64String string, local bool) float64 {
	spaceBetween := 12.0 // Space between elements

	// Calculate positions based on the dimensions and space between elements
	tipoServicoX := x

	//! TIPO SERVIÇO LOGO
	var tipoServicoImagemPath string
	if local {
		// Get workspace root directory (go up one level from current directory)
		workspaceRoot := filepath.Dir(GetWorkingDir())

		// Try direct path first (assuming we're in pdf-generator/pkg/helpers)
		tipoServicoImagemPath = filepath.Join(workspaceRoot, "layers/images", tipoServicoImagem)
		log.Printf("Trying image path: %s", tipoServicoImagemPath)

		// Check if file exists
		if _, err := os.Stat(tipoServicoImagemPath); os.IsNotExist(err) {
			log.Printf("WARNING: Image file not found at %s", tipoServicoImagemPath)

			// Try alternate path formats
			alternativePaths := []string{
				filepath.Join(workspaceRoot, "layers", "images", tipoServicoImagem),
				filepath.Join(GetWorkingDir(), "..", "..", "layers", "images", tipoServicoImagem),
				filepath.Join(GetWorkingDir(), "..", "layers", "images", tipoServicoImagem),
				filepath.Join("/home/gabriel/Favor/Github/pdf-generator/layers/images", tipoServicoImagem),
			}

			for _, altPath := range alternativePaths {
				log.Printf("Trying alternative path: %s", altPath)
				if _, err := os.Stat(altPath); err == nil {
					log.Printf("Found image at path: %s", altPath)
					tipoServicoImagemPath = altPath
					break
				}
			}
		}
	} else {
		tipoServicoImagemPath = filepath.Join("/opt", "bin", "images", tipoServicoImagem)
	}

	log.Printf("Loading image from: %s", tipoServicoImagemPath)

	keepAspectRatio := false
	ratio := 1.4
	if tipoServicoImagem == MINI_ENVIOS_FILEPATH {
		keepAspectRatio = true
		ratio = 1.0
	}

	imgErr := addImage(pdf, tipoServicoImagemPath, tipoServicoX, y, ratio*tipoServicoSize, tipoServicoSize, keepAspectRatio)
	if imgErr != nil {
		log.Printf("Error adding service type image: %v", imgErr)
	}

	idPlpX := tipoServicoX - 0.7
	idPLpY := y + tipoServicoSize + 0.25
	// idPlpText := fmt.Sprintf("ID: %s", idPrePostagem)
	pdf.SetXY(idPlpX, idPLpY)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, 8, "suporte:", "", 0, "LM", false, 0, "")

	//! DATA MATRIX
	// errDataMatrix := generateDataMatrix(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	dataMatrixX := tipoServicoX + tipoServicoSize + 1.3*spaceBetween
	errDataMatrix := addBase64ImageToPDF(pdf, dataMatrixBase64String, dataMatrixX, y, dataMatrixSize, dataMatrixSize)
	if errDataMatrix != nil {
		errDataMatrixString := fmt.Sprintf("Erro generateDataMatrix %s", errDataMatrix.Error())
		log.Fatalf(errDataMatrixString)
		panic(errDataMatrixString)
	}

	//! LOGO FAVOR
	var favorLogoImagePath string
	brandingX := dataMatrixX + dataMatrixSize + 0.7*spaceBetween
	if local {
		// Get workspace root directory (go up one level from current directory)
		workspaceRoot := filepath.Dir(GetWorkingDir())

		// Try direct path first
		favorLogoImagePath = filepath.Join(workspaceRoot, "layers/images", "favor-logo.png")
		log.Printf("Trying favor logo path: %s", favorLogoImagePath)

		// Check if file exists
		if _, err := os.Stat(favorLogoImagePath); os.IsNotExist(err) {
			log.Printf("WARNING: Favor logo image file not found at %s", favorLogoImagePath)

			// Try alternate path formats
			alternativePaths := []string{
				filepath.Join(workspaceRoot, "layers", "images", "favor-logo.png"),
				filepath.Join(GetWorkingDir(), "..", "..", "layers", "images", "favor-logo.png"),
				filepath.Join(GetWorkingDir(), "..", "layers", "images", "favor-logo.png"),
				filepath.Join("/home/gabriel/Favor/Github/pdf-generator/layers/images", "favor-logo.png"),
			}

			for _, altPath := range alternativePaths {
				log.Printf("Trying alternative path for favor logo: %s", altPath)
				if _, err := os.Stat(altPath); err == nil {
					log.Printf("Found favor logo at path: %s", altPath)
					favorLogoImagePath = altPath
					break
				}
			}
		}
	} else {
		favorLogoImagePath = filepath.Join("/opt", "bin", "images", "favor-logo.png")
	}

	log.Printf("Loading favor logo from: %s", favorLogoImagePath)

	logoErr := addImage(pdf, favorLogoImagePath, brandingX, y, logoWidth, logoHeight, false)
	if logoErr != nil {
		log.Printf("Error adding favor logo image: %v", logoErr)
	}

	nextY := y + dataMatrixSize

	return nextY
}

//-----------------------------------------------------------------

func DrawSecondRow(pdf *gofpdf.Fpdf, x, y float64, idPrePostagem string, peso float64) float64 {
	spaceBetween := 12.0
	lineHeight := 6.0

	pedidoTextX := x - 0.7
	// pedidoText := fmt.Sprintf("Id: %s", idPrePostagem)
	pdf.SetXY(pedidoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, "favordespaches.com.br/suporte", "", 0, "L", false, 0, "")

	nfTextX := x + tipoServicoSize + 1.3*spaceBetween - 0.7
	pdf.SetXY(nfTextX, y)
	pdf.SetFont("Arial", "", 8)
	// pdf.CellFormat(tipoServicoSize, lineHeight, "favordespaches.com.br/suporte", "", 0, "L", false, 0, "")

	pesoTextX := nfTextX + dataMatrixSize + 0.7*spaceBetween - 0.7
	pesoText := fmt.Sprintf("Peso (g): %.0f", peso)
	pdf.SetXY(pesoTextX, y)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(tipoServicoSize, lineHeight, pesoText, "", 0, "L", false, 0, "")

	nextY := y + 2.0
	return nextY
}

// -----------------------------------------------------------------
func mmToPixels(mm int, dpi int) int {
	return int(float64(mm*dpi)/25.4 + 0.5) // Adding 0.5 for rounding to nearest integer
}

func ConvertImage(img image.Image) *image.RGBA {
	b := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			newImg.Set(x, y, img.At(x, y))
		}
	}
	return newImg
}

// ! ===== CÓDIGO DE RASTREIO =====
func CreateBarcodeBaseString(width, height int, code string) string {
	baseCode, err := code128.Encode(code)
	if err != nil {
		log.Fatalf("code128.Encode(code): %v", err)
	}

	scaledCode, err := barcode.Scale(baseCode, mmToPixels(width, DPI), mmToPixels(height, DPI))
	if err != nil {
		log.Fatalf("barcode.Scale(): %v", err)
	}

	//fmt.Printf("%s", scaledCode.Metadata())
	var buf bytes.Buffer
	err = png.Encode(&buf, ConvertImage(scaledCode))
	if err != nil {
		log.Fatalf("png.Encode(&buf, ConvertImage(scaledCode)): %v", err)
	}

	// fmt.Printf("Encode: %s", buf.Bytes())

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func CreateDatamatrixBaseString(code string, w, h int) string {
	baseCode, _ := datamatrix.Encode(code)
	scaledCode, _ := barcode.Scale(baseCode, mmToPixels(w, DPI), mmToPixels(h, DPI))

	var buf bytes.Buffer
	err := png.Encode(&buf, ConvertImage(scaledCode))
	if err != nil {
		log.Fatalf("Failed to encode barcode as PNG: %v", err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func DrawTrackingCode(pdf *gofpdf.Fpdf, x, y float64, trackingCode string, isRightColumn bool) float64 {
	pdf.SetFont("Arial", "B", 12)

	textWidth := pdf.GetStringWidth(trackingCode)
	var startX = (labelWidth / 2) - (textWidth / 2)

	if isRightColumn {
		startX += labelWidth + 6.5
	}

	pdf.SetXY(startX, y)
	pdf.CellFormat(textWidth, 10, trackingCode, "", 0, "C", false, 0, "")

	return y + 8.0
}

// ! ========== BARRA DE CÓDIGO MAIOR ==========
func DrawBarcode(pdf *gofpdf.Fpdf, x, y float64, barcodeBase64String string) float64 {
	centerX := x + (labelWidth / 2) - (barcodeWidth / 2)
	// generateBarcode128(pdf, barcodeBase64String, centerX, y, barcodeWidth, barcodeHeight)
	errBarcode := addBase64ImageToPDF(pdf, barcodeBase64String, centerX, y, barcodeWidth, barcodeHeight)

	if errBarcode != nil {
		errBarcodeString := fmt.Sprintf("Erro DrawBarcode addBase64ImageToPDF %s", errBarcode.Error())
		log.Fatalf(errBarcodeString)
		panic(errBarcodeString)
	}

	barcodePaddingBottom := 5.0
	return y + barcodeHeight + barcodePaddingBottom
}

// ! ========== ASSINATURAS ==========
func DrawRecebedorAssinaturaDocumentoLines(pdf *gofpdf.Fpdf, x, y float64) float64 {
	const RECEBEDOR = "Recebedor: "
	const ASSINATURA = "Assinatura: "
	const DOCUMENTO = "Documento: "

	pdf.SetFont("Arial", "", 8)

	lineHeight := 5.0

	//! RECEBEDOR
	recebedorX := x
	recebedorY := y

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

	recebedorAssinaturaPaddingBottom := 2.0
	nextY := documentoY + recebedorAssinaturaPaddingBottom

	return nextY
}

//-----------------------------------------------------------------

// ! ========== DIVISOR DESTINATÁRIO ==========
func DrawDestinatarioCorreiosLogoDivisor(pdf *gofpdf.Fpdf, x, y float64, local bool) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	const DESTINATARIO = "DESTINATÁRIO:"
	// destinatarioTextWidth := pdf.GetStringWidth(DESTINATARIO) + 10
	lineHeight := 8.0
	fontSize := 9.0

	destinatarioText := translator(DESTINATARIO)
	pdf.SetFont("Arial", "B", fontSize)
	pdf.Text(x+paddingDestinatario, y+lineHeight/2, destinatarioText)

	//pdf.SetLineWidth(0.3)
	pdf.Line(x, y, x, y+dadosDestinatarioHeight)
	pdf.Line(x, y, x+labelWidth-paddingRight+1, y)
	pdf.Line(x+labelWidth-paddingRight+1, y, x+labelWidth-paddingRight+1, y+dadosDestinatarioHeight)
	//pdf.SetLineWidth(defaultLineWidth)

	//! DESENHA O RETANGULO COM FUNDO PRETO
	/*
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
	*/
	return y + lineHeight
}

//-----------------------------------------------------------------

func buildLogradouroDestinatarioString(destinatario types.SolicitarEtiquetaDestinatario) string {
	var hasNumerodestinatario = destinatario.NumeroDestinatario != ""

	var logradouroDestinatarioString string

	logradouroDestinatarioString += destinatario.LogradouroDestinatario

	if hasNumerodestinatario {
		logradouroDestinatarioString += ", "
		logradouroDestinatarioString += destinatario.NumeroDestinatario
	}

	return logradouroDestinatarioString
}

func buildComplementoBairroDestinatarioString(destinatario types.SolicitarEtiquetaDestinatario) string {
	var hasComplemento bool
	if destinatario.ComplementoDestinatario != nil && *destinatario.ComplementoDestinatario != "" {
		hasComplemento = true
	} else {
		hasComplemento = false
	}
	var hasBairro = destinatario.BairroDestinatario != ""

	var complementoBairroDestinatarioString string

	if hasComplemento {
		complementoBairroDestinatarioString += *destinatario.ComplementoDestinatario
	}
	if hasComplemento && hasBairro {
		complementoBairroDestinatarioString += ", "
	}
	if hasBairro {
		complementoBairroDestinatarioString += destinatario.BairroDestinatario
	}

	return complementoBairroDestinatarioString
}

func buildCepDestinatarioString(destinatario types.SolicitarEtiquetaDestinatario) string {
	var formattedCEP string
	cep := destinatario.CepDestinatario
	lenCEP := len(destinatario.CepDestinatario)

	if lenCEP != 8 {
		if lenCEP == 7 {
			formattedCEP += "0"
			formattedCEP += destinatario.CepDestinatario
		} else {
			log.Fatalf("CEP INVÁLIDO")
			panic("CEP INVÁLIDO")
		}
	} else {
		formattedCEP = cep[:5] + "-" + cep[5:]
	}

	return formattedCEP
}

func buildCidadeUfDestinatarioString(destinatario types.SolicitarEtiquetaDestinatario) string {
	cidadeUfDestinatarioString := fmt.Sprintf("%s / %s", destinatario.CidadeDestinatario, destinatario.UfDestinatario)
	return cidadeUfDestinatarioString
}

// ! ========== DADOS DO DESTINATÁRIO ==========
func DrawDadosDestinatario(pdf *gofpdf.Fpdf, x, y float64, destinatario types.SolicitarEtiquetaDestinatario, isCarta bool) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	fontSize := 9.0
	lineHeight := 4.0

	if isCarta {
		fontSize = 8.5
		lineHeight = 3.0

		pdf.SetFont("Arial", "B", fontSize)
		destinatarioX := x
		destinatarioY := y
		destinatarioText := translator("Destinatário:")
		pdf.Text(destinatarioX, destinatarioY, destinatarioText)
	}

	pdf.SetFont("Arial", "", fontSize)
	nomeDestinatarioX := x
	nomeDestinatarioY := y
	if isCarta {
		nomeDestinatarioY = y + lineHeight
	}

	nomeDestinatarioText := translator(destinatario.NomeDestinatario)

	// Adjust font size if name is too long
	originalFontSize := fontSize
	availableWidth := labelWidth - paddingRight - paddingDestinatario
	textWidth := pdf.GetStringWidth(nomeDestinatarioText)

	// Reduce font size if text is too wide
	adjustedFontSize := originalFontSize
	for textWidth > availableWidth && adjustedFontSize > 6.0 {
		adjustedFontSize -= 0.5
		pdf.SetFont("Arial", "", adjustedFontSize)
		textWidth = pdf.GetStringWidth(nomeDestinatarioText)
	}

	pdf.Text(nomeDestinatarioX, nomeDestinatarioY, nomeDestinatarioText)

	// Restore original font size
	pdf.SetFont("Arial", "", originalFontSize)

	logradouroDestinatarioX := x
	logradouroDestinatarioY := nomeDestinatarioY + lineHeight
	logradouroDestinatarioString := buildLogradouroDestinatarioString(destinatario)
	logradouroDestinatarioText := translator(logradouroDestinatarioString)
	pdf.Text(logradouroDestinatarioX, logradouroDestinatarioY, logradouroDestinatarioText)

	complementoBairroDestinatarioX := x
	complementoBairroDestinatarioY := logradouroDestinatarioY + lineHeight
	complementoBairroDestinatarioString := buildComplementoBairroDestinatarioString(destinatario)
	complementoBairroDestinatarioText := translator(complementoBairroDestinatarioString)
	pdf.Text(complementoBairroDestinatarioX, complementoBairroDestinatarioY, complementoBairroDestinatarioText)

	cepCidadeUfDestinatarioY := complementoBairroDestinatarioY + lineHeight*1.1

	pdf.SetFont("Arial", "B", fontSize)
	cepDestinatarioX := x
	cepDestinatarioString := buildCepDestinatarioString(destinatario)
	cepDestinatarioText := translator(cepDestinatarioString)
	pdf.Text(cepDestinatarioX, cepCidadeUfDestinatarioY, cepDestinatarioText)

	pdf.SetFont("Arial", "", fontSize)
	cidadeUfDestinatarioPaddingLeft := 1.5
	cidadeUfDestinatarioX := x + pdf.GetStringWidth(cepDestinatarioString) + cidadeUfDestinatarioPaddingLeft
	cidadeUfDestinatarioString := buildCidadeUfDestinatarioString(destinatario)
	cidadeUfDestinatarioText := translator(cidadeUfDestinatarioString)
	pdf.Text(cidadeUfDestinatarioX, cepCidadeUfDestinatarioY, cidadeUfDestinatarioText)

	nextY := cepCidadeUfDestinatarioY + 2.0

	return nextY
}

// ! ========== BARRA DE CÓDIGO DESTINATÁRIO ==========
func DrawDestinatarioBarCode(pdf *gofpdf.Fpdf, x, y float64, destinatarioBarcodeBase64String string) float64 {
	// generateBarcode128(pdf, destinatarioBarcodeBase64String, x, y, destinatarioBarCodeWidth, destinatarioBarCodeHeight)
	errDestinatarioBarCode := addBase64ImageToPDF(pdf, destinatarioBarcodeBase64String, x, y, destinatarioBarCodeWidth, destinatarioBarCodeHeight)

	if errDestinatarioBarCode != nil {
		errDestinatarioBarCodeString := fmt.Sprintf("Erro DrawDestinatarioBarCode generateBarcode128 %s", errDestinatarioBarCode.Error())
		log.Fatalf(errDestinatarioBarCodeString)
		panic(errDestinatarioBarCodeString)
	}

	return y + destinatarioBarCodeHeight
}

// ! ========== OBSERVAÇÕES ==========
func DrawObservacoes(pdf *gofpdf.Fpdf, x, y float64, servicoAdicional *types.ServicoAdicional) {
	if servicoAdicional == nil {
		return
	}

	codigoServicoAdicional := servicoAdicional.CodigoServicoAdicional

	translator := pdf.UnicodeTranslatorFromDescriptor("")
	fontSize := 9.0
	pdf.SetFont("Arial", "", fontSize)
	observacoesTextString := "Obs.:"

	var observacoes []string
	for _, codigoStr := range codigoServicoAdicional {
		codigo, err := strconv.Atoi(codigoStr)
		if err != nil {
			continue
		}

		switch codigo {
		case 1:
			observacoes = append(observacoes, "Aviso de Recebimento")
		case 2:
			observacoes = append(observacoes, "Mão Própria")
		case 19, 64, 65, 75, 76:
			observacoes = append(observacoes, "Valor Declarado")
		}
	}

	if len(observacoes) > 0 {
		observacoesX := x + 1.3*destinatarioBarCodeWidth + 1.0
		currentY := y + 2.0
		pdf.Text(observacoesX, currentY, translator("Obs.:"))
		servicoasAdicionaisX := observacoesX + pdf.GetStringWidth("Obs.:") + 0.5
		for _, servAdicional := range observacoes {
			pdf.Text(servicoasAdicionaisX, currentY, translator(servAdicional))
			currentY += 4.0
		}
		observacoesTextString += " " + strings.Join(observacoes, ", ")
	}
	// pdf.Text(observacoesX, observacoesY, observacoesText)
}

// ! ========== SEPARADOR REMETENTE ==========
func DrawSeparadorRemetente(pdf *gofpdf.Fpdf, x, y float64) float64 {
	paddingTop := 2.0
	paddingBottom := 4.0

	//pdf.SetLineWidth(0.5)
	pdf.Line(x, y+paddingTop, x+labelWidth-paddingRight+1, y+paddingTop)
	//pdf.SetLineWidth(defaultLineWidth)

	nextY := paddingTop + y + paddingBottom
	return nextY
}

//----------------------------------------------------

func buildLogradouroRemetenteString(remetente types.SolicitarEtiquetaRemetente) string {
	var hasNumeroRemetente = remetente.NumeroRemetente != ""

	var logradouroRemetenteString string

	logradouroRemetenteString += remetente.LogradouroRemetente

	if hasNumeroRemetente {
		logradouroRemetenteString += ", "
		logradouroRemetenteString += remetente.NumeroRemetente
	}

	return logradouroRemetenteString
}

func buildComplementoBairroRemetenteString(remetente types.SolicitarEtiquetaRemetente) string {
	var hasComplemento bool
	if remetente.ComplementoRemetente != nil && *remetente.ComplementoRemetente != "" {
		hasComplemento = true
	} else {
		hasComplemento = false
	}
	var hasBairro = remetente.BairroRemetente != ""

	var complementoBairroRemetenteString string

	if hasComplemento {
		complementoBairroRemetenteString += *remetente.ComplementoRemetente
	}
	if hasComplemento && hasBairro {
		complementoBairroRemetenteString += ", "
	}
	if hasBairro {
		complementoBairroRemetenteString += remetente.BairroRemetente
	}

	return complementoBairroRemetenteString
}

func buildCepRemetenteString(remetente types.SolicitarEtiquetaRemetente) string {
	var formattedCEP string
	cep := remetente.CepRemetente
	lenCEP := len(remetente.CepRemetente)

	if lenCEP != 8 {
		if lenCEP == 7 {
			formattedCEP += "0"
			formattedCEP += remetente.CepRemetente
		} else {
			log.Fatalf("CEP INVÁLIDO")
			panic("CEP INVÁLIDO")
		}
	} else {
		formattedCEP = cep[:5] + "-" + cep[5:]
	}
	return formattedCEP
}

func buildCidadeUfRemetenteString(remetente types.SolicitarEtiquetaRemetente) string {
	cidadeUfRemetenteString := fmt.Sprintf("%s / %s", remetente.CidadeRemetente, remetente.UfRemetente)
	return cidadeUfRemetenteString
}

// ! ========== DADOS DO REMETENTE ==========
func DrawDadosRemetente(pdf *gofpdf.Fpdf, x, y float64, remetente types.SolicitarEtiquetaRemetente, isCarta bool) float64 {
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	fontSize := 8.5
	lineHeight := 3.5

	if isCarta {
		fontSize = 6.0
		lineHeight = 2.0
	}

	pdf.SetFont("Arial", "B", fontSize)
	nomeRemetenteY := y
	remetenteX := x
	remetenteText := translator("Remetente: ")
	pdf.Text(remetenteX, nomeRemetenteY, remetenteText)

	pdf.SetFont("Arial", "", fontSize)
	nomeRemetentePaddingLeft := 1.0
	nomeRemetenteX := x + pdf.GetStringWidth("Remetente: ") + nomeRemetentePaddingLeft
	nomeRemetenteText := translator(remetente.NomeRemetente)

	// Adjust font size if name is too long
	originalFontSize := fontSize
	availableWidth := labelWidth - paddingRight - (nomeRemetenteX - x)
	textWidth := pdf.GetStringWidth(nomeRemetenteText)

	// Reduce font size if text is too wide
	adjustedFontSize := originalFontSize
	for textWidth > availableWidth && adjustedFontSize > 4.0 {
		adjustedFontSize -= 0.5
		pdf.SetFont("Arial", "", adjustedFontSize)
		textWidth = pdf.GetStringWidth(nomeRemetenteText)
	}

	pdf.Text(nomeRemetenteX, nomeRemetenteY, nomeRemetenteText)

	// Restore original font size
	pdf.SetFont("Arial", "", originalFontSize)

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

// ! ========== NÚMERO DA NOTA FISCAL ==========
func DrawNumeroNotaFiscal(pdf *gofpdf.Fpdf, x, y float64, numeroNotaFiscal *string) {
	if numeroNotaFiscal == nil {
		return
	}

	pdf.SetFont("Arial", "B", 7)
	pdf.Text(x, y, "NF:")
	pdf.SetFont("Arial", "", 7)
	pdf.Text(x+pdf.GetStringWidth("NF: "), y, *numeroNotaFiscal)
}

func DrawChancelaCarta(pdf *gofpdf.Fpdf, x, y float64, tipoServicoImagem string, local bool) {
	//! TIPO SERVIÇO LOGO
	var tipoServicoImagemPath string
	if local {
		// Get workspace root directory (go up one level from current directory)
		workspaceRoot := filepath.Dir(GetWorkingDir())

		// Try direct path first
		tipoServicoImagemPath = filepath.Join(workspaceRoot, "layers/images", tipoServicoImagem)
		log.Printf("Trying chancela image path: %s", tipoServicoImagemPath)

		// Check if file exists
		if _, err := os.Stat(tipoServicoImagemPath); os.IsNotExist(err) {
			log.Printf("WARNING: Chancela image file not found at %s", tipoServicoImagemPath)

			// Try alternate path formats
			alternativePaths := []string{
				filepath.Join(workspaceRoot, "layers", "images", tipoServicoImagem),
				filepath.Join(GetWorkingDir(), "..", "..", "layers", "images", tipoServicoImagem),
				filepath.Join(GetWorkingDir(), "..", "layers", "images", tipoServicoImagem),
				filepath.Join("/home/gabriel/Favor/Github/pdf-generator/layers/images", tipoServicoImagem),
			}

			for _, altPath := range alternativePaths {
				log.Printf("Trying alternative path for chancela: %s", altPath)
				if _, err := os.Stat(altPath); err == nil {
					log.Printf("Found chancela image at path: %s", altPath)
					tipoServicoImagemPath = altPath
					break
				}
			}
		}
	} else {
		tipoServicoImagemPath = filepath.Join("/opt", "bin", "images", tipoServicoImagem)
	}

	log.Printf("Loading chancela image from: %s", tipoServicoImagemPath)

	chancelaX := x
	chancelaY := y
	imgErr := addImage(pdf, tipoServicoImagemPath, chancelaX, chancelaY, tipoServicoSize, tipoServicoSize, true)
	if imgErr != nil {
		log.Printf("Error adding chancela image: %v", imgErr)
	}

	textPaddingTop := 9.5
	textHeight := 2.5
	textX := x - 1.0
	textY := y + 20.5 + textPaddingTop
	pdf.SetXY(textX, textY)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(tipoServicoSize, textHeight, "Data de Postagem", "", 0, "L", false, 0, "")

	dataText := getCurrentDateAsString()
	dataTextX := x + pdf.GetStringWidth(dataText)/2 - 2.5
	dataTextY := y + 20.5 + textHeight + textPaddingTop
	pdf.SetXY(dataTextX, dataTextY)
	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(tipoServicoSize, textHeight, dataText, "", 0, "L", false, 0, "")
}

func DrawDataMatrixCarta(pdf *gofpdf.Fpdf, x, y float64, dataMatrixBase64String string, lote string) {
	errDataMatrix := addBase64ImageToPDF(pdf, dataMatrixBase64String, x, y, 15, 15)
	if errDataMatrix != nil {
		errDataMatrixString := fmt.Sprintf("Erro generateDataMatrix %s", errDataMatrix.Error())
		log.Fatalf(errDataMatrixString)
		panic(errDataMatrixString)
	}

	loteTextHeight := 3.0
	loteX := x
	loteY := y + 16.5
	pdf.SetXY(loteX, loteY)
	pdf.SetFont("Arial", "B", 7)
	pdf.CellFormat(tipoServicoSize, loteTextHeight, "ID:", "", 0, "L", false, 0, "")

	loteTextX := x
	loteTextY := y + 16.5 + loteTextHeight
	pdf.SetXY(loteTextX, loteTextY)
	pdf.SetFont("Arial", "", 6)
	pdf.CellFormat(tipoServicoSize, loteTextHeight, lote, "", 0, "L", false, 0, "")
}

// ! ========== FORMATADOR DO CÓDIGO DE RASTREIO ==========
func FormatTrackingCode(code string) string {
	if len(code) != 13 {
		return code
	}

	return code[:2] + " " + code[2:5] + " " + code[5:8] + " " + code[8:11] + " " + code[11:]
}

func getCurrentDateAsString() string {
	currentTime := time.Now()
	return currentTime.Format("02/01/2006")
}

// GetWorkingDir returns the current working directory
func GetWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory: %v", err)
		return "unknown"
	}
	return dir
}
