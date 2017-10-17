import * as React from 'react';
import * as ReactDOM from 'react-dom';
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

const target: HTMLElement = document.getElementById('root') as HTMLElement;

ReactDOM.render(element, target);
registerServiceWorker();
