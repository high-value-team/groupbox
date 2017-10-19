import React from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import registerServiceWorker from './registerServiceWorker';

import 'typeface-roboto';
import 'material-design-icons/iconfont/material-icons.css';

import moment from 'moment';
import 'moment/locale/de';
moment.locale('de');

import Theme from './themes/standard';
import App from './components/app';
import HomePage from './components/home';
import BoxPage from './components/box';
import NotFoundPage from './components/notfound';

let page = null;
const path = window.location.pathname;
if (path === '/') {
  page = <HomePage />;
} else if (path === '/notfound') {
  page = <NotFoundPage />;
} else if (path.match(/\/\w+/)) {
  const boxkey = path.substr(1);
  page = <BoxPage boxkey={boxkey} />;
} else {
  page = <NotFoundPage />;
}

const element = (
  <MuiThemeProvider theme={Theme}>
    <App>
      {page}
    </App>
  </MuiThemeProvider>
);

const target = document.getElementById('root');

ReactDOM.render(element, target);
registerServiceWorker();
