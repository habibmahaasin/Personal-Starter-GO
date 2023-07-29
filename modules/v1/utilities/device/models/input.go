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
	Header                  int     `json:"header"`
	Aerator_mode            int     `json:"aerator_mode"`
	Aerator_status          int     `json:"aerator_status"`
	Temperature             float64 `json:"temperature"`
	Ph                      float64 `json:"ph"`
	Dissolved_oxygen        float64 `json:"dissolved_oxygen"`
	Ph_calibration_firstval float64 `json:"value_ph_k"`
	Ph_calibration_secval   float64 `json:"value_ph_x"`
	Ph_adc                  int     `json:"ph_adc"`
	Device_id               string
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
	Brand_id        string `json:"brand_id" form:"brand_id" binding:"required"`
	Mode_id         string `json:"mode_id" form:"mode_id" binding:"required"`
}

type PhCalibration struct {
	Antares_id              string `json:"antares_id" form:"antares_id" binding:"required"`
	Ph_calibration_firstval string `json:"ph_calibration_firstval" form:"ph_calibration_firstval"`
	Ph_calibration_secval   string `json:"ph_calibration_secval" form:"ph_calibration_secval"`
}

type ControllingAPI struct {
	User_id   string `json:"user_id" form:"user_id" binding:"required"`
	Device_id string `json:"device_id" form:"device_id" binding:"required"`
	Mode      string `json:"mode" binding:"required"`
	Power     string `json:"power" binding:"required"`
}
