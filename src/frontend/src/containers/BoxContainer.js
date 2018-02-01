import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import * as boxActionCreators from '../redux/box';
import moment from 'moment';

import { withStyles } from 'material-ui/styles';
import red from 'material-ui/colors/red';
import Avatar from 'material-ui/Avatar';
import Button from 'material-ui/Button';
import Card, { CardHeader, CardContent } from 'material-ui/Card';
import Dialog, { DialogActions, DialogContent, DialogTitle } from 'material-ui/Dialog';
import { LinearProgress } from 'material-ui/Progress';
import TextField from 'material-ui/TextField';
import Tooltip from 'material-ui/Tooltip';
import Typography from 'material-ui/Typography';
import IconButton from 'material-ui/IconButton';

import AddIcon from 'material-ui-icons/Add';
import Linkify from 'react-linkify';
import TextTruncate from 'react-text-truncate';
import ModeEditIcon from 'material-ui-icons/ModeEdit';
import ActionDeleteIcon from 'material-ui-icons/Delete';

const styles = theme => ({
  root: {
    width: '100%',
  },
  avatar: {
    backgroundColor: red[500],
    float: 'left'
  },
  buttonAdd: {
    margin: theme.spacing.unit,
    flip: false,
    position: 'absolute',
    top: 32,
    right: 32,
  },
  previewCard: {
    float: 'left',
    margin: '0.5em',
    width: '20em',
    height: '12em',
  },
  card: {
    float: 'left',
    margin: '0.5em',
    maxWidth: '20em',
  },
  dialog: {
    width: '30em',
  },
  dialogMessage: {
  },
  flexGrow: {
    flex: '1 1 auto',
  },
  greeting: {
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'nowrap',
    alignItems: 'flex-start ',
    width: '100%',
    height: '3em',
    marginTop: '1em', // top right bottom left
    marginBottom: theme.spacing.unit,
    marginLeft: theme.spacing.unit,
  },
  greetingName: {
    margin: theme.spacing.unit
  },
  greetingAvatar: {
  },
  greetingCount: {
    margin: theme.spacing.unit,
    marginLeft: '5em',
  },
  showCursor: {
    cursor: 'pointer'
  }
});

