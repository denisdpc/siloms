// Package siloms inclui funções para filtragem de requisições
package siloms

import (
	_ "fmt"
	"strings"
)

// IsDeserto verifica se a requisição apresenta status de item deserto
func IsDeserto(r Requisicao) bool {
	return r.status == "Item Deserto"
}

func IsMapa(r Requisicao) bool {
	return r.status == "Mapa aprovado" || r.status == "Mapa gerado"
}

// IsNacionalizado verifica se a requisição é de material nacionalizado
func IsNacionalizado(r Requisicao) bool {
	return strings.HasPrefix(r.partNumber, "DCN")
}

// IsPreMapa verifica se a requisição não atingiu o status de mapa
func IsPreMapa(r Requisicao) bool {

	return false
}

// IsReqAtendida verifica se a requisição atingiu o status de mapa e posteriores
func IsReqAtendida(r Requisicao) bool {
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
