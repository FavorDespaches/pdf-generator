package main

import (
	"fmt"
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func newCarta(id string, nomeDestinatario string, logradouro string, numero string, bairro string, cidade string, uf string, cep string) types.SolicitarEtiquetasPDFObjetoPostal {
	return types.SolicitarEtiquetasPDFObjetoPostal{
		IdPrePostagem:         id,
		CodigoServicoPostagem: "80160",
		CodigoRastreio:        "BR123456789XX",
		Destinatario: types.SolicitarEtiquetaDestinatario{
			NomeDestinatario:        nomeDestinatario,
			TelefoneDestinatario:    ptrToString("11987654321"),
			LogradouroDestinatario:  logradouro,
			ComplementoDestinatario: ptrToString(""),
			NumeroDestinatario:      numero,
			CpfCnpjDestinatario:     ptrToString("123.456.789-00"),
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
		Peso:             20,
		DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789801600100037Residencial Eden 300050000000000000-00.000000-00.000000|",
	}
}

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
			newCarta("ID001", "Maria Silva", "Rua das Flores", "100", "Centro", "São Paulo", "SP", "01001000"),
			newCarta("ID002", "João Santos", "Av. Paulista", "200", "Bela Vista", "São Paulo", "SP", "01310100"),
			newCarta("ID003", "Ana Oliveira", "Rua Augusta", "300", "Consolação", "São Paulo", "SP", "01304001"),
			newCarta("ID004", "Pedro Costa", "Rua Oscar Freire", "400", "Jardins", "São Paulo", "SP", "01426001"),
			newCarta("ID005", "Lucas Souza", "Av. Brasil", "500", "Centro", "Rio de Janeiro", "RJ", "20040020"),
			newCarta("ID006", "Fernanda Lima", "Rua da Praia", "600", "Copacabana", "Rio de Janeiro", "RJ", "22041080"),
			newCarta("ID007", "Carlos Mendes", "Av. Atlântica", "700", "Leme", "Rio de Janeiro", "RJ", "22010000"),
			newCarta("ID008", "Juliana Rocha", "Rua XV de Novembro", "800", "Centro", "Curitiba", "PR", "80020310"),
			newCarta("ID009", "Roberto Alves", "Av. Beira Mar", "900", "Centro", "Florianópolis", "SC", "88015400"),
			newCarta("ID010", "Camila Ferreira", "Rua da Bahia", "1000", "Funcionários", "Belo Horizonte", "MG", "30160011"),
			newCarta("ID011", "Ricardo Gomes", "Av. Independência", "1100", "Centro", "Porto Alegre", "RS", "90035070"),
		},
	}

	fmt.Println("Starting Carta Simples PDF generation test...")
	fmt.Println("Working directory: ", helpers.GetWorkingDir())

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF, "carta-simples.pdf")
	if err != nil {
		log.Fatalf("ERRO AO GERAR PDF DE CARTA SIMPLES: %v", err)
	} else {
		fmt.Println("Carta Simples PDF generated successfully at carta-simples.pdf")
	}
}

func ptrToString(s string) *string {
	return &s
}
