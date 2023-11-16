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
			NomeRemetente:           "Carioca Star",
			CepRemetente:            "13560648",
			NumeroRemetente:         "1762",
			LogradouroRemetente:     "Rua Conde do Pinhal",
			ComplementoRemetente:    "Loja",
			BairroRemetente:         "Jardim São Carlos",
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
				CodigoServicoPostagem: "03220",
				Peso:                  20,
				Destinatario: types.Destinatario{
					NomeDestinatario:        "Reda Ali Gharib",
					LogradouroDestinatario:  "Rua Passos",
					NumeroEndDestinatario:   "249",
					ComplementoDestinatario: "APTO - 132 Bloco A",
				},
				Nacional: types.Nacional{
					CepDestinatario:    "03058010",
					BairroDestinatario: "Belenzinho",
					CidadeDestinatario: "São Paulo",
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
		},
	}

	err := helpers.GenerateLabelsPDFLocal(correiosLog)
	if err != nil {
		panic(err)
	}
}
