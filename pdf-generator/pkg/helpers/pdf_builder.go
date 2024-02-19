package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/jung-kurt/gofpdf"
)

func chunkifyObjetosPostais(objetos []types.SolicitarEtiquetasPDFObjetoPostal, chunkSize int) [][]types.SolicitarEtiquetasPDFObjetoPostal {
	var chunks [][]types.SolicitarEtiquetasPDFObjetoPostal
	for i := 0; i < len(objetos); i += chunkSize {
		end := i + chunkSize
		if end > len(objetos) {
			end = len(objetos)
		}
		chunks = append(chunks, objetos[i:end])
	}
	return chunks
}

func GenerateLabelsPDFLocal(solicitarEtiquetasPDF types.SolicitarEtiquetasPDF) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	remetente := solicitarEtiquetasPDF.Remetente
	chunkifiedObjetoPostal := chunkifyObjetosPostais(solicitarEtiquetasPDF.ObjetosPostais, 4)
	fmt.Println(" - Número de Páginas do PDF: ", len(chunkifiedObjetoPostal))
	for k, objetoPostalChunk := range chunkifiedObjetoPostal {
		fmt.Println("   - Desenhando a página ", k)
		pdf.AddPage()
		//DrawDottedLines(pdf)

		subdivisionStartPoints := []struct{ x, y float64 }{
			{0, 0},
			{pageWidth / 2, 0},
			{0, pageHeight / 2},
			{pageWidth / 2, pageHeight / 2},
		}

		for i, objetoPostal := range objetoPostalChunk {
			fmt.Println("     - Desenhando a etiqueta ", i)
			startPoint := subdivisionStartPoints[i]
			DrawSmallLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, remetente, objetoPostal, true)
		}
	}

	return pdf.OutputFileAndClose("label.pdf")
}

func GenerateLabelsPDF(solicitarEtiquetasPDF types.SolicitarEtiquetasPDF) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	remetente := solicitarEtiquetasPDF.Remetente
	chunkifiedObjetoPostal := chunkifyObjetosPostais(solicitarEtiquetasPDF.ObjetosPostais, 4)

	fmt.Println(" - Número de Páginas do PDF: ", len(chunkifiedObjetoPostal))
	for k, objetoPostalChunk := range chunkifiedObjetoPostal {
		fmt.Println("   - Desenhando a página ", k+1)
		pdf.AddPage()
		//DrawDottedLines(pdf)

		subdivisionStartPoints := []struct{ x, y float64 }{
			{0, 0},
			{pageWidth / 2, 0},
			{0, pageHeight / 2},
			{pageWidth / 2, pageHeight / 2},
		}

		for i, objetoPostal := range objetoPostalChunk {
			fmt.Println("     - Desenhando a etiqueta ", i+1)
			startPoint := subdivisionStartPoints[i]
			DrawSmallLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, remetente, objetoPostal, false)
		}
	}

	var buffer bytes.Buffer

	err := pdf.Output(&buffer)
	if err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	fmt.Println("base64 string")
	fmt.Println(base64Str)
	return base64Str, nil
}

func DrawSmallLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, remetente types.SolicitarEtiquetaRemetente, objetoPostal types.SolicitarEtiquetasPDFObjetoPostal, local bool) {
	pesoObjeto := objetoPostal.Peso
	codRastreio := FormatTrackingCode(objetoPostal.CodigoRastreio)
	codServicoPostagem := objetoPostal.CodigoServicoPostagem
	tipoServicoImagem := findTipoServicoImagemByCodServicoPostagem(codServicoPostagem)
	idPrePostagem := objetoPostal.IdPrePostagem

	fmt.Println("        * Tipo serviço postagem: ", tipoServicoImagem)
	fmt.Println()

	dataMatrixBase64String := objetoPostal.Base64.Datamatrix
	barcodeBase64String := objetoPostal.Base64.Code
	destinatarioBarcodeBase64String := objetoPostal.Base64.CepBarcode

	paddingTop := 3.5 + 4.75
	paddingLeft := 6.0 + 3.5
	var nextY = y + paddingTop

	//DrawSmallDelimiter(pdf, x, y)
	//! LOGO FAVOR, DATAMATRIX, TIPO SERVIÇO LOGO E numero PLP
	nextY = DrawSmallFirstRow(pdf, x+paddingLeft, nextY, idPrePostagem, tipoServicoImagem, dataMatrixBase64String, local)
	//! PEDIDO, NF E PESO
	nextY = DrawSmallSecondRow(pdf, x+paddingLeft, nextY, idPrePostagem, pesoObjeto)
	//! CÓDIGO DE RASTREIO
	nextY = DrawTrackingCode(pdf, x, nextY, codRastreio)
	//! BARRA DE CÓDIGO
	nextY = DrawBarcode(pdf, x, nextY, barcodeBase64String)
	//! RECEBEDOR, ASSINATURA e DOCUMENTO
	nextY = DrawSmallRecebedorAssinaturaDocumentoLines(pdf, x+paddingLeft, nextY)
	//! SEPARADOR DESTINATÁRIO E LOGO CORREIOS
	nextY = DrawSmallDestinatarioCorreiosLogoDivisor(pdf, x+paddingLeft, nextY, local)
	//! DADOS DESTINATÁRIO
	paddingDestinatario := 3.0
	nextY = DrawDadosDestinatario(pdf, x+paddingLeft+paddingDestinatario, nextY, objetoPostal.Destinatario)
	//! BARRA DE CODIGO DESTINATARIO
	DrawObservacoes(pdf, x, nextY, objetoPostal.ServicoAdicional)
	nextY = DrawDestinatarioBarCode(pdf, x+paddingLeft+paddingDestinatario, nextY, destinatarioBarcodeBase64String)
	//fmt.Printf("nextY DrawSmallSeparadorRemetente %f\n", nextY)
	//! SEPARADOR REMETENTE
	nextY = DrawSmallSeparadorRemetente(pdf, x+paddingLeft, nextY)
	//! DADOS REMETENTE
	DrawDadosRemetente(pdf, x+paddingLeft, nextY, remetente)
}
