package models

type RegisterUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type Received struct {
	First AntaresDetail `json:"m2m:cin"`
}

type AntaresDetail struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type ConnectionDat struct {
	Device_mode      int     `json:"aeratorMode"`
	Status_device    int     `json:"statusDevice"`
	Temperature      float64 `json:"temperature"`
	Ph               float64 `json:"ph"`
	Dissolved_oxygen float64 `json:"dissolvedOxygen"`
	Device_id        string
}

type ObjectAntares5 struct {
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con string `json:"con"`
}

type ObjectAntares4 struct {
	M2m_cin ObjectAntares5 `json:"m2m:cin"`
}

type ObjectAntares3 struct {
	M2m_rep ObjectAntares4 `json:"m2m:rep"`
	M2m_rss int            `json:"m2m:rss"`
}

type ObjectAntares2 struct {
	M2m_nev ObjectAntares3 `json:"m2m:nev"`
	M2m_sud bool           `json:"m2m:sud"`
	M2m_sur string         `json:"m2m:sur"`
}

type ObjectAntares1 struct {
	First ObjectAntares2 `json:"m2m:sgn"`
}

type DeviceInput struct {
	Device_name     string `json:"device_name" form:"device_name" binding:"required"`
	Antares_id      string `json:"antares_id" form:"antares_id" binding:"required"`
	Device_location string `json:"device_location" form:"device_location" binding:"required"`
	Latitude        string `json:"latitude" form:"latitude" binding:"required"`
	Longitude       string `json:"longitude" form:"longitude" binding:"required"`
	Brand           string `json:"brand" form:"brand" binding:"required"`
	Mode_id         string `json:"mode_id" form:"mode_id" binding:"required"`
}
