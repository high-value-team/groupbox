import {API_ROOT} from './Config';

export function saveBox(box) {
  const body = JSON.stringify({
    title: box.title,
    owner: box.ownerEmail,
    members: box.memberEmails,
  });
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');
  headers.append('Content-Length', body.length);

  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/boxes`, {
      method: 'POST',
      headers,
      body,
    }).then(resp => {
      if (resp.ok) {
        resp.json().then(respObj => {
          resolve({...box, key: respObj.boxKey});
        });
      } else {
        console.warn(`saveBox():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

export function loadBox(boxkey) {
  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/boxes/${boxkey}`, {
      method: 'GET',
    }).then(resp => {
      if (resp.ok) {
        resp.json().then(box => {
          resolve(box);
        });
      } else {
        console.warn(`loadBox():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

export function getVersion() {
  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/version`, {
      method: 'GET',
    }).then(resp => {
      if (resp.ok) {
        resp.text().then(version => {
          console.log(`version:${version}`);
          resolve(version);
        });
      } else {
        console.warn(`getVersion():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

export function saveItem(boxkey, item) {
  const body = JSON.stringify({message: item.message});
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');
  headers.append('Content-Length', body.length);

  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/boxes/${boxkey}/items`, {
      method: 'POST',
      headers,
      body,
    }).then(resp => {
      if (resp.ok) {
        resolve();
      } else {
        console.warn(`saveItem():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

export function updateItem(boxkey, item) {
  const body = JSON.stringify({subject: item.subject, message: item.message});
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');
  headers.append('Content-Length', body.length);

  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/boxes/${boxkey}/items/${item.itemID}`, {
      method: 'PUT',
      headers,
      body,
    }).then(resp => {
      if (resp.ok) {
        resolve();
      } else {
        console.warn(`updateItem():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

export function deleteItem(boxkey, item) {
  const body = JSON.stringify({});
  const headers = new Headers();
  headers.append('Content-Type', 'application/json');
  headers.append('Content-Length', body.length);

  return new Promise((resolve, reject) => {
    fetch(`${API_ROOT}/api/boxes/${boxkey}/items/${item.itemID}`, {
      method: 'DELETE',
      headers,
      body,
    }).then(resp => {
      if (resp.ok) {
        resolve();
      } else {
        console.warn(`deleteItem():${JSON.stringify(resp, null, 2)}`);
        reject(`API endpoint failed: ${resp.status} - ${resp.statusText}`);
      }
    }).catch(err => {
      reject(err);
    });
  });
}

