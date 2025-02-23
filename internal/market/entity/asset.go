package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

// funcao construtora
func NewAsset(id, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
