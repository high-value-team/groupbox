import React from 'react';
import ReactDOM from 'react-dom';
import thunk from 'redux-thunk';
import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import { routerReducer, syncHistoryWithStore } from 'react-router-redux';
import { browserHistory } from 'react-router';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Theme from './Theme';
import getRoutes from './Routes';
import Reboot from 'material-ui/Reboot';

import box from './Redux';

const store = createStore(combineReducers({box, routing: routerReducer}), compose(
  applyMiddleware(thunk),
  window.devToolsExtension ? window.devToolsExtension() : (f) => f
));

const history = syncHistoryWithStore(browserHistory, store);

const element = (
  <Provider store={store}>
    <MuiThemeProvider theme={Theme}>
      <Reboot>
        {getRoutes(history)}
      </Reboot>
    </MuiThemeProvider>
  </Provider>
);
const target = document.getElementById('root');

ReactDOM.render(element, target);



