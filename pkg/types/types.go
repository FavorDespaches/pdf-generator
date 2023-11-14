package types

type FechaPlpVariosServicosParams struct {
	XML            string   `json:"xml"`            // PLP (pré-lista de postagem) em formato XML.
	IdPlpCliente   int      `json:"idPlpCliente"`   // Código gerado pelo cliente para identificação da PLP. (Pode ser aleatório)
	CartaoPostagem string   `json:"cartaoPostagem"` // Número do Cartão de Postagem, que é o código que autoriza a postagem dos serviços no contrato.
	ListaEtiquetas []string `json:"listaEtiquetas"` // Lista de códigos de rastreios, sem o dígito verificador.
	Usuario        string   `json:"usuario"`        // Login do usuário para autenticação. (Fornecido pelo Representante Comercial dos Correios mediante carta de solicitação)
	Senha          string   `json:"senha"`          // Senha de autenticação. (Fornecida pelo Representante Comercial dos Correios mediante carta de solicitação)
}

type PLP struct {
	IdPlp               int     `json:"id_plp"`                // Identifica o registro da PLP
	ValorGlobal         *int    `json:"valorGlobal"`           // Valor em reais do total da tarificação dos objetos da PLP
	McuUnidadePostagem  *string `json:"mcu_unidade_postagem"`  // Identifica a unidade de postagem dos Correios.
	NomeUnidadePostagem *string `json:"nome_unidade_postagem"` // Nome da unidade de postagem dos Correios.
	CartaoPostagem      string  `json:"cartao_postagem"`       // Numero do cartão do cliente associado à PLP (Com 10 caracteres, incluindo os zeros à esquerda).
}

type Remetente struct {
	NumeroContrato          string  `json:"numero_contrato"`           // Número do contrato do cliente
	NumeroDiretoria         int     `json:"numero_diretoria"`          // Número da Diretoria Regional do contrato do cliente
	CodigoAdministrativo    string  `json:"codigo_administrativo"`     // Código administrativo do contrato
	NomeRemetente           string  `json:"nome_remetente"`            // Nome do remetente
	LogradouroRemetente     string  `json:"logradouro_remetente"`      // Logradouro do remetente
	NumeroRemetente         string  `json:"numero_remetente"`          // Número do endereço do remetente
	ComplementoRemetente    *string `json:"complemento_remetente"`     // Complemento do endereço do remetente
	BairroRemetente         string  `json:"bairro_remetente"`          // Bairro do remetente
	CepRemetente            string  `json:"cep_remetente"`             // CEP do remetente
	CidadeRemetente         string  `json:"cidade_remetente"`          // Cidade do remetente
	UfRemetente             string  `json:"uf_remetente"`              // Unidade de Federação
	TelefoneRemetente       *int    `json:"telefone_remetente"`        // Telefone do remetente
	FaxRemetente            *int    `json:"fax_remetente"`             // Fax do remetente
	EmailRemetente          *string `json:"email_remetente"`           // Email do remetente
	CelularRemetente        *int    `json:"celular_remetente"`         // Celular do remetente
	CpfCnpjRemetente        *int    `json:"cpf_cnpj_remetente"`        // CPF ou CNPJ do remetente
	CienciaConteudoProibido string  `json:"ciencia_conteudo_proibido"` // Coletar a ciência após mostrar a mensagem que o cliente está ciente do risco da postagem de objetos proibidos e coletar a ciência.
}

type ObjetoPostal struct {
	NumeroEtiqueta            string               `json:"numero_etiqueta"`             // Código do objeto a ser postado.
	Sscc                      *string              `json:"sscc"`                        // Serial Shipping Container Code (sscc)
	CodigoObjetoCliente       *string              `json:"codigo_objeto_cliente"`       // Código de controle do cliente
	CodigoServicoPostagem     string               `json:"codigo_servico_postagem"`     // Código do serviço a ser utilizado na postagem do objeto.
	Cubagem                   float64              `json:"cubagem"`                     // Cubagem do Objeto (em centímetros cúbicos)
	Peso                      float64              `json:"peso"`                        // Peso do objeto (em gramas)
	Rt1                       *string              `json:"rt1"`                         // Reservado para observação do cliente
	Rt2                       *string              `json:"rt2"`                         // Reservado para o SIGEP Web
	RestricaoAnac             string               `json:"restricao_anac"`              // Coletar o ciente do remetente para esta encomenda
	Destinatario              Destinatario         `json:"destinatario"`                // Dados do destinatário
	Nacional                  Nacional             `json:"nacional"`                    // Dados relevantes a postagem
	ServicoAdicional          ServicoAdicional     `json:"servico_adicional"`           // Identifica os serviços adicionais do objeto
	DimensaoObjeto            DimensaoObjeto       `json:"dimensao_objeto"`             // Deve conter as dimensões do objeto
	DataPostagemSara          *string              `json:"data_postagem_sara"`          // Data de efetivação da postagem
	StatusProcessamento       string               `json:"status_processamento"`        // STATUS do processamento do objeto
	NumeroComprovantePostagem *int                 `json:"numero_comprovante_postagem"` // Numero de comprovante de postagem
	ValorCobrado              *float64             `json:"valor_cobrado"`               // Valor que foi tarifado no Sistema de Atendimento dos Correios
	DeclaracaoConteudo        []DeclaracaoConteudo `json:"declaracao_conteudo"`         // Declaração do conteúdo sendo transportado
	Base64                    *Base64AuxParams     `json:"base64"`                      // Parâmetro auxiliar
}

