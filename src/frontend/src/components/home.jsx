import React from 'react';
import PropTypes from 'prop-types';

import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import Paper from 'material-ui/Paper';
import TextField from 'material-ui/TextField';

const styles = theme => ({
  root: {
    flexGrow: 1,
    marginTop: 30,
    fontFamily: 'Roboto, sans-serif',
  },
  button: {
    margin: theme.spacing.unit,
  },
  paper: {
    padding: 16,
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    width: 200,
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
    };
  }

  doSubmit = () => {
    console.log('doSubmit', this.state.name, this.state.ownerEmail);
  };

  updateName = e => this.setState({ name: e.target.value });
  updateOwnerEmail = e => this.setState({ ownerEmail: e.target.value });

  render() {

    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <Grid container={true} spacing={24} justify="center">

          <Grid item={true} xs={12}>
            <Paper className={classes.paper} elevation={0}>
              Groupbox anlegen
            </Paper>
          </Grid>

          <Grid item={true} xs={12}>
            <TextField
              id="name"
              label="Etikett"
              className={classes.textField}
              value={this.state.name}
              onChange={this.updateName}
              margin="normal"
            />
          </Grid>
          <Grid item={true} xs={12}>
            <TextField
              id="ownerEmail"
              label="Wer bist Du?"
              placeholder="Deine Mailadresse"
              className={classes.textField}
              value={this.state.ownerEmail}
              onChange={this.updateOwnerEmail}
              margin="normal"
            />
          </Grid>
          <Grid item={true} xs={12}>
            <Button raised={true} color="primary" className={classes.button} onClick={this.doSubmit}>
              Los!
            </Button>
          </Grid>

        </Grid>
      </div>
    );

  }

}


export default withStyles(styles)(Home);
