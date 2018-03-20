import React  from 'react';
import HomeContainer from "./HomeContainer";

export default {
  component: HomeContainer,
  reduxState: {
    box: {
      title:"",
      ownerEmail: "",
      memberEmailsTextfield: ""
    }
  },
  fetch: [
    {
      matcher: '/api/boxes',
      method: 'POST',
      response: {boxKey:"a"},
    }
  ]
};

