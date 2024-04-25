package main

import (
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func main() {
	objetoPostalCarta := types.SolicitarEtiquetasPDFObjetoPostal{
		IdPrePostagem:         "PRVH7npdHXQIm9dWsExqky4A",
		CodigoServicoPostagem: "80160",
		CodigoRastreio:        "",
		Destinatario: types.SolicitarEtiquetaDestinatario{
			NomeDestinatario:        "Cliente Final",
			TelefoneDestinatario:    ptrToString("11987654321"),
			LogradouroDestinatario:  "Logradouro Teste",
			ComplementoDestinatario: ptrToString("Apto 1234"),
			NumeroDestinatario:      "321",
			CpfCnpjDestinatario:     ptrToString("123.456.789-00"),
			BairroDestinatario:      "Bairro Teste",
			CidadeDestinatario:      "Cidade do Cliente",
			UfDestinatario:          "CL",
			CepDestinatario:         "13560520",
		},
		DimensaoObjeto: types.DimensaoObjeto{
			TipoObjeto:          "Envelope",
			DimensaoAltura:      0.0,
			DimensaoLargura:     0.0,
			DimensaoComprimento: 0.0,
			DimensaoDiametro:    0.0,
		},
		Peso:             500,
		DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
	}

	solicitarEtiquetasPDF := types.SolicitarEtiquetasPDF{
		Remetente: types.SolicitarEtiquetaRemetente{
			NomeRemetente:        "Empresa Exemplo",
			LogradouroRemetente:  "Logradouro Teste",
			NumeroRemetente:      "123",
			ComplementoRemetente: ptrToString("Apto 101"),
			BairroRemetente:      "Bairro Exemplo",
			CepRemetente:         "12282220",
			CidadeRemetente:      "Cidade Exemplo",
			UfRemetente:          "EX",
			TelefoneRemetente:    ptrToString("11987654321"),
			CpfCnpjRemetente:     "12345678000199",
		},
		ObjetosPostais: []types.SolicitarEtiquetasPDFObjetoPostal{objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta, objetoPostalCarta},
	}

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("ERRO AO GERAR ETIQUETA: %s", errMsg)
		panic(errMsg)
	}
}

func ptrToString(s string) *string {
	return &s
}
