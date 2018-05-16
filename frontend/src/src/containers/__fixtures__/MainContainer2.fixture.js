import MainContainer from "../MainContainer";
import React  from 'react';

export default {
  component: MainContainer,
  reduxState: {
    box: {
      version: "version",
      title: ""
    }
  },
  fetch: [
    {
      matcher: '/api/version',
      response: {versionNumber: "versionNumber*", timestamp: "2018-03-19T10:32:40+01:00"}
    }
  ],
  props: {
    children: (
      <div>
        <p>No Title</p>
      </div>
    ),
  }
};
