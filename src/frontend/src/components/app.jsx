import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';

import Topbar from './topbar';

import { getVersion } from '../services/version';

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

  componentWillMount() {
    this.loadVersion();
  }

  loadVersion = () => {
    getVersion().then((version) => {
      const { versionNumber } = version;
      this.setState({ versionNumber });
    });
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
