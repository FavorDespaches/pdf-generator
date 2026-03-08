package main

import (
	"fmt"
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func newCartaRegistrada(codigoRastreio string, nomeDestinatario string, cpfCnpj string, logradouro string, numero string, complemento string, bairro string, cidade string, uf string, cep string) types.SolicitarEtiquetasPDFObjetoPostal {
	return types.SolicitarEtiquetasPDFObjetoPostal{
		IdPrePostagem:         "9912559667",
		CodigoServicoPostagem: "80799",
		CodigoRastreio:        codigoRastreio,
		Destinatario: types.SolicitarEtiquetaDestinatario{
			NomeDestinatario:        nomeDestinatario,
			LogradouroDestinatario:  logradouro,
			ComplementoDestinatario: ptrToString(complemento),
			NumeroDestinatario:      numero,
			CpfCnpjDestinatario:     ptrToString(cpfCnpj),
			BairroDestinatario:      bairro,
			CidadeDestinatario:      cidade,
			UfDestinatario:          uf,
			CepDestinatario:         cep,
		},
		DimensaoObjeto: types.DimensaoObjeto{
			TipoObjeto:          "Envelope",
			DimensaoAltura:      0.0,
			DimensaoLargura:     16.0,
			DimensaoComprimento: 24.0,
			DimensaoDiametro:    0.0,
		},
		Peso:             50,
		DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789807990100037Residencial Eden 300050000000000000-00.000000-00.000000|",
	}
}

func main() {
	solicitarEtiquetasPDF := types.SolicitarEtiquetasPDF{
		Remetente: types.SolicitarEtiquetaRemetente{
			NomeRemetente:        "Pitpet",
			LogradouroRemetente:  "Rua José Alvim",
			NumeroRemetente:      "18",
			ComplementoRemetente: ptrToString(""),
			BairroRemetente:      "Centro",
			CepRemetente:         "12940750",
			CidadeRemetente:      "Atibaia",
			UfRemetente:          "SP",
			CpfCnpjRemetente:     "32550349000193",
		},
		ObjetosPostais: []types.SolicitarEtiquetasPDFObjetoPostal{
			newCartaRegistrada("AA123456789BR", "Maria Eduarda da Silva Bernatzki", "16293947770", "Rua Real Grandeza", "96", "apartamento 401 bloco 2", "Botafogo", "Rio de Janeiro", "RJ", "22281034"),
			newCartaRegistrada("AA234567890BR", "ANA PAULA DE MOURA MOREIRA", "06465669625", "Rua dos Eucalíptos", "335", "", "Cidade Jardim", "Uberlândia", "MG", "38412144"),
			newCartaRegistrada("AA345678901BR", "Igor Rangel A Ferreira Silva", "70149234120", "Rua 7 Quadra 13 Lote 36", "N/A", "Condominio Imperial", "Trindade", "Trindade", "GO", "75385259"),
			newCartaRegistrada("AA456789012BR", "JULIANA BONG ALMEIDA DINIZ", "25252826803", "Rua Joao Batista Bianchini", "204", "casa", "Assunção", "São Bernardo Do Campo", "SP", "09861580"),
			newCartaRegistrada("AA567890123BR", "Andrea De Podestá G de Mattos", "92022375734", "Rua Pereira Nunes", "135", "903", "Vila Isabel", "Rio de Janeiro", "RJ", "20540134"),
		},
	}

	fmt.Println("Starting Carta Registrada PDF generation test...")
	fmt.Println("Working directory: ", helpers.GetWorkingDir())

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF, "carta-registrada.pdf")
	if err != nil {
		log.Fatalf("ERRO AO GERAR PDF DE CARTA REGISTRADA: %v", err)
	} else {
		fmt.Println("Carta Registrada PDF generated successfully at carta-registrada.pdf")
	}
}

func ptrToString(s string) *string {
	return &s
}
