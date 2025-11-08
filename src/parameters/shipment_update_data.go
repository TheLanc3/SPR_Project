package parameters

import "spr-project/enums"

type ShipmentUpdateData struct {
	Id     int64
	Status enums.Status
}
