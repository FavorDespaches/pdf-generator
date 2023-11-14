package helpers

import (
	"bytes"
	"encoding/base64"

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

func GenerateLabelsPDF(correiosLog types.CorreiosLog) (string, []string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("Arial", "", "ttf/Arial.ttf")
	pdf.AddUTF8Font("Arial", "B", "ttf/Arial_Bold.ttf")

	idPlp := correiosLog.Plp.IdPlp
	remetente := correiosLog.Remetente
	chunkifiedObjetoPostal := chunkifyObjetosPostais(correiosLog.ObjetoPostal, 4)

	var etiquetas []string
	for _, objetoPostalChunk := range chunkifiedObjetoPostal {
		pdf.AddPage()
		DrawDottedLines(pdf)

		subdivisionStartPoints := []struct{ x, y float64 }{
			{0, 0},
			{pageWidth / 2, 0},
			{0, pageHeight / 2},
			{pageWidth / 2, pageHeight / 2},
		}

		for i, startPoint := range subdivisionStartPoints {
			objetoPostal := objetoPostalChunk[i]
			etiquetas = append(etiquetas, objetoPostal.NumeroEtiqueta)
			DrawLabel(pdf, startPoint.x, startPoint.y, labelWidth, labelHeight, i, idPlp, remetente, objetoPostal)
		}
	}

	var buffer bytes.Buffer

	err := pdf.Output(&buffer)
	if err != nil {
		return "", []string{""}, err
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return base64Str, etiquetas, nil
}

func DrawLabel(pdf *gofpdf.Fpdf, x, y, width, height float64, index int, idPlp int, remetente types.Remetente, objetoPostal types.ObjetoPostal) {
	pesoObjeto := objetoPostal.Peso
	codRastreio := FormatTrackingCode(objetoPostal.NumeroEtiqueta)

	//dataMatrixBase64String := objetoPostal.Base64.Datamatrix
	//barcodeBase64String := objetoPostal.Base64.Code
	//destinatarioBarcodeBase64String := objetoPostal.Base64.CepBarcode

	paddingTop := 2.0
	paddingLeft := 5.0
	var nextY = y + paddingTop

	//! LOGO FAVOR, DATAMATRIX, TIPO SERVIÇO LOGO E numero PLP
	nextY = DrawFirstRow(pdf, x+paddingLeft, nextY, idPlp)

	//! PEDIDO, NF E PESO
	nextY = DrawSecondRow(pdf, x+paddingLeft, nextY, pesoObjeto)

	//! CÓDIGO DE RASTREIO
	nextY = DrawTrackingCode(pdf, x, nextY, codRastreio)

	//! BARRA DE CÓDIGO
	nextY = DrawBarcodePlaceholder(pdf, x, nextY)

	//! RECEBEDOR, ASSINATURA e DOCUMENTO
	nextY = DrawRecebedorAssinaturaDocumentoLines(pdf, x+paddingLeft, nextY)

	//! SEPARADOR DESTINATÁRIO E LOGO CORREIOS
	nextY = DrawDestinatarioCorreiosLogoDivisor(pdf, x, nextY)

	//! DADOS DESTINATÁRIO
	nextY = DrawDadosDestinatario(pdf, x+paddingLeft, nextY, objetoPostal.Destinatario, objetoPostal.Nacional)

	//! BARRA DE CODIGO DESTINATARIO
	nextY = DrawDestinatarioBarCodePlaceholder(pdf, x+paddingLeft, nextY)

	//! SEPARADOR REMETENTE
	nextY = DrawSeparadorRemetente(pdf, x, nextY)

	//! DADOS REMETENTE
	DrawDadosRemetente(pdf, x+paddingLeft/2, nextY, remetente)
}
