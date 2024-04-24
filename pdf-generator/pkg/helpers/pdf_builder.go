package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/jung-kurt/gofpdf"
)

func chunkifyObjetosPostais(objetos []types.SolicitarEtiquetasPDFObjetoPostal, chunkSize int) [][]types.SolicitarEtiquetasPDFObjetoPostal {
	var chunks [][]types.SolicitarEtiquetasPDFObjetoPostal

	if len(objetos) == 0 {
		chunks = append(chunks, objetos)
		return chunks
	}

	for i := 0; i < len(objetos); i += chunkSize {
		end := i + chunkSize
		if end > len(objetos) {
			end = len(objetos)
		}
		chunks = append(chunks, objetos[i:end])
	}

	return chunks
}

func separateObjetosPostais(objetos []types.SolicitarEtiquetasPDFObjetoPostal) ([]types.SolicitarEtiquetasPDFObjetoPostal, []types.SolicitarEtiquetasPDFObjetoPostal) {
	objetosPostaisCartas := make([]types.SolicitarEtiquetasPDFObjetoPostal, 0)
	objetosPostaisPadrao := make([]types.SolicitarEtiquetasPDFObjetoPostal, 0)
	codigosServicosPostagemCarta := []string{"80160"}

	// Create a map for efficient checking against service codes
	serviceCodeMap := make(map[string]bool)
	for _, code := range codigosServicosPostagemCarta {
		serviceCodeMap[code] = true
	}

	for _, objeto := range objetos {
		if _, found := serviceCodeMap[objeto.CodigoServicoPostagem]; found {
			objetosPostaisCartas = append(objetosPostaisCartas, objeto)
		} else {
			objetosPostaisPadrao = append(objetosPostaisPadrao, objeto)
		}
	}

	return objetosPostaisPadrao, objetosPostaisCartas
}

func GenerateLabelsPDFLocal(solicitarEtiquetasPDF types.SolicitarEtiquetasPDF) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	remetente := solicitarEtiquetasPDF.Remetente
	objetosPostaisPadrao, objetosPostaisCartas := separateObjetosPostais(solicitarEtiquetasPDF.ObjetosPostais)

	chunkifiedObjetosPostaisPadrao := chunkifyObjetosPostais(objetosPostaisPadrao, 4)
	chunkifiedObjetosPostaisCarta := chunkifyObjetosPostais(objetosPostaisCartas, 10)

	fmt.Println(" - Número de Páginas do PDF: ", len(chunkifiedObjetosPostaisPadrao)+len(chunkifiedObjetosPostaisCarta))

	//! CONSTRUINDO AS ETIQUETAS NORMAIS
	if len(objetosPostaisPadrao) != 0 {
		for k, objetoPostalChunk := range chunkifiedObjetosPostaisPadrao {
			fmt.Println("   - Desenhando a página ", k)
			pdf.AddPage()

			subdivisionStartPoints := []struct{ x, y float64 }{
				{0, 0},
				{pageWidth / 2, 0},
				{0, pageHeight / 2},
				{pageWidth / 2, pageHeight / 2},
			}

			for i, objetoPostal := range objetoPostalChunk {
				fmt.Println("     - Desenhando a etiqueta ", i)
				startPoint := subdivisionStartPoints[i]
				DrawLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, remetente, objetoPostal, true)
			}
		}
	}

	if len(objetosPostaisCartas) != 0 {
		pdf.SetMargins(0, 0, 0)
		pdf.SetAutoPageBreak(false, 0)
		//! CONSTRUINDO AS ETIQUETAS DE CARTA
		for k, objetoPostalCartaChunk := range chunkifiedObjetosPostaisCarta {
			fmt.Println("   - Desenhando a página ", k)
			pdf.AddPage()

			subdivisionStartPoints := []struct{ x, y float64 }{
				{0, 0},
				{pageWidth / 2, 0},
				{0, 1 * pageHeight / 5},
				{pageWidth / 2, 1 * pageHeight / 5},
				{0, 2 * pageHeight / 5},
				{pageWidth / 2, 2 * pageHeight / 5},
				{0, 3 * pageHeight / 5},
				{pageWidth / 2, 3 * pageHeight / 5},
				{0, 4 * pageHeight / 5},
				{pageWidth / 2, 4 * pageHeight / 5},
			}

			for i, objetoPostal := range objetoPostalCartaChunk {
				fmt.Println("     - Desenhando a carta ", i+1)
				startPoint := subdivisionStartPoints[i]
				DrawCartaLabel(pdf, startPoint.x, startPoint.y, cartaWidth, cartaHeight, i, remetente, objetoPostal, true)
			}
		}
	}

	var buffer bytes.Buffer
	err := pdf.Output(&buffer)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("ERRO AO TRANSFORMAR PDF EM BASE64STRING: %s", errMsg)
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	fmt.Printf("%s", base64Str)

	return pdf.OutputFileAndClose("label.pdf")
}

