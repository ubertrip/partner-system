import {
  PAYMENTS_LOAD,
  PAYMENTS_LOAD_STATEMENTS,
  PAYMENT_SET_STATEMENT_UUID,
} from './reducer'

import {PaymentsApi} from '../../_api';

export const loadPayments = (statementUUID) => (dispatch) => {
  return PaymentsApi.getPaymentsByStatementUUID(statementUUID).then(({data}) => {
    if(data.status === 'ok') {
      dispatch({
        type: PAYMENTS_LOAD,
        list: data.result,
      });
    }
  })
};

export const loadStatements = () => (dispatch) => {
  return PaymentsApi.getStatements().then(({data}) => {
    if(data.status === 'ok') {
      dispatch({
        type: PAYMENTS_LOAD_STATEMENTS,
        statements: data.result,
      });
    }
  })
};

export const changeStatement = statementUUID => (dispatch, getState) => {
  return loadPayments(statementUUID)(dispatch).then(() => {
    dispatch(onChangeStatementUUID(statementUUID))
  })
};

export const onChangeStatementUUID = statementUUID => ({
  type: PAYMENT_SET_STATEMENT_UUID,
  statementUUID,
});