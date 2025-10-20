package repository

type NetworkStruct struct{}

func NewNetworkRepository() *NetworkStruct {
	return &NetworkStruct{}
}

// func (r *NetworkStruct) FetchDomain(domain string) (models.Network, error) {
// 	var payload models.Network

// }
