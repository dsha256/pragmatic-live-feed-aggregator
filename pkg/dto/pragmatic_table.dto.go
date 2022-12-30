package dto

type PragmaticTable struct {
	TotalSeatedPlayers int `json:"totalSeatedPlayers"`
	Last20Results      []struct {
		Time               string `json:"time"`
		Result             int    `json:"result"`
		Color              string `json:"color"`
		GameId             string `json:"gameId"`
		PowerUpList        []any  `json:"powerUpList"`
		PowerUpMultipliers []any  `json:"powerUpMultipliers"`
	} `json:"last20Results"`
	TableId                   string `json:"tableId"`
	TableName                 string `json:"tableName"`
	NewTable                  bool   `json:"newTable"`
	LanguageSpecificTableInfo string `json:"languageSpecificTableInfo"`
	TableImage                string `json:"tableImage"`
	TableLimits               struct {
		Ranges     []float64 `json:"ranges"`
		MinBet     float64   `json:"minBet"`
		MaxBet     float64   `json:"maxBet"`
		MaxPlayers int       `json:"maxPlayers"`
	} `json:"tableLimits"`
	Dealer struct {
		Name string `json:"name"`
	} `json:"dealer"`
	TableOpen    bool   `json:"tableOpen"`
	TableType    string `json:"tableType"`
	TableSubtype string `json:"tableSubtype"`
	Currency     string `json:"currency"`
}

type PragmaticTableWithID struct {
	// tID = 100; cID = 200 => TableAndCurrencyID = "100:200"
	TableAndCurrencyID string         `json:"tableAndCurrencyID,omitempty"`
	PragmaticTable     PragmaticTable `json:"pragmaticTable"`
}

//func (pt PragmaticTable) Marshal() ([]byte, error) {
//	jsonPT, err := json.Marshal(pt)
//	if err != nil {
//		return nil, err
//	}
//	return jsonPT, nil
//}
//
//func (pt PragmaticTable)  {
//
//}

//// CreateTableCurrencyID returns a unique ID based on PragmaticTable's TableId and Currency
//// Generated ID => TableId + Currency
//func (pt PragmaticTable) CreateTableCurrencyID() string {
//	return pt.TableId + ":" + pt.Currency
//}
