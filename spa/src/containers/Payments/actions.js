import {
  PAYMENTS_LOAD,
  PAYMENTS_LOAD_STATEMENTS,
  PAYMENT_SET_STATEMENT_UUID,
} from './reducer'
import {toggleLoading} from '../../_reducer'
import {PaymentsApi} from '../../_api';
import {history} from '../../store';

export const loadPayments = (statementUUID) => (dispatch) => {
  dispatch(toggleLoading(true, 'Payments'));
  return PaymentsApi.getPaymentsByStatementUUID(statementUUID).then(({data}) => {
    if(data.status === 'ok') {
      dispatch({
        type: PAYMENTS_LOAD,
        list: data.result,
      });
    }

    dispatch(toggleLoading(false));
  })
};

export const loadStatements = () => (dispatch) => {
  dispatch(toggleLoading(true, 'Statements'));
  return PaymentsApi.getStatements().then(({data}) => {
    if(data.status === 'ok') {
      dispatch({
        type: PAYMENTS_LOAD_STATEMENTS,
        statements: data.result,
      });
    }else{
      console.log('redir2');
      history.push(`/login`)
    }
    dispatch(toggleLoading(false));
    return data.result;
  })
};

export const changeStatement = statementUUID => (dispatch, getState) => {
  return loadPayments(statementUUID)(dispatch).then(() => {
    dispatch(onChangeStatementUUID(statementUUID))
  })
};

export const onChangeStatementUUID = (statementUUID, driverUUID, toEdit) => (dispatch) => {
  dispatch({
    type: PAYMENT_SET_STATEMENT_UUID,
    statementUUID,
  });

  if(driverUUID) {
    console.log('redirect3');
    history.push(`/credit/${statementUUID}/${driverUUID}${toEdit ? '/add' : ''}`);
  }
};