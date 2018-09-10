import Server from './server';

export default class {
  static getPaymentsByStatementUUID(statementUUID) {
    return Server.get(`credit/${statementUUID}`);
  }

  static getStatements() {
    return Server.get(`statements`);
  }
}