package disks

import (
	"database/sql"
)

type Machine struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	CpuCount          int64   `json:"cpuCount"`
	TotalDiskSpace 		int64 	`json:"totalDiskSpace"`
}

type MachineDisk struct {
	MachineId         int64   `json:"machineId"`
	DiskId            int64   `json:"diskId"`
}


type Store struct {
	Db *sql.DB
}
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListMachines() ([]*Machine, error) {
	rows, err := s.Db.Query(`
		SELECT 
			m."id",
			m."name",
			m."cpu" as "cpuCount",
			sum(d."space") as "totalDiskSpace"
		FROM machines m
		JOIN machines_disks md on md."machineId" = m."id"
		JOIN disks d on d."id" = md."diskId"
		GROUP BY m."id"
		ORDER BY m."id"
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Machine
	for rows.Next() {
		var machine Machine
		if err := rows.Scan(&machine.Id, &machine.Name, &machine.CpuCount, &machine.TotalDiskSpace); err != nil {
			return nil, err
		}

		res = append(res, &machine)
	}
	if res == nil {
		res = make([]*Machine, 0)
	}
	return res, nil
}

func (s *Store) MachineConnect(machine int64, disk int64) (MachineDisk, error) {
	rows, err := s.Db.Query(`
		INSERT INTO machines_disks ("machineId", "diskId") VALUES ($1, $2) RETURNING "machineId", "diskId"
	`, machine, disk)
	var machineDisk MachineDisk

	rows.Next()
	rows.Scan(&machineDisk.MachineId, &machineDisk.DiskId)
	return machineDisk, err
}
