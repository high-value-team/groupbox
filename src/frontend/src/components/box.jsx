import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

import BoxService from '../services/box';

import { withStyles } from 'material-ui/styles';
import { LinearProgress } from 'material-ui/Progress';

import Avatar from 'material-ui/Avatar';
import Button from 'material-ui/Button';
import Card, { CardHeader, CardContent, CardActions } from 'material-ui/Card';
import Dialog, { DialogActions, DialogContent, DialogTitle } from 'material-ui/Dialog';

import FavoriteIcon from 'material-ui-icons/Favorite';
import IconButton from 'material-ui/IconButton';
import TextField from 'material-ui/TextField';
import Tooltip from 'material-ui/Tooltip';
import Typography from 'material-ui/Typography';

import red from 'material-ui/colors/red';
import AddIcon from 'material-ui-icons/Add';


const styles = theme => ({
  root: {
    width: '100%',
  },
  avatar: {
    backgroundColor: red[500],
  },
  buttonAdd: {
    margin: theme.spacing.unit,
    flip: false,
    position: 'absolute',
    top: 32,
    right: 32,
  },
  card: {
    float: 'left',
    margin: '0.5em',
    maxWidth: 400,
  },
  dialog: {
    width: '30em',
  },
  dialogMessage: {
  },
  flexGrow: {
    flex: '1 1 auto',
  },
});

class Box extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
    boxkey: PropTypes.string,
  }

  constructor(props) {
    super(props);
    this.state = {
      status: 0, // 0: none, 1: loading, 2: loaded - ok, 3: loaded - error
      box: null,
      dialogOpen: false,
      itemMessage: '',
    };
  }

  componentDidMount() {
    this.loadBox();
  }

  componentWillUnmount() {
    BoxService.next(null);
  }


  hideDialog = () => {
    this.setState({ dialogOpen: false });
  }

  loadBox = () => {
    this.setState({ status: 1 });
    fetch(`/api/boxes/${this.props.boxkey}`).then(rsp => {
      if (rsp.ok) {
        rsp.json().then(box => {
          BoxService.next(box);
          this.setState({ box, status: 2 });
        });
      } else {
        this.setState({ status: 3 });
      }
    });
  }

  saveItem = () => {

    const { boxkey } = this.props;

    const body = JSON.stringify({
      message: this.state.itemMessage,
    });

    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Content-Length', body.length);

    fetch(`/api/boxes/${boxkey}/items`, {
      method: 'POST',
      headers,
      body,
    }).then(rsp => {
      if (rsp.ok) {
        window.location.href = `${window.location.href}${boxkey}`;
      } else {
        console.log(`Error creating new box: ${rsp.status} - ${rsp.statusText}`);
      }
    });

    this.setState({itemMessage: ''});

  }

  showDialog = () => {
    this.setState({ dialogOpen: true });
  }

  updateItemMessage = e => this.setState({itemMessage: e.target.value});

  renderBlank = () => {
    return <div className={this.props.classes.root} />;
  }

  renderDialog = () => {

    const { classes } = this.props;

    return (
      <Dialog open={this.state.dialogOpen} onRequestClose={this.hideDialog}>
        <DialogTitle>neuer Eintrag</DialogTitle>
        <DialogContent className={classes.dialog}>
          <TextField
            className={classes.dialogMessage}
            autoFocus={true}
            margin="dense"
            multiline={true}
            fullWidth={true}
            value={this.state.itemMessage}
            onChange={this.updateItemMessage}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={this.hideDialog} color="primary">
            Abbrechen
          </Button>
          <Button onClick={this.saveItem} color="primary" raised={true}>
            Hinzufügen
          </Button>
        </DialogActions>
      </Dialog>
    );

  }

  renderDialogButton = () => {
    const { classes } = this.props;
    return (
      <Tooltip id="tooltip-icon" title="neuer Eintrag anlegen">
        <Button
          className={classes.buttonAdd}
          fab={true}
          color="accent"
          aria-label="neuer Eintrag anlegen"
          onClick={this.showDialog} >
          <AddIcon />
        </Button>
      </Tooltip>
    );
  }

  renderLoading = () => {
    return (
      <div className={this.props.classes.root}>
        <LinearProgress />
      </div>
    );
  }

  renderLoadedError = () => {
    return (
      <div className={this.props.classes.root}>
        <p>Die Box kann leider nicht gefunden werden.</p>
      </div>
    );
  }

  renderItem = (index, item) => {
    const { classes } = this.props;

    const authorInitials = item.authorNickname
      .split(/\s+/)
      .map(word => word.substr(0, 1))
      .reduce((initials, letter) => `${initials}${letter}`, '');

    return (
      <Card className={classes.card} key={`item-${index}`} >
        <CardHeader
          avatar={
            <Tooltip id="avatar-tooltip" title={item.authorNickname} placement="right">
              <Avatar aria-label="Author" className={classes.avatar}>
                {authorInitials}
              </Avatar>
            </Tooltip>
          }
          title={item.subject}
          subheader={moment(item.creationDate).fromNow()}
        />
        <CardContent>
          <Typography component="p">
            {item.message}
          </Typography>
        </CardContent>
        <CardActions disableActionSpacing={true}>
          <Tooltip id="favorite-tooltip" title="Finde ich gut!" placement="right">
            <IconButton aria-label="Finde ich gut!">
              <FavoriteIcon />
            </IconButton>
          </Tooltip>
          <div className={classes.flexGrow} />
        </CardActions>
      </Card>
    );
  }

  renderLoadedOK = () => {

    const { classes } = this.props;
    const { box } = this.state;

    const items = box.items ? box.items.sort((a, b) => b.creationDate > a.creationDate) : [];
    const addButton = this.renderDialogButton();
    const dialog = this.renderDialog();

    return (
      <div className={classes.root}>
        {addButton}
        {items.map((item, index) => this.renderItem(index, item))}
        {dialog}
      </div>
    );

  }

  render = () => {

    switch (this.state.status) {
      case 1:
        return this.renderLoading();
      case 2:
        return this.renderLoadedOK();
      case 3:
        return this.renderLoadedError();
    }

    return this.renderBlank();

  }

}

export default withStyles(styles)(Box);
