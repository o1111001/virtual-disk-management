const http = require('../common/http');

const Client = url => {

  const client = new http.Client(url);

  return {
    listMachines: () => client.listMachines('/disks'),
    addDisk: (diskId, machineId) => client.addDisk('/disks', { diskId, machineId })
  }
};

module.exports = { Client };
