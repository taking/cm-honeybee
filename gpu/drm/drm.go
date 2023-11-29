package drm

import (
	"github.com/NeowayLabs/drm"
	"github.com/jollaman999/utils/logger"
	"strconv"
)

type DRM struct {
	DriverName        string `json:"driver_name"`
	DriverVersion     string `json:"driver_version"`
	DriverDate        string `json:"driver_date"`
	DriverDescription string `json:"driver_description"`
}

func GetDRMInfo() (DRM, error) {
	v, err := drm.Available()
	if err != nil {
		logger.Println(logger.DEBUG, true, "DRM: DRM is not available.")
		return DRM{}, nil
	}

	d := DRM{
		DriverName:        v.Name,
		DriverVersion:     strconv.Itoa(int(v.Major)) + "." + strconv.Itoa(int(v.Minor)) + "." + strconv.Itoa(int(v.Patch)),
		DriverDate:        v.Date,
		DriverDescription: v.Desc,
	}

	return d, nil
}
