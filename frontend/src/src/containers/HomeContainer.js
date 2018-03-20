import React from 'react';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import * as boxActionCreators from '../redux/box';
import PropTypes from 'prop-types';

import {withStyles} from 'material-ui/styles/index';
import Button from 'material-ui/Button';
import TextField from 'material-ui/TextField';
import Typography from 'material-ui/Typography';

const styles = theme => ({
  root: {
    fontFamily: 'Roboto, sans-serif',
    width: '100%',
  },
  container: {
  },
  title: {
    margin: theme.spacing.unit,
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: '300px'
  },
  textFieldBig: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: '90%',
  },
  button: {
    marginLeft: theme.spacing.unit,
    marginTop: theme.spacing.unit * 3,
  },
});


class HomeContainer extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
    updateTitle: PropTypes.func.isRequired,
    updateOwnerEmail: PropTypes.func.isRequired,
    updateMemberEmailsTextfield: PropTypes.func.isRequired,
    saveBox: PropTypes.func.isRequired,

    title: PropTypes.string.isRequired,
    ownerEmail: PropTypes.string.isRequired,
    memberEmailsTextfield: PropTypes.string.isRequired,
  };

  constructor(props) {
    super(props);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onSubmit () {
    const memberEmails = this.props.memberEmailsTextfield.split(/[\s,;]+/);
    const box = {
      title: this.props.title,
      ownerEmail: this.props.ownerEmail,
      memberEmails: memberEmails,
    };
    this.props.saveBox(box);
  }

  render () {
    const classes = this.props.classes;
    return (
      <div className={classes.root}>
        <form className={classes.container} noValidate={true} autoComplete="off">
          <Typography className={classes.title} type="display1" gutterBottom={true}>Groupbox anlegen</Typography>
          <div>
            <TextField
              id="title"
              label="Titel"
              className={classes.textField}
              value={this.props.title}
              onChange={(e) => this.props.updateTitle(e.target.value)}
              margin="normal"
              type="email"
            />
          </div>
          <div>
            <TextField
              id="ownerEmail"
              label="Deine Mailadresse"
              className={classes.textField}
              value={this.props.ownerEmail}
              onChange={(e) => this.props.updateOwnerEmail(e.target.value)}
              margin="normal"
              type="email"
            />
          </div>
          <div>
            <TextField
              id="memberEmails"
              label="Mitglieder Mailadressen"
              placeholder="Die Mailadressen von die Box-Mitglieder bitte hier eintragen."
              className={classes.textFieldBig}
              value={this.props.memberEmailsTextfield}
              onChange={(e) => this.props.updateMemberEmailsTextfield(e.target.value)}
              fullWidth={true}
              margin="normal"
            />
          </div>
          <div>
            <Button raised={true} color="primary" className={classes.button} onClick={this.onSubmit}>
              Los!
            </Button>
          </div>
        </form>
      </div>
    );
  }
}

function mapStateToProps (state) {
  return {...state.box};
}

function mapDispatchToProps (dispatch) {
  return bindActionCreators(boxActionCreators, dispatch);
}

export default withStyles(styles)(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(HomeContainer)
);
