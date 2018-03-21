import React  from 'react';
import BoxContainer from "./BoxContainer";

export default {
  component: BoxContainer,
  route: '/:boxkey',
  props: {
    router: {
      params: {
        boxkey: "my-box-key"
      }
    },
  },
  reduxState: {
  },
  fetch: [
    {
      matcher: '/api/boxes/my-box-key',
      method: 'GET',
      response: 500,
    }
  ]
};
