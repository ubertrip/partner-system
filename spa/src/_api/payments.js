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

  static addPayment(statementUuid, driverUUID, credit) {
    return Server.post(`credit/${driverUUID}`, {statementUuid, credit: parseFloat(credit)});
  }

  static getDriverByID(id) {
    return Server.get(`drivers/${id}`);
  }
}