import axios from 'axios';
import config from '../_environment';

const {baseURL} = config;


const Server = axios.create({
  baseURL,
  validateStatus: function (status) {
    return status < 400;
  },
});

export default Server;