func GenerateLabelsPDF(solicitarEtiquetasPDF types.SolicitarEtiquetasPDF) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	remetente := solicitarEtiquetasPDF.Remetente
	objetosPostaisPadrao, objetosPostaisCartas := separateObjetosPostais(solicitarEtiquetasPDF.ObjetosPostais)

	chunkifiedObjetosPostaisPadrao := chunkifyObjetosPostais(objetosPostaisPadrao, 4)
	chunkifiedObjetosPostaisCarta := chunkifyObjetosPostais(objetosPostaisCartas, 10)

	fmt.Println(" - Número de Páginas do PDF: ", len(chunkifiedObjetosPostaisPadrao)+len(chunkifiedObjetosPostaisCarta))

	if len(objetosPostaisPadrao) != 0 {
		for k, objetoPostalChunk := range chunkifiedObjetosPostaisPadrao {
			fmt.Println("   - Desenhando a página ", k+1)
			pdf.AddPage()

			subdivisionStartPoints := []struct{ x, y float64 }{
				{0, 0},
				{pageWidth / 2, 0},
				{0, pageHeight / 2},
				{pageWidth / 2, pageHeight / 2},
			}

			for i, objetoPostal := range objetoPostalChunk {
				fmt.Println("     - Desenhando a etiqueta ", i+1)
				startPoint := subdivisionStartPoints[i]
				DrawLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, remetente, objetoPostal, false)
			}
		}
	}

	if len(objetosPostaisCartas) != 0 {
		pdf.SetMargins(0, 0, 0)
		pdf.SetAutoPageBreak(false, 0)
		//! CONSTRUINDO AS ETIQUETAS DE CARTA
		for k, objetoPostalCartaChunk := range chunkifiedObjetosPostaisCarta {
			fmt.Println("   - Desenhando a página ", k)
			pdf.AddPage()

			subdivisionStartPoints := []struct{ x, y float64 }{
				{0, 0},
				{pageWidth / 2, 0},
				{0, 1 * pageHeight / 5},
				{pageWidth / 2, 1 * pageHeight / 5},
				{0, 2 * pageHeight / 5},
				{pageWidth / 2, 2 * pageHeight / 5},
				{0, 3 * pageHeight / 5},
				{pageWidth / 2, 3 * pageHeight / 5},
				{0, 4 * pageHeight / 5},
				{pageWidth / 2, 4 * pageHeight / 5},
			}

			for i, objetoPostal := range objetoPostalCartaChunk {
				fmt.Println("     - Desenhando a carta ", i+1)
				startPoint := subdivisionStartPoints[i]
				DrawCartaLabel(pdf, startPoint.x, startPoint.y, cartaWidth, cartaHeight, i, remetente, objetoPostal, false)
			}
		}
	}

	var buffer bytes.Buffer
	err := pdf.Output(&buffer)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("ERRO AO TRANSFORMAR PDF EM BASE64STRING: %s", errMsg)
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return base64Str, nil
}

func DrawLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, remetente types.SolicitarEtiquetaRemetente, objetoPostal types.SolicitarEtiquetasPDFObjetoPostal, local bool) {
	pesoObjeto := objetoPostal.Peso
	codRastreio := FormatTrackingCode(objetoPostal.CodigoRastreio)
	codServicoPostagem := objetoPostal.CodigoServicoPostagem
	tipoServicoImagem := findTipoServicoImagemByCodServicoPostagem(codServicoPostagem)
	idPrePostagem := objetoPostal.IdPrePostagem

	fmt.Printf("        * x: %.2f\n", x)
	fmt.Printf("        * y: %.2f\n", y)
	fmt.Println("        * Tipo serviço postagem: ", tipoServicoImagem)

	dataMatrixBase64String := CreateDatamatrixBaseString(objetoPostal.DatamatrixString, 25, 25)
	barcodeBase64String := CreateBarcodeBaseString(80, 18, objetoPostal.CodigoRastreio)
	destinatarioBarcodeBase64String := CreateBarcodeBaseString(40, 18, objetoPostal.Destinatario.CepDestinatario)
	paddingTop := 14.0
	if y == pageHeight/2 {
		paddingTop = 8.0
	}
	paddingLeft := 6.0 + 3.5
	var nextY = y + paddingTop

	//! LOGO FAVOR, DATAMATRIX, TIPO SERVIÇO LOGO E numero PLP
	nextY = DrawFirstRow(pdf, x+paddingLeft, nextY, idPrePostagem, tipoServicoImagem, dataMatrixBase64String, local)
	//! PEDIDO, NF E PESO
	nextY = DrawSecondRow(pdf, x+paddingLeft, nextY, idPrePostagem, pesoObjeto)
	//! CÓDIGO DE RASTREIO
	nextY = DrawTrackingCode(pdf, x, nextY, codRastreio)
	//! BARRA DE CÓDIGO
	nextY = DrawBarcode(pdf, x, nextY, barcodeBase64String)
	//! RECEBEDOR, ASSINATURA e DOCUMENTO
	nextY = DrawRecebedorAssinaturaDocumentoLines(pdf, x+paddingLeft, nextY)
	//! SEPARADOR DESTINATÁRIO E LOGO CORREIOS
	nextY = DrawDestinatarioCorreiosLogoDivisor(pdf, x+paddingLeft, nextY, local)
	//! DADOS DESTINATÁRIO
	paddingDestinatario := 3.0
	nextY = DrawDadosDestinatario(pdf, x+paddingLeft+paddingDestinatario, nextY, objetoPostal.Destinatario)
	//! BARRA DE CODIGO DESTINATARIO
	DrawObservacoes(pdf, x, nextY, objetoPostal.ServicoAdicional)
	nextY = DrawDestinatarioBarCode(pdf, x+paddingLeft+paddingDestinatario, nextY, destinatarioBarcodeBase64String)
	//fmt.Printf("nextY DrawSmallSeparadorRemetente %f\n", nextY)
	//! SEPARADOR REMETENTE
	nextY = DrawSeparadorRemetente(pdf, x+paddingLeft, nextY)
	//! DADOS REMETENTE
	DrawDadosRemetente(pdf, x+paddingLeft, nextY, remetente)
}

func DrawCartaLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, remetente types.SolicitarEtiquetaRemetente, objetoPostal types.SolicitarEtiquetasPDFObjetoPostal, local bool) {
	colunaEsquerda := x == 0
	linhaInicial := y == 0
	paddingTop := 9.0

	//! ==================== REMETENTE ====================
	remetenteX := 10.5
	remetenteY := y + 13.0
	if !colunaEsquerda {
		remetenteX = x + 8.5
	}
	if !linhaInicial {
		remetenteY = y + 3.0 + paddingTop
	}
	DrawDadosRemetente(pdf, remetenteX, remetenteY, remetente)

	//! ==================== CHANCELA E DATA DE POSTAGEM ====================
	tipoServicoImagem := findTipoServicoImagemByCodServicoPostagem(objetoPostal.CodigoServicoPostagem)
	chancelaX := x + 80.0
	chancelaY := y + 12.0

	if !linhaInicial {
		chancelaY = y + 7.0 + paddingTop
	}

	DrawChancelaCarta(pdf, chancelaX, chancelaY, tipoServicoImagem, local)

	//! ==================== DATAMATRIX ====================
	dataMatrixBase64String := CreateDatamatrixBaseString(objetoPostal.DatamatrixString, 15, 15)
	dataMatrixX := x + 5.0
	dataMatrixY := y + 29.0

	if !colunaEsquerda {
		dataMatrixX = x + 2.0
	}
	if !linhaInicial {
		dataMatrixY = y + 18.0 + paddingTop
	}

	DrawDataMatrixCarta(pdf, dataMatrixX, dataMatrixY, dataMatrixBase64String, objetoPostal.IdPrePostagem)

	//! ==================== DESTINATARIO ====================
	destinatarioX := x + 20.0
	destinatarioY := y + 32.0

	if !colunaEsquerda {
		destinatarioX = x + 18.0
	}
	if !linhaInicial {
		destinatarioY = y + 22.0 + paddingTop
	}
	DrawDadosDestinatario(pdf, destinatarioX, destinatarioY, objetoPostal.Destinatario)
}
