import React from 'react';
import PropTypes from 'prop-types';

import BoxService from '../services/box';
import VersionService from '../services/version';

import { withStyles } from 'material-ui/styles';
import Topbar from './topbar';

const styles = (/*theme*/) => ({
  root: {
  },
});

class App extends React.Component {

  static propTypes = {
    classes: PropTypes.object.isRequired,
    children: PropTypes.node.isRequired,
  }

  constructor(props) {
    super(props);
    this.state = { title: '', versionNumber: '' };
  }

  componentDidMount() {
    this.boxService = BoxService.subscribe(this.onNewBox);
    this.versionService = VersionService
      .subscribe(
        version => this.setState({versionNumber: version.versionNumber}, this.onNewBox),
        err => console.log(err)
      );
  }

  componentWillUnmount() {
    this.boxService.unsubscribe();
    this.versionService.unsubscribe();
  }

  onNewBox = box => {

    if (box) {
      const title = `${box.title} - Groupbox`;
      this.setState({title});
      window.history.replaceState({}, title, '/');
      document.title = title;
    } else {
      const title = this.state.versionNumber ? `Groupbox ${this.state.versionNumber}` : 'Groupbox';
      this.setState({title});
      document.title = title;
    }

  }

  render() {
    const { children, classes } = this.props;
    return (
      <div className={classes.root}>
        <Topbar title={this.state.title} />
        {children}
      </div>
    );
  }

}

export default withStyles(styles)(App);
