import React  from 'react';
import BoxContainer from "./BoxContainer";

export default {
  component: BoxContainer,
  route: '/:boxkey',
  // url: '/my-box-key',
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
      matcher: '/api/boxes',
      method: 'POST',
      response: {boxKey:"a"},
    },
    {
      matcher: '/api/boxes/my-box-key',
      method: 'GET',
      response: {
        "title": "Klassiker der Weltliteratur",
        "memberNickname": "Golden Panda",
        "creationDate": "2017-10-01T10:30:59Z",
        "items": [
          {
            "itemID": "0",
            "authorNickname": "Golden Panda",
            "creationDate": "2017-10-01T10:35:20Z",
            "subject": "Die drei Muske...",
            "message": "Die drei Musketiere, Alexandre Dumas"
          },
          {
            "itemID": "1",
            "authorNickname": "Fierce Tiger",
            "creationDate": "2017-10-02T14:40:30Z",
            "subject": "Der Zauberer v...",
            "message": "Der Zauberer von Oz, Frank Baum"
          },
          {
            "itemID": "2",
            "authorNickname": "Flying Fox",
            "creationDate": "2017-10-03T20:55:10Z",
            "subject": "Schuld und Süh...",
            "message": "Schuld und Sühne, Dostojewski, www.amazon.de/Schuld-Sühne-Fjodr-Michailowitsch-Dostojewski-ebook/dp/B004UBCWK6"
          }
        ]
      }
    }
  ]
};
