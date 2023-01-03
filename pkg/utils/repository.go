package utils

func GenerateIDFromTableAndCurrencyIDs(tableID, currencyID string) string {
	return tableID + ":" + currencyID
}
