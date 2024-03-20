package main

import (
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
				},
				Peso:             500,
				DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
			},
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
				},
				Peso:             500,
				DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
			},
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
				},
				Peso:             500,
				DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
			},
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
				},
				Peso:             500,
				DatamatrixString: "13560520000371228222000037851AA123456789BB2501190000000123456789032980100037Residencial Eden 300050000000000000-00.000000-00.000000|",
			},
		},
	}

	err := helpers.GenerateLabelsPDFLocal(solicitarEtiquetasPDF)
	if err != nil {
		log.Fatalf("ERRO AO TRANSFORMAR PDF EM BASE64STRING")
		panic(err)
	}
}

func ptrToString(s string) *string {
	return &s
}
