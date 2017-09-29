package main

import (
	"context"
	"dummydb"
	"fmt"
)


func main() {
	dmdb := dummydb.DMDatabase(context.Background())
	
	imaginaryDocumentID := dmdb.AddQuery("INSERT INTO documents (date, amount, vat) VALUES('2013-02-10','123.23','32.2') RETURNING id")
	imaginaryParentItemID := dmdb.AddQuery(fmt.Sprintf("INSERT INTO documents_item (document_id, parent_id, quantity) VALUES('%s','-1','23') RETURNING id", imaginaryDocumentID))
	imaginarySubParentItemID := dmdb.AddQuery(fmt.Sprintf("INSERT INTO documents_item (document_id, parent_id, quantity) VALUES('%s','%s','23') RETURNING id", imaginaryDocumentID, imaginaryParentItemID))
	dmdb.AddQuery(fmt.Sprintf("INSERT INTO documents_item (document_id, parent_id, quantity) VALUES('%s','%s','23') RETURNING id", imaginaryDocumentID, imaginarySubParentItemID))
	dmdb.AddQuery(fmt.Sprintf("INSERT INTO transactions (document_id, account_id, amount) VALUES('%s','53','123.23') RETURNING id", imaginaryDocumentID))
	dmdb.AddQuery(fmt.Sprintf("INSERT INTO transactions (document_id, account_id, amount) VALUES('%s','54','-123.23') RETURNING id", imaginaryDocumentID))

	dmdb.ExecuteQueries()
}
