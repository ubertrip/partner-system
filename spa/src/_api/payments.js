import Server from './server';

export default class {
  static getPaymentsByStatementUUID(statementUUID) {
    return Server.get(`credit/${statementUUID}`);
  }

  static getStatements() {
    return Server.get(`statements`);
  }

  static getDriverPayments(statementUUID, driverUUID) {
    return Server.get(`credit/${statementUUID}/${driverUUID}`);
  }

  static addPayment(statementUuid, driverUUID, credit, extra, gas = 0, petrol = 0) {
    return Server.post(`credit/${driverUUID}`, {
      statementUuid,
      credit: parseFloat(credit),
      extra,
      gas: parseFloat(gas),
      petrol: parseFloat(petrol),
    });
  }

  static getDriverByID(id) {
    return Server.get(`drivers/${id}`);
  }

  // static getLoginForm (login, password){
  //   return Server.get(`driver/${login}/${password}`);
  // }

  static getUserByLogin (login){
    return Server.get(`driver/${login}`);
  }

  static auth(login, password) {
    return Server.post(`login`, {
      login, password
    });
  }

  static logout(){
    return Server.get(`logout`);
  }


}