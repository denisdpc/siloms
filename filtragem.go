// Package siloms inclui funções para filtragem de requisições
package siloms

var statusAquisicaoPendente = map[string]bool{
	"Anulada":                  false,
	"Cancelada":                false,
	"Item Deserto":             true,
	"Suspensa temporariamente": true,
	"Aguardando validação":     true,
	"Validada":                 true,
	"Análise do Pedido":        true,
	"Empenhada":                false,
}

// IsAquisicaoPendente verifica necessidade de ação interna para aquisição
func IsAquisicaoPendente(r Requisicao) bool {
	return statusAquisicaoPendente[r.Status]
	return false
}

// IsReqAtendida verifica se a requisição atingiu o status de mapa e posteriores
func IsReqAtendida(r Requisicao) bool {
	return !statusAquisicaoPendente[r.Status]
	return false
}

// IsDeserto verifica se a requisição apresenta status de item deserto
func IsDeserto(r Requisicao) bool {
	return r.Status == "Item Deserto"
}

//IsMapa verifica se a requisição está em mapa gerado ou aprovado
func IsMapa(r Requisicao) bool {
	return r.Status == "Mapa Gerado" || r.Status == "Mapa Aprovado"
}

// IsPreMapa verifica se a requisição não atingiu o status de mapa
func IsPreMapa(r Requisicao) bool {

	return false
}

// FiltrarRequisicoes extrai requisições que atende a determinado requisito
func FiltrarRequisicoes(reqs []Requisicao, f func(Requisicao) bool) []Requisicao {
	var requisicoes []Requisicao
	for _, r := range reqs {
		if f(r) {
			requisicoes = append(requisicoes, r)
		}

	}
	return requisicoes
}
