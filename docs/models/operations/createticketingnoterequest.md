# CreateTicketingNoteRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `TicketingNote`                                               | [*shared.TicketingNote](../../models/shared/ticketingnote.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `ConnectionID`                                                | *string*                                                      | :heavy_check_mark:                                            | ID of the connection                                          |
| `Fields`                                                      | []*string*                                                    | :heavy_minus_sign:                                            | Comma-delimited fields to return                              |
| `TicketID`                                                    | *string*                                                      | :heavy_check_mark:                                            | ID of the ticket                                              |