class BoxContainer extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,

    loadBox: PropTypes.func.isRequired,
    saveItem: PropTypes.func.isRequired,
    deleteItem: PropTypes.func.isRequired,
    updateItem: PropTypes.func.isRequired,

    boxkey: PropTypes.string.isRequired,
    box: PropTypes.object.isRequired,
  };

  constructor(props) {
    super(props);

    this.newItem = this.newItem.bind(this);

    this.state = {
      item: this.newItem(),
      openNewItemDialog: false,
      openShowItemDialog: false,
      openEditItemDialog: false,
    };

    this.hideDialog = this.hideDialog.bind(this);
    this.openEditItemDialog = this.openEditItemDialog.bind(this);
    this.openShowItemDialog = this.openShowItemDialog.bind(this);

    this.updateItemMessage = this.updateItemMessage.bind(this);
    this.getAuthorInitials = this.getAuthorInitials.bind(this);

    this.saveItem = this.saveItem.bind(this);
    this.deleteItem = this.deleteItem.bind(this);
    this.updateItem = this.updateItem.bind(this);

    this.getAuthorInitials = this.getAuthorInitials.bind(this);

    this.renderLoading = this.renderLoading.bind(this);
    this.renderBox = this.renderBox.bind(this);
    this.renderLoadedError = this.renderLoadedError.bind(this);
    this.renderBlank = this.renderBlank.bind(this);

    this.updateItemSubject = this.updateItemSubject.bind(this);
    this.updateItemMessage = this.updateItemMessage.bind(this);

  }

  componentDidMount() {
    this.props.loadBox(this.props.boxkey);
  }

  hideDialog() {
    this.setState({
      openNewItemDialog: false,
      openShowItemDialog: false,
      openEditItemDialog: false,
    });
  }

  newItem(ID, message, subject, creationDate, authorNickname) {
    return {
      itemID: '',
      message: message ? message : '',
      subject: subject ? subject : '',
      creationDate: creationDate ? creationDate : '',
      authorNickname: authorNickname ? authorNickname : '',
    };
  }

  openEditItemDialog() {
    this.hideDialog();
    this.setState({
      openEditItemDialog: true,
    });
  }

  openShowItemDialog(item) {
    this.setState({item, openShowItemDialog: true});
  }

  updateItemMessage(e) {
    const item = this.state.item;
    item.message = e.target.value;
    this.setState({item});
  }

  updateItemSubject(e) {
    const item = this.state.item;
    item.subject = e.target.value;
    this.setState({item});
  }

  saveItem() {
    this.props.saveItem(this.props.boxkey, this.state.item);
    this.hideDialog();
    this.setState({item: this.newItem()});
    // this.props.loadBox(this.props.boxkey);
  }

  deleteItem() {
    this.props.deleteItem(this.props.boxkey, this.state.item);
    this.hideDialog();
    this.setState({item: this.newItem()});
    // this.props.loadBox(this.props.boxkey);
  }

  updateItem() {
    // console.log(`updateItem:${JSON.stringify(this.state.item, null, 2)}`);
    this.props.updateItem(this.props.boxkey, this.state.item);
    this.hideDialog();
    this.setState({item: this.newItem()});
    // this.props.loadBox(this.props.boxkey);
  }

  getAuthorInitials (nickname) {
    return nickname
      .split(/\s+/)
      .map(word => word.substr(0, 1))
      .reduce((initials, letter) => `${initials}${letter}`, '');
  }

  render () {
    switch (this.props.box.status) {
      case 1:
        return this.renderLoading();
      case 2:
        return this.renderBox();
      case 3:
        return this.renderLoadedError();
      default:
        return this.renderBlank();
    }
  }

  renderLoading () {
    return (
      <div className={this.props.classes.root}>
        <LinearProgress />
      </div>
    );
  }

  renderLoadedError () {
    return (
      <div className={this.props.classes.root}>
        <p>Die Box kann leider nicht gefunden werden.</p>
      </div>
    );
  }

  renderBlank () {
    return <div className={this.props.classes.root} />;
  }

  renderBox () {

    const {classes, box} = this.props;
    const items = box.items ? box.items.sort((a, b) => b.creationDate > a.creationDate) : [];

    return (
      <div className={classes.root}>

        {/*AddButton*/}
        <Tooltip id="tooltip-icon" title="Neuen Eintrag anlegen">
          <Button
            className={classes.buttonAdd}
            fab={true}
            // color="accent"
            color="default"
            aria-label="Neuen Eintrag anlegen"
            onClick={() => this.setState({openNewItemDialog: true})}>
            <AddIcon/>
          </Button>
        </Tooltip>

        {/*Greeting*/}
        <div className={classes.greeting}>
          <div className={classes.greetingName}>
            <Typography type="title">
              Hallo, {box.memberNickname}!
            </Typography>
          </div>
          <div className={classes.greetingAvatar}>
            <Avatar aria-label="Author" className={classes.avatar}>
              {this.getAuthorInitials(box.memberNickname)}
            </Avatar>
          </div>
          <div className={classes.greetingCount}>
            <Typography type="subheading">
              Einträge: {box.items ? box.items.length : 0}
            </Typography>
          </div>
          <div className={classes.flexGrow}/>
        </div>

        {/*Items*/}
        {items.map((item, index) => (
          <Card className={classes.previewCard} key={`item-${index}`}>
            <CardHeader
              avatar={
                <Tooltip id="avatar-tooltip" title={item.authorNickname} placement="right">
                  <Avatar aria-label="Author" className={classes.avatar}>
                    {this.getAuthorInitials(item.authorNickname)}
                  </Avatar>
                </Tooltip>
              }
              title={item.subject}
              subheader={moment(item.creationDate).fromNow()}
              onClick={() => this.openShowItemDialog(item)}
              className={classes.showCursor}
            />
            <CardContent>
              {/*TODO: <div> cannot appear as a descendant of <p>.*/}
              <Typography component="p">
                <TextTruncate
                  line={3}
                  truncateText="…"
                  text={item.message}
                />
              </Typography>
            </CardContent>
          </Card>
        ))}

        {/*New Item Dialog*/}
        <Dialog open={this.state.openNewItemDialog} onClose={this.hideDialog}>
          <DialogTitle>Neuer Eintrag</DialogTitle>
          <DialogContent className={classes.dialog}>
            <TextField
              className={classes.dialogMessage}
              autoFocus={true}
              margin="dense"
              multiline={true}
              fullWidth={true}
              value={this.state.item.message}
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

        {/*Show Item Dialog*/}
        <Dialog open={this.state.openShowItemDialog} onClose={this.hideDialog}>
          <Card className={classes.card}>
            <CardHeader
              avatar={
                <Tooltip id="avatar-tooltip" title={this.state.item.authorNickname} placement="right">
                  <Avatar aria-label="Author" className={classes.avatar}>
                    {this.getAuthorInitials(this.state.item.authorNickname)}
                  </Avatar>
                </Tooltip>
              }
              title={this.state.item.subject}
              subheader={moment(this.state.item.creationDate).fromNow()}
              action={
                <div>
                  <IconButton onClick={this.deleteItem}>
                    <ActionDeleteIcon/>
                  </IconButton>
                  <IconButton onClick={this.openEditItemDialog}>
                    <ModeEditIcon/>
                  </IconButton>
                </div>
              }
            />
            <CardContent>
              <Typography component="p">
                <Linkify properties={{target: '_blank'}}>
                  {this.state.item.message}
                </Linkify>
              </Typography>
            </CardContent>
          </Card>
        </Dialog>

        {/*Edit Item Dialog*/}
        <Dialog open={this.state.openEditItemDialog} onClose={this.hideDialog}>
          <DialogContent className={classes.dialog}>
            <TextField
              className={classes.dialogMessage}
              autoFocus={true}
              margin="dense"
              multiline={false}
              fullWidth={true}
              value={this.state.item.subject}
              onChange={this.updateItemSubject}
            />
            <TextField
              className={classes.dialogMessage}
              autoFocus={true}
              margin="dense"
              multiline={true}
              fullWidth={true}
              value={this.state.item.message}
              onChange={this.updateItemMessage}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={this.hideDialog} color="primary">
              Abbrechen
            </Button>
            <Button onClick={this.updateItem} color="primary" raised={true}>
              Speichern
            </Button>
          </DialogActions>
        </Dialog>

      </div>
    );
  }
}

function mapStateToProps ({box}, {router}) {
  return {
    box: box,
    boxkey: router.params.boxkey,
  };
}

function mapDispatchToProps (dispatch) {
  return bindActionCreators(boxActionCreators, dispatch);
}

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(BoxContainer)
);
