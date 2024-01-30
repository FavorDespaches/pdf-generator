package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/FavorDespaches/pdf-generator/pkg/types"
	"github.com/jung-kurt/gofpdf"
)

func chunkifyObjetosPostais(objetos []types.ObjetoPostal, chunkSize int) [][]types.ObjetoPostal {
	var chunks [][]types.ObjetoPostal
	for i := 0; i < len(objetos); i += chunkSize {
		end := i + chunkSize
		if end > len(objetos) {
			end = len(objetos)
		}
		chunks = append(chunks, objetos[i:end])
	}
	return chunks
}

func GenerateLabelsPDFLocal(correiosLog types.CorreiosLog) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	idPlp := correiosLog.Plp.IdPlp
	remetente := correiosLog.Remetente
	chunkifiedObjetoPostal := chunkifyObjetosPostais(correiosLog.ObjetoPostal, 4)
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
			DrawSmallLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, idPlp, remetente, objetoPostal, true)
		}
	}

	return pdf.OutputFileAndClose("label.pdf")
}

func GenerateLabelsPDF(correiosLog types.CorreiosLog) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	idPlp := correiosLog.Plp.IdPlp
	remetente := correiosLog.Remetente
	chunkifiedObjetoPostal := chunkifyObjetosPostais(correiosLog.ObjetoPostal, 4)

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
			DrawSmallLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, idPlp, remetente, objetoPostal, false)
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

func DrawLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, idPlp int, remetente types.Remetente, objetoPostal types.ObjetoPostal, local bool) {
	pesoObjeto := objetoPostal.Peso
	codRastreio := FormatTrackingCode(objetoPostal.NumeroEtiqueta)
	codServicoPostagem := objetoPostal.CodigoServicoPostagem
	tipoServicoImagem := findTipoServicoImagemByCodServicoPostagem(codServicoPostagem)
	fmt.Println("     - Tipo serviço postagem etiqueta", tipoServicoImagem)

	dataMatrixBase64String := objetoPostal.Base64.Datamatrix
	barcodeBase64String := objetoPostal.Base64.Code
	destinatarioBarcodeBase64String := objetoPostal.Base64.CepBarcode

	paddingTop := 2.0
	paddingLeft := 5.0
	var nextY = y + paddingTop

	//! LOGO FAVOR, DATAMATRIX, TIPO SERVIÇO LOGO E numero PLP
	nextY = DrawFirstRow(pdf, x+paddingLeft, nextY, idPlp, tipoServicoImagem, dataMatrixBase64String, local)

	//! PEDIDO, NF E PESO
	nextY = DrawSecondRow(pdf, x+paddingLeft, nextY, pesoObjeto)

	//! CÓDIGO DE RASTREIO
	nextY = DrawTrackingCode(pdf, x, nextY, codRastreio)

	//! BARRA DE CÓDIGO
	nextY = DrawBarcode(pdf, x, nextY, barcodeBase64String)

	//! RECEBEDOR, ASSINATURA e DOCUMENTO
	nextY = DrawRecebedorAssinaturaDocumentoLines(pdf, x+paddingLeft, nextY)

	//! SEPARADOR DESTINATÁRIO E LOGO CORREIOS
	nextY = DrawDestinatarioCorreiosLogoDivisor(pdf, x, nextY, local)

	//! DADOS DESTINATÁRIO
	nextY = DrawDadosDestinatario(pdf, x+paddingLeft, nextY, objetoPostal.Destinatario, objetoPostal.Nacional)

	//! BARRA DE CODIGO DESTINATARIO
	nextY = DrawDestinatarioBarCode(pdf, x+paddingLeft, nextY, destinatarioBarcodeBase64String)

	//! SEPARADOR REMETENTE
	nextY = DrawSeparadorRemetente(pdf, x, nextY)

	//! DADOS REMETENTE
	DrawDadosRemetente(pdf, x+paddingLeft/2, nextY, remetente)
}

func DrawSmallLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, idPlp int, remetente types.Remetente, objetoPostal types.ObjetoPostal, local bool) {
	pesoObjeto := objetoPostal.Peso
	codRastreio := FormatTrackingCode(objetoPostal.NumeroEtiqueta)
	codServicoPostagem := objetoPostal.CodigoServicoPostagem
	tipoServicoImagem := findTipoServicoImagemByCodServicoPostagem(codServicoPostagem)
	fmt.Println("     - Tipo serviço postagem etiqueta", tipoServicoImagem)

	dataMatrixBase64String := objetoPostal.Base64.Datamatrix
	barcodeBase64String := objetoPostal.Base64.Code
	destinatarioBarcodeBase64String := objetoPostal.Base64.CepBarcode

	paddingTop := 2.0 + 4.75
	paddingLeft := 5.0 + 3.5
	var nextY = y + paddingTop

	DrawSmallDelimiter(pdf, x, y)
	//! LOGO FAVOR, DATAMATRIX, TIPO SERVIÇO LOGO E numero PLP
	nextY = DrawSmallFirstRow(pdf, x+paddingLeft, nextY, idPlp, tipoServicoImagem, dataMatrixBase64String, local)
	//! PEDIDO, NF E PESO
	nextY = DrawSmallSecondRow(pdf, x+paddingLeft, nextY, pesoObjeto)
	//! CÓDIGO DE RASTREIO
	nextY = DrawTrackingCode(pdf, x, nextY, codRastreio)

	//! BARRA DE CÓDIGO
	nextY = DrawBarcode(pdf, x, nextY, barcodeBase64String)
	//! RECEBEDOR, ASSINATURA e DOCUMENTO
	nextY = DrawSmallRecebedorAssinaturaDocumentoLines(pdf, x+paddingLeft, nextY)

	//! SEPARADOR DESTINATÁRIO E LOGO CORREIOS
	nextY = DrawSmallDestinatarioCorreiosLogoDivisor(pdf, x+3.5, nextY, local)

	//! DADOS DESTINATÁRIO
	nextY = DrawDadosDestinatario(pdf, x+paddingLeft, nextY, objetoPostal.Destinatario, objetoPostal.Nacional)

	//! BARRA DE CODIGO DESTINATARIO
	nextY = DrawDestinatarioBarCode(pdf, x+paddingLeft, nextY, destinatarioBarcodeBase64String)

	//! SEPARADOR REMETENTE
	nextY = DrawSmallSeparadorRemetente(pdf, x, nextY)

	//! DADOS REMETENTE
	DrawDadosRemetente(pdf, x+paddingLeft/2, nextY, remetente)
}
