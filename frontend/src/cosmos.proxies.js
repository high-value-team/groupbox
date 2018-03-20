import createFetchProxy from 'react-cosmos-fetch-proxy';
import createReduxProxy from 'react-cosmos-redux-proxy';
import createRouterProxy from 'react-cosmos-router-proxy';
import configureStore from './cosmos.configureStore';

//
// Redux
//

const ReduxProxy = createReduxProxy({
  createStore: state => configureStore(state)
});


//
// proxy order
//

export default [
  createFetchProxy(),
  ReduxProxy,
  createRouterProxy()
];
