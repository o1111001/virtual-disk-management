const axios = require('axios');

class Client {
  constructor(url) {
    this.url = url;
  }

  listMachines(){
    return axios.get(this.url);
  }

  addDisk(data){
    return axios.post(this.url, data);
  }
}

module.exports = { Client };
