import {combineReducers} from 'redux';
import {routerReducer} from 'react-router-redux';
import global from './_reducer';


export default combineReducers({
  routing: routerReducer,
  global,
})