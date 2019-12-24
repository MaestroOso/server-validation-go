package entities

type DomainDataModel struct {
  Host string
}

type HistoryModel struct {
  Items []DomainDataModel
}
