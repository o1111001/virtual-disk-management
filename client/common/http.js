const axios = require('axios');

class Client {
  constructor(url) {
    this.url = url;
  }

  async listMachines(route){
    const res = await axios.get(`${this.url}${route}`);
    return res.data
  }

  async addDisk(route, data){
    const res = await axios.post(`${this.url}${route}`, data);
    return res.data;
  }
}

module.exports = { Client };
