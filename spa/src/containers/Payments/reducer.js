export const PAYMENTS_LOAD = 'payments/load-list';
export const PAYMENTS_LOAD_STATEMENTS = 'payments/load-statements';
export const PAYMENT_SET_STATEMENT_UUID = 'payment/set-statement-uuid';

const initialState = {
  list: [],
  statements: [],
  statementUUID: ''
};

export default (state = initialState, action) => {
  switch (action.type) {

    case PAYMENTS_LOAD:
      return {
        ...state,
        list: [...action.list],
      };

    case PAYMENTS_LOAD_STATEMENTS:

      return {
        ...state,
        statements: [...action.statements],
      };

    case PAYMENT_SET_STATEMENT_UUID:

      return {
        ...state,
        statementUUID: action.statementUUID,
      };

    default:
      return state
  }
}