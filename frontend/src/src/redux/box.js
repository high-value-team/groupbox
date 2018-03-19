import { browserHistory } from 'react-router';
import * as api from '../Api';

//
// action constants
//

const ADD_BOX = 'ADD_BOX';
const UPDATE_TITLE = 'UPDATE_TITLE';
const UPDATE_OWNER_EMAIL = 'UPDATE_OWNER_EMAIL';
const UPDATE_MEMBER_EMAILS = 'UPDATE_MEMBER_EMAILS';
const LOAD_VERSION = 'LOAD_VERSION';

const LOADING_BOX = 'LOADING_BOX';
const LOADING_BOX_SUCCESS = 'LOADING_BOX_SUCCESS';
const LOADING_BOX_ERROR = 'LOADING_BOX_ERROR';

// const SAVING_BOX = 'SAVING_BOX';
// const SAVING_BOX_SUCCESS = 'SAVING_BOX_SUCCESS';
// const SAVING_BOX_ERROR = 'SAVING_BOX_ERROR';

const SAVING_ITEM = 'SAVING_ITEM';
const SAVING_ITEM_SUCCESS = 'SAVING_ITEM_SUCCESS';
const SAVING_ITEM_ERROR = 'SAVING_ITEM_ERROR';

const UPDATING_ITEM_ERROR = 'UPDATING_ITEM_ERROR';
const DELETING_ITEM_ERROR = 'DELETING_ITEM_ERROR';

//
// actions
//

function loadingBox () {
  return {
    type: LOADING_BOX,
    status: 1,
  };
}

function loadingBoxSuccess (box) {
  return {
    type: LOADING_BOX_SUCCESS,
    status: 2,
    box,
  };
}

function loadingBoxError (error) {
  return {
    type: LOADING_BOX_ERROR,
    status: 3,
    error,
  };
}

function updatingItemError (error) {
  return {
    type: UPDATING_ITEM_ERROR,
    itemStatus: 3,
    error,
  };
}


function savingItem () {
  return {
    type: SAVING_ITEM,
    itemStatus: 1,
  };
}

function savingItemSuccess () {
  return {
    type: SAVING_ITEM_SUCCESS,
    itemStatus: 2,
    box,
  };
}

function savingItemError (error) {
  return {
    type: SAVING_ITEM_ERROR,
    itemStatus: 3,
    error,
  };
}

function deletingItemError (error) {
  return {
    type: DELETING_ITEM_ERROR,
    itemStatus: 3,
    error,
  };
}

function addBox (box) {
  return {
    type: ADD_BOX,
    box,
  };
}

//
// action creators
//

export function updateTitle (title) {
  return function(dispatch) {
    dispatch({
      type: UPDATE_TITLE,
      title,
    });
  };
}

export function updateOwnerEmail (ownerEmail) {
  return function(dispatch) {
    dispatch({
      type: UPDATE_OWNER_EMAIL,
      ownerEmail,
    });
  };
}

export function updateMemberEmailsTextfield (memberEmailsTextfield) {
  return function(dispatch) {
    dispatch({
      type: UPDATE_MEMBER_EMAILS,
      memberEmailsTextfield,
    });
  };
}

export function loadBox (boxkey) {
  return function (dispatch) {
    dispatch(loadingBox());
    api.loadBox(boxkey).then((box) => {
      dispatch(loadingBoxSuccess(box));
    })
      .catch((err) => {
        dispatch(loadingBoxError(err));
        console.warn('Error in loadBox:', err);
      });
  };
}

export function saveBox (box) {
  return function (dispatch) {
    api.saveBox(box)
      .then((newBox) => {
        dispatch(addBox(newBox));
        browserHistory.push(`/${newBox.key}`);
      })
      .catch((err) => {
        console.warn('Error in boxFanout', err);
      });
  };
}

export function saveItem ( boxkey, item ){
  return function (dispatch) {
    dispatch(savingItem());
    api.saveItem(boxkey, item)
      .then(() => {
        dispatch(savingItemSuccess());
        loadBox(boxkey)(dispatch);
      })
      .catch((err) => {
        dispatch(savingItemError(err));
        console.warn('Error in saveItem', err);
      });
  };
}

export function updateItem (boxkey, item){
  return function (dispatch) {
    api.updateItem(boxkey, item)
      .then(() => {
        // console.log(`updateItem:${boxkey}`);
        loadBox(boxkey)(dispatch);
      })
      .catch((err) => {
        dispatch(updatingItemError(err));
        console.warn('Error in updateItem', err);
      });
  };
}

export function deleteItem (boxkey, item){
  return function (dispatch) {
    api.deleteItem(boxkey, item)
      .then(() => {
        loadBox(boxkey)(dispatch);
      })
      .catch((err) => {
        dispatch(deletingItemError(err));
        console.warn('Error in deleteItem', err);
      });
  };
}

export function loadVersion () {
  return function(dispatch) {
    api.getVersion()
      .then((version) => {
        dispatch({
          type: LOAD_VERSION,
          version: version,
        });
      })
      .catch((err) => {
        console.warn('Error in loadVersion:', err);
      });
  };
}

//
// reducers
//

const initialState = {
  version: '',
  error: '',

  title: '',
  ownerEmail: '',
  memberEmailsTextfield: '',
  memberEmails: [],

  // status:
  // 0: none
  // 1: loading
  // 2: loaded - ok
  // 3: loaded - error
  status: 0,
  itemStatus: 0,

  key: '',
  memberNickname: '',
  creationDate: '',
  items: [],
  // items: [
  //   {
  //     itemID: '0',
  //     authorNickname: '',
  //     creationDate: '',
  //     subject: '',
  //     message: '',
  //   }
  // ],

};

export default function box (state = initialState, action) {
  switch (action.type) {
    case UPDATE_TITLE :
      return {
        ...state,
        title: action.title,
      };
    case UPDATE_OWNER_EMAIL :
      return {
        ...state,
        ownerEmail: action.ownerEmail
      };
    case UPDATE_MEMBER_EMAILS :
      return {
        ...state,
        memberEmailsTextfield: action.memberEmailsTextfield,
      };
    case ADD_BOX :
      return {
        ...state,
        ...action.box,
        error: '',
      };
    case LOAD_VERSION :
      return {
        ...state,
        version: action.version,
      };
    case LOADING_BOX :
      return {
        ...state,
        status: action.status,
      };
    case LOADING_BOX_SUCCESS :
      return {
        ...state,
        ...action.box,
        status: action.status,
        error: '',
      };
    case LOADING_BOX_ERROR :
      return {
        ...state,
        status: action.status,
        error: action.error,
      };
    case SAVING_ITEM :
      return {
        ...state,
        itemStatus: action.status,
      };
    case SAVING_ITEM_SUCCESS :
      return {
        ...state,
        itemStatus: action.status,
      };
    case SAVING_ITEM_ERROR :
      return {
        ...state,
        itemStatus: action.status,
        error: action.error,
      };
    case UPDATING_ITEM_ERROR :
      return {
        ...state,
        itemStatus: action.status,
        error: action.error,
      };
    case DELETING_ITEM_ERROR :
      return {
        ...state,
        itemStatus: action.status,
        error: action.error,
      };
    default :
      return state;
  }
}
