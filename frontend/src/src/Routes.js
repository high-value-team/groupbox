import React from 'react';
import { Router, Route, IndexRoute } from 'react-router';
import MainContainer from './containers/MainContainer';
import BoxContainer from './containers/BoxContainer';
import HomeContainer from './containers/HomeContainer';

export default function getRoutes (history) {
  return (
    <Router history={history}>
      <Route path="/" component={MainContainer}>
        <Route path="/:boxkey" component={BoxContainer} />
        <IndexRoute component={HomeContainer}/>
      </Route>
    </Router>
  );
}
