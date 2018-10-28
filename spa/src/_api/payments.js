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
}