import {combineReducers} from 'redux';
import {routerReducer} from 'react-router-redux';
import global from './_reducer';
import payments from './containers/Payments/reducer';
import driverPayments from './containers/Payment/reducer';


export default combineReducers({
  routing: routerReducer,
  global,
  payments,
  driverPayments,
});