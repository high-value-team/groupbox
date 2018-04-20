import thunk from 'redux-thunk';
import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import { devToolsEnhancer } from 'redux-devtools-extension';
import box from './src/Redux'


export default function(initialState) {
  const store = createStore(combineReducers({ box }), initialState, compose(applyMiddleware(thunk), devToolsEnhancer()));
  return store;
}