type Base64AuxParams struct {
	Datamatrix    string `json:"datamatrix"`
	Code          string `json:"code"`
	CepBarcode    string `json:"cepBarcode"`
	LogoCorreios  string `json:"logo_correios"`
	LogoCliente   string `json:"logo_cliente"`
	TipoEncomenda string `json:"tipoEncomenda"`
}

type DeclaracaoConteudo struct {
	Conteudo      string  `json:"conteudo"`      // Nome do produto sendo transportado
	Quantidade    int     `json:"quantidade"`    // Quantidade do produto sendo transportado
	ValorUnitario float64 `json:"valorUnitario"` // Preço individual do produto
}

type Destinatario struct {
	NomeDestinatario        string  `json:"nome_destinatario"`        // Nome do destinatário
	TelefoneDestinatario    *int    `json:"telefone_destinatario"`    // Telefone do Destinatário
	CelularDestinatario     *int    `json:"celular_destinatario"`     // Celular do Destinatário
	EmailDestinatario       *string `json:"email_destinatario"`       // Email do Destinatário
	LogradouroDestinatario  string  `json:"logradouro_destinatario"`  // Logradouro do destinatário
	ComplementoDestinatario *string `json:"complemento_destinatario"` // Complemento do endereço
	NumeroEndDestinatario   string  `json:"numero_end_destinatario"`  // Parte do endereço
	CpfCnpjDestinatario     *int    `json:"cpf_cnpj_destinatario"`    // CPF ou CNPJ do Destinatário
}

type Nacional struct {
	BairroDestinatario  string   `json:"bairro_destinatario"`   // Bairro do destinatário
	CidadeDestinatario  string   `json:"cidade_destinatario"`   // Cidade do destinatário
	UfDestinatario      string   `json:"uf_destinatario"`       // Sigla da UF do destinatário
	CepDestinatario     string   `json:"cep_destinatario"`      // CEP do destinatário
	CodigoUsuarioPostal *string  `json:"codigo_usuario_postal"` // Código do usuário postal
	CentroCustoCliente  *string  `json:"centro_custo_cliente"`  // Centro de custo do cliente
	NumeroNotaFiscal    *int     `json:"numero_nota_fiscal"`    // Número da nota fiscal
	SerieNotaFiscal     *string  `json:"serie_nota_fiscal"`     // Série da nota fiscal
	ValorNotaFiscal     *float64 `json:"valor_nota_fiscal"`     // Valor da nota fiscal
	NaturezaNotaFiscal  *string  `json:"natureza_nota_fiscal"`  // Natureza da nota fiscal
	DescricaoObjeto     *string  `json:"descricao_objeto"`      // Descrição do objeto
	ValorACobrar        *float64 `json:"valor_a_cobrar"`        // Valor a cobrar do destinatário
}

type ServicoAdicional struct {
	CodigoServicoAdicional []string  `json:"codigo_servico_adicional"` // Código do serviço adicional
	ValorDeclarado         *float64  `json:"valor_declarado"`          // Valor do seguro adicional declarado pelo cliente
	EnderecoVizinho        *string   `json:"endereco_vizinho"`         // Endereço para a entrega no vizinho
	SiglaServicoAdicional  *[]string `json:"sigla_servico_adicional"`  // Parâmetro auxiliar para a criação da etiqueta (Não pertence ao XML dos correios)
}

type DimensaoObjeto struct {
	TipoObjeto          string   `json:"tipo_objeto"`          // Contém o código do tipo de objeto que foi postado (embalagem)
	DimensaoAltura      *float64 `json:"dimensao_altura"`      // Altura do objeto (em cm)
	DimensaoLargura     *float64 `json:"dimensao_largura"`     // Largura do objeto (em cm)
	DimensaoComprimento *float64 `json:"dimensao_comprimento"` // Comprimento do objeto (em cm)
	DimensaoDiametro    *float64 `json:"dimensao_diametro"`    // Diâmetro do objeto (em cm)
}

type CorreiosLog struct {
	TipoArquivo               string         `json:"tipo_arquivo"`                // Para este layout, deverá ser preenchido com a palavra 'Postagem'
	VersaoArquivo             string         `json:"versao_arquivo"`              // Identifica a versão do layout do arquivo XML. A versão deste layout é 2.3
	Plp                       PLP            `json:"plp"`                         // Pré-Lista de Postagem
	Remetente                 Remetente      `json:"remetente"`                   // Remetente
	FormaPagamento            *string        `json:"forma_pagamento"`             // Valor numérico indicando a forma de pagamento utilizada pelo cliente para realizar a postagem.
	ObjetoPostal              []ObjetoPostal `json:"objeto_postal"`               // Tag delimitadora do objeto que será postado. Esta tag contém as características do objeto.
	DataPostagemSara          *string        `json:"data_postagem_sara"`          // Deve conter a data de efetivação da postagem
	StatusProcessamento       string         `json:"status_processamento"`        // Contém o STATUS do processamento do objeto
	NumeroComprovantePostagem *int           `json:"numero_comprovante_postagem"` // Contém o numero de comprovante de postagem
	ValorCobrado              *float64       `json:"valor_cobrado"`               // Valor que foi tarifado no Sistema de Atendimento dos Correios
}
