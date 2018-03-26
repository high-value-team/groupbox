import React from 'react';
import { calculateParts }from '../TruncateText';

test('calculateParts: trim simple text', () => {
  let received = calculateParts({
    "originalText": "Die drei Musketiere, Alexandre Dumas",
    "URLPositions": [],
    "suffix": "…",
    "lengthWithSuffix": 28,
    "lengthNoSuffix": 27,
  }, 0);
  let expected = [{"URL": "", "isLast": true, "isURL": false, "text": "Die drei Musketiere, Alexan…"}];
  expect(received).toEqual(expected);
});

test('calculateParts: trim text with url', () => {
  let received = calculateParts({
    "originalText": "Schuld und Sühne, Dostojewski, www.amazon.de/Schuld-Sühne-Fjodr-Michailowitsch-Dostojewski-ebook/dp/B004UBCWK6",
    "URLPositions": [
      [
        31,
        53
      ]
    ],
    "suffix": "…",
    "lengthWithSuffix": 51,
    "lengthNoSuffix": 50,
  }, 0)
  let expected = [{"URL": "", "isLast": false, "isURL": false, "text": "Schuld und Sühne, Dostojewski, "}, {"URL": "www.amazon.de/Schuld-S", "isLast": true, "isURL": true, "text": "www.amazon.de/Schul…"}]
  expect(received).toEqual(expected);
});


test('calculateParts: short text with no trim', () => {
  let received = calculateParts({
    "originalText": "Hallo World",
    "URLPositions": [],
    "suffix": "…",
    "lengthWithSuffix": 11,
    "lengthNoSuffix": 11,
  }, 0);
  let expected = [{"URL": "", "isLast": true, "isURL": false, "text": "Hallo World"}];
  expect(received).toEqual(expected);
});




