
package service

type CEPService interface {
    ResolveCity(cep string) (string, error)
}
