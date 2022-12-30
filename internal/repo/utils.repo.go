package repo

func generateIDFromTableAndCurrencyIDs(tableID, currencyID string) string {
	return tableID + ":" + currencyID
}
