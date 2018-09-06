package repositories

import (
	"github.com/ubertrip/partner-system/models"
	"fmt"
)

func UpdateDriver(driver models.Driver) (ok bool) {
	_, err := Get().Exec("INSERT INTO drivers (uuid, id, name, status, photo) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE id=VALUES(id), name=VALUES(name), status=VALUES(status), photo=VALUES(photo)",
		driver.Uuid, driver.ID, driver.Name, driver.Status, driver.Photo)

	if err != nil {
		fmt.Println(err)
		ok = false
		return
	}

	ok = true

	return
}

func GetDriverById(driverID string) (driver models.Driver, err error) {
	err = Get().QueryRow("SELECT uuid, id, name, status, photo  FROM `drivers` WHERE id = ?", driverID).Scan(
		&driver.Uuid,
		&driver.ID,
		&driver.Name,
		&driver.Status,
		&driver.Photo)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetDriverByUUID(driverUUID string) (driver models.Driver, err error) {
	err = Get().QueryRow("SELECT uuid, id, name, status, photo  FROM `drivers` WHERE uuid = ?", driverUUID).Scan(
		&driver.Uuid,
		&driver.ID,
		&driver.Name,
		&driver.Status,
		&driver.Photo)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetDriversByStatus(status string) (drivers []models.Driver, err error) {
	rows, err := Get().Query("SELECT uuid, id, name, status, photo  FROM `drivers` WHERE status = ? ORDER BY name", status)

	if err != nil {
		return drivers, err
	}

	for rows.Next() {
		driver := models.Driver{}
		errRow := rows.Scan(
			&driver.Uuid,
			&driver.ID,
			&driver.Name,
			&driver.Status,
			&driver.Photo)

		if errRow != nil {
			continue
		}

		drivers = append(drivers, driver)
	}

	return drivers, err
}
