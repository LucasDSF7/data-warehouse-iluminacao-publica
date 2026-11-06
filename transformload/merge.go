package transformload

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Extração de dados do sistema da Exati para table de staging genérica
// com todas colunas tipo TEXT. Necessário casting para types apropriádos.
func MergeDimPontoServico(dbpool *pgxpool.Pool) {
	query := `MERGE INTO dw.dim_ponto_servico t
	USING staging.ponto_servico s
	ON t.id_ponto_servico = s."ID Ponto de serviço"::integer
	WHEN MATCHED AND (
		t.plaqueta IS DISTINCT FROM COALESCE(s."Número de identificação", 'Não informado') OR
		t.endereco IS DISTINCT FROM COALESCE(s."Endereço", 'Não informado') OR
		t.latitude IS DISTINCT FROM s."Latitude"::float OR
		t.longitude IS DISTINCT FROM s."Longitude"::float OR
		t.bairro IS DISTINCT FROM COALESCE(s."Bairro - Atributo do ponto", 'Não informado') OR
		t.logradouro_eixo IS DISTINCT FROM COALESCE(s."Logradouro Eixo", 'Não informado') OR
		t.logradouro_google IS DISTINCT FROM COALESCE(s."Logradouro Google", 'Não informado') OR
		t.documento_evento IS DISTINCT FROM COALESCE(s."Documento Evento", 'Não preenchido') OR
		t.e_condominio IS DISTINCT FROM COALESCE(s."Condominio", 'Não') OR
		t.e_banco_credito IS DISTINCT FROM COALESCE(s."Banco de Crédito", 'Não') OR
		t.e_ponto_escuro IS DISTINCT FROM COALESCE(s."Ponto Escuro", 'Não') OR
		t.valor_creditos IS DISTINCT FROM COALESCE(REPLACE(REPLACE(s."Créditos", ',', '.'), '-', '0.0')::DECIMAL(5,2), 0.00) OR
		t.e_exclusividade_ip IS DISTINCT FROM COALESCE(s."Exclusividade IP", 'Não') OR
		t.altura_poste IS DISTINCT FROM COALESCE(s."Altura do Poste"::DECIMAL(4,1), 0.0) OR
		t.tipo_poste IS DISTINCT FROM COALESCE(s."Tipo do Poste", 'Não informado') OR
		t.e_telegestao IS DISTINCT FROM COALESCE(s."Telegestão", 'Não') OR
		t.marco_telegestao IS DISTINCT FROM COALESCE(s."Marco", 'Sem telegestão') OR
		t.largura_via_veiculos IS DISTINCT FROM COALESCE(s."Largura da via veículo"::DECIMAL(4,1), 0.0) OR
		t.largura_via_pedestres IS DISTINCT FROM COALESCE(s."Largura Via Pedestres"::DECIMAL(4,1), 0.0) OR
		t.distancia_entre_postes IS DISTINCT FROM COALESCE(s."Distância Entre Postes"::DECIMAL(4,1), 0.0) OR
		t.classificacao_via_veiculos IS DISTINCT FROM COALESCE(s."Classe de iluminação da via de veículos", 'Não se aplica') OR
		t.classificacao_via_pedestres IS DISTINCT FROM COALESCE(s."Classe de Iluminação da Via Pedestre", 'Não se aplica') OR
		t.idcomando IS DISTINCT FROM COALESCE(s."idComando", 'Não preenchido')
	) THEN
	 UPDATE SET
	 plaqueta = COALESCE(s."Número de identificação", t.plaqueta),
	 endereco = COALESCE(s."Endereço", t.endereco),
	 latitude = COALESCE(s."Latitude"::float, t.latitude),
	 longitude = COALESCE(s."Longitude"::float, t.longitude),
	 bairro = COALESCE(s."Bairro - Atributo do ponto", t.bairro),
	 logradouro_eixo = COALESCE(s."Logradouro Eixo", t.logradouro_eixo),
	 logradouro_google = COALESCE(s."Logradouro Google", t.logradouro_google),
	 documento_evento = COALESCE(s."Documento Evento", t.documento_evento),
	 e_condominio = COALESCE(s."Condominio", t.e_condominio),
	 e_banco_credito = COALESCE(s."Banco de Crédito", t.e_banco_credito),
	 e_ponto_escuro = COALESCE(s."Ponto Escuro", t.e_ponto_escuro),
	 valor_creditos = COALESCE(REPLACE(REPLACE(s."Créditos", ',', '.'), '-', '0.0')::DECIMAL(5,2), t.valor_creditos),
	 e_exclusividade_ip = COALESCE(s."Exclusividade IP", t.e_exclusividade_ip),
	 altura_poste = COALESCE(s."Altura do Poste"::DECIMAL(4,1), t.altura_poste),
	 tipo_poste = COALESCE(s."Tipo do Poste", t.tipo_poste),
	 e_telegestao = COALESCE(s."Telegestão", t.e_telegestao),
	 marco_telegestao = COALESCE(s."Marco", t.marco_telegestao),
	 largura_via_veiculos = COALESCE(s."Largura da via veículo"::DECIMAL(4,1), t.largura_via_veiculos),
	 largura_via_pedestres = COALESCE(s."Largura Via Pedestres"::DECIMAL(4,1), t.largura_via_pedestres),
	 distancia_entre_postes = COALESCE(s."Distância Entre Postes"::DECIMAL(4,1), t.distancia_entre_postes),
	 classificacao_via_veiculos = COALESCE(s."Classe de iluminação da via de veículos", t.classificacao_via_veiculos),
	 classificacao_via_pedestres = COALESCE(s."Classe de Iluminação da Via Pedestre", t.classificacao_via_pedestres),
	 idcomando = COALESCE(s."idComando", t.idcomando),
	 tipo_altura_poste = CASE
	 	WHEN COALESCE(s."Altura do Poste"::DECIMAL(4,1), 0.0) <= 12 THEN 'Baixo'
		ELSE 'Alto'
		END,
	 data_modificacao = CURRENT_DATE
	WHEN NOT MATCHED THEN
	INSERT (
		id_ponto_servico,
		plaqueta,
		endereco,
		latitude,
		longitude,
		cidade,
		bairro,
		logradouro_eixo,
		logradouro_google,
		documento_evento,
		e_condominio,
		e_banco_credito,
		e_ponto_escuro,
		valor_creditos,
		e_exclusividade_ip,
		altura_poste,
		tipo_altura_poste,
		tipo_poste,
		e_telegestao,
		marco_telegestao,
		largura_via_veiculos,
		largura_via_pedestres,
		distancia_entre_postes,
		classificacao_via_veiculos,
		classificacao_via_pedestres,
		idcomando
	)
	VALUES (
		s."ID Ponto de serviço"::integer,
		COALESCE(s."Número de identificação", 'Não informado'),
		COALESCE(s."Endereço", 'Não informado'),
		s."Latitude"::float,
		s."Longitude"::float,
		'Aracaju',
		COALESCE(s."Bairro - Atributo do ponto", 'Não informado'),
		COALESCE(s."Logradouro Eixo", 'Não informado'),
		COALESCE(s."Logradouro Google", 'Não informado'),
		COALESCE(s."Documento Evento", 'Não preenchido'),
		COALESCE(s."Condominio", 'Não'),
		COALESCE(s."Banco de Crédito", 'Não'),
		COALESCE(s."Ponto Escuro", 'Não'),
		COALESCE(REPLACE(REPLACE(s."Créditos", ',', '.'), '-', '0.0')::DECIMAL(5,2), 0.00),
		COALESCE(s."Exclusividade IP", 'Não'),
		COALESCE(s."Altura do Poste"::DECIMAL(4,1), 0.0),
		CASE 
            WHEN COALESCE(s."Altura do Poste"::DECIMAL(4,1), 0.0) <= 12 THEN 'Baixo'
            ELSE 'Alto'
        END,
		COALESCE(s."Tipo do Poste", 'Não informado'),
		COALESCE(s."Telegestão", 'Não'),
		COALESCE(s."Marco", 'Sem telegestão'),
		COALESCE(s."Largura da via veículo"::DECIMAL(4,1), 0.0),
		COALESCE(s."Largura Via Pedestres"::DECIMAL(4,1), 0.0),
		COALESCE(s."Distância Entre Postes"::DECIMAL(4,1), 0.0),
		COALESCE(s."Classe de iluminação da via de veículos", 'Não se aplica'),
		COALESCE(s."Classe de Iluminação da Via Pedestre", 'Não se aplica'),
		COALESCE(s."idComando", 'Não preenchido')
	);
	`
	_, err := dbpool.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
