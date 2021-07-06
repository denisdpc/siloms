package siloms

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

// Requisicao contempla campos de interesse da requisicao
type Requisicao struct {
	numero     string
	partNumber string
	status     string
}

// IsReqNacionalizado verifica se a requisição é de material nacionalizado
func IsReqNacionalizado(r Requisicao) bool {
	return strings.HasPrefix(r.partNumber, "DCN")
}

// IsReqPendente verifica se a requisição não atingiu o status de mapa
func IsReqPendente(r Requisicao) bool {

	return false
}

// IsReqAtendida verifica se a requisição atingiu o status de mapa e posteriores
func IsReqAtendida(r Requisicao) bool {
	return false
}

// FiltrarRequisicoes extrai requisições que atende a determinado requisito
func FiltrarRequisicoes(v func(Requisicao) bool) []Requisicao {
	return nil
}
