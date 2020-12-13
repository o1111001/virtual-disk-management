const disks = require('./disks/client');
const client = disks.Client('http://localhost:8080');

const listMachines = async () => {
  try {
    const list = await client.listMachines();
    console.log(list);
  } catch (error) {
    console.log(`Problem listing available virtual machines: ${error.message || error}`);
  }
}

const addDisk = async () => {
  try {
    const data = await client.addDisk(5, 15);
    console.log(data);
  } catch (error) {
    console.log(`Problem adding disk: ${error.message || error}`);
  }
}

listMachines();
addDisk();
