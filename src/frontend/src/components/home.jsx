import React from 'react';
import PropTypes from 'prop-types';

import { withStyles } from 'material-ui/styles';
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

class Home extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
  }

  constructor(props) {
    super(props);
    this.state = {
      name: '',
      ownerEmail: '',
      memberEmails: '',
    };
  }

  doSubmit = () => {
    const memberEmails = this.state.memberEmails.split(/[\s,;]+/);
    //console.log('doSubmit', this.state.name, this.state.ownerEmail, memberEmails);

    const body = JSON.stringify({
      title: this.state.name,
      owner: this.state.ownerEmail,
      members: memberEmails,
    });

    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Content-Length', body.length);

    fetch('/api/boxes', {
      method: 'POST',
      headers,
      body,
    }).then(rsp => {
      if (rsp.ok) {
        try {
          rsp.json().then(newbox => {
            const boxkey = newbox.boxKey;
            if (boxkey) {
              window.location.href = `${window.location.href}${boxkey}`;
            } else {
              console.log('Error creating new box: missing box key');
            }
          });
        } catch (err) {
          console.log(`Error creating new box: ${err}`);
        }
      } else {
        console.log(`Error creating new box: ${rsp.status} - ${rsp.statusText}`);
      }
    });

  };

  updateName = e => this.setState({ name: e.target.value });
  updateOwnerEmail = e => this.setState({ ownerEmail: e.target.value });
  updateMemberEmails = e => this.setState({ memberEmails: e.target.value });

  render() {

    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <form className={classes.container} noValidate={true} autoComplete="off">
          <Typography className={classes.title} type="display1" gutterBottom={true}>Groupbox anlegen</Typography>
          <div>
            <TextField
              id="name"
              label="Titel"
              className={classes.textField}
              value={this.state.name}
              onChange={this.updateName}
              margin="normal"
            />
          </div>
          <div>
            <TextField
              id="ownerEmail"
              label="Deine Mailadresse"
              className={classes.textField}
              value={this.state.ownerEmail}
              onChange={this.updateOwnerEmail}
              margin="normal"
            />
          </div>
          <div>
            <TextField
              id="memberEmails"
              label="Mitglieder Mailadressen"
              placeholder="Die Mailadressen von die Box-Mitglieder bitte hier eintragen."
              className={classes.textFieldBig}
              value={this.state.memberEmails}
              onChange={this.updateMemberEmails}
              fullWidth={true}
              margin="normal"
            />
          </div>
          <div>
            <Button raised={true} color="primary" className={classes.button} onClick={this.doSubmit}>
              Los!
            </Button>
          </div>
        </form>
      </div>
    );

  }

}


export default withStyles(styles)(Home);
