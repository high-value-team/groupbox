
// export const API_ROOT = 'https://us-central1-totalorder-4bafb.cloudfunctions.net/api';
// export const API_ROOT = 'https://totalorder-backend.cloud.dropstack.run';
// export const API_ROOT = ''; // local

export const NODE_ENV =  process.env.NODE_ENV;
export const API_ROOT =  process.env.REACT_APP_BACKEND_URL || '';
