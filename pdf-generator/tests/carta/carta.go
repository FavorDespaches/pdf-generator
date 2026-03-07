package main

import (
	"fmt"
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func main() {
	solicitarEtiquetasPDF := types.SolicitarEtiquetasPDF{
		Remetente: types.SolicitarEtiquetaRemetente{
			NomeRemetente:        "Empresa Exemplo",
			LogradouroRemetente:  "Rua Exemplo",
			NumeroRemetente:      "123",
			ComplementoRemetente: ptrToString("Apto 101"),
			BairroRemetente:      "Bairro Exemplo",
			CepRemetente:         "12282220",
			CidadeRemetente:      "Cidade Exemplo",
			UfRemetente:          "EX",
			TelefoneRemetente:    ptrToString("11987654321"),
			CpfCnpjRemetente:     "12345678000199",
		},
		ObjetosPostais: []types.SolicitarEtiquetasPDFObjetoPostal{
			{
				IdPrePostagem:         "ID123456789",
				CodigoServicoPostagem: "80160",
				CodigoRastreio:        "BR123456789XX",
				Destinatario: types.SolicitarEtiquetaDestinatario{
					NomeDestinatario:        "Cliente Final",
					TelefoneDestinatario:    ptrToString("11987654321"),
					LogradouroDestinatario:  "Rua do Cliente",
					ComplementoDestinatario: ptrToString("Casa 2"),
					NumeroDestinatario:      "321",
					CpfCnpjDestinatario:     ptrToString("123.456.789-00"),
					BairroDestinatario:      "Bairro do Cliente",
					CidadeDestinatario:      "Cidade do Cliente",
					UfDestinatario:          "CL",
					CepDestinatario:         "13560520",
				},
				DimensaoObjeto: types.DimensaoObjeto{
					TipoObjeto:          "Envelope",
					DimensaoAltura:      0.0,
					DimensaoLargura:     16.0,
					DimensaoComprimento: 24.0,
					DimensaoDiametro:    0.0,
				},
				Peso:             20,
				DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789801600100037Residencial Eden 300050000000000000-00.000000-00.000000|",
			},
		},
	}

	fmt.Println("Starting Carta PDF generation test...")
	fmt.Println("Working directory: ", helpers.GetWorkingDir())

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF)
	if err != nil {
		log.Fatalf("ERRO AO GERAR PDF DE CARTA: %v", err)
	} else {
		fmt.Println("Carta PDF generated successfully at label.pdf")
	}
}

func ptrToString(s string) *string {
	return &s
}
