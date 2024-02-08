package main

import (
	"github.com/FavorDespaches/pdf-generator/pkg/codes"
	"github.com/FavorDespaches/pdf-generator/pkg/helpers"
	"github.com/FavorDespaches/pdf-generator/pkg/types"
)

func main() {
	correiosLog := types.CorreiosLog{
		TipoArquivo:   "Postagem",
		VersaoArquivo: "2.3",
		Plp: types.PLP{
			IdPlp:          0,
			CartaoPostagem: "0076822516",
		},
		Remetente: types.Remetente{
			NomeRemetente:           "Gabriel Faglioni Mendes",
			CepRemetente:            "13560520",
			NumeroRemetente:         "37",
			LogradouroRemetente:     "Rua Luis Vaz de Toledo Piza",
			ComplementoRemetente:    "Apto 32",
			BairroRemetente:         "Jardim Lutfalla",
			CidadeRemetente:         "São Carlos",
			UfRemetente:             "SP",
			CienciaConteudoProibido: "S",
			CodigoAdministrativo:    "21431817",
			NumeroDiretoria:         74,
			NumeroContrato:          "9912559667",
		},
		FormaPagamento: "5",
		ObjetoPostal: []types.ObjetoPostal{
			{
				CodigoServicoPostagem: "03298",
				Peso:                  20,
				Destinatario: types.Destinatario{
					NomeDestinatario:        "Nelson Mendes Jr.",
					LogradouroDestinatario:  "Rua Teste",
					NumeroEndDestinatario:   "249",
					ComplementoDestinatario: "Clinica Docvalle",
				},
				Nacional: types.Nacional{
					CepDestinatario:    "12282220",
					BairroDestinatario: "Bairro Teste",
					CidadeDestinatario: "Caçapava",
					UfDestinatario:     "SP",
				},
				DimensaoObjeto: types.DimensaoObjeto{
					TipoObjeto:          "001",
					DimensaoAltura:      0.0,
					DimensaoComprimento: 20.0,
					DimensaoLargura:     20.0,
				},
				ServicoAdicional: types.ServicoAdicional{
					CodigoServicoAdicional: []string{"025", "001", "002", "019"},
					SiglaServicoAdicional:  []string{},
				},
				NumeroEtiqueta:      "TI708197798BR",
				Cubagem:             0,
				RestricaoAnac:       "S",
				StatusProcessamento: "0",
				Base64: types.Base64AuxParams{
					Datamatrix: codes.GetDatamatrix(),
					Code:       codes.GetCode(),
					CepBarcode: codes.GetCepBarcode(),
				},
			},
			{
				CodigoServicoPostagem: "03220",
				Peso:                  20,
				Destinatario: types.Destinatario{
					NomeDestinatario:        "Nelson Mendes Jr.",
					LogradouroDestinatario:  "Rua Teste",
					NumeroEndDestinatario:   "249",
					ComplementoDestinatario: "Clinica Docvalle",
				},
				Nacional: types.Nacional{
					CepDestinatario:    "12282220",
					BairroDestinatario: "Bairro Teste",
					CidadeDestinatario: "Caçapava",
					UfDestinatario:     "SP",
				},
				DimensaoObjeto: types.DimensaoObjeto{
					TipoObjeto:          "001",
					DimensaoAltura:      0.0,
					DimensaoComprimento: 20.0,
					DimensaoLargura:     20.0,
				},
				ServicoAdicional: types.ServicoAdicional{
					CodigoServicoAdicional: []string{"025", "002"},
					SiglaServicoAdicional:  []string{},
				},
				NumeroEtiqueta:      "TI708197798BR",
				Cubagem:             0,
				RestricaoAnac:       "S",
				StatusProcessamento: "0",
				Base64: types.Base64AuxParams{
					Datamatrix: codes.GetDatamatrix(),
					Code:       codes.GetCode(),
					CepBarcode: codes.GetCepBarcode(),
				},
			},
			{
				CodigoServicoPostagem: "04227",
				Peso:                  20,
				Destinatario: types.Destinatario{
					NomeDestinatario:        "Nelson Mendes Jr.",
					LogradouroDestinatario:  "Rua Teste",
					NumeroEndDestinatario:   "249",
					ComplementoDestinatario: "Clinica Docvalle",
				},
				Nacional: types.Nacional{
					CepDestinatario:    "12282220",
					BairroDestinatario: "Bairro Teste",
					CidadeDestinatario: "Caçapava",
					UfDestinatario:     "SP",
				},
				DimensaoObjeto: types.DimensaoObjeto{
					TipoObjeto:          "001",
					DimensaoAltura:      0.0,
					DimensaoComprimento: 20.0,
					DimensaoLargura:     20.0,
				},
				ServicoAdicional: types.ServicoAdicional{
					CodigoServicoAdicional: []string{"025"},
					SiglaServicoAdicional:  []string{},
				},
				NumeroEtiqueta:      "TI708197798BR",
				Cubagem:             0,
				RestricaoAnac:       "S",
				StatusProcessamento: "0",
				Base64: types.Base64AuxParams{
					Datamatrix: codes.GetDatamatrix(),
					Code:       codes.GetCode(),
					CepBarcode: codes.GetCepBarcode(),
				},
			},
			{
				CodigoServicoPostagem: "03158",
				Peso:                  20,
				Destinatario: types.Destinatario{
					NomeDestinatario:        "Nelson Mendes Jr.",
					LogradouroDestinatario:  "Rua Teste",
					NumeroEndDestinatario:   "249",
					ComplementoDestinatario: "Clinica Docvalle",
				},
				Nacional: types.Nacional{
					CepDestinatario:    "12282220",
					BairroDestinatario: "Bairro Teste",
					CidadeDestinatario: "Caçapava",
					UfDestinatario:     "SP",
				},
				DimensaoObjeto: types.DimensaoObjeto{
					TipoObjeto:          "001",
					DimensaoAltura:      0.0,
					DimensaoComprimento: 20.0,
					DimensaoLargura:     20.0,
				},
				ServicoAdicional: types.ServicoAdicional{
					CodigoServicoAdicional: []string{"025", "019"},
					SiglaServicoAdicional:  []string{},
				},
				NumeroEtiqueta:      "TI708197798BR",
				Cubagem:             0,
				RestricaoAnac:       "S",
				StatusProcessamento: "0",
				Base64: types.Base64AuxParams{
					Datamatrix: codes.GetDatamatrix(),
					Code:       codes.GetCode(),
					CepBarcode: codes.GetCepBarcode(),
				},
			},
		},
	}

	err := helpers.GenerateLabelsPDFLocal(correiosLog)
	if err != nil {
		panic(err)
	}
}
