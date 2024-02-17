package main

import (
	"github.com/FavorDespaches/pdf-generator/pkg/codes"
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
			TelefoneRemetente:    ptrToInt64(11987654321),
			CpfCnpjRemetente:     "12345678000199",
		},
		ObjetosPostais: []types.SolicitarEtiquetasPDFObjetoPostal{
			{
				IdPrePostagem:         "ID123456789",
				CodigoServicoPostagem: "03298",
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
					TipoObjeto:          "Caixa",
					DimensaoAltura:      10.0,
					DimensaoLargura:     20.0,
					DimensaoComprimento: 30.0,
					DimensaoDiametro:    0.0,
				},
				ServicoAdicional: &types.ServicoAdicional{
					CodigoServicoAdicional: []string{"001", "002"},
					ValorDeclarado:         1000.00,
					EnderecoVizinho:        "Rua do Vizinho, 123",
					SiglaServicoAdicional:  []string{"SA01", "SA02"},
				},
				Peso: 500,
				Base64: types.Base64Strings{
					Datamatrix: codes.GetDatamatrix(),
					Code:       codes.GetCode(),
					CepBarcode: codes.GetCepBarcode(),
				},
			},
		},
	}

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF)
	if err != nil {
		panic(err)
	}
}

func ptrToString(s string) *string {
	return &s
}

func ptrToInt64(i int64) *int64 {
	return &i
}
