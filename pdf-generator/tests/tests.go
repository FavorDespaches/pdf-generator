package main

import (
	"encoding/json"
	"log"

	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func main() {
	// solicitarEtiquetasPDF := types.SolicitarEtiquetasPDF{
	// 	Remetente: types.SolicitarEtiquetaRemetente{
	// 		NomeRemetente:        "Empresa Exemplo",
	// 		LogradouroRemetente:  "Rua Exemplo",
	// 		NumeroRemetente:      "123",
	// 		ComplementoRemetente: ptrToString("Apto 101"),
	// 		BairroRemetente:      "Bairro Exemplo",
	// 		CepRemetente:         "12282220",
	// 		CidadeRemetente:      "Cidade Exemplo",
	// 		UfRemetente:          "EX",
	// 		TelefoneRemetente:    ptrToString("11987654321"),
	// 		CpfCnpjRemetente:     "12345678000199",
	// 	},
	// 	ObjetosPostais: []types.SolicitarEtiquetasPDFObjetoPostal{
	// 		// {
	// 		// 	IdPrePostagem:         "ID123456789",
	// 		// 	CodigoServicoPostagem: "03298",
	// 		// 	CodigoRastreio:        "BR123456789XX",
	// 		// 	Destinatario: types.SolicitarEtiquetaDestinatario{
	// 		// 		NomeDestinatario:        "Cliente Final",
	// 		// 		TelefoneDestinatario:    ptrToString("11987654321"),
	// 		// 		LogradouroDestinatario:  "Rua do Cliente",
	// 		// 		ComplementoDestinatario: ptrToString("Casa 2"),
	// 		// 		NumeroDestinatario:      "321",
	// 		// 		CpfCnpjDestinatario:     ptrToString("123.456.789-00"),
	// 		// 		BairroDestinatario:      "Bairro do Cliente",
	// 		// 		CidadeDestinatario:      "Cidade do Cliente",
	// 		// 		UfDestinatario:          "CL",
	// 		// 		CepDestinatario:         "13560520",
	// 		// 	},
	// 		// 	DimensaoObjeto: types.DimensaoObjeto{
	// 		// 		TipoObjeto:          "Caixa",
	// 		// 		DimensaoAltura:      10.0,
	// 		// 		DimensaoLargura:     20.0,
	// 		// 		DimensaoComprimento: 30.0,
	// 		// 		DimensaoDiametro:    0.0,
	// 		// 	},
	// 		// 	ServicoAdicional: &types.ServicoAdicional{
	// 		// 		CodigoServicoAdicional: []string{"001", "002"},
	// 		// 		ValorDeclarado:         1000.00,
	// 		// 	},
	// 		// 	Peso:             500,
	// 		// 	DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
	// 		// },
	// 		{
	// 			IdPrePostagem:         "PRVH7npdHXQIm9dWsExqky4A",
	// 			CodigoServicoPostagem: "80160",
	// 			CodigoRastreio:        "",
	// 			Destinatario: types.SolicitarEtiquetaDestinatario{
	// 				NomeDestinatario:        "Cliente Final",
	// 				TelefoneDestinatario:    ptrToString("11987654321"),
	// 				LogradouroDestinatario:  "Rua do Cliente",
	// 				ComplementoDestinatario: ptrToString("Casa 2"),
	// 				NumeroDestinatario:      "321",
	// 				CpfCnpjDestinatario:     ptrToString("123.456.789-00"),
	// 				BairroDestinatario:      "Bairro do Cliente",
	// 				CidadeDestinatario:      "Cidade do Cliente",
	// 				UfDestinatario:          "CL",
	// 				CepDestinatario:         "13560520",
	// 			},
	// 			DimensaoObjeto: types.DimensaoObjeto{
	// 				TipoObjeto:          "Envelope",
	// 				DimensaoAltura:      0.0,
	// 				DimensaoLargura:     0.0,
	// 				DimensaoComprimento: 0.0,
	// 				DimensaoDiametro:    0.0,
	// 			},
	// 			Peso:             500,
	// 			DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
	// 		},
	// 	},
	// }
	jsonData := `{"remetente":{"nome_remetente":"Gabriel Faglioni Mendes","cep_remetente":"13560520","numero_remetente":"37","logradouro_remetente":"Rua Luiz Vaz de Toledo Piza","complemento_remetente":"Residencial Eden 3","bairro_remetente":"Jardim Lutfalla","cidade_remetente":"São Carlos","uf_remetente":"SP","cpf_cnpj_remetente":"45741156886"},"objetosPostais":[{"idPrePostagem":"PRRGEMooGySSyizvwl4zSzaQ","codigoRastreio":"OQDOJ5905","codigoServicoPostagem":"80160","destinatario":{"nome_destinatario":"Nelson Mendes Jr.","logradouro_destinatario":"Rua Job Vaz do Amaral","complemento_destinatario":"","numero_destinatario":"156","cpf_cnpj_destinatario":"70074159119","bairro_destinatario":"Jardim Lallo","cidade_destinatario":"São Paulo","uf_destinatario":"SP","cep_destinatario":"04812240","telefone_destinatario":""},"dimensaoObjeto":{"tipo_objeto":"001","dimensao_altura":0,"dimensao_largura":0,"dimensao_comprimento":0},"peso":20,"datamatrixString":"0481224000156135605200003790350000000000000PRRGEMooGySSyizvwl4zSzaQOQDOJ590543485925000182|"}]}`

	var solicitarEtiquetasPDF types.SolicitarEtiquetasPDF
	errJson := json.Unmarshal([]byte(jsonData), &solicitarEtiquetasPDF)
	if errJson != nil {
		log.Fatalf("Error unmarshalling JSON: %v", errJson)
	}

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("ERRO AO GERAR ETIQUETA: %s", errMsg)
		panic(errMsg)
	}
}

// func ptrToString(s string) *string {
// 	return &s
// }
