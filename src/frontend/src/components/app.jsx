import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';

import Topbar from './topbar';
import VersionService from '../services/version';

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
    this.state = { versionNumber: '' };
  }

  componentDidMount() {
    this.version = VersionService
      .subscribe(version => this.setState({versionNumber: version.versionNumber}), err => console.log(err));
  }

  componentWillUnmount() {
    this.version.unsubscribe();
  }

  render() {
    const { children, classes } = this.props;
    return (
      <div className={classes.root}>
        <Topbar version={this.state.versionNumber} />
        {children}
      </div>
    );
  }

}

export default withStyles(styles)(App);
