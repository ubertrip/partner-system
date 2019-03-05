import axios from 'axios';
import config from '../_environment';

const {baseURL} = config;

const Server = axios.create({
  baseURL,
  validateStatus: function (status) {
    return status < 500;
  },
});

axios.defaults.withCredentials = true;

export default Server;
