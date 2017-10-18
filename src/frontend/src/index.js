import React from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import registerServiceWorker from './registerServiceWorker';

import 'typeface-roboto';
import 'material-design-icons/iconfont/material-icons.css';

import Theme from './themes/standard';
import App from './components/app';
import Home from './components/home';

const element = (
  <MuiThemeProvider theme={Theme}>
    <App>
      <Home />
    </App>
  </MuiThemeProvider>
);

const target = document.getElementById('root');

ReactDOM.render(element, target);
registerServiceWorker();
