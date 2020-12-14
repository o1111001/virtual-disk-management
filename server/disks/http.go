package disks

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/o1111001/virtual-disk-management/server/tools"
)

type HTTPHandlerFunc http.HandlerFunc

func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListMachines(store, rw)
		} else if r.Method == "POST" {
			handleMachineConnect(store, r, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListMachines(store *Store, rw http.ResponseWriter) {
	res, err := store.ListMachines()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleMachineConnect(store *Store, r *http.Request, rw http.ResponseWriter) {
	machineDisk := MachineDisk{}
	if err := json.NewDecoder(r.Body).Decode(&machineDisk); err != nil {
		log.Printf("machineDisk input is empty")
	}
	machineDisk, err := store.MachineConnect(machineDisk.MachineId, machineDisk.DiskId)
	if err == nil {
		tools.WriteJsonOk(rw, &machineDisk)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
