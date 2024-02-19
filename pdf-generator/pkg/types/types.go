package types

type ServicoAdicional struct {
	CodigoServicoAdicional []string `json:"codigo_servico_adicional"` // Código do serviço adicional
	ValorDeclarado         float64  `json:"valor_declarado"`          // Valor do seguro adicional declarado pelo cliente
	EnderecoVizinho        string   `json:"endereco_vizinho"`         // Endereço para a entrega no vizinho
	SiglaServicoAdicional  []string `json:"sigla_servico_adicional"`  // Parâmetro auxiliar para a criação da etiqueta (Não pertence ao XML dos correios)
}

type DimensaoObjeto struct {
	TipoObjeto          string  `json:"tipo_objeto"`          // Contém o código do tipo de objeto que foi postado (embalagem)
	DimensaoAltura      float64 `json:"dimensao_altura"`      // Altura do objeto (em cm)
	DimensaoLargura     float64 `json:"dimensao_largura"`     // Largura do objeto (em cm)
	DimensaoComprimento float64 `json:"dimensao_comprimento"` // Comprimento do objeto (em cm)
	DimensaoDiametro    float64 `json:"dimensao_diametro"`    // Diâmetro do objeto (em cm)
}

// SolicitarEtiquetaRemetente represents the sender's label request information.
type SolicitarEtiquetaRemetente struct {
	NomeRemetente        string  `json:"nome_remetente"`        // Nome do remetente
	LogradouroRemetente  string  `json:"logradouro_remetente"`  // Logradouro do remetente
	NumeroRemetente      string  `json:"numero_remetente"`      // Número do endereço do remetente
	ComplementoRemetente *string `json:"complemento_remetente"` // Complemento do endereço do remetente
	BairroRemetente      string  `json:"bairro_remetente"`      // Bairro do remetente
	CepRemetente         string  `json:"cep_remetente"`         // CEP do remetente
	CidadeRemetente      string  `json:"cidade_remetente"`      // Cidade do remetente
	UfRemetente          string  `json:"uf_remetente"`          // Unidade de Federação
	TelefoneRemetente    *string `json:"telefone_remetente"`    // Telefone do remetente
	CpfCnpjRemetente     string  `json:"cpf_cnpj_remetente"`    // CPF ou CNPJ do remetente
}

// SolicitarEtiquetaObjetoPostal represents the postal object label request information.
type SolicitarEtiquetaObjetoPostal struct {
	CodigoServicoPostagem string                                `json:"codigo_servico_postagem"` // Código do serviço a ser utilizado na postagem do objeto
	Peso                  int                                   `json:"peso"`                    // Peso do objeto (em gramas)
	Destinatario          SolicitarEtiquetaDestinatario         `json:"destinatario"`            // Destinatário information
	ServicoAdicional      *ServicoAdicional                     `json:"servico_adicional"`       // Serviço adicional information
	DimensaoObjeto        DimensaoObjeto                        `json:"dimensao_objeto"`         // Dimensão do objeto
	DeclaracaoConteudo    []SolicitarEtiquetaDeclaracaoConteudo `json:"declaracao_conteudo"`     // Declaração de conteúdo
}

// SolicitarEtiquetaDeclaracaoConteudo represents the content declaration for a label request.
type SolicitarEtiquetaDeclaracaoConteudo struct {
	Conteudo      string  `json:"conteudo"`       // Nome do produto sendo transportado
	Quantidade    int     `json:"quantidade"`     // Quantidade do produto sendo transportado
	ValorUnitario float64 `json:"valor_unitario"` // Preço individual do produto
}

// SolicitarEtiquetaDestinatario represents the recipient's information for a label request.
type SolicitarEtiquetaDestinatario struct {
	NomeDestinatario        string  `json:"nome_destinatario"`        // Nome do destinatário
	TelefoneDestinatario    *string `json:"telefone_destinatario"`    // Telefone do Destinatário
	LogradouroDestinatario  string  `json:"logradouro_destinatario"`  // Logradouro do destinatário
	ComplementoDestinatario *string `json:"complemento_destinatario"` // Complemento do endereço
	NumeroDestinatario      string  `json:"numero_destinatario"`      // Parte do endereço
	CpfCnpjDestinatario     *string `json:"cpf_cnpj_destinatario"`    // CPF ou CNPJ do Destinatário
	BairroDestinatario      string  `json:"bairro_destinatario"`      // Bairro do destinatário
	CidadeDestinatario      string  `json:"cidade_destinatario"`      // Cidade do destinatário
	UfDestinatario          string  `json:"uf_destinatario"`          // Sigla da UF do destinatário
	CepDestinatario         string  `json:"cep_destinatario"`         // CEP do destinatário
	NumeroNotaFiscal        *int64  `json:"numero_nota_fiscal"`       // Número da nota fiscal
}

// SolicitarDeclaracaoConteudo represents the content declaration request.
type SolicitarDeclaracaoConteudo struct {
	Remetente      SolicitarEtiquetaRemetente                `json:"remetente"`      // Remetente information
	ObjetosPostais []SolicitarDeclaracaoConteudoObjetoPostal `json:"objetosPostais"` // Postal objects information
}

// SolicitarDeclaracaoConteudoObjetoPostal represents the postal object for a content declaration request.
type SolicitarDeclaracaoConteudoObjetoPostal struct {
	Peso                float64                               `json:"peso"`                // Peso do objeto (em gramas)
	Destinatario        SolicitarEtiquetaDestinatario         `json:"destinatario"`        // Destinatário information
	DeclaracoesConteudo []SolicitarEtiquetaDeclaracaoConteudo `json:"declaracoesConteudo"` // Declarações de conteúdo
}

type SolicitarEtiquetasPDF struct {
	Remetente      SolicitarEtiquetaRemetente          `json:"remetente"`
	ObjetosPostais []SolicitarEtiquetasPDFObjetoPostal `json:"objetosPostais"`
}

type SolicitarEtiquetasPDFObjetoPostal struct {
	IdPrePostagem         string                        `json:"idPrePostagem"`
	CodigoServicoPostagem string                        `json:"codigoServicoPostagem"`
	CodigoRastreio        string                        `json:"codigoRastreio"`
	Destinatario          SolicitarEtiquetaDestinatario `json:"destinatario"`
	DimensaoObjeto        DimensaoObjeto                `json:"dimensaoObjeto"`
	ServicoAdicional      *ServicoAdicional             `json:"servicoAdicional"`
	Peso                  float64                       `json:"peso"`
	Base64                Base64Strings                 `json:"base64"`
}

type Base64Strings struct {
	Datamatrix string `json:"datamatrix"`
	Code       string `json:"code"`
	CepBarcode string `json:"cepBarcode"`
}
