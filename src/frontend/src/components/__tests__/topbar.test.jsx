import React from 'react';
import ReactDOM from 'react-dom';
import Topbar from '../old/topbar';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<Topbar version="hello" />, div);
});
